package leetcode

// RemoveElement solves LeetCode problem 0027: Remove Element
// Difficulty: Easy
// Tags: Array, Two Pointers
//
// Given an integer array nums and an integer val, remove all occurrences of val
// in nums in-place. The order of the elements may be changed. Then return the
// number of elements in nums which are not equal to val.
//
// Time complexity: O(n), Space complexity: O(1)
func RemoveElement(nums []int, val int) int {
	// Two pointers approach
	index := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}

	return index
}

// RemoveElementTwoPointers uses two pointers from both ends (better when val is rare)
func RemoveElementTwoPointers(nums []int, val int) int {
	i := 0
	n := len(nums)

	for i < n {
		if nums[i] == val {
			// Swap with last element
			nums[i] = nums[n-1]
			n--
		} else {
			i++
		}
	}

	return n
}