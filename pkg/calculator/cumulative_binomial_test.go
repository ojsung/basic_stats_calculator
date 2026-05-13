package calculator

import (
	"math/big"
	"testing"

	bu "github.com/ojsung/basic_stats_calculator/internal/big_utils"
)

func Test_CumulativeBinomialProbability(t *testing.T) {
	half := bu.StrToFloat("0.5")
	tests := []struct {
		name      string
		p         *big.Float
		n, k      int64
		wantCumul string
		wantTerms int
		wantErr   bool
	}{
		{
			name: "k=0 returns P(X=0) only",
			p:    half, n: 3, k: 0,
			wantCumul: "0.125",
			wantTerms: 1,
		},
		{
			name: "k=1 accumulates two terms",
			p:    half, n: 3, k: 1,
			wantCumul: "0.5",
			wantTerms: 2,
		},
		{
			name: "k=n sums to 1",
			p:    half, n: 3, k: 3,
			wantCumul: "1.0",
			wantTerms: 4,
		},
		{
			name: "n < k returns error",
			p:    half, n: 2, k: 3,
			wantErr: true,
		},
		{
			name: "k < 0 returns error",
			p:    half, n: 3, k: -1,
			wantErr: true,
		},
		{
			name: "p > 1 returns error",
			p:    bu.StrToFloat("1.1"), n: 3, k: 1,
			wantErr: true,
		},
		{
			name: "p < 0 returns error",
			p:    bu.StrToFloat("-0.1"), n: 3, k: 1,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cumulative, terms, err := CumulativeBinomialProbability(tt.p, tt.n, tt.k)
			if (err != nil) != tt.wantErr {
				t.Fatalf("CumulativeBinomialProbability() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if len(terms) != tt.wantTerms {
				t.Errorf("expected %d terms, got %d", tt.wantTerms, len(terms))
			}
			if compare := bu.NewCompare(&cumulative, tt.wantCumul); !compare.Equal() {
				t.Errorf("CumulativeBinomialProbability() = %v, want %v", compare.ActualAsString, compare.Expected)
			}
		})
	}
}

func Test_CumulativeBinomialProbabilityTermValues(t *testing.T) {
	_, terms, err := CumulativeBinomialProbability(bu.StrToFloat("0.5"), 3, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if compare := bu.NewCompare(&terms[0], "0.125"); !compare.Equal() {
		t.Errorf("terms[0] = %v, want 0.125", compare.ActualAsString)
	}
	if compare := bu.NewCompare(&terms[1], "0.375"); !compare.Equal() {
		t.Errorf("terms[1] = %v, want 0.375", compare.ActualAsString)
	}
}
