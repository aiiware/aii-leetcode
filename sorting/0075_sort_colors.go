package sorting

// SortColors solves LeetCode problem 0075: Sort Colors
// Difficulty: Medium
// Tags: Array, Two Pointers, Sorting
//
// Given an array nums with n objects colored red, white, or blue, sort them in-place
// so that objects of the same color are adjacent, with the colors in the order red, white, and blue.
//
// We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.
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
// Constraints:
// n == nums.length
// 1 <= n <= 300
// nums[i] is either 0, 1, or 2.
//
// Time complexity: O(n), Space complexity: O(1)
func SortColors(nums []int) {
	if len(nums) <= 1 {
		return
	}

	// Dutch National Flag algorithm (three-way partitioning)
	// Use three pointers:
	// - low: boundary for 0s (exclusive)
	// - mid: current element being processed
	// - high: boundary for 2s (exclusive)
	low, mid, high := 0, 0, len(nums)-1

	for mid <= high {
		switch nums[mid] {
		case 0:
			// Move 0 to the low boundary
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			// 1 is in the correct position, just move mid forward
			mid++
		case 2:
			// Move 2 to the high boundary
			nums[mid], nums[high] = nums[high], nums[mid]
			high--
			// Don't increment mid because we need to check the swapped element
		}
	}
}

// SortColorsCounting is an alternative solution using counting sort.
// This approach counts the number of 0s, 1s, and 2s, then overwrites the array.
// While this is O(n) time and O(1) space, it requires two passes through the array.
func SortColorsCounting(nums []int) {
	if len(nums) == 0 {
		return
	}

	// Count occurrences of each color
	counts := [3]int{}
	for _, num := range nums {
		counts[num]++
	}

	// Overwrite the array with sorted values
	index := 0
	for color, count := range counts {
		for i := 0; i < count; i++ {
			nums[index] = color
			index++
		}
	}
}

// SortColorsTwoPass is another alternative that first moves all 0s to the front,
// then moves all 2s to the back. This is also O(n) time and O(1) space.
func SortColorsTwoPass(nums []int) {
	if len(nums) <= 1 {
		return
	}

	// First pass: move all 0s to the front
	pos := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i], nums[pos] = nums[pos], nums[i]
			pos++
		}
	}

	// Second pass: move all 2s to the back
	pos = len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == 2 {
			nums[i], nums[pos] = nums[pos], nums[i]
			pos--
		}
	}
}