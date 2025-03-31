package big_helpers

import (
	"math/big"
	"reflect"
	"testing"
)

func Test_bigFloatFromString(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want *big.Float
	}{
		{
			"It should parse decimal values",
			"0.0123",
			func() *big.Float {
				f, _ := PrecFloat().SetString("0.0123")
				return f
			}(),
		},
		{
			"It should parse mixed values",
			"123.456",
			func() *big.Float {
				f, _ := PrecFloat().SetString("123.456")
				return f
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BigFloatFromString(tt.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bigFloatFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

