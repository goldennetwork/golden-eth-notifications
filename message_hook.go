package ethNotification

import "fmt"

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
	return fmt.Sprintf("Wallet (%s) has new transaction with value: %s ", ws.WalletName, tran.Value)
}

func newMessageHook() MessageHook {
	return MessageHook{
		BeforeSend:     beforeSend,
		MessageTitle:   messageTitle,
		MessagePayload: messagePayload,
		AfterSend:      afterSend,
	}
}
