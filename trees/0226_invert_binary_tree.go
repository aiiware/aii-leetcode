package trees

import "leetcode/utils"

// InvertTree solves LeetCode problem 0226: Invert Binary Tree
// Difficulty: Easy
// Tags: Tree, Depth-First Search, Breadth-First Search, Binary Tree
//
// Given the root of a binary tree, invert the tree, and return its root.
//
// Example 1:
// Input: root = [4,2,7,1,3,6,9]
// Output: [4,7,2,9,6,3,1]
//
// Example 2:
// Input: root = [2,1,3]
// Output: [2,3,1]
//
// Example 3:
// Input: root = []
// Output: []
//
// Constraints:
// - The number of nodes in the tree is in the range [0, 100].
// - -100 <= Node.val <= 100
//
// Time complexity: O(n), Space complexity: O(n) for recursion stack, O(1) for iterative
func InvertTree(root *utils.TreeNode) *utils.TreeNode {
	// Recursive approach (most elegant)
	if root == nil {
		return nil
	}

	// Swap left and right subtrees
	root.Left, root.Right = root.Right, root.Left

	// Recursively invert subtrees
	InvertTree(root.Left)
	InvertTree(root.Right)

	return root
}

// InvertTreeIterative solves the same problem using iterative approach (BFS)
func InvertTreeIterative(root *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return nil
	}

	// Use a queue for level-order traversal
	queue := []*utils.TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// Swap left and right children
		node.Left, node.Right = node.Right, node.Left

		// Add children to queue
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return root
}

// InvertTreeDFS solves using iterative DFS with stack
func InvertTreeDFS(root *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return nil
	}

	stack := []*utils.TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Swap left and right children
		node.Left, node.Right = node.Right, node.Left

		// Push children to stack (right then left for DFS)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return root
}

// InvertTreeFunctional returns a new tree without modifying the original
func InvertTreeFunctional(root *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return nil
	}

	// Create a new tree with swapped children
	return &utils.TreeNode{
		Val:   root.Val,
		Left:  InvertTreeFunctional(root.Right),  // Right becomes left
		Right: InvertTreeFunctional(root.Left),   // Left becomes right
	}
}