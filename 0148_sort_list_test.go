package leetcode

import (
	"testing"
)

func TestSortList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Example 1",
			input:    []int{4, 2, 1, 3},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Example 2",
			input:    []int{-1, 5, 3, 4, 0},
			expected: []int{-1, 0, 3, 4, 5},
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
			name:     "Already sorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse sorted",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "All same elements",
			input:    []int{3, 3, 3, 3, 3},
			expected: []int{3, 3, 3, 3, 3},
		},
		{
			name:     "Large list",
			input:    []int{9, 2, 7, 4, 5, 6, 3, 8, 1, 0},
			expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "Negative numbers",
			input:    []int{-5, -1, -3, -2, -4},
			expected: []int{-5, -4, -3, -2, -1},
		},
		{
			name:     "Mixed positive and negative",
			input:    []int{-3, 5, -1, 2, 0, -2, 4, -4, 1, 3},
			expected: []int{-4, -3, -2, -1, 0, 1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create input list
			inputList := createList(tt.input)
			
			// Sort the list
			result := SortList(inputList)
			
			// Convert result to slice
			resultSlice := listToSlice(result)
			
			// Check if result matches expected
			if !slicesEqual(resultSlice, tt.expected) {
				t.Errorf("SortList(%v) = %v, expected %v", tt.input, resultSlice, tt.expected)
			}
		})
	}
}

func TestSortListBottomUp(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Example 1",
			input:    []int{4, 2, 1, 3},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Example 2",
			input:    []int{-1, 5, 3, 4, 0},
			expected: []int{-1, 0, 3, 4, 5},
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
			name:     "Already sorted",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create input list
			inputList := createList(tt.input)
			
			// Sort the list using bottom-up approach
			result := SortListBottomUp(inputList)
			
			// Convert result to slice
			resultSlice := listToSlice(result)
			
			// Check if result matches expected
			if !slicesEqual(resultSlice, tt.expected) {
				t.Errorf("SortListBottomUp(%v) = %v, expected %v", tt.input, resultSlice, tt.expected)
			}
		})
	}
}

func BenchmarkSortList(b *testing.B) {
	// Create a large list for benchmarking
	size := 10000
	input := make([]int, size)
	for i := 0; i < size; i++ {
		input[i] = size - i - 1 // Reverse sorted
	}
	
	inputList := createList(input)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a fresh copy for each iteration
		listCopy := copyList(inputList)
		SortList(listCopy)
	}
}

func BenchmarkSortListBottomUp(b *testing.B) {
	// Create a large list for benchmarking
	size := 10000
	input := make([]int, size)
	for i := 0; i < size; i++ {
		input[i] = size - i - 1 // Reverse sorted
	}
	
	inputList := createList(input)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a fresh copy for each iteration
		listCopy := copyList(inputList)
		SortListBottomUp(listCopy)
	}
}

// Helper function to create a linked list from a slice
func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	
	head := &ListNode{Val: nums[0]}
	current := head
	
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}
	
	return head
}

// Helper function to convert a linked list to a slice
func listToSlice(head *ListNode) []int {
	var result []int
	current := head
	
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	
	return result
}

// Helper function to check if two slices are equal
func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	
	return true
}

// Helper function to copy a linked list
func copyList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	
	newHead := &ListNode{Val: head.Val}
	current := newHead
	original := head.Next
	
	for original != nil {
		current.Next = &ListNode{Val: original.Val}
		current = current.Next
		original = original.Next
	}
	
	return newHead
}