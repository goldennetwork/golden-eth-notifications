package ethNotification

import (
	"context"
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
		return err
	}

	if allowPush(transaction) {
		hdl.pushPendingTransaction(transaction)
	}
	return nil
}

func (hdl txHashHandler) fetchTxInfo() (*Transaction, error) {
	var result *Transaction

	if err := hdl.engine.cT.CallContext(context.Background(), &result, "eth_getTransactionByHash", hdl.hash); err != nil {
		return nil, err
	}

	if result == nil {
		return nil, ErrTransactionNotFound
	}

	result.Status = Pending
	result.IsSeft = result.From == result.To
	result.ChainName = hdl.engine.ChainName

	if isERCTransaction(result) {
		fillTokenInfo(hdl.engine.tokenDataSource, result)
	}

	if result.Value != "0x0" {
		bigInt, _ := ConvertHexStringToBigInt(result.Value)
		result.Value = bigInt.String()
	} else {
		result.Value = "0"
	}

	return result, nil
}

func (hdl txHashHandler) pushPendingTransaction(tran *Transaction) {

	walletSubsResult := hdl.engine.dataSource.FindWalletSubscribers([]Transaction{*tran})

	if len(walletSubsResult) > 0 {
		walletSubscribers := walletSubsResult[0].Subscribers
		hdl.engine.cacheData.Set(tran.Hash, walletSubscribers, *tran)
		go hdl.engine.pushMessage(tran, walletSubscribers)
	}
}
