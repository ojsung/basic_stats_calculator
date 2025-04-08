package matrix

import (
	"reflect"
	"testing"
)

func Test_sort(t *testing.T) {
	type testCase[T Number] struct {
		name     string
		matrix   *NumberMatrix[T]
		expected *[]Cell[T]
	}

	tests := []testCase[int]{
		{
			name: "Sort matrix cells by row and column",
			matrix: &NumberMatrix[int]{
				&Matrix[int, float64]{
					cells: &[]Cell[int]{
						{Operand: NewOperand(6), Row: 1, Column: 1},
						{Operand: NewOperand(4), Row: 0, Column: 1},
						{Operand: NewOperand(5), Row: 1, Column: 0},
						{Operand: NewOperand(3), Row: 0, Column: 0},
					},
				},
			},
			expected: &[]Cell[int]{
				{Operand: NewOperand(3), Row: 0, Column: 0},
				{Operand: NewOperand(4), Row: 0, Column: 1},
				{Operand: NewOperand(5), Row: 1, Column: 0},
				{Operand: NewOperand(6), Row: 1, Column: 1},
			},
		},
		{
			name: "Already sorted matrix cells",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{3, 4},
					{5, 6},
				},
			)),
			expected: &[]Cell[int]{
				{Operand: NewOperand(3), Row: 0, Column: 0},
				{Operand: NewOperand(4), Row: 0, Column: 1},
				{Operand: NewOperand(5), Row: 1, Column: 0},
				{Operand: NewOperand(6), Row: 1, Column: 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.matrix.sort()
			if !reflect.DeepEqual(tt.matrix.cells, tt.expected) {
				t.Errorf("Matrix.sort() failed. Got %v, expected %v", tt.matrix.cells, tt.expected)
			}
		})
	}
}

func Test_reindex(t *testing.T) {
	type testCase[T Number] struct {
		name            string
		matrix          NumberMatrix[T]
		expectedCells *[]Cell[T]
		expectedRows    int
		expectedColumns int
	}
	tests := []testCase[int]{
		{
			name: "should decrement following rows when rows are removed",
			matrix: NumberMatrix[int]{
				Matrix: &Matrix[int, float64]{cells: &[]Cell[int]{
					{Operand: NewOperand(1), Row: 1, Column: 0},
					{Operand: NewOperand(2), Row: 1, Column: 1},
					{Operand: NewOperand(3), Row: 3, Column: 0},
					{Operand: NewOperand(4), Row: 3, Column: 1},
				},
					columns: 2,
					rows:    4,
				},
			},
			expectedCells: &[]Cell[int]{
				{Operand: NewOperand(1), Row: 0, Column: 0},
				{Operand: NewOperand(2), Row: 0, Column: 1},
				{Operand: NewOperand(3), Row: 1, Column: 0},
				{Operand: NewOperand(4), Row: 1, Column: 1},
			},
			expectedRows:    2,
			expectedColumns: 2,
		},
		{
			name: "should decrement following columns when columns are removed",
			matrix: NumberMatrix[int]{
				Matrix: &Matrix[int, float64]{
					cells: &[]Cell[int]{
						{Operand: NewOperand(1), Row: 0, Column: 1},
						{Operand: NewOperand(2), Row: 0, Column: 4},
						{Operand: NewOperand(3), Row: 1, Column: 1},
						{Operand: NewOperand(4), Row: 1, Column: 4},
					},
					columns: 5,
					rows:    2,
				},
			},
			expectedCells: &[]Cell[int]{
				{Operand: NewOperand(1), Row: 0, Column: 0},
				{Operand: NewOperand(2), Row: 0, Column: 1},
				{Operand: NewOperand(3), Row: 1, Column: 0},
				{Operand: NewOperand(4), Row: 1, Column: 1},
			},
			expectedRows:    2,
			expectedColumns: 2,
		},
		{
			name: "should decrement following columns and columns when rows and columns are removed",
			matrix: NumberMatrix[int]{
				Matrix: &Matrix[int, float64]{cells: &[]Cell[int]{
					{Operand: NewOperand(1), Row: 2, Column: 1},
					{Operand: NewOperand(2), Row: 2, Column: 4},
					{Operand: NewOperand(3), Row: 5, Column: 1},
					{Operand: NewOperand(4), Row: 5, Column: 4},
				},
					rows:    6,
					columns: 5,
				},
			},
			expectedCells: &[]Cell[int]{
				{Operand: NewOperand(1), Row: 0, Column: 0},
				{Operand: NewOperand(2), Row: 0, Column: 1},
				{Operand: NewOperand(3), Row: 1, Column: 0},
				{Operand: NewOperand(4), Row: 1, Column: 1},
			},
			expectedRows:    2,
			expectedColumns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.matrix.reindex()
			if cells := tt.matrix.Cells(); !reflect.DeepEqual(cells, tt.expectedCells) {
				t.Errorf("Matrix.reindex() failed. Expected %v, got %v", tt.expectedCells, cells)
			}
			if tt.matrix.rows != tt.expectedRows {
				t.Errorf("Matrix.reindex() row assignment failed. Expected %v, got %v", tt.expectedRows, tt.matrix.rows)
			}
			if tt.matrix.columns != tt.expectedColumns {
				t.Errorf("Matrix.reindex() column assignment failed. Expected %v, got %v", tt.expectedColumns, tt.matrix.columns)
			}
		})
	}
}

