package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
    "leetcode/utils"
)

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Example 1",
			input:    []int{1, 1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "Example 2",
			input:    []int{1, 1, 2, 3, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "No duplicates",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "All duplicates",
			input:    []int{1, 1, 1, 1, 1},
			expected: []int{1},
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
			expected: []int{1},
		},
		{
			name:     "Two different elements",
			input:    []int{1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "Duplicates at beginning",
			input:    []int{1, 1, 1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Duplicates at end",
			input:    []int{1, 2, 3, 4, 5, 5, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Duplicates in middle",
			input:    []int{1, 2, 2, 2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Multiple duplicate groups",
			input:    []int{1, 1, 2, 2, 3, 3, 4, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Alternating duplicates",
			input:    []int{1, 2, 2, 3, 3, 4, 5, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Large duplicates block",
			input:    []int{1, 1, 1, 1, 2, 2, 2, 3, 3, 4, 5, 5, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Negative numbers",
			input:    []int{-5, -5, -3, -3, -1, 0, 0, 2, 2},
			expected: []int{-5, -3, -1, 0, 2},
		},
		{
			name:     "Mixed positive and negative",
			input:    []int{-2, -1, -1, 0, 0, 1, 1, 2, 2},
			expected: []int{-2, -1, 0, 1, 2},
		},
		{
			name:     "Consecutive unique then duplicates",
			input:    []int{1, 2, 3, 4, 4, 5, 6, 6, 7},
			expected: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:     "Only one unique element in middle",
			input:    []int{1, 1, 2, 3, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Complex case 1",
			input:    []int{0, 1, 1, 1, 2, 2, 3, 3, 3, 4, 5, 5},
			expected: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name:     "Complex case 2",
			input:    []int{1, 1, 2, 3, 3, 4, 4, 5, 6, 6, 7, 8, 8, 9},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "All same except one at end",
			input:    []int{1, 1, 1, 1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "All same except one at beginning",
			input:    []int{1, 2, 2, 2, 2},
			expected: []int{1, 2},
		},
		{
			name:     "Single duplicate pair",
			input:    []int{1, 1, 2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Triple duplicates",
			input:    []int{1, 1, 1, 2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Descending order (should not happen with sorted input)",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "Random order with duplicates",
			input:    []int{1, 2, 2, 3, 3, 3, 4, 5, 5, 5, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := utils.NewListFromSlice(tt.input)
			result := DeleteDuplicates(head)
			resultSlice := result.ToSlice()
			
			assert.Equal(t, tt.expected, resultSlice,
				"DeleteDuplicates(%v) = %v, expected %v",
				tt.input, resultSlice, tt.expected)
		})
	}
}

func TestDeleteDuplicatesRecursive(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Recursive test 1",
			input:    []int{1, 1, 2, 3, 3},
			expected: []int{1, 2, 3},
		},
		{
			name:     "All duplicates",
			input:    []int{1, 1, 1, 1},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := utils.NewListFromSlice(tt.input)
			result := DeleteDuplicatesRecursive(head)
			resultSlice := result.ToSlice()
			
			assert.Equal(t, tt.expected, resultSlice)
		})
	}
}

func TestDeleteDuplicatesTwoPointers(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Two pointers test",
			input:    []int{1, 1, 1, 2, 3},
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := utils.NewListFromSlice(tt.input)
			result := DeleteDuplicatesTwoPointers(head)
			resultSlice := result.ToSlice()
			
			assert.Equal(t, tt.expected, resultSlice)
		})
	}
}

func TestDeleteDuplicatesWithDummy(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Dummy test",
			input:    []int{1, 2, 2, 3, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := utils.NewListFromSlice(tt.input)
			result := DeleteDuplicatesWithDummy(head)
			resultSlice := result.ToSlice()
			
			assert.Equal(t, tt.expected, resultSlice)
		})
	}
}

func TestDeleteDuplicatesEarlyExit(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Early exit test",
			input:    []int{1, 1, 2, 3, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Single node",
			input:    []int{1},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := utils.NewListFromSlice(tt.input)
			result := DeleteDuplicatesEarlyExit(head)
			resultSlice := result.ToSlice()
			
			assert.Equal(t, tt.expected, resultSlice)
		})
	}
}

