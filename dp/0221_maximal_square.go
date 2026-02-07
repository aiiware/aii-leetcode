package dp

// MaximalSquare solves LeetCode problem 0221: Maximal Square
// Difficulty: Medium
// Tags: Dynamic Programming, Matrix
//
// Given an m x n binary matrix filled with 0's and 1's,
// find the largest square containing only 1's and return its area.
//
// Example:
// Input: matrix = [
//   ["1","0","1","0","0"],
//   ["1","0","1","1","1"],
//   ["1","1","1","1","1"],
//   ["1","0","0","1","0"]
// ]
// Output: 4
//
// Time complexity: O(m*n), Space complexity: O(m*n)
func MaximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	m, n := len(matrix), len(matrix[0])
	// dp[i][j] represents the side length of the largest square
	// whose bottom-right corner is at (i-1, j-1)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	maxSide := 0

	// Fill dp table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				// The side length of the square ending at (i-1, j-1) is
				// 1 + min of squares ending at left, top, and top-left
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				if dp[i][j] > maxSide {
					maxSide = dp[i][j]
				}
			}
		}
	}

	return maxSide * maxSide
}

// Helper function to find minimum of three integers
func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}