package arrays

// FindDisappearedNumbers solves LeetCode problem 0448: Find All Numbers Disappeared in an Array
// Difficulty: Easy
// Tags: Array, Hash Table
//
// Given an array nums of n integers where nums[i] is in the range [1, n],
// return an array of all the integers in the range [1, n] that do not appear in nums.
//
// Example 1:
// Input: nums = [4,3,2,7,8,2,3,1]
// Output: [5,6]
//
// Example 2:
// Input: nums = [1,1]
// Output: [2]
//
// Constraints:
// n == nums.length
// 1 <= n <= 10^5
// 1 <= nums[i] <= n
//
// Time complexity: O(n), Space complexity: O(1) excluding output array
func FindDisappearedNumbers(nums []int) []int {
	n := len(nums)

	// First pass: mark numbers that appear by using negative marking
	for i := 0; i < n; i++ {
		// Get the absolute value since we might have marked it negative already
		val := abs(nums[i])

		// Use 0-based indexing: val-1 gives us the correct index
		// val is in range [1, n], so idx is in range [0, n-1]
		idx := val - 1

		// Mark this number as seen by making the value at this index negative
		// Only mark if it's positive to avoid double-negating
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}

	// Second pass: collect indices where values are still positive
	// These indices correspond to missing numbers
	result := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			// Convert back to 1-based number
			result = append(result, i+1)
		}
	}

	return result
}

// Helper function to get absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
