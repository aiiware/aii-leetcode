package arrays

import "leetcode/utils"

// SortedArrayToBST is an exported wrapper for sortedArrayToBST
// Converts a sorted array to a height-balanced binary search tree
func SortedArrayToBST(nums []int) *utils.TreeNode {
	return sortedArrayToBST(nums)
}