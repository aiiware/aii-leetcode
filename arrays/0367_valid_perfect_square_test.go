package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPerfectSquare(t *testing.T) {
	tests := []struct {
		name     string
		num      int
		expected bool
	}{
		{
			name:     "Example 1: 16 is perfect square",
			num:      16,
			expected: true,
		},
		{
			name:     "Example 2: 14 is not perfect square",
			num:      14,
			expected: false,
		},
		{
			name:     "Zero is perfect square",
			num:      0,
			expected: true,
		},
		{
			name:     "One is perfect square",
			num:      1,
			expected: true,
		},
		{
			name:     "Small perfect square: 4",
			num:      4,
			expected: true,
		},
		{
			name:     "Small perfect square: 9",
			num:      9,
			expected: true,
		},
		{
			name:     "Small perfect square: 25",
			num:      25,
			expected: true,
		},
		{
			name:     "Small non-perfect square: 2",
			num:      2,
			expected: false,
		},
		{
			name:     "Small non-perfect square: 3",
			num:      3,
			expected: false,
		},
		{
			name:     "Small non-perfect square: 5",
			num:      5,
			expected: false,
		},
		{
			name:     "Medium perfect square: 100",
			num:      100,
			expected: true,
		},
		{
			name:     "Medium perfect square: 144",
			num:      144,
			expected: true,
		},
		{
			name:     "Medium non-perfect square: 101",
			num:      101,
			expected: false,
		},
		{
			name:     "Medium non-perfect square: 143",
			num:      143,
			expected: false,
		},
		{
			name:     "Large perfect square: 10000",
			num:      10000,
			expected: true,
		},
		{
			name:     "Large perfect square: 104976", // 324²
			num:      104976,
			expected: true,
		},
		{
			name:     "Large non-perfect square: 10001",
			num:      10001,
			expected: false,
		},
		{
			name:     "Large non-perfect square: 104975",
			num:      104975,
			expected: false,
		},
		{
			name:     "Very large perfect square: 46340² = 2147395600",
			num:      2147395600,
			expected: true,
		},
		{
			name:     "Very large non-perfect square: 2147395601",
			num:      2147395601,
			expected: false,
		},
		{
			name:     "Negative number",
			num:      -4,
			expected: false,
		},
		{
			name:     "Perfect square of prime: 289 (17²)",
			num:      289,
			expected: true,
		},
		{
			name:     "Perfect square of large prime: 5329 (73²)",
			num:      5329,
			expected: true,
		},
		{
			name:     "Number just below perfect square: 255 (16² = 256)",
			num:      255,
			expected: false,
		},
		{
			name:     "Number just above perfect square: 257 (16² = 256)",
			num:      257,
			expected: false,
		},
		{
			name:     "Max int32 perfect square: 46340² = 2147395600",
			num:      2147395600,
			expected: true,
		},
		{
			name:     "Max int32 value: 2147483647 (not a perfect square)",
			num:      2147483647,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPerfectSquare(tt.num)
			assert.Equal(t, tt.expected, result,
				"IsPerfectSquare(%d) = %v, expected %v",
				tt.num, result, tt.expected)
		})
	}
}

func TestIsPerfectSquareOptimized(t *testing.T) {
	tests := []struct {
		name     string
		num      int
		expected bool
	}{
		{
			name:     "16 is perfect square",
			num:      16,
			expected: true,
		},
		{
			name:     "14 is not perfect square",
			num:      14,
			expected: false,
		},
		{
			name:     "Zero is perfect square",
			num:      0,
			expected: true,
		},
		{
			name:     "One is perfect square",
			num:      1,
			expected: true,
		},
		{
			name:     "Large perfect square: 10000",
			num:      10000,
			expected: true,
		},
		{
			name:     "Large non-perfect square: 10001",
			num:      10001,
			expected: false,
		},
		{
			name:     "Max int32 perfect square",
			num:      2147395600,
			expected: true,
		},
		{
			name:     "Max int32 value",
			num:      2147483647,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPerfectSquareOptimized(tt.num)
			assert.Equal(t, tt.expected, result,
				"IsPerfectSquareOptimized(%d) = %v, expected %v",
				tt.num, result, tt.expected)
		})
	}
}

func TestIsPerfectSquare_Consistency(t *testing.T) {
	// Test that both implementations give the same result
	testCases := []int{
		0, 1, 2, 3, 4, 5, 9, 10, 16, 25, 100, 101, 144, 145,
		10000, 10001, 2147395600, 2147395601, 2147483647,
	}

	for _, num := range testCases {
		t.Run("", func(t *testing.T) {
			result1 := IsPerfectSquare(num)
			result2 := IsPerfectSquareOptimized(num)
			assert.Equal(t, result1, result2,
				"Both implementations should give same result for %d: %v vs %v",
				num, result1, result2)
		})
	}
}

func TestIsPerfectSquare_EdgeCases(t *testing.T) {
	t.Run("Negative numbers", func(t *testing.T) {
		assert.False(t, IsPerfectSquare(-1))
		assert.False(t, IsPerfectSquare(-4))
		assert.False(t, IsPerfectSquare(-100))
	})

	t.Run("Very small numbers", func(t *testing.T) {
		assert.True(t, IsPerfectSquare(0))
		assert.True(t, IsPerfectSquare(1))
		assert.False(t, IsPerfectSquare(2))
		assert.False(t, IsPerfectSquare(3))
		assert.True(t, IsPerfectSquare(4))
	})

	t.Run("Perfect squares of consecutive integers", func(t *testing.T) {
		// Test squares from 0 to 100
		for i := 0; i <= 100; i++ {
			square := i * i
			assert.True(t, IsPerfectSquare(square),
				"%d should be a perfect square (%d²)", square, i)

			// Numbers just before and after perfect squares should be false
			// Skip checking square-1 when square = 1 because 0 is a perfect square
			if square > 1 {
				assert.False(t, IsPerfectSquare(square-1),
					"%d should not be a perfect square", square-1)
			}
			// Skip checking square+1 when square = 0 because 1 is a perfect square
			if square > 0 && square < 10000 { // Avoid overflow
				assert.False(t, IsPerfectSquare(square+1),
					"%d should not be a perfect square", square+1)
			}
		}
	})
}

func BenchmarkIsPerfectSquare(b *testing.B) {
	testCases := []int{
		16, 100, 10000, 1000000, 2147395600,
	}

	for _, num := range testCases {
		b.Run("", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsPerfectSquare(num)
			}
		})
	}
}

func BenchmarkIsPerfectSquareOptimized(b *testing.B) {
	testCases := []int{
		16, 100, 10000, 1000000, 2147395600,
	}

	for _, num := range testCases {
		b.Run("", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsPerfectSquareOptimized(num)
			}
		})
	}
}

func BenchmarkIsPerfectSquare_WorstCase(b *testing.B) {
	// Worst case: number is not a perfect square and binary search goes to the end
	num := 2147483647 // Max int32, not a perfect square

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsPerfectSquare(num)
	}
}

func BenchmarkIsPerfectSquare_BestCase(b *testing.B) {
	// Best case: number is a perfect square found quickly
	num := 65536 // 256²

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsPerfectSquare(num)
	}
}