package ethNotification

import (
	"fmt"
)

type MessageHook struct {
	MessageTitle   func(*Transaction, WalletSubscriber) string
	MessagePayload func(*Transaction, WalletSubscriber) map[string]interface{}
	BeforeSend     func(*Transaction, WalletSubscriber, PushMessage)
	AllowSend      func(*Transaction, WalletSubscriber, PushMessage) bool
	AfterSend      func(*Transaction, WalletSubscriber, PushMessage)
}

func beforeSend(tran *Transaction, ws WalletSubscriber, pm PushMessage) {
}

func afterSend(tran *Transaction, ws WalletSubscriber, pm PushMessage) {
}

func messagePayload(tran *Transaction, ws WalletSubscriber) map[string]interface{} {
	return map[string]interface{}{}
}

func messageTitle(tran *Transaction, ws WalletSubscriber) string {
	value := ""
	symbol := "ETH"

	// Format value base on ETH or Token
	if tran.TokenSymbol != "" {
		value = ConvertInputValueWithDecimal(tran.Value, int8(tran.TokenDecimal))
		// Change symbol if token
		symbol = tran.TokenSymbol
	} else {
		value = ConvertInputValueWithDecimal(tran.Value, 18)
	}

	content := fmt.Sprintf("Wallet %s received %s %s. Status: %s.", ws.WalletName, value, symbol, tran.Status.String())
	if tran.From == ws.WalletAddress {
		content = fmt.Sprintf("Wallet %s sent %s %s. Status: %s.", ws.WalletName, value, symbol, tran.Status.String())
	}

	if tran.ChainName != "mainnet" {
		content = tran.ChainName + ": " + content
	}
	return content
}

func allow(tran *Transaction, ws WalletSubscriber, pm PushMessage) bool {
	return true
}

func newMessageHook() MessageHook {
	return MessageHook{
		BeforeSend:     beforeSend,
		MessageTitle:   messageTitle,
		MessagePayload: messagePayload,
		AllowSend:      allow,
		AfterSend:      afterSend,
	}
}
