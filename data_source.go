package ethNotification

import (
	"sync"
)

type EngineDataSource interface {
	FindWalletSubscribers(transactions []Transaction) []WalletSubscriberResult
	SubscribeWallet(walletName, walletAddress, deviceToken string)
	UnsubscribeWallet(walletAddress, deviceToken string)
	UnsubscribeWalletAllDevice(walletAddress string)
}

type defaultDataSource struct {
	Data map[string][]WalletSubscriber
	lock *sync.RWMutex
}

func newDefaultDataSource() *defaultDataSource {
	return &defaultDataSource{
		Data: make(map[string][]WalletSubscriber),
		lock: &sync.RWMutex{},
	}
}

func (ds *defaultDataSource) FindWalletSubscribers(transactions []Transaction) []WalletSubscriberResult {
	var result []WalletSubscriberResult
	ds.lock.Lock()

	for i, transaction := range transactions {
		walletSubsFrom, foundFrom := ds.Data[transaction.From]
		walletSubsTo, foundTo := ds.Data[transaction.To]
		wsr := WalletSubscriberResult{
			Transaction: &transactions[i],
			Subscribers: []WalletSubscriber{},
		}

		if foundFrom {
			wsr.Subscribers = append(wsr.Subscribers, walletSubsFrom...)
		}

		if foundTo {
			wsr.Subscribers = append(wsr.Subscribers, walletSubsTo...)
		}
		if len(wsr.Subscribers) > 0 {
			result = append(result, wsr)
		}
	}

	ds.lock.Unlock()
	return result
}

func (ds *defaultDataSource) SubscribeWallet(walletName, walletAddress, deviceToken string) {
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

func (ds *defaultDataSource) UnsubscribeWallet(walletAddress, deviceToken string) {
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

func (ds *defaultDataSource) UnsubscribeWalletAllDevice(walletAddress string) {
	ds.lock.Lock()
	delete(ds.Data, walletAddress)
	ds.lock.Unlock()
}
