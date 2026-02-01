package arrays

// BinarySearch solves LeetCode problem 0704: Binary Search
// Difficulty: Easy
// Tags: Array, Binary Search
//
// Given an array of integers nums which is sorted in ascending order, and an
// integer target, write a function to search target in nums. If target exists,
// then return its index. Otherwise, return -1.
//
// You must write an algorithm with O(log n) runtime complexity.
//
// Time complexity: O(log n), Space complexity: O(1)
func BinarySearch(nums []int, target int) int {
	// Edge case: empty array
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1

	for left <= right {
		// Calculate mid point (prevents overflow for large arrays)
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			// Target is in right half
			left = mid + 1
		} else {
			// Target is in left half
			right = mid - 1
		}
	}

	// Target not found
	return -1
}