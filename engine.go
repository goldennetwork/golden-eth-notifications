package ethNotification

import (
	"sync"

	"github.com/ethereum/go-ethereum/rpc"
)

type Engine struct {
	c          *rpc.Client
	pushKey    string
	DataSource EngineDataSource
}

func NewEngine(config EngineConfig) Engine {
	client, err := rpc.Dial(config.WSURL)
	if err != nil {
		panic("Can not connect to " + config.WSURL)
	}

	return Engine{
		c:       client,
		pushKey: config.FCM_PUSH_KEY,
		DataSource: DefaultDataSouce{
			Data: make(map[string][]string),
			lock: &sync.RWMutex{},
		},
	}
}

func (e *Engine) Start() {
	ethSub := newETHSub(e)
	ethSub.StartEtherSub()
}

func (e *Engine) SetDataSource(ds EngineDataSource) {
	e.DataSource = ds
}

func (e *Engine) SubscribeWallet(address string, deviceToken string) {
	go e.DataSource.SubscribeWallet(address, deviceToken)
}

func (e *Engine) UnsubscribeWallet(address string) {
	go e.DataSource.UnsubscribeWalletAllDevice(address)
}
