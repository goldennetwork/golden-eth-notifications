package ethNotification

import (
	"github.com/ethereum/go-ethereum/rpc"
)

type Engine struct {
	c       *rpc.Client
	pushKey string
}

func NewEngine(config EngineConfig) Engine {
	client, err := rpc.Dial(config.WSURL)
	if err != nil {
		panic("Can not connect to " + config.WSURL)
	}

	return Engine{
		c:       client,
		pushKey: config.FCM_PUSH_KEY,
	}
}

func (e *Engine) Start() {
	ethSub := newETHSub(e)
	ethSub.StartEtherSub()
}
