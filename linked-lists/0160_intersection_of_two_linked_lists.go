package linkedlists

import "leetcode/utils"

// 0160. Intersection of Two Linked Lists
// https://leetcode.com/problems/intersection-of-two-linked-lists

// utils.ListNode definition (from common.go)
// type utils.ListNode struct {
//     Val  int
//     Next *utils.ListNode
// }

// getIntersectionNode is the main solution function
func getIntersectionNode(headA, headB *utils.ListNode) *utils.ListNode {
	// Solution 1: Two pointers technique
	return getIntersectionNodeTwoPointers(headA, headB)
}

// ===== Solution 1: Two pointers technique =====
// Time complexity: O(m + n) where m and n are lengths of the lists
// Space complexity: O(1)

func getIntersectionNodeTwoPointers(headA, headB *utils.ListNode) *utils.ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	ptrA, ptrB := headA, headB

	// Traverse both lists
	// When ptrA reaches end, switch to headB
	// When ptrB reaches end, switch to headA
	// If they intersect, they will meet at intersection point
	// If they don't intersect, both will become nil at the same time

	for ptrA != ptrB {
		// Move ptrA forward
		if ptrA == nil {
			ptrA = headB
		} else {
			ptrA = ptrA.Next
		}

		// Move ptrB forward
		if ptrB == nil {
			ptrB = headA
		} else {
			ptrB = ptrB.Next
		}
	}

	return ptrA // This will be either intersection node or nil
}

// ===== Solution 2: Calculate lengths difference =====
// Time complexity: O(m + n)
// Space complexity: O(1)

func getIntersectionNodeLength(headA, headB *utils.ListNode) *utils.ListNode {
	// Helper function to get length of a linked list
	getLength := func(head *utils.ListNode) int {
		length := 0
		for head != nil {
			length++
			head = head.Next
		}
		return length
	}

	lenA := getLength(headA)
	lenB := getLength(headB)

	// Align the starting points
	currA, currB := headA, headB

	// Move the longer list's pointer forward by the difference
	if lenA > lenB {
		for i := 0; i < lenA-lenB; i++ {
			currA = currA.Next
		}
	} else if lenB > lenA {
		for i := 0; i < lenB-lenA; i++ {
			currB = currB.Next
		}
	}

	// Now both lists are aligned, traverse together
	for currA != nil && currB != nil {
		if currA == currB {
			return currA
		}
		currA = currA.Next
		currB = currB.Next
	}

	return nil
}

// ===== Solution 3: Hash set approach =====
// Time complexity: O(m + n)
// Space complexity: O(m) or O(n)

func getIntersectionNodeHash(headA, headB *utils.ListNode) *utils.ListNode {
	visited := make(map[*utils.ListNode]bool)

	// Traverse first list and mark all nodes
	curr := headA
	for curr != nil {
		visited[curr] = true
		curr = curr.Next
	}

	// Traverse second list and check for visited nodes
	curr = headB
	for curr != nil {
		if visited[curr] {
			return curr
		}
		curr = curr.Next
	}

	return nil
}

// ===== Solution 4: Cycle detection approach =====
// Time complexity: O(m + n)
// Space complexity: O(1)

func getIntersectionNodeCycle(headA, headB *utils.ListNode) *utils.ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	// Find tail of list A and connect it to headB
	// This creates a cycle if there's an intersection
	tailA := headA
	for tailA.Next != nil {
		tailA = tailA.Next
	}

	// Save the original next pointer to restore later
	originalNext := tailA.Next
	tailA.Next = headB

	// Now use Floyd's cycle detection algorithm
	slow, fast := headA, headA

	// Find meeting point if cycle exists
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			// Cycle detected, find intersection
			ptr1 := headA
			ptr2 := slow

			for ptr1 != ptr2 {
				ptr1 = ptr1.Next
				ptr2 = ptr2.Next
			}

			// Restore the original list structure
			tailA.Next = originalNext
			return ptr1
		}
	}

	// No cycle detected
	tailA.Next = originalNext
	return nil
}