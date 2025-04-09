package matrix

import (
	"fmt"
	op "github.com/ojsung/basic_stats_calculator/pkg/operand"
)


type Cell[T op.Number | op.BigNumber] struct {
	op.Operand[T]
	Row    int
	Column int
}

func (m Cell[T]) Add(summand Cell[T]) (sum Cell[T]) {
	m.Operand = m.Operand.Add(summand.Operand)
	return m
}

func (m Cell[T]) Sub(subtrahend Cell[T]) (difference Cell[T]) {
	m.Operand = m.Operand.Sub(subtrahend.Operand)
	return m
}

func (m Cell[T]) Mul(multiplier Cell[T]) (product Cell[T]) {
	m.Operand = m.Operand.Mul(multiplier.Operand)
	return m
}

func (m Cell[T]) Div(divisor Cell[T]) (quotient Cell[T]) {
	m.Operand = m.Operand.Div(divisor.Operand)
	return m
}

func (m Cell[T]) AddValue(summand T) (sum Cell[T]) {
	m.Operand = m.Operand.AddValue(summand)
	return m
}

func (m Cell[T]) SubValue(subtrahend T) (difference Cell[T]) {
	m.Operand = m.Operand.SubValue(subtrahend)
	return m
}

func (m Cell[T]) MulValue(multiplier T) (product Cell[T]) {
	m.Operand = m.Operand.MulValue(multiplier)
	return m
}

func (m Cell[T]) DivValue(divisor T) (quotient Cell[T]) {
	m.Operand = m.Operand.DivValue(divisor)
	return m
}

func (m Cell[T]) Zero() (zero Cell[T]) {
	m.Operand = m.Operand.Zero()
	return m
}

func (m Cell[T]) Identity() (identity Cell[T]) {
	m.Operand = m.Operand.Identity()
	return m
}

func (m Cell[T]) Cmp(value Cell[T]) (comparison int) {
	return m.Operand.Cmp(value.Operand)
}

func (m Cell[T]) String() string {
	return fmt.Sprintf("{ Value: %v, Column: %v, Row: %v }", m.Value, m.Column, m.Row)
}
