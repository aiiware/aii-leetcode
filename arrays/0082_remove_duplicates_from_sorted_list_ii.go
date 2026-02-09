package arrays

// RemoveDuplicatesFromSortedListII removes all duplicate nodes from a sorted linked list
// that appear more than once, leaving only distinct nodes.
//
// LeetCode problem 0082: Remove Duplicates from Sorted List II
// Difficulty: Medium
// Tags: Linked Lists, Two Pointers
//
// Time complexity: O(n), Space complexity: O(1)
func RemoveDuplicatesFromSortedListII(head *ListNode) *ListNode {
	// Create a dummy node to simplify edge cases
	dummy := &ListNode{Val: 0, Next: head}
	prev := dummy
	current := head

	for current != nil {
		// Check if current node has duplicates
		if current.Next != nil && current.Val == current.Next.Val {
			// Skip all nodes with the same value
			val := current.Val
			for current != nil && current.Val == val {
				current = current.Next
			}
			// Connect the previous node to skip the duplicates
			prev.Next = current
		} else {
			// No duplicates, move forward
			prev = current
			current = current.Next
		}
	}

	return dummy.Next
}
