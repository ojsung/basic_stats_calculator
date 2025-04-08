package matrix

import (
	"errors"
	"fmt"
	"math/big"
	"slices"
	"strings"

	su "github.com/ojsung/basic_stats_calculator/internal"
)

/*
Need to be able to
- Define a matrix
- Find the determinant of the matrix
- Find the cofactor of the matrix
- Find the cofactor matrix of the matrix
- Find the transpose of the matrix
- Find the inverse of the matrix
- Do matrix multiplication
- Do scalar multiplication
- Set row
- Set column
- Add row
- Add column
- Remove row
- Remove column
*/

type Number interface {
	float64 | int
}

type BigNumber interface {
	*big.Float | *big.Int
}

type FloatNumber interface {
	*big.Float | float64
}

type Matrix[T Number | BigNumber, U FloatNumber] struct {
	cells   *[]Cell[T]
	rows    int
	columns int
}

type NumberMatrix[T Number] struct {
	*Matrix[T, float64]
}
type BigMatrix[T BigNumber] struct {
	*Matrix[T, *big.Float]
}

func (matrix *Matrix[T, U]) reindex() {
	columnIndex := -1
	rowIndex := -1
	rowIndexMap := make(map[int]int)
	columnIndexMap := make(map[int]int)
	for index, cell := range *matrix.cells {
		if value, exists := rowIndexMap[cell.Row]; exists {
			cell.Row = value
		} else {
			rowIndex += 1
			rowIndexMap[cell.Row] = rowIndex
			cell.Row = rowIndex
		}
		if value, exists := columnIndexMap[cell.Column]; exists {
			cell.Column = value
		} else {
			columnIndex += 1
			columnIndexMap[cell.Column] = columnIndex
			cell.Column = columnIndex
		}
		(*matrix.cells)[index] = cell
	}
	matrix.rows = rowIndex + 1
	matrix.columns = columnIndex + 1
}

func (matrix *Matrix[T, U]) sort() {
	slices.SortFunc(*matrix.cells, func(a Cell[T], b Cell[T]) int {
		if a.Row > b.Row || a.Row == b.Row && a.Column > b.Column {
			return 1
		}
		return -1
	})
}

func (matrix Matrix[T, U]) Cells() (cells *[]Cell[T]) {
	return matrix.cells
}

func (m Matrix[T, U]) RemoveRows(indices ...int) (subMatrix *Matrix[T, U], err error) {
	matrixLength := m.rows
	for _, rowIndex := range indices {
		if rowIndex >= matrixLength {
			return nil, fmt.Errorf("index out of range. len %v, index %v", matrixLength, rowIndex)
		}
	}
	cells := make([]Cell[T], len(*m.cells) - len(indices) * m.columns)
	index := 0
	for _, cell := range *m.cells {
		if !slices.Contains(indices, cell.Row) {
			cells[index] = cell
			index++
		}
	}
	subMatrix = &Matrix[T, U]{
		cells: &cells,
	}
	subMatrix.reindex()
	subMatrix.sort()
	return subMatrix, nil
}

func (m Matrix[T, U]) RemoveColumns(indices ...int) (subMatrix *Matrix[T, U], err error) {
	matrixLength := m.columns
	for _, colIndex := range indices {
		if colIndex >= matrixLength {
			return nil, fmt.Errorf("index out of range. len %v, index %v", matrixLength, colIndex)
		}
	}
	cells := make([]Cell[T], len(*m.cells) - len(indices) * m.rows)
	index := 0
	for _, cell := range *m.cells {
		if !slices.Contains(indices, cell.Column) {
			cells[index] = cell
			index++
		}
	}
	subMatrix = &Matrix[T, U] {
		cells: &cells,
	}
	subMatrix.reindex()
	subMatrix.sort()
	return subMatrix, nil
}

