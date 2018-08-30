package ethNotification

import (
	"context"
	"log"
)

type blockHashHandler struct {
	engine *Engine
	hash   string
}

func NewBlockHashHandler(e *Engine, hash string) blockHashHandler {
	return blockHashHandler{
		engine: e,
		hash:   hash,
	}
}

func (hdl blockHashHandler) Handle() error {
	blockInfo, err := hdl.fetchInfoBlock()
	if err != nil {
		log.Println("Block error: ", err.Error())
		return err
	}

	txs := blockInfo.Transactions
	hdl.pushTrackingTransaction(txs)
	return nil
}

func (hdl blockHashHandler) fetchInfoBlock() (*Block, error) {
	blockInfo := Block{}
	err := hdl.engine.c.CallContext(context.Background(), &blockInfo, "eth_getBlockByHash", hdl.hash, false)
	if err != nil {
		return nil, err
	}

	return &blockInfo, nil
}

func (hdl blockHashHandler) fetchTxReceipt(txHash string) (*Transaction, error) {
	txInfoReceipt := Transaction{}
	err := hdl.engine.c.CallContext(context.Background(), &txInfoReceipt, "eth_getTransactionReceipt", txHash)
	if err != nil {
		return nil, err
	}

	if txInfoReceipt.StatusReceipt == "0x1" {
		txInfoReceipt.Status = Success
	}
	if txInfoReceipt.StatusReceipt == "0x0" {
		txInfoReceipt.Status = Failure
	}
	txInfoReceipt.ChainName = hdl.engine.ChainName

	return &txInfoReceipt, nil
}

func (hdl blockHashHandler) pushTrackingTransaction(txs []string) {
	for _, txHash := range txs {
		cd, errCd := hdl.engine.CacheData.Get(txHash)
		txInfoReceipt, errTx := hdl.fetchTxReceipt(txHash)
		if errCd != nil || errTx != nil {
			continue
		}

		txInfoReceipt.Value = cd.Transaction.Value

		go func() {
			for _, ws := range cd.WalletSubscribers {
				message := PushMessage{
					Title:        hdl.engine.pushTitle,
					Sound:        "default",
					Content:      hdl.engine.MessageHook.MessageTitle(txInfoReceipt, ws),
					Badge:        "1",
					DeviceTokens: []string{ws.DeviceToken},
					Payload:      hdl.engine.MessageHook.MessagePayload(txInfoReceipt, ws),
				}
				hdl.engine.MessageHook.BeforeSend(txInfoReceipt, ws)
				sendMessage(hdl.engine.pushKey, &message)
				hdl.engine.MessageHook.AfterSend(txInfoReceipt, ws, message)
			}
		}()
		hdl.engine.CacheData.Remove(txHash)
	}
}
