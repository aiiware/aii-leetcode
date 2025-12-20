package leetcode

// TwoSum solves LeetCode problem 0001: Two Sum
// Difficulty: Easy
// Tags: Array, Hash Table
//
// Given an array of integers nums and an integer target, return indices of the
// two numbers such that they add up to target.
//
// You may assume that each input would have exactly one solution, and you may
// not use the same element twice.
//
// Time complexity: O(n), Space complexity: O(n)
func TwoSum(nums []int, target int) []int {
	// Create a map to store number -> index
	numMap := make(map[int]int)

	for i, num := range nums {
		// Calculate the complement needed to reach target
		complement := target - num

		// Check if complement exists in map
		if idx, found := numMap[complement]; found {
			// Return indices (complement index first, then current index)
			return []int{idx, i}
		}

		// Store current number and its index
		numMap[num] = i
	}

	// According to problem statement, exactly one solution exists
	// So we should never reach here, but return empty slice for safety
	return []int{}
}