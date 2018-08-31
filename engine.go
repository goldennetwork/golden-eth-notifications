package ethNotification

import (
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/rpc"
)

type Engine struct {
	c               *rpc.Client
	pushKey         string
	pushTitle       string
	DataSource      EngineDataSource
	TokenDataSource EngineTokenDataSource
	CacheData       EngineCache
	MessageHook     MessageHook
	ChainName       string
}

func NewEngine(config EngineConfig) Engine {
	client, err := rpc.Dial(config.WSURL)
	if err != nil {
		panic("Can not connect to " + config.WSURL)
	}

	if config.CHAIN_NAME == "" {
		config.CHAIN_NAME = "mainnet"
	}

	return Engine{
		c:         client,
		pushKey:   config.FCM_PUSH_KEY,
		pushTitle: config.FCM_PUSH_TITLE,
		DataSource: &DefaultDataSouce{
			Data: make(map[string][]WalletSubscriber),
			lock: &sync.RWMutex{},
		},
		TokenDataSource: newDefaultTokenDataSource(),
		CacheData: &DefaultEngineCache{
			Data: make(map[string]CacheData),
			l:    &sync.RWMutex{},
		},
		MessageHook: newMessageHook(),
		ChainName:   config.CHAIN_NAME,
	}
}

func (e *Engine) Start() {
	log.Println("ENGINE START!")
	ethSub := newETHSub(e)
	ethSub.StartEtherSub()
}

func (e *Engine) SetDataSource(ds EngineDataSource) {
	e.DataSource = ds
}

func (e *Engine) SetEngineCache(ec EngineCache) {
	e.CacheData = ec
}

func (e *Engine) SetTokenDataSource(etds EngineTokenDataSource) {
	e.TokenDataSource = etds
}

func (e *Engine) SubscribeWallet(walletName, address, deviceToken string) {
	go e.DataSource.SubscribeWallet(walletName, address, deviceToken)
}

func (e *Engine) UnsubscribeWallet(address string) {
	go e.DataSource.UnsubscribeWalletAllDevice(address)
}

// Hook
func (e *Engine) OnBeforeSendMessage(hdl func(*Transaction, WalletSubscriber, PushMessage)) {
	e.MessageHook.BeforeSend = hdl
}

func (e *Engine) OnAfterSendMessage(hdl func(*Transaction, WalletSubscriber, PushMessage)) {
	e.MessageHook.AfterSend = hdl
}

func (e *Engine) SetMessageTitle(hdl func(*Transaction, WalletSubscriber) string) {
	e.MessageHook.MessageTitle = hdl
}

func (e *Engine) SetMessagePayload(hdl func(*Transaction, WalletSubscriber) map[string]interface{}) {
	e.MessageHook.MessagePayload = hdl
}

func (e *Engine) pushMessage(tran *Transaction, walletSubs []WalletSubscriber) {
	for _, ws := range walletSubs {
		message := PushMessage{
			Title:        e.pushTitle,
			Sound:        "default",
			Content:      e.MessageHook.MessageTitle(tran, ws),
			Badge:        "1",
			DeviceTokens: []string{ws.DeviceToken},
			Payload:      e.MessageHook.MessagePayload(tran, ws),
		}
		e.MessageHook.BeforeSend(tran, ws, message)
		sendMessage(e.pushKey, &message)
		e.MessageHook.AfterSend(tran, ws, message)
	}
}
