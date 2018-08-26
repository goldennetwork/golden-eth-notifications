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

func NewPendingTransactionSubscribe(client *rpc.Client, url string, timeRetry time.Duration) *PendingTxSub {
	pendingTxSub := make(chan string)
	newBlockSub := make(chan interface{})
	return &PendingTxSub{
		client:       client,
		pendingTxSub: pendingTxSub,
		newBlockSub:  newBlockSub,
		timeRetry:    timeRetry,
	}
}

func StartSubscribe(client *rpc.Client, url string) {
	pendingTxSub := NewPendingTransactionSubscribe(client, url, 2*time.Second)
	pendingTxSub.startSubscribe()

	time.Sleep(pendingTxSub.timeRetry)
	StartSubscribe(client, url)
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
