package trees

// ListNode represents a node in a singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// DeleteDuplicates removes duplicates from a sorted linked list
// Difficulty: Easy
// Tags: Linked List
//
// Given the head of a sorted linked list, delete all duplicates such that each element appears only once.
// Return the linked list sorted as well.
//
// Time complexity: O(n), Space complexity: O(1)
func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	
	current := head
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	
	return head
}