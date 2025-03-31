package big_helpers

import (
	"fmt"
	"math/big"
)

type BinaryExp struct {
	R *big.Float
	A *big.Float
	X *big.Int
}

func (b BinaryExp) String() string {
	return fmt.Sprintf("r: %v\na: %v\nx: %v\n", b.R, b.A, b.X)
}

func (a BinaryExp) Equal(b BinaryExp) (isEqual bool, failed BinaryExpComparison) {
	compareA := a.A.Cmp(b.A) == 0
	compareR := a.R.Cmp(b.R) == 0
	compareX := a.X.Cmp(b.X) == 0
	var binComparison BinaryExpComparison
	if !compareA {
		binComparison.A = comparison[*big.Float]{
			Own:   a.A,
			Other: b.A,
		}
	}
	if !compareR {
		binComparison.R = comparison[*big.Float]{
			Own:   a.R,
			Other: b.R,
		}
	}
	if !compareX {
		binComparison.X = comparison[*big.Int]{
			Own:   a.X,
			Other: b.X,
		}
	}
	return a.A.Cmp(b.A) == 0 && a.R.Cmp(b.R) == 0 && a.X.Cmp(b.X) == 0, binComparison
}

type comparison[T any] struct {
	Own   T
	Other T
}
type BinaryExpComparison struct {
	R comparison[*big.Float]
	A comparison[*big.Float]
	X comparison[*big.Int]
}


func StrToFloat(value string) *big.Float {
	f, _ := PrecFloat().SetString(value)
	return f
}

func PrecFloat() *big.Float {
	return new(big.Float).SetPrec(256).SetMode(big.ToNearestEven)
}

func NormalizeReturn(value *big.Float) *big.Float {
	return value.SetMode(big.ToZero).SetPrec(128)
}


