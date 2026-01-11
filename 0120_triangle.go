package leetcode

// Problem 120: Triangle
// Given a triangle array, return the minimum path sum from top to bottom.
// For each step, you may move to an adjacent number on the row below.

// minimumTotal returns the minimum path sum from top to bottom of the triangle.
// This implementation uses O(n) extra space where n is the number of rows.
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	// Start from the second last row and work upwards
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			// Add the minimum of the two adjacent numbers from the row below
			left := triangle[i+1][j]
			right := triangle[i+1][j+1]
			if left < right {
				triangle[i][j] += left
			} else {
				triangle[i][j] += right
			}
		}
	}

	return triangle[0][0]
}

// minimumTotalDP uses dynamic programming with separate dp array.
func minimumTotalDP(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	// Create a dp array for the last row
	dp := make([]int, len(triangle))
	lastRow := triangle[len(triangle)-1]
	copy(dp, lastRow)

	// Work from second last row upwards
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			// Update dp[j] with current value + min of two below
			left := dp[j]
			right := dp[j+1]
			if left < right {
				dp[j] = triangle[i][j] + left
			} else {
				dp[j] = triangle[i][j] + right
			}
		}
	}

	return dp[0]
}

// minimumTotalDFS uses DFS with memoization (top-down approach).
func minimumTotalDFS(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	// Create memoization table
	memo := make([][]int, len(triangle))
	for i := range memo {
		memo[i] = make([]int, len(triangle[i]))
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return dfsTriangle(triangle, 0, 0, memo)
}

// dfsTriangle is the helper function for DFS with memoization for triangle problem.
func dfsTriangle(triangle [][]int, row, col int, memo [][]int) int {
	if row == len(triangle)-1 {
		return triangle[row][col]
	}

	if memo[row][col] != -1 {
		return memo[row][col]
	}

	// Explore both paths
	left := dfsTriangle(triangle, row+1, col, memo)
	right := dfsTriangle(triangle, row+1, col+1, memo)

	// Take minimum path
	minPath := triangle[row][col]
	if left < right {
		minPath += left
	} else {
		minPath += right
	}

	memo[row][col] = minPath
	return minPath
}

// Helper function to create a triangle from 2D slice
func createTriangle(nums [][]int) [][]int {
	triangle := make([][]int, len(nums))
	for i := range nums {
		triangle[i] = make([]int, len(nums[i]))
		copy(triangle[i], nums[i])
	}
	return triangle
}

// Helper function to find minimum path sum using brute force (for testing)
func minimumTotalBruteForce(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	return bruteForce(triangle, 0, 0)
}

func bruteForce(triangle [][]int, row, col int) int {
	if row == len(triangle)-1 {
		return triangle[row][col]
	}

	left := bruteForce(triangle, row+1, col)
	right := bruteForce(triangle, row+1, col+1)

	if left < right {
		return triangle[row][col] + left
	}
	return triangle[row][col] + right
}