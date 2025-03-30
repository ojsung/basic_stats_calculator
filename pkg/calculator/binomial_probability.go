package calculator

import (
	"errors"
	"math/big"
	mu "github.com/ojsung/basic_stats_calculator/internal/math_utils"
)

func calculateBinomialProbability(chanceOfSuccess *big.Float, trials int64, successes int64) (probability big.Float, err error) {
	if trials < successes {
		return big.Float{}, errors.New("binomial probability trials (n) cannot be less than successes (k)")
	}

}

func calculateProbabilityOfKSuccesses(chanceOfSuccess *big.Float, trials int64, successes int64) (probability big.Float, err error) {
	if trials < successes {
		return big.Float{}, errors.New("probability of k successes trials (n) cannot be less than successes(k)")
	}
	if successes == 0 {
		exponent := new(big.Float).SetInt64(trials)
		base := chanceOfSuccess.
	}

}

func calculateBinomialCoefficient(trials, successes int64) (coefficient Int, err error) {
	if trials < successes {
		return nil, errors.New("binomial coefficient trials (n) cannot be less than successes (k)")
	}
	numerator, numerError := factorial(trials)
	if numerError != nil {
		return nil, numerError
	}
	denominatorMultiplicand, denomMultiplicandError := factorial(trials - successes)
	if denomMultiplicandError != nil {
		return nil, denomMultiplicandError
	}
	denominatorMultiplier, denomMultiplierError := factorial(successes)
	if denomMultiplierError != nil {
		return nil, denomMultiplierError
	}
	return numerator.Div(numerator, denominatorMultiplicand.Mul(denominatorMultiplicand, denominatorMultiplier)), nil
}

type Int = *big.Int

func factorial(argument int64) (factorial Int, err error) {
	if argument < 0 {
		return nil, errors.New("factorial argument (n) cannot be negative")
	}
	// I'm using 1 a lot
	one := big.NewInt(1)
	if argument == 1 || argument == 0 {
		return one, nil
	}
	var carry Int = one
	for i := argument; i > 1; i-- {
		carry.MulRange(2, argument)
	}
	return carry, nil
}
