package ethNotification

import (
	"log"

	"github.com/ethereum/go-ethereum/rpc"
)

type Engine struct {
	c                *rpc.Client
	ethSub           *ethSub
	pushKey          string
	pushTitle        string
	dataSource       EngineDataSource
	tokenDataSource  EngineTokenDataSource
	cacheData        EngineCache
	messageHook      MessageHook
	ChainName        string
	isAllowPendingTx bool
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
		c:                client,
		pushKey:          config.FCM_PUSH_KEY,
		pushTitle:        config.FCM_PUSH_TITLE,
		dataSource:       newDefaultDataSource(),
		tokenDataSource:  newDefaultTokenDataSource(),
		cacheData:        newDefaultEngineCache(),
		messageHook:      newMessageHook(),
		ChainName:        config.CHAIN_NAME,
		isAllowPendingTx: true,
	}
}

func (e *Engine) Start() {
	log.Println("ENGINE START!")
	ethSub := newETHSub(e)
	e.ethSub = &ethSub
	e.ethSub.StartEtherSub()
}

func (e *Engine) Stop() {
	log.Println("ENGINE STOPPED !")
	e.ethSub.cancel()
	e.ethSub = nil
}

func (e *Engine) SetDataSource(ds EngineDataSource) {
	e.dataSource = ds
}

func (e *Engine) SetEngineCache(ec EngineCache) {
	e.cacheData = ec
}

func (e *Engine) SetTokenDataSource(etds EngineTokenDataSource) {
	e.tokenDataSource = etds
}

func (e *Engine) SubscribeWallet(walletName, address, deviceToken string) {
	go e.dataSource.SubscribeWallet(walletName, address, deviceToken)
}

func (e *Engine) UnsubscribeWallet(address string) {
	go e.dataSource.UnsubscribeWalletAllDevice(address)
}

// Hook
func (e *Engine) OnBeforeSendMessage(hdl func(*Transaction, WalletSubscriber, PushMessage)) {
	e.messageHook.BeforeSend = hdl
}

func (e *Engine) OnAfterSendMessage(hdl func(*Transaction, WalletSubscriber, PushMessage)) {
	e.messageHook.AfterSend = hdl
}

func (e *Engine) SetMessageTitle(hdl func(*Transaction, WalletSubscriber) string) {
	e.messageHook.MessageTitle = hdl
}

func (e *Engine) SetMessagePayload(hdl func(*Transaction, WalletSubscriber) map[string]interface{}) {
	e.messageHook.MessagePayload = hdl
}

func (e *Engine) SetAllowSendMessage(hdl func(*Transaction, WalletSubscriber, PushMessage) bool) {
	e.messageHook.AllowSend = hdl
}

func (e *Engine) pushMessage(tran *Transaction, walletSubs []WalletSubscriber) {
	for _, ws := range walletSubs {
		message := PushMessage{
			Title:        e.pushTitle,
			Sound:        "default",
			Content:      e.messageHook.MessageTitle(tran, ws),
			Badge:        "1",
			DeviceTokens: []string{ws.DeviceToken},
			Payload:      e.messageHook.MessagePayload(tran, ws),
		}
		e.messageHook.BeforeSend(tran, ws, message)

		if e.messageHook.AllowSend(tran, ws, message) {
			sendMessage(e.pushKey, &message)
			e.messageHook.AfterSend(tran, ws, message)
		}
	}
}
