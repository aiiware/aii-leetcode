package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteDuplicatesII(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Example 1",
			input:    []int{1, 2, 3, 3, 4, 4, 5},
			expected: []int{1, 2, 5},
		},
		{
			name:     "Example 2",
			input:    []int{1, 1, 1, 2, 3},
			expected: []int{2, 3},
		},
		{
			name:     "No duplicates",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "All duplicates",
			input:    []int{1, 1, 1, 1, 1},
			expected: []int{},
		},
		{
			name:     "Empty list",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single element",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "Two same elements",
			input:    []int{1, 1},
			expected: []int{},
		},
		{
			name:     "Two different elements",
			input:    []int{1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "Duplicates at beginning",
			input:    []int{1, 1, 1, 2, 3, 4, 5},
			expected: []int{2, 3, 4, 5},
		},
		{
			name:     "Duplicates at end",
			input:    []int{1, 2, 3, 4, 5, 5, 5},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Duplicates in middle",
			input:    []int{1, 2, 2, 2, 3, 4},
			expected: []int{1, 3, 4},
		},
		{
			name:     "Multiple duplicate groups",
			input:    []int{1, 1, 2, 2, 3, 3, 4, 4},
			expected: []int{},
		},
		{
			name:     "Alternating duplicates",
			input:    []int{1, 2, 2, 3, 3, 4, 5, 5},
			expected: []int{1, 4},
		},
		{
			name:     "Large duplicates block",
			input:    []int{1, 1, 1, 1, 2, 2, 2, 3, 3, 4, 5, 5, 5},
			expected: []int{4},
		},
		{
			name:     "Negative numbers",
			input:    []int{-5, -5, -3, -3, -1, 0, 0, 2, 2},
			expected: []int{-1},
		},
		{
			name:     "Mixed positive and negative",
			input:    []int{-2, -1, -1, 0, 0, 1, 1, 2, 2},
			expected: []int{-2},
		},
		{
			name:     "Consecutive unique then duplicates",
			input:    []int{1, 2, 3, 4, 4, 5, 6, 6, 7},
			expected: []int{1, 2, 3, 5, 7},
		},
		{
			name:     "Only one unique element in middle",
			input:    []int{1, 1, 2, 3, 3},
			expected: []int{2},
		},
		{
			name:     "Complex case 1",
			input:    []int{0, 1, 1, 1, 2, 2, 3, 3, 3, 4, 5, 5},
			expected: []int{0, 4},
		},
		{
			name:     "Complex case 2",
			input:    []int{1, 1, 2, 3, 3, 4, 4, 5, 6, 6, 7, 8, 8, 9},
			expected: []int{2, 5, 7, 9},
		},
		{
			name:     "All same except one at end",
			input:    []int{1, 1, 1, 1, 2},
			expected: []int{2},
		},
		{
			name:     "All same except one at beginning",
			input:    []int{1, 2, 2, 2, 2},
			expected: []int{1},
		},
		{
			name:     "Single duplicate pair",
			input:    []int{1, 1, 2, 3, 4},
			expected: []int{2, 3, 4},
		},
		{
			name:     "Triple duplicates",
			input:    []int{1, 1, 1, 2, 3, 4},
			expected: []int{2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := NewListFromSlice(tt.input)
			result := DeleteDuplicatesII(head)
			resultSlice := result.ToSlice()
			
			assert.Equal(t, tt.expected, resultSlice,
				"DeleteDuplicatesII(%v) = %v, expected %v",
				tt.input, resultSlice, tt.expected)
		})
	}
}

func TestDeleteDuplicatesIIRecursive(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Recursive test 1",
			input:    []int{1, 2, 3, 3, 4, 4, 5},
			expected: []int{1, 2, 5},
		},
		{
			name:     "All duplicates",
			input:    []int{1, 1, 1, 1},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := NewListFromSlice(tt.input)
			result := DeleteDuplicatesIIRecursive(head)
			resultSlice := result.ToSlice()
			
			assert.Equal(t, tt.expected, resultSlice)
		})
	}
}

func TestDeleteDuplicatesIITwoPointers(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Two pointers test",
			input:    []int{1, 1, 1, 2, 3},
			expected: []int{2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := NewListFromSlice(tt.input)
			result := DeleteDuplicatesIITwoPointers(head)
			resultSlice := result.ToSlice()
			
			assert.Equal(t, tt.expected, resultSlice)
		})
	}
}

