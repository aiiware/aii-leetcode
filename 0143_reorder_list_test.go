package leetcode

import (
	"reflect"
	"testing"
)

func TestReorderList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Single node",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "Two nodes",
			input:    []int{1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "Three nodes",
			input:    []int{1, 2, 3},
			expected: []int{1, 3, 2},
		},
		{
			name:     "Four nodes",
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 4, 2, 3},
		},
		{
			name:     "Five nodes",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 5, 2, 4, 3},
		},
		{
			name:     "Six nodes",
			input:    []int{1, 2, 3, 4, 5, 6},
			expected: []int{1, 6, 2, 5, 3, 4},
		},
		{
			name:     "Seven nodes",
			input:    []int{1, 2, 3, 4, 5, 6, 7},
			expected: []int{1, 7, 2, 6, 3, 5, 4},
		},
		{
			name:     "Eight nodes",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8},
			expected: []int{1, 8, 2, 7, 3, 6, 4, 5},
		},
		{
			name:     "Large list (10 nodes)",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test all implementations
			implementations := []struct {
				name string
				fn   func(*ListNode)
			}{
				{"ReorderList", ReorderList},
				{"ReorderListStack", ReorderListStack},
				{"ReorderListArray", ReorderListArray},
			}

			for _, impl := range implementations {
				t.Run(impl.name, func(t *testing.T) {
					// Create a fresh list for each test
					head := NewListFromSlice(tt.input)
					
					// Apply the reorder function
					impl.fn(head)
					
					// Convert result to slice
					result := head.ToSlice()
					
					// Compare with expected
					if !reflect.DeepEqual(result, tt.expected) {
						t.Errorf("%s() = %v, expected %v", impl.name, result, tt.expected)
					}
				})
			}
		})
	}
}

func BenchmarkReorderList(b *testing.B) {
	// Create a large linked list
	values := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		values[i] = i + 1
	}

	b.Run("OptimalSolution", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			head := NewListFromSlice(values)
			ReorderList(head)
		}
	})

	b.Run("StackApproach", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			head := NewListFromSlice(values)
			ReorderListStack(head)
		}
	})

	b.Run("ArrayApproach", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			head := NewListFromSlice(values)
			ReorderListArray(head)
		}
	})
}

func TestReorderListEdgeCases(t *testing.T) {
	t.Run("Nil list", func(t *testing.T) {
		// Should not panic
		ReorderList(nil)
		ReorderListStack(nil)
		ReorderListArray(nil)
	})

	t.Run("Already reordered list", func(t *testing.T) {
		// Test that reordering an already reordered list doesn't change it
		// Note: The reorder operation is not idempotent for lists with 3+ nodes
		// For 2 nodes it is idempotent, but for 4+ nodes it changes the list
		tests := []struct {
			name     string
			input    []int
			expected []int
		}{
			{"Two nodes", []int{1, 2}, []int{1, 2}},
			{"Four nodes", []int{1, 4, 2, 3}, []int{1, 3, 4, 2}},
			{"Five nodes", []int{1, 5, 2, 4, 3}, []int{1, 3, 5, 4, 2}},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				head := NewListFromSlice(tt.input)
				
				ReorderList(head)
				result := head.ToSlice()
				
				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("ReorderList() changed already reordered list: got %v, expected %v", result, tt.expected)
				}
			})
		}
	})

	t.Run("Verify list structure after reorder", func(t *testing.T) {
		// Create a list and reorder it
		head := NewListFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8})
		ReorderList(head)
		
		// Verify the structure by traversing
		result := head.ToSlice()
		expected := []int{1, 8, 2, 7, 3, 6, 4, 5}
		
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("ReorderList() = %v, expected %v", result, expected)
		}
		
		// Also verify that the list is properly terminated (no cycles)
		if HasCycle(head) {
			t.Error("ReorderList() created a cycle in the list")
		}
	})
}

func TestReverseListHelper(t *testing.T) {
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
			name:     "Multiple nodes",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := NewListFromSlice(tt.input)
			reversed := reverseListHelper(head)
			result := reversed.ToSlice()
			
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("reverseListHelper() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestMergeListsHelper(t *testing.T) {
	tests := []struct {
		name     string
		first    []int
		second   []int
		expected []int
	}{
		{
			name:     "First longer than second",
			first:    []int{1, 2, 3},
			second:   []int{4, 5},
			expected: []int{1, 4, 2, 5, 3},
		},
		{
			name:     "Second longer than first",
			first:    []int{1, 2},
			second:   []int{3, 4, 5},
			expected: []int{1, 3, 2, 4, 5},
		},
		{
			name:     "Equal length",
			first:    []int{1, 2, 3},
			second:   []int{4, 5, 6},
			expected: []int{1, 4, 2, 5, 3, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			first := NewListFromSlice(tt.first)
			second := NewListFromSlice(tt.second)
			
			mergeLists(first, second)
			result := first.ToSlice()
			
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("mergeLists() = %v, expected %v", result, tt.expected)
			}
		})
	}
}