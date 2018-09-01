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
	err := hdl.engine.c.CallContext(context.Background(), &blockInfo, "eth_getBlockByHash", hdl.hash, true)
	if err != nil {
		return nil, err
	}

	return &blockInfo, nil
}

func (hdl blockHashHandler) fetchTransactionsInBock(b *Block) error {
	batchElems := generateTxReceiptBatchElements(b)
	return hdl.engine.c.BatchCallContext(context.Background(), batchElems)
}

func (hdl blockHashHandler) hanlePushWithCache(b *Block) {
	hdl.fetchTransactionsInBock(b)
}

func (hdl blockHashHandler) hanlePushWithoutCache(b *Block) {
	err := hdl.fetchTransactionsInBock(b)
	if err != nil {
		log.Println("Error fetch tx from block: ", err)
		return
	}
	b.Transactions = updateTransactionFromReceipt(hdl.engine.tokenDataSource, b.Transactions)
	walletSubscribersRes := hdl.engine.dataSource.FindWalletSubscribers(b.Transactions)

	for _, wsr := range walletSubscribersRes {
		hdl.engine.pushMessage(wsr.Transaction, wsr.Subscribers)
	}
}
