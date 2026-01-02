package leetcode

// TotalNQueens solves LeetCode problem 52: N-Queens II
// Difficulty: Hard
// Tags: Backtracking, Bit Manipulation
//
// The n-queens puzzle is the problem of placing n queens on an n×n chessboard
// such that no two queens attack each other.
//
// Given an integer n, return the number of distinct solutions to the n-queens puzzle.
//
// Example 1:
// Input: n = 4
// Output: 2
// Explanation: There are two distinct solutions to the 4-queens puzzle as shown.
//
// Example 2:
// Input: n = 1
// Output: 1
//
// Constraints:
// 1 <= n <= 9
//
// Time complexity: O(n!), Space complexity: O(n)
func TotalNQueens(n int) int {
	if n <= 0 {
		return 0
	}

	// Use bit manipulation for faster checking
	// We'll track columns, diagonals, and anti-diagonals
	var count int
	var backtrack func(row, cols, diags, antiDiags int)
	
	backtrack = func(row, cols, diags, antiDiags int) {
		// If all queens are placed, we found a solution
		if row == n {
			count++
			return
		}

		// Get available positions for this row
		// ~(cols | diags | antiDiags) gives us positions where we CAN place a queen
		// & ((1 << n) - 1) masks to only n bits
		available := ^(cols | diags | antiDiags) & ((1 << n) - 1)

		// Try each available position
		for available > 0 {
			// Get the rightmost available position
			pos := available & -available
			
			// Place queen at this position
			// For next row:
			// - cols: mark this column as occupied
			// - diags: shift left by 1 (diagonal attacks move down-left)
			// - antiDiags: shift right by 1 (anti-diagonal attacks move down-right)
			backtrack(
				row+1,
				cols|pos,
				(diags|pos)<<1,
				(antiDiags|pos)>>1,
			)
			
			// Remove this position from available
			available &^= pos
		}
	}

	backtrack(0, 0, 0, 0)
	return count
}

// TotalNQueensDFS is an alternative implementation using DFS with arrays
// This is easier to understand but slightly slower than bit manipulation
func TotalNQueensDFS(n int) int {
	if n <= 0 {
		return 0
	}

	count := 0
	// Track which columns, diagonals, and anti-diagonals are occupied
	cols := make([]bool, n)
	diags := make([]bool, 2*n-1)    // diagonals: row - col + (n-1)
	antiDiags := make([]bool, 2*n-1) // anti-diagonals: row + col

	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			count++
			return
		}

		for col := 0; col < n; col++ {
			diagIdx := row - col + (n - 1)
			antiDiagIdx := row + col

			// Check if position is safe
			if !cols[col] && !diags[diagIdx] && !antiDiags[antiDiagIdx] {
				// Place queen
				cols[col] = true
				diags[diagIdx] = true
				antiDiags[antiDiagIdx] = true

				// Move to next row
				dfs(row + 1)

				// Remove queen (backtrack)
				cols[col] = false
				diags[diagIdx] = false
				antiDiags[antiDiagIdx] = false
			}
		}
	}

	dfs(0)
	return count
}