func (m *Matrix[T, U]) RemoveColumnsInPlace(indices ...int) (columns map[int][]T) {
	columns = make(map[int][]T)
	for _, index := range indices {
		columns[index] = nil
	}
	*m.cells = slices.DeleteFunc(*m.cells, func(cell Cell[T]) bool {
		if isRemove := slices.Contains(indices, cell.Column); isRemove {
			columns[cell.Column] = append(columns[cell.Column], cell.Value)
			return isRemove
		}
		return false
	})
	m.reindex()
	m.sort()
	return
}

func (m *Matrix[T, U]) RemoveRowsInPlace(indices ...int) (rows map[int][]T) {
	rows = make(map[int][]T)
	for _, index := range indices {
		rows[index] = nil
	}
	*m.cells = slices.DeleteFunc(*m.cells, func(cell Cell[T]) bool {
		if isRemove := slices.Contains(indices, cell.Row); isRemove {
			rows[cell.Row] = append(rows[cell.Row], cell.Value)
			return isRemove
		}
		return false
	})
	m.reindex()
	m.sort()
	return
}

func (m *Matrix[T, U]) ScalarMul(scalar T) *Matrix[T, U] {
	newCells := make([]Cell[T], len(*m.cells))
	for i, cell := range *m.cells {
		newCells[i] = Cell[T]{
			Operand: NewOperand(cell.Operand.MulValue(scalar).Value),
			Row:     cell.Row,
			Column:  cell.Column,
		}
	}
	return &Matrix[T, U]{
		cells:   &newCells,
		rows:    m.rows,
		columns: m.columns,
	}
}

func (m *Matrix[T, U]) FractionalScalarMul(scalar U) (floatMatrix *Matrix[U, U]) {
	newCells := su.Map(*m.cells, func(cell Cell[T]) Cell[U] {
		newCell := Cell[U]{
			Row:     cell.Row,
			Column:  cell.Column,
			Operand: ToFloat[T, U](cell.Operand).MulValue(scalar),
		}
		return newCell
	})
	floatMatrix = &Matrix[U, U]{
		rows:    m.rows,
		columns: m.columns,
		cells:   &newCells,
	}
	*floatMatrix.cells = newCells
	return
}

func (m Matrix[T, U]) Scale(scalar T) (product *Matrix[T, U]) {
	product = m.ScalarMul(scalar)
	return
}

func (matrixA Matrix[T, U]) Mul(matrixB *Matrix[T, U]) (product *Matrix[T, U], err error) {
	if matrixA.columns != matrixB.rows {
		return nil, fmt.Errorf("columns of multiplicand (%v) must match rows of multiplier (%v)", matrixA.columns, matrixB.rows)
	}
	product, _ = NewMatrix[T, U]([][]T{})
	*product.cells = make([]Cell[T], matrixA.rows*matrixB.columns)
	matrixACells := *matrixA.Cells()
	matrixBCells := *matrixB.Cells()
	for index := range *product.cells {
		rowIndex := index / matrixB.columns
		columnIndex := index % matrixB.columns
		cell := Cell[T]{
			Row:     rowIndex,
			Column:  columnIndex,
			Operand: matrixACells[0].Operand.Zero(),
		}
		for k := range matrixA.columns {
			matrixACell := matrixACells[rowIndex*matrixA.columns+k]
			matrixBCell := matrixBCells[matrixB.columns*k+columnIndex]
			cell = cell.Add(matrixACell.Mul(matrixBCell))
		}
		(*product.cells)[index] = cell
	}
	product.sort()
	product.reindex()
	return
}

func (matrixA Matrix[T, U]) Add(matrixB *Matrix[T, U]) (sum *Matrix[T, U], err error) {
	if matrixA.columns != matrixB.columns || matrixA.rows != matrixB.rows {
		return nil, fmt.Errorf("dimensions of matrices must match. MatrixA rows: %v, columns: %v. MatrixB rows: %v, columns: %v", matrixA.rows, matrixA.columns, matrixB.rows, matrixB.columns)
	}
	sumCells := make([]Cell[T], len(*matrixA.cells))
	sum = &Matrix[T, U]{
		cells: &sumCells,
	}
	matrixACells := *matrixA.cells
	matrixBCells := *matrixB.cells
	for index, cell := range matrixACells {
		sumCell := cell.Add(matrixBCells[index])
		(*sum.cells)[index] = sumCell
	}
	sum.reindex()
	return
}

