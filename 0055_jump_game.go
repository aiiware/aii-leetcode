package leetcode

// CanJump solves LeetCode problem 0055: Jump Game
// Difficulty: Medium
// Tags: Array, Greedy, Dynamic Programming
//
// You are given an integer array nums. You are initially positioned at the
// array's first index, and each element in the array represents your maximum
// jump length at that position.
//
// Return true if you can reach the last index, or false otherwise.
//
// Example 1:
// Input: nums = [2,3,1,1,4]
// Output: true
// Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
//
// Example 2:
// Input: nums = [3,2,1,0,4]
// Output: false
// Explanation: You will always arrive at index 3 no matter what. Its maximum
// jump length is 0, which makes it impossible to reach the last index.
//
// Time complexity: O(n), Space complexity: O(1)
func CanJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	// Greedy approach: track the farthest reachable index
	maxReach := 0

	for i := 0; i < len(nums); i++ {
		// If current index is beyond what we can reach, return false
		if i > maxReach {
			return false
		}

		// Update the farthest reachable index
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
		}

		// If we can reach or go beyond the last index, return true
		if maxReach >= len(nums)-1 {
			return true
		}
	}

	return false
}