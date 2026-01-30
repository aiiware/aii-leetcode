package linkedlists

import (
	"fmt"
	"testing"
    "leetcode/utils"
)

func TestReverseBetween(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		left     int
		right    int
		expected []int
	}{
		{
			name:     "Example 1",
			input:    []int{1, 2, 3, 4, 5},
			left:     2,
			right:    4,
			expected: []int{1, 4, 3, 2, 5},
		},
		{
			name:     "Example 2",
			input:    []int{5},
			left:     1,
			right:    1,
			expected: []int{5},
		},
		{
			name:     "Reverse entire list",
			input:    []int{1, 2, 3, 4, 5},
			left:     1,
			right:    5,
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "Reverse first two nodes",
			input:    []int{1, 2, 3, 4, 5},
			left:     1,
			right:    2,
			expected: []int{2, 1, 3, 4, 5},
		},
		{
			name:     "Reverse last two nodes",
			input:    []int{1, 2, 3, 4, 5},
			left:     4,
			right:    5,
			expected: []int{1, 2, 3, 5, 4},
		},
		{
			name:     "Reverse middle three nodes",
			input:    []int{1, 2, 3, 4, 5, 6},
			left:     2,
			right:    4,
			expected: []int{1, 4, 3, 2, 5, 6},
		},
		{
			name:     "Single node reversal",
			input:    []int{1, 2, 3},
			left:     2,
			right:    2,
			expected: []int{1, 2, 3},
		},
		{
			name:     "Two nodes list reverse both",
			input:    []int{1, 2},
			left:     1,
			right:    2,
			expected: []int{2, 1},
		},
		{
			name:     "Three nodes reverse middle",
			input:    []int{1, 2, 3},
			left:     2,
			right:    2,
			expected: []int{1, 2, 3},
		},
		{
			name:     "Large list reverse section",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			left:     3,
			right:    7,
			expected: []int{1, 2, 7, 6, 5, 4, 3, 8, 9, 10},
		},
		{
			name:     "Reverse from beginning to middle",
			input:    []int{1, 2, 3, 4, 5},
			left:     1,
			right:    3,
			expected: []int{3, 2, 1, 4, 5},
		},
		{
			name:     "Reverse from middle to end",
			input:    []int{1, 2, 3, 4, 5},
			left:     3,
			right:    5,
			expected: []int{1, 2, 5, 4, 3},
		},
		{
			name:     "Adjacent reversal",
			input:    []int{1, 2, 3, 4, 5, 6},
			left:     2,
			right:    3,
			expected: []int{1, 3, 2, 4, 5, 6},
		},
		{
			name:     "Empty list",
			input:    []int{},
			left:     1,
			right:    1,
			expected: []int{},
		},
		{
			name:     "Negative values",
			input:    []int{-5, -4, -3, -2, -1},
			left:     2,
			right:    4,
			expected: []int{-5, -2, -3, -4, -1},
		},
		{
			name:     "All same values",
			input:    []int{1, 1, 1, 1, 1},
			left:     2,
			right:    4,
			expected: []int{1, 1, 1, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := utils.NewListFromSlice(tt.input)
			result := ReverseBetween(head, tt.left, tt.right)
			actual := result.ToSlice()

			if !utils.SlicesEqual(actual, tt.expected) {
				t.Errorf("ReverseBetween(%v, %d, %d) = %v, expected %v",
					tt.input, tt.left, tt.right, actual, tt.expected)
			}
		})
	}
}

