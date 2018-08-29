package ethNotification

import (
	"context"
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
	message := PushMessage{
		Title:   "Golden",
		Sound:   "default",
		Content: "Rinkeby: You have a new pending transaction: " + tran.Hash,
		Badge:   "1",
		DeviceTokens: []string{
			"cnMWyehW6oU:APA91bE5SgWL9ZiVZXUtldC3_EIOHJKLbjqrorRN67kRyVex2ETV6axzx8RFB60pcqp4IduCPr9diTCtEixkGzoc7KM7Fp2IdjxAnm657757JxB0Kwng-I_aWXNsyTyPHiCFUvopKB3x",
			"e8AVw-dZ8IM:APA91bEpzXshrUogZb8iPcJwoJzPxhJoX_WwS4cx0IuNcsfipWVLy_AvpROh9kbXaVzsoiAVcdmt-BoIuh8s56YeFzWKgUgaz93GMARLyxmpHQ0PvA4xizHtEwd3O6VIDoIXGSCMti-J",
		},
	}

	SendMessage(hdl.engine.pushKey, &message)
}
