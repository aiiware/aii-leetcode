package linkedlists

import (
	"fmt"
	"testing"
    "leetcode/utils"
)

func TestPartitionList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		x        int
		expected []int
	}{
		{
			name:     "Example 1",
			input:    []int{1, 4, 3, 2, 5, 2},
			x:        3,
			expected: []int{1, 2, 2, 4, 3, 5},
		},
		{
			name:     "Example 2",
			input:    []int{2, 1},
			x:        2,
			expected: []int{1, 2},
		},
		{
			name:     "Empty list",
			input:    []int{},
			x:        5,
			expected: []int{},
		},
		{
			name:     "Single element less than x",
			input:    []int{3},
			x:        5,
			expected: []int{3},
		},
		{
			name:     "Single element greater than x",
			input:    []int{7},
			x:        5,
			expected: []int{7},
		},
		{
			name:     "All elements less than x",
			input:    []int{1, 2, 3, 4},
			x:        5,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "All elements greater than or equal to x",
			input:    []int{5, 6, 7, 8},
			x:        5,
			expected: []int{5, 6, 7, 8},
		},
		{
			name:     "Mixed elements with duplicates",
			input:    []int{3, 1, 4, 3, 2, 5, 3, 2},
			x:        3,
			expected: []int{1, 2, 2, 3, 4, 3, 5, 3},
		},
		{
			name:     "Already partitioned",
			input:    []int{1, 2, 3, 4, 5},
			x:        3,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Reverse order",
			input:    []int{5, 4, 3, 2, 1},
			x:        3,
			expected: []int{2, 1, 5, 4, 3},
		},
		{
			name:     "All equal to x",
			input:    []int{3, 3, 3, 3},
			x:        3,
			expected: []int{3, 3, 3, 3},
		},
		{
			name:     "Large x value",
			input:    []int{1, 2, 3, 4, 5},
			x:        10,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Small x value",
			input:    []int{1, 2, 3, 4, 5},
			x:        0,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Negative values",
			input:    []int{-1, -2, -3, -4, -5},
			x:        -3,
			expected: []int{-4, -5, -1, -2, -3},
		},
		{
			name:     "Mixed positive and negative",
			input:    []int{-2, 5, -1, 3, 0, -4},
			x:        0,
			expected: []int{-2, -1, -4, 5, 3, 0},
		},
		{
			name:     "Complex case 1",
			input:    []int{4, 3, 2, 5, 2, 1, 6, 3},
			x:        4,
			expected: []int{3, 2, 2, 1, 3, 4, 5, 6},
		},
		{
			name:     "Complex case 2",
			input:    []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			x:        5,
			expected: []int{4, 3, 2, 1, 9, 8, 7, 6, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := utils.NewListFromSlice(tt.input)
			result := PartitionList(head, tt.x)
			actual := result.ToSlice()
			
			if !utils.SlicesEqual(actual, tt.expected) {
				t.Errorf("PartitionList(%v, %d) = %v, expected %v", 
					tt.input, tt.x, actual, tt.expected)
			}
		})
	}
}

func TestAllPartitionListImplementations(t *testing.T) {
	testCases := []struct {
		input    []int
		x        int
	}{
		{[]int{1, 4, 3, 2, 5, 2}, 3},
		{[]int{2, 1}, 2},
		{[]int{}, 5},
		{[]int{3, 1, 4, 3, 2, 5, 3, 2}, 3},
		{[]int{-2, 5, -1, 3, 0, -4}, 0},
	}

	implementations := []struct {
		name string
		fn   func(*utils.ListNode, int) *utils.ListNode
	}{
		{"partitionList", partitionList},
		{"partitionListTwoPass", partitionListTwoPass},
		{"partitionListInPlace", partitionListInPlace},
		{"partitionListRecursive", partitionListRecursive},
		{"partitionListOptimized", partitionListOptimized},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input:%v_x:%d", tc.input, tc.x), func(t *testing.T) {
			head := utils.NewListFromSlice(tc.input)
			expected := PartitionList(head, tc.x).ToSlice()
			
			for _, impl := range implementations {
				t.Run(impl.name, func(t *testing.T) {
					head := utils.NewListFromSlice(tc.input)
					result := impl.fn(head, tc.x)
					actual := result.ToSlice()
					
					if !utils.SlicesEqual(actual, expected) {
						t.Errorf("%s(%v, %d) = %v, expected %v", 
							impl.name, tc.input, tc.x, actual, expected)
					}
				})
			}
		})
	}
}

