package graphs

import (
	"testing"
)

func TestSolve0994(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			name: "Example 1",
			grid: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			expected: 4,
		},
		{
			name: "Example 2",
			grid: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			expected: -1,
		},
		{
			name: "Example 3",
			grid: [][]int{
				{0, 2},
			},
			expected: 0,
		},
		{
			name: "All rotten initially",
			grid: [][]int{
				{2, 2, 2},
				{2, 2, 2},
			},
			expected: 0,
		},
		{
			name: "Single fresh orange",
			grid: [][]int{
				{2, 1},
			},
			expected: 1,
		},
		{
			name: "Fresh orange isolated",
			grid: [][]int{
				{2, 0, 1},
				{0, 0, 0},
				{1, 0, 2},
			},
			expected: -1,
		},
		{
			name: "Chain reaction",
			grid: [][]int{
				{2, 1, 0},
				{1, 1, 1},
				{0, 1, 1},
			},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of the grid since Solve0994 modifies it
			gridCopy := make([][]int, len(tt.grid))
			for i := range tt.grid {
				gridCopy[i] = make([]int, len(tt.grid[i]))
				copy(gridCopy[i], tt.grid[i])
			}

			result := Solve0994(gridCopy)

			if result != tt.expected {
				t.Errorf("Solve0994(%v) = %d, expected %d", tt.grid, result, tt.expected)
			}
		})
	}
}