func (matrixA Matrix[T, U]) Sub(matrixB *Matrix[T, U]) (difference *Matrix[T, U], err error) {
	if matrixA.columns != matrixB.columns || matrixA.rows != matrixB.rows {
		return nil, fmt.Errorf("dimensions of matrices must match. MatrixA rows: %v, columns: %v. MatrixB rows: %v, columns: %v", matrixA.rows, matrixA.columns, matrixB.rows, matrixB.columns)
	}
	differenceCells := make([]Cell[T], len(*matrixA.cells))
	difference = &Matrix[T, U]{cells: &differenceCells}
	matrixACells := *matrixA.cells
	matrixBCells := *matrixB.cells
	for index, cell := range matrixACells {
		diffCell := cell.Sub(matrixBCells[index])
		(*difference.cells)[index] = diffCell
	}
	difference.reindex()
	return
}

// Find the determinant by Gaussian Elimination to find the eigenvalues
// Then multiply the eigenvalues together. There is likely a better way to find them programmatically
// But this is what I could think of
// We have to be careful not to do any division for our int or *big.Int values, as it could be an int matrix, and to track any scaling we do of rows
func (m *Matrix[T, U]) Determinant() (determinant Operand[T], err error) {
	if len(*m.cells) == 0 {
		return new(Operand[T]).Zero(), errors.New("matrix must have a length")
	}
	baseValue := (*m.cells)[0].Operand
	zero := baseValue.Zero()
	identity := baseValue.Identity()
	if !m.IsSquare() {
		return zero, errors.New("cannot find determinant for non-square matrix")
	}
	orderedMatrix, sign := m.reorderRowsPerfectMatching()
	if sign == 0 {
		return zero, nil
	}
	triangularized, scale := orderedMatrix.UpperTriangularize()
	determinantOperand := identity
	determinantOperand = determinantOperand.Mul(baseValue.FromInt(sign))
	triangularRows := triangularized.GetRowOperands()
	for index := range triangularized.Rows() {
		determinantOperand = determinantOperand.Mul(triangularRows[index][index])
	}
	determinantOperand = determinantOperand.Div(scale)

	return determinantOperand, nil
}

func (nonZeroDiagonalMatrix *Matrix[T, U]) UpperTriangularize() (triangularizedMatrix *Matrix[T, U], scale Operand[T]) {
	rows := nonZeroDiagonalMatrix.GetRowOperands()
	zero := rows[0][0].Zero()
	scale = rows[0][0].Identity()
	var rowReducer func(rowIndex int, row []Operand[T], nextRow []Operand[T], scale Operand[T]) (reducedRow []Operand[T], determinantScale Operand[T])
	switch any(zero.Value).(type) {
	case *big.Int, int:
		rowReducer = intTriangularRowReducer[T, U]
	case *big.Float, float64:
		rowReducer = floatTriangularRowReducer[T, U]
	}
	lenRows := len(rows)
	operandRows := make([][]Operand[T], lenRows)
	operandRows[0] = rows[0]
	valueRows := make([][]T, lenRows)
	for rowIndex, row := range operandRows {
		for nextRowIndex := rowIndex + 1; nextRowIndex < lenRows; nextRowIndex++ {
			if lenRows > nextRowIndex {
				nextRow := rows[nextRowIndex]
				nextRow, scale = rowReducer(rowIndex, row, nextRow, scale)
				operandRows[nextRowIndex] = nextRow
			}
		}
		valueRow := su.Map(row, func(value Operand[T]) T {
			return value.Value
		})
		valueRows[rowIndex] = valueRow
	}
	triangularizedMatrix, _ = NewMatrix[T, U](valueRows)
	return triangularizedMatrix, scale
}

