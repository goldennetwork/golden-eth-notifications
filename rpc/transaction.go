package rpc

import (
	"golden_gate/appError"

	"github.com/goldennetwork/golden-eth-notifications/types"
)

func (r RPCQuery) GetTransactionInfoByHash(txHash string) (*types.Transaction, error) {
	var result *types.Transaction
	if err := r.client.Call(&result, "eth_getTransactionByHash", txHash); err != nil {
		return nil, err
	}
	if result == nil {
		return nil, appError.ErrNotFound
	}
	return result, nil
}