func Test_Cells(t *testing.T) {
	type testCase[T Number] struct {
		name     string
		matrix   NumberMatrix[T]
		expected *[]Cell[T]
	}

	tests := []testCase[int]{
		{
			name: "Retrieve cells from matrix",
			matrix: NumberMatrix[int]{
				Matrix: &Matrix[int, float64]{cells: &[]Cell[int]{
					{Operand: NewOperand(3), Row: 1, Column: 1},
					{Operand: NewOperand(4), Row: 1, Column: 2},
					{Operand: NewOperand(5), Row: 2, Column: 1},
					{Operand: NewOperand(6), Row: 2, Column: 2},
				}},
			},
			expected: &[]Cell[int]{
				{Operand: NewOperand(3), Row: 1, Column: 1},
				{Operand: NewOperand(4), Row: 1, Column: 2},
				{Operand: NewOperand(5), Row: 2, Column: 1},
				{Operand: NewOperand(6), Row: 2, Column: 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cells := tt.matrix.Cells()
			if !reflect.DeepEqual(cells, tt.expected) {
				t.Errorf("Matrix.Cells() failed. Got %v, expected %v", cells, tt.expected)
			}
		})
	}
}

// Placeholder removed as no additional code was required to be inserted.
func Test_RemoveColumnsInPlace(t *testing.T) {
	type testCase[T Number] struct {
		name           string
		matrix         *NumberMatrix[T]
		indices        []int
		expected       [][]T
		expectedReturn map[int][]T
	}

	tests := []testCase[int]{
		{
			name: "Remove single column",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			)),
			indices: []int{1},
			expected: [][]int{
				{1, 4, 7},
				{3, 6, 9},
			},
			expectedReturn: map[int][]int{
				1: {2, 5, 8},
			},
		},
		{
			name: "Remove multiple columns",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3, 4},
					{5, 6, 7, 8},
					{9, 10, 11, 12},
				},
			)),
			indices: []int{0, 2},
			expected: [][]int{
				{2, 6, 10},
				{4, 8, 12},
			},
			expectedReturn: map[int][]int{
				0: {1, 5, 9},
				2: {3, 7, 11},
			},
		},
		{
			name: "Remove all columns",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2},
					{3, 4},
				},
			)),
			indices:  []int{0, 1},
			expected: [][]int{},
			expectedReturn: map[int][]int{
				0: {1, 3},
				1: {2, 4},
			},
		},
		{
			name: "Remove no columns",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
			)),
			indices: []int{},
			expected: [][]int{
				{1, 4},
				{2, 5},
				{3, 6},
			},
			expectedReturn: map[int][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.matrix.RemoveColumnsInPlace(tt.indices...)
			if actual := tt.matrix.GetColumns(); !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("Matrix.RemoveColumnsInPlace() failed. Got %v, expected %v", actual, tt.expected)
			}
			if !reflect.DeepEqual(result, tt.expectedReturn) {
				t.Errorf("Matrix.RemoveColumnsInPlace() return value failed. Got %v, expected %v", result, tt.expectedReturn)
			}
		})
	}
}

func Test_Transpose(t *testing.T) {
	type testCase[T Number] struct {
		name     string
		matrix   *NumberMatrix[T]
		expected *NumberMatrix[T]
	}

	tests := []testCase[int]{
		{
			name: "Transpose square matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2},
					{3, 4},
				},
			)),
			expected: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 3},
					{2, 4},
				},
			)),
		},
		{
			name: "Transpose rectangular matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
			)),
			expected: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 4},
					{2, 5},
					{3, 6},
				},
			)),
		},
		{
			name: "Transpose single row matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
				},
			)),
			expected: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1},
					{2},
					{3},
				},
			)),
		},
		{
			name: "Transpose single column matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1},
					{2},
					{3},
				},
			)),
			expected: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
				},
			)),
		},
		{
			name:     "Transpose empty matrix",
			matrix:   getAssumedNoErrorMatrix(NewNumberMatrix[int]([][]int{})),
			expected: getAssumedNoErrorMatrix(NewNumberMatrix[int]([][]int{})),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.matrix.Transpose()
			if !reflect.DeepEqual(result.GetRows(), tt.expected.GetRows()) {
				t.Errorf("Matrix.Transpose() failed. Got %v, expected %v", result.GetRows(), tt.expected.GetRows())
			}
		})
	}
}
func Test_RemoveRowsInPlace(t *testing.T) {
	type testCase[T Number] struct {
		name           string
		matrix         *NumberMatrix[T]
		indices        []int
		expected       [][]T
		expectedReturn map[int][]int
	}

	tests := []testCase[int]{
		{
			name: "Remove single row",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			)),
			indices: []int{1},
			expected: [][]int{
				{1, 2, 3},
				{7, 8, 9},
			},
			expectedReturn: map[int][]int{
				1: {4, 5, 6},
			},
		},
		{
			name: "Remove multiple rows",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
					{10, 11, 12},
				},
			)),
			indices: []int{0, 2},
			expected: [][]int{
				{4, 5, 6},
				{10, 11, 12},
			},
			expectedReturn: map[int][]int{
				0: {1, 2, 3},
				2: {7, 8, 9},
			},
		},
		{
			name: "Remove all rows",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2},
					{3, 4},
				},
			)),
			indices:  []int{0, 1},
			expected: [][]int{},
			expectedReturn: map[int][]int{
				0: {1, 2},
				1: {3, 4},
			},
		},
		{
			name: "Remove no rows",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
			)),
			indices: []int{},
			expected: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			expectedReturn: map[int][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.matrix.RemoveRowsInPlace(tt.indices...)
			if !reflect.DeepEqual(tt.matrix.GetRows(), tt.expected) {
				t.Errorf("Matrix.RemoveRowsInPlace() failed. Result matrix: %v, expected matrix: %v", tt.matrix.GetRows(), tt.expected)
			}
			if !reflect.DeepEqual(result, tt.expectedReturn) {
				t.Errorf("Matrix.RemoveRowsInPlace() failed. Got %v, expected %v", result, tt.expectedReturn)
			}
		})
	}
}

