package dp

// UniquePaths solves LeetCode problem 0062: Unique Paths
// Difficulty: Medium
// Tags: Math, Dynamic Programming, Combinatorics
//
// There is a robot on an m x n grid. The robot is initially located at the top-left corner (grid[0][0]).
// The robot tries to move to the bottom-right corner (grid[m-1][n-1]). The robot can only move either
// down or right at any point in time.
//
// Given the two integers m and n, return the number of possible unique paths that the robot can take
// to reach the bottom-right corner.
//
// Example 1:
// Input: m = 3, n = 7
// Output: 28
//
// Example 2:
// Input: m = 3, n = 2
// Output: 3
// Explanation: From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
// 1. Right -> Down -> Down
// 2. Down -> Down -> Right
// 3. Down -> Right -> Down
//
// Time complexity: O(m*n), Space complexity: O(n)
func UniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}

	// Create DP array for current row
	dp := make([]int, n)
	
	// Initialize first row with 1 (only one way to reach any cell in first row: move right)
	for j := 0; j < n; j++ {
		dp[j] = 1
	}

	// Calculate for remaining rows
	for i := 1; i < m; i++ {
		// First column always has 1 way (move down)
		// For other columns: dp[j] = dp[j] (from above) + dp[j-1] (from left)
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}

	return dp[n-1]
}