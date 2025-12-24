package leetcode

import (
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		dividend int
		divisor  int
		expected int
	}{
		{10, 3, 3},
		{7, -3, -2},
		{0, 1, 0},
		{1, 1, 1},
		{-1, 1, -1},
		{1, -1, -1},
		{-1, -1, 1},
		{2147483647, 1, 2147483647},
		{-2147483648, -1, 2147483647}, // Overflow case
		{-2147483648, 1, -2147483648},
		{100, 25, 4},
		{100, -25, -4},
		{-100, 25, -4},
		{-100, -25, 4},
		{1, 2, 0},
		{2, 3, 0},
		{5, 2, 2},
	}

	for i, test := range tests {
		// Test bit manipulation approach
		result := Divide(test.dividend, test.divisor)

		if result != test.expected {
			t.Errorf("Test %d (bit manipulation) failed: %d / %d = %d, expected %d",
				i, test.dividend, test.divisor, result, test.expected)
		}

		// Test simple subtraction approach
		result = DivideSimple(test.dividend, test.divisor)

		if result != test.expected {
			t.Errorf("Test %d (simple) failed: %d / %d = %d, expected %d",
				i, test.dividend, test.divisor, result, test.expected)
		}
	}
}

func BenchmarkDivide(b *testing.B) {
	dividend := 2147483647
	divisor := 3

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Divide(dividend, divisor)
	}
}

func BenchmarkDivideSimple(b *testing.B) {
	dividend := 2147483647
	divisor := 3

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DivideSimple(dividend, divisor)
	}
}