func Test_ScalarMul(t *testing.T) {
	type testCase[T Number] struct {
		name     string
		matrix   *NumberMatrix[T]
		scalar   T
		expected *NumberMatrix[T]
	}

	tests := []testCase[int]{
		{
			name: "Multiply matrix by scalar",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			)),
			scalar: 2,
			expected: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{2, 4, 6},
					{8, 10, 12},
					{14, 16, 18},
				},
			)),
		},
		{
			name: "Multiply matrix by zero",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
			)),
			scalar: 0,
			expected: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{0, 0, 0},
					{0, 0, 0},
				},
			)),
		},
		{
			name:     "Multiply empty matrix by scalar",
			matrix:   getAssumedNoErrorMatrix(NewNumberMatrix([][]int{})),
			scalar:   5,
			expected: getAssumedNoErrorMatrix(NewNumberMatrix([][]int{})),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.matrix.ScalarMul(tt.scalar)
			if !reflect.DeepEqual(result.GetRows(), tt.expected.GetRows()) {
				t.Errorf("Matrix.ScalarMul() failed. Got %v, expected %v", result.GetRows(), tt.expected.GetRows())
			}
		})
	}

	t.Run("should not mutate original matrix", func(t *testing.T) {
		matrix := getAssumedNoErrorMatrix(NewNumberMatrix(
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
		))
		expected := getAssumedNoErrorMatrix(NewNumberMatrix(
			[][]int{
				{2, 4, 6},
				{8, 10, 12},
			},
		))
		originalMatrix := getAssumedNoErrorMatrix(NewNumberMatrix(
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
		))
		scalar := 2

		result := matrix.ScalarMul(scalar)
		if !reflect.DeepEqual(result.cells, expected.cells) {
			t.Errorf("Matrix.ScalarMul() failed. Got %v, expected %v", result.Cells(), expected.Cells())
		}
		if !reflect.DeepEqual(matrix.cells, originalMatrix.cells) {
			t.Errorf("Matrix.ScalarMul() failed. The original matrix was mutated. Expected %v, got %v", originalMatrix, matrix)
		}
	})
}
func Test_MatrixMul(t *testing.T) {
	type testCase[T Number] struct {
		name     string
		matrixA  *NumberMatrix[T]
		matrixB  *NumberMatrix[T]
		expected *NumberMatrix[T]
		wantErr  bool
	}

	tests := []testCase[int]{
		{
			name: "Multiply two 2x2 matrices",
			matrixA: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2},
					{3, 4},
				},
			)),
			matrixB: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{5, 6},
					{7, 8},
				},
			)),
			expected: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{19, 22},
					{43, 50},
				},
			)),
			wantErr: false,
		},
		{
			name: "Multiply 2x3 and 3x2 matrices",
			matrixA: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
			)),
			matrixB: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{7, 8},
					{9, 10},
					{11, 12},
				},
			)),
			expected: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{58, 64},
					{139, 154},
				},
			)),
			wantErr: false,
		},
		{
			name: "Multiply matrix by identity matrix",
			matrixA: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2},
					{3, 4},
				},
			)),
			matrixB: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 0},
					{0, 1},
				},
			)),
			expected: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2},
					{3, 4},
				},
			)),
			wantErr: false,
		},
		{
			name: "Multiply matrix by zero matrix",
			matrixA: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2},
					{3, 4},
				},
			)),
			matrixB: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{0, 0},
					{0, 0},
				},
			)),
			expected: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{0, 0},
					{0, 0},
				},
			)),
			wantErr: false,
		},
		{
			name: "Multiply incompatible matrices",
			matrixA: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
				},
			)),
			matrixB: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{3, 4},
					{5, 6},
				},
			)),
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.matrixA.Mul(tt.matrixB.Matrix)
			if (err != nil) != tt.wantErr {
				t.Errorf("Matrix.Mul() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.expected == nil {
				if result != nil {
					t.Errorf("Matrix.Mul() failed. Expected nil, got %v", result)
				}
			} else if !reflect.DeepEqual(result.GetRows(), tt.expected.GetRows()) {
				t.Errorf("Matrix.Mul() failed. Got %v, expected %v", result.GetRows(), tt.expected.GetRows())
			}
		})
	}
}
func Test_Add(t *testing.T) {
	type testCase[T Number] struct {
		name     string
		matrixA  *NumberMatrix[T]
		matrixB  *NumberMatrix[T]
		expected *NumberMatrix[T]
		wantErr  bool
	}

	tests := []testCase[int]{
		{
			name: "Add two 2x2 matrices",
			matrixA: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2},
					{3, 4},
				},
			)),
			matrixB: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{5, 6},
					{7, 8},
				},
			)),
			expected: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{6, 8},
					{10, 12},
				},
			)),
			wantErr: false,
		},
		{
			name: "Add incompatible matrices",
			matrixA: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2},
				},
			)),
			matrixB: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{3, 4},
					{5, 6},
				},
			)),
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "Add empty matrices",
			matrixA:  getAssumedNoErrorMatrix(NewNumberMatrix([][]int{})),
			matrixB:  getAssumedNoErrorMatrix(NewNumberMatrix([][]int{})),
			expected: getAssumedNoErrorMatrix(NewNumberMatrix([][]int{})),
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.matrixA.Add(tt.matrixB.Matrix)
			if (err != nil) != tt.wantErr {
				t.Errorf("Matrix.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.expected == nil {
				if result != nil {
					t.Errorf("Matrix.Add() failed. Expected nil, got %v", result)
				}
			} else if !reflect.DeepEqual(result.GetRows(), tt.expected.GetRows()) {
				t.Errorf("Matrix.Add() failed. Got %v, expected %v", result.GetRows(), tt.expected.GetRows())
			}
		})
	}
}
func Test_Sub(t *testing.T) {
	type testCase[T Number] struct {
		name     string
		matrixA  *NumberMatrix[T]
		matrixB  *NumberMatrix[T]
		expected *NumberMatrix[T]
		wantErr  bool
	}

	tests := []testCase[int]{
		{
			name: "Subtract two 2x2 matrices",
			matrixA: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{5, 6},
					{7, 8},
				},
			)),
			matrixB: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2},
					{3, 4},
				},
			)),
			expected: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{4, 4},
					{4, 4},
				},
			)),
			wantErr: false,
		},
		{
			name: "Subtract matrix from itself",
			matrixA: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2},
					{3, 4},
				},
			)),
			matrixB: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2},
					{3, 4},
				},
			)),
			expected: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{0, 0},
					{0, 0},
				},
			)),
			wantErr: false,
		},
		{
			name: "Subtract incompatible matrices",
			matrixA: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2},
				},
			)),
			matrixB: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{3, 4},
					{5, 6},
				},
			)),
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "Subtract empty matrices",
			matrixA:  getAssumedNoErrorMatrix(NewNumberMatrix[int]([][]int{})),
			matrixB:  getAssumedNoErrorMatrix(NewNumberMatrix[int]([][]int{})),
			expected: getAssumedNoErrorMatrix(NewNumberMatrix[int]([][]int{})),
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.matrixA.Sub(tt.matrixB.Matrix)
			if (err != nil) != tt.wantErr {
				t.Errorf("Matrix.Subtract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.expected == nil {
				if result != nil {
					t.Errorf("Matrix.Subtract() failed. Expected nil, got %v", result)
				}
			} else if !reflect.DeepEqual(result.GetRows(), tt.expected.GetRows()) {
				t.Errorf("Matrix.Subtract() failed. Got %v, expected %v", result.GetRows(), tt.expected.GetRows())
			}
		})
	}
}
func Test_Trace(t *testing.T) {
	type testCase[T Number] struct {
		name     string
		matrix   *NumberMatrix[T]
		expected T
		wantErr  bool
	}

	tests := []testCase[int]{
		{
			name: "Trace of 2x2 matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2},
					{3, 4},
				},
			)),
			expected: 5, // Trace = 1 + 4
			wantErr:  false,
		},
		{
			name: "Trace of 3x3 matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 0, 0},
					{0, 2, 0},
					{0, 0, 3},
				},
			)),
			expected: 6, // Trace = 1 + 2 + 3
			wantErr:  false,
		},
		{
			name: "Trace of non-square matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
			)),
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "Trace of empty matrix",
			matrix:   getAssumedNoErrorMatrix(NewNumberMatrix[int]([][]int{})),
			expected: 0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.matrix.Trace()
			if (err != nil) != tt.wantErr {
				t.Errorf("Matrix.Trace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.expected != 0 && result != tt.expected {
				t.Errorf("Matrix.Trace() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
}
func Test_IsSquare(t *testing.T) {
	type testCase[T Number] struct {
		name     string
		matrix   *NumberMatrix[T]
		expected bool
	}

	tests := []testCase[int]{
		{
			name: "Square matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2},
					{3, 4},
				},
			)),
			expected: true,
		},
		{
			name: "Non-square matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
			)),
			expected: false,
		},
		{
			name:     "Empty matrix",
			matrix:   getAssumedNoErrorMatrix(NewNumberMatrix[int]([][]int{})),
			expected: false,
		},
		{
			name: "1x1 matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1},
				},
			)),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.matrix.IsSquare()
			if result != tt.expected {
				t.Errorf("Matrix.IsSquare() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
}

func Test_NewMatrix(t *testing.T) {
	type testCase[T Number, U float64] struct {
		name     string
		rows     [][]T
		expected *Matrix[T, U]
		wantErr  bool
	}

	tests := []testCase[int, float64]{
		{
			name: "Create valid 2x2 matrix",
			rows: [][]int{
				{1, 2},
				{3, 4},
			},
			expected: &Matrix[int, float64]{
				cells: &[]Cell[int]{
					{Operand: NewOperand(1), Row: 0, Column: 0},
					{Operand: NewOperand(2), Row: 0, Column: 1},
					{Operand: NewOperand(3), Row: 1, Column: 0},
					{Operand: NewOperand(4), Row: 1, Column: 1},
				},
			},
			wantErr: false,
		},
		{
			name: "Create valid 3x3 matrix",
			rows: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: &Matrix[int, float64]{
				cells: &[]Cell[int]{
					{Operand: NewOperand(1), Row: 0, Column: 0},
					{Operand: NewOperand(2), Row: 0, Column: 1},
					{Operand: NewOperand(3), Row: 0, Column: 2},
					{Operand: NewOperand(4), Row: 1, Column: 0},
					{Operand: NewOperand(5), Row: 1, Column: 1},
					{Operand: NewOperand(6), Row: 1, Column: 2},
					{Operand: NewOperand(7), Row: 2, Column: 0},
					{Operand: NewOperand(8), Row: 2, Column: 1},
					{Operand: NewOperand(9), Row: 2, Column: 2},
				},
			},
			wantErr: false,
		},
		{
			name: "Empty matrix",
			rows: [][]int{},
			expected: &Matrix[int, float64]{
				cells: &[]Cell[int]{},
			},
			wantErr: false,
		},
		{
			name: "Mismatched row lengths",
			rows: [][]int{
				{1, 2},
				{3},
			},
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := NewNumberMatrix(tt.rows)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMatrix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(result.cells, tt.expected.cells) {
				t.Errorf("NewMatrix() failed. Got %v, expected %v", result.cells, tt.expected.cells)
			}
		})
	}
}
func Test_Determinant(t *testing.T) {
	type testCase[T Number] struct {
		name     string
		matrix   *NumberMatrix[T]
		expected T
		wantErr  bool
	}

	tests := []testCase[int]{
		{
			name: "Determinant of 2x2 matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2},
					{3, 4},
				},
			)),
			expected: -2, // Determinant = (1*4) - (2*3) = -2
			wantErr:  false,
		},
		{
			name: "Determinant of 3x3 matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{6, 1, 1},
					{4, -2, 5},
					{2, 8, 7},
				},
			)),
			expected: -306, // Determinant calculated manually
			wantErr:  false,
		},
		{
			name: "Determinant of 1x1 matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{5},
				},
			)),
			expected: 5, // Determinant of a 1x1 matrix is the single element
			wantErr:  false,
		},
		{
			name:     "Determinant of empty matrix",
			matrix:   getAssumedNoErrorMatrix(NewNumberMatrix[int]([][]int{})),
			expected: 0,
			wantErr:  true,
		},
		{
			name: "Determinant of non-square matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
			)),
			expected: 0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.matrix.Determinant()
			if (err != nil) != tt.wantErr {
				t.Errorf("Matrix.Determinant() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result.Value != tt.expected {
				t.Errorf("Matrix.Determinant() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
}
func Test_Rows(t *testing.T) {
	type testCase[T Number] struct {
		name     string
		matrix   *NumberMatrix[T]
		expected int
	}

	tests := []testCase[int]{
		{
			name: "Matrix with 3 rows",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			)),
			expected: 3,
		},
		{
			name: "Matrix with 1 row",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
				},
			)),
			expected: 1,
		},
		{
			name:     "Empty matrix",
			matrix:   getAssumedNoErrorMatrix(NewNumberMatrix[int]([][]int{})),
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.matrix.Rows()
			if result != tt.expected {
				t.Errorf("Matrix.Rows() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
}

func Test_Columns(t *testing.T) {
	type testCase[T Number] struct {
		name     string
		matrix   *NumberMatrix[T]
		expected int
	}

	tests := []testCase[int]{
		{
			name: "Matrix with 3 columns",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			)),
			expected: 3,
		},
		{
			name: "Matrix with 1 column",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1},
					{2},
					{3},
				},
			)),
			expected: 1,
		},
		{
			name:     "Empty matrix",
			matrix:   getAssumedNoErrorMatrix(NewNumberMatrix[int]([][]int{})),
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.matrix.Columns()
			if result != tt.expected {
				t.Errorf("Matrix.Columns() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
}

func Test_GetRows(t *testing.T) {
	type testCase[T Number] struct {
		name     string
		matrix   *NumberMatrix[T]
		expected [][]T
	}

	tests := []testCase[int]{
		{
			name: "Matrix with multiple rows",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			)),
			expected: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		},
		{
			name: "Matrix with a single row",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
				},
			)),
			expected: [][]int{
				{1, 2, 3},
			},
		},
		{
			name:     "Matrix with no rows (empty matrix)",
			matrix:   getAssumedNoErrorMatrix(NewNumberMatrix[int]([][]int{})),
			expected: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.matrix.GetRows()
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Matrix.GetRows() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
}

