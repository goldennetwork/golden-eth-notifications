package ethNotification

import (
	"context"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/rpc"
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
	} else {
		log.Println(b)
	}
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
	// for _, txHash := range txs {
	// 	cd, errCd := hdl.engine.cacheData.Get(txHash)
	// 	txInfoReceipt, errTx := hdl.fetchTxReceipt(txHash)
	// 	if errCd != nil || errTx != nil {
	// 		continue
	// 	}
	// 	txInfoReceipt.Value = cd.Transaction.Value

	// 	go hdl.engine.pushMessage(txInfoReceipt, cd.WalletSubscribers)
	// 	hdl.engine.cacheData.Remove(txHash)
	// }
}

func generateTxReceiptBatchElements(b *Block) []rpc.BatchElem {
	result := []rpc.BatchElem{}

	for i, tran := range b.Transactions {
		var arg interface{} = tran.Hash
		be := rpc.BatchElem{
			Method: "eth_getTransactionReceipt",
			Args:   []interface{}{arg},
			Result: &b.Transactions[i],
		}
		result = append(result, be)
	}
	return result
}
