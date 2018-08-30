package ethNotification

import (
	"fmt"
	"math/big"
)

type MessageHook struct {
	BeforeSend     func(*Transaction, WalletSubscriber)
	MessageTitle   func(*Transaction, WalletSubscriber) string
	MessagePayload func(*Transaction, WalletSubscriber) map[string]interface{}
	AfterSend      func(*Transaction, WalletSubscriber, PushMessage)
}

func beforeSend(tran *Transaction, ws WalletSubscriber) {
}

func afterSend(tran *Transaction, ws WalletSubscriber, pm PushMessage) {
}

func messagePayload(tran *Transaction, ws WalletSubscriber) map[string]interface{} {
	return map[string]interface{}{}
}

func messageTitle(tran *Transaction, ws WalletSubscriber) string {
	bigInt, _ := new(big.Int).SetString(tran.Value, 10)
	value := CoinToNumberInString(bigInt, 18, 5)
	content := fmt.Sprintf("Wallet %s received %s ETH. Status: %s.", ws.WalletName, value, tran.Status.String())
	if tran.From == ws.WalletAddress {
		content = fmt.Sprintf("Wallet %s sent %s ETH. Status: %s.", ws.WalletName, value, tran.Status.String())
	}

	if tran.ChainName != "mainnet" {
		content = tran.ChainName + ": " + content
	}
	return content
}

func newMessageHook() MessageHook {
	return MessageHook{
		BeforeSend:     beforeSend,
		MessageTitle:   messageTitle,
		MessagePayload: messagePayload,
		AfterSend:      afterSend,
	}
}
