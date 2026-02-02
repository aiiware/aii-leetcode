package math

import (
	"testing"
)

func TestRangeBitwiseAnd(t *testing.T) {
	tests := []struct {
		left     int
		right    int
		expected int
	}{
		{5, 7, 4},           // 5: 101, 6: 110, 7: 111 -> AND = 100 (4)
		{0, 0, 0},           // Single number 0
		{1, 1, 1},           // Single number 1
		{1, 2147483647, 0},  // Large range, result should be 0
		{10, 11, 10},        // 10: 1010, 11: 1011 -> AND = 1010 (10)
		{12, 15, 12},        // 12: 1100, 13: 1101, 14: 1110, 15: 1111 -> AND = 1100 (12)
		{6, 7, 6},           // 6: 110, 7: 111 -> AND = 110 (6)
	}

	for _, test := range tests {
		result := RangeBitwiseAnd(test.left, test.right)
		if result != test.expected {
			t.Errorf("RangeBitwiseAnd(%d, %d) = %d, expected %d", test.left, test.right, result, test.expected)
		}

		// Also test the alternative implementation
		result2 := RangeBitwiseAnd2(test.left, test.right)
		if result2 != test.expected {
			t.Errorf("RangeBitwiseAnd2(%d, %d) = %d, expected %d", test.left, test.right, result2, test.expected)
		}
	}
}

func TestRangeBitwiseAnd_EdgeCases(t *testing.T) {
	// Test with minimum and maximum values
	tests := []struct {
		left     int
		right    int
		expected int
	}{
		{0, 1, 0},                    // 0: 0, 1: 1 -> AND = 0
		{2147483646, 2147483647, 2147483646}, // Adjacent large numbers
		{100, 200, 0},               // Random medium range - result should be 0
	}

	for _, test := range tests {
		result := RangeBitwiseAnd(test.left, test.right)
		if result != test.expected {
			t.Errorf("RangeBitwiseAnd(%d, %d) = %d, expected %d", test.left, test.right, result, test.expected)
		}
	}
}

func BenchmarkRangeBitwiseAnd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RangeBitwiseAnd(1, 1000000)
	}
}

func BenchmarkRangeBitwiseAnd2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RangeBitwiseAnd2(1, 1000000)
	}
}