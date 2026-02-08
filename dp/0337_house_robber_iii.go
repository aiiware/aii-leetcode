package dp

// RobTree solves LeetCode problem 0337: House Robber III
// Difficulty: Medium
// Tags: Tree, Dynamic Programming
//
// The thief has found himself a new place for his thievery again. There is only one entrance to this area, called root.
// Besides the root, each house has one and only one parent house. After a tour, the smart thief realized that all houses
// form a binary tree. It will automatically contact the police if two directly-linked houses were broken into on the same night.
// Given the root of the binary tree, return the maximum amount of money the thief can rob without alerting the police.
//
// Time complexity: O(n), Space complexity: O(h) where h is the height of the tree
func RobTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// Return the maximum of robbing or not robbing the root
	rob, notRob := robTreeHelper(root)
	return max(rob, notRob)
}

// robTreeHelper returns (max money if we rob current node, max money if we don't rob current node)
func robTreeHelper(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}

	// Recursively get results for left and right subtrees
	leftRob, leftNotRob := robTreeHelper(node.Left)
	rightRob, rightNotRob := robTreeHelper(node.Right)

	// If we rob current node, we can't rob children
	robCurrent := node.Val + leftNotRob + rightNotRob

	// If we don't rob current node, we can take max from both children
	// (whether we rob or not rob the children)
	notRobCurrent := max(leftRob, leftNotRob) + max(rightRob, rightNotRob)

	return robCurrent, notRobCurrent
}

// TreeNode represents a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
