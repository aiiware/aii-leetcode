package trees


/*
Difficulty: Easy
Tags: [Add relevant tags]
Companies: [Add company names]
*/

import "leetcode/utils"

// Problem 0094: Binary Tree Inorder Traversal
//
// Given the root of a binary tree, return the inorder traversal of its nodes' values.
//
// Example 1:
// Input: root = [1,null,2,3]
// Output: [1,3,2]
//
// Example 2:
// Input: root = []
// Output: []
//
// Example 3:
// Input: root = [1]
// Output: [1]
//
// Constraints:
// - The number of nodes in the tree is in the range [0, 100].
// - -100 <= Node.val <= 100

// inorderTraversal is the main solution function using recursion.
// Time complexity: O(n), Space complexity: O(n) worst case (skewed tree)
func inorderTraversal(root *utils.TreeNode) []int {
	result := []int{}
	inorderRecursive(root, &result)
	return result
}

func inorderRecursive(node *utils.TreeNode, result *[]int) {
	if node == nil {
		return
	}
	inorderRecursive(node.Left, result)
	*result = append(*result, node.Val)
	inorderRecursive(node.Right, result)
}

// inorderTraversalIterative uses iterative approach with stack.
func inorderTraversalIterative(root *utils.TreeNode) []int {
	result := []int{}
	stack := []*utils.TreeNode{}
	current := root

	for current != nil || len(stack) > 0 {
		// Go to the leftmost node
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		// Process the node
		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, current.Val)

		// Go to right subtree
		current = current.Right
	}

	return result
}

// inorderTraversalMorris uses Morris traversal (threaded binary tree).
// Time complexity: O(n), Space complexity: O(1)
func inorderTraversalMorris(root *utils.TreeNode) []int {
	result := []int{}
	current := root

	for current != nil {
		if current.Left == nil {
			// If no left child, visit current node and go to right
			result = append(result, current.Val)
			current = current.Right
		} else {
			// Find the inorder predecessor
			predecessor := current.Left
			for predecessor.Right != nil && predecessor.Right != current {
				predecessor = predecessor.Right
			}

			if predecessor.Right == nil {
				// Create a thread to current node
				predecessor.Right = current
				current = current.Left
			} else {
				// Remove the thread and visit current node
				predecessor.Right = nil
				result = append(result, current.Val)
				current = current.Right
			}
		}
	}

	return result
}

// inorderTraversalDFS uses DFS with explicit stack.
func inorderTraversalDFS(root *utils.TreeNode) []int {
	if root == nil {
		return []int{}
	}

	type stackItem struct {
		node *utils.TreeNode
		visited bool
	}
	
	stack := []stackItem{{node: root, visited: false}}
	result := []int{}
	
	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		
		if item.node == nil {
			continue
		}
		
		if item.visited {
			result = append(result, item.node.Val)
		} else {
			// Push in reverse order: right, node (visited), left
			stack = append(stack, stackItem{node: item.node.Right, visited: false})
			stack = append(stack, stackItem{node: item.node, visited: true})
			stack = append(stack, stackItem{node: item.node.Left, visited: false})
		}
	}
	
	return result
}

// inorderTraversalBFS is not applicable for inorder (BFS is level order).
// This is included for completeness but uses different approach.
func inorderTraversalBFS(root *utils.TreeNode) []int {
	// BFS doesn't give inorder traversal, but we can simulate with state
	if root == nil {
		return []int{}
	}

	type state struct {
		node *utils.TreeNode
		visited bool
	}
	
	queue := []state{{node: root, visited: false}}
	result := []int{}
	
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		
		if current.node == nil {
			continue
		}
		
		if current.visited {
			result = append(result, current.node.Val)
		} else {
			// For BFS-like but inorder, we need to process differently
			// This is not a true BFS for inorder
			queue = append([]state{
				{node: current.node.Right, visited: false},
				{node: current.node, visited: true},
				{node: current.node.Left, visited: false},
			}, queue...)
		}
	}
	
	return result
}

// inorderTraversalSimple is a simple recursive implementation.
func inorderTraversalSimple(root *utils.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	
	left := inorderTraversalSimple(root.Left)
	right := inorderTraversalSimple(root.Right)
	
	result := append(left, root.Val)
	result = append(result, right...)
	return result
}

// inorderTraversalOptimized is an optimized iterative version.
func inorderTraversalOptimized(root *utils.TreeNode) []int {
	result := []int{}
	stack := []*utils.TreeNode{}
	
	for root != nil || len(stack) > 0 {
		// Push all left nodes
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		
		// Pop and process
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, root.Val)
		
		// Move to right
		root = root.Right
	}
	
	return result
}

// InorderTraversal is the public interface function.
// It uses the optimized iterative solution by default.
func InorderTraversal(root *utils.TreeNode) []int {
	return inorderTraversalOptimized(root)
}