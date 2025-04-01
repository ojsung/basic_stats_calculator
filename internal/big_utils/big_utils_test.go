package big_utils

import (
	"math/big"
	"testing"
)

func Test_RoundDown(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"It should round 1.1 down to 1", "1.1", "1"},
		{"It should round 2.9 down to 2", "2.9", "2"},
		{"It should not round integers", "0.0", "0"},
		{"It should not round negative integers", "-1", "-1"},
		{"It should round -2.9 down to -3", "-2.9", "-3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := new(big.Float)
			input.SetString(tt.input)

			expected := new(big.Int)
			expected.SetString(tt.expected, 10)

			result := RoundDown(input)

			if result.Cmp(expected) != 0 {
				t.Errorf("RoundDown(%s) = %s; want %s", tt.input, result.String(), tt.expected)
			}
		})
	}
}

func Test_ToStr(t *testing.T) {
	tests := []struct {
		name        string
		input       interface{} // *big.Int or *big.Float
		places      []int       // decimal places
		expectedStr string
	}{
		{"It should convert big.Int to string", new(big.Int).SetInt64(12345), nil, "12345"},
		{"It should convert negative big.Int to string", new(big.Int).SetInt64(-98765), nil, "-98765"},
		{"It should convert big.Float to string with default precision", new(big.Float).SetPrec(128).SetFloat64(123.456789), []int{6}, "123.456789"},
		{"It should convert negative big.Float to string with default precision", new(big.Float).SetPrec(128).SetFloat64(-987.654321), []int{5}, "-987.65432"},
		{"It should convert big.Float to string with specified precision", new(big.Float).SetFloat64(123.456789), []int{2}, "123.46"},
		{"It should convert negative big.Float to string with specified precision", new(big.Float).SetFloat64(-987.654321), []int{3}, "-987.654"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result string
			switch v := tt.input.(type) {
			case *big.Int:
				result = ToStr(v)
			case *big.Float:
				result = ToStr(v, tt.places...)
			default:
				t.Fatalf("unsupported type in test case: %T", tt.input)
			}

			if result != tt.expectedStr {
				t.Errorf("ToStr(%v, %v) = %s; want %s", tt.input, tt.places, result, tt.expectedStr)
			}
		})
	}
}

func Test_NewCompareInt(t *testing.T) {
	tests := []struct {
		name        string
		actual      *big.Int
		expected    string
		expectedStr string
		shouldEqual bool
	}{
		{"It should match equal integers", new(big.Int).SetInt64(12345), "12345", "12345", true},
		{"It should match equal negative integers", new(big.Int).SetInt64(-98765), "-98765", "-98765", true},
		{"It should detect mismatched integers", new(big.Int).SetInt64(12345), "54321", "12345", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			compare := NewCompare[big.Int](tt.actual, tt.expected)

			if compare.ActualAsString != tt.expectedStr {
				t.Errorf("NewCompare(%v, %s).ActualAsString = %s; want %s", tt.actual, tt.expected, compare.ActualAsString, tt.expectedStr)
			}

			if compare.Equal() != tt.shouldEqual {
				t.Errorf("NewCompare(%v, %s).Equal() = %v; want %v", tt.actual, tt.expected, compare.Equal(), tt.shouldEqual)
			}
		})
	}
}

func Test_NewCompareFloat(t *testing.T) {
	tests := []struct {
		name        string
		actual      *big.Float
		expected    string
		expectedStr string
		shouldEqual bool
	}{
		{"It should match equal floats", StrToFloat("123.456"), "123.456", "123.456", true},
		{"It should match equal negative floats", StrToFloat("-987.654"), "-987.654", "-987.654", true},
		{"It should detect mismatched floats", StrToFloat("123.456"), "123.45", "123.456", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			compare := NewCompare(tt.actual, tt.expected)

			if tt.shouldEqual && compare.ActualAsString != tt.expectedStr {
				t.Errorf("NewCompare(%v, %s).ActualAsString = %s; want %s", tt.actual, tt.expected, compare.ActualAsString, tt.expectedStr)
			}

			if compare.Equal() != tt.shouldEqual {
				t.Errorf("NewCompare(%v, %s).Equal() = %v; want %v", tt.actual, tt.expected, compare.Equal(), tt.shouldEqual)
			}
		})
	}
}

func Test_StrBigCompareEqual(t *testing.T) {
	tests := []struct {
		name           string
		actualAsString string
		expected       string
		shouldEqual    bool
	}{
		{"It should match equal integers as strings", "12345", "12345", true},
		{"It should match equal negative integers as strings", "-98765", "-98765", true},
		{"It should match equal floats as strings", "123.456", "123.456", true},
		{"It should match equal negative floats as strings", "-987.654", "-987.654", true},
		{"It should detect mismatched floats as strings", "123.456", "123.45", false},
		{"It should detect mismatched integers as strings", "12345", "54321", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			container := StrBigCompare[big.Float]{
				ActualAsString: tt.actualAsString,
				Expected:       tt.expected,
			}

			if container.Equal() != tt.shouldEqual {
				t.Errorf("StrBigCompare{ActualAsString: %s, Expected: %s}.Equal() = %v; want %v",
					tt.actualAsString, tt.expected, container.Equal(), tt.shouldEqual)
			}
		})
	}
}

func Test_RoundUp(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"It should round 1.1 to 2", "1.1", "2"},
		{"It should round 2.9 to 3", "2.9", "3"},
		{"It should not round integers", "0.0", "0"},
		{"It should not round negative integers", "-1", "-1"},
		{"It should round -2.9 to -2", "-2.9", "-2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := new(big.Float)
			input.SetString(tt.input)

			expected := new(big.Int)
			expected.SetString(tt.expected, 10)

			result := RoundUp(input)

			if result.Cmp(expected) != 0 {
				t.Errorf("RoundUp(%s) = %s; want %s", tt.input, result.String(), tt.expected)
			}
		})
	}
}
