package arrays

import (
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		target   int
		expected bool
	}{
		{
			name: "Example 1 - target found",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   3,
			expected: true,
		},
		{
			name: "Example 2 - target not found",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   13,
			expected: false,
		},
		{
			name: "Single element matrix - target found",
			matrix: [][]int{
				{5},
			},
			target:   5,
			expected: true,
		},
		{
			name: "Single element matrix - target not found",
			matrix: [][]int{
				{5},
			},
			target:   3,
			expected: false,
		},
		{
			name: "Target at first position",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   1,
			expected: true,
		},
		{
			name: "Target at last position",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   60,
			expected: true,
		},
		{
			name: "Target in middle of matrix",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   34,
			expected: true,
		},
		{
			name: "Target less than first element",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   0,
			expected: false,
		},
		{
			name: "Target greater than last element",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   100,
			expected: false,
		},
		{
			name:     "Empty matrix",
			matrix:   [][]int{},
			target:   5,
			expected: false,
		},
		{
			name: "Matrix with empty row",
			matrix: [][]int{
				{},
			},
			target:   5,
			expected: false,
		},
		{
			name: "Single row matrix - target found",
			matrix: [][]int{
				{1, 3, 5, 7, 9},
			},
			target:   5,
			expected: true,
		},
		{
			name: "Single row matrix - target not found",
			matrix: [][]int{
				{1, 3, 5, 7, 9},
			},
			target:   4,
			expected: false,
		},
		{
			name: "Single column matrix - target found",
			matrix: [][]int{
				{1},
				{3},
				{5},
				{7},
				{9},
			},
			target:   5,
			expected: true,
		},
		{
			name: "Single column matrix - target not found",
			matrix: [][]int{
				{1},
				{3},
				{5},
				{7},
				{9},
			},
			target:   4,
			expected: false,
		},
		{
			name: "Large matrix - target found",
			matrix: [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
				{21, 22, 23, 24, 25},
			},
			target:   13,
			expected: true,
		},
		{
			name: "Large matrix - target not found",
			matrix: [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
				{21, 22, 23, 24, 25},
			},
			target:   26,
			expected: false,
		},
		{
			name: "Matrix with negative numbers - target found",
			matrix: [][]int{
				{-10, -8, -6, -4},
				{-2, 0, 2, 4},
				{6, 8, 10, 12},
			},
			target:   -6,
			expected: true,
		},
		{
			name: "Matrix with negative numbers - target not found",
			matrix: [][]int{
				{-10, -8, -6, -4},
				{-2, 0, 2, 4},
				{6, 8, 10, 12},
			},
			target:   -5,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SearchMatrix(tt.matrix, tt.target)
			if result != tt.expected {
				t.Errorf("SearchMatrix() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestSearchMatrixTwoStep(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		target   int
		expected bool
	}{
		{
			name: "Example 1 - target found",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   3,
			expected: true,
		},
		{
			name: "Example 2 - target not found",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   13,
			expected: false,
		},
		{
			name: "Single element matrix - target found",
			matrix: [][]int{
				{5},
			},
			target:   5,
			expected: true,
		},
		{
			name: "Target at first position",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   1,
			expected: true,
		},
		{
			name: "Target at last position",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   60,
			expected: true,
		},
		{
			name: "Target in middle of matrix",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   34,
			expected: true,
		},
		{
			name: "Target less than first element",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   0,
			expected: false,
		},
		{
			name: "Target greater than last element",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   100,
			expected: false,
		},
		{
			name:     "Empty matrix",
			matrix:   [][]int{},
			target:   5,
			expected: false,
		},
		{
			name: "Matrix with empty row",
			matrix: [][]int{
				{},
			},
			target:   5,
			expected: false,
		},
		{
			name: "Single row matrix - target found",
			matrix: [][]int{
				{1, 3, 5, 7, 9},
			},
			target:   5,
			expected: true,
		},
		{
			name: "Single column matrix - target found",
			matrix: [][]int{
				{1},
				{3},
				{5},
				{7},
				{9},
			},
			target:   5,
			expected: true,
		},
		{
			name: "Matrix with negative numbers - target found",
			matrix: [][]int{
				{-10, -8, -6, -4},
				{-2, 0, 2, 4},
				{6, 8, 10, 12},
			},
			target:   -6,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SearchMatrixTwoStep(tt.matrix, tt.target)
			if result != tt.expected {
				t.Errorf("SearchMatrixTwoStep() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestSearchMatrixLinear(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		target   int
		expected bool
	}{
		{
			name: "Example 1 - target found",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   3,
			expected: true,
		},
		{
			name: "Example 2 - target not found",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   13,
			expected: false,
		},
		{
			name: "Single element matrix - target found",
			matrix: [][]int{
				{5},
			},
			target:   5,
			expected: true,
		},
		{
			name: "Target at first position",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   1,
			expected: true,
		},
		{
			name: "Target at last position",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   60,
			expected: true,
		},
		{
			name: "Target in middle of matrix",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   34,
			expected: true,
		},
		{
			name: "Target less than first element",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   0,
			expected: false,
		},
		{
			name: "Target greater than last element",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target:   100,
			expected: false,
		},
		{
			name:     "Empty matrix",
			matrix:   [][]int{},
			target:   5,
			expected: false,
		},
		{
			name: "Matrix with empty row",
			matrix: [][]int{
				{},
			},
			target:   5,
			expected: false,
		},
		{
			name: "Single row matrix - target found",
			matrix: [][]int{
				{1, 3, 5, 7, 9},
			},
			target:   5,
			expected: true,
		},
		{
			name: "Single column matrix - target found",
			matrix: [][]int{
				{1},
				{3},
				{5},
				{7},
				{9},
			},
			target:   5,
			expected: true,
		},
		{
			name: "Matrix with negative numbers - target found",
			matrix: [][]int{
				{-10, -8, -6, -4},
				{-2, 0, 2, 4},
				{6, 8, 10, 12},
			},
			target:   -6,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SearchMatrixLinear(tt.matrix, tt.target)
			if result != tt.expected {
				t.Errorf("SearchMatrixLinear() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func BenchmarkSearchMatrix(b *testing.B) {
	// Create a large matrix for benchmarking
	matrix := make([][]int, 100)
	for i := range matrix {
		matrix[i] = make([]int, 100)
		for j := range matrix[i] {
			matrix[i][j] = i*100 + j + 1
		}
	}

	targets := []int{1, 5000, 10000, 50000, 9999, 10001}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		target := targets[i%len(targets)]
		SearchMatrix(matrix, target)
	}
}

func BenchmarkSearchMatrixTwoStep(b *testing.B) {
	// Create a large matrix for benchmarking
	matrix := make([][]int, 100)
	for i := range matrix {
		matrix[i] = make([]int, 100)
		for j := range matrix[i] {
			matrix[i][j] = i*100 + j + 1
		}
	}

	targets := []int{1, 5000, 10000, 50000, 9999, 10001}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		target := targets[i%len(targets)]
		SearchMatrixTwoStep(matrix, target)
	}
}

func BenchmarkSearchMatrixLinear(b *testing.B) {
	// Create a large matrix for benchmarking
	matrix := make([][]int, 100)
	for i := range matrix {
		matrix[i] = make([]int, 100)
		for j := range matrix[i] {
			matrix[i][j] = i*100 + j + 1
		}
	}

	targets := []int{1, 5000, 10000, 50000, 9999, 10001}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		target := targets[i%len(targets)]
		SearchMatrixLinear(matrix, target)
	}
}
