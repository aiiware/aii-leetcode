package arrays

// LongestConsecutive solves LeetCode problem 0128: Longest Consecutive Sequence
// Difficulty: Medium
// Tags: Array, Hash Table, Union Find
//
// Given an unsorted array of integers nums, return the length of the longest consecutive
// elements sequence.
//
// You must write an algorithm that runs in O(n) time.
//
// Example 1:
// Input: nums = [100,4,200,1,3,2]
// Output: 4
// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
//
// Example 2:
// Input: nums = [0,3,7,2,5,8,4,6,0,1]
// Output: 9
//
// Constraints:
// 0 <= nums.length <= 10^5
// -10^9 <= nums[i] <= 10^9
//
// Time complexity: O(n), Space complexity: O(n)
func LongestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Create a set of all numbers for O(1) lookups
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	longestStreak := 0

	// For each number, check if it's the start of a sequence
	for num := range numSet {
		// Check if this number is the start of a sequence
		// (i.e., num-1 is not in the set)
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1

			// Count consecutive numbers
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}

			// Update longest streak
			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}

	return longestStreak
}