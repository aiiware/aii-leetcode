package trees

import "leetcode/utils"

// IsBalancedBottomUp is the exported wrapper for isBalancedBottomUp
func IsBalancedBottomUp(root *utils.TreeNode) bool {
	return isBalancedBottomUp(root)
}

// IsBalanced is the exported wrapper for isBalanced
func IsBalanced(root *utils.TreeNode) bool {
	return isBalanced(root)
}