func TestDeleteDuplicatesStack(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Stack test",
			input:    []int{1, 1, 2, 3, 3},
			expected: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := utils.NewListFromSlice(tt.input)
			result := DeleteDuplicatesStack(head)
			resultSlice := result.ToSlice()
			
			assert.Equal(t, tt.expected, resultSlice)
		})
	}
}

func TestDeleteDuplicatesGeneric(t *testing.T) {
	tests := []struct {
		name        string
		input       []int
		maxDuplicates int
		expected    []int
	}{
		{
			name:        "Allow 1 duplicate (standard remove duplicates)",
			input:       []int{1, 1, 1, 2, 2, 3},
			maxDuplicates: 1,
			expected:    []int{1, 2, 3},
		},
		{
			name:        "Allow 2 duplicates",
			input:       []int{1, 1, 1, 2, 2, 3},
			maxDuplicates: 2,
			expected:    []int{1, 1, 2, 2, 3},
		},
		{
			name:        "Allow 3 duplicates",
			input:       []int{1, 1, 1, 1, 2, 2, 2, 3},
			maxDuplicates: 3,
			expected:    []int{1, 1, 1, 2, 2, 2, 3},
		},
		{
			name:        "Allow 0 duplicates (remove all)",
			input:       []int{1, 1, 2, 2, 3, 3},
			maxDuplicates: 0,
			expected:    []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := utils.NewListFromSlice(tt.input)
			result := DeleteDuplicatesGeneric(head, tt.maxDuplicates)
			resultSlice := result.ToSlice()
			
			assert.Equal(t, tt.expected, resultSlice)
		})
	}
}

