package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "Example 1: n=2",
			n:        2,
			expected: 2,
		},
		{
			name:     "Example 2: n=3",
			n:        3,
			expected: 3,
		},
		{
			name:     "n=1",
			n:        1,
			expected: 1,
		},
		{
			name:     "n=0",
			n:        0,
			expected: 0,
		},
		{
			name:     "n=4",
			n:        4,
			expected: 5, // 1+1+1+1, 1+1+2, 1+2+1, 2+1+1, 2+2
		},
		{
			name:     "n=5",
			n:        5,
			expected: 8, // Fibonacci: 1,2,3,5,8
		},
		{
			name:     "n=6",
			n:        6,
			expected: 13,
		},
		{
			name:     "n=7",
			n:        7,
			expected: 21,
		},
		{
			name:     "n=8",
			n:        8,
			expected: 34,
		},
		{
			name:     "n=9",
			n:        9,
			expected: 55,
		},
		{
			name:     "n=10",
			n:        10,
			expected: 89,
		},
		{
			name:     "n=20",
			n:        20,
			expected: 10946,
		},
		{
			name:     "n=30",
			n:        30,
			expected: 1346269,
		},
		{
			name:     "n=45",
			n:        45,
			expected: 1836311903,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ClimbStairs(tt.n)
			assert.Equal(t, tt.expected, result,
				"ClimbStairs(%d) = %d, expected %d",
				tt.n, result, tt.expected)
		})
	}
}

func TestClimbStairs_EdgeCases(t *testing.T) {
	t.Run("Negative n returns 0", func(t *testing.T) {
		result := ClimbStairs(-1)
		assert.Equal(t, 0, result)
	})

	t.Run("Large n doesn't overflow", func(t *testing.T) {
		// Test up to n=46 (Fibonacci(46) fits in int32)
		// Fibonacci sequence: 0,1,1,2,3,5,8,13,21,34,55,89,144,...
		// We need f(46) which should be 1836311903
		result := ClimbStairs(46)
		// f(45) = 1836311903, f(46) = 2971215073 which overflows int32
		// But Go int is platform dependent, so just check it's positive
		assert.True(t, result > 0, "Result should be positive for n=46")
	})

	t.Run("Fibonacci property", func(t *testing.T) {
		// Verify that f(n) = f(n-1) + f(n-2) for n >= 3
		// Note: For n=2, f(2) = 2, f(1) = 1, f(0) = 0, so 2 != 1 + 0
		// This is because f(0) is defined as 0 for this problem (0 ways to climb 0 stairs)
		// But mathematically, if we define f(0) = 1, then recurrence holds for all n >= 2
		for n := 3; n <= 20; n++ {
			fn := ClimbStairs(n)
			fn1 := ClimbStairs(n - 1)
			fn2 := ClimbStairs(n - 2)
			assert.Equal(t, fn, fn1+fn2,
				"ClimbStairs(%d) = %d should equal ClimbStairs(%d) + ClimbStairs(%d) = %d + %d = %d",
				n, fn, n-1, n-2, fn1, fn2, fn1+fn2)
		}
	})
}

func TestClimbStairs_PropertyBased(t *testing.T) {
	// Test that ClimbStairs follows Fibonacci sequence (shifted by 1)
	// Fibonacci: 0,1,1,2,3,5,8,13,21,34,55,89,...
	// ClimbStairs: f(0)=0, f(1)=1, f(2)=2, f(3)=3, f(4)=5, f(5)=8,...
	// So ClimbStairs(n) = Fibonacci(n+1) for n >= 1
	t.Run("Matches Fibonacci sequence", func(t *testing.T) {
		// Pre-calculated Fibonacci numbers
		fibonacci := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765}
		
		for i := 0; i < len(fibonacci)-1; i++ {
			n := i
			expected := fibonacci[i+1] // ClimbStairs(n) = Fibonacci(n+1)
			if n == 0 {
				expected = 0 // Our implementation returns 0 for n=0
			}
			result := ClimbStairs(n)
			assert.Equal(t, expected, result,
				"ClimbStairs(%d) = %d, Fibonacci(%d)=%d",
				n, result, n+1, expected)
		}
	})

	// Test monotonic property: f(n) > f(n-1) for n >= 2
	t.Run("Monotonic increasing", func(t *testing.T) {
		for n := 2; n <= 20; n++ {
			fn := ClimbStairs(n)
			fn1 := ClimbStairs(n - 1)
			assert.True(t, fn > fn1,
				"ClimbStairs(%d) = %d should be > ClimbStairs(%d) = %d",
				n, fn, n-1, fn1)
		}
	})

	// Test that number of ways grows exponentially
	t.Run("Exponential growth", func(t *testing.T) {
		// For large n, f(n) ≈ φ^n / √5 where φ = (1+√5)/2 ≈ 1.618
		// So f(n) should be roughly φ times f(n-1)
		for n := 10; n <= 20; n++ {
			fn := ClimbStairs(n)
			fn1 := ClimbStairs(n - 1)
			ratio := float64(fn) / float64(fn1)
			// Ratio should approach φ ≈ 1.618
			assert.True(t, ratio > 1.5 && ratio < 1.7,
				"Ratio f(%d)/f(%d) = %f should be close to φ ≈ 1.618",
				n, n-1, ratio)
		}
	})
}

func TestClimbStairs_AlternativeApproaches(t *testing.T) {
	// Test that our implementation matches recursive definition
	t.Run("Matches recursive definition", func(t *testing.T) {
		// Simple recursive implementation for verification
		var recursiveClimbStairs func(int) int
		recursiveClimbStairs = func(n int) int {
			if n <= 0 {
				return 0
			}
			if n == 1 {
				return 1
			}
			if n == 2 {
				return 2
			}
			return recursiveClimbStairs(n-1) + recursiveClimbStairs(n-2)
		}

		for n := 0; n <= 20; n++ {
			iterativeResult := ClimbStairs(n)
			recursiveResult := recursiveClimbStairs(n)
			assert.Equal(t, recursiveResult, iterativeResult,
				"For n=%d, iterative=%d, recursive=%d",
				n, iterativeResult, recursiveResult)
		}
	})
}

func BenchmarkClimbStairs(b *testing.B) {
	testCases := []struct {
		name string
		n    int
	}{
		{"Small", 10},
		{"Medium", 30},
		{"Large", 45},
		{"Very large", 60},
		{"Maximum practical", 90},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ClimbStairs(tc.n)
			}
		})
	}
}