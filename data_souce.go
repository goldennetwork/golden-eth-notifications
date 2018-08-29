package ethNotification

import (
	"sync"
)

type EngineDataSource interface {
	FindDeviceTokens(walletAddress string) []string
	SubscribeWallet(walletAddress string, deviceToken string)
	UnsubscribeWallet(walletAddress string, deviceToken string)
	UnsubscribeWalletAllDevice(walletAddress string)
}

type DefaultDataSouce struct {
	Data map[string][]string
	lock *sync.RWMutex
}

func (ds DefaultDataSouce) FindDeviceTokens(walletAddress string) []string {
	tokens, ok := ds.Data[walletAddress]
	if ok {
		return tokens
	}
	return []string{}
}

func (ds DefaultDataSouce) SubscribeWallet(walletAddress string, deviceToken string) {
	ds.lock.Lock()
	tokens, ok := ds.Data[walletAddress]
	if ok {
		tokens = append(tokens, deviceToken)
		ds.Data[walletAddress] = tokens
	} else {
		ds.Data[walletAddress] = []string{deviceToken}
	}
	ds.lock.Unlock()
}

func (ds DefaultDataSouce) UnsubscribeWallet(walletAddress string, deviceToken string) {
	ds.lock.Lock()
	tokens, ok := ds.Data[walletAddress]
	if ok {
		dts := []string{}
		for i, td := range tokens {
			if deviceToken != td {
				dts = append(dts, tokens[i])
			}
		}
		ds.Data[walletAddress] = dts
	}
	ds.lock.Unlock()
}

func (ds DefaultDataSouce) UnsubscribeWalletAllDevice(walletAddress string) {
	ds.lock.Lock()
	delete(ds.Data, walletAddress)
	ds.lock.Unlock()
}