func Test_GetColumns(t *testing.T) {
	type testCase[T Number] struct {
		name     string
		matrix   *NumberMatrix[T]
		expected [][]T
	}

	tests := []testCase[int]{
		{
			name: "Matrix with multiple columns",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			)),
			expected: [][]int{
				{1, 4, 7},
				{2, 5, 8},
				{3, 6, 9},
			},
		},
		{
			name: "Matrix with a single column",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1},
					{2},
					{3},
				},
			)),
			expected: [][]int{
				{1, 2, 3},
			},
		},
		{
			name:     "Matrix with no rows (empty matrix)",
			matrix:   getAssumedNoErrorMatrix(NewNumberMatrix[int]([][]int{})),
			expected: [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.matrix.GetColumns()
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Matrix.GetColumns() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
}

func Test_GetRow(t *testing.T) {
	type testCase[T Number] struct {
		name     string
		matrix   *NumberMatrix[T]
		index    int
		expected []T
		wantErr  bool
	}

	tests := []testCase[int]{
		{
			name: "Get valid row from 3x3 matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			)),
			index:    1,
			expected: []int{4, 5, 6},
			wantErr:  false,
		},
		{
			name: "Get first row from 2x2 matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2},
					{3, 4},
				},
			)),
			index:    0,
			expected: []int{1, 2},
			wantErr:  false,
		},
		{
			name: "Get last row from rectangular matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2},
					{3, 4},
					{5, 6},
				},
			)),
			index:    2,
			expected: []int{5, 6},
			wantErr:  false,
		},
		{
			name: "Get row from single-column matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1},
					{2},
					{3},
				},
			)),
			index:    1,
			expected: []int{2},
			wantErr:  false,
		},
		{
			name: "Get row from single-row matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
				},
			)),
			index:    0,
			expected: []int{1, 2, 3},
			wantErr:  false,
		},
		{
			name:     "Get row from empty matrix",
			matrix:   getAssumedNoErrorMatrix(NewNumberMatrix[int]([][]int{})),
			index:    0,
			expected: nil,
			wantErr:  true,
		},
		{
			name: "Get row with out-of-bounds index",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2},
					{3, 4},
				},
			)),
			index:    3,
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.matrix.GetRow(tt.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Matrix.GetColumn() error. Got %v, expected %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Matrix.GetColumn() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
}

