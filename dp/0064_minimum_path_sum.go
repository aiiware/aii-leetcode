package dp

// MinPathSum solves LeetCode problem 0064: Minimum Path Sum
// Difficulty: Medium
// Tags: Array, Dynamic Programming, Matrix
//
// Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right,
// which minimizes the sum of all numbers along its path.
//
// Note: You can only move either down or right at any point in time.
//
// Example 1:
// Input: grid = [[1,3,1],[1,5,1],[4,2,1]]
// Output: 7
// Explanation: Because the path 1 → 3 → 1 → 1 → 1 minimizes the sum.
//
// Example 2:
// Input: grid = [[1,2,3],[4,5,6]]
// Output: 12
//
// Time complexity: O(m*n), Space complexity: O(n)
func MinPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])

	// Create DP array for current row
	dp := make([]int, n)

	// Initialize first cell
	dp[0] = grid[0][0]

	// Initialize first row (can only come from left)
	for j := 1; j < n; j++ {
		dp[j] = dp[j-1] + grid[0][j]
	}

	// Calculate for remaining rows
	for i := 1; i < m; i++ {
		// Update first column (can only come from above)
		dp[0] = dp[0] + grid[i][0]

		// Update other columns (min of coming from left or above)
		for j := 1; j < n; j++ {
			// dp[j] currently holds value from above (previous row)
			// dp[j-1] holds value from left (current row)
			if dp[j] < dp[j-1] {
				dp[j] = dp[j] + grid[i][j]
			} else {
				dp[j] = dp[j-1] + grid[i][j]
			}
		}
	}

	return dp[n-1]
}