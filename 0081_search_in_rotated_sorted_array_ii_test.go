package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInRotatedSortedArrayII(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected bool
	}{
		{
			name:     "Example 1: target found",
			nums:     []int{2, 5, 6, 0, 0, 1, 2},
			target:   0,
			expected: true,
		},
		{
			name:     "Example 2: target not found",
			nums:     []int{2, 5, 6, 0, 0, 1, 2},
			target:   3,
			expected: false,
		},
		{
			name:     "Single element found",
			nums:     []int{1},
			target:   1,
			expected: true,
		},
		{
			name:     "Single element not found",
			nums:     []int{1},
			target:   0,
			expected: false,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			target:   1,
			expected: false,
		},
		{
			name:     "Not rotated, target found",
			nums:     []int{1, 2, 3, 4, 5},
			target:   3,
			expected: true,
		},
		{
			name:     "Not rotated, target not found",
			nums:     []int{1, 2, 3, 4, 5},
			target:   6,
			expected: false,
		},
		{
			name:     "All same elements, target found",
			nums:     []int{1, 1, 1, 1, 1},
			target:   1,
			expected: true,
		},
		{
			name:     "All same elements, target not found",
			nums:     []int{1, 1, 1, 1, 1},
			target:   2,
			expected: false,
		},
		{
			name:     "Rotated with duplicates at pivot",
			nums:     []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 13, 1, 1, 1, 1},
			target:   13,
			expected: true,
		},
		{
			name:     "Complex case with many duplicates",
			nums:     []int{4, 5, 6, 6, 7, 0, 1, 2, 4, 4},
			target:   0,
			expected: true,
		},
		{
			name:     "Complex case with many duplicates, target not found",
			nums:     []int{4, 5, 6, 6, 7, 0, 1, 2, 4, 4},
			target:   3,
			expected: false,
		},
		{
			name:     "Target at beginning",
			nums:     []int{3, 1, 2, 2, 2},
			target:   3,
			expected: true,
		},
		{
			name:     "Target at end",
			nums:     []int{2, 2, 2, 1, 3},
			target:   3,
			expected: true,
		},
		{
			name:     "Large rotation",
			nums:     []int{6, 7, 8, 9, 10, 1, 2, 3, 4, 5},
			target:   8,
			expected: true,
		},
		{
			name:     "Small rotation",
			nums:     []int{1, 3, 1, 1, 1},
			target:   3,
			expected: true,
		},
		{
			name:     "Negative numbers",
			nums:     []int{-5, -3, -3, -1, -1, 0, 2, 2},
			target:   -3,
			expected: true,
		},
		{
			name:     "Mixed positive and negative",
			nums:     []int{-1, 0, 2, 2, -5, -3},
			target:   0,
			expected: true,
		},
		{
			name:     "All duplicates except target",
			nums:     []int{1, 1, 1, 1, 2, 1, 1},
			target:   2,
			expected: true,
		},
		{
			name:     "Worst case for binary search (many duplicates)",
			nums:     []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1},
			target:   2,
			expected: true,
		},
		{
			name:     "Rotation at middle with duplicates",
			nums:     []int{2, 2, 2, 0, 2, 2},
			target:   0,
			expected: true,
		},
		{
			name:     "Target in left sorted portion",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   5,
			expected: true,
		},
		{
			name:     "Target in right sorted portion",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   1,
			expected: true,
		},
		{
			name:     "Full rotation (back to original)",
			nums:     []int{1, 2, 3, 4, 5},
			target:   3,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SearchInRotatedSortedArrayII(tt.nums, tt.target)
			assert.Equal(t, tt.expected, result,
				"SearchInRotatedSortedArrayII(%v, %d) = %v, expected %v",
				tt.nums, tt.target, result, tt.expected)
		})
	}
}

