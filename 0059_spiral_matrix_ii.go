package leetcode

// GenerateMatrix solves LeetCode problem 0059: Spiral Matrix II
// Difficulty: Medium
// Tags: Array, Matrix, Simulation
//
// Given a positive integer n, generate an n x n matrix filled with elements from 1 to n^2 in spiral order.
//
// Example 1:
// Input: n = 3
// Output: [[1,2,3],[8,9,4],[7,6,5]]
//
// Example 2:
// Input: n = 1
// Output: [[1]]
//
// Time complexity: O(n^2), Space complexity: O(n^2)
func GenerateMatrix(n int) [][]int {
	if n <= 0 {
		return [][]int{}
	}

	// Initialize matrix with zeros
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	// Define boundaries
	top, bottom := 0, n-1
	left, right := 0, n-1
	num := 1

	for top <= bottom && left <= right {
		// Fill top row from left to right
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++

		// Fill right column from top to bottom
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--

		// Fill bottom row from right to left (if still within bounds)
		if top <= bottom {
			for i := right; i >= left; i-- {
				matrix[bottom][i] = num
				num++
			}
			bottom--
		}

		// Fill left column from bottom to top (if still within bounds)
		if left <= right {
			for i := bottom; i >= top; i-- {
				matrix[i][left] = num
				num++
			}
			left++
		}
	}

	return matrix
}