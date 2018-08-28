package tracking

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/rpc"
)

type EthereumSubscribe struct {
	pendingTxSub chan string
	newBlockSub  chan interface{}
	timeRetry    time.Duration
}

func NewEthereumSubscribe(url string, timeRetry time.Duration) *EthereumSubscribe {
	pendingTxSub := make(chan string)
	newBlockSub := make(chan interface{})
	return &EthereumSubscribe{
		pendingTxSub: pendingTxSub,
		newBlockSub:  newBlockSub,
		timeRetry:    timeRetry,
	}
}

func (es *EthereumSubscribe) StartEtherSub(client *rpc.Client) {
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
		time.Sleep(es.timeRetry)
		// go es.StartEtherSub()
	}()

	if errSubTx != nil {
		return
	}

	if errSubBlock != nil {
		unsubsribe()
		return
	}

	for {
		select {
		case txHash := <-es.pendingTxSub:

		case <-es.newBlockSub:
			// block

		case <-subTx.Err():
		case <-subBlock.Err():
		case <-ctx.Done():
			break
		}
	}
}
