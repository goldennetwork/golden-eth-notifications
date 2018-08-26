package tracking

import (
	"context"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/rpc"
)

type PendingTxSub struct {
	client       *rpc.Client
	pendingTxSub chan string
	newBlockSub  chan interface{}
	timeRetry    time.Duration
}

func NewPendingTransactionSubscribe(url string, timeRetry time.Duration) *PendingTxSub {
	client := newRpcClient(url)
	pendingTxSub := make(chan string)
	newBlockSub := make(chan interface{})
	return &PendingTxSub{
		client:       client,
		pendingTxSub: pendingTxSub,
		newBlockSub:  newBlockSub,
		timeRetry:    timeRetry,
	}
}

func newRpcClient(url string) *rpc.Client {
	client, err := rpc.Dial(url)
	if err != nil {
		log.Fatalln("Cannot connect to full node > Error: ", err)
	}
	return client
}

func StartSubscribe(url string) {
	pendingTxSub := NewPendingTransactionSubscribe(url, 2*time.Second)
	pendingTxSub.startSubscribe()

	time.Sleep(pendingTxSub.timeRetry)
	StartSubscribe(url)
}

func (t PendingTxSub) startSubscribe() {
	client := t.client
	pendingTxSub := t.pendingTxSub

	sub, err := client.EthSubscribe(context.Background(), pendingTxSub, "newPendingTransactions")
	if err != nil {
		log.Println(err)
	}

	for {
		select {
		case txHash := <-pendingTxSub:
			log.Println(txHash)

		case err := <-sub.Err():
			log.Println(err.Error())
			sub.Unsubscribe()
			close(pendingTxSub)
			break
		}
	}
}
