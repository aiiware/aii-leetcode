package leetcode

// Problem 0104: Maximum Depth of Binary Tree
//
// Given the root of a binary tree, return its maximum depth.
//
// A binary tree's maximum depth is the number of nodes along the longest path
// from the root node down to the farthest leaf node.
//
// Example 1:
// Input: root = [3,9,20,null,null,15,7]
// Output: 3
// Explanation:
//    3
//   / \
//  9  20
//    /  \
//   15   7
//
// Example 2:
// Input: root = [1,null,2]
// Output: 2
//
// Example 3:
// Input: root = []
// Output: 0
//
// Constraints:
// - The number of nodes in the tree is in the range [0, 10^4].
// - -100 <= Node.val <= 100

// maxDepthRecursive is the standard recursive DFS solution.
// This approach recursively calculates the maximum depth by finding
// the maximum depth of left and right subtrees and adding 1 for the current node.
// Time complexity: O(n), Space complexity: O(h) where h is the height of the tree
func maxDepthRecursive(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepthRecursive(root.Left)
	rightDepth := maxDepthRecursive(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

// maxDepthBFS is the BFS (level order traversal) solution.
// This approach uses a queue to traverse the tree level by level,
// counting the number of levels.
// Time complexity: O(n), Space complexity: O(n) worst case (complete tree)
func maxDepthBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 0
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		depth++

		// Process all nodes at current level
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

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

// maxDepthDFS is the iterative DFS solution using a stack.
// This approach uses a stack to simulate recursion, tracking both the node
// and its current depth.
// Time complexity: O(n), Space complexity: O(h) where h is the height of the tree
func maxDepthDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxDepth := 0
	stack := []struct {
		node  *TreeNode
		depth int
	}{{root, 1}}

	for len(stack) > 0 {
		// Pop from stack
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		node := item.node
		currentDepth := item.depth

		// Update max depth
		if currentDepth > maxDepth {
			maxDepth = currentDepth
		}

		// Push children to stack
		if node.Right != nil {
			stack = append(stack, struct {
				node  *TreeNode
				depth int
			}{node.Right, currentDepth + 1})
		}
		if node.Left != nil {
			stack = append(stack, struct {
				node  *TreeNode
				depth int
			}{node.Left, currentDepth + 1})
		}
	}

	return maxDepth
}

// maxDepth is the main function that delegates to the optimized solution.
// It uses the recursive approach as it's the most intuitive and efficient
// for this problem.
func maxDepth(root *TreeNode) int {
	return maxDepthRecursive(root)
}