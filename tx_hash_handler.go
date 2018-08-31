package ethNotification

import (
	"context"
	"log"
	"strings"
)

type txHashHandler struct {
	engine *Engine
	hash   string
}

func NewTxHashHandler(e *Engine, hash string) txHashHandler {
	return txHashHandler{
		engine: e,
		hash:   hash,
	}
}

func (hdl txHashHandler) Handle() error {
	transaction, err := hdl.fetchTxInfo()
	if err != nil {
		log.Println("Transaction error: ", err.Error(), " txHash: ", hdl.hash)
		return err
	}

	if allowPush(transaction) {
		hdl.pushPendingTransaction(transaction)
	}
	return nil
}

func (hdl txHashHandler) fetchTxInfo() (*Transaction, error) {
	var result *Transaction

	if err := hdl.engine.c.CallContext(context.Background(), &result, "eth_getTransactionByHash", hdl.hash); err != nil {
		return nil, err
	}

	if result == nil {
		return nil, ErrTransactionNotFound
	}

	result.Status = Pending
	result.IsSeft = result.From == result.To
	result.ChainName = hdl.engine.ChainName

	if isERCTransaction(result) {
		hdl.fillTokenInfo(result)
	}

	if result.Value != "0x0" {
		bigInt, _ := ConvertHexStringToBigInt(result.Value)
		result.Value = bigInt.String()
	} else {
		result.Value = "0"
	}

	return result, nil
}

func (hdl txHashHandler) fillTokenInfo(tran *Transaction) {
	tokens := hdl.engine.TokenDataSource.FindTokens([]string{tran.To})
	if len(tokens) > 0 {
		token := tokens[0]
		inputData := ParseInputTx(tran.Input, token.Decimals)
		tran.To = inputData.ToAddress
		tran.Value = inputData.Value
		tran.ContractAddress = token.ContractAddress
		tran.TokenDecimal = int(token.Decimals)
		tran.TokenSymbol = token.Symbol
	}
}

func isETHTransaction(tran *Transaction) bool {
	return tran.Input == "0x"
}

func isERCTransaction(tran *Transaction) bool {
	return strings.Contains(tran.Input, MethodIDTransferERC20Token.String())
}

func allowPush(tran *Transaction) bool {
	return isETHTransaction(tran) || isERCTransaction(tran)
}

func (hdl txHashHandler) pushPendingTransaction(tran *Transaction) {

	walletSubsResult := hdl.engine.DataSource.FindWalletSubscribers([]Transaction{*tran})

	if len(walletSubsResult) > 0 {
		walletSubscribers := walletSubsResult[0].Subscribers
		hdl.engine.CacheData.Set(tran.Hash, walletSubscribers, *tran)

		go func() {
			for _, ws := range walletSubscribers {
				message := PushMessage{
					Title:        hdl.engine.pushTitle,
					Sound:        "default",
					Content:      hdl.engine.MessageHook.MessageTitle(tran, ws),
					Badge:        "1",
					DeviceTokens: []string{ws.DeviceToken},
					Payload:      hdl.engine.MessageHook.MessagePayload(tran, ws),
				}
				hdl.engine.MessageHook.BeforeSend(tran, ws, message)
				sendMessage(hdl.engine.pushKey, &message)
				hdl.engine.MessageHook.AfterSend(tran, ws, message)
			}
		}()
	}
}