func TestAllReverseBetweenImplementations(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
		left  int
		right int
	}{
		{
			name:  "Example 1",
			input: []int{1, 2, 3, 4, 5},
			left:  2,
			right: 4,
		},
		{
			name:  "Single node",
			input: []int{5},
			left:  1,
			right: 1,
		},
		{
			name:  "Reverse entire",
			input: []int{1, 2, 3},
			left:  1,
			right: 3,
		},
		{
			name:  "Reverse first two",
			input: []int{1, 2, 3, 4},
			left:  1,
			right: 2,
		},
		{
			name:  "Reverse last two",
			input: []int{1, 2, 3, 4},
			left:  3,
			right: 4,
		},
	}

	implementations := []struct {
		name string
		fn   func(*utils.ListNode, int, int) *utils.ListNode
	}{
		{"reverseBetween", reverseBetween},
		{"reverseBetweenTwoPass", reverseBetweenTwoPass},
		{"reverseBetweenRecursive", reverseBetweenRecursive},
		{"reverseBetweenStack", reverseBetweenStack},
		{"reverseBetweenInPlace", reverseBetweenInPlace},
		{"reverseBetweenOptimized", reverseBetweenOptimized},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := ReverseBetween(utils.NewListFromSlice(tc.input), tc.left, tc.right)
			expectedSlice := expected.ToSlice()

			for _, impl := range implementations {
				t.Run(impl.name, func(t *testing.T) {
					head := utils.NewListFromSlice(tc.input)
					result := impl.fn(head, tc.left, tc.right)
					actual := result.ToSlice()

					if !utils.SlicesEqual(actual, expectedSlice) {
						t.Errorf("%s(%v, %d, %d) = %v, expected %v",
							impl.name, tc.input, tc.left, tc.right, actual, expectedSlice)
					}
				})
			}
		})
	}
}

func TestReverseBetweenEdgeCases(t *testing.T) {
	t.Run("Empty list", func(t *testing.T) {
		result := ReverseBetween(nil, 1, 1)
		if result != nil {
			t.Errorf("Expected nil for empty list, got %v", result.ToSlice())
		}
	})

	t.Run("Single node list", func(t *testing.T) {
		head := &utils.ListNode{Val: 1}
		result := ReverseBetween(head, 1, 1)
		if result.Val != 1 || result.Next != nil {
			t.Errorf("Expected single node 1, got %v", result.ToSlice())
		}
	})

	t.Run("left equals right", func(t *testing.T) {
		head := utils.NewListFromSlice([]int{1, 2, 3, 4, 5})
		result := ReverseBetween(head, 3, 3)
		expected := []int{1, 2, 3, 4, 5}
		actual := result.ToSlice()

		if !utils.SlicesEqual(actual, expected) {
			t.Errorf("When left == right, list should remain unchanged. Got %v, expected %v",
				actual, expected)
		}
	})

	t.Run("left = 1, right = length", func(t *testing.T) {
		head := utils.NewListFromSlice([]int{1, 2, 3, 4, 5})
		result := ReverseBetween(head, 1, 5)
		expected := []int{5, 4, 3, 2, 1}
		actual := result.ToSlice()

		if !utils.SlicesEqual(actual, expected) {
			t.Errorf("Reverse entire list failed. Got %v, expected %v",
				actual, expected)
		}
	})

	t.Run("Invalid left/right (left > right)", func(t *testing.T) {
		// According to constraints, left <= right, but we should handle gracefully
		head := utils.NewListFromSlice([]int{1, 2, 3})
		result := ReverseBetween(head, 3, 2)
		// Should return original list or handle error
		// Our implementation assumes left <= right per constraints
		actual := result.ToSlice()
		expected := []int{1, 2, 3} // Should remain unchanged or error

		if !utils.SlicesEqual(actual, expected) {
			t.Errorf("When left > right, behavior undefined. Got %v", actual)
		}
	})

	t.Run("Large list", func(t *testing.T) {
		// Create a large list
		size := 100
		input := make([]int, size)
		for i := 0; i < size; i++ {
			input[i] = i + 1
		}

		head := utils.NewListFromSlice(input)
		result := ReverseBetween(head, 25, 75)

		// Verify the reversed section
		actual := result.ToSlice()

		// Check first part unchanged
		for i := 0; i < 24; i++ {
			if actual[i] != input[i] {
				t.Errorf("First part changed at index %d: got %d, expected %d",
					i, actual[i], input[i])
			}
		}

		// Check reversed section
		for i := 24; i < 75; i++ {
			expected := input[74 - (i - 24)]
			if actual[i] != expected {
				t.Errorf("Reversed section mismatch at index %d: got %d, expected %d",
					i, actual[i], expected)
			}
		}

		// Check last part unchanged
		for i := 75; i < size; i++ {
			if actual[i] != input[i] {
				t.Errorf("Last part changed at index %d: got %d, expected %d",
					i, actual[i], input[i])
			}
		}
	})
}

