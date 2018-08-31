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
		cd, errCd := hdl.engine.cacheData.Get(txHash)
		txInfoReceipt, errTx := hdl.fetchTxReceipt(txHash)
		if errCd != nil || errTx != nil {
			continue
		}
		txInfoReceipt.Value = cd.Transaction.Value

		go hdl.engine.pushMessage(txInfoReceipt, cd.WalletSubscribers)
		hdl.engine.cacheData.Remove(txHash)
	}
}