func TestDeleteDuplicates_Consistency(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
	}{
		{
			name:  "Standard case",
			input: []int{1, 1, 2, 3, 3},
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
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test all implementations
			implementations := []struct {
				name string
				fn   func(*utils.ListNode) *utils.ListNode
			}{
				{"DeleteDuplicates", DeleteDuplicates},
				{"DeleteDuplicatesRecursive", DeleteDuplicatesRecursive},
				{"DeleteDuplicatesTwoPointers", DeleteDuplicatesTwoPointers},
				{"DeleteDuplicatesWithDummy", DeleteDuplicatesWithDummy},
				{"DeleteDuplicatesEarlyExit", DeleteDuplicatesEarlyExit},
				{"DeleteDuplicatesStack", DeleteDuplicatesStack},
			}

			// Also test generic version with maxDuplicates=1
			genericFn := func(head *utils.ListNode) *utils.ListNode {
				return DeleteDuplicatesGeneric(head, 1)
			}
			implementations = append(implementations, struct {
				name string
				fn   func(*utils.ListNode) *utils.ListNode
			}{"DeleteDuplicatesGeneric(1)", genericFn})

			results := make([][]int, len(implementations))
			
			for i, impl := range implementations {
				head := utils.NewListFromSlice(tc.input)
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

func TestDeleteDuplicates_PropertyBased(t *testing.T) {
	t.Run("Result contains no consecutive duplicates", func(t *testing.T) {
		testCases := [][]int{
			{1, 1, 2, 3, 3},
			{1, 1, 1, 2, 3},
			{1, 1, 1, 1, 1},
			{1, 2, 3, 4, 5},
			{1, 1, 2, 2, 3, 3},
		}

		for _, input := range testCases {
			head := utils.NewListFromSlice(input)
			result := DeleteDuplicates(head)
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
			{1, 1, 2, 3, 3},
			{1, 1, 1, 1},
		}

		for _, input := range testCases {
			head := utils.NewListFromSlice(input)
			result := DeleteDuplicates(head)
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
			{1, 1, 2, 3, 3},
			{1, 1, 1, 2, 3},
		}

		for _, input := range testCases {
			head := utils.NewListFromSlice(input)
			result := DeleteDuplicates(head)
			resultSlice := result.ToSlice()
			
			// Create a set of original input
			originalSet := make(map[int]bool)
			for _, val := range input {
				originalSet[val] = true
			}
			
			// Check that each element in result was in original
			for _, val := range resultSlice {
				assert.True(t, originalSet[val],
					"Element %d in result %v was not in original %v",
					val, resultSlice, input)
			}
		}
	})

	t.Run("At least one copy of each unique element is preserved", func(t *testing.T) {
		testCases := [][]int{
			{1, 1, 2, 3, 3},
			{1, 1, 1, 2, 3},
		}

		for _, input := range testCases {
			head := utils.NewListFromSlice(input)
			result := DeleteDuplicates(head)
			resultSlice := result.ToSlice()
			
			// Create a set of unique elements in original
			originalUnique := make(map[int]bool)
			for _, val := range input {
				originalUnique[val] = true
			}
			
			// Create a set of elements in result
			resultSet := make(map[int]bool)
			for _, val := range resultSlice {
				resultSet[val] = true
			}
			
			// Check that all unique elements from original are in result
			for val := range originalUnique {
				assert.True(t, resultSet[val],
					"Unique element %d from original %v is missing in result %v",
					val, input, resultSlice)
			}
		}
	})

	t.Run("Idempotent", func(t *testing.T) {
		// Applying the function twice should give the same result
		testCases := [][]int{
			{1, 1, 2, 3, 3},
			{1, 1, 1, 2, 3},
			{1, 2, 3, 4, 5},
		}

		for _, input := range testCases {
			head := utils.NewListFromSlice(input)
			result1 := DeleteDuplicates(head)
			result1Slice := result1.ToSlice()
			
			// Apply function again to result
			head2 := utils.NewListFromSlice(result1Slice)
			result2 := DeleteDuplicates(head2)
			result2Slice := result2.ToSlice()
			
			assert.Equal(t, result1Slice, result2Slice,
				"DeleteDuplicates should be idempotent for input %v",
				input)
		}
	})

	t.Run("Length property", func(t *testing.T) {
		// Result length should be <= original length
		// Result length should be >= number of unique elements
		testCases := [][]int{
			{1, 1, 2, 3, 3},
			{1, 1, 1, 2, 3},
			{1, 2, 3, 4, 5},
		}

		for _, input := range testCases {
			head := utils.NewListFromSlice(input)
			result := DeleteDuplicates(head)
			resultSlice := result.ToSlice()
			
			// Count unique elements in original
			uniqueCount := 0
			seen := make(map[int]bool)
			for _, val := range input {
				if !seen[val] {
					seen[val] = true
					uniqueCount++
				}
			}
			
			assert.True(t, len(resultSlice) <= len(input),
				"Result length %d should be <= original length %d for input %v",
				len(resultSlice), len(input), input)
			
			assert.True(t, len(resultSlice) >= uniqueCount,
				"Result length %d should be >= unique count %d for input %v",
				len(resultSlice), uniqueCount, input)
			
			// For this problem specifically, result length should equal unique count
			assert.Equal(t, uniqueCount, len(resultSlice),
				"Result length %d should equal unique count %d for input %v",
				len(resultSlice), uniqueCount, input)
		}
	})
}

func BenchmarkDeleteDuplicates(b *testing.B) {
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
		fn   func(*utils.ListNode) *utils.ListNode
	}{
		{"Standard", DeleteDuplicates},
		{"Recursive", DeleteDuplicatesRecursive},
		{"TwoPointers", DeleteDuplicatesTwoPointers},
		{"WithDummy", DeleteDuplicatesWithDummy},
		{"EarlyExit", DeleteDuplicatesEarlyExit},
		{"Stack", DeleteDuplicatesStack},
		{"Generic(1)", func(head *utils.ListNode) *utils.ListNode {
			return DeleteDuplicatesGeneric(head, 1)
		}},
	}

	for _, tc := range testCases {
		for _, impl := range implementations {
			b.Run(tc.name+"_"+impl.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					head := utils.NewListFromSlice(tc.input)
					impl.fn(head)
				}
			})
		}
	}
}