func TestDeleteDuplicatesIIWithCounter(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Counter test",
			input:    []int{1, 2, 2, 3, 3, 4},
			expected: []int{1, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := NewListFromSlice(tt.input)
			result := DeleteDuplicatesIIWithCounter(head)
			resultSlice := result.ToSlice()
			
			assert.Equal(t, tt.expected, resultSlice)
		})
	}
}

func TestDeleteDuplicatesIIEarlyExit(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Early exit test",
			input:    []int{1, 1, 2, 3, 3, 4, 5},
			expected: []int{2, 4, 5},
		},
		{
			name:     "Single node",
			input:    []int{1},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := NewListFromSlice(tt.input)
			result := DeleteDuplicatesIIEarlyExit(head)
			resultSlice := result.ToSlice()
			
			assert.Equal(t, tt.expected, resultSlice)
		})
	}
}

func TestDeleteDuplicatesIIStack(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Stack test",
			input:    []int{1, 2, 3, 3, 4, 4, 5},
			expected: []int{1, 2, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := NewListFromSlice(tt.input)
			result := DeleteDuplicatesIIStack(head)
			resultSlice := result.ToSlice()
			
			assert.Equal(t, tt.expected, resultSlice)
		})
	}
}

func TestDeleteDuplicatesII_Consistency(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
	}{
		{
			name:  "Standard case",
			input: []int{1, 2, 3, 3, 4, 4, 5},
		},
		{
			name:  "All duplicates",
			input: []int{1, 1, 1, 1, 1},
		},
		{
			name:  "No duplicates",
			input: []int{1, 2, 3, 4, 5},
		},
		{
			name:  "Empty list",
			input: []int{},
		},
		{
			name:  "Single element",
			input: []int{42},
		},
		{
			name:  "Complex duplicates",
			input: []int{1, 1, 2, 2, 3, 3, 4, 4, 5},
		},
		{
			name:  "Only one unique",
			input: []int{1, 1, 2, 2, 2, 3, 3},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test all implementations
			implementations := []struct {
				name string
				fn   func(*ListNode) *ListNode
			}{
				{"DeleteDuplicatesII", DeleteDuplicatesII},
				{"DeleteDuplicatesIIRecursive", DeleteDuplicatesIIRecursive},
				{"DeleteDuplicatesIITwoPointers", DeleteDuplicatesIITwoPointers},
				{"DeleteDuplicatesIIWithCounter", DeleteDuplicatesIIWithCounter},
				{"DeleteDuplicatesIIEarlyExit", DeleteDuplicatesIIEarlyExit},
				{"DeleteDuplicatesIIStack", DeleteDuplicatesIIStack},
			}

			results := make([][]int, len(implementations))
			
			for i, impl := range implementations {
				head := NewListFromSlice(tc.input)
				result := impl.fn(head)
				results[i] = result.ToSlice()
			}

			// All implementations should return the same result
			for i := 1; i < len(results); i++ {
				assert.Equal(t, results[0], results[i],
					"%s and %s should return same result for input %v",
					implementations[0].name, implementations[i].name, tc.input)
			}
		})
	}
}

