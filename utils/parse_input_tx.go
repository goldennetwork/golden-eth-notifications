package utils

import (
	"strings"

	"github.com/goldennetwork/golden-eth-notifications/types"
)

func ParseInputTx(input string) *types.InputData {
	if len(input) == 2 {
		return nil
	}
	methodID := input[0:10]
	toAddress := strings.TrimLeft(input[10:74], "0")
	value := strings.TrimLeft(input[75:], "0")
	return &types.InputData{
		MethodID:  methodID,
		ToAddress: toAddress,
		Value:     value,
	}
}
