package arrays

// FindDuplicates solves LeetCode problem 0442: Find All Duplicates in an Array
// Difficulty: Medium
// Tags: Array, Hash Table
//
// Given an integer array nums of length n where all the integers of nums are in the range [1, n]
// and each integer appears once or twice, return an array of all the integers that appear twice.
//
// You must write an algorithm that runs in O(n) time and uses only constant extra space.
//
// Example 1:
// Input: nums = [4,3,2,7,8,2,3,1]
// Output: [2,3]
//
// Example 2:
// Input: nums = [1,1,2]
// Output: [1]
//
// Example 3:
// Input: nums = [1]
// Output: []
//
// Constraints:
// n == nums.length
// 1 <= n <= 10^5
// 1 <= nums[i] <= n
// Each element in nums appears once or twice.
//
// Time complexity: O(n), Space complexity: O(1) excluding output array
func FindDuplicates(nums []int) []int {
	result := []int{}

	// Use negative marking technique
	for i := 0; i < len(nums); i++ {
		// Get absolute value (since we might have marked it negative)
		val := absInt(nums[i])
		idx := val - 1 // Convert to 0-based index

		// If the value at this index is already negative,
		// it means we've seen this number before (it's a duplicate)
		if nums[idx] < 0 {
			result = append(result, val)
		} else {
			// Mark this number as seen by making it negative
			nums[idx] = -nums[idx]
		}
	}

	return result
}