func TestDeleteDuplicatesII_PropertyBased(t *testing.T) {
	t.Run("Result contains no duplicates", func(t *testing.T) {
		testCases := [][]int{
			{1, 2, 3, 3, 4, 4, 5},
			{1, 1, 1, 2, 3},
			{1, 1, 1, 1, 1},
			{1, 2, 3, 4, 5},
			{1, 1, 2, 2, 3, 3},
		}

		for _, input := range testCases {
			head := NewListFromSlice(input)
			result := DeleteDuplicatesII(head)
			resultSlice := result.ToSlice()
			
			// Check that result contains no consecutive duplicates
			for i := 1; i < len(resultSlice); i++ {
				assert.NotEqual(t, resultSlice[i], resultSlice[i-1],
					"Result %v contains consecutive duplicates at index %d",
					resultSlice, i)
			}
		}
	})

	t.Run("Result is sorted", func(t *testing.T) {
		testCases := [][]int{
			{1, 2, 3, 3, 4, 4, 5},
			{1, 1, 1, 1},
		}

		for _, input := range testCases {
			head := NewListFromSlice(input)
			result := DeleteDuplicatesII(head)
			resultSlice := result.ToSlice()
			
			// Check that result is sorted (non-decreasing)
			for i := 1; i < len(resultSlice); i++ {
				assert.True(t, resultSlice[i] >= resultSlice[i-1],
					"Result %v is not sorted at index %d",
					resultSlice, i)
			}
		}
	})

	t.Run("All elements in result were in original", func(t *testing.T) {
		testCases := [][]int{
			{1, 2, 3, 3, 4, 4, 5},
			{1, 1, 1, 2, 3},
		}

		for _, input := range testCases {
			head := NewListFromSlice(input)
			result := DeleteDuplicatesII(head)
			resultSlice := result.ToSlice()
			
			// Create a frequency map of original input
			originalMap := make(map[int]int)
			for _, val := range input {
				originalMap[val]++
			}
			
			// Check that each element in result was in original
			for _, val := range resultSlice {
				assert.True(t, originalMap[val] > 0,
					"Element %d in result %v was not in original %v",
					val, resultSlice, input)
			}
		}
	})

	t.Run("Elements with count > 1 are removed", func(t *testing.T) {
		testCases := [][]int{
			{1, 2, 3, 3, 4, 4, 5},
			{1, 1, 1, 2, 3},
		}

		for _, input := range testCases {
			head := NewListFromSlice(input)
			result := DeleteDuplicatesII(head)
			resultSlice := result.ToSlice()
			
			// Create a frequency map of original input
			originalMap := make(map[int]int)
			for _, val := range input {
				originalMap[val]++
			}
			
			// Check that no element with count > 1 is in result
			for _, val := range resultSlice {
				assert.Equal(t, 1, originalMap[val],
					"Element %d appears %d times in original but is in result %v",
					val, originalMap[val], resultSlice)
			}
		}
	})

	t.Run("Idempotent", func(t *testing.T) {
		// Applying the function twice should give the same result
		testCases := [][]int{
			{1, 2, 3, 3, 4, 4, 5},
			{1, 1, 1, 2, 3},
			{1, 2, 3, 4, 5},
		}

		for _, input := range testCases {
			head := NewListFromSlice(input)
			result1 := DeleteDuplicatesII(head)
			result1Slice := result1.ToSlice()
			
			// Apply function again to result
			head2 := NewListFromSlice(result1Slice)
			result2 := DeleteDuplicatesII(head2)
			result2Slice := result2.ToSlice()
			
			assert.Equal(t, result1Slice, result2Slice,
				"DeleteDuplicatesII should be idempotent for input %v",
				input)
		}
	})
}

func BenchmarkDeleteDuplicatesII(b *testing.B) {
	// Create test cases of different sizes
	testCases := []struct {
		name  string
		input []int
	}{
		{
			name:  "Small (10 elements)",
			input: []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5},
		},
		{
			name:  "Medium (100 elements) with few duplicates",
			input: func() []int {
				nums := make([]int, 100)
				for i := 0; i < 100; i++ {
					nums[i] = i
				}
				return nums
			}(),
		},
		{
			name:  "Medium (100 elements) with many duplicates",
			input: func() []int {
				nums := make([]int, 100)
				for i := 0; i < 100; i++ {
					nums[i] = i / 5 // Creates many duplicates
				}
				return nums
			}(),
		},
		{
			name:  "Large (1000 elements)",
			input: func() []int {
				nums := make([]int, 1000)
				for i := 0; i < 1000; i++ {
					nums[i] = i / 10 // Creates duplicates
				}
				return nums
			}(),
		},
		{
			name:  "All same elements",
			input: func() []int {
				nums := make([]int, 1000)
				for i := range nums {
					nums[i] = 7
				}
				return nums
			}(),
		},
		{
			name:  "No duplicates",
			input: func() []int {
				nums := make([]int, 1000)
				for i := range nums {
					nums[i] = i
				}
				return nums
			}(),
		},
		{
			name:  "Alternating pattern",
			input: func() []int {
				nums := make([]int, 1000)
				for i := range nums {
					if i%3 == 0 {
						nums[i] = i / 3
					} else {
						nums[i] = i/3 + 1
					}
				}
				return nums
			}(),
		},
	}

	implementations := []struct {
		name string
		fn   func(*ListNode) *ListNode
	}{
		{"Standard", DeleteDuplicatesII},
		{"Recursive", DeleteDuplicatesIIRecursive},
		{"TwoPointers", DeleteDuplicatesIITwoPointers},
		{"WithCounter", DeleteDuplicatesIIWithCounter},
		{"EarlyExit", DeleteDuplicatesIIEarlyExit},
		{"Stack", DeleteDuplicatesIIStack},
	}

	for _, tc := range testCases {
		for _, impl := range implementations {
			b.Run(tc.name+"_"+impl.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					head := NewListFromSlice(tc.input)
					impl.fn(head)
				}
			})
		}
	}
}