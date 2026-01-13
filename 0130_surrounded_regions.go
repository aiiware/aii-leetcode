package leetcode

// Solve solves LeetCode problem 0130: Surrounded Regions
// Difficulty: Medium
// Tags: Array, Depth-First Search, Breadth-First Search, Union Find, Matrix
//
// Given an m x n matrix board containing 'X' and 'O', capture all regions that are 4-directionally
// surrounded by 'X'.
//
// A region is captured by flipping all 'O's into 'X's in that surrounded region.
//
// Example 1:
// Input: board = [["X","X","X","X"],["X","O","O","X"],["X","X","O","X"],["X","O","X","X"]]
// Output: [["X","X","X","X"],["X","X","X","X"],["X","X","X","X"],["X","O","X","X"]]
// Explanation:
// Surrounded regions should not be on the border, which means that any 'O' on the border of the board
// are not flipped to 'X'. Any 'O' that is not on the border and it is not connected to an 'O' on the
// border will be flipped to 'X'. Two cells are connected if they are adjacent cells connected
// horizontally or vertically.
//
// Example 2:
// Input: board = [["X"]]
// Output: [["X"]]
//
// Constraints:
// m == board.length
// n == board[i].length
// 1 <= m, n <= 200
// board[i][j] is 'X' or 'O'.
//
// Time complexity: O(m * n), Space complexity: O(m * n) in worst case for recursion stack
func Solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}

	m, n := len(board), len(board[0])

	// Mark all 'O's connected to border with temporary marker 'T'
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != 'O' {
			return
		}

		// Mark as temporary (connected to border)
		board[i][j] = 'T'

		// Explore 4 directions
		dfs(i-1, j) // up
		dfs(i+1, j) // down
		dfs(i, j-1) // left
		dfs(i, j+1) // right
	}

	// Mark 'O's on border and connected to border
	for i := 0; i < m; i++ {
		// First column
		if board[i][0] == 'O' {
			dfs(i, 0)
		}
		// Last column
		if board[i][n-1] == 'O' {
			dfs(i, n-1)
		}
	}

	for j := 0; j < n; j++ {
		// First row
		if board[0][j] == 'O' {
			dfs(0, j)
		}
		// Last row
		if board[m-1][j] == 'O' {
			dfs(m-1, j)
		}
	}

	// Process the board
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'O' {
				// Surrounded region, flip to 'X'
				board[i][j] = 'X'
			} else if board[i][j] == 'T' {
				// Restore border-connected 'O's
				board[i][j] = 'O'
			}
		}
	}
}