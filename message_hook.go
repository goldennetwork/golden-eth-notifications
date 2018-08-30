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
	content := "Wallet %s received %s ETH (%s)."
	if tran.From == ws.WalletAddress {
		content = "Wallet %s sent %s ETH (%s)."
	}

	return fmt.Sprintf(content, ws.WalletName, value, tran.Status.String())
}

func newMessageHook() MessageHook {
	return MessageHook{
		BeforeSend:     beforeSend,
		MessageTitle:   messageTitle,
		MessagePayload: messagePayload,
		AfterSend:      afterSend,
	}
}
