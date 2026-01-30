package math

// SpiralOrder solves LeetCode problem 0054: Spiral Matrix
// Difficulty: Medium
// Tags: Array, Matrix, Simulation
//
// Given an m x n matrix, return all elements of the matrix in spiral order.
//
// Example 1:
// Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
// Output: [1,2,3,6,9,8,7,4,5]
//
// Example 2:
// Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
// Output: [1,2,3,4,8,12,11,10,9,5,6,7]
//
// Time complexity: O(m*n), Space complexity: O(1) excluding output
func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	m, n := len(matrix), len(matrix[0])
	result := make([]int, 0, m*n)

	// Define boundaries
	top, bottom := 0, m-1
	left, right := 0, n-1

	for top <= bottom && left <= right {
		// Traverse from left to right along the top row
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		// Traverse from top to bottom along the right column
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		// Check if we still have rows to traverse
		if top <= bottom {
			// Traverse from right to left along the bottom row
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}

		// Check if we still have columns to traverse
		if left <= right {
			// Traverse from bottom to top along the left column
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}

	return result
}