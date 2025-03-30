package math_utils

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
				f, _ := new(big.Float).SetString("0.0123")
				return f
			}(),
		},
		{
			"It should parse mixed values",
			"123.456",
			func() *big.Float {
				f, _ := new(big.Float).SetString("123.456")
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
			if got, err := Pow(tt.args.base, tt.args.exponent); err != nil {
				t.Errorf("Pow() error = %v", err)
			} else if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pow() = %v, want %v", got, tt.want)
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
		name string
		argument *big.Float
		want *big.Float
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
			"It should error for negative numbers",
			big.NewFloat(-1.0),
			nil,
			true,
		},
		{
			"It should return a a value for positive, real numbers (argument = 4)",
			new(big.Float).SetInt64(4),
			BigFloatFromString("1.3862943611198906"),
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