func TestPartitionListEdgeCases(t *testing.T) {
	t.Run("Nil head", func(t *testing.T) {
		result := PartitionList(nil, 5)
		if result != nil {
			t.Errorf("Expected nil, got %v", result)
		}
	})

	t.Run("Single node", func(t *testing.T) {
		head := &utils.ListNode{Val: 3}
		result := PartitionList(head, 5)
		if result.Val != 3 || result.Next != nil {
			t.Errorf("Expected single node with value 3, got %v", result.ToSlice())
		}
	})

	t.Run("Preserve relative order", func(t *testing.T) {
		// Test that relative order is preserved within each partition
		input := []int{1, 4, 3, 2, 5, 2}
		head := utils.NewListFromSlice(input)
		result := PartitionList(head, 3)
		actual := result.ToSlice()
		expected := []int{1, 2, 2, 4, 3, 5}
		
		if !utils.SlicesEqual(actual, expected) {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
		
		// Verify that all elements < 3 are in original order: 1, 2, 2
		// and all elements >= 3 are in original order: 4, 3, 5
		lessPartition := []int{}
		greaterPartition := []int{}
		for _, val := range input {
			if val < 3 {
				lessPartition = append(lessPartition, val)
			} else {
				greaterPartition = append(greaterPartition, val)
			}
		}
		
		// Check less partition
		resultLess := []int{}
		current := result
		for current != nil && current.Val < 3 {
			resultLess = append(resultLess, current.Val)
			current = current.Next
		}
		
		if !utils.SlicesEqual(resultLess, lessPartition) {
			t.Errorf("Less partition order not preserved. Expected %v, got %v", 
				lessPartition, resultLess)
		}
	})
}

func TestPartitionListProperties(t *testing.T) {
	// Property-based testing
	tests := []struct {
		name string
		fn   func(*utils.ListNode, int) *utils.ListNode
	}{
		{"partitionList", partitionList},
		{"partitionListTwoPass", partitionListTwoPass},
		{"partitionListInPlace", partitionListInPlace},
		{"partitionListOptimized", partitionListOptimized},
	}

	for _, impl := range tests {
		t.Run(impl.name+"_properties", func(t *testing.T) {
			// Test 1: Result should contain all original elements
			input := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
			head := utils.NewListFromSlice(input)
			result := impl.fn(head, 5)
			resultSlice := result.ToSlice()
			
			// Count occurrences
			originalCount := make(map[int]int)
			resultCount := make(map[int]int)
			
			for _, val := range input {
				originalCount[val]++
			}
			for _, val := range resultSlice {
				resultCount[val]++
			}
			
			for val, count := range originalCount {
				if resultCount[val] != count {
					t.Errorf("Element %d appears %d times in result, expected %d", 
						val, resultCount[val], count)
				}
			}
			
			// Test 2: All elements before partition point should be < x
			// and all elements after should be >= x
			foundPartitionPoint := false
			current := result
			for current != nil {
				if !foundPartitionPoint {
					if current.Val >= 5 {
						foundPartitionPoint = true
					}
				} else {
					if current.Val < 5 {
						t.Errorf("Found element %d < 5 after partition point", current.Val)
					}
				}
				current = current.Next
			}
		})
	}
}

func BenchmarkPartitionList(b *testing.B) {
	// Create a large linked list for benchmarking
	size := 10000
	input := make([]int, size)
	for i := 0; i < size; i++ {
		input[i] = i % 100 // Values from 0 to 99
	}
	x := 50

	b.ResetTimer()
	
	b.Run("partitionList", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// Create a fresh copy for each iteration
			h := utils.NewListFromSlice(input)
			partitionList(h, x)
		}
	})
	
	b.Run("partitionListTwoPass", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h := utils.NewListFromSlice(input)
			partitionListTwoPass(h, x)
		}
	})
	
	b.Run("partitionListInPlace", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h := utils.NewListFromSlice(input)
			partitionListInPlace(h, x)
		}
	})
	
	b.Run("partitionListRecursive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h := utils.NewListFromSlice(input)
			partitionListRecursive(h, x)
		}
	})
	
	b.Run("partitionListOptimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h := utils.NewListFromSlice(input)
			partitionListOptimized(h, x)
		}
	})
}

func BenchmarkPartitionListSmall(b *testing.B) {
	input := []int{1, 4, 3, 2, 5, 2}
	x := 3

	b.ResetTimer()
	
	b.Run("partitionList", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h := utils.NewListFromSlice(input)
			partitionList(h, x)
		}
	})
	
	b.Run("partitionListOptimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			h := utils.NewListFromSlice(input)
			partitionListOptimized(h, x)
		}
	})
}