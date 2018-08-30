package ethNotification

import (
	"sync"
)

type EngineDataSource interface {
	FindWalletSubscribers(fromAddress, toAddress string) []WalletSubscriber
	SubscribeWallet(walletName, walletAddress, deviceToken string)
	UnsubscribeWallet(walletAddress, deviceToken string)
	UnsubscribeWalletAllDevice(walletAddress string)
}

type DefaultDataSouce struct {
	Data []WalletSubscriber
	lock *sync.RWMutex
}

func (ds DefaultDataSouce) FindWalletSubscribers(fromAddress, toAddress string) []WalletSubscriber {
	var result []WalletSubscriber
	ds.lock.Lock()
	for _, ws := range ds.Data {
		if ws.WalletAddress == fromAddress || ws.WalletAddress == toAddress {
			result = append(result, ws)
		}
	}
	ds.lock.Unlock()
	return result
}

func (ds DefaultDataSouce) SubscribeWallet(walletName, walletAddress, deviceToken string) {
	ds.lock.Lock()
	walletSubscriber := WalletSubscriber{
		WalletName:    walletName,
		WalletAddress: walletAddress,
		DeviceToken:   deviceToken,
	}
	ds.Data = append(ds.Data, walletSubscriber)
	ds.lock.Unlock()
}

func (ds DefaultDataSouce) UnsubscribeWallet(walletAddress, deviceToken string) {
	ds.lock.Lock()

	walletSubscribers := []WalletSubscriber{}
	for _, ws := range ds.Data {
		if ws.WalletAddress == walletAddress && ws.DeviceToken == deviceToken {
			continue
		}
		walletSubscribers = append(walletSubscribers, ws)
	}
	ds.Data = walletSubscribers
	ds.lock.Unlock()
}

func (ds DefaultDataSouce) UnsubscribeWalletAllDevice(walletAddress string) {
	ds.lock.Lock()
	walletSubscribers := []WalletSubscriber{}
	for _, ws := range ds.Data {
		if ws.WalletAddress == walletAddress {
			continue
		}
		walletSubscribers = append(walletSubscribers, ws)
	}
	ds.Data = walletSubscribers
	ds.lock.Unlock()
}
