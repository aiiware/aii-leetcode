package leetcode

// Problem 0112: Path Sum
//
// Given the root of a binary tree and an integer targetSum,
// return true if the tree has a root-to-leaf path such that
// adding up all the values along the path equals targetSum.
//
// A leaf is a node with no children.
//
// Example 1:
// Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
// Output: true
// Explanation:
//       5
//      / \
//     4   8
//    /   / \
//   11  13  4
//  /  \      \
// 7    2      1
// The root-to-leaf path 5->4->11->2 sums to 22.
//
// Example 2:
// Input: root = [1,2,3], targetSum = 5
// Output: false
// Explanation:
//   1
//  / \
// 2   3
// There is no root-to-leaf path with sum = 5.
//
// Example 3:
// Input: root = [], targetSum = 0
// Output: false
// Explanation: Since the tree is empty, there are no root-to-leaf paths.
//
// Constraints:
// - The number of nodes in the tree is in the range [0, 5000].
// - -1000 <= Node.val <= 1000
// - -1000 <= targetSum <= 1000

// hasPathSumRecursive is the standard recursive DFS solution.
// This approach recursively checks if there's a path from current node to leaf
// that sums to the remaining target.
// Time complexity: O(n), Space complexity: O(h) where h is the height of the tree
func hasPathSumRecursive(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	// Check if current node is a leaf and matches target
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	// Recursively check left and right subtrees with reduced target
	remaining := targetSum - root.Val
	return hasPathSumRecursive(root.Left, remaining) || hasPathSumRecursive(root.Right, remaining)
}

// hasPathSumDFS is the iterative DFS solution using a stack.
// This approach uses a stack to simulate recursion, tracking both the node
// and the remaining sum needed from that node to the leaf.
// Time complexity: O(n), Space complexity: O(h) where h is the height of the tree
func hasPathSumDFS(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	stack := []struct {
		node   *TreeNode
		remain int
	}{{root, targetSum}}

	for len(stack) > 0 {
		// Pop from stack
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		node := item.node
		remain := item.remain

		// Check if current node is a leaf and matches remaining sum
		if node.Left == nil && node.Right == nil && node.Val == remain {
			return true
		}

		// Push children to stack with updated remaining sum
		if node.Right != nil {
			stack = append(stack, struct {
				node   *TreeNode
				remain int
			}{node.Right, remain - node.Val})
		}
		if node.Left != nil {
			stack = append(stack, struct {
				node   *TreeNode
				remain int
			}{node.Left, remain - node.Val})
		}
	}

	return false
}

// hasPathSumBFS is the BFS solution using a queue.
// This approach uses two queues: one for nodes and one for remaining sums.
// Time complexity: O(n), Space complexity: O(n) worst case (complete tree)
func hasPathSumBFS(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	nodeQueue := []*TreeNode{root}
	sumQueue := []int{targetSum}

	for len(nodeQueue) > 0 {
		node := nodeQueue[0]
		nodeQueue = nodeQueue[1:]

		remain := sumQueue[0]
		sumQueue = sumQueue[1:]

		// Check if current node is a leaf and matches remaining sum
		if node.Left == nil && node.Right == nil && node.Val == remain {
			return true
		}

		// Add children to queues with updated remaining sums
		if node.Left != nil {
			nodeQueue = append(nodeQueue, node.Left)
			sumQueue = append(sumQueue, remain-node.Val)
		}
		if node.Right != nil {
			nodeQueue = append(nodeQueue, node.Right)
			sumQueue = append(sumQueue, remain-node.Val)
		}
	}

	return false
}

// hasPathSumOptimized is an optimized recursive solution with early termination.
// This version uses a helper function that tracks the current sum.
func hasPathSumOptimized(root *TreeNode, targetSum int) bool {
	return hasPathSumHelper(root, 0, targetSum)
}

func hasPathSumHelper(node *TreeNode, currentSum, targetSum int) bool {
	if node == nil {
		return false
	}

	currentSum += node.Val

	// Check if current node is a leaf and sum matches target
	if node.Left == nil && node.Right == nil {
		return currentSum == targetSum
	}

	// Recursively check left and right subtrees
	return hasPathSumHelper(node.Left, currentSum, targetSum) ||
		hasPathSumHelper(node.Right, currentSum, targetSum)
}

// hasPathSum is the main function that delegates to the optimized solution.
// It uses the recursive approach as it's the most intuitive and efficient.
func hasPathSum(root *TreeNode, targetSum int) bool {
	return hasPathSumRecursive(root, targetSum)
}