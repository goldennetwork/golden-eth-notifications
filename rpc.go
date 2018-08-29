package main

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/rpc"
	"github.com/goldennetwork/golden-eth-notifications/types"
)

var (
	ErrNotFound = errors.New("Not found")
)

type RPCQuery struct {
	client *rpc.Client
	// context context.Context
}

func NewRPCQuery(client *rpc.Client) RPCQuery {
	return RPCQuery{client}
}

func (r RPCQuery) GetTransactionInfoByHash(txHash string) (*types.Transaction, error) {
	var result *types.Transaction
	if err := r.client.CallContext(context.Background(), &result, types.MethodNameGetTxByHash, txHash); err != nil {
		return nil, err
	}
	if result == nil {
		return nil, ErrNotFound
	}
	return result, nil
}

// func (r RPCQuery) GetBlockInfoByBlockNumber()
