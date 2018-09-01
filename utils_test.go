package ethNotification

import (
	"math/big"
	"reflect"
	"testing"
	"time"
)

func TestParseInputTx(t *testing.T) {
	type args struct {
		input    string
		decimals int8
	}
	tests := []struct {
		name string
		args args
		want *InputData
	}{
		{
			"Test case 1: Transaction VECHAIN",
			args{
				input:    "0xa9059cbb00000000000000000000000019bbb1c4407dcee04059682949f32cb3b513bf7600000000000000000000000000000000000000000000043c2b2187680e25a000",
				decimals: 18,
			},
			&InputData{
				ToAddress:         "19bbb1c4407dcee04059682949f32cb3b513bf76",
				Value:             "19999378490000000000000",
				ValueWithDecimals: "19999.37849",
				MethodID:          "0xa9059cbb",
			},
		},
		{
			"Test case 2: Transaction Bytom",
			args{
				input:    "0xa9059cbb000000000000000000000000f8e2f119e4c9f5bd939cad8ba59dbbf68af109490000000000000000000000000000000000000000000000000000022d191fac50",
				decimals: 8,
			},
			&InputData{
				ToAddress:         "f8e2f119e4c9f5bd939cad8ba59dbbf68af10949",
				Value:             "2392718290000",
				ValueWithDecimals: "23927.1829",
				MethodID:          "0xa9059cbb",
			},
		},
		{
			"Test case 3: Transaction",
			args{
				input:    "0xa9059cbb000000000000000000000000af654d6b7254746edb974fe292d36fc8f9da10eb000000000000000000000000000000000000000000000000117c6b5300fe0000",
				decimals: 18,
			},
			&InputData{
				ToAddress:         "af654d6b7254746edb974fe292d36fc8f9da10eb",
				Value:             "1260000000000000000",
				ValueWithDecimals: "1.26",
				MethodID:          "0xa9059cbb",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseInputTx(tt.args.input, tt.args.decimals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseInputTx() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoinToNumberInString(t *testing.T) {
	bigIntTest0 := big.NewInt(19831202373)
	bigIntTest1, _ := new(big.Int).SetString("198312023182371823973", 10)
	bigIntTest2, _ := new(big.Int).SetString("49820000000000000000", 10)
	BigIntTest := []*big.Int{
		bigIntTest0,
		bigIntTest1,
		bigIntTest2,
	}
	type args struct {
		value            *big.Int
		decimal          int
		number_precision int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test case 0",
			args{
				value:            BigIntTest[0],
				decimal:          8,
				number_precision: 5,
			},
			"198.31",
		},
		{"Test case 1",
			args{
				value:            BigIntTest[1],
				decimal:          18,
				number_precision: 5,
			},
			"198.31",
		},
		{"Test case 2",
			args{
				value:            BigIntTest[2],
				decimal:          18,
				number_precision: 4,
			},
			"49.82",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			got := CoinToNumberInString(tt.args.value, tt.args.decimal, tt.args.number_precision)
			t.Log(time.Now().Sub(start))
			if got != tt.want {
				t.Errorf("CoinToNumberInString() = %v, want %v", got, tt.want)
			}
			t.Log(got)
		})
	}
}

func TestConvertInputValueWithDecimal(t *testing.T) {
	type args struct {
		valStr   string
		decimals int8
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test case 1", args{"19999378490000000000000", 18}, "19999.37849"},
		{"Test case 2", args{"0.000", 3}, "0"},
		{"Test case 3", args{"0", 5}, "0"},
		{"Test case 4", args{"8000", 8}, "0.0008"},
		{"Test case 5", args{"9000", 4}, "0.9"},
		{"Test case 6", args{"91237123", 0}, "91237123"},
		{"Test case 7", args{"72000", 3}, "72"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertInputValueWithDecimal(tt.args.valStr, tt.args.decimals); got != tt.want {
				t.Errorf("ConvertInputValueWithDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
