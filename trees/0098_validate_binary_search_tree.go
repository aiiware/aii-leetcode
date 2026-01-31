package trees


/*
Difficulty: Medium
Tags: [Add relevant tags]
Companies: [Add company names]
*/

import (
    "leetcode/utils"
    "math"
)

// Problem 0098: Validate Binary Search Tree
//
// Given the root of a binary tree, determine if it is a valid binary search tree (BST).
//
// A valid BST is defined as follows:
// - The left subtree of a node contains only nodes with keys less than the node's key.
// - The right subtree of a node contains only nodes with keys greater than the node's key.
// - Both the left and right subtrees must also be binary search trees.
//
// Example 1:
// Input: root = [2,1,3]
// Output: true
//
// Example 2:
// Input: root = [5,1,4,null,null,3,6]
// Output: false
// Explanation: The root node's value is 5 but its right child's value is 4.
//
// Constraints:
// - The number of nodes in the tree is in the range [1, 10^4].
// - -2^31 <= Node.val <= 2^31 - 1

// isValidBST is the main solution function using inorder traversal.
// Time complexity: O(n), Space complexity: O(n) worst case (skewed tree)
func isValidBST(root *utils.TreeNode) bool {
	// Use inorder traversal - BST inorder should be sorted
	var prev *int
	return isValidBSTInorder(root, &prev)
}

func isValidBSTInorder(node *utils.TreeNode, prev **int) bool {
	if node == nil {
		return true
	}
	
	// Check left subtree
	if !isValidBSTInorder(node.Left, prev) {
		return false
	}
	
	// Check current node
	if *prev != nil && node.Val <= **prev {
		return false
	}
	*prev = &node.Val
	
	// Check right subtree
	return isValidBSTInorder(node.Right, prev)
}

// isValidBSTRecursive uses recursion with min/max bounds.
func isValidBSTRecursive(root *utils.TreeNode) bool {
	return isValidBSTHelper(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTHelper(node *utils.TreeNode, minVal, maxVal int64) bool {
	if node == nil {
		return true
	}
	
	val := int64(node.Val)
	if val <= minVal || val >= maxVal {
		return false
	}
	
	return isValidBSTHelper(node.Left, minVal, val) &&
		isValidBSTHelper(node.Right, val, maxVal)
}

// isValidBSTIterative uses iterative inorder traversal with stack.
func isValidBSTIterative(root *utils.TreeNode) bool {
	if root == nil {
		return true
	}
	
	stack := []*utils.TreeNode{}
	current := root
	var prev *int = nil
	
	for current != nil || len(stack) > 0 {
		// Go to leftmost node
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}
		
		// Process node
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		// Check order
		if prev != nil && current.Val <= *prev {
			return false
		}
		prev = &current.Val
		
		// Go to right subtree
		current = current.Right
	}
	
	return true
}

// isValidBSTMorris uses Morris traversal (threaded binary tree).
// Time complexity: O(n), Space complexity: O(1)
func isValidBSTMorris(root *utils.TreeNode) bool {
	var prev *int = nil
	current := root
	
	for current != nil {
		if current.Left == nil {
			// Visit current node
			if prev != nil && current.Val <= *prev {
				return false
			}
			prev = &current.Val
			current = current.Right
		} else {
			// Find inorder predecessor
			predecessor := current.Left
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
			}
			
			if predecessor.Right == nil {
				// Create thread to current
				predecessor.Right = current
				current = current.Left
			} else {
				// Remove thread and visit current
				predecessor.Right = nil
				if prev != nil && current.Val <= *prev {
					return false
				}
				prev = &current.Val
				current = current.Right
			}
		}
	}
	
	return true
}

// isValidBSTDFS uses DFS with explicit stack.
func isValidBSTDFS(root *utils.TreeNode) bool {
	if root == nil {
		return true
	}
	
	type stackItem struct {
		node *utils.TreeNode
		min  int64
		max  int64
	}
	
	stack := []stackItem{{node: root, min: math.MinInt64, max: math.MaxInt64}}
	
	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		if item.node == nil {
			continue
		}
		
		val := int64(item.node.Val)
		if val <= item.min || val >= item.max {
			return false
		}
		
		// Push children
		stack = append(stack, stackItem{
			node: item.node.Right,
			min:  val,
			max:  item.max,
		})
		stack = append(stack, stackItem{
			node: item.node.Left,
			min:  item.min,
			max:  val,
		})
	}
	
	return true
}

// isValidBSTBFS uses BFS (level order) with bounds.
func isValidBSTBFS(root *utils.TreeNode) bool {
	if root == nil {
		return true
	}
	
	type queueItem struct {
		node *utils.TreeNode
		min  int64
		max  int64
	}
	
	queue := []queueItem{{node: root, min: math.MinInt64, max: math.MaxInt64}}
	
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		
		val := int64(item.node.Val)
		if val <= item.min || val >= item.max {
			return false
		}
		
		// Add children to queue
		if item.node.Left != nil {
			queue = append(queue, queueItem{
				node: item.node.Left,
				min:  item.min,
				max:  val,
			})
		}
		if item.node.Right != nil {
			queue = append(queue, queueItem{
				node: item.node.Right,
				min:  val,
				max:  item.max,
			})
		}
	}
	
	return true
}

// isValidBSTSimple uses simple recursive approach.
func isValidBSTSimple(root *utils.TreeNode) bool {
	// Get inorder traversal
	values := inorderTraversalForBST(root)
	
	// Check if sorted
	for i := 1; i < len(values); i++ {
		if values[i] <= values[i-1] {
			return false
		}
	}
	
	return true
}

// inorderTraversalForBST returns inorder traversal of tree (internal helper)
func inorderTraversalForBST(root *utils.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	
	left := inorderTraversalForBST(root.Left)
	right := inorderTraversalForBST(root.Right)
	
	result := append(left, root.Val)
	result = append(result, right...)
	return result
}

// isValidBSTOptimized is an optimized version.
func isValidBSTOptimized(root *utils.TreeNode) bool {
	var prev *int = nil
	var dfs func(*utils.TreeNode) bool
	
	dfs = func(node *utils.TreeNode) bool {
		if node == nil {
			return true
		}
		
		// Check left subtree
		if !dfs(node.Left) {
			return false
		}
		
		// Check current node
		if prev != nil && node.Val <= *prev {
			return false
		}
		prev = &node.Val
		
		// Check right subtree
		return dfs(node.Right)
	}
	
	return dfs(root)
}

// IsValidBST is the public interface function.
// It uses the optimized recursive solution by default.
func IsValidBST(root *utils.TreeNode) bool {
	return isValidBSTOptimized(root)
}