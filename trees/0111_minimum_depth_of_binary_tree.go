package trees


/*
Difficulty: Easy
Tags: [Add relevant tags]
Companies: [Add company names]
*/

import "leetcode/utils"

// Problem 0111: Minimum Depth of Binary Tree
//
// Given a binary tree, find its minimum depth.
//
// The minimum depth is the number of nodes along the shortest path
// from the root node down to the nearest leaf node.
//
// Note: A leaf is a node with no children.
//
// Example 1:
// Input: root = [3,9,20,null,null,15,7]
// Output: 2
// Explanation:
//    3
//   / \
//  9  20
//    /  \
//   15   7
// The minimum depth is 2 (path: 3 -> 9).
//
// Example 2:
// Input: root = [2,null,3,null,4,null,5,null,6]
// Output: 5
// Explanation:
//   2
//    \
//     3
//      \
//       4
//        \
//         5
//          \
//           6
// The minimum depth is 5 (path: 2 -> 3 -> 4 -> 5 -> 6).
//
// Example 3:
// Input: root = []
// Output: 0
//
// Constraints:
// - The number of nodes in the tree is in the range [0, 10^5].
// - -1000 <= Node.val <= 1000

// minDepthRecursive is the standard recursive DFS solution.
// This approach recursively calculates the minimum depth by finding
// the minimum depth of left and right subtrees and adding 1 for the current node.
// Special case: if a node has only one child, we must follow that child.
// Time complexity: O(n), Space complexity: O(h) where h is the height of the tree
func minDepthRecursive(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	// If node is a leaf
	if root.Left == nil && root.Right == nil {
		return 1
	}

	// If left subtree is empty, we must go right
	if root.Left == nil {
		return minDepthRecursive(root.Right) + 1
	}

	// If right subtree is empty, we must go left
	if root.Right == nil {
		return minDepthRecursive(root.Left) + 1
	}

	// Both subtrees exist, take the minimum
	leftDepth := minDepthRecursive(root.Left)
	rightDepth := minDepthRecursive(root.Right)
	if leftDepth < rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

// minDepthBFS is the BFS (level order traversal) solution.
// This approach uses a queue to traverse the tree level by level,
// returning the depth when we encounter the first leaf node.
// This is more efficient for finding minimum depth as we stop early.
// Time complexity: O(n) worst case, but O(minDepth) on average
// Space complexity: O(n) worst case (complete tree)
func minDepthBFS(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 0
	queue := []*utils.TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		depth++

		// Process all nodes at current level
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// Check if this is a leaf node
			if node.Left == nil && node.Right == nil {
				return depth
			}

			// Add children to queue for next level
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return depth
}

// minDepthDFS is the iterative DFS solution using a stack.
// This approach uses a stack to simulate recursion, tracking both the node
// and its current depth. We keep track of the minimum depth found so far.
// Time complexity: O(n), Space complexity: O(h) where h is the height of the tree
func minDepthDFS(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	minDepth := int(^uint(0) >> 1) // Max int
	stack := []struct {
		node  *utils.TreeNode
		depth int
	}{{root, 1}}

	for len(stack) > 0 {
		// Pop from stack
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		node := item.node
		currentDepth := item.depth

		// If this is a leaf node, update minDepth
		if node.Left == nil && node.Right == nil {
			if currentDepth < minDepth {
				minDepth = currentDepth
			}
			continue
		}

		// Push children to stack
		if node.Right != nil {
			stack = append(stack, struct {
				node  *utils.TreeNode
				depth int
			}{node.Right, currentDepth + 1})
		}
		if node.Left != nil {
			stack = append(stack, struct {
				node  *utils.TreeNode
				depth int
			}{node.Left, currentDepth + 1})
		}
	}

	return minDepth
}

// minDepthOptimized is an optimized recursive solution.
// This version is more concise and handles edge cases elegantly.
func minDepthOptimized(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	left := minDepthOptimized(root.Left)
	right := minDepthOptimized(root.Right)

	// If one subtree is empty, we must use the other
	if left == 0 || right == 0 {
		return left + right + 1
	}

	// Both subtrees exist, take the minimum
	if left < right {
		return left + 1
	}
	return right + 1
}

// minDepth is the main function that delegates to the optimized solution.
// It uses the BFS approach as it's more efficient for finding minimum depth
// (stops early when first leaf is found).
func minDepth(root *utils.TreeNode) int {
	return minDepthBFS(root)
}