func Test_GetColumn(t *testing.T) {
	type testCase[T Number] struct {
		name     string
		matrix   *NumberMatrix[T]
		index    int
		expected []T
		wantErr  bool
	}

	tests := []testCase[int]{
		{
			name: "Get valid column from 3x3 matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			)),
			index:    1,
			expected: []int{2, 5, 8},
			wantErr:  false,
		},
		{
			name: "Get first column from 2x2 matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2},
					{3, 4},
				},
			)),
			index:    0,
			expected: []int{1, 3},
			wantErr:  false,
		},
		{
			name: "Get last column from rectangular matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
			)),
			index:    2,
			expected: []int{3, 6},
			wantErr:  false,
		},
		{
			name: "Get column from single-row matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
				},
			)),
			index:    1,
			expected: []int{2},
			wantErr:  false,
		},
		{
			name: "Get column from single-column matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1},
					{2},
					{3},
				},
			)),
			index:    0,
			expected: []int{1, 2, 3},
			wantErr:  false,
		},
		{
			name:     "Get column from empty matrix",
			matrix:   getAssumedNoErrorMatrix(NewNumberMatrix[int]([][]int{})),
			index:    0,
			expected: nil,
			wantErr:  true,
		},
		{
			name: "Get column with out-of-bounds index",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2},
					{3, 4},
				},
			)),
			index:    3,
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.matrix.GetColumn(tt.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Matrix.GetColumn() error. Got %v, expected %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Matrix.GetColumn() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
}

