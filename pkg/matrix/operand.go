package matrix

import (
	"fmt"
	"math/big"

	bu "github.com/ojsung/basic_stats_calculator/internal/big_utils"
)


type Operand[T Number | BigNumber] struct {
	Value T
}

func NewOperand[T Number | BigNumber](value T) Operand[T] {
	return Operand[T]{
		Value: value,
	}
}

func (m Operand[T]) String() string {
	var typeString string
	switch any(m.Value).(type) {
	case int:
		typeString = "int"
	case float64:
		typeString = "float64"
	case *big.Float:
		typeString = "*big.Float"
	case *big.Int:
		typeString = "*big.Int"
	}
	return fmt.Sprintf("Operand[%v]{%v}", typeString, m.Value)
}

func (m Operand[T]) FromInt(integer int) (operand Operand[T]) {
	switch v := any(m.Value).(type) {
	case int:
		m.Value = any(integer).(T)
	case float64:
		m.Value = any(float64(integer)).(T)
	case *big.Int:
		m.Value = any(big.NewInt(int64(integer))).(T)
	case *big.Float:
		var prec uint
		if v == nil {
			println("passed nil as Value for Operand[*big.Float]. Using default precision of 64 bits")
			prec = 64
		} else {
			prec = v.Prec()
		}
		m.Value = any(new(big.Float).SetPrec(prec).SetInt64(int64(integer))).(T)
	}
	return m
}

func (m Operand[T]) AddValue(summand T) (sum Operand[T]) {
	switch v := any(m.Value).(type) {
	case int:
		m.Value = any(any(m.Value).(int) + any(summand).(int)).(T)
	case float64:
		m.Value = any(any(m.Value).(float64) + any(summand).(float64)).(T)
	case *big.Float:
		var prec uint
		if v == nil {
			println("adding to nil for Operand[*big.Float]. Using default value 0 with 52 bit precision")
			prec = 52
			m.Value = any(bu.PrecFloat(prec).Add(bu.PrecFloat(prec).SetInt64(0), any(summand).(*big.Float))).(T)
		} else {
			prec = v.Prec()
			m.Value = any(bu.PrecFloat(prec).Add(any(m.Value).(*big.Float), any(summand).(*big.Float))).(T)
		}
	case *big.Int:
		m.Value = any(big.NewInt(0).Add(any(m.Value).(*big.Int), any(summand).(*big.Int))).(T)
	default:
		panic("unsupported type for Add")
	}
	return m
}

func (m Operand[T]) Add(summand Operand[T]) (sum Operand[T]) {
	m = m.AddValue(summand.Value)
	return m
}

func (m Operand[T]) SubValue(subtrahend T) (difference Operand[T]) {
	switch v := any(m.Value).(type) {
	case int:
		m.Value = any(any(m.Value).(int) - any(subtrahend).(int)).(T)
	case float64:
		m.Value = any(any(m.Value).(float64) - any(subtrahend).(float64)).(T)
	case *big.Float:
		var prec uint
		if v == nil {
			println("subtracting from nil for Operand[*big.Float]. Using default value 0 with 52 bit precision")
			prec = 52
			m.Value = any(bu.PrecFloat(prec).Sub(bu.PrecFloat(prec).SetInt64(0), any(subtrahend).(*big.Float))).(T)
		} else {
			prec = v.Prec()
			m.Value = any(bu.PrecFloat(prec).Sub(any(m.Value).(*big.Float), any(subtrahend).(*big.Float))).(T)
		}
	case *big.Int:
		m.Value = any(big.NewInt(0).Sub(any(m.Value).(*big.Int), any(subtrahend).(*big.Int))).(T)
	}
	return m
}

func (m Operand[T]) Sub(subtrahend Operand[T]) (difference Operand[T]) {
	m = m.SubValue(subtrahend.Value)
	return m
}

