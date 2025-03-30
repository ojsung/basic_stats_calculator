package math_utils

import (
	"errors"
	"fmt"
	"math/big"
)

func BigFloatFromString(value string) *big.Float {
	f, _ := new(big.Float).SetString(value)
	return f
}

func Pow(base *big.Float, exponent *big.Float) (power *big.Float, err error) {
	return
}

func NaturalPow(exponent *big.Float) (power *big.Float, err error) {
	return
}

func Ln(argument *big.Float) (logarithm *big.Float, err error) {
	return
}

func taylorApproximationLn(argument *big.Float) (logarithm *big.Float, err error) {
	if argument.Sign() <= 0 {
		return nil, errors.New("log is only defined for numbers greater than 0")
	}
	if new(big.Float).Abs(new(big.Float).Sub(argument, new(big.Float).SetInt64(1))).Cmp(new(big.Float).SetInt64(1)) < 1 {
		return nil, errors.New(fmt.Sprintf("Taylor expansion appx of ln only converges for values (s + 1), s <= 1, received %v", argument))
	}
	return
}
