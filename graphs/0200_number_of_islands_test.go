package graphs

import (
	"testing"
)

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]byte
		expected int
	}{
		{
			name: "Example 1",
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 1,
		},
		{
			name: "Example 2",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			expected: 3,
		},
		{
			name:     "Empty grid",
			grid:     [][]byte{},
			expected: 0,
		},
		{
			name: "Single island",
			grid: [][]byte{
				{'1'},
			},
			expected: 1,
		},
		{
			name: "Single water",
			grid: [][]byte{
				{'0'},
			},
			expected: 0,
		},
		{
			name: "All land",
			grid: [][]byte{
				{'1', '1', '1'},
				{'1', '1', '1'},
				{'1', '1', '1'},
			},
			expected: 1,
		},
		{
			name: "All water",
			grid: [][]byte{
				{'0', '0', '0'},
				{'0', '0', '0'},
				{'0', '0', '0'},
			},
			expected: 0,
		},
		{
			name: "Diagonal islands",
			grid: [][]byte{
				{'1', '0', '0', '0'},
				{'0', '1', '0', '0'},
				{'0', '0', '1', '0'},
				{'0', '0', '0', '1'},
			},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of the grid for DFS solution
			gridDFS := make([][]byte, len(tt.grid))
			for i := range tt.grid {
				gridDFS[i] = make([]byte, len(tt.grid[i]))
				copy(gridDFS[i], tt.grid[i])
			}
			
			// Create a copy of the grid for BFS solution
			gridBFS := make([][]byte, len(tt.grid))
			for i := range tt.grid {
				gridBFS[i] = make([]byte, len(tt.grid[i]))
				copy(gridBFS[i], tt.grid[i])
			}
			
			// Test DFS solution
			resultDFS := numIslands(gridDFS)
			if resultDFS != tt.expected {
				t.Errorf("numIslands(DFS) = %d, expected %d", resultDFS, tt.expected)
			}
			
			// Test BFS solution
			resultBFS := numIslandsBFS(gridBFS)
			if resultBFS != tt.expected {
				t.Errorf("numIslands(BFS) = %d, expected %d", resultBFS, tt.expected)
			}
		})
	}
}

func BenchmarkNumIslandsDFS(b *testing.B) {
	grid := [][]byte{
		{'1', '1', '0', '0', '0', '1', '0', '1', '1', '0'},
		{'1', '1', '0', '0', '0', '0', '1', '0', '0', '1'},
		{'0', '0', '1', '0', '0', '1', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1', '0', '0', '0', '1', '1'},
		{'1', '1', '0', '0', '0', '1', '0', '1', '1', '0'},
		{'1', '1', '0', '0', '0', '0', '1', '0', '0', '1'},
		{'0', '0', '1', '0', '0', '1', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1', '0', '0', '0', '1', '1'},
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a fresh copy for each iteration
		gridCopy := make([][]byte, len(grid))
		for j := range grid {
			gridCopy[j] = make([]byte, len(grid[j]))
			copy(gridCopy[j], grid[j])
		}
		numIslands(gridCopy)
	}
}

func BenchmarkNumIslandsBFS(b *testing.B) {
	grid := [][]byte{
		{'1', '1', '0', '0', '0', '1', '0', '1', '1', '0'},
		{'1', '1', '0', '0', '0', '0', '1', '0', '0', '1'},
		{'0', '0', '1', '0', '0', '1', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1', '0', '0', '0', '1', '1'},
		{'1', '1', '0', '0', '0', '1', '0', '1', '1', '0'},
		{'1', '1', '0', '0', '0', '0', '1', '0', '0', '1'},
		{'0', '0', '1', '0', '0', '1', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1', '0', '0', '0', '1', '1'},
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a fresh copy for each iteration
		gridCopy := make([][]byte, len(grid))
		for j := range grid {
			gridCopy[j] = make([]byte, len(grid[j]))
			copy(gridCopy[j], grid[j])
		}
		numIslandsBFS(gridCopy)
	}
}