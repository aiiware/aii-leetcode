package math

import (
	"fmt"
	"testing"
)

func TestTrailingZeroes(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		// Basic examples from problem description
		{0, 0},  // 0! = 1
		{1, 0},  // 1! = 1
		{2, 0},  // 2! = 2
		{3, 0},  // 3! = 6
		{4, 0},  // 4! = 24
		{5, 1},  // 5! = 120
		{6, 1},  // 6! = 720
		{7, 1},  // 7! = 5040
		{8, 1},  // 8! = 40320
		{9, 1},  // 9! = 362880
		{10, 2}, // 10! = 3628800
		
		// Edge cases with multiples of 5
		{14, 2}, // 14! has 2 trailing zeroes
		{15, 3}, // 15! has 3 trailing zeroes (5, 10, 15)
		{19, 3}, // 19! has 3 trailing zeroes
		{20, 4}, // 20! has 4 trailing zeroes (5, 10, 15, 20)
		{24, 4}, // 24! has 4 trailing zeroes
		{25, 6}, // 25! has 6 trailing zeroes (5, 10, 15, 20, 25 contributes 2)
		{29, 6}, // 29! has 6 trailing zeroes
		{30, 7}, // 30! has 7 trailing zeroes
		
		// Larger numbers
		{100, 24},    // 100! has 24 trailing zeroes
		{125, 31},    // 125! has 31 trailing zeroes (125 contributes 3 factors of 5)
		{250, 62},    // 250! has 62 trailing zeroes
		{500, 124},   // 500! has 124 trailing zeroes
		{1000, 249},  // 1000! has 249 trailing zeroes
		{2000, 499},  // 2000! has 499 trailing zeroes
		
		// Maximum constraint value
		{10000, 2499}, // 10000! has 2499 trailing zeroes
		
		// Additional test cases
		{26, 6},  // 26! has 6 trailing zeroes
		{50, 12}, // 50! has 12 trailing zeroes
		{75, 18}, // 75! has 18 trailing zeroes
		{99, 22}, // 99! has 22 trailing zeroes
		{101, 24}, // 101! has 24 trailing zeroes
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test case %d: n=%d", i, tt.n), func(t *testing.T) {
			result := trailingZeroes(tt.n)
			if result != tt.expected {
				t.Errorf("trailingZeroes(%d) = %v, want %v", tt.n, result, tt.expected)
			}
		})
	}
}

// Benchmark tests for performance
func BenchmarkTrailingZeroes(b *testing.B) {
	testCases := []int{0, 5, 10, 25, 100, 1000, 10000}
	
	for _, tc := range testCases {
		b.Run(fmt.Sprintf("n=%d", tc), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				trailingZeroes(tc)
			}
		})
	}
}

// Property-based test: trailingZeroes should be non-decreasing
func TestTrailingZeroesProperty(t *testing.T) {
	// Test that trailingZeroes(n) <= trailingZeroes(n+1) for all n
	// Actually, trailingZeroes should be non-decreasing
	for n := 0; n < 1000; n++ {
		curr := trailingZeroes(n)
		next := trailingZeroes(n + 1)
		if next < curr {
			t.Errorf("trailingZeroes should be non-decreasing: trailingZeroes(%d)=%d > trailingZeroes(%d)=%d",
				n, curr, n+1, next)
		}
	}
}

// Test that our implementation matches the mathematical formula
func TestTrailingZeroesFormula(t *testing.T) {
	// The mathematical formula is: sum_{i=1}^{∞} floor(n / 5^i)
	// We can test this for small values by computing factorial directly
	// (though this is inefficient for large n)
	
	// Test small values where we can compute factorial
	testCases := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	
	for _, n := range testCases {
		// Compute factorial
		fact := 1
		for i := 2; i <= n; i++ {
			fact *= i
		}
		
		// Count trailing zeroes manually
		expected := 0
		temp := fact
		for temp > 0 && temp%10 == 0 {
			expected++
			temp /= 10
		}
		
		result := trailingZeroes(n)
		if result != expected {
			t.Errorf("trailingZeroes(%d) = %v, but %d! = %d has %v trailing zeroes",
				n, result, n, fact, expected)
		}
	}
}