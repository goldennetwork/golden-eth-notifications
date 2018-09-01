package ethNotification

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

func ParseInputTx(input string, decimals int8) *InputData {
	if len(input) != 138 {
		return nil
	}
	methodID := input[0:10]
	if methodID != MethodIDTransferERC20Token.String() {
		return nil
	}
	toAddress := strings.TrimLeft(input[10:74], "0")
	valueRaw := strings.TrimLeft(input[75:], "0")
	valBigInt, _ := ConvertHexStringToBigInt(valueRaw)
	valDcms := ConvertInputValueWithDecimal(valBigInt.String(), decimals)

	return &InputData{
		MethodID:          methodID,
		ToAddress:         toAddress,
		Value:             valueRaw,
		ValueWithDecimals: valDcms,
	}
}

func ConvertHexStringToBigInt(str string) (*big.Int, error) {
	if !strings.HasPrefix(str, "0x") {
		str = "0x" + str
	}
	bigInt, err := hexutil.DecodeBig(str)
	return bigInt, err
}

func ConvertInputValueWithDecimal(valStr string, decimals int8) string {
	trimStr := strings.Trim(valStr, "0")
	if trimStr == "" || decimals == 0 {
		return valStr
	}
	if trimStr == "." {
		return "0"
	}

	decimalsInt := int(decimals)
	lenVal := len(valStr)
	arrStr := strings.Split(valStr, "")

	if lenVal < decimalsInt {
		numberOfZero := decimalsInt - lenVal - 1
		return "0." + strings.Repeat("0", numberOfZero) + strings.TrimRight(valStr, "0")
	}

	if lenVal > decimalsInt {
		index := lenVal - decimalsInt
		pre := strings.Join(arrStr[:index], "")
		suf := strings.Join(arrStr[index:], "")
		if strings.Trim(suf, "0") == "" {
			return pre
		}
		return pre + "." + strings.TrimRight(suf, "0")
	}

	return "0." + strings.TrimRight(valStr, "0")
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

func getTransactionHashesFromBlock(b *Block) []string {
	result := []string{}

	for _, tran := range b.Transactions {
		result = append(result, tran.Hash)
	}

	return result
}

func updateTransactionFromReceipt(ds EngineTokenDataSource, trans []Transaction) []Transaction {
	for i, tran := range trans {
		if tran.Value != "0x0" {
			bigInt, _ := ConvertHexStringToBigInt(tran.Value)
			trans[i].Value = bigInt.String()
		} else {
			trans[i].Value = "0"
		}

		if tran.Receipt.Status == "0x1" {
			trans[i].Status = Success
		} else {
			trans[i].Status = Failure
		}
		fillTokenInfo(ds, &trans[i])
	}
	return trans
}

func fillTokenInfo(ds EngineTokenDataSource, tran *Transaction) {
	tokens := ds.FindTokens([]string{tran.To})
	if len(tokens) > 0 {
		token := tokens[0]
		inputData := ParseInputTx(tran.Input, token.Decimals)
		tran.To = inputData.ToAddress
		tran.Value = inputData.Value
		tran.Receipt.ContractAddress = token.ContractAddress
		tran.TokenDecimal = int(token.Decimals)
		tran.TokenSymbol = token.Symbol
	}
}

func generateTxReceiptBatchElements(b *Block) []rpc.BatchElem {
	result := []rpc.BatchElem{}

	for i, tran := range b.Transactions {
		var arg interface{} = tran.Hash
		be := rpc.BatchElem{
			Method: "eth_getTransactionReceipt",
			Args:   []interface{}{arg},
			Result: &b.Transactions[i].Receipt,
		}
		result = append(result, be)
	}
	return result
}
