package calculator

import (
	"math/big"
	"testing"

	bu "github.com/ojsung/basic_stats_calculator/internal/big_utils"
)

func Test_BinomialPValue(t *testing.T) {
	half := bu.StrToFloat("0.5")
	tests := []struct {
		name    string
		p       *big.Float
		n, k    int64
		tail    string
		want    string
		wantErr bool
	}{
		{
			name: "left tail k=1",
			p:    half, n: 3, k: 1, tail: "left",
			want: "0.5",
		},
		{
			name: "right tail k=1",
			p:    half, n: 3, k: 1, tail: "right",
			want: "0.875",
		},
		{
			name: "right tail k=0 returns 1",
			p:    half, n: 3, k: 0, tail: "right",
			want: "1.0",
		},
		{
			name: "two-tailed k=0",
			p:    half, n: 3, k: 0, tail: "two",
			want: "0.25",
		},
		{
			// p=0.9, n=1, k=1: left=1.0, right=0.9, two=2*0.9=1.8 → capped to 1
			name: "two-tailed capped at 1",
			p:    bu.StrToFloat("0.9"), n: 1, k: 1, tail: "two",
			want: "1.0",
		},
		{
			name: "invalid tail returns error",
			p:    half, n: 3, k: 1, tail: "center",
			wantErr: true,
		},
		{
			name: "p > 1 returns error",
			p:    bu.StrToFloat("1.1"), n: 3, k: 1, tail: "left",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pval, err := BinomialPValue(tt.p, tt.n, tt.k, tt.tail)
			if (err != nil) != tt.wantErr {
				t.Fatalf("BinomialPValue() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if compare := bu.NewCompare(&pval, tt.want); !compare.Equal() {
				t.Errorf("BinomialPValue() = %v, want %v", compare.ActualAsString, compare.Expected)
			}
		})
	}
}
