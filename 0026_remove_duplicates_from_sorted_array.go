package leetcode

// RemoveDuplicates solves LeetCode problem 0026: Remove Duplicates from Sorted Array
// Difficulty: Easy
// Tags: Array, Two Pointers
//
// Given an integer array nums sorted in non-decreasing order, remove the duplicates
// in-place such that each unique element appears only once. The relative order of
// the elements should be kept the same. Then return the number of unique elements in nums.
//
// Time complexity: O(n), Space complexity: O(1)
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Two pointers approach
	uniqueIndex := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[uniqueIndex] {
			uniqueIndex++
			nums[uniqueIndex] = nums[i]
		}
	}

	return uniqueIndex + 1
}

// RemoveDuplicatesGeneral removes duplicates from sorted array (general version)
func RemoveDuplicatesGeneral(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	uniqueIndex := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[uniqueIndex] {
			uniqueIndex++
			nums[uniqueIndex] = nums[i]
		}
	}

	return nums[:uniqueIndex+1]
}