package dp

import (
	"testing"
)

func TestLongestIncreasingPath(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected int
	}{
		{
			name: "Example 1",
			matrix: [][]int{
				{9, 9, 4},
				{6, 6, 8},
				{2, 1, 1},
			},
			expected: 4,
		},
		{
			name: "Example 2",
			matrix: [][]int{
				{3, 4, 5},
				{3, 2, 6},
				{2, 2, 1},
			},
			expected: 4,
		},
		{
			name: "Example 3",
			matrix: [][]int{
				{1},
			},
			expected: 1,
		},
		{
			name: "Single row increasing",
			matrix: [][]int{
				{1, 2, 3, 4, 5},
			},
			expected: 5,
		},
		{
			name: "Single column increasing",
			matrix: [][]int{
				{1},
				{2},
				{3},
				{4},
				{5},
			},
			expected: 5,
		},
		{
			name: "All same values",
			matrix: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			expected: 1,
		},
		{
			name: "Decreasing matrix",
			matrix: [][]int{
				{5, 4, 3},
				{4, 3, 2},
				{3, 2, 1},
			},
			expected: 1,
		},
		{
			name: "Complex case 1",
			matrix: [][]int{
				{7, 8, 9},
				{9, 7, 6},
				{7, 2, 3},
			},
			expected: 6, // Path: 2->3->6->7->8->9 or 2->3->6->7->8->9
		},
		{
			name: "Complex case 2",
			matrix: [][]int{
				{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
				{19, 18, 17, 16, 15, 14, 13, 12, 11, 10},
				{20, 21, 22, 23, 24, 25, 26, 27, 28, 29},
				{39, 38, 37, 36, 35, 34, 33, 32, 31, 30},
				{40, 41, 42, 43, 44, 45, 46, 47, 48, 49},
				{59, 58, 57, 56, 55, 54, 53, 52, 51, 50},
				{60, 61, 62, 63, 64, 65, 66, 67, 68, 69},
				{79, 78, 77, 76, 75, 74, 73, 72, 71, 70},
				{80, 81, 82, 83, 84, 85, 86, 87, 88, 89},
				{99, 98, 97, 96, 95, 94, 93, 92, 91, 90},
				{100, 101, 102, 103, 104, 105, 106, 107, 108, 109},
				{119, 118, 117, 116, 115, 114, 113, 112, 111, 110},
				{120, 121, 122, 123, 124, 125, 126, 127, 128, 129},
				{139, 138, 137, 136, 135, 134, 133, 132, 131, 130},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			expected: 140, // Snake pattern through the matrix
		},
		{
			name: "Empty matrix",
			matrix: [][]int{},
			expected: 0,
		},
		{
			name: "Single element matrix",
			matrix: [][]int{
				{42},
			},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestIncreasingPath(tt.matrix)
			if result != tt.expected {
				t.Errorf("longestIncreasingPath() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func BenchmarkLongestIncreasingPath(b *testing.B) {
	matrix := [][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longestIncreasingPath(matrix)
	}
}

func BenchmarkLongestIncreasingPathLarge(b *testing.B) {
	// Create a 50x50 matrix with increasing values
	matrix := make([][]int, 50)
	for i := range matrix {
		matrix[i] = make([]int, 50)
		for j := range matrix[i] {
			matrix[i][j] = i*50 + j
		}
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		longestIncreasingPath(matrix)
	}
}