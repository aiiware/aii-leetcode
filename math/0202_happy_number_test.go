package math

import (
	"testing"
)

func TestIsHappy(t *testing.T) {
	tests := []struct {
		n        int
		expected bool
	}{
		{19, true},   // 19 is happy: 1²+9²=82, 8²+2²=68, 6²+8²=100, 1²+0²+0²=1
		{2, false},   // 2 is not happy (enters cycle: 4, 16, 37, 58, 89, 145, 42, 20, 4)
		{1, true},    // 1 is happy by definition
		{7, true},    // 7 is happy
		{10, true},   // 10 is happy: 1²+0²=1
		{13, true},   // 13 is happy: 1²+3²=10 -> 1
		{4, false},   // 4 is not happy (enters cycle)
		{100, true},  // 100 is happy: 1²+0²+0²=1
		{999, false}, // 999 is not happy
	}

	for _, test := range tests {
		result := IsHappy(test.n)
		if result != test.expected {
			t.Errorf("IsHappy(%d) = %v, expected %v", test.n, result, test.expected)
		}

		// Test Floyd's algorithm
		result2 := IsHappyFloyd(test.n)
		if result2 != test.expected {
			t.Errorf("IsHappyFloyd(%d) = %v, expected %v", test.n, result2, test.expected)
		}

		// Test mathematical approach
		result3 := IsHappyMath(test.n)
		if result3 != test.expected {
			t.Errorf("IsHappyMath(%d) = %v, expected %v", test.n, result3, test.expected)
		}
	}
}

func TestSumOfSquares(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{19, 82},   // 1² + 9² = 1 + 81 = 82
		{100, 1},   // 1² + 0² + 0² = 1 + 0 + 0 = 1
		{123, 14},  // 1² + 2² + 3² = 1 + 4 + 9 = 14
		{7, 49},    // 7² = 49
		{0, 0},     // 0 has no digits
		{999, 243}, // 9² + 9² + 9² = 81 + 81 + 81 = 243
	}

	for _, test := range tests {
		result := sumOfSquares(test.n)
		if result != test.expected {
			t.Errorf("sumOfSquares(%d) = %d, expected %d", test.n, result, test.expected)
		}
	}
}

func TestIsHappy_EdgeCases(t *testing.T) {
	// Test large numbers
	tests := []struct {
		n        int
		expected bool
	}{
		{2147483647, false}, // Maximum 32-bit signed integer
		{1000000, true},     // 1,000,000 -> 1
		{9999999, false},    // Large number that's not happy
	}

	for _, test := range tests {
		result := IsHappy(test.n)
		if result != test.expected {
			t.Errorf("IsHappy(%d) = %v, expected %v", test.n, result, test.expected)
		}
	}
}

func BenchmarkIsHappy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsHappy(9999999)
	}
}

func BenchmarkIsHappyFloyd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsHappyFloyd(9999999)
	}
}

func BenchmarkIsHappyMath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsHappyMath(9999999)
	}
}