package math_utils

import "math/big"
func BigFloatFromString(value string) *big.Float {
	f, _ := new(big.Float).SetString(value)
	return f
}

func Pow(base *big.Float, exponent *big.Float) *big.Float {
	
}

func NaturalPow(exponent *big.Float) *big.Float {
	
}

func ln(argument *big.Float) *big.Float {

}


func taylorApproximationLn(argument *big.Float) (*big.Float, error) {
	if argument.Sign() <= 0 {
		return nil, errors.New("log is only defined for numbers greater than 0")
	}
	if new(big.Float).Abs(new(big.Float).Sub(big.NewInt(1))) > 1 {
		return nil, errors.New("Taylor expansion appx of ln only converges for values (s + 1), s <= 1, received %v", argument)
	}
	
}