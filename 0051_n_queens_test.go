package leetcode

import (
	"fmt"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	tests := []struct {
		n        int
		expected int // number of solutions
	}{
		{0, 0},
		{1, 1},
		{2, 0},
		{3, 0},
		{4, 2},
		{5, 10},
		{6, 4},
		{7, 40},
		{8, 92},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d", tt.n), func(t *testing.T) {
			solutions := SolveNQueens(tt.n)
			if len(solutions) != tt.expected {
				t.Errorf("SolveNQueens(%d) = %d solutions, want %d", tt.n, len(solutions), tt.expected)
			}

			// Verify each solution is valid
			for i, solution := range solutions {
				if !isValidNQueensSolution(solution) {
					t.Errorf("Solution %d for n=%d is invalid:\n%s", i, tt.n, formatSolution(solution))
				}
			}
		})
	}
}

func TestSolveNQueensOptimized(t *testing.T) {
	tests := []struct {
		n        int
		expected int // number of solutions
	}{
		{0, 0},
		{1, 1},
		{2, 0},
		{3, 0},
		{4, 2},
		{5, 10},
		{6, 4},
		{7, 40},
		{8, 92},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d", tt.n), func(t *testing.T) {
			solutions := SolveNQueensOptimized(tt.n)
			if len(solutions) != tt.expected {
				t.Errorf("SolveNQueensOptimized(%d) = %d solutions, want %d", tt.n, len(solutions), tt.expected)
			}

			// Verify each solution is valid
			for i, solution := range solutions {
				if !isValidNQueensSolution(solution) {
					t.Errorf("Solution %d for n=%d is invalid:\n%s", i, tt.n, formatSolution(solution))
				}
			}
		})
	}
}

func TestSolveNQueensCount(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 0},
		{3, 0},
		{4, 2},
		{5, 10},
		{6, 4},
		{7, 40},
		{8, 92},
		{9, 352},
		{10, 724},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d", tt.n), func(t *testing.T) {
			count := SolveNQueensCount(tt.n)
			if count != tt.expected {
				t.Errorf("SolveNQueensCount(%d) = %d, want %d", tt.n, count, tt.expected)
			}
		})
	}
}

func TestBothNQueensImplementationsMatch(t *testing.T) {
	// Test that both implementations produce the same results
	for n := 1; n <= 8; n++ {
		t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
			solutions1 := SolveNQueens(n)
			solutions2 := SolveNQueensOptimized(n)

			if len(solutions1) != len(solutions2) {
				t.Errorf("n=%d: Basic has %d solutions, Optimized has %d", n, len(solutions1), len(solutions2))
				return
			}

			// Convert solutions to a map for comparison
			solutionsMap := make(map[string]bool)
			for _, sol := range solutions1 {
				solutionsMap[formatSolution(sol)] = true
			}

			for _, sol := range solutions2 {
				key := formatSolution(sol)
				if !solutionsMap[key] {
					t.Errorf("n=%d: Solution from optimized not found in basic:\n%s", n, key)
				}
			}
		})
	}
}

func TestNQueensSolutions(t *testing.T) {
	// Test specific known solutions
	testCases := []struct {
		n        int
		solution []string
	}{
		{
			n: 4,
			solution: []string{
				".Q..",
				"...Q",
				"Q...",
				"..Q.",
			},
		},
		{
			n: 4,
			solution: []string{
				"..Q.",
				"Q...",
				"...Q",
				".Q..",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("n=%d_solution", tc.n), func(t *testing.T) {
			solutions := SolveNQueens(tc.n)
			found := false
			for _, sol := range solutions {
				if solutionsEqual(sol, tc.solution) {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Expected solution not found for n=%d:\n%s", tc.n, formatSolution(tc.solution))
			}
		})
	}
}

// Helper functions

func isValidNQueensSolution(board []string) bool {
	n := len(board)
	if n == 0 {
		return true
	}

	// Count queens and check board dimensions
	queenCount := 0
	for i := 0; i < n; i++ {
		if len(board[i]) != n {
			return false
		}
		for j := 0; j < n; j++ {
			if board[i][j] == 'Q' {
				queenCount++
			} else if board[i][j] != '.' {
				return false // Invalid character
			}
		}
	}

	if queenCount != n {
		return false // Wrong number of queens
	}

	// Check each queen doesn't attack another
	queens := make([][2]int, 0, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'Q' {
				queens = append(queens, [2]int{i, j})
			}
		}
	}

	// Check all pairs of queens
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			q1, q2 := queens[i], queens[j]
			// Same row, column, or diagonal?
			if q1[0] == q2[0] || q1[1] == q2[1] || absNQueens(q1[0]-q2[0]) == absNQueens(q1[1]-q2[1]) {
				return false
			}
		}
	}

	return true
}

func solutionsEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func formatSolution(solution []string) string {
	result := ""
	for _, row := range solution {
		result += row + "\n"
	}
	return result
}

func absNQueens(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func BenchmarkSolveNQueens(b *testing.B) {
	for n := 4; n <= 8; n++ {
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SolveNQueens(n)
			}
		})
	}
}

func BenchmarkSolveNQueensOptimized(b *testing.B) {
	for n := 4; n <= 12; n++ {
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SolveNQueensOptimized(n)
			}
		})
	}
}

func BenchmarkSolveNQueensCount(b *testing.B) {
	for n := 4; n <= 12; n++ {
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SolveNQueensCount(n)
			}
		})
	}
}