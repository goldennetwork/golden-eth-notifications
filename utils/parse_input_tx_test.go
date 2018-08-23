package utils

import (
	"reflect"
	"testing"

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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertInputValueWithDecimal(tt.args.val, tt.args.decimals)
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
