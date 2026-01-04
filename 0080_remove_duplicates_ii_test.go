package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicatesII(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
		expectedNums []int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 1, 1, 2, 2, 3},
			expected: 5,
			expectedNums: []int{1, 1, 2, 2, 3},
		},
		{
			name:     "Example 2",
			nums:     []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			expected: 7,
			expectedNums: []int{0, 0, 1, 1, 2, 3, 3},
		},
		{
			name:     "No duplicates",
			nums:     []int{1, 2, 3, 4, 5},
			expected: 5,
			expectedNums: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "All same elements (more than 2)",
			nums:     []int{1, 1, 1, 1, 1},
			expected: 2,
			expectedNums: []int{1, 1},
		},
		{
			name:     "Exactly 2 duplicates",
			nums:     []int{1, 1, 2, 2, 3, 3},
			expected: 6,
			expectedNums: []int{1, 1, 2, 2, 3, 3},
		},
		{
			name:     "Single element",
			nums:     []int{1},
			expected: 1,
			expectedNums: []int{1},
		},
		{
			name:     "Two elements same",
			nums:     []int{1, 1},
			expected: 2,
			expectedNums: []int{1, 1},
		},
		{
			name:     "Two elements different",
			nums:     []int{1, 2},
			expected: 2,
			expectedNums: []int{1, 2},
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: 0,
			expectedNums: []int{},
		},
		{
			name:     "Mixed with negative numbers",
			nums:     []int{-5, -5, -5, -3, -3, -3, -1, 0, 0, 2, 2, 2},
			expected: 8,
			expectedNums: []int{-5, -5, -3, -3, -1, 0, 0, 2, 2},
		},
		{
			name:     "Large duplicates block",
			nums:     []int{1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3},
			expected: 5,
			expectedNums: []int{1, 1, 2, 2, 3},
		},
		{
			name:     "Alternating duplicates",
			nums:     []int{1, 1, 1, 2, 2, 2, 3, 3, 3},
			expected: 6,
			expectedNums: []int{1, 1, 2, 2, 3, 3},
		},
		{
			name:     "Three different duplicates patterns",
			nums:     []int{1, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 3},
			expected: 6,
			expectedNums: []int{1, 1, 2, 2, 3, 3},
		},
		{
			name:     "Zeros and ones",
			nums:     []int{0, 0, 0, 0, 1, 1, 1, 1},
			expected: 4,
			expectedNums: []int{0, 0, 1, 1},
		},
		{
			name:     "Descending then ascending (should not happen with sorted input)",
			nums:     []int{5, 4, 3, 2, 1},
			expected: 5,
			expectedNums: []int{5, 4, 3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of nums since we'll modify it
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)
			
			result := RemoveDuplicatesII(numsCopy)
			
			assert.Equal(t, tt.expected, result,
				"RemoveDuplicatesII(%v) returned %d, expected %d",
				tt.nums, result, tt.expected)
			
			// Check first 'result' elements match expectedNums
			for i := 0; i < tt.expected; i++ {
				if i < len(tt.expectedNums) {
					assert.Equal(t, tt.expectedNums[i], numsCopy[i],
						"After RemoveDuplicatesII, nums[%d] = %d, expected %d",
						i, numsCopy[i], tt.expectedNums[i])
				}
			}
		})
	}
}

