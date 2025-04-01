package calculator

import (
	"errors"
	"math/big"

)

func CalculateBinomialProbability(chanceOfSuccess *big.Float, trials int64, successes int64) (probability big.Float, err error) {
	if trials < successes {
		return big.Float{}, errors.New("binomial probability trials (n) cannot be less than successes (k)")
	}
	return
}

func calculateProbabilityOfKSuccesses(chanceOfSuccess *big.Float, trials int64, successes int64) (probability big.Float, err error) {
	if trials < successes {
		return big.Float{}, errors.New("probability of k successes trials (n) cannot be less than successes(k)")
	}
	if successes == 0 {

	}
	return
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
