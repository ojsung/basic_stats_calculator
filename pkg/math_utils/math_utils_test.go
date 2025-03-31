package math_utils

import (
	"fmt"
	"math/big"
	"testing"

	bh "github.com/ojsung/basic_stats_calculator/internal/big_helpers"
)

const Ln2String = "0.6931471805599453"
const EulerString = "2.7182818284590452"

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
			if compare := bh.NewCompare(got, tt.want); !compare.Equal() {
				t.Errorf("factorial() = %v, want %v", compare.ActualAsString, tt.want)
			}
		})
	}
}

func Test_IntPow(t *testing.T) {
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
				bh.StrToFloat("0.8"),
				big.NewInt(5),
			},
			"0.32768",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntPow(tt.args.base, tt.args.exponent)
			if compare := bh.NewCompare(got, tt.want); !compare.Equal() {
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

func Test_Exp(t *testing.T) {
	tests := []struct {
		name          string
		exponent      *big.Float
		maxIterations int64
		want          string
	}{
		{
			name:     "It should return 1 for e^0",
			exponent: bh.PrecFloat().SetInt64(0),
			want:     "1",
		},
		{
			name:     "It should return 2.718281828459045 for e^1",
			exponent: bh.PrecFloat().SetInt64(1),
			want:     "2.718281828459045",
		},
		{
			name:     "It should return 7.38905609893065 for e^2",
			exponent: bh.PrecFloat().SetInt64(2),
			want:     "7.38905609893065",
		},
		{
			name:     "It should return 0.36787944117144232 for e^-1",
			exponent: bh.PrecFloat().SetFloat64(-1),
			want:     "0.36787944117144232",
		},
		{
			name:          "It should return 1.6487212707001281 for e^0.5 with maxIterations=100",
			exponent:      bh.PrecFloat().SetFloat64(0.5),
			maxIterations: 100,
			want:          "1.6487212707001281",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got *big.Float
			if tt.maxIterations > 0 {
				got = Exp(tt.exponent, tt.maxIterations)
			} else {
				got = Exp(tt.exponent)
			}
			asString := got.Text('f', len(tt.want)-2) // Match precision of expected value
			if asString != tt.want {
				t.Errorf("Exp() = %v, want %v", asString, tt.want)
			}
		})
	}
}

func Test_Ln(t *testing.T) {
	tests := []struct {
		name     string
		argument *big.Float
		want     string
		wantErr  bool
	}{
		{
			name:     "It should error for negative numbers",
			argument: bh.StrToFloat("-1"),
			wantErr:  true,
		},
		{
			name:     "It should error for 0",
			argument: bh.StrToFloat("0"),
			wantErr:  true,
		},
		{
			name:     "It should return 0.6931471805599453 for ln(2)",
			argument: bh.StrToFloat("2"),
			want:     "0.6931471805599453",
			wantErr:  false,
		},
		{
			name:     "It should return 1.9459101490553133 for ln(7)",
			argument: bh.StrToFloat("7"),
			want:     "1.9459101490553133",
			wantErr:  false,
		},
		{
			name:     "It should return -0.6931471805599453 for ln(0.5)",
			argument: bh.StrToFloat("0.5"),
			want:     "-0.6931471805599453",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Ln(tt.argument)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ln() error = %v", err)
			} else if tt.wantErr && err != nil {
				return
			} else if compare := bh.NewCompare(got, tt.want); !compare.Equal() {
				t.Errorf("ln() = %v, want %v", compare.ActualAsString, compare.Expected)
			}
		})
	}
}

func Test_taylorApproximationLn(t *testing.T) {
	tests := []struct {
		name     string
		argument *big.Float
		want     string
		wantErr  bool
	}{
		{
			name:     "It should error for zero",
			argument: bh.PrecFloat().SetInt64(0),
			wantErr:  true,
		},
		{
			name:     "It should error for real negative numbers",
			argument: bh.PrecFloat().SetInt64(-1),
			wantErr:  true,
		},
		{
			name:     "It should error for real numbers greater than 1",
			argument: bh.PrecFloat().SetInt64(4),
			wantErr:  true,
		},
		{
			name:     "It should return 0.6931471805599453 for ln(2)",
			argument: bh.PrecFloat().SetInt64(2),
			want:     "0.6931471805599453",
			wantErr:  false,
		},
		{
			name:     "It should return -0.6931471805599453 for ln(0.5)",
			argument: bh.PrecFloat().SetFloat64(0.5),
			// ln(0.5) = ln(1/2) = ln(2^-1) = -1 * ln(2) = -ln(2)
			want:    fmt.Sprintf("-%v", Ln2String),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := taylorApproximationLn(tt.argument)
			if (err != nil) != tt.wantErr {
				t.Errorf("taylorApproximationLn() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want == "" && got == nil {
				return
			}
			if compare := bh.NewCompare(got, tt.want); !compare.Equal() {
				t.Errorf("taylorApproximationLn() = %v, want %v", compare.ActualAsString, compare.Expected)
			}
		})
	}
}

