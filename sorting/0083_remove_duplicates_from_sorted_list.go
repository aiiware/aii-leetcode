package sorting

import "leetcode/utils"

// DeleteDuplicates solves LeetCode problem 0083: Remove Duplicates from Sorted List
// Difficulty: Easy
// Tags: Linked List
//
// Given the head of a sorted linked list, delete all duplicates such that each element
// appears only once. Return the linked list sorted as well.
//
// Example 1:
// Input: head = [1,1,2]
// Output: [1,2]
//
// Example 2:
// Input: head = [1,1,2,3,3]
// Output: [1,2,3]
//
// Constraints:
// The number of nodes in the list is in the range [0, 300].
// -100 <= Node.val <= 100
// The list is guaranteed to be sorted in ascending order.
//
// Time complexity: O(n)
// Space complexity: O(1)
func DeleteDuplicates(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	current := head
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			// Skip duplicate node
			current.Next = current.Next.Next
		} else {
			// Move to next node
			current = current.Next
		}
	}

	return head
}

// DeleteDuplicatesRecursive is a recursive implementation
func DeleteDuplicatesRecursive(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// Remove duplicates from the rest of the list
	head.Next = DeleteDuplicatesRecursive(head.Next)

	// If current node duplicates next node, skip next node
	if head.Next != nil && head.Val == head.Next.Val {
		return head.Next
	}

	return head
}

// DeleteDuplicatesTwoPointers uses two pointers explicitly
func DeleteDuplicatesTwoPointers(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev := head
	curr := head.Next

	for curr != nil {
		if prev.Val == curr.Val {
			// Skip duplicate
			prev.Next = curr.Next
		} else {
			// Move prev forward
			prev = curr
		}
		curr = curr.Next
	}

	return head
}

// DeleteDuplicatesWithDummy uses a dummy node for consistency
func DeleteDuplicatesWithDummy(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &utils.ListNode{Next: head}
	prev := dummy
	curr := head

	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			// Skip all duplicates
			for curr.Next != nil && curr.Val == curr.Next.Val {
				curr = curr.Next
			}
			prev.Next = curr
		} else {
			prev = curr
			curr = curr.Next
		}
	}

	return dummy.Next
}

// DeleteDuplicatesEarlyExit adds early exit optimizations
func DeleteDuplicatesEarlyExit(head *utils.ListNode) *utils.ListNode {
	// Early exit for empty or single node list
	if head == nil || head.Next == nil {
		return head
	}

	current := head
	for current.Next != nil {
		if current.Val == current.Next.Val {
			// Skip duplicate
			current.Next = current.Next.Next
		} else {
			// Only move forward if next node is different
			current = current.Next
		}
	}

	return head
}

// DeleteDuplicatesStack uses a stack-like approach
func DeleteDuplicatesStack(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// Create result list
	dummy := &utils.ListNode{}
	tail := dummy
	current := head

	for current != nil {
		// Append current node to result
		tail.Next = current
		tail = tail.Next
		
		// Skip all duplicates
		for current.Next != nil && current.Val == current.Next.Val {
			current = current.Next
		}
		
		current = current.Next
		// Important: disconnect the node from the rest
		tail.Next = nil
	}

	return dummy.Next
}

// DeleteDuplicatesGeneric allows specifying max duplicates (generalized version)
func DeleteDuplicatesGeneric(head *utils.ListNode, maxDuplicates int) *utils.ListNode {
	if head == nil || maxDuplicates <= 0 {
		return nil
	}

	dummy := &utils.ListNode{Next: head}
	prev := dummy
	current := head

	for current != nil {
		// Count duplicates of current value
		count := 1
		runner := current
		for runner.Next != nil && runner.Val == runner.Next.Val {
			runner = runner.Next
			count++
		}

		if count > maxDuplicates {
			// Skip extra duplicates
			for i := 0; i < count-maxDuplicates; i++ {
				current = current.Next
			}
			prev.Next = current
		} else {
			// Keep all nodes (within limit)
			for i := 0; i < count; i++ {
				prev = current
				current = current.Next
			}
		}
	}

	return dummy.Next
}