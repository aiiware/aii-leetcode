package dp

// WiggleMaxLength solves LeetCode problem 0376: Wiggle Subsequence
// Difficulty: Medium
// Tags: Array, Dynamic Programming
//
// A wiggle sequence is a sequence where the differences between successive numbers 
// strictly alternate between positive and negative. The first difference (if one exists) 
// may be either positive or negative. A sequence with one element and a sequence with 
// two non-equal elements are trivially wiggle sequences.
//
// Given an integer array nums, return the length of the longest wiggle subsequence.
//
// Time complexity: O(n), Space complexity: O(1)
func WiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	
	// up represents the length of wiggle subsequence ending at current position with a rising trend
	// down represents the length of wiggle subsequence ending at current position with a falling trend
	up := 1
	down := 1
	
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			// Rising trend - extend the down sequence
			up = down + 1
		} else if nums[i] < nums[i-1] {
			// Falling trend - extend the up sequence
			down = up + 1
		}
		// If nums[i] == nums[i-1], we don't change either up or down
	}
	
	return max(up, down)
}