func intTriangularRowReducer[T Number | BigNumber, U FloatNumber](rowIndex int, row []Operand[T], nextRow []Operand[T], scale Operand[T]) (reducedRow []Operand[T], determinantScale Operand[T]) {
	rowMultiplier := nextRow[rowIndex]
	nextRowMultiplier := row[rowIndex]
	for colIndex, operand := range nextRow {
		nextRow[colIndex] = operand.Mul(nextRowMultiplier).Sub(row[colIndex].Mul(rowMultiplier))
	}
	scale = scale.Mul(nextRowMultiplier)
	return nextRow, scale
}

func floatTriangularRowReducer[T Number | BigNumber, U FloatNumber](rowIndex int, row []Operand[T], nextRow []Operand[T], scale Operand[T]) (reducedRow []Operand[T], determinantScale Operand[T]) {
	ratio := nextRow[rowIndex].Div(row[rowIndex])
	for colIndex, operand := range nextRow {
		nextRow[colIndex] = operand.Sub(row[colIndex].Mul(ratio))
	}
	return nextRow, scale
}

func (m Matrix[T, U]) Trace() (trace T, err error) {
	if len(*m.cells) == 0 {
		return new(Operand[T]).Zero().Value, errors.New("matrix must have a length")
	}
	baseValue := (*m.cells)[0].Operand
	zero := baseValue.Zero()
	if !m.IsSquare() {
		return zero.Value, errors.New("cannot find determinant for non-square matrix")
	}
	traceOperand := zero
	rows := m.GetRowOperands()
	for index := range m.Rows() {
		traceOperand = traceOperand.Add(rows[index][index])
	}
	trace = traceOperand.Value
	return
}

func (m Matrix[T, U]) Rows() (length int) {
	return m.rows
}

func (m Matrix[T, U]) Columns() (length int) {
	return m.columns
}

func (m Matrix[T, U]) IsSquare() (isSquare bool) {
	if len(*m.cells) == 0 {
		return false
	}
	return m.rows == m.columns
}

func (m Matrix[T, U]) GetRows() (rows [][]T) {
	rows = make([][]T, m.rows)
	for _, cell := range *m.cells {
		rows[cell.Row] = append(rows[cell.Row], cell.Value)
	}
	return
}

func (m Matrix[T, U]) GetColumns() (columns [][]T) {
	columns = make([][]T, m.columns)
	for _, cell := range *m.cells {
		columns[cell.Column] = append(columns[cell.Column], cell.Value)
	}
	return
}

func (m Matrix[T, U]) GetRowOperands() (rows [][]Operand[T]) {
	rows = make([][]Operand[T], m.rows)
	for _, cell := range *m.cells {
		rows[cell.Row] = append(rows[cell.Row], cell.Operand)
	}
	return
}

func (m Matrix[T, U]) GetRowCells() (rows [][]Cell[T]) {
	rows = make([][]Cell[T], m.rows)
	for _, cell := range *m.cells {
		rows[cell.Row] = append(rows[cell.Row], cell)
	}
	return
}

func (m Matrix[T, U]) GetColumnOperands() (columns [][]Operand[T]) {
	columns = make([][]Operand[T], m.columns)
	for _, cell := range *m.cells {
		columns[cell.Column] = append(columns[cell.Column], cell.Operand)
	}
	return
}

func (m Matrix[T, U]) GetColumnCells() (columns [][]Cell[T]) {
	columns = make([][]Cell[T], m.columns)
	for _, cell := range *m.cells {
		columns[cell.Column] = append(columns[cell.Column], cell)
	}
	return
}

func (Matrix[T, U]) Get(col int, row int) (value T) {
	return
}

func (m Matrix[T, U]) GetColumn(index int) (column []T, err error) {
	if index > m.columns-1 {
		return nil, fmt.Errorf("out of bounds. Columns: %v, zero-based index: %v", m.columns, index)
	}
	column = make([]T, m.rows)
	for _, cell := range *m.cells {
		if cell.Column == index {
			column[cell.Row] = cell.Value
		}
	}
	return
}

func (m Matrix[T, U]) GetRow(index int) (row []T, err error) {
	if index > m.rows-1 {
		return nil, fmt.Errorf("out of bounds. Rows: %v, zero-based index: %v", m.rows, index)
	}
	row = make([]T, m.columns)
	for _, cell := range *m.cells {
		if cell.Row == index {
			row[cell.Column] = cell.Value
		}
	}
	return
}

