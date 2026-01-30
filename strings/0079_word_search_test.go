package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExist(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]byte
		word     string
		expected bool
	}{
		{
			name: "Example 1: ABCCED",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "ABCCED",
			expected: true,
		},
		{
			name: "Example 2: SEE",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "SEE",
			expected: true,
		},
		{
			name: "Example 3: ABCB",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "ABCB",
			expected: false,
		},
		{
			name: "Single cell board with matching word",
			board: [][]byte{
				{'A'},
			},
			word:     "A",
			expected: true,
		},
		{
			name: "Single cell board with non-matching word",
			board: [][]byte{
				{'A'},
			},
			word:     "B",
			expected: false,
		},
		{
			name: "Word longer than board cells",
			board: [][]byte{
				{'A', 'B'},
				{'C', 'D'},
			},
			word:     "ABCDE",
			expected: false,
		},
		{
			name: "Word uses same cell twice",
			board: [][]byte{
				{'A', 'B'},
				{'C', 'D'},
			},
			word:     "ABA",
			expected: false,
		},
		{
			name: "Word in straight line horizontal",
			board: [][]byte{
				{'A', 'B', 'C'},
				{'D', 'E', 'F'},
				{'G', 'H', 'I'},
			},
			word:     "ABC",
			expected: true,
		},
		{
			name: "Word in straight line vertical",
			board: [][]byte{
				{'A', 'B', 'C'},
				{'D', 'E', 'F'},
				{'G', 'H', 'I'},
			},
			word:     "ADG",
			expected: true,
		},
		{
			name: "Word in L shape",
			board: [][]byte{
				{'A', 'B', 'C'},
				{'D', 'E', 'F'},
				{'G', 'H', 'I'},
			},
			word:     "ABE",
			expected: true,
		},
		{
			name: "Word with backtracking needed",
			board: [][]byte{
				{'A', 'A', 'A'},
				{'A', 'B', 'A'},
				{'A', 'A', 'A'},
			},
			word:     "AB",
			expected: true,
		},
		{
			name: "Empty word should return true",
			board: [][]byte{
				{'A', 'B'},
				{'C', 'D'},
			},
			word:     "",
			expected: true,
		},
		{
			name:     "Empty board with non-empty word",
			board:    [][]byte{},
			word:     "A",
			expected: false,
		},
		{
			name: "Empty board with empty word",
			board:    [][]byte{},
			word:     "",
			expected: true,
		},
		{
			name: "Board with duplicate letters",
			board: [][]byte{
				{'A', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'A'},
				{'A', 'A', 'A', 'A'},
			},
			word:     "AAAAAA",
			expected: true,
		},
		{
			name: "Case sensitive test",
			board: [][]byte{
				{'a', 'b'},
				{'c', 'd'},
			},
			word:     "ab",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Exist(tt.board, tt.word)
			assert.Equal(t, tt.expected, result,
				"Exist(board, %q) = %v, expected %v",
				tt.word, result, tt.expected)
		})
	}
}

func TestExistOptimized(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]byte
		word     string
		expected bool
	}{
		{
			name: "Example 1: ABCCED",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "ABCCED",
			expected: true,
		},
		{
			name: "Example 2: SEE",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word:     "SEE",
			expected: true,
		},
		{
			name: "Word with backtracking",
			board: [][]byte{
				{'A', 'A', 'A'},
				{'A', 'B', 'A'},
				{'A', 'A', 'A'},
			},
			word:     "AB",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of the board since ExistOptimized modifies it
			boardCopy := make([][]byte, len(tt.board))
			for i := range tt.board {
				boardCopy[i] = make([]byte, len(tt.board[i]))
				copy(boardCopy[i], tt.board[i])
			}
			
			result := ExistOptimized(boardCopy, tt.word)
			assert.Equal(t, tt.expected, result,
				"ExistOptimized(board, %q) = %v, expected %v",
				tt.word, result, tt.expected)
		})
	}
}

