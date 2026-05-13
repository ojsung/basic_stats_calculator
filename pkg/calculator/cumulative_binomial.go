package calculator

import (
	"errors"
	"math/big"

	bu "github.com/ojsung/basic_stats_calculator/internal/big_utils"
)

func CumulativeBinomialProbability(p *big.Float, n, k int64) (cumulative big.Float, terms []big.Float, err error) {
	if k < 0 {
		return big.Float{}, nil, errors.New("cumulative binomial probability k cannot be negative")
	}
	if n < k {
		return big.Float{}, nil, errors.New("cumulative binomial probability n cannot be less than k")
	}
	acc := bu.PrecFloat().SetInt64(0)
	terms = make([]big.Float, 0, k+1)
	for i := int64(0); i <= k; i++ {
		prob, probErr := CalculateBinomialProbability(p, n, i)
		if probErr != nil {
			return big.Float{}, nil, probErr
		}
		acc.Add(acc, &prob)
		terms = append(terms, prob)
	}
	return *acc, terms, nil
}