func Test_determineLeastIterations(t *testing.T) {
	tests := []struct {
		name     string
		argument *big.Float
		want     int64
	}{
		{
			name:     "It should return 16 for 0.1",
			argument: bh.StrToFloat("0.1"),
			want:     16,
		},
		{
			name:     "It should return 22 for 0.2",
			argument: bh.StrToFloat("0.2"),
			want:     22,
		},
		{
			name:     "It should return 50 for argument 0.5",
			argument: bh.StrToFloat("0.5"),
			want:     50,
		},
		{
			name:     "It should return 92 for argument 0.7",
			argument: bh.StrToFloat("0.7"),
			want:     92,
		},
		{
			name:     "It should return 200 for argument 0.85",
			argument: bh.StrToFloat("0.85"),
			want:     200,
		},
		{
			name:     "It should return 327 for argument 0.9",
			argument: bh.StrToFloat("0.9"),
			want:     327,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := determineLeastIterations(tt.argument)
			// we can be a bit fuzzy here

			if got < tt.want || got > int64(float64(tt.want) * 1.1 + 1) {
				t.Errorf("determineLeastIterations() = %v, want a value greater than %v and within 10%% of it", got, tt.want)
			}
		})
	}
}

func Test_Euler(t *testing.T) {
	tests := []struct {
		name     string
		decimals uint16
		want     string
	}{
		{
			"It should return 3 for 0 decimals",
			0,
			"3",
		},
		{
			"It should return 2.7182818284590452 for 16 decimals",
			16,
			"2.7182818284590452",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Euler(tt.decimals)
			asString := got.Text('f', int(tt.decimals))
			if asString != tt.want {
				t.Errorf("Euler() = %v, want %v", asString, tt.want)
			}
		})
	}
}
func Test_FloatPow(t *testing.T) {
	tests := []struct {
		name     string
		base     *big.Float
		exponent *big.Float
		want     string
		wantErr  bool
	}{
		{
			name:     "It should return 8 for 2^3",
			base:     bh.PrecFloat().SetInt64(2),
			exponent: bh.PrecFloat().SetInt64(3),
			want:     "8",
			wantErr:  false,
		},
		{
			name:     "It should return 0.125 for 2^-3",
			base:     bh.PrecFloat().SetInt64(2),
			exponent: bh.PrecFloat().SetFloat64(-3),
			want:     "0.125",
			wantErr:  false,
		},
		{
			name:     "It should return 1 for any number raised to 0",
			base:     bh.PrecFloat().SetInt64(5),
			exponent: bh.PrecFloat().SetInt64(0),
			want:     "1",
			wantErr:  false,
		},
		{
			name:     "It should return 0.707106781165475 for 2^(-0.5)",
			base:     bh.PrecFloat().SetInt64(2),
			exponent: bh.PrecFloat().SetFloat64(-0.5),
			want:     "0.7071067811865475",
			wantErr:  false,
		},
		{
			name:     "It should return 1.4142135623730950 for 2^0.5",
			base:     bh.PrecFloat().SetInt64(2),
			exponent: bh.PrecFloat().SetFloat64(0.5),
			want:     "1.4142135623730950",
			wantErr:  false,
		},
		{
			name:     "It should error for negative base with non-integer exponent",
			base:     bh.PrecFloat().SetFloat64(-2),
			exponent: bh.PrecFloat().SetFloat64(0.5),
			wantErr:  true,
		},
		{
			name:     "It should return 1 for 1^any exponent",
			base:     bh.PrecFloat().SetInt64(1),
			exponent: bh.PrecFloat().SetFloat64(123.456),
			want:     "1",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FloatPow(tt.base, tt.exponent)
			if (err != nil) != tt.wantErr {
				t.Errorf("FloatPow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && tt.wantErr {
				return
			}
			if compare := bh.NewCompare(got, tt.want); !compare.Equal() {
				t.Errorf("FloatPow() = %v, want %v", compare.Actual, compare.Expected)
			}
		})
	}
}
