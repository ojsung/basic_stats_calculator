package calculator

import (
	"math/big"
)

/*
Need to be able to
- Define a matrix
- Find the determinant of the matrix
- Find the cofactor of the matrix
- Find the cofactor matrix of the matrix
- Find the inverse of the matrix
- Do cross-product multiplication
- Do scalar multiplication
- Set row
- Set column
- Add row
- Add column
- Remove row
- Remove column
*/

type Number interface {
	float32 | float64 | int | *big.Float | *big.Int
}

type Matrix[T Number] struct {
	rows           [][]T
	determinant    T
	cofactor       *Matrix[T]
	cofactorMatrix *Matrix[T]
	inverse        *Matrix[T]
}

func (Matrix[T]) ScalarMul(Matrix[T]) (product Matrix[T]) {
	return
}

func (multiplicand Matrix[T]) VectorMul(vectorMultiplier Matrix[T]) (vectorProduct Matrix[T]) {
	return multiplicand.Cross(vectorMultiplier)
}

func (Matrix[T]) Cross(Matrix[T]) (crossProduct Matrix[T]) {
	return
}

func (Matrix[T]) Determinant() (determinant T) {
	return
}

func (Matrix[T]) RemoveColumns(index ...int) (columns [][]T) {
	return
}

func (Matrix[T]) RemoveRows(index ...int) (rows [][]T) {
	return
}

func (Matrix[T]) Rows() (rows [][]T) {
	return
}

func (Matrix[T]) Columns() (columns [][]T) {
	return
}

func (Matrix[T]) GetColumn(index int) (column []T) {
	return
}

func (Matrix[T]) GetRow(index int) (row []T) {
	return
}

func (Matrix[T]) Inverse() (inverse *Matrix[T]) {
	return
}

func (Matrix[T]) Cofactor() (cofactor T) {
	return
}

func (Matrix[T]) CofactorMatrix() (cofactorMatrix *Matrix[T]) {
	return
}

func NewMatrix[T Number](rows []T, columns []T) (matrix *Matrix[T]) {
	return
}
