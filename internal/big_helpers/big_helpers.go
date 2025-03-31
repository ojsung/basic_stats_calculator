package big_helpers

import (
	"math/big"
	"strings"
)

type StrBigCompare[T big.Float | big.Int] struct {
	Actual         *T
	Expected       string
	ActualAsString string
}

func (container StrBigCompare[T]) Equal() bool {
	return container.ActualAsString == container.Expected
}
func NewCompare[T big.Float | big.Int](actual *T, expected string) *StrBigCompare[T] {
	container := StrBigCompare[T]{
		Actual:   actual,
		Expected: expected,
	}
	absStr := strings.Replace(container.Expected, "-", "", 1)
	decimalIndex := strings.Index(absStr, ".")
	strLen := len(absStr)
	if decimalIndex == -1 {
		strLen = 0
	} else {
		strLen -= (decimalIndex + 1)
	}
	switch v := any(container.Actual).(type) {
	case *big.Float:
		container.ActualAsString = v.Text('f', strLen)
	case *big.Int:
		container.ActualAsString = v.Text(10)
	default:
		panic("unsupported type")
	}
	return &container
}

func ToStr[T big.Int | big.Float](value *T, places ...int) string {
	decimals := 16
	if len(places) > 0 {
		decimals = places[0]
	}
	switch v := any(value).(type) {
	case *big.Float:
		return v.Text('f', decimals)
	case *big.Int:
		return v.Text(10)
	default:
		panic("unsupported type")
	}
}