func Test_RemoveRows(t *testing.T) {
	type testCase[T Number] struct {
		name           string
		matrix         *NumberMatrix[T]
		indices        []int
		expected       *NumberMatrix[T]
		expectedErr    bool
	}

	tests := []testCase[int]{
		{
			name: "Remove single row",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			)),
			indices: []int{1},
			expected: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
					{7, 8, 9},
				},
			)),
			expectedErr: false,
		},
		{
			name: "Remove multiple rows",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
					{10, 11, 12},
				},
			)),
			indices: []int{0, 2},
			expected: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{4, 5, 6},
					{10, 11, 12},
				},
			)),
			expectedErr: false,
		},
		{
			name: "Remove all rows",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2},
					{3, 4},
				},
			)),
			indices: []int{0, 1},
			expected: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{},
			)),
			expectedErr: false,
		},
		{
			name: "Remove no rows",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
			)),
			indices: []int{},
			expected: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
			)),
			expectedErr: false,
		},
		{
			name: "Remove row with out-of-bounds index",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2},
					{3, 4},
				},
			)),
			indices:     []int{2},
			expected:    nil,
			expectedErr: true,
		},
		{
			name: "Remove row from empty matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix[int]([][]int{})),
			indices:     []int{0},
			expected:    nil,
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.matrix.RemoveRows(tt.indices...)
			if (err != nil) != tt.expectedErr {
				t.Errorf("Matrix.RemoveRows() error = %v, expectedErr %v", err, tt.expectedErr)
				return
			}
			if tt.expected == nil {
				if result != nil {
					t.Errorf("Matrix.RemoveRows() failed. Expected nil, got %v", result)
				}
			} else if !reflect.DeepEqual(result.GetRows(), tt.expected.GetRows()) {
				t.Errorf("Matrix.RemoveRows() failed. Got %v, expected %v", result.GetRows(), tt.expected.GetRows())
			}
		})
	}
}

