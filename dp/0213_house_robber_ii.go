package dp

import (
	"leetcode/utils"
)

// RobII solves LeetCode problem 0213: House Robber II
// Difficulty: Medium
// Tags: Dynamic Programming, Array
//
// You are a professional robber planning to rob houses along a street.
// Each house has a certain amount of money stashed. All houses at this place
// are arranged in a circle. That means the first house is the neighbor of the last one.
// Meanwhile, adjacent houses have a security system connected, and it will
// automatically contact the police if two adjacent houses were broken into on the same night.
//
// Given an integer array nums representing the amount of money of each house,
// return the maximum amount of money you can rob tonight without alerting the police.
//
// Example 1:
// Input: nums = [2,3,2]
// Output: 3
// Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2),
// because they are adjacent houses.
//
// Example 2:
// Input: nums = [1,2,3,1]
// Output: 4
// Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
// Total amount you can rob = 1 + 3 = 4.
//
// Example 3:
// Input: nums = [1,2,3]
// Output: 3
//
// Constraints:
// - 1 <= nums.length <= 100
// - 0 <= nums[i] <= 1000
//
// Time complexity: O(n), Space complexity: O(1)
func RobII(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return utils.Max(nums[0], nums[1])
	}

	// Since houses are arranged in a circle, we have two cases:
	// 1. Rob houses from 0 to n-2 (exclude last house)
	// 2. Rob houses from 1 to n-1 (exclude first house)
	// Take the maximum of these two cases
	case1 := robRange(nums, 0, len(nums)-2)
	case2 := robRange(nums, 1, len(nums)-1)

	return utils.Max(case1, case2)
}

// robRange solves the house robber problem for a specific range [l, r]
// This is the same as the original House Robber problem but for a subarray
func robRange(nums []int, l, r int) int {
	if l > r {
		return 0
	}
	if l == r {
		return nums[l]
	}

	prev2 := 0 // dp[i-2]
	prev1 := 0 // dp[i-1]

	for i := l; i <= r; i++ {
		current := utils.Max(prev1, nums[i]+prev2)
		prev2 = prev1
		prev1 = current
	}

	return prev1
}

// RobIIDPArray is an alternative implementation using DP array for clarity
func RobIIDPArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return utils.Max(nums[0], nums[1])
	}

	// Case 1: Rob houses 0 to n-2
	dp1 := make([]int, len(nums)-1)
	dp1[0] = nums[0]
	dp1[1] = utils.Max(nums[0], nums[1])

	for i := 2; i < len(nums)-1; i++ {
		dp1[i] = utils.Max(dp1[i-1], nums[i]+dp1[i-2])
	}

	// Case 2: Rob houses 1 to n-1
	dp2 := make([]int, len(nums)-1)
	dp2[0] = nums[1]
	dp2[1] = utils.Max(nums[1], nums[2])

	for i := 2; i < len(nums)-1; i++ {
		dp2[i] = utils.Max(dp2[i-1], nums[i+1]+dp2[i-2])
	}

	return utils.Max(dp1[len(dp1)-1], dp2[len(dp2)-1])
}

// RobIIRecursive is a recursive solution with memoization
func RobIIRecursive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return utils.Max(nums[0], nums[1])
	}

	// Two cases with memoization
	memo1 := make([]int, len(nums))
	memo2 := make([]int, len(nums))
	for i := range memo1 {
		memo1[i] = -1
		memo2[i] = -1
	}

	// Case 1: Include first house, exclude last house
	case1 := robIIHelper(nums, 0, len(nums)-2, memo1)
	// Case 2: Exclude first house, include last house
	case2 := robIIHelper(nums, 1, len(nums)-1, memo2)

	return utils.Max(case1, case2)
}

func robIIHelper(nums []int, start, end int, memo []int) int {
	if start > end {
		return 0
	}
	if memo[start] != -1 {
		return memo[start]
	}

	// Two choices:
	// 1. Rob current house and skip next house
	// 2. Skip current house and consider next house
	robCurrent := nums[start] + robIIHelper(nums, start+2, end, memo)
	skipCurrent := robIIHelper(nums, start+1, end, memo)

	memo[start] = utils.Max(robCurrent, skipCurrent)
	return memo[start]
}