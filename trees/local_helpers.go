package trees

import "leetcode/utils"

// sortedArrayToBSTLocal is a local version to avoid circular imports
func sortedArrayToBSTLocal(nums []int) *utils.TreeNode {
	if len(nums) == 0 {
		return nil
	}
	
	mid := len(nums) / 2
	root := &utils.TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBSTLocal(nums[:mid])
	root.Right = sortedArrayToBSTLocal(nums[mid+1:])
	return root
}