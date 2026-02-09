package arrays

// SearchMatrix solves LeetCode problem 0074: Search a 2D Matrix
// Difficulty: Medium
// Tags: Array, Binary Search, Matrix
//
// You are given an m x n integer matrix matrix with the following two properties:
// 1. Each row is sorted in non-decreasing order.
// 2. The first integer of each row is greater than the last integer of the previous row.
//
// Given an integer target, return true if target is in matrix or false otherwise.
//
// You must write a solution in O(log(m * n)) time complexity.
//
// Example 1:
// Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
// Output: true
//
// Example 2:
// Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
// Output: false
//
// Constraints:
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 100
// -10^4 <= matrix[i][j], target <= 10^4
//
// Time complexity: O(log(m*n)), Space complexity: O(1)
func SearchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1

	for left <= right {
		mid := left + (right-left)/2
		// Convert 1D index to 2D coordinates
		row := mid / n
		col := mid % n
		midValue := matrix[row][col]

		if midValue == target {
			return true
		} else if midValue < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

// SearchMatrixTwoStep is an alternative solution that first finds the correct row,
// then searches within that row. This approach also has O(log(m) + log(n)) = O(log(m*n)) complexity.
func SearchMatrixTwoStep(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])

	// Step 1: Find the correct row using binary search
	top, bottom := 0, m-1
	for top <= bottom {
		row := top + (bottom-top)/2
		if matrix[row][0] <= target && target <= matrix[row][n-1] {
			// Step 2: Search within the row using binary search
			left, right := 0, n-1
			for left <= right {
				mid := left + (right-left)/2
				if matrix[row][mid] == target {
					return true
				} else if matrix[row][mid] < target {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
			return false
		} else if matrix[row][0] > target {
			bottom = row - 1
		} else {
			top = row + 1
		}
	}

	return false
}

// SearchMatrixLinear is a simpler O(m+n) solution that starts from the top-right corner
// and moves left or down based on comparison with target.
// This is included for educational purposes and comparison.
func SearchMatrixLinear(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])
	row, col := 0, n-1

	for row < m && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			col-- // Move left
		} else {
			row++ // Move down
		}
	}

	return false
}