func (m Matrix[T, U]) Inverse() (inverse *Matrix[U, U], isSingular bool, err error) {
	if len(*m.cells) == 0 {
		return nil, false, errors.New("cannot find inverse of empty matrix")
	}
	if !m.IsSquare() {
		return nil, false, errors.New("cannot find inverse of non-square matrix")
	}
	base := (*m.cells)[0].Operand
	zero := base.Zero()
	identity := (*m.cells)[0].Operand.Identity()
	determinant, err := m.Determinant()
	if determinant.Cmp(zero) == 0 {
		return nil, true, err
	}
	if err != nil {
		return nil, false, err
	}
	floatIdentity := ToFloat[T, U](identity)
	floatDeterminant := ToFloat[T, U](determinant)
	cofactorMatrix, err := m.CofactorMatrix() 
	if err != nil {
		return nil, false, err
	}
	transpose := cofactorMatrix.Transpose()
	inverse = transpose.FractionalScalarMul(floatIdentity.Div(floatDeterminant).Value)
	return
}

func (m Matrix[T, U]) Transpose() (transpose *Matrix[T, U]) {
	remappedCells := su.Map(*m.cells, func(cell Cell[T]) Cell[T] {
		return Cell[T]{
			Operand: cell.Operand,
			Row:     cell.Column,
			Column:  cell.Row,
		}
	})
	transpose = &Matrix[T, U]{
		rows:    m.rows,
		columns: m.columns,
		cells:   &remappedCells,
	}
	transpose.reindex()
	transpose.sort()
	return
}

func (m Matrix[T, U]) Cofactor(rowIndex int, colIndex int) (cofactor Operand[T], err error) {
	if len(*m.cells) == 0 {
		return Operand[T]{}.Zero(), errors.New("cannot find cofactor of empty matrix")
	}
	base := (*m.cells)[0].Operand
	zero := base.Zero();
	identity := base.Identity()
	if m.rows == 1 && m.columns == 1 {
		return identity, nil
	}
	if !m.IsSquare() {
		return zero, errors.New("cannot find cofactor of non-square matrix")
	}
	subMatrix, err := m.RemoveColumns(colIndex)
	if err != nil {
		return zero, err
	}
	subMatrix, err = subMatrix.RemoveRows(rowIndex)
	if err != nil {
		return zero, err
	}
	return subMatrix.Determinant()
}

func (m Matrix[T, U]) CofactorMatrix() (cofactorMatrix *Matrix[T, U], err error) {
	mCells := *m.cells
	lenCells := len(mCells)
	if lenCells == 0 {
		return nil, errors.New("cannot create cofactor matrix of empty matrix")
	}
	if !m.IsSquare() {
		return nil, errors.New("cannot create cofactor matrix of non-square matrix")
	}
	cofactorCells := make([]Cell[T], lenCells)
	for index, cell := range (mCells) {
		cofactor, err := m.Cofactor(cell.Row, cell.Column)
		if err != nil {
			return nil, err
		}
		var sign Operand[T] 
		if (cell.Row + cell.Column) % 2 == 0 {
			sign = mCells[0].Operand.Identity()
		} else {
			sign = mCells[0].Operand.FromInt(-1)
		}
		cofactorCells[index] = Cell[T]{
			Row: cell.Row,
			Column: cell.Column,
			Operand: cofactor.Mul(sign),
		}
	}
	cofactorMatrix = &Matrix[T, U]{
		cells: &cofactorCells,
	}
	cofactorMatrix.reindex()
	cofactorMatrix.sort()
	return
}

func (m Matrix[T, U]) String() string {
	rows := m.GetRows()
	rowsAsString := su.Map(rows, func(row []T) string {
		stringRow := su.Map(row, func(value T) string {
			return NewOperand(value).String()
		})
		return fmt.Sprintln("[", strings.Join(stringRow, ", "), "]")
	})
	return fmt.Sprintln("[\n", strings.Join(rowsAsString, "\n"), "\n]")
}