func TestRemoveDuplicatesIISimple(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 1, 1, 2, 2, 3},
			expected: 5,
		},
		{
			name:     "All same",
			nums:     []int{1, 1, 1, 1},
			expected: 2,
		},
		{
			name:     "No duplicates",
			nums:     []int{1, 2, 3},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)
			
			result := RemoveDuplicatesIISimple(numsCopy)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRemoveDuplicatesIIBruteForce(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Simple case",
			nums:     []int{1, 1, 2, 2, 2, 3},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)
			
			result := RemoveDuplicatesIIBruteForce(numsCopy)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRemoveDuplicatesIIGeneric(t *testing.T) {
	tests := []struct {
		name        string
		nums        []int
		maxDuplicates int
		expected    int
	}{
		{
			name:        "Allow 1 duplicate (standard remove duplicates)",
			nums:        []int{1, 1, 1, 2, 2, 3},
			maxDuplicates: 1,
			expected:    3,
		},
		{
			name:        "Allow 2 duplicates (problem requirement)",
			nums:        []int{1, 1, 1, 2, 2, 3},
			maxDuplicates: 2,
			expected:    5,
		},
		{
			name:        "Allow 3 duplicates",
			nums:        []int{1, 1, 1, 1, 2, 2, 2, 3},
			maxDuplicates: 3,
			expected:    7,
		},
		{
			name:        "Allow 0 duplicates (remove all duplicates)",
			nums:        []int{1, 1, 2, 2, 3, 3},
			maxDuplicates: 0,
			expected:    3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)
			
			result := RemoveDuplicatesIIGeneric(numsCopy, tt.maxDuplicates)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRemoveDuplicatesIIWithMap(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Simple case",
			nums:     []int{1, 1, 2, 2, 2, 3},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)
			
			result := RemoveDuplicatesIIWithMap(numsCopy)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRemoveDuplicatesIIEarlyExit(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Large duplicate block",
			nums:     []int{1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 3},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)
			
			result := RemoveDuplicatesIIEarlyExit(numsCopy)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestRemoveDuplicatesII_Consistency(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
	}{
		{
			name: "Standard case",
			nums: []int{1, 1, 1, 2, 2, 3},
		},
		{
			name: "All duplicates",
			nums: []int{2, 2, 2, 2, 2},
		},
		{
			name: "No duplicates",
			nums: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Mixed",
			nums: []int{0, 0, 0, 1, 1, 1, 1, 2, 2, 3, 3, 3, 3, 3},
		},
		{
			name: "Single element",
			nums: []int{42},
		},
		{
			name: "Empty",
			nums: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test all implementations
			implementations := []struct {
				name string
				fn   func([]int) int
			}{
				{"RemoveDuplicatesII", RemoveDuplicatesII},
				{"RemoveDuplicatesIISimple", RemoveDuplicatesIISimple},
				{"RemoveDuplicatesIIBruteForce", RemoveDuplicatesIIBruteForce},
				{"RemoveDuplicatesIIWithMap", RemoveDuplicatesIIWithMap},
				{"RemoveDuplicatesIIEarlyExit", RemoveDuplicatesIIEarlyExit},
			}

			// Also test generic version with maxDuplicates=2
			genericFn := func(nums []int) int {
				return RemoveDuplicatesIIGeneric(nums, 2)
			}
			implementations = append(implementations, struct {
				name string
				fn   func([]int) int
			}{"RemoveDuplicatesIIGeneric(2)", genericFn})

			results := make([]int, len(implementations))
			
			for i, impl := range implementations {
				numsCopy := make([]int, len(tc.nums))
				copy(numsCopy, tc.nums)
				results[i] = impl.fn(numsCopy)
			}

			// All implementations should return the same result
			for i := 1; i < len(results); i++ {
				assert.Equal(t, results[0], results[i],
					"%s and %s should return same result for %v",
					implementations[0].name, implementations[i].name, tc.nums)
			}
		})
	}
}

