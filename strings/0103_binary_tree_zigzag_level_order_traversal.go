package strings

import "leetcode/utils"

// Problem 0103: Binary Tree Zigzag Level Order Traversal
//
// Given the root of a binary tree, return the zigzag level order traversal of its nodes' values.
// (i.e., from left to right, then right to left for the next level and alternate between).
//
// Example 1:
// Input: root = [3,9,20,null,null,15,7]
// Output: [[3],[20,9],[15,7]]
// Explanation:
// Level 0: [3] (left to right)
// Level 1: [20,9] (right to left)
// Level 2: [15,7] (left to right)
//
// Example 2:
// Input: root = [1]
// Output: [[1]]
//
// Example 3:
// Input: root = []
// Output: []
//
// Constraints:
// - The number of nodes in the tree is in the range [0, 2000].
// - -1000 <= Node.val <= 1000

// zigzagLevelOrderBFS is the standard BFS solution using a queue with direction flag.
// This approach performs level order traversal and reverses the order of nodes
// at alternating levels based on a direction flag.
// Time complexity: O(n), Space complexity: O(n) worst case (complete tree)
func zigzagLevelOrderBFS(root *utils.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*utils.TreeNode{root}
	leftToRight := true // Direction flag: true for left->right, false for right->left

	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, 0, levelSize)

		// Process all nodes at current level
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// Add node value to level slice
			level = append(level, node.Val)

			// Add children to queue for next level
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// Reverse the level if we're going right to left
		if !leftToRight {
			reverseSlice(level)
		}

		result = append(result, level)
		leftToRight = !leftToRight // Toggle direction for next level
	}

	return result
}

// zigzagLevelOrderTwoStacks uses two stacks to alternate traversal direction.
// This approach uses two stacks to naturally handle the zigzag pattern:
// - When processing left->right, push children in left then right order
// - When processing right->left, push children in right then left order
// Time complexity: O(n), Space complexity: O(n)
func zigzagLevelOrderTwoStacks(root *utils.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	currentLevel := []*utils.TreeNode{root} // Use as stack
	leftToRight := true

	for len(currentLevel) > 0 {
		levelSize := len(currentLevel)
		level := make([]int, 0, levelSize)
		nextLevel := []*utils.TreeNode{} // Use as stack for next level

		// Process current level
		for i := 0; i < levelSize; i++ {
			// Pop from stack (process in reverse order of push)
			node := currentLevel[len(currentLevel)-1]
			currentLevel = currentLevel[:len(currentLevel)-1]

			level = append(level, node.Val)

			if leftToRight {
				// For left->right direction, push left then right
				// This ensures right is processed before left in next iteration
				if node.Left != nil {
					nextLevel = append(nextLevel, node.Left)
				}
				if node.Right != nil {
					nextLevel = append(nextLevel, node.Right)
				}
			} else {
				// For right->left direction, push right then left
				// This ensures left is processed before right in next iteration
				if node.Right != nil {
					nextLevel = append(nextLevel, node.Right)
				}
				if node.Left != nil {
					nextLevel = append(nextLevel, node.Left)
				}
			}
		}

		result = append(result, level)
		currentLevel = nextLevel
		leftToRight = !leftToRight
	}

	return result
}

// zigzagLevelOrderDFS uses DFS with level tracking and direction-based insertion.
// This approach uses recursion to traverse the tree while tracking the current level.
// Values are inserted at the beginning or end of the level slice based on direction.
// Time complexity: O(n), Space complexity: O(n) worst case (skewed tree)
func zigzagLevelOrderDFS(root *utils.TreeNode) [][]int {
	result := [][]int{}
	dfsZigzag(root, 0, &result)
	return result
}

// dfsZigzag helper function for DFS solution
func dfsZigzag(node *utils.TreeNode, level int, result *[][]int) {
	if node == nil {
		return
	}

	// If we need a new level, add it
	if level >= len(*result) {
		*result = append(*result, []int{})
	}

	// Add current node value to its level based on direction
	if level%2 == 0 {
		// Even level (0-indexed): left to right, append to end
		(*result)[level] = append((*result)[level], node.Val)
	} else {
		// Odd level: right to left, prepend to beginning
		(*result)[level] = append([]int{node.Val}, (*result)[level]...)
	}

	// Recursively process children at next level
	dfsZigzag(node.Left, level+1, result)
	dfsZigzag(node.Right, level+1, result)
}

// zigzagLevelOrderDeque uses a deque (double-ended queue) for efficient operations.
// This approach uses a deque to efficiently add/remove from both ends.
// Time complexity: O(n), Space complexity: O(n)
func zigzagLevelOrderDeque(root *utils.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	deque := []*utils.TreeNode{root}
	leftToRight := true

	for len(deque) > 0 {
		levelSize := len(deque)
		level := make([]int, 0, levelSize)

		if leftToRight {
			// Process from front to back
			for i := 0; i < levelSize; i++ {
				node := deque[0]
				deque = deque[1:]
				level = append(level, node.Val)

				// Add children to back for next level
				if node.Left != nil {
					deque = append(deque, node.Left)
				}
				if node.Right != nil {
					deque = append(deque, node.Right)
				}
			}
		} else {
			// Process from back to front
			for i := 0; i < levelSize; i++ {
				node := deque[len(deque)-1]
				deque = deque[:len(deque)-1]
				level = append(level, node.Val)

				// Add children to front in reverse order for next level
				// Note: We need to add right then left when processing from back
				if node.Right != nil {
					deque = append([]*utils.TreeNode{node.Right}, deque...)
				}
				if node.Left != nil {
					deque = append([]*utils.TreeNode{node.Left}, deque...)
				}
			}
		}

		result = append(result, level)
		leftToRight = !leftToRight
	}

	return result
}

// zigzagLevelOrderOptimized is an optimized BFS solution with pre-allocation.
// This version pre-allocates slices and uses index-based reversal for better performance.
// Time complexity: O(n), Space complexity: O(n)
func zigzagLevelOrderOptimized(root *utils.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)
	queue := []*utils.TreeNode{root}
	leftToRight := true

	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, levelSize)

		// Process all nodes at current level
		for i := 0; i < levelSize; i++ {
			node := queue[i] // Direct access without shifting

			// Store value based on direction
			if leftToRight {
				level[i] = node.Val
			} else {
				level[levelSize-1-i] = node.Val
			}

			// Add children to queue
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// Remove processed nodes from queue
		queue = queue[levelSize:]
		result = append(result, level)
		leftToRight = !leftToRight
	}

	return result
}

// zigzagLevelOrder is the main entry point that uses the BFS solution.
// This follows LeetCode's expected function signature.
func zigzagLevelOrder(root *utils.TreeNode) [][]int {
	return zigzagLevelOrderBFS(root)
}

// reverseSlice reverses a slice of integers in place
func reverseSlice(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}