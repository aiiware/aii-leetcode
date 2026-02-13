package sliding_window

// 1343. Number of Sub-arrays of Size K and Average Greater than or Equal to Threshold
// https://leetcode.com/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold/
//
// Given an array of integers arr and two integers k and threshold, return the number of sub-arrays
// of size k whose average is greater than or equal to threshold.
//
// Example 1:
// Input: arr = [2,2,2,2,5,5,5,8], k = 3, threshold = 4
// Output: 3
// Explanation: Sub-arrays [2,5,5],[5,5,5] and [5,5,8] have averages 4, 5 and 6 respectively.
// All other sub-arrays of size 3 have averages less than 4 (the threshold).
//
// Example 2:
// Input: arr = [11,13,17,23,29,31,7,5,2,3], k = 3, threshold = 5
// Output: 6
// Explanation: The first 6 sub-arrays of size 3 have averages greater than 5.
//
// Constraints:
// - 1 <= arr.length <= 10^5
// - 1 <= arr[i] <= 10^4
// - 1 <= k <= arr.length
// - 0 <= threshold <= 10^4
//
// Difficulty: Medium
// Tags: Array, Sliding Window

// numOfSubarrays uses sliding window technique to count subarrays with average >= threshold.
// Time complexity: O(n), Space complexity: O(1)
func numOfSubarrays(arr []int, k int, threshold int) int {
	if len(arr) < k {
		return 0
	}

	count := 0
	// Calculate sum of first window
	currentSum := 0
	for i := 0; i < k; i++ {
		currentSum += arr[i]
	}

	// Check first window
	if currentSum >= k*threshold {
		count++
	}

	// Slide the window
	for i := k; i < len(arr); i++ {
		// Add new element and remove old element
		currentSum = currentSum + arr[i] - arr[i-k]
		if currentSum >= k*threshold {
			count++
		}
	}

	return count
}

// numOfSubarraysBruteForce is a naive O(n*k) solution for comparison.
func numOfSubarraysBruteForce(arr []int, k int, threshold int) int {
	if len(arr) < k {
		return 0
	}

	count := 0
	for i := 0; i <= len(arr)-k; i++ {
		currentSum := 0
		for j := 0; j < k; j++ {
			currentSum += arr[i+j]
		}
		if currentSum >= k*threshold {
			count++
		}
	}
	return count
}