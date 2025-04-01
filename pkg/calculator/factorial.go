package calculator
import (
	"errors"
	"math/big"

)

func Factorial(argument int64) (factorial *big.Int, err error) {
	if argument < 0 {
		return nil, errors.New("factorial argument (n) cannot be negative")
	}
	// I'm using 1 a lot
	one := big.NewInt(1)
	if argument == 1 || argument == 0 {
		return one, nil
	}
	var carry *big.Int = one
	for i := argument; i > 1; i-- {
		carry.MulRange(2, argument)
	}
	return carry, nil
}
