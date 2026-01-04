package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinPathSum(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "Example 1",
			grid: [][]int{
				{1, 3, 1},
				{1, 5, 1},
				{4, 2, 1},
			},
			expected: 7,
		},
		{
			name: "Example 2",
			grid: [][]int{
				{1, 2, 3},
				{4, 5, 6},
			},
			expected: 12,
		},
		{
			name: "Single cell",
			grid: [][]int{
				{5},
			},
			expected: 5,
		},
		{
			name: "Single row",
			grid: [][]int{
				{1, 2, 3, 4, 5},
			},
			expected: 15, // Sum of all elements (only path is right)
		},
		{
			name: "Single column",
			grid: [][]int{
				{1},
				{2},
				{3},
				{4},
				{5},
			},
			expected: 15, // Sum of all elements (only path is down)
		},
		{
			name: "2x2 grid",
			grid: [][]int{
				{1, 2},
				{1, 1},
			},
			expected: 3, // Path: 1 → 1 → 1
		},
		{
			name: "3x3 all ones",
			grid: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			expected: 5, // Any path has sum 5 (2 right + 2 down)
		},
		{
			name: "Grid with zeros",
			grid: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			expected: 0,
		},
		{
			name: "Large numbers",
			grid: [][]int{
				{100, 200, 300},
				{400, 500, 600},
				{700, 800, 900},
			},
			expected: 2100, // Path: 100 → 200 → 300 → 600 → 900
		},
		{
			name: "Mixed positive",
			grid: [][]int{
				{1, 2, 5},
				{3, 2, 1},
				{4, 3, 2},
			},
			expected: 8, // Path: 1 → 2 → 2 → 1 → 2
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinPathSum(tt.grid)
			assert.Equal(t, tt.expected, result,
				"MinPathSum(%v) = %d, expected %d",
				tt.grid, result, tt.expected)
		})
	}
}

func TestMinPathSum_EdgeCases(t *testing.T) {
	t.Run("Empty grid returns 0", func(t *testing.T) {
		result := MinPathSum([][]int{})
		assert.Equal(t, 0, result)
	})

	t.Run("Empty row returns 0", func(t *testing.T) {
		result := MinPathSum([][]int{{}})
		assert.Equal(t, 0, result)
	})

	t.Run("1x1 grid with negative (if allowed)", func(t *testing.T) {
		// Note: Problem says non-negative numbers, but testing edge case
		result := MinPathSum([][]int{{-5}})
		assert.Equal(t, -5, result)
	})

	t.Run("All same value", func(t *testing.T) {
		grid := [][]int{
			{7, 7, 7, 7},
			{7, 7, 7, 7},
			{7, 7, 7, 7},
			{7, 7, 7, 7},
		}
		result := MinPathSum(grid)
		// Any path: 3 right + 3 down = 6 moves, plus start = 7 cells
		assert.Equal(t, 7*7, result)
	})

	t.Run("Increasing diagonally", func(t *testing.T) {
		grid := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}
		result := MinPathSum(grid)
		// Best path: 1 → 2 → 3 → 6 → 9 = 21
		// Alternative: 1 → 4 → 7 → 8 → 9 = 29
		assert.Equal(t, 21, result)
	})

	t.Run("Decreasing diagonally", func(t *testing.T) {
		grid := [][]int{
			{9, 8, 7},
			{6, 5, 4},
			{3, 2, 1},
		}
		result := MinPathSum(grid)
		// Best path: 9 → 8 → 7 → 4 → 1 = 29
		// Alternative: 9 → 6 → 3 → 2 → 1 = 21
		assert.Equal(t, 21, result)
	})
}

func TestMinPathSum_PropertyBased(t *testing.T) {
	// Test that result is at least sum of first row or first column
	grid := [][]int{
		{1, 4, 2, 8},
		{3, 1, 5, 2},
		{2, 6, 3, 1},
		{4, 2, 1, 3},
	}

	result := MinPathSum(grid)

	// Calculate sum of first row (only moving right)
	firstRowSum := 0
	for j := 0; j < len(grid[0]); j++ {
		firstRowSum += grid[0][j]
	}

	// Calculate sum of first column (only moving down)
	firstColSum := 0
	for i := 0; i < len(grid); i++ {
		firstColSum += grid[i][0]
	}

	// Result should be <= min of these two
	minPossible := firstRowSum
	if firstColSum < minPossible {
		minPossible = firstColSum
	}

	assert.True(t, result <= minPossible,
		"Result %d should be <= min(firstRowSum=%d, firstColSum=%d)=%d",
		result, firstRowSum, firstColSum, minPossible)

	// Result should be >= grid[0][0] + grid[m-1][n-1]
	minValue := grid[0][0] + grid[len(grid)-1][len(grid[0])-1]
	assert.True(t, result >= minValue,
		"Result %d should be >= start+end = %d",
		result, minValue)
}

func BenchmarkMinPathSum(b *testing.B) {
	// Create different sized grids for benchmarking
	testCases := []struct {
		name string
		grid [][]int
	}{
		{
			name: "10x10",
			grid: func() [][]int {
				grid := make([][]int, 10)
				for i := range grid {
					grid[i] = make([]int, 10)
					for j := range grid[i] {
						grid[i][j] = (i + j) % 10
					}
				}
				return grid
			}(),
		},
		{
			name: "20x20",
			grid: func() [][]int {
				grid := make([][]int, 20)
				for i := range grid {
					grid[i] = make([]int, 20)
					for j := range grid[i] {
						grid[i][j] = (i*2 + j*3) % 20
					}
				}
				return grid
			}(),
		},
		{
			name: "50x50",
			grid: func() [][]int {
				grid := make([][]int, 50)
				for i := range grid {
					grid[i] = make([]int, 50)
					for j := range grid[i] {
						grid[i][j] = (i + j*2) % 50
					}
				}
				return grid
			}(),
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MinPathSum(tc.grid)
			}
		})
	}
}