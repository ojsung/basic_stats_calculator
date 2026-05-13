package calculator

import (
	"errors"
	"math/big"
)

func Factorial(argument int64) (factorial *big.Int, err error) {
	if argument < 0 {
		return nil, errors.New("factorial argument (n) cannot be negative")
	}
	return new(big.Int).MulRange(1, argument), nil
}
