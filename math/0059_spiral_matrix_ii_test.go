package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateMatrix(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected [][]int
	}{
		{
			name: "Example 1: n = 3",
			n:    3,
			expected: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			name: "Example 2: n = 1",
			n:    1,
			expected: [][]int{
				{1},
			},
		},
		{
			name: "n = 2",
			n:    2,
			expected: [][]int{
				{1, 2},
				{4, 3},
			},
		},
		{
			name: "n = 4",
			n:    4,
			expected: [][]int{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 16, 15, 6},
				{10, 9, 8, 7},
			},
		},
		{
			name: "n = 5",
			n:    5,
			expected: [][]int{
				{1, 2, 3, 4, 5},
				{16, 17, 18, 19, 6},
				{15, 24, 25, 20, 7},
				{14, 23, 22, 21, 8},
				{13, 12, 11, 10, 9},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GenerateMatrix(tt.n)
			assert.Equal(t, tt.expected, result,
				"GenerateMatrix(%d) = %v, expected %v",
				tt.n, result, tt.expected)
		})
	}
}

func TestGenerateMatrix_EdgeCases(t *testing.T) {
	t.Run("n = 0 returns empty matrix", func(t *testing.T) {
		result := GenerateMatrix(0)
		assert.Empty(t, result)
	})

	t.Run("n = -1 returns empty matrix", func(t *testing.T) {
		result := GenerateMatrix(-1)
		assert.Empty(t, result)
	})

	t.Run("n = 6 produces correct spiral", func(t *testing.T) {
		n := 6
		result := GenerateMatrix(n)

		// Verify dimensions
		assert.Len(t, result, n)
		for i := range result {
			assert.Len(t, result[i], n)
		}

		// Verify all numbers from 1 to n^2 are present exactly once
		seen := make(map[int]bool, n*n)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				num := result[i][j]
				assert.True(t, num >= 1 && num <= n*n,
					"Number %d at position [%d][%d] is out of range 1-%d",
					num, i, j, n*n)
				assert.False(t, seen[num],
					"Duplicate number %d found at position [%d][%d]",
					num, i, j)
				seen[num] = true
			}
		}
		assert.Len(t, seen, n*n, "Should have exactly %d unique numbers", n*n)
	})
}

func BenchmarkGenerateMatrix(b *testing.B) {
	sizes := []int{10, 50, 100, 200}
	for _, n := range sizes {
		b.Run("n="+string(rune(n)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				GenerateMatrix(n)
			}
		})
	}
}