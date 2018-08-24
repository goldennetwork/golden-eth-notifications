package utils

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
	"strings"

	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/goldennetwork/golden-eth-notifications/types"
)

func ParseInputTx(input string, decimals int8) *types.InputData {
	if len(input) != 138 {
		return nil
	}
	methodID := input[0:10]
	if methodID != types.MethodIDTransferERC20Token {
		return nil
	}
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

	start := time.Now()
	valBigIntString := valBigInt.String()
	if valBigIntString == "0" {
		log.Println(time.Now().Sub(start))
		return valBigIntString, nil
	}
	if decimals == 0 {
		log.Println(time.Now().Sub(start))
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
	log.Println(time.Now().Sub(start))
	return pre + suf, nil
}

/**
 * @author: thienthongthai
 */
func CoinToNumberInString(value *big.Int, decimal int, number_precision int) string {
	valueRat := new(big.Rat).SetInt(value)
	powDeicmal := new(big.Rat).SetFloat64(math.Pow(float64(10), float64(decimal)))
	valFloat, _ := new(big.Rat).Quo(valueRat, powDeicmal).Float64()

	pow := math.Pow(float64(10), float64(number_precision))
	intValue, _ := strconv.Atoi(fmt.Sprintf("%.0f", valFloat*pow))
	return fmt.Sprintf("%."+fmt.Sprintf("%d", number_precision)+"g", float64(intValue)/pow)
}
