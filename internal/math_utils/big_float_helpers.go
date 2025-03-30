package math_utils

import (
	"errors"
	"fmt"
	"math/big"
)
type binaryExp struct {
	r *big.Float
	a *big.Float
	x *big.Int
}
func (b binaryExp) String() string {
	return fmt.Sprintf("r: %v\na: %v\nx: %v\n", b.r, b.a, b.x)
}

func (a binaryExp) Equal(b binaryExp) (isEqual bool, failed binaryExpComparison) {
	compareA := a.a.Cmp(b.a) == 0
	compareR := a.r.Cmp(b.r) == 0
	compareX := a.x.Cmp(b.x) == 0
	var binComparison binaryExpComparison
	if !compareA {
		binComparison.a = comparison[*big.Float]{
			own:   a.a,
			other: b.a,
		}
	}
	if !compareR {
		binComparison.r = comparison[*big.Float]{
			own: a.r,
			other: b.r,
		}
	}
	if !compareX {
		binComparison.x = comparison[*big.Int]{
			own: a.x,
			other: b.x,
		}
	}
	return a.a.Cmp(b.a) == 0 && a.r.Cmp(b.r) == 0 && a.x.Cmp(b.x) == 0, binComparison
}

type comparison[T any] struct {
	own T
	other T
}
type binaryExpComparison struct {
	r comparison[*big.Float]
	a comparison[*big.Float]
	x comparison[*big.Int]
}

const Ln2String = "0.6931471805599453"
var zero *big.Int = big.NewInt(0)
var one *big.Int = big.NewInt(1)
var two *big.Int = big.NewInt(2)
var floatTwo *big.Float = new(big.Float).SetInt64(2)

func BigFloatFromString(value string) *big.Float {
	f, _ := precFloat().SetString(value)
	return f
}

func Pow(base *big.Float, exponent *big.Int) (power *big.Float) {
	composition := binaryExp{
		BigFloatFromString("1"),
		base.SetPrec(128),
		exponent,
	}
	i := new(big.Int).Set(composition.x)
	for i.Cmp(zero) == 1 {
		if (new(big.Int).Mod(i, two).Cmp(one) == 0) {
			composition = binaryPowOdd(composition)
		} else {
			composition = binaryPowEven(composition)
		}
		i = composition.x
	}

	return normalizeReturn(composition.r)
}

func binaryPowEven(binaryPow binaryExp) (evenPow binaryExp) {
	return binaryExp{
		binaryPow.r,
		precFloat().Mul(binaryPow.a, binaryPow.a),
		new(big.Int).Div(binaryPow.x, two),
	}
}

func binaryPowOdd(binaryPow binaryExp) (oddPow binaryExp) {
	r := precFloat().Mul(binaryPow.a, binaryPow.r)
	return binaryPowEven(binaryExp{
		r,
		binaryPow.a,
		binaryPow.x,
	})
}

func precFloat() *big.Float {
	return new(big.Float).SetPrec(256).SetMode(big.ToNearestEven)
}

func normalizeReturn(value *big.Float) *big.Float {
	return value.SetMode(big.ToZero).SetPrec(128)
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
	if precFloat().Abs(precFloat().Sub(argument, precFloat().SetInt64(1))).Cmp(precFloat().SetInt64(1)) < 1 {
		return nil, fmt.Errorf("taylor expansion appx of ln only converges for real values (s + 1), where s <= 1, received %v", argument)
	}
	// If a function f(x) is infinitely differentiable at a point (a), its approximation, given by the Taylor expansion series, about that point is
	// given by f(x) = (f(a) * (x-a)^0)/0! + (f'(a) * (x-a)^1)/1! + (f''(a) * (x-a)^2)/2! + (f'''(a) * (x-a)^3)/3! + ... + Σ [(-1)^(n+1) * (x - 1)^n] / n, n=1 to ∞
	// For f(x) = ln(x), f'(x) = (1/x), f''(x) = -(1/x^2), etc..
	
	return
}
