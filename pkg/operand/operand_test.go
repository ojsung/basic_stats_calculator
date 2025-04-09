package operand

import (
	"math/big"
	"reflect"
	"testing"
)

func Test_NewOperand(t *testing.T) {
	type testCase[T Number | BigNumber] struct {
		name     string
		value    T
		expected Operand[T]
	}

	intTests := []testCase[int]{
		{
			name:     "Create Operand with int",
			value:    5,
			expected: Operand[int]{Value: 5},
		},
	}

	float64Tests := []testCase[float64]{
		{
			name:     "Create Operand with float64",
			value:    3.14,
			expected: Operand[float64]{Value: 3.14},
		},
	}

	bigIntTests := []testCase[*big.Int]{
		{
			name:     "Create Operand with *big.Int",
			value:    big.NewInt(42),
			expected: Operand[*big.Int]{Value: big.NewInt(42)},
		},
	}

	bigFloatTests := []testCase[*big.Float]{
		{
			name:     "Create Operand with *big.Float",
			value:    big.NewFloat(1.23),
			expected: Operand[*big.Float]{Value: big.NewFloat(1.23)},
		},
	}

	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewOperand(tt.value)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("NewOperand() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
	for _, tt := range float64Tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewOperand(tt.value)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("NewOperand() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
	for _, tt := range bigIntTests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewOperand(tt.value)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("NewOperand() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
	for _, tt := range bigFloatTests {
		t.Run(tt.name, func(t *testing.T) {
			result := NewOperand(tt.value)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("NewOperand() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
}
func Test_Operand_String(t *testing.T) {
	type testCase[T Number | BigNumber] struct {
		name     string
		operand  Operand[T]
		expected string
	}

	intTests := []testCase[int]{
		{
			name:     "String representation of Operand with int",
			operand:  Operand[int]{Value: 5},
			expected: "Operand[int]{5}",
		},
	}

	float64Tests := []testCase[float64]{
		{
			name:     "String representation of Operand with float64",
			operand:  Operand[float64]{Value: 3.14},
			expected: "Operand[float64]{3.14}",
		},
	}

	bigIntTests := []testCase[*big.Int]{
		{
			name:     "String representation of Operand with *big.Int",
			operand:  Operand[*big.Int]{Value: big.NewInt(42)},
			expected: "Operand[*big.Int]{42}",
		},
	}

	bigFloatTests := []testCase[*big.Float]{
		{
			name:     "String representation of Operand with *big.Float",
			operand:  Operand[*big.Float]{Value: big.NewFloat(1.23)},
			expected: "Operand[*big.Float]{1.23}",
		},
	}

	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.String()
			if result != tt.expected {
				t.Errorf("Operand.String() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
	for _, tt := range float64Tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.String()
			if result != tt.expected {
				t.Errorf("Operand.String() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
	for _, tt := range bigIntTests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.String()
			if result != tt.expected {
				t.Errorf("Operand.String() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
	for _, tt := range bigFloatTests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.String()
			if result != tt.expected {
				t.Errorf("Operand.String() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
}
func Test_Operand_AddValue(t *testing.T) {
	type testCase[T Number | BigNumber] struct {
		name     string
		operand  Operand[T]
		summand  T
		expected Operand[T]
	}

	intTests := []testCase[int]{
		{
			name:     "AddValue with int",
			operand:  Operand[int]{Value: 5},
			summand:  3,
			expected: Operand[int]{Value: 8},
		},
	}

	float64Tests := []testCase[float64]{
		{
			name:     "AddValue with float64",
			operand:  Operand[float64]{Value: 3.14},
			summand:  1.86,
			expected: Operand[float64]{Value: 5.0},
		},
	}

	bigIntTests := []testCase[*big.Int]{
		{
			name:     "AddValue with *big.Int",
			operand:  Operand[*big.Int]{Value: big.NewInt(42)},
			summand:  big.NewInt(8),
			expected: Operand[*big.Int]{Value: big.NewInt(50)},
		},
	}

	bigFloatTests := []testCase[*big.Float]{
		{
			name:     "AddValue with *big.Float",
			operand:  Operand[*big.Float]{Value: big.NewFloat(1.23)},
			summand:  big.NewFloat(0.77),
			expected: Operand[*big.Float]{Value: big.NewFloat(2.0)},
		},
	}

	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.AddValue(tt.summand)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Operand.AddValue() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
	for _, tt := range float64Tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.AddValue(tt.summand)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Operand.AddValue() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
	for _, tt := range bigIntTests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.AddValue(tt.summand)
			if result.Value.Cmp(tt.expected.Value) != 0 {
				t.Errorf("Operand.AddValue() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
	for _, tt := range bigFloatTests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.AddValue(tt.summand)
			if result.Value.Cmp(tt.expected.Value) != 0 {
				t.Errorf("Operand.AddValue() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
}

func Test_Operand_SubValue(t *testing.T) {
	type testCase[T Number | BigNumber] struct {
		name        string
		operand     Operand[T]
		subtrahend  T
		expected    Operand[T]
	}

	intTests := []testCase[int]{
		{
			name:        "SubValue with int",
			operand:     Operand[int]{Value: 10},
			subtrahend:  3,
			expected:    Operand[int]{Value: 7},
		},
	}

	float64Tests := []testCase[float64]{
		{
			name:        "SubValue with float64",
			operand:     Operand[float64]{Value: 5.5},
			subtrahend:  2.5,
			expected:    Operand[float64]{Value: 3.0},
		},
	}

	bigIntTests := []testCase[*big.Int]{
		{
			name:        "SubValue with *big.Int",
			operand:     Operand[*big.Int]{Value: big.NewInt(50)},
			subtrahend:  big.NewInt(20),
			expected:    Operand[*big.Int]{Value: big.NewInt(30)},
		},
	}

	bigFloatTests := []testCase[*big.Float]{
		{
			name:        "SubValue with *big.Float",
			operand:     Operand[*big.Float]{Value: big.NewFloat(3.5)},
			subtrahend:  big.NewFloat(1.5),
			expected:    Operand[*big.Float]{Value: big.NewFloat(2.0)},
		},
	}

	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.SubValue(tt.subtrahend)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Operand.SubValue() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
	for _, tt := range float64Tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.SubValue(tt.subtrahend)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Operand.SubValue() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
	for _, tt := range bigIntTests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.SubValue(tt.subtrahend)
			if result.Value.Cmp(tt.expected.Value) != 0 {
				t.Errorf("Operand.SubValue() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
	for _, tt := range bigFloatTests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.SubValue(tt.subtrahend)
			if result.Value.Cmp(tt.expected.Value) != 0 {
				t.Errorf("Operand.SubValue() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
}
func Test_Operand_MulValue(t *testing.T) {
	type testCase[T Number | BigNumber] struct {
		name       string
		operand    Operand[T]
		multiplier T
		expected   Operand[T]
	}

	intTests := []testCase[int]{
		{
			name:       "MulValue with int",
			operand:    Operand[int]{Value: 5},
			multiplier: 3,
			expected:   Operand[int]{Value: 15},
		},
	}

	float64Tests := []testCase[float64]{
		{
			name:       "MulValue with float64",
			operand:    Operand[float64]{Value: 2.5},
			multiplier: 2.0,
			expected:   Operand[float64]{Value: 5.0},
		},
	}

	bigIntTests := []testCase[*big.Int]{
		{
			name:       "MulValue with *big.Int",
			operand:    Operand[*big.Int]{Value: big.NewInt(6)},
			multiplier: big.NewInt(7),
			expected:   Operand[*big.Int]{Value: big.NewInt(42)},
		},
	}

	bigFloatTests := []testCase[*big.Float]{
		{
			name:       "MulValue with *big.Float",
			operand:    Operand[*big.Float]{Value: big.NewFloat(1.5)},
			multiplier: big.NewFloat(2.0),
			expected:   Operand[*big.Float]{Value: big.NewFloat(3.0)},
		},
	}

	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.MulValue(tt.multiplier)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Operand.MulValue() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
	for _, tt := range float64Tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.MulValue(tt.multiplier)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Operand.MulValue() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
	for _, tt := range bigIntTests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.MulValue(tt.multiplier)
			if result.Value.Cmp(tt.expected.Value) != 0 {
				t.Errorf("Operand.MulValue() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
	for _, tt := range bigFloatTests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.MulValue(tt.multiplier)
			if result.Value.Cmp(tt.expected.Value) != 0 {
				t.Errorf("Operand.MulValue() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
}

func Test_Operand_DivValue(t *testing.T) {
	type testCase[T Number | BigNumber] struct {
		name     string
		operand  Operand[T]
		divisor  T
		expected Operand[T]
	}

	intTests := []testCase[int]{
		{
			name:     "DivValue with int",
			operand:  Operand[int]{Value: 10},
			divisor:  2,
			expected: Operand[int]{Value: 5},
		},
	}

	float64Tests := []testCase[float64]{
		{
			name:     "DivValue with float64",
			operand:  Operand[float64]{Value: 5.0},
			divisor:  2.0,
			expected: Operand[float64]{Value: 2.5},
		},
	}

	bigIntTests := []testCase[*big.Int]{
		{
			name:     "DivValue with *big.Int",
			operand:  Operand[*big.Int]{Value: big.NewInt(42)},
			divisor:  big.NewInt(7),
			expected: Operand[*big.Int]{Value: big.NewInt(6)},
		},
	}

	bigFloatTests := []testCase[*big.Float]{
		{
			name:     "DivValue with *big.Float",
			operand:  Operand[*big.Float]{Value: big.NewFloat(3.0)},
			divisor:  big.NewFloat(1.5),
			expected: Operand[*big.Float]{Value: big.NewFloat(2.0)},
		},
	}

	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.DivValue(tt.divisor)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Operand.DivValue() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
	for _, tt := range float64Tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.DivValue(tt.divisor)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Operand.DivValue() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
	for _, tt := range bigIntTests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.DivValue(tt.divisor)
			if result.Value.Cmp(tt.expected.Value) != 0 {
				t.Errorf("Operand.DivValue() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
	for _, tt := range bigFloatTests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.DivValue(tt.divisor)
			if result.Value.Cmp(tt.expected.Value) != 0 {
				t.Errorf("Operand.DivValue() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
}
func Test_Operand_Zero(t *testing.T) {
	type testCase[T Number | BigNumber] struct {
		name     string
		operand  Operand[T]
		expected Operand[T]
	}

	intTests := []testCase[int]{
		{
			name:     "Zero value for int",
			operand:  Operand[int]{Value: 5},
			expected: Operand[int]{Value: 0},
		},
	}

	float64Tests := []testCase[float64]{
		{
			name:     "Zero value for float64",
			operand:  Operand[float64]{Value: 3.14},
			expected: Operand[float64]{Value: 0.0},
		},
	}

	bigIntTests := []testCase[*big.Int]{
		{
			name:     "Zero value for *big.Int",
			operand:  Operand[*big.Int]{Value: big.NewInt(42)},
			expected: Operand[*big.Int]{Value: big.NewInt(0)},
		},
	}

	bigFloatTests := []testCase[*big.Float]{
		{
			name:     "Zero value for *big.Float",
			operand:  Operand[*big.Float]{Value: big.NewFloat(1.23)},
			expected: Operand[*big.Float]{Value: big.NewFloat(0.0)},
		},
	}

	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.Zero()
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Operand.Zero() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
	for _, tt := range float64Tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.Zero()
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Operand.Zero() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
	for _, tt := range bigIntTests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.Zero()
			if result.Value.Cmp(tt.expected.Value) != 0 {
				t.Errorf("Operand.Zero() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
	for _, tt := range bigFloatTests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.Zero()
			if result.Value.Cmp(tt.expected.Value) != 0 {
				t.Errorf("Operand.Zero() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
}

func Test_Operand_Identity(t *testing.T) {
	type testCase[T Number | BigNumber] struct {
		name     string
		operand  Operand[T]
		expected Operand[T]
	}

	intTests := []testCase[int]{
		{
			name:     "Identity value for int",
			operand:  Operand[int]{Value: 5},
			expected: Operand[int]{Value: 1},
		},
	}

	float64Tests := []testCase[float64]{
		{
			name:     "Identity value for float64",
			operand:  Operand[float64]{Value: 3.14},
			expected: Operand[float64]{Value: 1.0},
		},
	}

	bigIntTests := []testCase[*big.Int]{
		{
			name:     "Identity value for *big.Int",
			operand:  Operand[*big.Int]{Value: big.NewInt(42)},
			expected: Operand[*big.Int]{Value: big.NewInt(1)},
		},
	}

	bigFloatTests := []testCase[*big.Float]{
		{
			name:     "Identity value for *big.Float",
			operand:  Operand[*big.Float]{Value: big.NewFloat(1.23)},
			expected: Operand[*big.Float]{Value: big.NewFloat(1.0)},
		},
	}

	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.Identity()
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Operand.Identity() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
	for _, tt := range float64Tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.Identity()
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Operand.Identity() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
	for _, tt := range bigIntTests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.Identity()
			if result.Value.Cmp(tt.expected.Value) != 0 {
				t.Errorf("Operand.Identity() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
	for _, tt := range bigFloatTests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.Identity()
			if result.Value.Cmp(tt.expected.Value) != 0 {
				t.Errorf("Operand.Identity() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
}

func Test_Operand_Cmp(t *testing.T) {
	type testCase[T Number | BigNumber] struct {
		name     string
		operand  Operand[T]
		value    Operand[T]
		expected int
	}

	intTests := []testCase[int]{
		{
			name:     "Cmp with int - equal",
			operand:  Operand[int]{Value: 5},
			value:    Operand[int]{Value: 5},
			expected: 0,
		},
		{
			name:     "Cmp with int - greater",
			operand:  Operand[int]{Value: 10},
			value:    Operand[int]{Value: 5},
			expected: 1,
		},
		{
			name:     "Cmp with int - less",
			operand:  Operand[int]{Value: 3},
			value:    Operand[int]{Value: 5},
			expected: -1,
		},
	}

	float64Tests := []testCase[float64]{
		{
			name:     "Cmp with float64 - equal",
			operand:  Operand[float64]{Value: 3.14},
			value:    Operand[float64]{Value: 3.14},
			expected: 0,
		},
		{
			name:     "Cmp with float64 - greater",
			operand:  Operand[float64]{Value: 5.0},
			value:    Operand[float64]{Value: 3.14},
			expected: 1,
		},
		{
			name:     "Cmp with float64 - less",
			operand:  Operand[float64]{Value: 2.0},
			value:    Operand[float64]{Value: 3.14},
			expected: -1,
		},
	}

	bigIntTests := []testCase[*big.Int]{
		{
			name:     "Cmp with *big.Int - equal",
			operand:  Operand[*big.Int]{Value: big.NewInt(42)},
			value:    Operand[*big.Int]{Value: big.NewInt(42)},
			expected: 0,
		},
		{
			name:     "Cmp with *big.Int - greater",
			operand:  Operand[*big.Int]{Value: big.NewInt(50)},
			value:    Operand[*big.Int]{Value: big.NewInt(42)},
			expected: 1,
		},
		{
			name:     "Cmp with *big.Int - less",
			operand:  Operand[*big.Int]{Value: big.NewInt(30)},
			value:    Operand[*big.Int]{Value: big.NewInt(42)},
			expected: -1,
		},
	}

	bigFloatTests := []testCase[*big.Float]{
		{
			name:     "Cmp with *big.Float - equal",
			operand:  Operand[*big.Float]{Value: big.NewFloat(1.23)},
			value:    Operand[*big.Float]{Value: big.NewFloat(1.23)},
			expected: 0,
		},
		{
			name:     "Cmp with *big.Float - greater",
			operand:  Operand[*big.Float]{Value: big.NewFloat(2.0)},
			value:    Operand[*big.Float]{Value: big.NewFloat(1.23)},
			expected: 1,
		},
		{
			name:     "Cmp with *big.Float - less",
			operand:  Operand[*big.Float]{Value: big.NewFloat(0.5)},
			value:    Operand[*big.Float]{Value: big.NewFloat(1.23)},
			expected: -1,
		},
	}

	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.Cmp(tt.value)
			if result != tt.expected {
				t.Errorf("Operand.Cmp() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
	for _, tt := range float64Tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.Cmp(tt.value)
			if result != tt.expected {
				t.Errorf("Operand.Cmp() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
	for _, tt := range bigIntTests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.Cmp(tt.value)
			if result != tt.expected {
				t.Errorf("Operand.Cmp() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
	for _, tt := range bigFloatTests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.operand.Cmp(tt.value)
			if result != tt.expected {
				t.Errorf("Operand.Cmp() failed. Got %v, expected %v", result, tt.expected)
			}
		})
	}
}