func TestExistBFS(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]byte
		word     string
		expected bool
	}{
		{
			name: "Simple horizontal word",
			board: [][]byte{
				{'A', 'B', 'C'},
				{'D', 'E', 'F'},
			},
			word:     "ABC",
			expected: true,
		},
		{
			name: "Word not found",
			board: [][]byte{
				{'A', 'B', 'C'},
				{'D', 'E', 'F'},
			},
			word:     "ABD",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExistBFS(tt.board, tt.word)
			assert.Equal(t, tt.expected, result,
				"ExistBFS(board, %q) = %v, expected %v",
				tt.word, result, tt.expected)
		})
	}
}

func TestExistEarlyPruning(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]byte
		word     string
		expected bool
	}{
		{
			name: "Board missing required character",
			board: [][]byte{
				{'A', 'B', 'C'},
				{'D', 'E', 'F'},
			},
			word:     "ABZ",
			expected: false,
		},
		{
			name: "Board has all characters",
			board: [][]byte{
				{'A', 'B', 'C'},
				{'D', 'E', 'F'},
			},
			word:     "ABC",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of the board since ExistEarlyPruning modifies it
			boardCopy := make([][]byte, len(tt.board))
			for i := range tt.board {
				boardCopy[i] = make([]byte, len(tt.board[i]))
				copy(boardCopy[i], tt.board[i])
			}
			
			result := ExistEarlyPruning(boardCopy, tt.word)
			assert.Equal(t, tt.expected, result,
				"ExistEarlyPruning(board, %q) = %v, expected %v",
				tt.word, result, tt.expected)
		})
	}
}

func TestExistDirectional(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]byte
		word     string
		expected bool
	}{
		{
			name: "Word found from multiple starting points",
			board: [][]byte{
				{'A', 'A', 'A'},
				{'A', 'B', 'A'},
				{'A', 'A', 'A'},
			},
			word:     "AB",
			expected: true,
		},
		{
			name: "Word not found",
			board: [][]byte{
				{'A', 'B', 'C'},
				{'D', 'E', 'F'},
			},
			word:     "XYZ",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of the board since ExistDirectional modifies it
			boardCopy := make([][]byte, len(tt.board))
			for i := range tt.board {
				boardCopy[i] = make([]byte, len(tt.board[i]))
				copy(boardCopy[i], tt.board[i])
			}
			
			result := ExistDirectional(boardCopy, tt.word)
			assert.Equal(t, tt.expected, result,
				"ExistDirectional(board, %q) = %v, expected %v",
				tt.word, result, tt.expected)
		})
	}
}

func TestExist_Consistency(t *testing.T) {
	// Test that all implementations return the same result
	testCases := []struct {
		name  string
		board [][]byte
		word  string
	}{
		{
			name: "Simple case",
			board: [][]byte{
				{'A', 'B'},
				{'C', 'D'},
			},
			word: "AB",
		},
		{
			name: "Complex case",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "ABCCED",
		},
		{
			name: "Not found case",
			board: [][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			word: "ABCB",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test standard implementation
			result1 := Exist(tc.board, tc.word)
			
			// Test optimized implementation (needs copy)
			boardCopy2 := make([][]byte, len(tc.board))
			for i := range tc.board {
				boardCopy2[i] = make([]byte, len(tc.board[i]))
				copy(boardCopy2[i], tc.board[i])
			}
			result2 := ExistOptimized(boardCopy2, tc.word)
			
			// Test BFS implementation
			result3 := ExistBFS(tc.board, tc.word)
			
			// Test early pruning implementation (needs copy)
			boardCopy4 := make([][]byte, len(tc.board))
			for i := range tc.board {
				boardCopy4[i] = make([]byte, len(tc.board[i]))
				copy(boardCopy4[i], tc.board[i])
			}
			result4 := ExistEarlyPruning(boardCopy4, tc.word)
			
			// Test directional implementation (needs copy)
			boardCopy5 := make([][]byte, len(tc.board))
			for i := range tc.board {
				boardCopy5[i] = make([]byte, len(tc.board[i]))
				copy(boardCopy5[i], tc.board[i])
			}
			result5 := ExistDirectional(boardCopy5, tc.word)
			
			// All implementations should agree
			assert.Equal(t, result1, result2, "Exist and ExistOptimized should match")
			assert.Equal(t, result1, result3, "Exist and ExistBFS should match")
			assert.Equal(t, result1, result4, "Exist and ExistEarlyPruning should match")
			assert.Equal(t, result1, result5, "Exist and ExistDirectional should match")
		})
	}
}

