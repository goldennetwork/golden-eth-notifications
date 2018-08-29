package ethNotification

import (
	"context"
	"log"
)

type ethSub struct {
	engine       *Engine
	pendingTxSub chan string
	newBlockSub  chan interface{}
}

func newETHSub(engine *Engine) ethSub {
	return ethSub{
		engine:       engine,
		pendingTxSub: make(chan string),
		newBlockSub:  make(chan interface{}),
	}
}

func (es *ethSub) StartEtherSub() {
	client := es.engine.c
	ctxPar := context.Background()
	ctx, cancelFunc := context.WithCancel(ctxPar)

	subTx, errSubTx := client.EthSubscribe(ctx, es.pendingTxSub, "newPendingTransactions")
	subBlock, errSubBlock := client.EthSubscribe(ctx, es.newBlockSub, "newHeads")

	unsubsribe := func() {
		if subBlock != nil {
			subBlock.Unsubscribe()
		}

		if subTx != nil {
			subTx.Unsubscribe()
		}
	}

	defer func() {
		cancelFunc()
		unsubsribe()
		log.Println("Stop engine !")
		// go es.StartEtherSub()
	}()

	if errSubTx != nil || errSubBlock != nil {
		return
	}

	for {
		select {
		case txHash := <-es.pendingTxSub:
			go func() {
				log.Println(txHash)
				NewTxHashHandler(es.engine, txHash).Handle()
			}()

		case <-es.newBlockSub:
			// block

		case <-subTx.Err():
		case <-subBlock.Err():
		case <-ctx.Done():
			break
		}
	}
}
