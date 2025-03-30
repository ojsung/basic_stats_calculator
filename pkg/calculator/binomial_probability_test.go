package calculator

import (
	"math/big"
	"reflect"
	"testing"
	mu "github.com/ojsung/basic_stats_calculator/internal/math_utils"
)

func Test_calculateBinomialProbability(t *testing.T) {
	type args = struct {
		chanceOfSuccess *big.Float
		trials int64
		successes int64
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calculateBinomialProbability(tt.args.chanceOfSuccess, tt.args.trials, tt.args.successes)
		})
	}
}

func Test_calculateBinomialCoefficient(t *testing.T) {
	type args = struct {
		n int64
		k int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"It should return 10 for n=10, k=1", args{10, 1},
		},
		{
			"It should return 3003 for n=15, k=5", args{15, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calculateBinomialCoefficient(tt.args.n, tt.args.k)
		})
	}
}

func Test_factorial(t *testing.T) {
	tests := []struct {
		name     string
		argument int64
		want     Int
	}{
		{
			"It should return 1 for 0", 0, big.NewInt(1),
		},
		{
			"It should return 1 for 1", 1, big.NewInt(1),
		},
		{
			"It should return 6 for 3", 3, big.NewInt(6),
		},
		{
			"It should return 3628800 for 10", 10, big.NewInt(3628800),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := factorial(tt.argument); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("factorial() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_calculateProbabilityOfKSuccesses(t *testing.T) {
	type args struct {
		chanceOfSuccess *big.Float
		trials          int64
		successes       int64
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Float
		wantErr bool
	}{
		{
			name:    "It should return an error when trials < successes",
			args:    args{mu.BigFloatFromString("0.5"), 3, 5},
			wantErr: true,
			want:    big.NewFloat(0),
		},
		{
			name:    "It should return .0005497558139 for p = 0.2, n = 15, k = 3",
			args:    args{mu.BigFloatFromString("0.2"), 15, 3},
			wantErr: false,
			want:    mu.BigFloatFromString(".0005497558139"),
		},
		{
			name:    "It should return .08192 for p = 0.2, n = k = 5",
			args:    args{mu.BigFloatFromString("0.2"), 5, 5},
			wantErr: false,
			want: mu.BigFloatFromString("0.08192"),
		},
		{
			name:    "It should use special case for k = 0",
			args:    args{mu.BigFloatFromString("0.1"), 5, 0},
			wantErr: false,
			want: mu.BigFloatFromString("0.59049"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := calculateProbabilityOfKSuccesses(tt.args.chanceOfSuccess, tt.args.trials, tt.args.successes)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculateBinomialProbabilityOfFailure() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
