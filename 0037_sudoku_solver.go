// 0037 - Sudoku Solver
// https://leetcode.com/problems/sudoku-solver/
// Hard - Array, Backtracking, Matrix

package leetcode

// SolveSudoku solves a Sudoku puzzle by filling the empty cells.
// Time Complexity: O(9^(n*n)) in worst case, but much better with pruning
// Space Complexity: O(n*n) for recursion stack
func SolveSudoku(board [][]byte) {
	solve(board)
}

// solve is a recursive backtracking function
func solve(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				// Try numbers 1-9
				for num := byte('1'); num <= '9'; num++ {
					if isValidPlacement(board, i, j, num) {
						board[i][j] = num
						if solve(board) {
							return true
						}
						board[i][j] = '.' // Backtrack
					}
				}
				return false // No valid number found
			}
		}
	}
	return true // All cells filled
}

// isValidPlacement checks if placing num at (row, col) is valid
func isValidPlacement(board [][]byte, row, col int, num byte) bool {
	// Check row
	for j := 0; j < 9; j++ {
		if board[row][j] == num {
			return false
		}
	}

	// Check column
	for i := 0; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
	}

	// Check 3x3 sub-box
	boxRow := (row / 3) * 3
	boxCol := (col / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[boxRow+i][boxCol+j] == num {
				return false
			}
		}
	}

	return true
}

// SolveSudokuOptimized uses bitmask optimization
func SolveSudokuOptimized(board [][]byte) {
	rows := make([]int, 9)
	cols := make([]int, 9)
	boxes := make([]int, 9)

	// Initialize bitmasks
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := int(board[i][j] - '0')
				bit := 1 << (num - 1)
				boxIndex := (i/3)*3 + (j / 3)

				rows[i] |= bit
				cols[j] |= bit
				boxes[boxIndex] |= bit
			}
		}
	}

	solveOptimized(board, rows, cols, boxes, 0, 0)
}

// solveOptimized is the recursive function with bitmask optimization
func solveOptimized(board [][]byte, rows, cols, boxes []int, row, col int) bool {
	if row == 9 {
		return true // Reached end of board
	}

	if col == 9 {
		return solveOptimized(board, rows, cols, boxes, row+1, 0)
	}

	if board[row][col] != '.' {
		return solveOptimized(board, rows, cols, boxes, row, col+1)
	}

	boxIndex := (row/3)*3 + (col / 3)
	available := ^(rows[row] | cols[col] | boxes[boxIndex]) & 0x1FF // 9 bits

	for num := 1; num <= 9; num++ {
		bit := 1 << (num - 1)
		if (available & bit) != 0 {
			// Place number
			board[row][col] = byte('0' + num)
			rows[row] |= bit
			cols[col] |= bit
			boxes[boxIndex] |= bit

			if solveOptimized(board, rows, cols, boxes, row, col+1) {
				return true
			}

			// Backtrack
			board[row][col] = '.'
			rows[row] &^= bit
			cols[col] &^= bit
			boxes[boxIndex] &^= bit
		}
	}

	return false
}

// SolveSudokuMRV uses Minimum Remaining Values heuristic
func SolveSudokuMRV(board [][]byte) {
	// Find empty cell with fewest possibilities
	row, col, possibilities := findMRVCell(board)
	if row == -1 {
		return // Board is solved
	}

	// Try each possibility
	for _, num := range possibilities {
		board[row][col] = num
		SolveSudokuMRV(board)
		if isSolved(board) {
			return
		}
		board[row][col] = '.' // Backtrack
	}
}

// findMRVCell finds empty cell with minimum remaining values
func findMRVCell(board [][]byte) (int, int, []byte) {
	bestRow, bestCol := -1, -1
	var bestPossibilities []byte
	minCount := 10 // More than maximum possible (9)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				possibilities := getPossibleNumbers(board, i, j)
				if len(possibilities) < minCount {
					minCount = len(possibilities)
					bestRow, bestCol = i, j
					bestPossibilities = possibilities
					if minCount == 1 {
						return bestRow, bestCol, bestPossibilities
					}
				}
			}
		}
	}

	return bestRow, bestCol, bestPossibilities
}

// getPossibleNumbers returns all possible numbers for a cell
func getPossibleNumbers(board [][]byte, row, col int) []byte {
	possibilities := []byte{}
	used := make([]bool, 10)

	// Check row and column
	for k := 0; k < 9; k++ {
		if board[row][k] != '.' {
			used[board[row][k]-'0'] = true
		}
		if board[k][col] != '.' {
			used[board[k][col]-'0'] = true
		}
	}

	// Check 3x3 sub-box
	boxRow := (row / 3) * 3
	boxCol := (col / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[boxRow+i][boxCol+j] != '.' {
				used[board[boxRow+i][boxCol+j]-'0'] = true
			}
		}
	}

	// Collect unused numbers
	for num := byte(1); num <= 9; num++ {
		if !used[num] {
			possibilities = append(possibilities, '0'+num)
		}
	}

	return possibilities
}

// isSolved checks if board is completely filled
func isSolved(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				return false
			}
		}
	}
	return true
}