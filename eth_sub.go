package ethNotification

import (
	"context"
	"log"
)

type ethSub struct {
	engine             *Engine
	pendingTxSubChan   chan string
	newBlockSubChan    chan Block
	context            context.Context
	cancel             context.CancelFunc
	currentBlockNumber string
}

func newETHSub(engine *Engine) ethSub {
	ctx, cancelFunc := context.WithCancel(context.Background())

	return ethSub{
		engine:           engine,
		pendingTxSubChan: make(chan string),
		newBlockSubChan:  make(chan Block),
		context:          ctx,
		cancel:           cancelFunc,
	}
}

func (es *ethSub) StartEtherSub() {
	client := es.engine.c

	subTx, errSubTx := client.EthSubscribe(context.Background(), es.pendingTxSubChan, "newPendingTransactions")
	subBlock, errSubBlock := client.EthSubscribe(context.Background(), es.newBlockSubChan, "newHeads")

	unsubsribe := func() {
		if subBlock != nil {
			subBlock.Unsubscribe()
		}

		if subTx != nil {
			subTx.Unsubscribe()
		}
	}

	defer func() {
		unsubsribe()
		// go es.StartEtherSub()
	}()

	if errSubTx != nil || errSubBlock != nil {
		log.Println("No channel")
		return
	}

	for {
		select {
		case txHash := <-es.pendingTxSubChan:
			if !es.engine.isAllowPendingTx {
				subTx.Unsubscribe()
			} else {
				go func() {
					log.Println("Transaction - " + txHash)
					NewTxHashHandler(es.engine, txHash).Handle()
				}()
			}
		case blockHeader := <-es.newBlockSubChan:
			if es.currentBlockNumber != blockHeader.Number {
				es.currentBlockNumber = blockHeader.Number

				go func() {
					log.Println("Block - " + blockHeader.Number)
					NewBlockHashHandler(es.engine, blockHeader.Hash).Handle()
				}()
			}
		case <-subTx.Err():
		case blockErr := <-subBlock.Err():
			log.Println("Block sub error: ", blockErr)
		case <-es.context.Done():
			log.Println("Engine Done.")
			break
		}
	}
}
