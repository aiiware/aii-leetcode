package sliding_window

// 0239. Sliding Window Maximum
// https://leetcode.com/problems/sliding-window-maximum/
//
// You are given an array of integers nums, there is a sliding window of size k
// which is moving from the very left of the array to the very right.
// You can only see the k numbers in the window. Each time the sliding window moves
// right by one position.
//
// Return the max sliding window.

// maxSlidingWindow uses a deque (double-ended queue) to track indices of potential maximums.
// Time complexity: O(n), Space complexity: O(k)
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}

	// Result slice to store maximums for each window
	result := make([]int, 0, len(nums)-k+1)

	// Deque to store indices of elements in the current window
	// The deque maintains indices of elements in decreasing order of their values
	deque := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		// Remove indices from the front that are out of the current window
		if len(deque) > 0 && deque[0] < i-k+1 {
			deque = deque[1:]
		}

		// Remove indices from the back while the current element is greater
		// This maintains decreasing order in the deque
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}

		// Add current index to the back of deque
		deque = append(deque, i)

		// Once we've processed at least k elements, add the maximum to result
		// The maximum is always at the front of the deque
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}

// maxSlidingWindowBruteForce is a naive O(n*k) solution for comparison.
// Not efficient for large inputs but useful for understanding the problem.
func maxSlidingWindowBruteForce(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}

	result := make([]int, len(nums)-k+1)
	for i := 0; i <= len(nums)-k; i++ {
		max := nums[i]
		for j := 1; j < k; j++ {
			if nums[i+j] > max {
				max = nums[i+j]
			}
		}
		result[i] = max
	}
	return result
}