package ethNotification

import (
	"context"
	"fmt"
	"log"
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
		log.Println("Transaction error: ", err.Error())
		return err
	}

	if isETHTransaction(transaction) {
		log.Println(transaction)
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
	return result, nil
}

func isETHTransaction(tran *Transaction) bool {
	return tran.Input == "0x"
}

func isERCTransaction(tran *Transaction) bool {
	return false // To do
}

func (hdl txHashHandler) pushPendingTransaction(tran *Transaction) {

	walletSubscribers := hdl.engine.DataSource.FindWalletSubscribers(tran.From, tran.To)

	for _, ws := range walletSubscribers {
		message := PushMessage{
			Title:        hdl.engine.pushTitle,
			Sound:        "default",
			Content:      fmt.Sprintf("Wallet (%s) has new transaction with value: %s ", ws.WalletName, tran.Value),
			Badge:        "1",
			DeviceTokens: []string{ws.DeviceToken},
		}
		sendMessage(hdl.engine.pushKey, &message)
	}
}