func NewNumberMatrix[T Number](rows [][]T) (matrix *NumberMatrix[T], err error) {
	returnedMatrix, err := NewMatrix[T, float64](rows)
	matrix = &NumberMatrix[T]{Matrix: returnedMatrix}
	return
}

func NewBigMatrix[T BigNumber](rows [][]T) (matrix *BigMatrix[T], err error) {
	returnedMatrix, err := NewMatrix[T, *big.Float](rows)
	matrix = &BigMatrix[T]{Matrix: returnedMatrix}
	return
}

func NewMatrix[T Number | BigNumber, U FloatNumber](rows [][]T) (matrix *Matrix[T, U], err error) {
	if len(rows) == 0 {
		return &Matrix[T, U]{cells: &[]Cell[T]{}}, nil
	}
	var cells *[]Cell[T]
	cells, err = buildCells(rows)
	columnLen := len(rows[0])
	rowLen := len(rows)
	matrix = &Matrix[T, U]{cells: cells,
		columns: columnLen,
		rows:    rowLen}
	return
}

func buildCells[T Number | BigNumber](rows [][]T) (cells *[]Cell[T], err error) {
	cells = &[]Cell[T]{}
	firstRowLength := len(rows[0])
	for rowIndex, row := range rows {
		if currentRowLength := len(row); currentRowLength != firstRowLength {
			err = fmt.Errorf("length of row[%v] (%v) does not match length of row[0] (%v)", rowIndex, currentRowLength, firstRowLength)
			return
		}
		for colIndex, element := range row {
			cell := Cell[T]{
				Row:     rowIndex,
				Column:  colIndex,
				Operand: NewOperand(element),
			}
			newCells := append(*cells, cell)
			cells = &newCells
		}
	}
	return
}

// perfectMatching returns a permutation perm such that for every column j, matrix[perm[j]][j] != 0,
// if such a perfect matching exists. It uses DFS backtracking.
func (m *Matrix[T, U]) perfectMatching() ([]int, bool) {
	n := m.rows
	perm := make([]int, n)
	used := make([]bool, n)

	matrix := m.GetRowCells()

	var dfs func(col int) bool
	dfs = func(col int) bool {
		if col == n {
			return true // found assignments for all columns
		}
		// Try every row for column col.
		for row := 0; row < n; row++ {
			entry := matrix[row][col]
			if !used[row] && entry.Cmp(entry.Zero()) != 0 {
				used[row] = true
				perm[col] = row
				if dfs(col + 1) {
					return true
				}
				used[row] = false // backtrack
			}
		}
		return false
	}

	if dfs(0) {
		return perm, true
	}
	return nil, false
}

// permutationSign computes the sign of the permutation (given as a slice)
// using an inversion count. Returns 1 if even and -1 if odd.
func (m *Matrix[T, U]) permutationSign(perm []int) int {
	inversions := 0
	n := len(perm)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if perm[i] > perm[j] {
				inversions++
			}
		}
	}
	if inversions%2 == 0 {
		return 1
	}
	return -1
}

// reorderRowsPerfectMatching reorders the rows of the matrix using a perfect matching so that
// every diagonal element is nonzero if possible.
// It returns 0 if no perfect matching exists,
// otherwise 1 if the permutation has an even sign or -1 if odd.
func (m Matrix[T, U]) reorderRowsPerfectMatching() (matrix *Matrix[T, U], sign int) {
	n := m.rows
	perm, ok := m.perfectMatching()
	if !ok {
		return nil, 0
	}

	// Compute the sign of the permutation.
	sign = m.permutationSign(perm)

	// Create a new matrix with rows ordered so that new row j is old row perm[j].
	newMat := make([][]T, n)
	rows := m.GetRows()
	for col, rowIndex := range perm {
		newMat[col] = rows[rowIndex]
	}

	// Update the cells of the matrix.
	cells, _ := buildCells(newMat)
	matrix = &Matrix[T, U]{
		cells: cells,
	}
	matrix.reindex()

	return matrix, sign
}
