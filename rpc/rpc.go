package rpc

import (
	"github.com/ethereum/go-ethereum/rpc"
)

type RPCQuery struct {
	client *rpc.Client
}

func NewRpcQuery struct {
	newRpcClient
}

func NewRpcClient(url string) *rpc.Client {
	client, err := rpc.Dial(url)
	if err != nil {
		log.Fatalln("Cannot connect to full node > Error: ", err)
	}
	return client
}
