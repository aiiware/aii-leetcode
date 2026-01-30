package math

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMySqrt(t *testing.T) {
	tests := []struct {
		name     string
		x        int
		expected int
	}{
		{
			name:     "Example 1: x=4",
			x:        4,
			expected: 2,
		},
		{
			name:     "Example 2: x=8",
			x:        8,
			expected: 2,
		},
		{
			name:     "x=0",
			x:        0,
			expected: 0,
		},
		{
			name:     "x=1",
			x:        1,
			expected: 1,
		},
		{
			name:     "x=2",
			x:        2,
			expected: 1,
		},
		{
			name:     "x=3",
			x:        3,
			expected: 1,
		},
		{
			name:     "x=9",
			x:        9,
			expected: 3,
		},
		{
			name:     "x=10",
			x:        10,
			expected: 3,
		},
		{
			name:     "x=15",
			x:        15,
			expected: 3,
		},
		{
			name:     "x=16",
			x:        16,
			expected: 4,
		},
		{
			name:     "x=25",
			x:        25,
			expected: 5,
		},
		{
			name:     "x=100",
			x:        100,
			expected: 10,
		},
		{
			name:     "x=1000",
			x:        1000,
			expected: 31, // 31^2=961, 32^2=1024
		},
		{
			name:     "x=10000",
			x:        10000,
			expected: 100,
		},
		{
			name:     "x=2147395599",
			x:        2147395599,
			expected: 46339, // Largest test case before overflow
		},
		{
			name:     "x=2147483647 (max int32)",
			x:        2147483647,
			expected: 46340, // sqrt(2^31-1) ≈ 46340.95
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MySqrt(tt.x)
			assert.Equal(t, tt.expected, result,
				"MySqrt(%d) = %d, expected %d",
				tt.x, result, tt.expected)
		})
	}
}

func TestMySqrt_EdgeCases(t *testing.T) {
	t.Run("Negative x returns -1", func(t *testing.T) {
		result := MySqrt(-1)
		assert.Equal(t, -1, result)
	})

	t.Run("Large perfect square", func(t *testing.T) {
		// Test perfect squares up to a reasonable limit
		for i := 0; i <= 10000; i++ {
			perfectSquare := i * i
			if perfectSquare < 0 {
				break // Overflow
			}
			result := MySqrt(perfectSquare)
			assert.Equal(t, i, result,
				"MySqrt(%d) = %d, expected %d",
				perfectSquare, result, i)
		}
	})

	t.Run("Numbers just below perfect squares", func(t *testing.T) {
		for i := 1; i <= 100; i++ {
			justBelow := i*i - 1
			result := MySqrt(justBelow)
			assert.Equal(t, i-1, result,
				"MySqrt(%d) = %d, expected %d",
				justBelow, result, i-1)
		}
	})

	t.Run("Numbers just above perfect squares", func(t *testing.T) {
		for i := 1; i <= 100; i++ {
			justAbove := i*i + 1
			result := MySqrt(justAbove)
			assert.Equal(t, i, result,
				"MySqrt(%d) = %d, expected %d",
				justAbove, result, i)
		}
	})
}

func TestMySqrt_PropertyBased(t *testing.T) {
	// Test that MySqrt(x) satisfies: result^2 <= x < (result+1)^2
	t.Run("Satisfies square root property", func(t *testing.T) {
		for x := 0; x <= 1000000; x += 1237 { // Skip some values for speed
			result := MySqrt(x)
			
			// Check lower bound: result^2 <= x
			lowerBound := result * result
			assert.True(t, lowerBound <= x,
				"For x=%d, result=%d: %d^2=%d should be <= %d",
				x, result, result, lowerBound, x)
			
			// Check upper bound: x < (result+1)^2
			// Be careful about overflow
			next := result + 1
			var upperBound int
			if next <= math.MaxInt64/next {
				upperBound = next * next
				assert.True(t, x < upperBound,
					"For x=%d, result=%d: %d should be < %d^2=%d",
					x, result, x, next, upperBound)
			}
		}
	})

	// Test that MySqrt is monotonic: if a < b, then sqrt(a) <= sqrt(b)
	t.Run("Monotonic property", func(t *testing.T) {
		for a := 0; a < 10000; a += 137 {
			for b := a + 1; b < 10000 && b < a+100; b += 29 {
				sqrtA := MySqrt(a)
				sqrtB := MySqrt(b)
				assert.True(t, sqrtA <= sqrtB,
					"For %d < %d, sqrt(%d)=%d should be <= sqrt(%d)=%d",
					a, b, a, sqrtA, b, sqrtB)
			}
		}
	})

	// Test that MySqrt matches math.Sqrt for small values
	t.Run("Matches math.Sqrt for small values", func(t *testing.T) {
		for x := 0; x <= 10000; x += 73 {
			expected := int(math.Sqrt(float64(x)))
			result := MySqrt(x)
			assert.Equal(t, expected, result,
				"MySqrt(%d) = %d, math.Sqrt = %d",
				x, result, expected)
		}
	})
}

func TestMySqrt_Performance(t *testing.T) {
	// Test that it handles large numbers without overflow
	t.Run("No overflow for large numbers", func(t *testing.T) {
		largeNumbers := []int{
			1000000000,
			2147483647, // Max int32
			10000000000, // Larger than int32 but still fits in int64
		}

		for _, x := range largeNumbers {
			result := MySqrt(x)
			// Just verify it doesn't panic and returns a reasonable value
			assert.True(t, result >= 0,
				"MySqrt(%d) = %d should be non-negative",
				x, result)
			assert.True(t, result*result <= x,
				"MySqrt(%d) = %d: %d^2=%d should be <= %d",
				x, result, result, result*result, x)
		}
	})
}

func BenchmarkMySqrt(b *testing.B) {
	testCases := []struct {
		name string
		x    int
	}{
		{"Small", 100},
		{"Medium", 1000000},
		{"Large", 1000000000},
		{"Very large", 2147483647},
		{"Perfect square", 100000000},
		{"Near perfect square", 99999999},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MySqrt(tc.x)
			}
		})
	}
}