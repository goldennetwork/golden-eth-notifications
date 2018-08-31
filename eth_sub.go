package ethNotification

import (
	"context"
	"log"
)

type ethSub struct {
	engine       *Engine
	pendingTxSub chan string
	newBlockSub  chan Block
	context      context.Context
	cancel       context.CancelFunc
}

func newETHSub(engine *Engine) ethSub {
	ctx, cancelFunc := context.WithCancel(context.Background())

	return ethSub{
		engine:       engine,
		pendingTxSub: make(chan string),
		newBlockSub:  make(chan Block),
		context:      ctx,
		cancel:       cancelFunc,
	}
}

func (es *ethSub) StartEtherSub() {
	client := es.engine.c

	subTx, errSubTx := client.EthSubscribe(es.context, es.pendingTxSub, "newPendingTransactions")
	subBlock, errSubBlock := client.EthSubscribe(es.context, es.newBlockSub, "newHeads")

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
		return
	}

	for {
		select {
		case txHash := <-es.pendingTxSub:
			go func() {
				log.Println("Transaction - " + txHash)
				// NewTxHashHandler(es.engine, txHash).Handle()
			}()

		case blockHeader := <-es.newBlockSub:
			go func() {
				log.Println("Block - " + blockHeader.Hash)
				// NewBlockHashHandler(es.engine, blockHeader.Hash).Handle()
			}()
		case <-subTx.Err():
		case blockErr := <-subBlock.Err():
			log.Println("Block sub error: ", blockErr)
		case <-es.context.Done():
			break
		}
	}
}
