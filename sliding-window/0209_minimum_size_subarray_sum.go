package sliding_window

// 0209. Minimum Size Subarray Sum
// https://leetcode.com/problems/minimum-size-subarray-sum/
//
// Given an array of positive integers nums and a positive integer target,
// return the minimal length of a contiguous subarray [nums_l, nums_l+1, ..., nums_r-1, nums_r]
// of which the sum is greater than or equal to target.
// If there is no such subarray, return 0 instead.

// minSubArrayLen uses sliding window technique to find the minimum length subarray
// with sum >= target.
// Time complexity: O(n), Space complexity: O(1)
func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	minLength := len(nums) + 1 // Initialize with a value larger than possible
	currentSum := 0
	left := 0

	for right := 0; right < len(nums); right++ {
		// Add current element to the window sum
		currentSum += nums[right]

		// While the current window sum is >= target, try to shrink the window
		// from the left to find the minimum length
		for currentSum >= target && left <= right {
			// Update minimum length if current window is smaller
			windowLength := right - left + 1
			if windowLength < minLength {
				minLength = windowLength
			}

			// Shrink window from the left
			currentSum -= nums[left]
			left++
		}
	}

	// If minLength was never updated, return 0
	if minLength > len(nums) {
		return 0
	}
	return minLength
}

// minSubArrayLenBruteForce is a naive O(n^2) solution for comparison.
func minSubArrayLenBruteForce(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	minLength := len(nums) + 1

	for i := 0; i < len(nums); i++ {
		currentSum := 0
		for j := i; j < len(nums); j++ {
			currentSum += nums[j]
			if currentSum >= target {
				length := j - i + 1
				if length < minLength {
					minLength = length
				}
				break // Found a valid subarray starting at i, move to next i
			}
		}
	}

	if minLength > len(nums) {
		return 0
	}
	return minLength
}