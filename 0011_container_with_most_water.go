package leetcode

// MaxArea solves LeetCode problem 0011: Container With Most Water
// Difficulty: Medium
// Tags: Array, Two Pointers
//
// You are given an integer array height of length n. There are n vertical lines
// drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
//
// Find two lines that together with the x-axis form a container, such that the
// container contains the most water.
//
// Return the maximum area of water that the container can store.
//
// Time complexity: O(n), Space complexity: O(1)
func MaxArea(height []int) int {
	maxArea := 0
	left := 0
	right := len(height) - 1

	for left < right {
		// Calculate width and current height (minimum of two lines)
		width := right - left
		currentHeight := min(height[left], height[right])
		currentArea := width * currentHeight

		// Update max area
		if currentArea > maxArea {
			maxArea = currentArea
		}

		// Move the pointer pointing to the shorter line
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

// Helper function to find minimum of two integers
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
