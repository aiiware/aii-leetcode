package linkedlists

import (
	"reflect"
	"testing"
    "leetcode/utils"
)

func TestInsertionSortList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
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
		{
			name:     "Already sorted ascending",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Already sorted descending",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Example 1 from LeetCode",
			input:    []int{4, 2, 1, 3},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Example 2 from LeetCode",
			input:    []int{-1, 5, 3, 4, 0},
			expected: []int{-1, 0, 3, 4, 5},
		},
		{
			name:     "Random order",
			input:    []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
			expected: []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9},
		},
		{
			name:     "All same values",
			input:    []int{5, 5, 5, 5, 5},
			expected: []int{5, 5, 5, 5, 5},
		},
		{
			name:     "Negative values",
			input:    []int{-5, -1, -3, -2, -4},
			expected: []int{-5, -4, -3, -2, -1},
		},
		{
			name:     "Mixed positive and negative",
			input:    []int{3, -2, 1, -5, 4, 0, -1, 2},
			expected: []int{-5, -2, -1, 0, 1, 2, 3, 4},
		},
		{
			name:     "Large values",
			input:    []int{10000, -10000, 5000, -5000, 0},
			expected: []int{-10000, -5000, 0, 5000, 10000},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test all implementations
			implementations := []struct {
				name string
				fn   func(*utils.ListNode) *utils.ListNode
			}{
				{"InsertionSortList", InsertionSortList},
				{"InsertionSortListOptimized", InsertionSortListOptimized},
				{"InsertionSortListWithArray", InsertionSortListWithArray},
			}

			for _, impl := range implementations {
				t.Run(impl.name, func(t *testing.T) {
					// Create a fresh list for each test
					head := utils.NewListFromSlice(tt.input)
					
					// Apply the sort function
					sorted := impl.fn(head)
					
					// Convert result to slice
					result := sorted.ToSlice()
					
					// Compare with expected
					if !reflect.DeepEqual(result, tt.expected) {
						t.Errorf("%s() = %v, expected %v", impl.name, result, tt.expected)
					}
				})
			}
		})
	}
}

func BenchmarkInsertionSortList(b *testing.B) {
	// Test cases for benchmarking
	testCases := []struct {
		name   string
		values []int
	}{
		{"Small sorted", []int{1, 2, 3, 4, 5}},
		{"Small reverse", []int{5, 4, 3, 2, 1}},
		{"Small random", []int{3, 1, 4, 2, 5}},
		{"Medium sorted", generateSortedArray(100)},
		{"Medium reverse", generateReverseArray(100)},
		{"Medium random", generateRandomArray(100)},
		{"Large sorted", generateSortedArray(1000)},
		{"Large reverse", generateReverseArray(1000)},
		{"Large random", generateRandomArray(1000)},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			// Create list once
			head := utils.NewListFromSlice(tc.values)
			
			b.Run("Standard", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					// Clone the list for each iteration
					cloned := cloneList(head)
					InsertionSortList(cloned)
				}
			})
			
			b.Run("Optimized", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					cloned := cloneList(head)
					InsertionSortListOptimized(cloned)
				}
			})
			
			b.Run("ArrayBased", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					cloned := cloneList(head)
					InsertionSortListWithArray(cloned)
				}
			})
		})
	}
}

func TestInsertionSortListEdgeCases(t *testing.T) {
	t.Run("Nil list", func(t *testing.T) {
		result := InsertionSortList(nil)
		if result != nil {
			t.Errorf("InsertionSortList(nil) = %v, expected nil", result)
		}

		result = InsertionSortListOptimized(nil)
		if result != nil {
			t.Errorf("InsertionSortListOptimized(nil) = %v, expected nil", result)
		}

		result = InsertionSortListWithArray(nil)
		if result != nil {
			t.Errorf("InsertionSortListWithArray(nil) = %v, expected nil", result)
		}
	})

	t.Run("Verify list is properly terminated", func(t *testing.T) {
		head := utils.NewListFromSlice([]int{4, 2, 1, 3})
		sorted := InsertionSortList(head)
		
		// Check for cycles
		if HasCycle(sorted) {
			t.Error("InsertionSortList created a cycle in the list")
		}
		
		// Verify length
		result := sorted.ToSlice()
		if len(result) != 4 {
			t.Errorf("Expected 4 nodes, got %d", len(result))
		}
	})

	t.Run("Stability test (not required but good to check)", func(t *testing.T) {
		// Create list with duplicate values but different "identities"
		// In Go, we can't easily track object identity, but we can verify
		// that the algorithm works correctly with duplicates
		head := utils.NewListFromSlice([]int{3, 1, 2, 1, 3})
		sorted := InsertionSortList(head)
		
		result := sorted.ToSlice()
		expected := []int{1, 1, 2, 3, 3}
		
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("InsertionSortList with duplicates = %v, expected %v", result, expected)
		}
	})

	t.Run("Large list with insertion at beginning", func(t *testing.T) {
		// Create list where smallest element is at the end
		values := make([]int, 100)
		for i := 0; i < 100; i++ {
			values[i] = i + 1
		}
		// Move smallest to end
		values[0], values[99] = values[99], values[0]
		
		head := utils.NewListFromSlice(values)
		sorted := InsertionSortList(head)
		
		result := sorted.ToSlice()
		
		// Verify sorted
		for i := 0; i < len(result)-1; i++ {
			if result[i] > result[i+1] {
				t.Errorf("List not sorted at position %d: %d > %d", i, result[i], result[i+1])
			}
		}
	})
}

// Helper functions for benchmarking
func generateSortedArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}

func generateReverseArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = n - i - 1
	}
	return arr
}

func generateRandomArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = (i * 73) % n // Simple pseudo-random
	}
	return arr
}

func cloneList(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return nil
	}
	
	// Create new nodes
	dummy := &utils.ListNode{}
	current := dummy
	original := head
	
	for original != nil {
		current.Next = &utils.ListNode{Val: original.Val}
		current = current.Next
		original = original.Next
	}
	
	return dummy.Next
}