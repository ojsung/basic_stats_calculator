package math_utils

import (
	"fmt"
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
				f, _ := precFloat().SetString("0.0123")
				return f
			}(),
		},
		{
			"It should parse mixed values",
			"123.456",
			func() *big.Float {
				f, _ := precFloat().SetString("123.456")
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

func TestPow(t *testing.T) {
	type args struct {
		base     *big.Float
		exponent *big.Int
	}
	tests := []struct {
		name string
		args args
		want *big.Float
	}{
		{
			"It should return 14348907 for 3^15",
			args{
				precFloat().SetInt64(3),
				big.NewInt(15),
			},
			precFloat().SetInt64(14348907),
		},
		{
			"It should return 0.32768 for 0.8^5",
			args{
				BigFloatFromString("0.8"),
				big.NewInt(5),
			},
			BigFloatFromString("0.32768"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Pow(tt.args.base, tt.args.exponent)
			got.SetMode(big.ToNearestEven).SetPrec(128)
			tt.want.SetMode(big.ToNearestEven).SetPrec(128)
			if got.Cmp(tt.want) != 0 {
				t.Errorf("Pow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binaryPowEven(t *testing.T) {
	tests := []struct {
		name        string
		binaryPow   binaryExp
		wantEvenPow binaryExp
	}{
		{
			"It should return the binary exponent form of 6^4 as 36^2 when r=1",
			binaryExp{
				precFloat().SetInt64(1),
				precFloat().SetInt64(6),
				big.NewInt(4),
			},
			binaryExp{
				precFloat().SetInt64(1).SetMode(big.ToZero).SetPrec(128),
				precFloat().SetInt64(36).SetMode(big.ToZero).SetPrec(128),
				big.NewInt(2),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEvenPow := binaryPowEven(tt.binaryPow);
			equal, failed := gotEvenPow.Equal(tt.wantEvenPow)
			if !equal {
				t.Errorf("comparison failed. a: %v, r: %v, x: %v", failed.a, failed.r, failed.x)
			}
		})
	}
}

func Test_binaryPowOdd(t *testing.T) {
	tests := []struct {
		name       string
		binaryPow  binaryExp
		wantOddPow binaryExp
	}{
		{
			"It should return r = 27, a = 81, x = 3 for 3 * 9^7",
			binaryExp{
				precFloat().SetInt64(3),
				precFloat().SetInt64(9),
				big.NewInt(7),
			},
			binaryExp{
				precFloat().SetInt64(27).SetMode(big.ToZero).SetPrec(128),
				precFloat().SetInt64(81).SetMode(big.ToZero).SetPrec(128),
				big.NewInt(3),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOddPow := binaryPowOdd(tt.binaryPow);
			equal, failed := gotOddPow.Equal(tt.wantOddPow)
			if !equal {
				t.Errorf("comparison failed. a: %v, r: %v, x: %v", failed.a, failed.r, failed.x)
			}
		})
	}
}

func TestNaturalPow(t *testing.T) {
	type args struct {
		exponent *big.Float
	}
	tests := []struct {
		name string
		args args
		want *big.Float
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := NaturalPow(tt.args.exponent); err != nil {
				t.Errorf("NaturalPow() error = %v", err)
			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NaturalPow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Ln(t *testing.T) {
	tests := []struct {
		name     string
		argument *big.Float
		want     *big.Float
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, err := Ln(tt.argument); err != nil {
				t.Errorf("Ln() error = %v", err)
			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ln() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_taylorApproximationLn(t *testing.T) {
	tests := []struct {
		name     string
		argument *big.Float
		want     *big.Float
		wantErr  bool
	}{
		{
			"It should error for zero",
			big.NewFloat(0.0),
			nil,
			true,
		},
		{
			"It should error for real negative numbers",
			big.NewFloat(-1.0),
			nil,
			true,
		},
		{
			"It should error for real numbers greater than 1",
			big.NewFloat(4.0),
			nil,
			true,
		},
		{
			"It should return a a value for r with domain (0, 1) in Reals (argument = 0.5)",
			precFloat().SetFloat64(0.5),
			// ln(0.5) = ln(1/2) = ln(2^-1) = -1 * ln(2) = -ln(2)
			BigFloatFromString(fmt.Sprintf("-%v", Ln2String)),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := taylorApproximationLn(tt.argument)
			if (err != nil) != tt.wantErr {
				t.Errorf("taylorApproximationLn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("taylorApproximationLn() = %v, want %v", got, tt.want)
			}
		})
	}
}
