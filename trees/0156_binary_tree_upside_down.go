package trees

import "leetcode/utils"

// 156. Binary Tree Upside Down
// https://leetcode.com/problems/binary-tree-upside-down/

// upsideDownBinaryTreeIterative transforms a binary tree upside down iteratively.
// The transformation rules:
// 1. The original left child becomes the new root
// 2. The original root becomes the new right child
// 3. The original right child becomes the new left child
// Time Complexity: O(n), Space Complexity: O(1)
func upsideDownBinaryTreeIterative(root *utils.TreeNode) *utils.TreeNode {
	if root == nil || root.Left == nil {
		return root
	}

	var parent, parentRight *utils.TreeNode
	current := root

	for current != nil && current.Left != nil {
		// Save the left child (will become new root)
		next := current.Left

		// Reassign pointers
		current.Left = parentRight
		parentRight = current.Right
		current.Right = parent

		// Move to next node
		parent = current
		current = next
	}

	// Handle the last node
	if current != nil {
		current.Left = parentRight
		current.Right = parent
	}

	return current
}

// upsideDownBinaryTreeRecursive transforms a binary tree upside down recursively.
// Time Complexity: O(n), Space Complexity: O(n) due to recursion stack
func upsideDownBinaryTreeRecursive(root *utils.TreeNode) *utils.TreeNode {
	if root == nil || root.Left == nil {
		return root
	}

	// Recursively transform the left subtree
	newRoot := upsideDownBinaryTreeRecursive(root.Left)

	// Reassign pointers according to transformation rules
	root.Left.Left = root.Right
	root.Left.Right = root

	// Clear the original root's children
	root.Left = nil
	root.Right = nil

	return newRoot
}

// upsideDownBinaryTree is the main function that uses iterative solution
func upsideDownBinaryTree(root *utils.TreeNode) *utils.TreeNode {
	return upsideDownBinaryTreeIterative(root)
}

// upsideDownBinaryTreeDFS is another recursive approach using DFS
func upsideDownBinaryTreeDFS(root *utils.TreeNode) *utils.TreeNode {
	return dfsUpsideDown(root, nil, nil)
}

func dfsUpsideDown(node, parent, rightSibling *utils.TreeNode) *utils.TreeNode {
	if node == nil {
		return parent
	}

	// Save the original left and right children
	originalLeft := node.Left
	originalRight := node.Right

	// Reassign pointers
	node.Left = rightSibling
	node.Right = parent

	// Recursively process the original left child
	return dfsUpsideDown(originalLeft, node, originalRight)
}

// upsideDownBinaryTreeStack uses an explicit stack for iterative solution
func upsideDownBinaryTreeStack(root *utils.TreeNode) *utils.TreeNode {
	if root == nil || root.Left == nil {
		return root
	}

	stack := []*utils.TreeNode{}
	current := root

	// Push all left nodes to stack
	for current != nil && current.Left != nil {
		stack = append(stack, current)
		current = current.Left
	}

	// The new root is the last left node (deepest left leaf)
	newRoot := current

	// Process nodes in reverse order
	for i := len(stack) - 1; i >= 0; i-- {
		node := stack[i]
		
		// Reassign pointers according to transformation rules
		current.Left = node.Right
		current.Right = node

		// Clear node's children
		node.Left = nil
		node.Right = nil

		// Move to next node
		current = node
	}

	return newRoot
}

// Helper function to create a binary tree from slice (for testing)
func createTreeFromSlice(values []interface{}) *utils.TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &utils.TreeNode{Val: values[0].(int)}
	queue := []*utils.TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(values) && values[i] != nil {
			node.Left = &utils.TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		// Right child
		if i < len(values) && values[i] != nil {
			node.Right = &utils.TreeNode{Val: values[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// Helper function to convert tree to slice (for testing)
func treeToSlice(root *utils.TreeNode) []interface{} {
	if root == nil {
		return []interface{}{}
	}

	result := []interface{}{}
	queue := []*utils.TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, nil)
		} else {
			result = append(result, node.Val)
			queue = append(queue, node.Left, node.Right)
		}
	}

	// Remove trailing nils
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	return result
}