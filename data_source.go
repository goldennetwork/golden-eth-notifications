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
	Data map[string][]WalletSubscriber
	lock *sync.RWMutex
}

func (ds *DefaultDataSouce) FindWalletSubscribers(fromAddress, toAddress string) []WalletSubscriber {
	var result []WalletSubscriber
	ds.lock.Lock()
	walletSubFrom, foundFrom := ds.Data[fromAddress]
	if foundFrom {
		result = append(result, walletSubFrom...)
	}

	walletSubTo, foundTo := ds.Data[toAddress]
	if foundTo {
		result = append(result, walletSubTo...)
	}
	ds.lock.Unlock()
	return result
}

func (ds *DefaultDataSouce) SubscribeWallet(walletName, walletAddress, deviceToken string) {
	ds.lock.Lock()
	walletSubs, found := ds.Data[walletAddress]
	walletSubscriber := WalletSubscriber{
		WalletName:    walletName,
		WalletAddress: walletAddress,
		DeviceToken:   deviceToken,
	}

	if found {
		ds.Data[walletAddress] = append(walletSubs, walletSubscriber)
	} else {
		ds.Data[walletAddress] = []WalletSubscriber{walletSubscriber}
	}
	ds.lock.Unlock()
}

func (ds *DefaultDataSouce) UnsubscribeWallet(walletAddress, deviceToken string) {
	ds.lock.Lock()
	walletSubs, found := ds.Data[walletAddress]
	walletSubscribers := []WalletSubscriber{}
	if found {
		for _, ws := range walletSubs {
			if ws.DeviceToken == deviceToken {
				continue
			}
			walletSubscribers = append(walletSubscribers, ws)
		}
		ds.Data[walletAddress] = walletSubscribers
	}
	ds.lock.Unlock()
}

func (ds *DefaultDataSouce) UnsubscribeWalletAllDevice(walletAddress string) {
	ds.lock.Lock()
	delete(ds.Data, walletAddress)
	ds.lock.Unlock()
}
