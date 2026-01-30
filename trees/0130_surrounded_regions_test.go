package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]byte
		expected [][]byte
	}{
		{
			name: "Example 1",
			board: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			expected: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		{
			name: "Example 2",
			board: [][]byte{
				{'X'},
			},
			expected: [][]byte{
				{'X'},
			},
		},
		{
			name: "All X's",
			board: [][]byte{
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
			},
			expected: [][]byte{
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
			},
		},
		{
			name: "All O's",
			board: [][]byte{
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
			},
			expected: [][]byte{
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
				{'O', 'O', 'O'},
			},
		},
		{
			name: "Single O in center",
			board: [][]byte{
				{'X', 'X', 'X'},
				{'X', 'O', 'X'},
				{'X', 'X', 'X'},
			},
			expected: [][]byte{
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
				{'X', 'X', 'X'},
			},
		},
		{
			name: "O on border only",
			board: [][]byte{
				{'O', 'X', 'O'},
				{'X', 'X', 'X'},
				{'O', 'X', 'O'},
			},
			expected: [][]byte{
				{'O', 'X', 'O'},
				{'X', 'X', 'X'},
				{'O', 'X', 'O'},
			},
		},
		{
			name: "Empty board",
			board: [][]byte{},
			expected: [][]byte{},
		},
		{
			name: "1x1 O",
			board: [][]byte{
				{'O'},
			},
			expected: [][]byte{
				{'O'},
			},
		},
		{
			name: "Complex pattern",
			board: [][]byte{
				{'X', 'O', 'X', 'O', 'X'},
				{'X', 'O', 'O', 'O', 'X'},
				{'X', 'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X', 'X'},
			},
			expected: [][]byte{
				{'X', 'O', 'X', 'O', 'X'},
				{'X', 'O', 'O', 'O', 'X'},
				{'X', 'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X', 'X'},
			},
		},
		{
			name: "Multiple surrounded regions",
			board: [][]byte{
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'O', 'X'},
				{'X', 'X', 'X', 'O', 'X'},
				{'X', 'O', 'O', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X'},
			},
			expected: [][]byte{
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X'},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of the board to modify
			board := make([][]byte, len(tt.board))
			for i := range tt.board {
				board[i] = make([]byte, len(tt.board[i]))
				copy(board[i], tt.board[i])
			}

			Solve(board)

			assert.Equal(t, tt.expected, board,
				"Solve() modified board incorrectly")
		})
	}
}

func TestSolve_EdgeCases(t *testing.T) {
	t.Run("Large board 200x200", func(t *testing.T) {
		// Create a 200x200 board with border O's and center X's
		m, n := 200, 200
		board := make([][]byte, m)
		for i := range board {
			board[i] = make([]byte, n)
			for j := range board[i] {
				if i == 0 || i == m-1 || j == 0 || j == n-1 {
					board[i][j] = 'O'
				} else {
					board[i][j] = 'X'
				}
			}
		}

		expected := make([][]byte, m)
		for i := range expected {
			expected[i] = make([]byte, n)
			for j := range expected[i] {
				expected[i][j] = board[i][j] // Should remain unchanged
			}
		}

		Solve(board)
		assert.Equal(t, expected, board)
	})

	t.Run("Board with only one row", func(t *testing.T) {
		board := [][]byte{
			{'X', 'O', 'X', 'O', 'X'},
		}
		expected := [][]byte{
			{'X', 'O', 'X', 'O', 'X'},
		}

		Solve(board)
		assert.Equal(t, expected, board)
	})

	t.Run("Board with only one column", func(t *testing.T) {
		board := [][]byte{
			{'X'},
			{'O'},
			{'X'},
			{'O'},
			{'X'},
		}
		expected := [][]byte{
			{'X'},
			{'O'},
			{'X'},
			{'O'},
			{'X'},
		}

		Solve(board)
		assert.Equal(t, expected, board)
	})

	t.Run("Alternating pattern", func(t *testing.T) {
		board := [][]byte{
			{'O', 'X', 'O', 'X', 'O'},
			{'X', 'O', 'X', 'O', 'X'},
			{'O', 'X', 'O', 'X', 'O'},
			{'X', 'O', 'X', 'O', 'X'},
			{'O', 'X', 'O', 'X', 'O'},
		}
		// Only O's on border should remain, interior O's should be flipped
		expected := [][]byte{
			{'O', 'X', 'O', 'X', 'O'},
			{'X', 'X', 'X', 'X', 'X'},
			{'O', 'X', 'X', 'X', 'O'},
			{'X', 'X', 'X', 'X', 'X'},
			{'O', 'X', 'O', 'X', 'O'},
		}

		Solve(board)
		assert.Equal(t, expected, board)
	})
}

func BenchmarkSolve(b *testing.B) {
	// Create a 100x100 board for benchmarking
	m, n := 100, 100
	board := make([][]byte, m)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			if (i+j)%3 == 0 {
				board[i][j] = 'O'
			} else {
				board[i][j] = 'X'
			}
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a copy for each benchmark iteration
		testBoard := make([][]byte, m)
		for idx := range board {
			testBoard[idx] = make([]byte, n)
			copy(testBoard[idx], board[idx])
		}
		Solve(testBoard)
	}
}