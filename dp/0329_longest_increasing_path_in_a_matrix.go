package dp

// longestIncreasingPath solves LeetCode problem 0329: Longest Increasing Path in a Matrix
// Difficulty: Hard
// Tags: Depth-First Search, Breadth-First Search, Graph, Topological Sort, Memoization, Dynamic Programming
//
// Given an m x n integers matrix, return the length of the longest increasing path in matrix.
// From each cell, you can either move in four directions: left, right, up, or down.
// You may not move diagonally or move outside the boundary (i.e., wrap-around is not allowed).
//
// Example 1:
// Input: matrix = [[9,9,4],[6,6,8],[2,1,1]]
// Output: 4
// Explanation: The longest increasing path is [1, 2, 6, 9].
//
// Example 2:
// Input: matrix = [[3,4,5],[3,2,6],[2,2,1]]
// Output: 4
// Explanation: The longest increasing path is [3, 4, 5, 6].
//
// Example 3:
// Input: matrix = [[1]]
// Output: 1
//
// Constraints:
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 200
// 0 <= matrix[i][j] <= 2^31 - 1
//
// Time complexity: O(m*n), Space complexity: O(m*n)
func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	m, n := len(matrix), len(matrix[0])
	
	// Memoization cache: dp[i][j] stores the longest increasing path starting from (i, j)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Directions: up, down, left, right
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// DFS with memoization
	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		// If already computed, return cached result
		if dp[row][col] != 0 {
			return dp[row][col]
		}

		// Initialize path length to 1 (the cell itself)
		maxLength := 1

		// Try all 4 directions
		for _, dir := range directions {
			newRow, newCol := row+dir[0], col+dir[1]

			// Check bounds and if next cell has greater value
			if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n && 
			   matrix[newRow][newCol] > matrix[row][col] {
				// Recursively compute path length from neighbor
				length := 1 + dfs(newRow, newCol)
				if length > maxLength {
					maxLength = length
				}
			}
		}

		// Cache the result
		dp[row][col] = maxLength
		return maxLength
	}

	// Compute longest path starting from each cell
	longestPath := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			pathLength := dfs(i, j)
			if pathLength > longestPath {
				longestPath = pathLength
			}
		}
	}

	return longestPath
}