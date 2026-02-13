package sliding_window

// 0643. Maximum Average Subarray I
// https://leetcode.com/problems/maximum-average-subarray-i/
//
// You are given an integer array nums consisting of n elements, and an integer k.
// Find a contiguous subarray whose length is equal to k that has the maximum average value
// and return this value. Any answer with a calculation error less than 10^-5 will be accepted.
//
// Example 1:
// Input: nums = [1,12,-5,-6,50,3], k = 4
// Output: 12.75000
// Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75
//
// Example 2:
// Input: nums = [5], k = 1
// Output: 5.00000
//
// Constraints:
// - n == nums.length
// - 1 <= k <= n <= 10^5
// - -10^4 <= nums[i] <= 10^4
//
// Difficulty: Easy
// Tags: Array, Sliding Window

// findMaxAverage uses sliding window technique to find maximum average of subarray of length k.
// Time complexity: O(n), Space complexity: O(1)
func findMaxAverage(nums []int, k int) float64 {
	if len(nums) == 0 || k == 0 {
		return 0
	}

	// Calculate sum of first window
	currentSum := 0
	for i := 0; i < k; i++ {
		currentSum += nums[i]
	}
	maxSum := currentSum

	// Slide the window
	for i := k; i < len(nums); i++ {
		// Add new element and remove old element
		currentSum = currentSum + nums[i] - nums[i-k]
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	// Return average
	return float64(maxSum) / float64(k)
}

// findMaxAverageBruteForce is a naive O(n*k) solution for comparison.
func findMaxAverageBruteForce(nums []int, k int) float64 {
	if len(nums) == 0 || k == 0 {
		return 0
	}

	maxSum := -1 << 31 // Initialize with minimum integer value

	for i := 0; i <= len(nums)-k; i++ {
		currentSum := 0
		for j := 0; j < k; j++ {
			currentSum += nums[i+j]
		}
		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return float64(maxSum) / float64(k)
}