package linkedlists


/*
Difficulty: Medium
Tags: [Add relevant tags]
Companies: [Add company names]
*/

import "leetcode/utils"

// Problem 0086: Partition List
// 
// Given the head of a linked list and a value x, partition it such that all nodes 
// less than x come before nodes greater than or equal to x. You should preserve 
// the original relative order of nodes in each of the two partitions.
//
// Example 1:
// Input: head = [1,4,3,2,5,2], x = 3
// Output: [1,2,2,4,3,5]
//
// Example 2:
// Input: head = [2,1], x = 2
// Output: [1,2]
//
// Constraints:
// - The number of nodes in the list is in the range [0, 200].
// - -100 <= Node.val <= 100
// - -200 <= x <= 200

// partitionList is the main solution function.
// Time complexity: O(n), Space complexity: O(1)
func partitionList(head *utils.ListNode, x int) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// Create dummy nodes for two partitions
	lessDummy := &utils.ListNode{}
	greaterDummy := &utils.ListNode{}
	
	// Pointers for building the two lists
	lessTail := lessDummy
	greaterTail := greaterDummy
	
	// Traverse the original list
	current := head
	for current != nil {
		if current.Val < x {
			// Add to less list
			lessTail.Next = current
			lessTail = lessTail.Next
		} else {
			// Add to greater or equal list
			greaterTail.Next = current
			greaterTail = greaterTail.Next
		}
		current = current.Next
	}
	
	// Terminate the greater list
	greaterTail.Next = nil
	
	// Connect less list to greater list
	lessTail.Next = greaterDummy.Next
	
	return lessDummy.Next
}

// partitionListTwoPass is an alternative solution using two passes.
// First pass collects nodes less than x, second pass collects nodes >= x.
// Time complexity: O(n), Space complexity: O(1)
func partitionListTwoPass(head *utils.ListNode, x int) *utils.ListNode {
	if head == nil {
		return nil
	}
	
	// Create a new list
	dummy := &utils.ListNode{}
	tail := dummy
	
	// First pass: add nodes with value < x
	current := head
	for current != nil {
		if current.Val < x {
			tail.Next = &utils.ListNode{Val: current.Val}
			tail = tail.Next
		}
		current = current.Next
	}
	
	// Second pass: add nodes with value >= x
	current = head
	for current != nil {
		if current.Val >= x {
			tail.Next = &utils.ListNode{Val: current.Val}
			tail = tail.Next
		}
		current = current.Next
	}
	
	return dummy.Next
}

// partitionListInPlace is another in-place solution.
// Time complexity: O(n), Space complexity: O(1)
func partitionListInPlace(head *utils.ListNode, x int) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	
	// Find the first node >= x
	var beforeX, afterX *utils.ListNode
	var beforeTail, afterTail *utils.ListNode
	
	current := head
	for current != nil {
		next := current.Next
		current.Next = nil // Detach current node
		
		if current.Val < x {
			if beforeX == nil {
				beforeX = current
				beforeTail = current
			} else {
				beforeTail.Next = current
				beforeTail = current
			}
		} else {
			if afterX == nil {
				afterX = current
				afterTail = current
			} else {
				afterTail.Next = current
				afterTail = current
			}
		}
		
		current = next
	}
	
	// Connect the two lists
	if beforeX == nil {
		return afterX
	}
	
	beforeTail.Next = afterX
	return beforeX
}

// partitionListRecursive is a recursive solution (not recommended for large lists).
// Time complexity: O(n), Space complexity: O(n) due to recursion stack
func partitionListRecursive(head *utils.ListNode, x int) *utils.ListNode {
	if head == nil {
		return nil
	}
	
	// Partition the rest of the list
	rest := partitionListRecursive(head.Next, x)
	
	// Insert current node at the beginning or end based on its value
	if head.Val < x {
		head.Next = rest
		return head
	} else {
		// Find the last node in the < x partition
		if rest == nil {
			return head
		}
		
		// Find the first node >= x in the partitioned list
		current := rest
		for current.Next != nil && current.Next.Val < x {
			current = current.Next
		}
		
		// Insert current node after the < x partition
		head.Next = current.Next
		current.Next = head
		return rest
	}
}

// partitionListOptimized is the most optimized solution.
// Uses two pointers and maintains relative order.
func partitionListOptimized(head *utils.ListNode, x int) *utils.ListNode {
	// Create two dummy nodes
	lessHead := &utils.ListNode{}
	greaterHead := &utils.ListNode{}
	
	less := lessHead
	greater := greaterHead
	
	for head != nil {
		if head.Val < x {
			less.Next = head
			less = less.Next
		} else {
			greater.Next = head
			greater = greater.Next
		}
		head = head.Next
	}
	
	// Connect the two lists
	greater.Next = nil
	less.Next = greaterHead.Next
	
	return lessHead.Next
}

// PartitionList is the public interface function.
// It uses the optimized solution by default.
func PartitionList(head *utils.ListNode, x int) *utils.ListNode {
	return partitionListOptimized(head, x)
}