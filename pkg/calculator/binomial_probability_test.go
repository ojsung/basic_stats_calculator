package calculator

import (
	"math/big"
	"testing"

	bu "github.com/ojsung/basic_stats_calculator/internal/big_utils"
)

func Test_CalculateBinomialProbability(t *testing.T) {
	type args = struct {
		chanceOfSuccess *big.Float
		trials          int64
		successes       int64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "It should return 0.375 for p=0.5, n=4, k=2",
			args: args{bu.StrToFloat("0.5"), 4, 2},
			want: "0.375",
		},
		{
			name: "It should return 0.32768 for p=0.2, n=5, k=0",
			args: args{bu.StrToFloat("0.2"), 5, 0},
			want: "0.32768",
		},
		{
			name: "It should return 0.00032 for p=0.2, n=5, k=5",
			args: args{bu.StrToFloat("0.2"), 5, 5},
			want: "0.00032",
		},
		{
			name: "It should return 0.24609375 for p=0.5, n=10, k=5",
			args: args{bu.StrToFloat("0.5"), 10, 5},
			want: "0.24609375",
		},
		{
			name: "It should return 0.343 for p=0.7, n=3, k=3",
			args: args{bu.StrToFloat("0.7"), 3, 3},
			want: "0.343",
		},
		{
			name: "It should return 0.441 for p=0.7, n=3, k=2",
			args: args{bu.StrToFloat("0.7"), 3, 2},
			want: "0.441",
		},
		{
			name: "It should return 0.3486784401 for p=0.1, n=10, k=0",
			args: args{bu.StrToFloat("0.1"), 10, 0},
			want: "0.3486784401",
		},
		{
			name: "It should return 1.0 for p=1.0, n=5, k=5",
			args: args{bu.StrToFloat("1.0"), 5, 5},
			want: "1.0",
		},
		{
			name: "It should return 1.0 for p=0.0, n=5, k=0",
			args: args{bu.StrToFloat("0.0"), 5, 0},
			want: "1.0",
		},
		{
			name: "It should return 0.0 for p=0.0, n=5, k=1",
			args: args{bu.StrToFloat("0.0"), 5, 1},
			want: "0.0",
		},
		{
			name: "It should return 1.0 for p=0.5, n=0, k=0",
			args: args{bu.StrToFloat("0.5"), 0, 0},
			want: "1.0",
		},
		{
			name: "It should return 0.2109375 for p=0.25, n=4, k=2",
			args: args{bu.StrToFloat("0.25"), 4, 2},
			want: "0.2109375",
		},
		{
			name: "It should return 0.3456 for p=0.6, n=5, k=3",
			args: args{bu.StrToFloat("0.6"), 5, 3},
			want: "0.3456",
		},
		{
			name:    "It should return an error when n < k",
			args:    args{bu.StrToFloat("0.5"), 3, 5},
			wantErr: true,
		},
		{
			name:    "It should return an error when p < 0",
			args:    args{bu.StrToFloat("-0.1"), 5, 2},
			wantErr: true,
		},
		{
			name:    "It should return an error when p > 1",
			args:    args{bu.StrToFloat("1.1"), 5, 2},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CalculateBinomialProbability(tt.args.chanceOfSuccess, tt.args.trials, tt.args.successes)
			if (err != nil) != tt.wantErr {
				t.Errorf("CalculateBinomialProbability() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if compare := bu.NewCompare(&got, tt.want); !compare.Equal() {
				t.Errorf("CalculateBinomialProbability() = %v, want %v", compare.ActualAsString, compare.Expected)
			}
		})
	}
}

func Test_BinomialProbabilitySumsToOne(t *testing.T) {
	tests := []struct {
		name string
		p    string
		n    int64
	}{
		{"p=0.3, n=5", "0.3", 5},
		{"p=0.5, n=10", "0.5", 10},
		{"p=0.7, n=4", "0.7", 4},
		{"p=0.1, n=20", "0.1", 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := bu.StrToFloat(tt.p)
			sum := bu.PrecFloat()
			for k := int64(0); k <= tt.n; k++ {
				prob, err := CalculateBinomialProbability(p, tt.n, k)
				if err != nil {
					t.Fatalf("unexpected error at k=%d: %v", k, err)
				}
				sum.Add(sum, &prob)
			}
			if compare := bu.NewCompare(sum, "1.000000000"); !compare.Equal() {
				t.Errorf("probabilities do not sum to 1: got %v", compare.ActualAsString)
			}
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

func Test_calculateProbabilityOfKSuccesses(t *testing.T) {
	type args struct {
		chanceOfSuccess *big.Float
		trials          int64
		successes       int64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "It should return an error when p < 0",
			args:    args{bu.StrToFloat("-0.1"), 5, 2},
			wantErr: true,
		},
		{
			name:    "It should return an error when p > 1",
			args:    args{bu.StrToFloat("1.1"), 5, 2},
			wantErr: true,
		},
		{
			name:    "It should return 0.0005497558139 for p = 0.2, n = 15, k = 3",
			args:    args{bu.StrToFloat("0.2"), 15, 3},
			wantErr: false,
			want:    "0.0005497558139",
		},
		{
			name:    "It should return 0.00032 for p = 0.2, n = 5, k = 5",
			args:    args{bu.StrToFloat("0.2"), 5, 5},
			wantErr: false,
			want:    "0.00032",
		},
		{
			name:    "It should return 0.59049 for p = 0.1, n = 5, k = 0",
			args:    args{bu.StrToFloat("0.1"), 5, 0},
			wantErr: false,
			want:    "0.59049",
		},
		{
			name:    "It should return 0 for p = 0",
			args:    args{bu.StrToFloat("0"), 5, 2},
			wantErr: false,
			want:    "0",
		},
		{
			name:    "It should return 0 for p = 1 when k < n",
			args:    args{bu.StrToFloat("1"), 5, 2},
			wantErr: false,
			want:    "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := calculateProbabilityOfKSuccesses(tt.args.chanceOfSuccess, tt.args.trials, tt.args.successes)
			if (err != nil) != tt.wantErr {
				t.Errorf("calculateProbabilityOfKSuccesses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if got == nil {
				t.Errorf("calculateProbabilityOfKSuccesses() returned nil, want %v", tt.want)
				return
			}
			if compare := bu.NewCompare(got, tt.want); !compare.Equal() {
				t.Errorf("calculateProbabilityOfKSuccesses() = %v, want %v", compare.ActualAsString, compare.Expected)
			}
		})
	}
}
