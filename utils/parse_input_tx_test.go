package utils

import (
	"reflect"
	"testing"

	"github.com/goldennetwork/golden-eth-notifications/types"
)

func TestParseInputTx(t *testing.T) {

	tests := []struct {
		name string
		args string
		want *types.InputData
	}{
		{"Test case 1",
			"0xa9059cbb0000000000000000000000005f099b04f1fcea17718e6b91e90990c6849055ee0000000000000000000000000000000000000000000000015a3d9f556fd9d800",
			&types.InputData{
				MethodID:  "0xa9059cbb",
				ToAddress: "5f099b04f1fcea17718e6b91e90990c6849055ee",
				Value:     "15a3d9f556fd9d800",
			},
		},
		{"Test case 2", "0x", nil},
		{"Test case 3", "0xa9059cbb000000000000000000000000b714fa495c0332b49d580f2aabd3da667d068b0c00000000000000000000000000000000000000000000000093bbaa5cda828800",
			&types.InputData{
				MethodID:  "0xa9059cbb",
				ToAddress: "b714fa495c0332b49d580f2aabd3da667d068b0c",
				Value:     "93bbaa5cda828800",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseInputTx(tt.args)
			t.Log(got)
			if !(reflect.DeepEqual(tt.want, got)) {
				t.Errorf("Failed")
			}
		})
	}
}
