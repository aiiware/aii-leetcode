package sliding_window

// 1004. Max Consecutive Ones III
// https://leetcode.com/problems/max-consecutive-ones-iii/
//
// Given a binary array nums and an integer k, return the maximum number of consecutive 1's
// in the array if you can flip at most k 0's.
//
// Example 1:
// Input: nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
// Output: 6
// Explanation: [1,1,1,0,0,1,1,1,1,1,1]
// Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
//
// Example 2:
// Input: nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
// Output: 10
// Explanation: [0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
// Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
//
// Constraints:
// - 1 <= nums.length <= 10^5
// - nums[i] is either 0 or 1.
// - 0 <= k <= nums.length
//
// Difficulty: Medium
// Tags: Array, Binary Search, Sliding Window, Prefix Sum

// longestOnes uses sliding window technique to find maximum consecutive 1's with at most k flips.
// Time complexity: O(n), Space complexity: O(1)
func longestOnes(nums []int, k int) int {
	left := 0
	zeroCount := 0
	maxLength := 0

	for right := 0; right < len(nums); right++ {
		// If current element is 0, increment zero count
		if nums[right] == 0 {
			zeroCount++
		}

		// If zero count exceeds k, shrink window from left
		for zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		// Update maximum length
		currentLength := right - left + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}

// longestOnesOptimized uses optimized sliding window that doesn't shrink unnecessarily.
// Time complexity: O(n), Space complexity: O(1)
func longestOnesOptimized(nums []int, k int) int {
	left := 0
	right := 0
	zeroCount := 0

	for right < len(nums) {
		// If current element is 0, increment zero count
		if nums[right] == 0 {
			zeroCount++
		}

		// If we have more than k zeros, move left pointer
		if zeroCount > k {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		right++
	}

	// The window size at the end is the maximum length
	return right - left
}

// longestOnesBruteForce is a naive O(n^2) solution for comparison.
func longestOnesBruteForce(nums []int, k int) int {
	maxLength := 0

	for i := 0; i < len(nums); i++ {
		zeroCount := 0
		for j := i; j < len(nums); j++ {
			if nums[j] == 0 {
				zeroCount++
			}
			if zeroCount > k {
				break
			}
			currentLength := j - i + 1
			if currentLength > maxLength {
				maxLength = currentLength
			}
		}
	}

	return maxLength
}