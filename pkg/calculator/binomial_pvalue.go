package calculator

import (
	"fmt"
	"math/big"

	bu "github.com/ojsung/basic_stats_calculator/internal/big_utils"
)

func BinomialPValue(p *big.Float, n, k int64, tail string) (pValue big.Float, err error) {
	if tail != "left" && tail != "right" && tail != "two" {
		return big.Float{}, fmt.Errorf("tail must be \"left\", \"right\", or \"two\", got %q", tail)
	}
	left, _, err := CumulativeBinomialProbability(p, n, k)
	if err != nil {
		return big.Float{}, err
	}
	if tail == "left" {
		return left, nil
	}
	var right big.Float
	if k == 0 {
		right = *bu.StrToFloat("1")
	} else {
		rightCum, _, err := CumulativeBinomialProbability(p, n, k-1)
		if err != nil {
			return big.Float{}, err
		}
		right = *bu.PrecFloat().Sub(bu.StrToFloat("1"), &rightCum)
	}
	if tail == "right" {
		return right, nil
	}
	var minVal big.Float
	if left.Cmp(&right) <= 0 {
		minVal = left
	} else {
		minVal = right
	}
	two := *bu.PrecFloat().Mul(bu.StrToFloat("2"), &minVal)
	if two.Cmp(bu.StrToFloat("1")) > 0 {
		return *bu.StrToFloat("1"), nil
	}
	return two, nil
}
