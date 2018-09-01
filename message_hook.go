package ethNotification

import (
	"fmt"
	"math/big"
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
	bigInt, _ := new(big.Int).SetString(tran.Value, 10)

	value := ""
	symbol := "ETH"

	// Format value base on ETH or Token
	if tran.TokenDecimal != 0 {
		value = CoinToNumberInString(bigInt, tran.TokenDecimal, 5)
		// Change symbol if token
		symbol = tran.TokenSymbol
	} else {
		value = CoinToNumberInString(bigInt, 18, 5)
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