func (m Operand[T]) MulValue(multiplier T) (product Operand[T]) {
	switch v := any(m.Value).(type) {
	case int:
		m.Value = any(any(m.Value).(int) * any(multiplier).(int)).(T)
	case float64:
		m.Value = any(any(m.Value).(float64) * any(multiplier).(float64)).(T)
	case *big.Float:
		var prec uint
		if v == nil {
			println("multiplying on nil for Operand[*big.Float]. Using default value 0 with 52 bit precision")
			prec = 52
			m.Value = any(bu.PrecFloat(prec).SetInt64(0)).(T)
		} else {
			prec = v.Prec()
			m.Value = any(bu.PrecFloat(prec).Mul(any(m.Value).(*big.Float), any(multiplier).(*big.Float))).(T)
		}
	case *big.Int:
		m.Value = any(big.NewInt(1).Mul(any(m.Value).(*big.Int), any(multiplier).(*big.Int))).(T)
	}
	return m
}

func (m Operand[T]) Mul(multiplier Operand[T]) (product Operand[T]) {
	m = m.MulValue(multiplier.Value)
	return m
}

func (m Operand[T]) DivValue(divisor T) (quotient Operand[T]) {
	switch v := any(m.Value).(type) {
	case int:
		m.Value = any(any(m.Value).(int) / any(divisor).(int)).(T)
	case float64:
		m.Value = any(any(m.Value).(float64) / any(divisor).(float64)).(T)
	case *big.Float:
		var prec uint
		if v == nil {
			println("dividing on nil for Operand[*big.Float]. Using default value 0 with 52 bit precision")
			prec = 52
			m.Value = any(bu.PrecFloat(prec).SetInt64(0)).(T)
		} else {
			prec = v.Prec()
			m.Value = any(bu.PrecFloat(prec).Quo(any(m.Value).(*big.Float), any(divisor).(*big.Float))).(T)
		}
	case *big.Int:
		m.Value = any(big.NewInt(1).Div(any(m.Value).(*big.Int), any(divisor).(*big.Int))).(T)
	}
	return m
}

func (m Operand[T]) Div(divisor Operand[T]) (quotient Operand[T]) {
	m = m.DivValue(divisor.Value)
	return m
}

func (m Operand[T]) Zero() (zero Operand[T]) {
	switch v := any(m.Value).(type) {
	case int:
		return Operand[T]{Value: any(int(0)).(T)}
	case float64:
		return Operand[T]{Value: any(float64(0.0)).(T)}
	case *big.Float:
		var prec uint
		if v == nil {
			println("called for zero value on nil for Operand[*big.Float]. Using default precision")
			prec = 52
		} else {
			prec = v.Prec()
		}
		return Operand[T]{Value: any(bu.PrecFloat(prec).SetInt64(0)).(T)}
	case *big.Int:
		return Operand[T]{Value: any(big.NewInt(0)).(T)}
	}
	return
}

func (m Operand[T]) Identity() (identity Operand[T]) {
	switch v := any(m.Value).(type) {
	case int:
		return Operand[T]{Value: any(int(1)).(T)}
	case float64:
		return Operand[T]{Value: any(float64(1.0)).(T)}
	case *big.Float:
		var prec uint
		if v == nil {
			println("called for identity on nil for Operand[*big.Float]. Using default precision")
			prec = 52
		} else {
			prec = v.Prec()
		}
		return Operand[T]{Value: any(bu.PrecFloat(prec).SetInt64(1)).(T)}
	case *big.Int:
		return Operand[T]{Value: any(big.NewInt(1)).(T)}
	}
	return
}

func (m Operand[T]) Cmp(value Operand[T]) (comparison int) {
	switch any(m.Value).(type) {
	case int:
		if m.Value == value.Value {
			return 0
		} else if any(m.Value).(int) > any(value.Value).(int) {
			return 1
		} else {
			return -1
		}
	case float64:
		if m.Value == value.Value {
			return 0
		} else if any(m.Value).(float64) > any(value.Value).(float64) {
			return 1
		} else {
			return -1
		}
	case *big.Float:
		return any(m.Value).(*big.Float).Cmp(any(value.Value).(*big.Float))
	case *big.Int:
		return any(m.Value).(*big.Int).Cmp(any(value.Value).(*big.Int))
	}
	return
}

func ToFloat[T Number | BigNumber, U FloatNumber](operand Operand[T]) (Operand[U]) {
	var value U
	switch v := any(operand.Value).(type) {
	case int:
		value = any(float64(v)).(U) 
	case *big.Int:
		value = any(new(big.Float).SetInt(v)).(U)
	case float64, *big.Float:
		value = v.(U)
	}
	return Operand[U]{Value: value}
}