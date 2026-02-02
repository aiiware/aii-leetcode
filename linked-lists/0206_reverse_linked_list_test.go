package linkedlists

import (
	"testing"
	"leetcode/utils"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Example 1",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "Example 2",
			input:    []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "Example 3 - empty list",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single node",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "Large list",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		},
		{
			name:     "List with negative numbers",
			input:    []int{-5, -4, -3, -2, -1},
			expected: []int{-1, -2, -3, -4, -5},
		},
		{
			name:     "List with zeros",
			input:    []int{0, 0, 0, 1, 2, 3},
			expected: []int{3, 2, 1, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := utils.NewListFromSlice(tt.input)
			expected := utils.NewListFromSlice(tt.expected)

			result := ReverseList(input)

			if !result.Equal(expected) {
				t.Errorf("ReverseList(%v) = %v, expected %v",
					tt.input, result.ToSlice(), tt.expected)
			}
		})
	}
}

func TestReverseListRecursive(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Example 1",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "Example 2",
			input:    []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "Empty list",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Single node",
			input:    []int{1},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := utils.NewListFromSlice(tt.input)
			expected := utils.NewListFromSlice(tt.expected)

			result := ReverseListRecursive(input)

			if !result.Equal(expected) {
				t.Errorf("ReverseListRecursive(%v) = %v, expected %v",
					tt.input, result.ToSlice(), tt.expected)
			}
		})
	}
}

func TestReverseListEdgeCases(t *testing.T) {
	// Test nil input
	t.Run("Nil input", func(t *testing.T) {
		result := ReverseList(nil)
		if result != nil {
			t.Errorf("ReverseList(nil) = %v, expected nil", result)
		}
	})

	// Test that function works correctly (original list will be modified in-place)
	t.Run("In-place reversal works correctly", func(t *testing.T) {
		original := utils.NewListFromSlice([]int{1, 2, 3})
		
		result := ReverseList(original)
		
		// Check that result is reversed
		expected := utils.NewListFromSlice([]int{3, 2, 1})
		if !result.Equal(expected) {
			t.Errorf("ReverseList([1,2,3]) = %v, expected [3,2,1]", result.ToSlice())
		}
		
		// Note: original is now modified and points to the last node (1) with Next = nil
		// This is expected behavior for in-place reversal
	})
}

func BenchmarkReverseList(b *testing.B) {
	// Create a large list for benchmarking
	vals := make([]int, 10000)
	for i := range vals {
		vals[i] = i
	}
	input := utils.NewListFromSlice(vals)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a fresh copy for each iteration since ReverseList modifies in-place
		ReverseList(utils.CopyList(input))
	}
}

func BenchmarkReverseListRecursive(b *testing.B) {
	// Create a smaller list for recursive benchmark (to avoid stack overflow)
	vals := make([]int, 1000)
	for i := range vals {
		vals[i] = i
	}
	input := utils.NewListFromSlice(vals)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a fresh copy for each iteration
		ReverseListRecursive(utils.CopyList(input))
	}
}

func TestReverseListBothMethods(t *testing.T) {
	// Test that both methods produce the same result
	tests := []struct {
		name  string
		input []int
	}{
		{"Empty", []int{}},
		{"Single", []int{1}},
		{"Two nodes", []int{1, 2}},
		{"Five nodes", []int{1, 2, 3, 4, 5}},
		{"Ten nodes", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input1 := utils.NewListFromSlice(tt.input)
			input2 := utils.NewListFromSlice(tt.input)
			
			iterativeResult := ReverseList(input1)
			recursiveResult := ReverseListRecursive(input2)
			
			if !iterativeResult.Equal(recursiveResult) {
				t.Errorf("Methods differ: iterative=%v, recursive=%v",
					iterativeResult.ToSlice(), recursiveResult.ToSlice())
			}
		})
	}
}