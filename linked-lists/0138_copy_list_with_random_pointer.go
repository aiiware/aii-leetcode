package linkedlists

/*
# 0138 - Copy List with Random Pointer
## Problem Description
A linked list of length n is given such that each node contains an additional random pointer, which could point to any node in the list, or null.

Construct a deep copy of the list. The deep copy should consist of exactly n brand new nodes, where each new node has its value set to the value of its corresponding original node. Both the next and random pointer of the new nodes should point to new nodes in the copied list such that the pointers in the original list and copied list represent the same list state. None of the pointers in the new list should point to nodes in the original list.

Return the head of the copied linked list.

## Examples
Example 1:
Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]

Example 2:
Input: head = [[1,1],[2,1]]
Output: [[1,1],[2,1]]

Example 3:
Input: head = [[3,null],[3,0],[3,null]]
Output: [[3,null],[3,0],[3,null]]

## Constraints
- 0 <= n <= 1000
- -10^4 <= Node.val <= 10^4
- Node.random is null or points to a node in the linked list.

## Solution Approach
This problem can be solved using a hash map:
1. First pass: Create a copy of each node and store mapping from original to copy
2. Second pass: Set next and random pointers for copied nodes using the mapping

Alternative approach using interweaving (O(1) space):
1. First pass: Create copy of each node and insert it after the original
2. Second pass: Set random pointers for copied nodes
3. Third pass: Separate the original and copied lists

Time Complexity: O(N) where N is the number of nodes
Space Complexity: O(N) for hash map approach, O(1) for interweaving approach
*/

// Node represents a node in a linked list with a random pointer
type RandomListNode struct {
	Val    int
	Next   *RandomListNode
	Random *RandomListNode
}

// NewRandomListNode creates a new RandomListNode with the given value
func NewRandomListNode(val int) *RandomListNode {
	return &RandomListNode{Val: val}
}

// CopyRandomList returns a deep copy of the linked list with random pointers
func CopyRandomList(head *RandomListNode) *RandomListNode {
	if head == nil {
		return nil
	}
	
	// Map to store mapping from original nodes to copied nodes
	nodeMap := make(map[*RandomListNode]*RandomListNode)
	
	// First pass: create copies of all nodes
	current := head
	for current != nil {
		nodeMap[current] = NewRandomListNode(current.Val)
		current = current.Next
	}
	
	// Second pass: set next and random pointers
	current = head
	for current != nil {
		copiedNode := nodeMap[current]
		
		// Set next pointer
		if current.Next != nil {
			copiedNode.Next = nodeMap[current.Next]
		}
		
		// Set random pointer
		if current.Random != nil {
			copiedNode.Random = nodeMap[current.Random]
		}
		
		current = current.Next
	}
	
	return nodeMap[head]
}

// CopyRandomListOptimized returns a deep copy using O(1) space (interweaving)
func CopyRandomListOptimized(head *RandomListNode) *RandomListNode {
	if head == nil {
		return nil
	}
	
	// First pass: create copy of each node and insert it after the original
	current := head
	for current != nil {
		copyNode := NewRandomListNode(current.Val)
		copyNode.Next = current.Next
		current.Next = copyNode
		current = copyNode.Next
	}
	
	// Second pass: set random pointers for copied nodes
	current = head
	for current != nil {
		if current.Random != nil {
			current.Next.Random = current.Random.Next
		}
		current = current.Next.Next
	}
	
	// Third pass: separate the original and copied lists
	current = head
	copyHead := head.Next
	copyCurrent := copyHead
	
	for current != nil {
		current.Next = current.Next.Next
		if copyCurrent.Next != nil {
			copyCurrent.Next = copyCurrent.Next.Next
		}
		current = current.Next
		copyCurrent = copyCurrent.Next
	}
	
	return copyHead
}