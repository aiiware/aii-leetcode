package leetcode

// ReorderList solves LeetCode problem 0143: Reorder List
// Difficulty: Medium
// Tags: Linked List, Two Pointers, Stack, Recursion
//
// You are given the head of a singly linked-list. The list can be represented as:
// L0 → L1 → … → Ln-1 → Ln
//
// Reorder the list to be on the following form:
// L0 → Ln → L1 → Ln-1 → L2 → Ln-2 → …
//
// You may not modify the values in the list's nodes. Only nodes themselves may be changed.
//
// Example 1:
// Input: head = [1,2,3,4]
// Output: [1,4,2,3]
//
// Example 2:
// Input: head = [1,2,3,4,5]
// Output: [1,5,2,4,3]
//
// Constraints:
// - The number of nodes in the list is in the range [1, 5 * 10^4].
// - 1 <= Node.val <= 1000
//
// Time complexity: O(n), Space complexity: O(1) for optimal solution, O(n) for stack approach
func ReorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	// Step 1: Find the middle of the list using slow-fast pointers
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Step 2: Reverse the second half of the list
	secondHalf := reverseListHelper(slow.Next)
	slow.Next = nil // Terminate the first half

	// Step 3: Merge the two halves
	mergeLists(head, secondHalf)
}

// ReorderListStack solves the same problem using stack approach
// Time complexity: O(n), Space complexity: O(n)
func ReorderListStack(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	// Step 1: Push all nodes to a stack
	stack := []*ListNode{}
	current := head
	for current != nil {
		stack = append(stack, current)
		current = current.Next
	}

	// Step 2: Reorder by alternating between head and stack
	current = head
	for i := len(stack) - 1; i > len(stack)/2; i-- {
		next := current.Next
		current.Next = stack[i]
		stack[i].Next = next
		current = next
	}

	// Step 3: Terminate the list properly
	if len(stack)%2 == 0 {
		current.Next.Next = nil
	} else {
		current.Next = nil
	}
}

// ReorderListArray solves the problem using array approach
// Time complexity: O(n), Space complexity: O(n)
func ReorderListArray(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	// Step 1: Store nodes in an array
	nodes := []*ListNode{}
	current := head
	for current != nil {
		nodes = append(nodes, current)
		current = current.Next
	}

	// Step 2: Reorder using two pointers
	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}

	// Step 3: Terminate the list
	nodes[i].Next = nil
}

// Helper function to reverse a linked list
func reverseListHelper(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

// Helper function to merge two lists in alternating fashion
func mergeLists(first, second *ListNode) {
	for second != nil {
		nextFirst := first.Next
		nextSecond := second.Next

		first.Next = second
		second.Next = nextFirst

		first = nextFirst
		second = nextSecond
	}
}