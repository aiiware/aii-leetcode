package leetcode

import (
	"testing"
)

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 3, -2, 4},
			expected: 6,
		},
		{
			name:     "Example 2",
			nums:     []int{-2, 0, -1},
			expected: 0,
		},
		{
			name:     "All positive",
			nums:     []int{1, 2, 3, 4},
			expected: 24,
		},
		{
			name:     "All negative odd count",
			nums:     []int{-2, -3, -1},
			expected: 6,
		},
		{
			name:     "All negative even count",
			nums:     []int{-2, -3, -1, -4},
			expected: 24,
		},
		{
			name:     "Single element positive",
			nums:     []int{5},
			expected: 5,
		},
		{
			name:     "Single element negative",
			nums:     []int{-5},
			expected: -5,
		},
		{
			name:     "With zeros",
			nums:     []int{-2, 0, -1, 0, 3, 2},
			expected: 6,
		},
		{
			name:     "Mixed positive and negative",
			nums:     []int{2, -5, 3, 1, -4, 0, -10, 2, 8},
			expected: 120,
		},
		{
			name:     "Large negative at end",
			nums:     []int{1, -2, 3, -4, 5, -6},
			expected: 360,
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "Alternating positive negative",
			nums:     []int{1, -2, 3, -4, 5},
			expected: 120,
		},
		{
			name:     "LeetCode test case 1",
			nums:     []int{1, 2, -3, 4},
			expected: 4,
		},
		{
			name:     "LeetCode test case 2",
			nums:     []int{-2, -1},
			expected: 2,
		},
		{
			name:     "Large numbers",
			nums:     []int{-10, 10, -10, 10},
			expected: 10000,
		},
		{
			name:     "Product resets at zero",
			nums:     []int{2, 3, 0, 4, 5},
			expected: 20,
		},
		{
			name:     "Negative prefix",
			nums:     []int{-2, 3, 4},
			expected: 12,
		},
		{
			name:     "Negative suffix",
			nums:     []int{3, 4, -2},
			expected: 12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxProduct(tt.nums)
			if result != tt.expected {
				t.Errorf("MaxProduct(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestMaxProductKadane(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Basic case",
			nums:     []int{2, 3, -2, 4},
			expected: 6,
		},
		{
			name:     "With zero",
			nums:     []int{-2, 0, -1},
			expected: 0,
		},
		{
			name:     "All negative",
			nums:     []int{-2, -3, -1},
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxProductKadane(tt.nums)
			if result != tt.expected {
				t.Errorf("MaxProductKadane(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestMaxProductPrefixSuffix(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Basic case",
			nums:     []int{2, 3, -2, 4},
			expected: 6,
		},
		{
			name:     "With zero",
			nums:     []int{-2, 0, -1},
			expected: 0,
		},
		{
			name:     "All negative even count",
			nums:     []int{-2, -3, -1, -4},
			expected: 24,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxProductPrefixSuffix(tt.nums)
			if result != tt.expected {
				t.Errorf("MaxProductPrefixSuffix(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestMaxProductBruteForce(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Small array",
			nums:     []int{2, 3, -2, 4},
			expected: 6,
		},
		{
			name:     "Single element",
			nums:     []int{-5},
			expected: -5,
		},
		{
			name:     "Three elements",
			nums:     []int{1, 2, 3},
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxProductBruteForce(tt.nums)
			if result != tt.expected {
				t.Errorf("MaxProductBruteForce(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestMaxProductDP(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Basic case",
			nums:     []int{2, 3, -2, 4},
			expected: 6,
		},
		{
			name:     "With zero",
			nums:     []int{-2, 0, -1},
			expected: 0,
		},
		{
			name:     "Mixed signs",
			nums:     []int{2, -5, 3, 1, -4, 0, -10, 2, 8},
			expected: 120,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxProductDP(tt.nums)
			if result != tt.expected {
				t.Errorf("MaxProductDP(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestAllImplementationsMatch(t *testing.T) {
	testCases := [][]int{
		{2, 3, -2, 4},
		{-2, 0, -1},
		{1, 2, 3, 4},
		{-2, -3, -1},
		{5},
		{-5},
		{0, 0, 0, 0},
		{1, -2, 3, -4, 5},
		{-10, 10, -10, 10},
		{2, 3, 0, 4, 5},
		{},
		{1},
		{-1},
		{0},
		{1, 0, 2},
		{-1, 0, -2},
	}

	for _, nums := range testCases {
		if len(nums) == 0 {
			continue // Skip empty array for implementations that don't handle it
		}

		result1 := MaxProductKadane(nums)
		result2 := MaxProductPrefixSuffix(nums)
		result3 := MaxProductDP(nums)
		
		// Only test brute force for small arrays (it's O(n^2))
		if len(nums) <= 10 {
			result4 := MaxProductBruteForce(nums)
			if result1 != result2 || result1 != result3 || result1 != result4 {
				t.Errorf("Implementations differ for %v: Kadane=%d, PrefixSuffix=%d, DP=%d, BruteForce=%d",
					nums, result1, result2, result3, result4)
			}
		} else {
			if result1 != result2 || result1 != result3 {
				t.Errorf("Implementations differ for %v: Kadane=%d, PrefixSuffix=%d, DP=%d",
					nums, result1, result2, result3)
			}
		}
	}
}

func BenchmarkMaxProductKadane(b *testing.B) {
	nums := []int{2, 3, -2, 4, 5, -1, 0, 3, 2, -4, 1, 2, 3, -5, 6, 7, -2, 3, 4, 1}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxProductKadane(nums)
	}
}

func BenchmarkMaxProductPrefixSuffix(b *testing.B) {
	nums := []int{2, 3, -2, 4, 5, -1, 0, 3, 2, -4, 1, 2, 3, -5, 6, 7, -2, 3, 4, 1}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxProductPrefixSuffix(nums)
	}
}

func BenchmarkMaxProductDP(b *testing.B) {
	nums := []int{2, 3, -2, 4, 5, -1, 0, 3, 2, -4, 1, 2, 3, -5, 6, 7, -2, 3, 4, 1}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxProductDP(nums)
	}
}

func BenchmarkMaxProductBruteForce(b *testing.B) {
	// Use smaller array for brute force benchmark
	nums := []int{2, 3, -2, 4, 5, -1, 0, 3}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxProductBruteForce(nums)
	}
}

func TestMaxProductEdgeCases(t *testing.T) {
	t.Run("Empty array", func(t *testing.T) {
		// All implementations should handle empty array
		result := MaxProductKadane([]int{})
		if result != 0 {
			t.Errorf("MaxProductKadane([]) = %d, expected 0", result)
		}

		result2 := MaxProductPrefixSuffix([]int{})
		if result2 != 0 {
			t.Errorf("MaxProductPrefixSuffix([]) = %d, expected 0", result2)
		}

		result3 := MaxProductDP([]int{})
		if result3 != 0 {
			t.Errorf("MaxProductDP([]) = %d, expected 0", result3)
		}

		result4 := MaxProductBruteForce([]int{})
		if result4 != 0 {
			t.Errorf("MaxProductBruteForce([]) = %d, expected 0", result4)
		}
	})

	t.Run("Single zero", func(t *testing.T) {
		result := MaxProduct([]int{0})
		if result != 0 {
			t.Errorf("MaxProduct([0]) = %d, expected 0", result)
		}
	})

	t.Run("Large array with pattern", func(t *testing.T) {
		// Create array: [1, -1, 1, -1, ...] of length 100
		nums := make([]int, 100)
		for i := range nums {
			if i%2 == 0 {
				nums[i] = 1
			} else {
				nums[i] = -1
			}
		}
		
		// For even length array of alternating 1 and -1, max product is 1
		result := MaxProduct(nums)
		if result != 1 {
			t.Errorf("MaxProduct(alternating 1,-1 x100) = %d, expected 1", result)
		}
	})

	t.Run("All ones", func(t *testing.T) {
		nums := make([]int, 50)
		for i := range nums {
			nums[i] = 1
		}
		
		result := MaxProduct(nums)
		if result != 1 {
			t.Errorf("MaxProduct(all ones x50) = %d, expected 1", result)
		}
	})

	t.Run("Product overflow test", func(t *testing.T) {
		// Test with values that would overflow 32-bit if multiplied
		// but problem guarantees product fits in 32-bit
		nums := []int{10, 10, 10, 10, 10, 10} // 10^6 = 1,000,000 fits in 32-bit
		result := MaxProduct(nums)
		expected := 1000000
		if result != expected {
			t.Errorf("MaxProduct([10,10,10,10,10,10]) = %d, expected %d", result, expected)
		}
	})
}