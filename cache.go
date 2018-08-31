package ethNotification

import (
	"sync"
)

type EngineCache interface {
	Get(txHash string) (CacheData, error)
	Set(txHash string, ws []WalletSubscriber, txInfo Transaction)
	Remove(txHash string)
}

type defaultEngineCache struct {
	Data map[string]CacheData
	l    *sync.RWMutex
}

func newDefaultEngineCache() *defaultEngineCache {
	return &defaultEngineCache{
		Data: make(map[string]CacheData),
		l:    &sync.RWMutex{},
	}
}

func (ec *defaultEngineCache) Get(txHash string) (CacheData, error) {
	cd, found := ec.Data[txHash]
	if !found {
		return CacheData{}, ErrTransactionNotFound
	}

	return cd, nil
}

func (ec *defaultEngineCache) Set(txHash string, ws []WalletSubscriber, txInfo Transaction) {
	ec.l.Lock()
	ec.Data[txHash] = CacheData{
		Transaction:       txInfo,
		WalletSubscribers: ws,
	}
	ec.l.Unlock()
}

func (ec *defaultEngineCache) Remove(txHash string) {
	ec.l.Lock()
	delete(ec.Data, txHash)
	ec.l.Unlock()
}
