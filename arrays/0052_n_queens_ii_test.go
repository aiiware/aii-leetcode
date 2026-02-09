package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotalNQueens(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "n = 1 (single queen)",
			n:        1,
			expected: 1,
		},
		{
			name:     "n = 2 (no solution)",
			n:        2,
			expected: 0,
		},
		{
			name:     "n = 3 (no solution)",
			n:        3,
			expected: 0,
		},
		{
			name:     "n = 4 (classic example)",
			n:        4,
			expected: 2,
		},
		{
			name:     "n = 5",
			n:        5,
			expected: 10,
		},
		{
			name:     "n = 6",
			n:        6,
			expected: 4,
		},
		{
			name:     "n = 7",
			n:        7,
			expected: 40,
		},
		{
			name:     "n = 8 (standard chessboard)",
			n:        8,
			expected: 92,
		},
		{
			name:     "n = 9 (maximum for LeetCode constraints)",
			n:        9,
			expected: 352,
		},
		{
			name:     "n = 0 (edge case)",
			n:        0,
			expected: 0,
		},
		{
			name:     "n = -1 (invalid input)",
			n:        -1,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TotalNQueens(tt.n)
			assert.Equal(t, tt.expected, result,
				"TotalNQueens(%d) = %d, expected %d",
				tt.n, result, tt.expected)
		})
	}
}

func TestTotalNQueensDFS(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "n = 1",
			n:        1,
			expected: 1,
		},
		{
			name:     "n = 4",
			n:        4,
			expected: 2,
		},
		{
			name:     "n = 8",
			n:        8,
			expected: 92,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TotalNQueensDFS(tt.n)
			assert.Equal(t, tt.expected, result,
				"TotalNQueensDFS(%d) = %d, expected %d",
				tt.n, result, tt.expected)
		})
	}
}

func TestBothImplementationsProduceSameResults(t *testing.T) {
	// Test that both implementations produce the same results for n = 1 to 9
	for n := 1; n <= 9; n++ {
		bitResult := TotalNQueens(n)
		dfsResult := TotalNQueensDFS(n)
		assert.Equal(t, bitResult, dfsResult,
			"For n = %d, TotalNQueens = %d, TotalNQueensDFS = %d (should be equal)",
			n, bitResult, dfsResult)
	}
}

func TestTotalNQueens_EdgeCases(t *testing.T) {
	t.Run("Large n (beyond constraints but should handle gracefully)", func(t *testing.T) {
		// Note: n > 9 is beyond LeetCode constraints but our implementation should handle it
		// The number of solutions grows extremely fast (n=10 has 724 solutions)
		result := TotalNQueens(10)
		assert.Equal(t, 724, result, "n=10 should have 724 solutions")
	})

	t.Run("Very large n (test recursion depth)", func(t *testing.T) {
		// This tests that our implementation doesn't crash on larger inputs
		// n=12 has 14200 solutions
		result := TotalNQueens(12)
		assert.Equal(t, 14200, result, "n=12 should have 14200 solutions")
	})
}

func BenchmarkTotalNQueens(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{"n=4", 4},
		{"n=8", 8},
		{"n=9", 9},
		{"n=10", 10},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				TotalNQueens(bm.n)
			}
		})
	}
}

func BenchmarkTotalNQueensDFS(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{"n=4", 4},
		{"n=8", 8},
		{"n=9", 9},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				TotalNQueensDFS(bm.n)
			}
		})
	}
}

func BenchmarkComparison(b *testing.B) {
	// Compare performance of both implementations
	n := 9

	b.Run("BitManipulation", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			TotalNQueens(n)
		}
	})

	b.Run("DFS", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			TotalNQueensDFS(n)
		}
	})
}
