package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "Example 1: obstacle in middle",
			grid: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			expected: 2,
		},
		{
			name: "Example 2: obstacle in first row",
			grid: [][]int{
				{0, 1},
				{0, 0},
			},
			expected: 1,
		},
		{
			name: "No obstacles 3x3",
			grid: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			expected: 6, // UniquePaths(3,3) = 6
		},
		{
			name: "Start blocked",
			grid: [][]int{
				{1, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			expected: 0,
		},
		{
			name: "End blocked",
			grid: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 1},
			},
			expected: 0,
		},
		{
			name: "Single cell no obstacle",
			grid: [][]int{
				{0},
			},
			expected: 1,
		},
		{
			name: "Single cell with obstacle",
			grid: [][]int{
				{1},
			},
			expected: 0,
		},
		{
			name: "First row blocked",
			grid: [][]int{
				{0, 1, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			expected: 3,
		},
		{
			name: "First column blocked",
			grid: [][]int{
				{0, 0, 0},
				{1, 0, 0},
				{0, 0, 0},
			},
			expected: 3,
		},
		{
			name: "Multiple obstacles",
			grid: [][]int{
				{0, 0, 0, 0},
				{0, 1, 0, 0},
				{0, 0, 1, 0},
				{0, 0, 0, 0},
			},
			expected: 4,
		},
		{
			name: "All obstacles",
			grid: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			expected: 0,
		},
		{
			name: "Large open grid 4x4",
			grid: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
			expected: 20, // UniquePaths(4,4) = 20
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := UniquePathsWithObstacles(tt.grid)
			assert.Equal(t, tt.expected, result,
				"UniquePathsWithObstacles(%v) = %d, expected %d",
				tt.grid, result, tt.expected)
		})
	}
}

func TestUniquePathsWithObstacles_EdgeCases(t *testing.T) {
	t.Run("Empty grid returns 0", func(t *testing.T) {
		result := UniquePathsWithObstacles([][]int{})
		assert.Equal(t, 0, result)
	})

	t.Run("Empty row returns 0", func(t *testing.T) {
		result := UniquePathsWithObstacles([][]int{{}})
		assert.Equal(t, 0, result)
	})

	t.Run("1x1 grid with 0", func(t *testing.T) {
		result := UniquePathsWithObstacles([][]int{{0}})
		assert.Equal(t, 1, result)
	})

	t.Run("1x1 grid with 1", func(t *testing.T) {
		result := UniquePathsWithObstacles([][]int{{1}})
		assert.Equal(t, 0, result)
	})

	t.Run("1xN grid with obstacle in middle", func(t *testing.T) {
		grid := [][]int{{0, 0, 1, 0, 0}}
		result := UniquePathsWithObstacles(grid)
		assert.Equal(t, 0, result, "Obstacle blocks the only path")
	})

	t.Run("Nx1 grid with obstacle in middle", func(t *testing.T) {
		grid := [][]int{{0}, {0}, {1}, {0}, {0}}
		result := UniquePathsWithObstacles(grid)
		assert.Equal(t, 0, result, "Obstacle blocks the only path")
	})

	t.Run("Grid with only one possible path", func(t *testing.T) {
		// This test case was incorrect - the original grid had no path
		// because robot can only move down/right, not left
		// Creating a correct test case with exactly one path
		grid := [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		}
		result := UniquePathsWithObstacles(grid)
		// Actually this has 2 paths, not 1
		// Let's create a truly single-path grid
		grid = [][]int{
			{0, 1},
			{0, 0},
		}
		result = UniquePathsWithObstacles(grid)
		assert.Equal(t, 1, result, "Should have exactly one path")
	})
}

func TestUniquePathsWithObstacles_Consistency(t *testing.T) {
	// Test that without obstacles, it matches UniquePaths
	testCases := []struct {
		m int
		n int
	}{
		{3, 7},
		{5, 5},
		{2, 8},
		{8, 2},
		{4, 6},
		{6, 4},
	}

	for _, tc := range testCases {
		// Create grid without obstacles
		grid := make([][]int, tc.m)
		for i := range grid {
			grid[i] = make([]int, tc.n)
		}

		resultWithObstacles := UniquePathsWithObstacles(grid)
		resultWithoutObstacles := UniquePaths(tc.m, tc.n)

		assert.Equal(t, resultWithoutObstacles, resultWithObstacles,
			"Without obstacles, UniquePathsWithObstacles should equal UniquePaths for %dx%d",
			tc.m, tc.n)
	}
}

func BenchmarkUniquePathsWithObstacles(b *testing.B) {
	// Create a 20x20 grid with some obstacles
	grid := make([][]int, 20)
	for i := range grid {
		grid[i] = make([]int, 20)
		for j := range grid[i] {
			// Place obstacles randomly (about 10% of cells)
			if (i+j)%10 == 0 {
				grid[i][j] = 1
			}
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		UniquePathsWithObstacles(grid)
	}
}