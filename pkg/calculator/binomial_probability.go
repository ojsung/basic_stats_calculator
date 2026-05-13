package calculator

import (
	"errors"
	"math/big"

	bu "github.com/ojsung/basic_stats_calculator/internal/big_utils"
)

func CalculateBinomialProbability(chanceOfSuccess *big.Float, trials int64, successes int64) (probability big.Float, err error) {
	if trials < successes {
		return big.Float{}, errors.New("binomial probability trials (n) cannot be less than successes (k)")
	}
	coefficient, err := calculateBinomialCoefficient(trials, successes)
	if err != nil {
		return big.Float{}, err
	}
	pOfKSuccesses, err := calculateProbabilityOfKSuccesses(chanceOfSuccess, trials, successes)
	if err != nil {
		return big.Float{}, err
	}
	coeffAsFloat := bu.PrecFloat().SetInt(coefficient)
	return *bu.PrecFloat().Mul(coeffAsFloat, pOfKSuccesses), nil
}

func calculateProbabilityOfKSuccesses(chanceOfSuccess *big.Float, trials int64, successes int64) (*big.Float, error) {
	if chanceOfSuccess.Cmp(bu.StrToFloat("0")) < 0 || chanceOfSuccess.Cmp(bu.StrToFloat("1")) > 0 {
		return nil, errors.New("probability of k successes chance of success (p) must be between 0 and 1")
	}
	oneMinusP := bu.PrecFloat().Sub(bu.StrToFloat("1"), chanceOfSuccess)
	pPowK := IntPow(chanceOfSuccess, big.NewInt(successes))
	qPowNMinusK := IntPow(oneMinusP, big.NewInt(trials-successes))
	return bu.PrecFloat().Mul(pPowK, qPowNMinusK), nil
}

func calculateBinomialCoefficient(trials, successes int64) (coefficient Int, err error) {
	if trials < successes {
		return nil, errors.New("binomial coefficient trials (n) cannot be less than successes (k)")
	}
	numerator, numerError := Factorial(trials)
	if numerError != nil {
		return nil, numerError
	}
	denominatorMultiplicand, denomMultiplicandError := Factorial(trials - successes)
	if denomMultiplicandError != nil {
		return nil, denomMultiplicandError
	}
	denominatorMultiplier, denomMultiplierError := Factorial(successes)
	if denomMultiplierError != nil {
		return nil, denomMultiplierError
	}
	return numerator.Div(numerator, denominatorMultiplicand.Mul(denominatorMultiplicand, denominatorMultiplier)), nil
}

type Int = *big.Int