func Test_RemoveColumns(t *testing.T) {
	type testCase[T Number] struct {
		name           string
		matrix         *NumberMatrix[T]
		indices        []int
		expected       *NumberMatrix[T]
		expectedErr    bool
	}

	tests := []testCase[int]{
		{
			name: "Remove single column",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			)),
			indices: []int{1},
			expected: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 3},
					{4, 6},
					{7, 9},
				},
			)),
			expectedErr: false,
		},
		{
			name: "Remove multiple columns",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3, 4},
					{5, 6, 7, 8},
					{9, 10, 11, 12},
				},
			)),
			indices: []int{0, 2},
			expected: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{2, 4},
					{6, 8},
					{10, 12},
				},
			)),
			expectedErr: false,
		},
		{
			name: "Remove all columns",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2},
					{3, 4},
				},
			)),
			indices: []int{0, 1},
			expected: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{},
			)),
			expectedErr: false,
		},
		{
			name: "Remove no columns",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
			)),
			indices: []int{},
			expected: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
			)),
			expectedErr: false,
		},
		{
			name: "Remove column with out-of-bounds index",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2},
					{3, 4},
				},
			)),
			indices:     []int{2},
			expected:    nil,
			expectedErr: true,
		},
		{
			name: "Remove column from empty matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix([][]int{})),
			indices:     []int{0},
			expected:    nil,
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.matrix.RemoveColumns(tt.indices...)
			if (err != nil) != tt.expectedErr {
				t.Errorf("Matrix.RemoveColumns() error = %v, expectedErr %v", err, tt.expectedErr)
				return
			}
			if tt.expected == nil {
				if result != nil {
					t.Errorf("Matrix.RemoveColumns() failed. Expected nil, got %v", result)
				}
			} else if !reflect.DeepEqual(result.GetRows(), tt.expected.GetRows()) {
				t.Errorf("Matrix.RemoveColumns() failed. Got %v, expected %v", result.GetRows(), tt.expected.GetRows())
			}
		})
	}
}

func Test_CofactorMatrix(t *testing.T) {
	type testCase[T Number] struct {
		name     string
		matrix   *NumberMatrix[T]
		expected *NumberMatrix[T]
		wantErr  bool
	}

	tests := []testCase[int]{
		{
			name: "Cofactor matrix of 2x2 matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2},
					{3, 4},
				},
			)),
			expected: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{4, -3},
					{-2, 1},
				},
			)),
			wantErr: false,
		},
		{
			name: "Cofactor matrix of 3x3 matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{6, 1, 1},
					{4, -2, 5},
					{2, 8, 7},
				},
			)),
			expected: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{-54, -18, 36},
					{1, 40, -46},
					{7, -26, -16},
				},
			)),
			wantErr: false,
		},
		{
			name: "Cofactor matrix of 1x1 matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{5},
				},
			)),
			expected: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1},
				},
			)),
			wantErr: false,
		},
		{
			name:     "Cofactor matrix of empty matrix",
			matrix:   getAssumedNoErrorMatrix(NewNumberMatrix[int]([][]int{})),
			expected: nil,
			wantErr:  true,
		},
		{
			name: "Cofactor matrix of non-square matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
			)),
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.matrix.CofactorMatrix()
			if (err != nil) != tt.wantErr {
				t.Errorf("Matrix.CofactorMatrix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.expected == nil {
				if result != nil {
					t.Errorf("Matrix.CofactorMatrix() failed. Expected nil, got %v", result)
				}
			} else if !reflect.DeepEqual(result.GetRows(), tt.expected.GetRows()) {
				t.Errorf("Matrix.CofactorMatrix() failed. Got %v, expected %v", result.GetRows(), tt.expected.GetRows())
			}
		})
	}
}

