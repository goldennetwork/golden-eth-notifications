package subscribe

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/goldennetwork/golden-eth-notifications/push"
	"github.com/goldennetwork/golden-eth-notifications/types"
)

type IRPCProvider interface {
	GetTransactionInfoByHash(txHash string) (*types.Transaction, error)
}

type EthereumSubscribe struct {
	pendingTxSub chan string
	newBlockSub  chan interface{}
	timeRetry    time.Duration
}

func NewEthereumSubscribe(timeRetry time.Duration) *EthereumSubscribe {
	pendingTxSub := make(chan string)
	newBlockSub := make(chan interface{})
	return &EthereumSubscribe{
		pendingTxSub: pendingTxSub,
		newBlockSub:  newBlockSub,
		timeRetry:    timeRetry,
	}
}

func (es *EthereumSubscribe) StartEtherSub(client *rpc.Client, rpcQuery IRPCProvider) {
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
			go HandleFillTransaction(client, txHash, rpcQuery)
		case <-es.newBlockSub:
			// block

		case <-subTx.Err():
		case <-subBlock.Err():
		case <-ctx.Done():
			break
		}
	}
}

func HandleFillTransaction(client *rpc.Client, txHash string, rpcQuery IRPCProvider) {

	txInfo, err := rpcQuery.GetTransactionInfoByHash(txHash)
	if err != nil {
		log.Fatalln(err)
	}

	if txInfo == nil {
		return
	}

	if txInfo.IsNormalTx() {
		//query contractinfor
		// decimals := int8(18)
		//parse Input
		// input := utils.ParseInputTx(txInfo.Input, decimals)
		//filter
		dvsToken, found := FilterToGetDeviceTokenByFromToFeilds(txInfo.From, txInfo.To)
		log.Println(found, dvsToken, txHash)
		if !found {
			return
		}
		//push notidication
		content := "PendingTx: " + txHash + " - Value: " + txInfo.Value
		message := types.PushMessage{
			Title:        os.Getenv("FCM_TITLE_PUSH"),
			Sound:        "default",
			Content:      content,
			Badge:        "1",
			DeviceTokens: dvsToken,
		}
		push.SendMessage(&message)
	}
}

var (
	trackingAddress = map[string]string{
		"Address1": "DeviceToken1",
	}
)

func FilterToGetDeviceTokenByFromToFeilds(from, to string) ([]string, bool) {
	result := []string{}
	log.Println(trackingAddress)
	dvToken1, found1 := trackingAddress[from]
	if found1 {
		result = append(result, dvToken1)
	}
	dvToken2, found2 := trackingAddress[to]
	if found2 {
		result = append(result, dvToken2)
	}
	return result, len(result) > 0

}