func TestSearchInRotatedSortedArrayIILinear(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected bool
	}{
		{
			name:     "Simple case",
			nums:     []int{2, 5, 6, 0, 0, 1, 2},
			target:   0,
			expected: true,
		},
		{
			name:     "Not found",
			nums:     []int{2, 5, 6, 0, 0, 1, 2},
			target:   3,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SearchInRotatedSortedArrayIILinear(tt.nums, tt.target)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSearchInRotatedSortedArrayIIFindPivot(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected bool
	}{
		{
			name:     "With duplicates",
			nums:     []int{4, 5, 6, 6, 7, 0, 1, 2, 4, 4},
			target:   0,
			expected: true,
		},
		{
			name:     "All same",
			nums:     []int{1, 1, 1, 1},
			target:   1,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SearchInRotatedSortedArrayIIFindPivot(tt.nums, tt.target)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSearchInRotatedSortedArrayIIRecursive(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected bool
	}{
		{
			name:     "Recursive test",
			nums:     []int{2, 5, 6, 0, 0, 1, 2},
			target:   1,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SearchInRotatedSortedArrayIIRecursive(tt.nums, tt.target)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSearchInRotatedSortedArrayIIEarlyExit(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected bool
	}{
		{
			name:     "Early exit first element",
			nums:     []int{5, 6, 0, 1, 2, 3, 4},
			target:   5,
			expected: true,
		},
		{
			name:     "Early exit last element",
			nums:     []int{5, 6, 0, 1, 2, 3, 4},
			target:   4,
			expected: true,
		},
		{
			name:     "Small array",
			nums:     []int{1, 2, 3},
			target:   2,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SearchInRotatedSortedArrayIIEarlyExit(tt.nums, tt.target)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSearchInRotatedSortedArrayIITwoPass(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected bool
	}{
		{
			name:     "Two pass test",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			target:   0,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SearchInRotatedSortedArrayIITwoPass(tt.nums, tt.target)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestSearchInRotatedSortedArrayII_Consistency(t *testing.T) {
	testCases := []struct {
		name   string
		nums   []int
		target int
	}{
		{
			name:   "Standard case",
			nums:   []int{2, 5, 6, 0, 0, 1, 2},
			target: 0,
		},
		{
			name:   "All duplicates",
			nums:   []int{1, 1, 1, 1, 1},
			target: 1,
		},
		{
			name:   "Not rotated",
			nums:   []int{1, 2, 3, 4, 5},
			target: 3,
		},
		{
			name:   "Complex duplicates",
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 13, 1, 1, 1, 1},
			target: 13,
		},
		{
			name:   "Empty array",
			nums:   []int{},
			target: 1,
		},
		{
			name:   "Single element",
			nums:   []int{42},
			target: 42,
		},
		{
			name:   "Target not found",
			nums:   []int{2, 5, 6, 0, 0, 1, 2},
			target: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test all implementations
			implementations := []struct {
				name string
				fn   func([]int, int) bool
			}{
				{"SearchInRotatedSortedArrayII", SearchInRotatedSortedArrayII},
				{"SearchInRotatedSortedArrayIILinear", SearchInRotatedSortedArrayIILinear},
				{"SearchInRotatedSortedArrayIIFindPivot", SearchInRotatedSortedArrayIIFindPivot},
				{"SearchInRotatedSortedArrayIIRecursive", SearchInRotatedSortedArrayIIRecursive},
				{"SearchInRotatedSortedArrayIIEarlyExit", SearchInRotatedSortedArrayIIEarlyExit},
				{"SearchInRotatedSortedArrayIITwoPass", SearchInRotatedSortedArrayIITwoPass},
			}

			results := make([]bool, len(implementations))
			
			for i, impl := range implementations {
				results[i] = impl.fn(tc.nums, tc.target)
			}

			// All implementations should return the same result
			for i := 1; i < len(results); i++ {
				assert.Equal(t, results[0], results[i],
					"%s and %s should return same result for nums=%v, target=%d",
					implementations[0].name, implementations[i].name, tc.nums, tc.target)
			}
		})
	}
}

func TestSearchInRotatedSortedArrayII_PropertyBased(t *testing.T) {
	t.Run("Linear search equivalence", func(t *testing.T) {
		// For small arrays, test that our implementation matches linear search
		testCases := [][]int{
			{2, 5, 6, 0, 0, 1, 2},
			{1, 1, 1, 1, 1},
			{1, 2, 3, 4, 5},
			{},
			{1},
			{1, 3, 1, 1, 1},
		}

		for _, nums := range testCases {
			// Test multiple targets including values that may or may not be in nums
			targets := append([]int{}, nums...)
			targets = append(targets, -100, 100, 0, 999)
			
			for _, target := range targets {
				binaryResult := SearchInRotatedSortedArrayII(nums, target)
				linearResult := SearchInRotatedSortedArrayIILinear(nums, target)
				
				assert.Equal(t, linearResult, binaryResult,
					"For nums=%v, target=%d: binary=%v, linear=%v",
					nums, target, binaryResult, linearResult)
			}
		}
	})

	t.Run("Monotonic with duplicates", func(t *testing.T) {
		// Test that if we find a target, it actually exists in the array
		testCases := []struct {
			nums   []int
			target int
		}{
			{nums: []int{2, 5, 6, 0, 0, 1, 2}, target: 0},
			{nums: []int{2, 5, 6, 0, 0, 1, 2}, target: 5},
			{nums: []int{2, 5, 6, 0, 0, 1, 2}, target: 2},
		}

		for _, tc := range testCases {
			if SearchInRotatedSortedArrayII(tc.nums, tc.target) {
				// Verify target actually exists
				found := false
				for _, num := range tc.nums {
					if num == tc.target {
						found = true
						break
					}
				}
				assert.True(t, found,
					"SearchInRotatedSortedArrayII returned true for target %d in %v, but target not found in array",
					tc.target, tc.nums)
			}
		}
	})

	t.Run("Idempotent", func(t *testing.T) {
		// Calling the function twice should return the same result
		testCases := []struct {
			nums   []int
			target int
		}{
			{nums: []int{2, 5, 6, 0, 0, 1, 2}, target: 0},
			{nums: []int{1, 1, 1, 1}, target: 2},
			{nums: []int{}, target: 1},
		}

		for _, tc := range testCases {
			result1 := SearchInRotatedSortedArrayII(tc.nums, tc.target)
			result2 := SearchInRotatedSortedArrayII(tc.nums, tc.target)
			
			assert.Equal(t, result1, result2,
				"SearchInRotatedSortedArrayII should be idempotent for nums=%v, target=%d",
				tc.nums, tc.target)
		}
	})
}

func BenchmarkSearchInRotatedSortedArrayII(b *testing.B) {
	// Create test cases of different sizes and characteristics
	testCases := []struct {
		name   string
		nums   []int
		target int
	}{
		{
			name:   "Small (10 elements)",
			nums:   []int{4, 5, 6, 7, 8, 9, 0, 1, 2, 3},
			target: 0,
		},
		{
			name:   "Medium (100 elements) with few duplicates",
			nums: func() []int {
				nums := make([]int, 100)
				for i := 0; i < 100; i++ {
					nums[i] = (i + 30) % 100 // Creates a rotation
				}
				return nums
			}(),
			target: 42,
		},
		{
			name:   "Medium (100 elements) with many duplicates",
			nums: func() []int {
				nums := make([]int, 100)
				for i := 0; i < 100; i++ {
					nums[i] = i / 10 // Creates many duplicates
				}
				// Rotate
				rotate := 40
				rotated := append(nums[rotate:], nums[:rotate]...)
				return rotated
			}(),
			target: 5,
		},
		{
			name:   "Large (1000 elements)",
			nums: func() []int {
				nums := make([]int, 1000)
				for i := 0; i < 1000; i++ {
					nums[i] = i
				}
				// Rotate
				rotate := 300
				rotated := append(nums[rotate:], nums[:rotate]...)
				return rotated
			}(),
			target: 500,
		},
		{
			name:   "All same elements",
			nums: func() []int {
				nums := make([]int, 1000)
				for i := range nums {
					nums[i] = 7
				}
				return nums
			}(),
			target: 7,
		},
		{
			name:   "Worst case (many duplicates, target at end)",
			nums: func() []int {
				nums := make([]int, 1000)
				for i := range nums {
					nums[i] = 1
				}
				nums[999] = 2
				return nums
			}(),
			target: 2,
		},
		{
			name:   "Not found case",
			nums: func() []int {
				nums := make([]int, 1000)
				for i := range nums {
					nums[i] = i * 2
				}
				rotate := 300
				rotated := append(nums[rotate:], nums[:rotate]...)
				return rotated
			}(),
			target: 999, // Odd number, won't be found
		},
	}

	implementations := []struct {
		name string
		fn   func([]int, int) bool
	}{
		{"Standard", SearchInRotatedSortedArrayII},
		{"Linear", SearchInRotatedSortedArrayIILinear},
		{"FindPivot", SearchInRotatedSortedArrayIIFindPivot},
		{"Recursive", SearchInRotatedSortedArrayIIRecursive},
		{"EarlyExit", SearchInRotatedSortedArrayIIEarlyExit},
		{"TwoPass", SearchInRotatedSortedArrayIITwoPass},
	}

	for _, tc := range testCases {
		for _, impl := range implementations {
			b.Run(tc.name+"_"+impl.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					impl.fn(tc.nums, tc.target)
				}
			})
		}
	}
}