func TestReverseBetweenProperties(t *testing.T) {
	// Property-based testing
	implementations := []struct {
		name string
		fn   func(*utils.ListNode, int, int) *utils.ListNode
	}{
		{"reverseBetween", reverseBetween},
		{"reverseBetweenTwoPass", reverseBetweenTwoPass},
		{"reverseBetweenInPlace", reverseBetweenInPlace},
		{"reverseBetweenOptimized", reverseBetweenOptimized},
	}

	testCases := []struct {
		input []int
		left  int
		right int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, 4},
		{[]int{1, 2, 3, 4, 5}, 1, 5},
		{[]int{1, 2, 3, 4, 5}, 1, 3},
		{[]int{1, 2, 3, 4, 5}, 3, 5},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4, 7},
	}

	for _, impl := range implementations {
		t.Run(impl.name+"_properties", func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(fmt.Sprintf("input=%v,left=%d,right=%d", tc.input, tc.left, tc.right), func(t *testing.T) {
					head := utils.NewListFromSlice(tc.input)
					result := impl.fn(head, tc.left, tc.right)
					actual := result.ToSlice()

					// Property 1: Length should remain the same
					if len(actual) != len(tc.input) {
						t.Errorf("Length changed: got %d, expected %d",
							len(actual), len(tc.input))
					}

					// Property 2: Elements outside [left, right] should be unchanged
					for i := 0; i < tc.left-1; i++ {
						if actual[i] != tc.input[i] {
							t.Errorf("Element before reversal changed at index %d: got %d, expected %d",
								i, actual[i], tc.input[i])
						}
					}
					for i := tc.right; i < len(tc.input); i++ {
						if actual[i] != tc.input[i] {
							t.Errorf("Element after reversal changed at index %d: got %d, expected %d",
								i, actual[i], tc.input[i])
						}
					}

					// Property 3: Elements in [left, right] should be reversed
					for i := tc.left - 1; i < tc.right; i++ {
						expectedIndex := tc.right - 1 - (i - (tc.left - 1))
						if actual[i] != tc.input[expectedIndex] {
							t.Errorf("Reversal mismatch at index %d: got %d, expected %d (from index %d)",
								i, actual[i], tc.input[expectedIndex], expectedIndex)
						}
					}

					// Property 4: The list should still be a valid linked list
					current := result
					count := 0
					for current != nil {
						count++
						current = current.Next
					}
					if count != len(tc.input) {
						t.Errorf("Linked list corrupted: expected %d nodes, found %d",
							len(tc.input), count)
					}
				})
			}
		})
	}
}

func BenchmarkReverseBetween(b *testing.B) {
	// Test cases of different sizes
	testCases := []struct {
		name  string
		input []int
		left  int
		right int
	}{
		{"Small", []int{1, 2, 3, 4, 5}, 2, 4},
		{"Medium", utils.MakeRange(1, 100), 25, 75},
		{"Large", utils.MakeRange(1, 1000), 250, 750},
		{"Reverse entire", utils.MakeRange(1, 100), 1, 100},
		{"Reverse beginning", utils.MakeRange(1, 100), 1, 50},
		{"Reverse end", utils.MakeRange(1, 100), 51, 100},
	}

	implementations := []struct {
		name string
		fn   func(*utils.ListNode, int, int) *utils.ListNode
	}{
		{"reverseBetween", reverseBetween},
		{"reverseBetweenTwoPass", reverseBetweenTwoPass},
		{"reverseBetweenRecursive", reverseBetweenRecursive},
		{"reverseBetweenStack", reverseBetweenStack},
		{"reverseBetweenInPlace", reverseBetweenInPlace},
		{"reverseBetweenOptimized", reverseBetweenOptimized},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for _, impl := range implementations {
				b.Run(impl.name, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						head := utils.NewListFromSlice(tc.input)
						impl.fn(head, tc.left, tc.right)
					}
				})
			}
		})
	}
}

func BenchmarkReverseBetweenWorstCase(b *testing.B) {
	// Worst case: reverse almost entire large list
	size := 1000
	input := utils.MakeRange(1, size)
	left, right := 2, size-1

	b.ResetTimer()

	b.Run("reverseBetween", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			head := utils.NewListFromSlice(input)
			reverseBetween(head, left, right)
		}
	})

	b.Run("reverseBetweenOptimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			head := utils.NewListFromSlice(input)
			reverseBetweenOptimized(head, left, right)
		}
	})

	b.Run("reverseBetweenRecursive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			head := utils.NewListFromSlice(input)
			reverseBetweenRecursive(head, left, right)
		}
	})
}