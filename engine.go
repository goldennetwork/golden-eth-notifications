package ethNotification

import (
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/rpc"
)

type Engine struct {
	c           *rpc.Client
	pushKey     string
	pushTitle   string
	DataSource  EngineDataSource
	MessageHook MessageHook
}

func NewEngine(config EngineConfig) Engine {
	client, err := rpc.Dial(config.WSURL)
	if err != nil {
		panic("Can not connect to " + config.WSURL)
	}

	return Engine{
		c:         client,
		pushKey:   config.FCM_PUSH_KEY,
		pushTitle: config.FCM_PUSH_TITLE,
		DataSource: &DefaultDataSouce{
			Data: make(map[string][]WalletSubscriber),
			lock: &sync.RWMutex{},
		},
		MessageHook: newMessageHook(),
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

func (e *Engine) SubscribeWallet(walletName, address, deviceToken string) {
	go e.DataSource.SubscribeWallet(walletName, address, deviceToken)
}

func (e *Engine) UnsubscribeWallet(address string) {
	go e.DataSource.UnsubscribeWalletAllDevice(address)
}

// Hook
func (e *Engine) OnBeforeSendMessage(hdl func(*Transaction, WalletSubscriber)) {
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
