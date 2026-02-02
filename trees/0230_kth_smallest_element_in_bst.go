package trees

import "leetcode/utils"

// KthSmallest solves LeetCode problem 0230: Kth Smallest Element in a BST
// Difficulty: Medium
// Tags: Tree, Depth-First Search, Binary Search Tree, Binary Tree
//
// Given the root of a binary search tree, and an integer k, return the kth smallest
// value (1-indexed) of all the values of the nodes in the tree.
//
// Example 1:
// Input: root = [3,1,4,null,2], k = 1
// Output: 1
//
// Example 2:
// Input: root = [5,3,6,2,4,null,null,1], k = 3
// Output: 3
//
// Constraints:
// - The number of nodes in the tree is n.
// - 1 <= k <= n <= 10^4
// - 0 <= Node.val <= 10^4
//
// Follow up: If the BST is modified often (i.e., we can do insert and delete operations)
// and you need to find the kth smallest frequently, how would you optimize?
//
// Time complexity: O(n) worst case, O(k) average for early termination
// Space complexity: O(h) where h is the height of the tree
func KthSmallest(root *utils.TreeNode, k int) int {
	// Use iterative in-order traversal to find the kth smallest element
	stack := []*utils.TreeNode{}
	current := root
	count := 0

	for current != nil || len(stack) > 0 {
		// Go to the leftmost node
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		// Process node
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		count++

		// If this is the kth node, return its value
		if count == k {
			return current.Val
		}

		// Move to right subtree
		current = current.Right
	}

	// This should never be reached if k is valid
	return -1
}

// KthSmallestRecursive solves the same problem using recursive in-order traversal
func KthSmallestRecursive(root *utils.TreeNode, k int) int {
	result := 0
	count := 0

	var inorder func(*utils.TreeNode)
	inorder = func(node *utils.TreeNode) {
		if node == nil || count >= k {
			return
		}

		// Traverse left subtree
		inorder(node.Left)

		// Process current node
		count++
		if count == k {
			result = node.Val
			return
		}

		// Traverse right subtree
		inorder(node.Right)
	}

	inorder(root)
	return result
}

// KthSmallestMorris solves using Morris traversal (O(1) space, no recursion stack)
func KthSmallestMorris(root *utils.TreeNode, k int) int {
	current := root
	count := 0

	for current != nil {
		if current.Left == nil {
			// Process current node
			count++
			if count == k {
				return current.Val
			}
			current = current.Right
		} else {
			// Find inorder predecessor
			predecessor := current.Left
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
			}

			if predecessor.Right == nil {
				// Create temporary link to current
				predecessor.Right = current
				current = current.Left
			} else {
				// Restore tree structure and process current node
				predecessor.Right = nil
				count++
				if count == k {
					return current.Val
				}
				current = current.Right
			}
		}
	}

	return -1
}

// KthSmallestWithCounter solves using a counter struct for cleaner recursive solution
func KthSmallestWithCounter(root *utils.TreeNode, k int) int {
	counter := &counter{target: k}
	return counter.find(root)
}

type counter struct {
	count  int
	target int
	result int
	found  bool
}

func (c *counter) find(node *utils.TreeNode) int {
	if node == nil || c.found {
		return c.result
	}

	// Traverse left subtree
	c.find(node.Left)

	// Process current node
	if !c.found {
		c.count++
		if c.count == c.target {
			c.result = node.Val
			c.found = true
			return c.result
		}
	}

	// Traverse right subtree
	c.find(node.Right)

	return c.result
}