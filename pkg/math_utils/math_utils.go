package math_utils

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ojsung/basic_stats_calculator/internal/big_helpers"
)

const Ln2String = "0.6931471805599453"
const EulerString =  "2.7182818284590452"

var zero *big.Int = big.NewInt(0)
var one *big.Int = big.NewInt(1)
var two *big.Int = big.NewInt(2)
var floatTwo *big.Float = new(big.Float).SetInt64(2)

func Pow(base *big.Float, exponent *big.Int) (power *big.Float) {
	composition := big_helpers.BinaryExp{
		R: big_helpers.BigFloatFromString("1"),
		A: base,
		X: exponent,
	}
	i := new(big.Int).Set(composition.X)
	for i.Cmp(zero) == 1 {
		if (new(big.Int).Mod(i, two).Cmp(one) == 0) {
			composition = binaryPowOdd(composition)
		} else {
			composition = binaryPowEven(composition)
		}
		i = composition.X
	}

	return composition.R
}

func binaryPowEven(binaryPow big_helpers.BinaryExp) (evenPow big_helpers.BinaryExp) {
	return big_helpers.BinaryExp{
		R: binaryPow.R,
		A: binaryPow.A.Mul(binaryPow.A, binaryPow.A),
		X: binaryPow.X.Div(binaryPow.X, two),
	}
}

func binaryPowOdd(binaryPow big_helpers.BinaryExp) (oddPow big_helpers.BinaryExp) {
	binaryPow.R.Mul(binaryPow.A, binaryPow.R)
	return binaryPowEven(binaryPow)
}


func NaturalPow(exponent *big.Float) (power *big.Float, err error) {
	return
}

func Ln(argument *big.Float) (logarithm *big.Float, err error) {
	return
}

func Euler(decimals ...uint16) (eulersNumber *big.Float, err error) {
	var places uint16
	if len(decimals) > 0 {
		places = decimals[0]
	} else {
		places = 16
	}
	if places == 0 {
		return big_helpers.PrecFloat().SetInt64(3), nil
	}
	minTerm, minTermRead := new(big.Int).SetString(strings.Join([]string{"1", strings.Repeat("0", int(places))}, ""), 10)
	if !minTermRead {
		return nil, errors.New("failed to create minimum term size from provided decimal places")
	}
	// The series to approximate Euler's number (e), is given by
	// e ~= 1 + 1 + 1/(2!) + 1/(3!) + ... + 1/(n!)
	eulersNumber = big_helpers.PrecFloat().SetInt64(0)
	one := big_helpers.PrecFloat().SetInt64(1)
	for i := int64(0); true; i++ {
		numerator := one
		denominator, err := Factorial(i)
		if err != nil {
			return nil, err
		}
		term := big_helpers.PrecFloat().Quo(numerator, big_helpers.PrecFloat().SetInt(denominator))
		eulersNumber.Add(eulersNumber, term)
		if denominator.Cmp(minTerm) > -1 {
			break;
		}
	}
	return
}

func taylorApproximationLn(argument *big.Float, maxIterations ...int64) (logarithm *big.Float, err error) {
	if zero := big_helpers.PrecFloat().SetInt64(0); argument.Cmp(zero) == -1 {
		return nil, errors.New("argument must be a positive, real number")
	}
	if zero := big_helpers.PrecFloat().SetInt64(0); argument.Cmp(zero) == 0 {
		return nil, errors.New("log is undefined at zero")
	}
	if two := big_helpers.PrecFloat().SetInt64(2); argument.Cmp(two) == 1 {
		return nil, errors.New("taylor approximation of natural log diverges for values greater than 2")
	}
	// Because our taylor appx is for ln(x+1), we have to mutate our argument. Don't mutate the original, copy it
	adjArgument := big_helpers.PrecFloat().Sub(argument, big_helpers.BigFloatFromString("1"))
	var iterations int64;
	if len(maxIterations) > 0 {
		iterations = maxIterations[0]
	} else {
		iterations = 500
	}
	if iterations < 1 {
		return nil, errors.New("maxIterations must be at least 1")
	}
	// If a function f(x) is infinitely differentiable at a point (a), its Taylor expansion series approximation about that point is
	// given by f(x) = (f(a) * (x-a)^0)/0! + (f'(a) * (x-a)^1)/1! + (f''(a) * (x-a)^2)/2! + (f'''(a) * (x-a)^3)/3! + ...
	// We center our taylor series around 0 (or Maclaurin series for this special case), giving us
	// f(x) = f(0) + (f'(0) * x)/1! + (f''(0) * x^2)/2! + ... 
	// For f(a) = ln(a + 1), f'(a) = 1/(a + 1), f''(a) = -1/(a + 1)^2, f'''(a) = 2 / (a + 1)^3, f''''(a) = -6 / (a + 1)^4...
	// For the derivatives at a = 0, we get f'(0) = 1, f''(0) = -1, f'''(0) = 2, f''''(0) = -6...
	// Then, plugging those back into our Taylor series, we get f(x) = ln(1 + x) = 0 + x - 1 * (x^2/2!) + 2 * (x^3/3!) - 6 * (x^4/4!) ...
	// Which we can re-write as f(x) = ln(1 + x) = 0 + x - (1/2)x^2 + ((1 * 2)/(1 * 2 * 3))x^3 + ((1 * 2 * 3)/(1 * 2 * 3 * 4))x^4 ...
	// When we cancel the terms across the numerators and denominators, we are left with x - (x^2)/2 + (x^3)/3 - (x^4)/4 + ... + (x^n)/n,
	// Which we can easily write as a for loop
	logarithm = big_helpers.BigFloatFromString("0")
	checkStart := min(iterations / 3, 120)
	for i := int64(1); i < iterations; i++ {
		numerator := Pow(adjArgument, big.NewInt(i))
		denominator := big_helpers.PrecFloat().SetInt64(i)
		term := big_helpers.PrecFloat().Quo(numerator, denominator)
		var original *big.Float;
		if (i > checkStart) {
			original = big_helpers.PrecFloat().Copy(logarithm)
		}
		if (i % 2 == 1) {
			logarithm.Sub(logarithm, term)
		} else {
			logarithm.Add(logarithm, term)
		}
		// After the first twenty iterations, start checking if we've hit the limit of our precision. Break if we did
		if i > checkStart && original.Cmp(logarithm) == 0 {
			fmt.Sprintln("Stopping at", logarithm)
			break
		}
	}
	return logarithm, nil
}

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