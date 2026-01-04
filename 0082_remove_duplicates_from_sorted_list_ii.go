package leetcode

// DeleteDuplicatesII solves LeetCode problem 0082: Remove Duplicates from Sorted List II
// Difficulty: Medium
// Tags: Linked List, Two Pointers
//
// Given the head of a sorted linked list, delete all nodes that have duplicate numbers,
// leaving only distinct numbers from the original list. Return the linked list sorted as well.
//
// Example 1:
// Input: head = [1,2,3,3,4,4,5]
// Output: [1,2,5]
//
// Example 2:
// Input: head = [1,1,1,2,3]
// Output: [2,3]
//
// Constraints:
// The number of nodes in the list is in the range [0, 300].
// -100 <= Node.val <= 100
// The list is guaranteed to be sorted in ascending order.
//
// Time complexity: O(n)
// Space complexity: O(1)
func DeleteDuplicatesII(head *ListNode) *ListNode {
	// Create a dummy node to handle edge cases (when head needs to be deleted)
	dummy := &ListNode{Next: head}
	prev := dummy
	current := head

	for current != nil {
		// Skip all duplicates
		if current.Next != nil && current.Val == current.Next.Val {
			// Move current until we find a different value or end of list
			for current.Next != nil && current.Val == current.Next.Val {
				current = current.Next
			}
			// Skip all duplicates
			prev.Next = current.Next
		} else {
			// No duplicates, move prev forward
			prev = prev.Next
		}
		current = current.Next
	}

	return dummy.Next
}

// DeleteDuplicatesIIRecursive is a recursive implementation
func DeleteDuplicatesIIRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// If current node is a duplicate
	if head.Next != nil && head.Val == head.Next.Val {
		// Skip all nodes with the same value
		for head.Next != nil && head.Val == head.Next.Val {
			head = head.Next
		}
		// Skip the last duplicate node too
		return DeleteDuplicatesIIRecursive(head.Next)
	}

	// Current node is not a duplicate
	head.Next = DeleteDuplicatesIIRecursive(head.Next)
	return head
}

// DeleteDuplicatesIITwoPointers uses two pointers more explicitly
func DeleteDuplicatesIITwoPointers(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	prev := dummy
	curr := head

	for curr != nil {
		// Check if current node has duplicates
		hasDuplicate := false
		for curr.Next != nil && curr.Val == curr.Next.Val {
			curr = curr.Next
			hasDuplicate = true
		}

		if hasDuplicate {
			// Skip all duplicates
			prev.Next = curr.Next
		} else {
			// No duplicates, keep this node
			prev = prev.Next
		}

		curr = curr.Next
	}

	return dummy.Next
}

// DeleteDuplicatesIIWithCounter uses a counter to track duplicates
func DeleteDuplicatesIIWithCounter(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	prev := dummy
	curr := head

	for curr != nil {
		// Count duplicates of current value
		count := 1
		runner := curr
		for runner.Next != nil && runner.Val == runner.Next.Val {
			runner = runner.Next
			count++
		}

		if count > 1 {
			// Skip all duplicates
			prev.Next = runner.Next
			curr = runner.Next
		} else {
			// Keep this unique node
			prev = curr
			curr = curr.Next
		}
	}

	return dummy.Next
}

// DeleteDuplicatesIIEarlyExit adds early exit optimizations
func DeleteDuplicatesIIEarlyExit(head *ListNode) *ListNode {
	// Early exit for empty or single node list
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{Next: head}
	prev := dummy
	curr := head

	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			// Found duplicates, skip them all
			duplicateValue := curr.Val
			for curr != nil && curr.Val == duplicateValue {
				curr = curr.Next
			}
			prev.Next = curr
		} else {
			// No duplicates, move both pointers
			prev = curr
			curr = curr.Next
		}
	}

	return dummy.Next
}

// DeleteDuplicatesIIStack uses a stack-like approach (though not really a stack)
func DeleteDuplicatesIIStack(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// Create result list
	dummy := &ListNode{}
	tail := dummy
	curr := head

	for curr != nil {
		// Check if current node is unique
		isUnique := true
		
		// Look ahead to see if there are duplicates
		if curr.Next != nil && curr.Val == curr.Next.Val {
			isUnique = false
			// Skip all duplicates
			duplicateValue := curr.Val
			for curr != nil && curr.Val == duplicateValue {
				curr = curr.Next
			}
		}

		if isUnique {
			// Append unique node to result
			tail.Next = curr
			tail = tail.Next
			curr = curr.Next
			// Important: disconnect the node from the rest
			tail.Next = nil
		}
	}

	return dummy.Next
}