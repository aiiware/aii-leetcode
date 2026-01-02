// 0051 - N-Queens
// https://leetcode.com/problems/n-queens/
// Hard - Array, Backtracking

package leetcode

// SolveNQueens returns all distinct solutions to the n-queens puzzle.
// Each solution contains a distinct board configuration of the n-queens' placement,
// where 'Q' and '.' indicate a queen and an empty space respectively.
// Time Complexity: O(n!) in worst case
// Space Complexity: O(n^2) for storing boards
func SolveNQueens(n int) [][]string {
	if n <= 0 {
		return [][]string{}
	}

	// Initialize board with all '.'
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	result := [][]string{}
	nQueensBacktrack(board, 0, &result)
	return result
}

// nQueensBacktrack is the recursive backtracking function for N-Queens
func nQueensBacktrack(board [][]byte, row int, result *[][]string) {
	n := len(board)
	if row == n {
		// Found a valid solution, convert board to []string format
		solution := make([]string, n)
		for i := 0; i < n; i++ {
			solution[i] = string(board[i])
		}
		*result = append(*result, solution)
		return
	}

	// Try placing queen in each column of current row
	for col := 0; col < n; col++ {
		if isValidQueenPlacement(board, row, col) {
			// Place queen
			board[row][col] = 'Q'
			// Recurse to next row
			nQueensBacktrack(board, row+1, result)
			// Backtrack
			board[row][col] = '.'
		}
	}
}

// isValidQueenPlacement checks if placing a queen at (row, col) is valid
func isValidQueenPlacement(board [][]byte, row, col int) bool {
	n := len(board)

	// Check column (above rows only, since we're placing row by row)
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// Check upper-left diagonal
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// Check upper-right diagonal
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	return true
}

// SolveNQueensOptimized uses bitmask optimization for better performance
func SolveNQueensOptimized(n int) [][]string {
	if n <= 0 {
		return [][]string{}
	}

	result := [][]string{}
	// Bitmask approach: use integers to represent occupied columns and diagonals
	nQueensBacktrackOptimized(n, 0, 0, 0, 0, []int{}, &result)
	return result
}

// nQueensBacktrackOptimized uses bitmask to track occupied positions
// cols: columns with queens
// diag1: main diagonals (row - col = constant)
// diag2: anti-diagonals (row + col = constant)
func nQueensBacktrackOptimized(n, row, cols, diag1, diag2 int, queens []int, result *[][]string) {
	if row == n {
		// Found a solution, build board from queens positions
		board := make([]string, n)
		for i := 0; i < n; i++ {
			rowStr := make([]byte, n)
			for j := 0; j < n; j++ {
				rowStr[j] = '.'
			}
			rowStr[queens[i]] = 'Q'
			board[i] = string(rowStr)
		}
		*result = append(*result, board)
		return
	}

	// Get available positions (bits that are 0)
	// cols | diag1 | diag2 gives all attacked positions
	// ^ gives complement, & mask keeps only n bits
	available := ^(cols | diag1 | diag2) & ((1 << n) - 1)

	// Try each available column
	for available != 0 {
		// Get least significant set bit
		pos := available & -available
		// Get column index from bit position
		col := 0
		for (pos>>col)&1 == 0 {
			col++
		}

		// Place queen
		queens = append(queens, col)

		// Update masks for next row
		// cols: mark this column as occupied
		// diag1: shift left for next row (row-col constant)
		// diag2: shift right for next row (row+col constant)
		newCols := cols | pos
		newDiag1 := (diag1 | pos) << 1
		newDiag2 := (diag2 | pos) >> 1

		nQueensBacktrackOptimized(n, row+1, newCols, newDiag1, newDiag2, queens, result)

		// Backtrack
		queens = queens[:len(queens)-1]

		// Remove this position from available
		available &= available - 1
	}
}

// SolveNQueensCount returns only the count of solutions (LeetCode 0052)
func SolveNQueensCount(n int) int {
	if n <= 0 {
		return 0
	}

	count := 0
	nQueensBacktrackCount(n, 0, 0, 0, 0, &count)
	return count
}

// nQueensBacktrackCount counts solutions using bitmask
func nQueensBacktrackCount(n, row, cols, diag1, diag2 int, count *int) {
	if row == n {
		*count++
		return
	}

	available := ^(cols | diag1 | diag2) & ((1 << n) - 1)

	for available != 0 {
		pos := available & -available
		newCols := cols | pos
		newDiag1 := (diag1 | pos) << 1
		newDiag2 := (diag2 | pos) >> 1

		nQueensBacktrackCount(n, row+1, newCols, newDiag1, newDiag2, count)

		available &= available - 1
	}
}