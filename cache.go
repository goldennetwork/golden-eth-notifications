package ethNotification

import (
	"sync"
)

type EngineCache interface {
	Get(txHash string) (CacheData, error)
	Set(txHash string, ws []WalletSubscriber, txInfo Transaction)
	Remove(txHash string)
}

type DefaultEngineCache struct {
	Data map[string]CacheData
	l    *sync.RWMutex
}

func (ec *DefaultEngineCache) Get(txHash string) (CacheData, error) {
	cd, found := ec.Data[txHash]
	if !found {
		return CacheData{}, ErrTransactionNotFound
	}

	return cd, nil
}

func (ec *DefaultEngineCache) Set(txHash string, ws []WalletSubscriber, txInfo Transaction) {
	ec.l.Lock()
	ec.Data[txHash] = CacheData{
		Transaction:       txInfo,
		WalletSubscribers: ws,
	}
	ec.l.Unlock()
}

func (ec *DefaultEngineCache) Remove(txHash string) {
	ec.l.Lock()
	delete(ec.Data, txHash)
	ec.l.Unlock()
}
