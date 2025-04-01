package calculator

import (
	"testing"

	bu "github.com/ojsung/basic_stats_calculator/internal/big_utils"
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
			if compare := bu.NewCompare(got, tt.want); !compare.Equal() {
				t.Errorf("factorial() = %v, want %v", compare.ActualAsString, tt.want)
			}
		})
	}
}