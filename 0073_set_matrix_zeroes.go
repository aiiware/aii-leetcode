package leetcode

// SetZeroes solves LeetCode problem 0073: Set Matrix Zeroes
// Difficulty: Medium
// Tags: Array, Hash Table, Matrix
//
// Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.
//
// You must do it in place.
//
// Example 1:
// Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
// Output: [[1,0,1],[0,0,0],[1,0,1]]
//
// Example 2:
// Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
// Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
//
// Constraints:
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 200
// -2^31 <= matrix[i][j] <= 2^31 - 1
//
// Follow up:
// 1. A straightforward solution using O(mn) space is probably a bad idea.
// 2. A simple improvement uses O(m + n) space, but still not the best solution.
// 3. Could you devise a constant space solution?
//
// Time complexity: O(m*n), Space complexity: O(1) (in-place)
func SetZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	m, n := len(matrix), len(matrix[0])
	firstRowHasZero, firstColHasZero := false, false

	// Check if first row has zero
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			firstRowHasZero = true
			break
		}
	}

	// Check if first column has zero
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			firstColHasZero = true
			break
		}
	}

	// Use first row and column as markers
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// Set zeros based on markers
	for i := 1; i < m; i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}

	for j := 1; j < n; j++ {
		if matrix[0][j] == 0 {
			for i := 1; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}

	// Handle first row
	if firstRowHasZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	// Handle first column
	if firstColHasZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}

// SetZeroesWithExtraSpace is an alternative solution using O(m+n) space
// This is included for comparison and educational purposes
func SetZeroesWithExtraSpace(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	m, n := len(matrix), len(matrix[0])
	rows := make([]bool, m)
	cols := make([]bool, n)

	// Mark rows and columns that contain zeros
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	// Set zeros based on markers
	for i := 0; i < m; i++ {
		if rows[i] {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}

	for j := 0; j < n; j++ {
		if cols[j] {
			for i := 0; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}
}