package leetcode

// MaxSubArray solves LeetCode problem 0053: Maximum Subarray
// Difficulty: Easy
// Tags: Array, Divide and Conquer, Dynamic Programming
//
// Given an integer array nums, find the subarray with the largest sum,
// and return its sum.
//
// Example 1:
// Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
// Output: 6
// Explanation: The subarray [4,-1,2,1] has the largest sum = 6.
//
// Example 2:
// Input: nums = [1]
// Output: 1
//
// Example 3:
// Input: nums = [5,4,-1,7,8]
// Output: 23
//
// Time complexity: O(n), Space complexity: O(1)
func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Kadane's algorithm
	maxCurrent := nums[0]
	maxGlobal := nums[0]

	for i := 1; i < len(nums); i++ {
		// Either extend the current subarray or start a new one
		maxCurrent = max(nums[i], maxCurrent+nums[i])

		// Update global maximum if needed
		if maxCurrent > maxGlobal {
			maxGlobal = maxCurrent
		}
	}

	return maxGlobal
}