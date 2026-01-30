package dp

// LargestRectangleArea solves LeetCode problem 0084: Largest Rectangle in Histogram
// Difficulty: Hard
// Tags: Array, Stack, Monotonic Stack
//
// Given an array of integers heights representing the histogram's bar height where
// the width of each bar is 1, return the area of the largest rectangle in the histogram.
//
// Example 1:
// Input: heights = [2,1,5,6,2,3]
// Output: 10
// Explanation: The above is a histogram where width of each bar is 1.
// The largest rectangle is shown in the shaded area, which has an area = 10 units.
//
// Example 2:
// Input: heights = [2,4]
// Output: 4
//
// Constraints:
// 1 <= heights.length <= 10^5
// 0 <= heights[i] <= 10^4
//
// Time complexity: O(n) where n is the length of heights
// Space complexity: O(n) for the stack
func LargestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return heights[0]
	}

	// Create arrays to store the index of the nearest smaller element to the left and right
	left := make([]int, n)  // left[i] = index of nearest smaller element to the left of i
	right := make([]int, n) // right[i] = index of nearest smaller element to the right of i

	// Initialize right array with n (no smaller element to the right)
	for i := 0; i < n; i++ {
		right[i] = n
	}

	// Use a monotonic increasing stack
	stack := make([]int, 0, n)

	// Calculate left boundaries and right boundaries in one pass
	for i := 0; i < n; i++ {
		// While stack is not empty and current height is less than height at stack top
		for len(stack) > 0 && heights[i] < heights[stack[len(stack)-1]] {
			// For the element at stack top, the current index i is the nearest smaller to the right
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			right[top] = i
		}

		// If stack is empty, there's no smaller element to the left
		if len(stack) == 0 {
			left[i] = -1
		} else {
			// The element at stack top is the nearest smaller to the left
			left[i] = stack[len(stack)-1]
		}

		// Push current index to stack
		stack = append(stack, i)
	}

	// Calculate maximum area
	maxArea := 0
	for i := 0; i < n; i++ {
		// Width = distance between right and left boundaries minus 1
		width := right[i] - left[i] - 1
		area := heights[i] * width
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

// LargestRectangleAreaOptimized is an optimized version that uses a single pass
// with a more efficient stack implementation
func LargestRectangleAreaOptimized(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}

	// Add sentinel values to handle boundaries
	// Append 0 to the end to force stack to empty at the end
	heightsWithSentinel := make([]int, n+1)
	copy(heightsWithSentinel, heights)
	heightsWithSentinel[n] = 0

	stack := make([]int, 0, n+1)
	maxArea := 0

	for i := 0; i <= n; i++ {
		// While stack is not empty and current height is less than height at stack top
		for len(stack) > 0 && heightsWithSentinel[i] < heightsWithSentinel[stack[len(stack)-1]] {
			// Pop the top
			h := heightsWithSentinel[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]

			// Calculate width
			var width int
			if len(stack) == 0 {
				width = i
			} else {
				width = i - stack[len(stack)-1] - 1
			}

			// Update max area
			area := h * width
			if area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}

	return maxArea
}

// LargestRectangleAreaBruteForce is a brute force solution for comparison
// Not recommended for large inputs (O(n^2) time complexity)
func LargestRectangleAreaBruteForce(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}

	maxArea := 0

	for i := 0; i < n; i++ {
		// Find left boundary
		left := i
		for left-1 >= 0 && heights[left-1] >= heights[i] {
			left--
		}

		// Find right boundary
		right := i
		for right+1 < n && heights[right+1] >= heights[i] {
			right++
		}

		// Calculate area
		width := right - left + 1
		area := heights[i] * width
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

// LargestRectangleAreaDivideConquer uses divide and conquer approach
// Time complexity: O(n log n) average, O(n^2) worst case
func LargestRectangleAreaDivideConquer(heights []int) int {
	return largestRectangleAreaDivideConquerHelper(heights, 0, len(heights)-1)
}

func largestRectangleAreaDivideConquerHelper(heights []int, left, right int) int {
	if left > right {
		return 0
	}
	if left == right {
		return heights[left]
	}

	// Find the index of the minimum height in the current range
	minIndex := left
	for i := left + 1; i <= right; i++ {
		if heights[i] < heights[minIndex] {
			minIndex = i
		}
	}

	// Calculate area with minimum height as the bar
	areaWithMin := heights[minIndex] * (right - left + 1)

	// Calculate maximum area in left and right subarrays
	areaLeft := largestRectangleAreaDivideConquerHelper(heights, left, minIndex-1)
	areaRight := largestRectangleAreaDivideConquerHelper(heights, minIndex+1, right)

	return max(areaWithMin, max(areaLeft, areaRight))
}

// LargestRectangleAreaDP uses dynamic programming approach
// This is similar to the first solution but separates left and right boundary calculations
func LargestRectangleAreaDP(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}

	// Arrays to store boundaries
	leftBoundary := make([]int, n)
	rightBoundary := make([]int, n)

	// Calculate left boundaries
	stack := make([]int, 0, n)
	for i := 0; i < n; i++ {
		// Pop elements while current height is less than or equal to stack top
		for len(stack) > 0 && heights[i] <= heights[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			leftBoundary[i] = -1
		} else {
			leftBoundary[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	// Clear stack for right boundary calculation
	stack = stack[:0]

	// Calculate right boundaries
	for i := n - 1; i >= 0; i-- {
		// Pop elements while current height is less than or equal to stack top
		for len(stack) > 0 && heights[i] <= heights[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			rightBoundary[i] = n
		} else {
			rightBoundary[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	// Calculate maximum area
	maxArea := 0
	for i := 0; i < n; i++ {
		width := rightBoundary[i] - leftBoundary[i] - 1
		area := heights[i] * width
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

// LargestRectangleAreaSegmentTree uses segment tree for range minimum queries
// Time complexity: O(n log n)
func LargestRectangleAreaSegmentTree(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}

	// Build segment tree for range minimum queries
	segTree := buildSegmentTree(heights)

	return largestRectangleAreaSegmentTreeHelper(heights, segTree, 0, n-1)
}

func largestRectangleAreaSegmentTreeHelper(heights []int, segTree []int, left, right int) int {
	if left > right {
		return 0
	}
	if left == right {
		return heights[left]
	}

	// Find minimum height index in range [left, right]
	minIndex := querySegmentTree(heights, segTree, left, right)

	// Calculate area with minimum height
	areaWithMin := heights[minIndex] * (right - left + 1)

	// Calculate areas in left and right subarrays
	areaLeft := largestRectangleAreaSegmentTreeHelper(heights, segTree, left, minIndex-1)
	areaRight := largestRectangleAreaSegmentTreeHelper(heights, segTree, minIndex+1, right)

	return max(areaWithMin, max(areaLeft, areaRight))
}

// Segment tree helper functions
func buildSegmentTree(arr []int) []int {
	n := len(arr)
	segTree := make([]int, 4*n)
	buildSegmentTreeHelper(arr, segTree, 0, 0, n-1)
	return segTree
}

func buildSegmentTreeHelper(arr []int, segTree []int, node, start, end int) {
	if start == end {
		segTree[node] = start
		return
	}

	mid := (start + end) / 2
	leftChild := 2*node + 1
	rightChild := 2*node + 2

	buildSegmentTreeHelper(arr, segTree, leftChild, start, mid)
	buildSegmentTreeHelper(arr, segTree, rightChild, mid+1, end)

	// Store index of minimum value
	if arr[segTree[leftChild]] <= arr[segTree[rightChild]] {
		segTree[node] = segTree[leftChild]
	} else {
		segTree[node] = segTree[rightChild]
	}
}

func querySegmentTree(arr []int, segTree []int, left, right int) int {
	n := len(arr)
	return querySegmentTreeHelper(arr, segTree, 0, 0, n-1, left, right)
}

func querySegmentTreeHelper(arr []int, segTree []int, node, start, end, left, right int) int {
	if right < start || left > end {
		return -1 // Invalid index
	}
	if left <= start && end <= right {
		return segTree[node]
	}

	mid := (start + end) / 2
	leftChild := 2*node + 1
	rightChild := 2*node + 2

	leftIndex := querySegmentTreeHelper(arr, segTree, leftChild, start, mid, left, right)
	rightIndex := querySegmentTreeHelper(arr, segTree, rightChild, mid+1, end, left, right)

	if leftIndex == -1 {
		return rightIndex
	}
	if rightIndex == -1 {
		return leftIndex
	}
	if arr[leftIndex] <= arr[rightIndex] {
		return leftIndex
	}
	return rightIndex
}

// Helper function for max
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}