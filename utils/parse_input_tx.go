package utils

import (
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/goldennetwork/golden-eth-notifications/types"
)

func ParseInputTx(input string, decimals int8) *types.InputData {
	if len(input) == 2 {
		return nil
	}
	methodID := input[0:10]
	toAddress := strings.TrimLeft(input[10:74], "0")
	value := strings.TrimLeft(input[75:], "0")
	valDcms, _ := ConvertInputValueWithDecimal(value, decimals)

	return &types.InputData{
		MethodID:          methodID,
		ToAddress:         toAddress,
		Value:             value,
		ValueWithDecimals: valDcms,
	}
}

func ConvertInputValueWithDecimal(valStr string, decimals int8) (string, error) {
	if !strings.HasPrefix(valStr, "0x") {
		valStr = "0x" + valStr
	}
	valBigInt, err := hexutil.DecodeBig(valStr)
	if err != nil {
		log.Println(err)
		return "", err
	}

	valBigIntString := valBigInt.String()
	if valBigIntString == "0" {
		return valBigIntString, nil
	}
	if decimals == 0 {
		return valBigIntString, nil
	}

	decimalsInt := int(decimals)
	lenVal := len(valBigIntString)
	pre := ""
	suf := ""

	if lenVal < decimalsInt {
		pre = "0."
		for i := 0; i < decimalsInt-lenVal-1; i++ {
			suf = suf + "0"
		}
		suf += valBigIntString
	} else {
		pre = valBigIntString[:lenVal-decimalsInt] + "."
		suf = strings.TrimRight(valBigIntString[lenVal-decimalsInt:], "0")
	}
	return pre + suf, nil
}
