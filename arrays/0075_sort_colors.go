package arrays

// SortColors solves LeetCode problem 0075: Sort Colors
// Difficulty: Medium
// Tags: Array, Two Pointers, Sorting
//
// Given an array nums with n objects colored red, white, or blue, sort them
// in-place so that objects of the same color are adjacent, with the colors in
// the order red, white, and blue.
//
// We will use the integers 0, 1, and 2 to represent the color red, white, and
// blue, respectively.
//
// You must solve this problem without using the library's sort function.
//
// Example 1:
// Input: nums = [2,0,2,1,1,0]
// Output: [0,0,1,1,2,2]
//
// Example 2:
// Input: nums = [2,0,1]
// Output: [0,1,2]
//
// Time complexity: O(n), Space complexity: O(1) (Dutch National Flag algorithm)
func SortColors(nums []int) {
	if len(nums) <= 1 {
		return
	}

	// Dutch National Flag algorithm (three-way partitioning)
	// Use three pointers:
	// - low: tracks the boundary of 0s (exclusive)
	// - mid: current element being processed
	// - high: tracks the boundary of 2s (exclusive)
	low, mid, high := 0, 0, len(nums)-1

	for mid <= high {
		switch nums[mid] {
		case 0:
			// Swap with low pointer and move both forward
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			// 1 is in correct position, just move mid forward
			mid++
		case 2:
			// Swap with high pointer and decrement high
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
			// Don't increment mid because we need to check the swapped element
		}
	}
}
