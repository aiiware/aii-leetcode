package leetcode

// MergeTwoLists solves LeetCode problem 0021: Merge Two Sorted Lists
// Difficulty: Easy
// Tags: Linked List, Recursion
//
// You are given the heads of two sorted linked lists list1 and list2.
// Merge the two lists into one sorted list. The list should be made by
// splicing together the nodes of the first two lists.
// Return the head of the merged linked list.
//
// Time complexity: O(n + m), Space complexity: O(1) (iterative) or O(n + m) (recursive)
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Create a dummy node to simplify edge cases
	dummy := &ListNode{}
	current := dummy

	// Merge while both lists have nodes
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

	// Append remaining nodes from either list
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return dummy.Next
}

// MergeTwoListsRecursive is a recursive implementation
func MergeTwoListsRecursive(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		list1.Next = MergeTwoListsRecursive(list1.Next, list2)
		return list1
	} else {
		list2.Next = MergeTwoListsRecursive(list1, list2.Next)
		return list2
	}
}