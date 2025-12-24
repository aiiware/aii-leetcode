package leetcode

import (
	"reflect"
	"testing"
)

func TestSolveSudoku(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]byte
		expected [][]byte
	}{
		{
			name: "Easy Sudoku",
			board: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			expected: [][]byte{
				{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
				{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
				{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
				{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
				{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
				{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
				{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
				{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
				{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
			},
		},
		{
			name: "Hard Sudoku",
			board: [][]byte{
				{'.', '.', '9', '7', '4', '8', '.', '.', '.'},
				{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
				{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
				{'.', '.', '7', '.', '.', '.', '2', '4', '.'},
				{'.', '6', '4', '.', '1', '.', '5', '9', '.'},
				{'.', '9', '8', '.', '.', '.', '3', '.', '.'},
				{'.', '.', '.', '8', '.', '3', '.', '2', '.'},
				{'.', '.', '.', '.', '.', '.', '.', '.', '6'},
				{'.', '.', '.', '2', '7', '5', '9', '.', '.'},
			},
			expected: [][]byte{
				{'5', '1', '9', '7', '4', '8', '6', '3', '2'},
				{'7', '8', '3', '6', '5', '2', '4', '1', '9'},
				{'4', '2', '6', '1', '3', '9', '8', '7', '5'},
				{'3', '5', '7', '9', '8', '6', '2', '4', '1'},
				{'2', '6', '4', '3', '1', '7', '5', '9', '8'},
				{'1', '9', '8', '5', '2', '4', '3', '6', '7'},
				{'9', '7', '5', '8', '6', '3', '1', '2', '4'},
				{'8', '3', '2', '4', '9', '1', '7', '5', '6'},
				{'6', '4', '1', '2', '7', '5', '9', '8', '3'},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Test basic backtracking
			board := make([][]byte, len(test.board))
			for i := range test.board {
				board[i] = make([]byte, len(test.board[i]))
				copy(board[i], test.board[i])
			}
			
			SolveSudoku(board)
			if !reflect.DeepEqual(board, test.expected) {
				t.Errorf("SolveSudoku() failed\nGot: %v\nExpected: %v", board, test.expected)
			}

			// Test optimized backtracking
			board = make([][]byte, len(test.board))
			for i := range test.board {
				board[i] = make([]byte, len(test.board[i]))
				copy(board[i], test.board[i])
			}
			
			SolveSudokuOptimized(board)
			if !reflect.DeepEqual(board, test.expected) {
				t.Errorf("SolveSudokuOptimized() failed\nGot: %v\nExpected: %v", board, test.expected)
			}

			// Test MRV heuristic
			board = make([][]byte, len(test.board))
			for i := range test.board {
				board[i] = make([]byte, len(test.board[i]))
				copy(board[i], test.board[i])
			}
			
			SolveSudokuMRV(board)
			if !reflect.DeepEqual(board, test.expected) {
				t.Errorf("SolveSudokuMRV() failed\nGot: %v\nExpected: %v", board, test.expected)
			}
		})
	}
}

func BenchmarkSolveSudoku(b *testing.B) {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a copy for each iteration
		testBoard := make([][]byte, len(board))
		for j := range board {
			testBoard[j] = make([]byte, len(board[j]))
			copy(testBoard[j], board[j])
		}
		SolveSudoku(testBoard)
	}
}

func BenchmarkSolveSudokuOptimized(b *testing.B) {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a copy for each iteration
		testBoard := make([][]byte, len(board))
		for j := range board {
			testBoard[j] = make([]byte, len(board[j]))
			copy(testBoard[j], board[j])
		}
		SolveSudokuOptimized(testBoard)
	}
}