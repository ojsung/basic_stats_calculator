package math_utils

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	bh "github.com/ojsung/basic_stats_calculator/internal/big_helpers"
)

var zero *big.Int = big.NewInt(0)
var one *big.Int = big.NewInt(1)
var two *big.Int = big.NewInt(2)
var lnCache map[string]string = make(map[string]string)
var eCache map[uint16]string = make(map[uint16]string)
var taylorCache map[string]string = map[string]string{"2.0000000000000000": "0.6931471805599453"}

func IntPow(base *big.Float, exponent *big.Int) (power *big.Float) {
	composition := bh.BinaryExp{
		R: bh.StrToFloat("1"),
		A: bh.PrecFloat().Copy(base),
		X: new(big.Int).Set(exponent),
	}
	i := new(big.Int).Set(composition.X)
	for i.Cmp(zero) == 1 {
		if new(big.Int).Mod(i, two).Cmp(one) == 0 {
			composition = binaryPowOdd(composition)
		} else {
			composition = binaryPowEven(composition)
		}
		i = composition.X
	}

	return composition.R
}

func binaryPowEven(binaryPow bh.BinaryExp) (evenPow bh.BinaryExp) {
	return bh.BinaryExp{
		R: binaryPow.R,
		A: binaryPow.A.Mul(binaryPow.A, binaryPow.A),
		X: binaryPow.X.Div(binaryPow.X, two),
	}
}

func binaryPowOdd(binaryPow bh.BinaryExp) (oddPow bh.BinaryExp) {
	binaryPow.R.Mul(binaryPow.A, binaryPow.R)
	return binaryPowEven(binaryPow)
}

func FloatPow(base *big.Float, exponent *big.Float) (power *big.Float, err error) {
	if zero := bh.StrToFloat("0"); base.Cmp(zero) == 0 {
		return zero, nil
	}
	if one := bh.StrToFloat("1"); base.Cmp(one) == 0 {
		return one, nil
	}
	if base.Sign() == -1 {
		return nil, errors.New("power functions are not uniformly defined for negative bases. For negative bases with integer exponents, use IntPow")
	}
		
	// a^n = a^n => ln(a^n) = n * ln(a) => e^ln(a^n) = e^(n * ln(a)) => a^n = e^(n * ln(a))
	lnBase, err := Ln(base)
	if err != nil {
		return nil, err
	}
	nDotLnBase := bh.PrecFloat().Mul(exponent, lnBase)
	power = Exp(nDotLnBase)
	return
}

func Exp(exponent *big.Float, maxIterations ...int64) *big.Float {
	zero := bh.PrecFloat().SetInt64(0)
	one := bh.PrecFloat().SetInt64(1)
	if exponent.Cmp(zero) == 0 {
		return one
	}
	iterations := int64(200)
	if len(maxIterations) > 0 {
		iterations = maxIterations[0]
	}
	checkStart := iterations / 3
	// 1 + x + x^2/2! + x^3/3! + ...
	power := bh.StrToFloat("0")
	for i := int64(0); i < iterations; i++ {
		var original *big.Float
		if i >= checkStart {
			original = power
		}
		numerator := IntPow(exponent, big.NewInt(i))
		denominator, _ := Factorial(i)
		term := numerator.Quo(numerator, bh.PrecFloat().SetInt(denominator))
		power.Add(power, term)
		if i >= checkStart && original.Cmp(power) == 0 {
			break
		}
	}
	return power
}

func Ln(argument *big.Float) (logarithm *big.Float, err error) {
	if value, ok := lnCache[bh.ToStr(argument)]; ok {
		return bh.StrToFloat(value), nil
	}
	if argument.Cmp(bh.StrToFloat("0")) == 0 {
		return nil, errors.New("natural log is not defined at 0")
	}
	if argument.Cmp(bh.StrToFloat("0")) == -1 {
		return nil, errors.New("argument of natural log must be a positive, real number")
	}
	if argument.Cmp(bh.StrToFloat("2")) <= 0 {
		return taylorApproximationLn(argument)
	} else {
		mantissa := bh.PrecFloat()
		exp := argument.MantExp(mantissa)
		// x = mantissa * 2^exp
		// ln(x) = ln(mantissa * 2^exp)
		// ln(x) = ln(mantissa) + exp * (ln(2))
		lnMantissa, lnMantissaErr := Ln(mantissa)
		if lnMantissaErr != nil {
			return nil, fmt.Errorf("error calculating natural log of mantissa: %v", mantissa)
		}
		ln2, ln2Err := Ln(bh.StrToFloat("2"))
		if ln2Err != nil {
			return nil, fmt.Errorf("error calculating natural log of 2")
		}
		expDotLn2 := bh.PrecFloat().Mul(bh.PrecFloat().SetInt64(int64(exp)), ln2)
		logarithm = bh.PrecFloat().Add(lnMantissa, expDotLn2)
		lnCache[bh.ToStr(argument)] = bh.ToStr(logarithm)
		return
	}
}

