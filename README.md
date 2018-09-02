# Golden Ethereum Notification #

A Go engine to track transaction and push notifications to mobile application (Firebase Cloud Message).

## Features

- [x] Tracking pending transaction.
- [x] Tracking new block.
- [x] Support ERC20 transaction.
- [x] Send message to mobile application.

## Installation ##

Golden-eth-notifications can be installed easily by running

	go get -u github.com/goldennetwork/golden-eth-notifications
	
Then you can import it in your project

    import "github.com/goldennetwork/golden-eth-notifications"

## Setting up a simple subscriber
    
#### Simple Usage

```golang
// file: main.go

package main
import (
	ethPush "github.com/goldennetwork/golden-eth-notifications"
)

func main() {
    engine := ethPush.NewEngine(ethPush.EngineConfig{
        WSURL: YOUR_WS_URL_FULLNODE,
        FCM_PUSH_KEY: YOUR_FIREBASE_PUSH_KEY,
        ENABLE_PUSH_PENDING_TX: true, // (*)
    }
    
    engine.SubscribeWallet("WALLET_NAME 1", "YOUR_WALLET_ADDRESS 1", "YOUR_DEVICE_TOKEN 1")
    engine.SubscribeWallet("WALLET_NAME 2", "YOUR_WALLET_ADDRESS 2", "YOUR_DEVICE_TOKEN 2")
    engine.Start()
}
```

> **Important Note:** Setting "ENABLE_PUSH_PENDING_TX: true" require your full node run with option --rpcapi miner,txpool

#### Tracking and Custom message

```golang
    engine.OnBeforeSendMessage(func(tran *ethPush.Transaction, ws ethPush.WalletSubscriber, pushMess ethPush.PushMessage) {
	// do something with your transaction and message
    })
    
    engine.SetMessageTitle(func(tran *ethPush.Transaction, ws ethPush.WalletSubscriber) string {
    	// Custom push message title
	return fmt.Sprintf("Wallet %s received %s from %s", ws.WalletName, tran.Value, tran.From)
    })
    
    engine.SetMessagePayload(func(tran *ethPush.Transaction, ws ethPush.WalletSubscriber) map[string]interface{} {
    	return map[string]interface{}{
		"address": ws.WalletAddress,
		"value":   tran.Value,
		"wallet":  ws.WalletName,
		"tx":      tran,
	}
    })
    
    engine.SetAllowSendMessage(func(tran *ethPush.Transaction, ws ethPush.WalletSubscriber, pushMess ethPush.PushMessage) bool {
    	return true
    })
    
    engine.OnAfterSendMessage(func(tran *ethPush.Transaction, ws ethPush.WalletSubscriber, pushMess ethPush.PushMessage) {
    	// do something with your transaction and message
    })
```

#### Custom Data Source
By default the engine use a simple data source to store and fetch wallet subscribers. To use your custom data source by implementing this interface:

```golang
type EngineDataSource interface {
	FindWalletSubscribers(transactions []Transaction) []WalletSubscriberResult
	SubscribeWallet(walletName, walletAddress, deviceToken string)
	UnsubscribeWallet(walletAddress, deviceToken string)
	UnsubscribeWalletAllDevice(walletAddress string)
}
```

```golang
type CustomDataSource struct {
}

func (cds *CustomDataSource) FindWalletSubscribers(transactions []Transaction) []WalletSubscriberResult{
	// Query your database
}

func (cds *CustomDataSource) SubscribeWallet(walletName, walletAddress, deviceToken string) {
	// Insert into your database
}

func (cds *CustomDataSource) UnsubscribeWallet(walletAddress, deviceToken string) {
	// Delete data from your database
}

func (cds *CustomDataSource) UnsubscribeWalletAllDevice(walletAddress string) {
	// Delete all wallet subscribers from your database
}

engine.SetDataSource(CustomDataSource)
```

#### Custom Token Data Source
```golang
type EngineTokenDataSource interface {
	FindTokens(tokenAddress []string) []TokenContract
}
```

```golang
type TokenDataSource struct {
}

func (ds TokenDataSource) FindTokens(tokenAddress []string) []TokenContract {
	// Query your database
}

engine.SetTokenDataSource(TokenDataSource)
```
#### Custom Cache Data Source
```golang
type EngineCache interface {
	Get(txHash string) (CacheData, error)
	Set(txHash string, ws []WalletSubscriber, txInfo Transaction)
	Remove(txHash string)
}
```

```golang
type Cache struct {
}

func (c *Cache) Get(txHash string) (CacheData, error) {
}
	
func (c *Cache) Set(txHash string, ws []WalletSubscriber, txInfo Transaction) {
}
	
func (c *Cache) Remove(txHash string) {
}

engine.SetEngineCache(Cache)
```

## License ##
MIT License

Copyright (C) 2018 Skylab Technology Cooperation
