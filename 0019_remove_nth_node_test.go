package leetcode

import (
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name     string
		head     []int
		n        int
		expected []int
	}{
		{"Example 1", []int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{"Single node", []int{1}, 1, []int{}},
		{"Remove head", []int{1, 2}, 2, []int{2}},
		{"Remove tail", []int{1, 2}, 1, []int{1}},
		{"Two nodes", []int{1, 2}, 1, []int{1}},
		{"Three nodes remove middle", []int{1, 2, 3}, 2, []int{1, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := NewListFromSlice(tt.head)
			result := RemoveNthFromEnd(head, tt.n)
			resultSlice := ListToSlice(result)
			if !intSlicesEqual(resultSlice, tt.expected) {
				t.Errorf("RemoveNthFromEnd(%v, %d) = %v, expected %v", tt.head, tt.n, resultSlice, tt.expected)
			}
		})
	}
}

func TestRemoveNthFromEndAlt(t *testing.T) {
	tests := []struct {
		name     string
		head     []int
		n        int
		expected []int
	}{
		{"Example 1", []int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}},
		{"Single node", []int{1}, 1, []int{}},
		{"Remove head", []int{1, 2}, 2, []int{2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := NewListFromSlice(tt.head)
			result := RemoveNthFromEndAlt(head, tt.n)
			resultSlice := ListToSlice(result)
			if !intSlicesEqual(resultSlice, tt.expected) {
				t.Errorf("RemoveNthFromEndAlt(%v, %d) = %v, expected %v", tt.head, tt.n, resultSlice, tt.expected)
			}
		})
	}
}

// Helper function to compare int slices
func intSlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Helper function to convert linked list to slice
func ListToSlice(head *ListNode) []int {
	var result []int
	for current := head; current != nil; current = current.Next {
		result = append(result, current.Val)
	}
	return result
}

func BenchmarkRemoveNthFromEnd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		head := NewListFromSlice([]int{1, 2, 3, 4, 5})
		RemoveNthFromEnd(head, 2)
	}
}

func BenchmarkRemoveNthFromEndAlt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		head := NewListFromSlice([]int{1, 2, 3, 4, 5})
		RemoveNthFromEndAlt(head, 2)
	}
}