func Euler(decimals ...uint16) (eulersNumber *big.Float) {
	var places uint16
	if len(decimals) > 0 {
		places = decimals[0]
	} else {
		places = 16
	}
	if places == 0 {
		return bh.PrecFloat().SetInt64(3)
	}
	if value, ok := eCache[places]; ok {
		return bh.StrToFloat(value)
	}
	minTerm, _ := new(big.Int).SetString(strings.Join([]string{"1", strings.Repeat("0", int(places))}, ""), 10)
	// The series to approximate Euler's number (e), is given by
	// e ~= 1 + 1 + 1/(2!) + 1/(3!) + ... + 1/(n!)
	eulersNumber = bh.PrecFloat().SetInt64(0)
	one := bh.PrecFloat().SetInt64(1)
	for i := int64(0); true; i++ {
		numerator := one
		denominator, _ := Factorial(i)
		term := bh.PrecFloat().Quo(numerator, bh.PrecFloat().SetInt(denominator))
		eulersNumber.Add(eulersNumber, term)
		if denominator.Cmp(minTerm) > -1 {
			break
		}
	}
	if len(eCache) > 10 {
		clear(eCache)
	}
	eCache[places] = bh.ToStr(eulersNumber)
	return
}

func taylorApproximationLn(argument *big.Float, maxIterations ...int64) (logarithm *big.Float, err error) {
	if zero := bh.PrecFloat().SetInt64(0); argument.Cmp(zero) == -1 {
		return nil, errors.New("argument must be a positive, real number")
	}
	if zero := bh.PrecFloat().SetInt64(0); argument.Cmp(zero) == 0 {
		return nil, errors.New("log is undefined at zero")
	}
	if two := bh.PrecFloat().SetInt64(2); argument.Cmp(two) == 1 {
		return nil, errors.New("taylor approximation of natural log diverges for values greater than 2")
	}
	if value, ok := taylorCache[bh.ToStr(argument)]; ok {
		return bh.StrToFloat(value), nil
	}
	// Because our taylor appx is for ln(x+1), we have to mutate our argument. Don't mutate the original, copy it
	adjArgument := bh.PrecFloat().Sub(argument, bh.StrToFloat("1"))
	var iterations int64
	if len(maxIterations) > 0 {
		iterations = maxIterations[0]
	} else {
		iterations = determineLeastIterations(argument)
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
	logarithm = bh.StrToFloat("0")
	checkStart := min(iterations/3, 120)
	for i := int64(1); i < iterations; i++ {
		numerator := IntPow(adjArgument, big.NewInt(i))
		denominator := bh.PrecFloat().SetInt64(i)
		term := bh.PrecFloat().Quo(numerator, denominator)
		var original *big.Float
		if i > checkStart {
			original = bh.PrecFloat().Copy(logarithm)
		}
		if i%2 == 0 {
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
	taylorCache[bh.ToStr(argument)] = bh.ToStr(logarithm)
	return logarithm, nil
}

func determineLeastIterations(argument *big.Float) int64 {
	abs := bh.PrecFloat().Abs(argument)
	if abs.Cmp(bh.StrToFloat("0.8")) < 0 {
		// y = a + be^(cx)
		a := bh.StrToFloat("4.63")
		b := bh.StrToFloat("8.28")
		c := bh.StrToFloat("3.43")
		cx := c.Mul(c, abs)
		ePowCX := Exp(cx)
		bDotEPowCX := b.Mul(b, ePowCX)
		y := a.Add(a, bDotEPowCX)
		yRoundDown, _ := y.Int(new(big.Int))
		return yRoundDown.Add(big.NewInt(1), yRoundDown).Int64()
	} else {
		// y = a + b/(1-x)^c
		a := bh.StrToFloat("-100.8")
		b := bh.StrToFloat("67.56")
		c := bh.StrToFloat("0.826")
		oneMinusX := bh.PrecFloat().Sub(bh.StrToFloat("1"), abs)
		denominator, _ := FloatPow(oneMinusX, c)
		quotient := b.Quo(b, denominator)
		y := a.Add(a, quotient)
		return RoundUp(y).Int64()
	}
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

func RoundUp(float *big.Float) *big.Int {
	intVal, _ := float.Int(new(big.Int))
	return intVal.Add(intVal, big.NewInt(1))
}
