package leetcode

// UniquePathsWithObstacles solves LeetCode problem 0063: Unique Paths II
// Difficulty: Medium
// Tags: Array, Dynamic Programming, Matrix
//
// You are given an m x n integer array grid. There is a robot initially located at the top-left corner (grid[0][0]).
// The robot tries to move to the bottom-right corner (grid[m-1][n-1]). The robot can only move either down or right
// at any point in time.
//
// An obstacle and space are marked as 1 or 0 respectively in grid. A path that the robot takes cannot include any
// square that is an obstacle.
//
// Return the number of possible unique paths that the robot can take to reach the bottom-right corner.
//
// Example 1:
// Input: obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
// Output: 2
// Explanation: There is one obstacle in the middle of the 3x3 grid above.
// There are two ways to reach the bottom-right corner:
// 1. Right -> Right -> Down -> Down
// 2. Down -> Down -> Right -> Right
//
// Example 2:
// Input: obstacleGrid = [[0,1],[0,0]]
// Output: 1
//
// Time complexity: O(m*n), Space complexity: O(n)
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}

	m, n := len(obstacleGrid), len(obstacleGrid[0])

	// If start or end is blocked, no paths exist
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}

	// Create DP array for current row
	dp := make([]int, n)
	
	// Initialize first cell
	dp[0] = 1

	// Initialize first row
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			dp[j] = 0
		} else {
			dp[j] = dp[j-1]
		}
	}

	// Calculate for remaining rows
	for i := 1; i < m; i++ {
		// Update first column
		if obstacleGrid[i][0] == 1 {
			dp[0] = 0
		}
		// dp[0] remains as is if not blocked

		// Update other columns
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
			} else {
				dp[j] = dp[j] + dp[j-1]
			}
		}
	}

	return dp[n-1]
}