package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name     string
		l1       *ListNode
		l2       *ListNode
		expected *ListNode
	}{
		{
			name:     "Example 1: 342 + 465 = 807",
			l1:       createList([]int{2, 4, 3}),
			l2:       createList([]int{5, 6, 4}),
			expected: createList([]int{7, 0, 8}),
		},
		{
			name:     "Example 2: 0 + 0 = 0",
			l1:       createList([]int{0}),
			l2:       createList([]int{0}),
			expected: createList([]int{0}),
		},
		{
			name:     "Example 3: 9999999 + 9999 = 10009998",
			l1:       createList([]int{9, 9, 9, 9, 9, 9, 9}),
			l2:       createList([]int{9, 9, 9, 9}),
			expected: createList([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
		{
			name:     "Different length lists",
			l1:       createList([]int{1, 2, 3}),
			l2:       createList([]int{4, 5}),
			expected: createList([]int{5, 7, 3}),
		},
		{
			name:     "Carry propagation",
			l1:       createList([]int{9, 9, 9}),
			l2:       createList([]int{1}),
			expected: createList([]int{0, 0, 0, 1}),
		},
		{
			name:     "One empty list",
			l1:       createList([]int{1, 2, 3}),
			l2:       createList([]int{}),
			expected: createList([]int{1, 2, 3}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AddTwoNumbers(tt.l1, tt.l2)
			assert.Equal(t, tt.expected, result,
				"AddTwoNumbers(%v, %v) = %v, expected %v",
				listToSlice(tt.l1), listToSlice(tt.l2), listToSlice(result), listToSlice(tt.expected))
		})
	}
}

func BenchmarkAddTwoNumbers(b *testing.B) {
	// Create two large numbers for benchmarking
	l1 := createList(make([]int, 1000))
	l2 := createList(make([]int, 1000))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AddTwoNumbers(l1, l2)
	}
}

// Helper function to create a linked list from slice
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

// Helper function to convert linked list to slice
func listToSlice(head *ListNode) []int {
	var result []int
	current := head

	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}

	return result
}
