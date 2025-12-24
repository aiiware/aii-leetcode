// 0036 - Valid Sudoku
// https://leetcode.com/problems/valid-sudoku/
// Medium - Array, Hash Table, Matrix

package leetcode

// IsValidSudoku checks if a 9x9 Sudoku board is valid.
// Time Complexity: O(1) since board is fixed 9x9, or O(n²) for n x n board
// Space Complexity: O(1) or O(n) for tracking seen numbers
func IsValidSudoku(board [][]byte) bool {
	// Check rows
	for i := 0; i < 9; i++ {
		seen := make([]bool, 10)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := int(board[i][j] - '0')
				if num < 1 || num > 9 || seen[num] {
					return false
				}
				seen[num] = true
			}
		}
	}

	// Check columns
	for j := 0; j < 9; j++ {
		seen := make([]bool, 10)
		for i := 0; i < 9; i++ {
			if board[i][j] != '.' {
				num := int(board[i][j] - '0')
				if num < 1 || num > 9 || seen[num] {
					return false
				}
				seen[num] = true
			}
		}
	}

	// Check 3x3 sub-boxes
	for box := 0; box < 9; box++ {
		seen := make([]bool, 10)
		rowStart := (box / 3) * 3
		colStart := (box % 3) * 3
		
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				cell := board[rowStart+i][colStart+j]
				if cell != '.' {
					num := int(cell - '0')
					if num < 1 || num > 9 || seen[num] {
						return false
					}
					seen[num] = true
				}
			}
		}
	}

	return true
}

// IsValidSudokuOptimized uses bitmask for more efficient checking
func IsValidSudokuOptimized(board [][]byte) bool {
	rows := make([]int, 9)
	cols := make([]int, 9)
	boxes := make([]int, 9)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}

			num := int(board[i][j] - '0')
			if num < 1 || num > 9 {
				return false
			}

			bit := 1 << (num - 1)
			boxIndex := (i/3)*3 + (j / 3)

			// Check if this number has been seen in row, column, or box
			if (rows[i]&bit) != 0 || (cols[j]&bit) != 0 || (boxes[boxIndex]&bit) != 0 {
				return false
			}

			// Mark as seen
			rows[i] |= bit
			cols[j] |= bit
			boxes[boxIndex] |= bit
		}
	}

	return true
}

// IsValidSudokuGeneric works for any n x n board where n is a perfect square
func IsValidSudokuGeneric(board [][]byte) bool {
	n := len(board)
	if n == 0 {
		return true
	}

	subSize := 3 // Standard Sudoku sub-box size
	if n == 4 {
		subSize = 2 // For 4x4 board
	} else if n == 16 {
		subSize = 4 // For 16x16 board
	}

	// Check rows and columns
	for i := 0; i < n; i++ {
		rowSeen := make([]bool, n+1)
		colSeen := make([]bool, n+1)
		
		for j := 0; j < n; j++ {
			// Check row
			if board[i][j] != '.' {
				num := int(board[i][j] - '0')
				if num < 1 || num > n || rowSeen[num] {
					return false
				}
				rowSeen[num] = true
			}

			// Check column
			if board[j][i] != '.' {
				num := int(board[j][i] - '0')
				if num < 1 || num > n || colSeen[num] {
					return false
				}
				colSeen[num] = true
			}
		}
	}

	// Check sub-boxes
	for boxRow := 0; boxRow < n; boxRow += subSize {
		for boxCol := 0; boxCol < n; boxCol += subSize {
			seen := make([]bool, n+1)
			
			for i := 0; i < subSize; i++ {
				for j := 0; j < subSize; j++ {
					cell := board[boxRow+i][boxCol+j]
					if cell != '.' {
						num := int(cell - '0')
						if num < 1 || num > n || seen[num] {
							return false
						}
						seen[num] = true
					}
				}
			}
		}
	}

	return true
}