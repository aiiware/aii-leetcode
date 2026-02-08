package arrays

// MergeTwoSortedLists merges two sorted linked lists and returns a new sorted list.
// The new list is made by splicing together the nodes of the first two lists.
//
// Example 1:
// Input: list1 = [1,2,4], list2 = [1,3,4]
// Output: [1,1,2,3,4,4]
//
// Example 2:
// Input: list1 = [], list2 = []
// Output: []
//
// Example 3:
// Input: list1 = [], list2 = [0]
// Output: [0]
//
// Time complexity: O(m + n), Space complexity: O(1)
func MergeTwoSortedLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Create a dummy node to simplify the logic
	dummy := &ListNode{}
	current := dummy

	// Traverse both lists and merge
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	// Attach remaining nodes (if any)
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return dummy.Next
}