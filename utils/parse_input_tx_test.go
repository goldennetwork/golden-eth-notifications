package utils

import (
	"math/big"
	"reflect"
	"testing"
	"time"

	"github.com/goldennetwork/golden-eth-notifications/types"
)

func TestParseInputTx(t *testing.T) {
	type args struct {
		input    string
		decimals int8
	}
	tests := []struct {
		name string
		args args
		want *types.InputData
	}{
		{
			"Test case 1: Transaction VECHAIN",
			args{
				input:    "0xa9059cbb00000000000000000000000019bbb1c4407dcee04059682949f32cb3b513bf7600000000000000000000000000000000000000000000043c2b2187680e25a000",
				decimals: 18,
			},
			&types.InputData{
				ToAddress:         "19bbb1c4407dcee04059682949f32cb3b513bf76",
				Value:             "43c2b2187680e25a000",
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
			&types.InputData{
				ToAddress:         "f8e2f119e4c9f5bd939cad8ba59dbbf68af10949",
				Value:             "22d191fac50",
				ValueWithDecimals: "23927.1829",
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
func TestConvertInputValueWithDecimal(t *testing.T) {
	type args struct {
		val      string
		decimals int8
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"Test case 1", args{"2b364320ad5f60000", 18}, "49.82", false},
		{"Test case 2", args{"0", 18}, "0", false},
		{"Test case 3", args{"0xfff", 2}, "40.95", false},
		{"Test case 4", args{"BCF", 5}, "0.3023", false},
		{"Test case 5", args{"B43132B3", 0}, "3023123123", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			got, err := ConvertInputValueWithDecimal(tt.args.val, tt.args.decimals)
			t.Log(time.Now().Sub(start))
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertInputValueWithDecimal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ConvertInputValueWithDecimal() = %v, want %v", got, tt.want)
			}
			t.Log(got)
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