func Test_Cofactor(t *testing.T) {
	type testCase[T Number] struct {
		name     string
		matrix   *NumberMatrix[T]
		rowIndex int
		colIndex int
		expected T
		wantErr  bool
	}

	tests := []testCase[int]{
		{
			name: "Cofactor of 2x2 matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2},
					{3, 4},
				},
			)),
			rowIndex: 0,
			colIndex: 0,
			expected: 4, // Determinant of submatrix [[4]]
			wantErr:  false,
		},
		{
			name: "Cofactor of 3x3 matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{6, 1, 1},
					{4, -2, 5},
					{2, 8, 7},
				},
			)),
			rowIndex: 1,
			colIndex: 1,
			expected: 40, // Determinant of submatrix [[6, 1], [2, 7]]
			wantErr:  false,
		},
		{
			name: "Cofactor of 1x1 matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{5},
				},
			)),
			rowIndex: 0,
			colIndex: 0,
			expected: 1, // Cofactor of a 1x1 matrix is 1
			wantErr:  false,
		},
		{
			name:     "Cofactor of empty matrix",
			matrix:   getAssumedNoErrorMatrix(NewNumberMatrix[int]([][]int{})),
			rowIndex: 0,
			colIndex: 0,
			expected: 0,
			wantErr:  true,
		},
		{
			name: "Cofactor of non-square matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
			)),
			rowIndex: 0,
			colIndex: 0,
			expected: 0,
			wantErr:  true,
		},
		{
			name: "Cofactor with out-of-bounds indices",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2},
					{3, 4},
				},
			)),
			rowIndex: 2,
			colIndex: 2,
			expected: 0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.matrix.Cofactor(tt.rowIndex, tt.colIndex)
			if (err != nil) != tt.wantErr {
				t.Errorf("Matrix.Cofactor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result.Value != tt.expected {
				t.Errorf("Matrix.Cofactor() failed. Got %v, expected %v", result.Value, tt.expected)
			}
		})
	}
}

func Test_Inverse(t *testing.T) {
	type testCase[T Number] struct {
		name     string
		matrix   *NumberMatrix[T]
		expected *NumberMatrix[float64]
		wantErr  bool
		isSingular bool
	}

	tests := []testCase[int]{
		{
			name: "Inverse of 2x2 matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{4, 7},
					{2, 6},
				},
			)),
			expected: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]float64{
					{0.6, -0.7},
					{-0.2, 0.4},
				},
			)),
			wantErr: false,
			isSingular: false,
		},
		{
			name: "Inverse of singular matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2},
					{2, 4},
				},
			)),
			expected: nil,
			wantErr:  false,
			isSingular: true,
		},
		{
			name: "Inverse of 1x1 matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{5},
				},
			)),
			expected: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]float64{
					{0.2},
				},
			)),
			wantErr: false,
			isSingular: false,
		},
		{
			name:     "Inverse of empty matrix",
			matrix:   getAssumedNoErrorMatrix(NewNumberMatrix[int]([][]int{})),
			expected: nil,
			wantErr:  true,
			isSingular: false,
		},
		{
			name: "Inverse of non-square matrix",
			matrix: getAssumedNoErrorMatrix(NewNumberMatrix(
				[][]int{
					{1, 2, 3},
					{4, 5, 6},
				},
			)),
			expected: nil,
			wantErr:  true,
			isSingular: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, isSingular, err := tt.matrix.Inverse()
			if isSingular != tt.isSingular {
				t.Errorf("Matrix.Inverse() isSingular = %v, expected %v", isSingular, tt.isSingular)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("Matrix.Inverse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.expected == nil {
				if result != nil {
					t.Errorf("Matrix.Inverse() failed. Expected nil, got %v", result)
				}
			} else if !floatMatricesEqual(result.GetRows(), tt.expected.GetRows(), .0001) {
				t.Errorf("Matrix.Inverse() failed. Got %v, expected %v", result.GetRows(), tt.expected.GetRows())
			}
		})
	}
}

func getAssumedNoErrorMatrix[T Number](matrix *NumberMatrix[T], err error) *NumberMatrix[T] {
	if err != nil {
		panic("error should be nil")
	}
	return matrix
}

func floatsEqual(a, b, epsilon float64) bool {
    return (a-b) < epsilon && (b-a) < epsilon
}

func floatMatricesEqual(a, b [][]float64, epsilon float64) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if len(a[i]) != len(b[i]) {
            return false
        }
        for j := range a[i] {
            if !floatsEqual(a[i][j], b[i][j], epsilon) {
                return false
            }
        }
    }
    return true
}