package arrays

// MoveZeroes solves LeetCode problem 0283: Move Zeroes
// Difficulty: Easy
// Tags: Array, Two Pointers
//
// Given an integer array nums, move all 0's to the end of it while maintaining
// the relative order of the non-zero elements.
//
// Note that you must do this in-place without making a copy of the array.
//
// Example 1:
// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]
//
// Example 2:
// Input: nums = [0]
// Output: [0]
//
// Time complexity: O(n), Space complexity: O(1)
func MoveZeroes(nums []int) {
	// Two pointers approach
	// j tracks the position for the next non-zero element
	j := 0

	// First pass: move all non-zero elements to the front
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}

	// Second pass: fill the remaining positions with zeros
	for i := j; i < len(nums); i++ {
		nums[i] = 0
	}
}

// MoveZeroesSinglePass uses a single pass with swapping
func MoveZeroesSinglePass(nums []int) {
	// Two pointers approach with swapping
	// j tracks the position for the next non-zero element
	j := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			// Swap nums[i] and nums[j]
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
}

// MoveZeroesSnowball is an alternative approach that counts zeros
func MoveZeroesSnowball(nums []int) {
	snowballSize := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			snowballSize++
		} else if snowballSize > 0 {
			// Swap current element with the first zero in the snowball
			nums[i], nums[i-snowballSize] = nums[i-snowballSize], nums[i]
		}
	}
}