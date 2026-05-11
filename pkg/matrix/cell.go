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

func (m Cell[T]) Cmp(value Cell[T]) (comparison int) {
	return m.Operand.Cmp(value.Operand)
}

func (m Cell[T]) String() string {
	return fmt.Sprintf("{ Value: %v, Column: %v, Row: %v }", m.Value, m.Column, m.Row)
}
