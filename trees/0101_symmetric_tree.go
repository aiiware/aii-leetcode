package trees

import "leetcode/utils"

// Problem 0101: Symmetric Tree
//
// Given the root of a binary tree, check whether it is a mirror of itself
// (i.e., symmetric around its center).
//
// Example 1:
// Input: root = [1,2,2,3,4,4,3]
// Output: true
//
// Example 2:
// Input: root = [1,2,2,null,3,null,3]
// Output: false
//
// Constraints:
// - The number of nodes in the tree is in the range [1, 1000].
// - -100 <= Node.val <= 100

// isSymmetricRecursive uses recursion to check symmetry.
// Time complexity: O(n), Space complexity: O(n) worst case (skewed tree)
func isSymmetricRecursive(root *utils.TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

// isMirror helper function for recursive solution
func isMirror(left, right *utils.TreeNode) bool {
	// Both nil
	if left == nil && right == nil {
		return true
	}
	// One nil
	if left == nil || right == nil {
		return false
	}
	// Values don't match
	if left.Val != right.Val {
		return false
	}
	// Check outer and inner pairs
	return isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
}

// isSymmetricIterative uses iterative approach with queue.
func isSymmetricIterative(root *utils.TreeNode) bool {
	if root == nil {
		return true
	}

	// Use a queue to process nodes in pairs
	queue := []*utils.TreeNode{root.Left, root.Right}

	for len(queue) > 0 {
		// Dequeue two nodes
		left := queue[0]
		right := queue[1]
		queue = queue[2:]

		// Both nil
		if left == nil && right == nil {
			continue
		}

		// One nil
		if left == nil || right == nil {
			return false
		}

		// Values don't match
		if left.Val != right.Val {
			return false
		}

		// Enqueue children in mirror order
		queue = append(queue, left.Left, right.Right)  // Outer pair
		queue = append(queue, left.Right, right.Left)  // Inner pair
	}

	return true
}

// isSymmetricStack uses stack for iterative solution.
func isSymmetricStack(root *utils.TreeNode) bool {
	if root == nil {
		return true
	}

	stack := []*utils.TreeNode{root.Left, root.Right}

	for len(stack) > 0 {
		// Pop two nodes
		right := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		left := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Both nil
		if left == nil && right == nil {
			continue
		}

		// One nil
		if left == nil || right == nil {
			return false
		}

		// Values don't match
		if left.Val != right.Val {
			return false
		}

		// Push children in mirror order
		stack = append(stack, left.Left, right.Right)  // Outer pair
		stack = append(stack, left.Right, right.Left)  // Inner pair
	}

	return true
}

// isSymmetricLevelOrder uses level order traversal.
func isSymmetricLevelOrder(root *utils.TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*utils.TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		levelValues := make([]*int, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// Store value (or nil)
			if node != nil {
				val := node.Val
				levelValues[i] = &val
				// Enqueue children
				queue = append(queue, node.Left, node.Right)
			} else {
				levelValues[i] = nil
				// Don't enqueue children for nil nodes
			}
		}

		// Check if level is symmetric
		if !isLevelSymmetric(levelValues) {
			return false
		}
	}

	return true
}

// isLevelSymmetric checks if a level is symmetric
func isLevelSymmetric(level []*int) bool {
	left, right := 0, len(level)-1
	for left < right {
		// Both nil
		if level[left] == nil && level[right] == nil {
			left++
			right--
			continue
		}
		// One nil
		if level[left] == nil || level[right] == nil {
			return false
		}
		// Values don't match
		if *level[left] != *level[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// isSymmetricDFS uses DFS with explicit stack.
func isSymmetricDFS(root *utils.TreeNode) bool {
	if root == nil {
		return true
	}

	type pair struct {
		left  *utils.TreeNode
		right *utils.TreeNode
	}

	stack := []pair{{root.Left, root.Right}}

	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Both nil
		if p.left == nil && p.right == nil {
			continue
		}

		// One nil
		if p.left == nil || p.right == nil {
			return false
		}

		// Values don't match
		if p.left.Val != p.right.Val {
			return false
		}

		// Push children in mirror order
		stack = append(stack, pair{p.left.Left, p.right.Right})  // Outer
		stack = append(stack, pair{p.left.Right, p.right.Left})  // Inner
	}

	return true
}

// isSymmetricOptimized is an optimized iterative solution.
func isSymmetricOptimized(root *utils.TreeNode) bool {
	if root == nil {
		return true
	}

	// Use two stacks for left and right subtrees
	leftStack := []*utils.TreeNode{root.Left}
	rightStack := []*utils.TreeNode{root.Right}

	for len(leftStack) > 0 && len(rightStack) > 0 {
		left := leftStack[len(leftStack)-1]
		leftStack = leftStack[:len(leftStack)-1]

		right := rightStack[len(rightStack)-1]
		rightStack = rightStack[:len(rightStack)-1]

		// Both nil
		if left == nil && right == nil {
			continue
		}

		// One nil
		if left == nil || right == nil {
			return false
		}

		// Values don't match
		if left.Val != right.Val {
			return false
		}

		// Push children in mirror order
		leftStack = append(leftStack, left.Left, left.Right)
		rightStack = append(rightStack, right.Right, right.Left)
	}

	return len(leftStack) == 0 && len(rightStack) == 0
}

// IsSymmetric is the public interface function.
// It uses the optimized iterative solution by default.
func IsSymmetric(root *utils.TreeNode) bool {
	return isSymmetricOptimized(root)
}