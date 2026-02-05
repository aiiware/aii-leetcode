package dp

import (
	"leetcode/utils"
)

// Rob solves LeetCode problem 0198: House Robber
// Difficulty: Medium
// Tags: Dynamic Programming, Array
//
// You are a professional robber planning to rob houses along a street.
// Each house has a certain amount of money stashed, the only constraint
// stopping you from robbing each of them is that adjacent houses have
// security systems connected and it will automatically contact the police
// if two adjacent houses were broken into on the same night.
//
// Given an integer array nums representing the amount of money of each house,
// return the maximum amount of money you can rob tonight without alerting the police.
//
// Example 1:
// Input: nums = [1,2,3,1]
// Output: 4
// Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
// Total amount you can rob = 1 + 3 = 4.
//
// Example 2:
// Input: nums = [2,7,9,3,1]
// Output: 12
// Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
// Total amount you can rob = 2 + 9 + 1 = 12.
//
// Constraints:
// - 1 <= nums.length <= 100
// - 0 <= nums[i] <= 400
//
// Time complexity: O(n), Space complexity: O(1)
func Rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	// dp[i] represents the maximum amount that can be robbed from first i houses
	// We only need to keep track of two previous values
	prev2 := 0 // dp[i-2]
	prev1 := 0 // dp[i-1]

	for i := 0; i < len(nums); i++ {
		// At each house, we have two choices:
		// 1. Rob this house: nums[i] + prev2 (can't rob previous house)
		// 2. Skip this house: prev1 (carry forward previous maximum)
		current := utils.Max(prev1, nums[i]+prev2)
		prev2 = prev1
		prev1 = current
	}

	return prev1
}

// RobDPArray is an alternative implementation using DP array for clarity
func RobDPArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	// dp[i] represents the maximum amount that can be robbed from first i+1 houses
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = utils.Max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		// Two choices: rob current house or skip it
		dp[i] = utils.Max(dp[i-1], nums[i]+dp[i-2])
	}

	return dp[len(nums)-1]
}

// RobRecursive is a recursive solution with memoization
func RobRecursive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	memo := make([]int, len(nums))
	for i := range memo {
		memo[i] = -1
	}
	return robHelper(nums, 0, memo)
}

func robHelper(nums []int, i int, memo []int) int {
	if i >= len(nums) {
		return 0
	}
	if memo[i] != -1 {
		return memo[i]
	}

	// Two choices:
	// 1. Rob current house and skip next house
	// 2. Skip current house and consider next house
	robCurrent := nums[i] + robHelper(nums, i+2, memo)
	skipCurrent := robHelper(nums, i+1, memo)

	memo[i] = utils.Max(robCurrent, skipCurrent)
	return memo[i]
}