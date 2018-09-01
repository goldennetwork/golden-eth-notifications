package ethNotification

import (
	"context"
	"log"
	"time"
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

	if blockInfo == nil {
		return nil
	}

	if hdl.engine.isAllowPendingTx {
		hdl.hanlePushWithCache(blockInfo)
	} else {
		hdl.hanlePushWithoutCache(blockInfo)
	}

	return nil
}

func (hdl blockHashHandler) fetchInfoBlock() (*Block, error) {
	time.Sleep(2 * time.Second)
	blockInfo := Block{}
	err := hdl.engine.cB.CallContext(context.Background(), &blockInfo, "eth_getBlockByHash", hdl.hash, true)
	if err != nil {
		return nil, err
	}

	return &blockInfo, nil
}

func (hdl blockHashHandler) fetchTransactionsInBlock(b *Block) error {
	batchElems := generateTxReceiptBatchElements(b)
	return hdl.engine.cB.BatchCallContext(context.Background(), batchElems)
}

func (hdl blockHashHandler) hanlePushWithCache(b *Block) {
	err := hdl.fetchTransactionsInBlock(b)
	if err != nil {
		log.Println("Error fetch tx from block: ", err)
		return
	}
	b.Transactions = updateTransactionFromReceipt(hdl.engine.tokenDataSource, b.Transactions)

	for i, tran := range b.Transactions {
		cd, err := hdl.engine.cacheData.Get(tran.Hash)
		if err == nil {
			go func(tx *Transaction) {
				hdl.engine.pushMessage(tx, cd.WalletSubscribers)
				hdl.engine.cacheData.Remove(tx.Hash)
			}(&b.Transactions[i])
		}
	}
}

func (hdl blockHashHandler) hanlePushWithoutCache(b *Block) {
	err := hdl.fetchTransactionsInBlock(b)
	if err != nil {
		log.Println("Error fetch tx from block: ", err)
		return
	}
	b.Transactions = updateTransactionFromReceipt(hdl.engine.tokenDataSource, b.Transactions)
	walletSubscribersRes := hdl.engine.dataSource.FindWalletSubscribers(b.Transactions)

	for _, wsr := range walletSubscribersRes {
		go hdl.engine.pushMessage(wsr.Transaction, wsr.Subscribers)
	}
}
