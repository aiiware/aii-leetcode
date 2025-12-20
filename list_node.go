package leetcode

// ListNode represents a node in a singly-linked list.
// This is the standard definition used in LeetCode problems.
type ListNode struct {
	Val  int
	Next *ListNode
}

// NewListNode creates a new ListNode with the given value.
func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

// NewListFromSlice creates a linked list from a slice of integers.
// Returns the head of the linked list.
func NewListFromSlice(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	current := head
	for i := 1; i < len(vals); i++ {
		current.Next = &ListNode{Val: vals[i]}
		current = current.Next
	}
	return head
}

// ToSlice converts a linked list to a slice of integers.
func (l *ListNode) ToSlice() []int {
	var result []int
	current := l
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}

// Equal compares two linked lists for equality.
func (l *ListNode) Equal(other *ListNode) bool {
	a, b := l, other
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return a == nil && b == nil
}