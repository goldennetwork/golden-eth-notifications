package ethNotification

import (
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/rpc"
)

type Engine struct {
	c          *rpc.Client
	pushKey    string
	pushTitle  string
	DataSource EngineDataSource
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