func TestRemoveDuplicatesII_PropertyBased(t *testing.T) {
	t.Run("Result length <= original length", func(t *testing.T) {
		// Generate random test cases
		testCases := [][]int{
			{1, 1, 1, 2, 2, 3},
			{0, 0, 1, 1, 1, 1, 2, 3, 3},
			{1, 2, 3, 4, 5},
			{1, 1, 1, 1, 1},
			{},
			{1},
			{1, 1},
		}

		for _, nums := range testCases {
			numsCopy := make([]int, len(nums))
			copy(numsCopy, nums)
			
			result := RemoveDuplicatesII(numsCopy)
			assert.True(t, result <= len(nums),
				"Result length %d should be <= original length %d for %v",
				result, len(nums), nums)
		}
	})

	t.Run("First k elements are valid (at most 2 duplicates)", func(t *testing.T) {
		testCases := [][]int{
			{1, 1, 1, 2, 2, 3},
			{0, 0, 1, 1, 1, 1, 2, 3, 3},
			{1, 1, 1, 1, 1},
		}

		for _, nums := range testCases {
			numsCopy := make([]int, len(nums))
			copy(numsCopy, nums)
			
			k := RemoveDuplicatesII(numsCopy)
			
			// Check that first k elements have at most 2 duplicates
			for i := 0; i < k; i++ {
				count := 1
				// Count occurrences of numsCopy[i] in first k elements
				for j := 0; j < k; j++ {
					if j != i && numsCopy[j] == numsCopy[i] {
						count++
					}
				}
				assert.True(t, count <= 2,
					"Element %d appears %d times in first %d elements for %v",
					numsCopy[i], count, k, nums)
			}
		}
	})

	t.Run("Result preserves order", func(t *testing.T) {
		testCases := [][]int{
			{1, 1, 1, 2, 2, 3},
			{0, 0, 1, 1, 1, 1, 2, 3, 3},
		}

		for _, nums := range testCases {
			numsCopy := make([]int, len(nums))
			copy(numsCopy, nums)
			
			k := RemoveDuplicatesII(numsCopy)
			
			// Check that first k elements are in non-decreasing order
			for i := 1; i < k; i++ {
				assert.True(t, numsCopy[i] >= numsCopy[i-1],
					"Elements not in order: nums[%d]=%d, nums[%d]=%d for %v",
					i-1, numsCopy[i-1], i, numsCopy[i], nums)
			}
		}
	})
}

func BenchmarkRemoveDuplicatesII(b *testing.B) {
	// Create test cases of different sizes
	testCases := []struct {
		name string
		nums []int
	}{
		{
			name: "Small (10 elements)",
			nums: []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4},
		},
		{
			name: "Medium (100 elements)",
			nums: func() []int {
				nums := make([]int, 100)
				for i := 0; i < 100; i++ {
					nums[i] = i / 3 // Creates duplicates
				}
				return nums
			}(),
		},
		{
			name: "Large (1000 elements)",
			nums: func() []int {
				nums := make([]int, 1000)
				for i := 0; i < 1000; i++ {
					nums[i] = i / 4 // Creates duplicates
				}
				return nums
			}(),
		},
		{
			name: "Very large (10000 elements)",
			nums: func() []int {
				nums := make([]int, 10000)
				for i := 0; i < 10000; i++ {
					nums[i] = i / 5 // Creates duplicates
				}
				return nums
			}(),
		},
		{
			name: "All same elements",
			nums: func() []int {
				nums := make([]int, 1000)
				for i := range nums {
					nums[i] = 42
				}
				return nums
			}(),
		},
		{
			name: "No duplicates",
			nums: func() []int {
				nums := make([]int, 1000)
				for i := range nums {
					nums[i] = i
				}
				return nums
			}(),
		},
	}

	implementations := []struct {
		name string
		fn   func([]int) int
	}{
		{"Standard", RemoveDuplicatesII},
		{"Simple", RemoveDuplicatesIISimple},
		{"BruteForce", RemoveDuplicatesIIBruteForce},
		{"WithMap", RemoveDuplicatesIIWithMap},
		{"EarlyExit", RemoveDuplicatesIIEarlyExit},
		{"Generic(2)", func(nums []int) int {
			return RemoveDuplicatesIIGeneric(nums, 2)
		}},
	}

	for _, tc := range testCases {
		for _, impl := range implementations {
			b.Run(tc.name+"_"+impl.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					// Create a fresh copy for each iteration
					numsCopy := make([]int, len(tc.nums))
					copy(numsCopy, tc.nums)
					impl.fn(numsCopy)
				}
			})
		}
	}
}