func TestExist_EdgeCases(t *testing.T) {
	t.Run("Empty word returns true", func(t *testing.T) {
		board := [][]byte{{'A', 'B'}, {'C', 'D'}}
		assert.True(t, Exist(board, ""))
		assert.True(t, ExistOptimized(board, ""))
		assert.True(t, ExistBFS(board, ""))
		assert.True(t, ExistEarlyPruning(board, ""))
		assert.True(t, ExistDirectional(board, ""))
	})

	t.Run("Empty board returns false for non-empty word", func(t *testing.T) {
		board := [][]byte{}
		assert.False(t, Exist(board, "A"))
		assert.False(t, ExistOptimized(board, "A"))
		assert.False(t, ExistBFS(board, "A"))
		assert.False(t, ExistEarlyPruning(board, "A"))
		assert.False(t, ExistDirectional(board, "A"))
	})

	t.Run("Empty board returns true for empty word", func(t *testing.T) {
		board := [][]byte{}
		assert.True(t, Exist(board, ""))
		assert.True(t, ExistOptimized(board, ""))
		assert.True(t, ExistBFS(board, ""))
		assert.True(t, ExistEarlyPruning(board, ""))
		assert.True(t, ExistDirectional(board, ""))
	})

	t.Run("Word longer than total cells returns false", func(t *testing.T) {
		board := [][]byte{{'A'}}
		assert.False(t, Exist(board, "AB"))
		assert.False(t, ExistOptimized(board, "AB"))
		assert.False(t, ExistBFS(board, "AB"))
		assert.False(t, ExistEarlyPruning(board, "AB"))
		assert.False(t, ExistDirectional(board, "AB"))
	})
}

func BenchmarkExist(b *testing.B) {
	// Create a test board
	board := [][]byte{
		{'A', 'B', 'C', 'D', 'E'},
		{'F', 'G', 'H', 'I', 'J'},
		{'K', 'L', 'M', 'N', 'O'},
		{'P', 'Q', 'R', 'S', 'T'},
		{'U', 'V', 'W', 'X', 'Y'},
	}
	
	testCases := []struct {
		name string
		word string
	}{
		{"Short word", "ABC"},
		{"Medium word", "ABCDE"},
		{"Long word", "ABCDEFGHIJ"},
		{"Not found", "XYZ"},
		{"Complex path", "AGMSRQVW"},
	}

	for _, tc := range testCases {
		b.Run("Standard_"+tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Exist(board, tc.word)
			}
		})
		
		b.Run("Optimized_"+tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Create a fresh copy for each iteration
				boardCopy := make([][]byte, len(board))
				for j := range board {
					boardCopy[j] = make([]byte, len(board[j]))
					copy(boardCopy[j], board[j])
				}
				ExistOptimized(boardCopy, tc.word)
			}
		})
		
		b.Run("BFS_"+tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ExistBFS(board, tc.word)
			}
		})
		
		b.Run("EarlyPruning_"+tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Create a fresh copy for each iteration
				boardCopy := make([][]byte, len(board))
				for j := range board {
					boardCopy[j] = make([]byte, len(board[j]))
					copy(boardCopy[j], board[j])
				}
				ExistEarlyPruning(boardCopy, tc.word)
			}
		})
		
		b.Run("Directional_"+tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Create a fresh copy for each iteration
				boardCopy := make([][]byte, len(board))
				for j := range board {
					boardCopy[j] = make([]byte, len(board[j]))
					copy(boardCopy[j], board[j])
				}
				ExistDirectional(boardCopy, tc.word)
			}
		})
	}
}