package math_utils

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"

	bh "github.com/ojsung/basic_stats_calculator/internal/big_helpers"
)

func Test_Factorial(t *testing.T) {
	tests := []struct {
		name     string
		argument int64
		want     string
	}{
		{
			"It should return 1 for 0", 0, "1",
		},
		{
			"It should return 1 for 1", 1, "1",
		},
		{
			"It should return 6 for 3", 3, "6",
		},
		{
			"It should return 3628800 for 10", 10, "3628800",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Factorial(tt.argument)
			if compare := bh.NewCompare(got, tt.want); !compare.Compare() {
				t.Errorf("factorial() = %v, want %v", compare.ActualAsString, tt.want)
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
		want string
	}{
		{
			"It should return 14348907 for 3^15",
			args{
				bh.PrecFloat().SetInt64(3),
				big.NewInt(15),
			},
			"14348907",
		},
		{
			"It should return 0.32768 for 0.8^5",
			args{
				bh.BigFloatFromString("0.8"),
				big.NewInt(5),
			},
			"0.32768",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Pow(tt.args.base, tt.args.exponent)
			if compare := bh.NewCompare(got, tt.want); !compare.Compare() {
				t.Errorf("comparison failed. expected %v but got %v", compare.Expected, compare.ActualAsString)
			}
		})
	}
}

func Test_binaryPowEven(t *testing.T) {
	tests := []struct {
		name        string
		binaryPow   bh.BinaryExp
		wantEvenPow bh.BinaryExp
	}{
		{
			"It should return the binary exponent form of 6^4 as 36^2 when r=1",
			bh.BinaryExp{
				R: bh.PrecFloat().SetInt64(1),
				A: bh.PrecFloat().SetInt64(6),
				X: big.NewInt(4),
			},
			bh.BinaryExp{
				R: bh.PrecFloat().SetInt64(1).SetMode(big.ToZero).SetPrec(128),
				A: bh.PrecFloat().SetInt64(36).SetMode(big.ToZero).SetPrec(128),
				X: big.NewInt(2),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotEvenPow := binaryPowEven(tt.binaryPow)
			equal, failed := gotEvenPow.Equal(tt.wantEvenPow)
			if !equal {
				t.Errorf("comparison failed. a: %v, r: %v, x: %v", failed.A, failed.R, failed.X)
			}
		})
	}
}

func Test_binaryPowOdd(t *testing.T) {
	tests := []struct {
		name       string
		binaryPow  bh.BinaryExp
		wantOddPow bh.BinaryExp
	}{
		{
			"It should return r = 27, a = 81, x = 3 for 3 * 9^7",
			bh.BinaryExp{
				R: bh.PrecFloat().SetInt64(3),
				A: bh.PrecFloat().SetInt64(9),
				X: big.NewInt(7),
			},
			bh.BinaryExp{
				R: bh.PrecFloat().SetInt64(27).SetMode(big.ToZero).SetPrec(128),
				A: bh.PrecFloat().SetInt64(81).SetMode(big.ToZero).SetPrec(128),
				X: big.NewInt(3),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOddPow := binaryPowOdd(tt.binaryPow)
			equal, failed := gotOddPow.Equal(tt.wantOddPow)
			if !equal {
				t.Errorf("comparison failed. a: %v, r: %v, x: %v", failed.A, failed.R, failed.X)
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
			bh.PrecFloat().SetInt64(0),
			nil,
			true,
		},
		{
			"It should error for real negative numbers",
			bh.PrecFloat().SetInt64(-1),
			nil,
			true,
		},
		{
			"It should error for real numbers greater than 1",
			bh.PrecFloat().SetInt64(4),
			nil,
			true,
		},
		{
			"It should return a a value for r with domain (0, 1) in Reals (argument = 0.5)",
			bh.PrecFloat().SetFloat64(0.5),
			// ln(0.5) = ln(1/2) = ln(2^-1) = -1 * ln(2) = -ln(2)
			bh.BigFloatFromString(fmt.Sprintf("-%v", Ln2String)),
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
			if tt.want == nil && got == nil {
				return
			}
			if got.Cmp(tt.want) == 0 {
				t.Errorf("taylorApproximationLn() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_Euler(t *testing.T) {
	tests := []struct {
		name     string
		decimals uint16
		want     string
		wantErr  bool
	}{
		{
			"It should return 3 for 0 decimals",
			0,
			"3",
			false,
		},
		{
			"It should return 2.7182818284590452 for 16 decimals",
			16,
			"2.7182818284590452",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Euler(tt.decimals)
			if (err != nil) != tt.wantErr {
				t.Errorf("Euler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			asString := got.Text('f', int(tt.decimals))
			if asString != tt.want {
				t.Errorf("Euler() = %v, want %v", asString, tt.want)
			}
		})
	}
}


