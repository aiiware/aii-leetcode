package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected []int
	}{
		{
			name: "Example 1",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "Example 2",
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
			},
			expected: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			name:     "Empty matrix",
			matrix:   [][]int{},
			expected: []int{},
		},
		{
			name: "Single row",
			matrix: [][]int{
				{1, 2, 3, 4, 5},
			},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Single column",
			matrix: [][]int{
				{1},
				{2},
				{3},
				{4},
				{5},
			},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name: "2x2 matrix",
			matrix: [][]int{
				{1, 2},
				{3, 4},
			},
			expected: []int{1, 2, 4, 3},
		},
		{
			name: "1x1 matrix",
			matrix: [][]int{
				{42},
			},
			expected: []int{42},
		},
		{
			name: "4x3 matrix",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
			},
			expected: []int{1, 2, 3, 6, 9, 12, 11, 10, 7, 4, 5, 8},
		},
		{
			name: "3x5 matrix",
			matrix: [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
			},
			expected: []int{1, 2, 3, 4, 5, 10, 15, 14, 13, 12, 11, 6, 7, 8, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SpiralOrder(tt.matrix)
			assert.Equal(t, tt.expected, result,
				"SpiralOrder(%v) = %v, expected %v",
				tt.matrix, result, tt.expected)
		})
	}
}

func TestSpiralOrder_EdgeCases(t *testing.T) {
	t.Run("Empty row matrix", func(t *testing.T) {
		result := SpiralOrder([][]int{{}})
		assert.Equal(t, []int{}, result)
	})

	t.Run("Large matrix", func(t *testing.T) {
		// Create a 5x5 matrix with sequential numbers
		matrix := make([][]int, 5)
		for i := 0; i < 5; i++ {
			matrix[i] = make([]int, 5)
			for j := 0; j < 5; j++ {
				matrix[i][j] = i*5 + j + 1
			}
		}

		expected := []int{1, 2, 3, 4, 5, 10, 15, 20, 25, 24, 23, 22, 21, 16, 11, 6, 7, 8, 9, 14, 19, 18, 17, 12, 13}
		result := SpiralOrder(matrix)
		assert.Equal(t, expected, result)
	})

	t.Run("Rectangular matrix 2x4", func(t *testing.T) {
		matrix := [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
		}
		expected := []int{1, 2, 3, 4, 8, 7, 6, 5}
		result := SpiralOrder(matrix)
		assert.Equal(t, expected, result)
	})
}

func BenchmarkSpiralOrder(b *testing.B) {
	// Create a 100x100 matrix for benchmarking
	size := 100
	matrix := make([][]int, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]int, size)
		for j := 0; j < size; j++ {
			matrix[i][j] = i*size + j
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SpiralOrder(matrix)
	}
}