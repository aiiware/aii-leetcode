package arrays

import (
	"testing"
)

func TestSetZeroes(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{
			name: "Example 1",
			input: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			expected: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			name: "Example 2",
			input: [][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			},
			expected: [][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
		{
			name: "Single element zero",
			input: [][]int{
				{0},
			},
			expected: [][]int{
				{0},
			},
		},
		{
			name: "Single element non-zero",
			input: [][]int{
				{5},
			},
			expected: [][]int{
				{5},
			},
		},
		{
			name: "All zeros",
			input: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			expected: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
		},
		{
			name: "No zeros",
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
		},
		{
			name: "First row has zero",
			input: [][]int{
				{0, 1, 2},
				{3, 4, 5},
				{6, 7, 8},
			},
			expected: [][]int{
				{0, 0, 0},
				{0, 4, 5},
				{0, 7, 8},
			},
		},
		{
			name: "First column has zero",
			input: [][]int{
				{1, 2, 3},
				{0, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{0, 2, 3},
				{0, 0, 0},
				{0, 8, 9},
			},
		},
		{
			name: "Multiple zeros in same row",
			input: [][]int{
				{1, 0, 1, 0},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
			expected: [][]int{
				{0, 0, 0, 0},
				{1, 0, 1, 0},
				{1, 0, 1, 0},
			},
		},
		{
			name: "Multiple zeros in same column",
			input: [][]int{
				{1, 1, 1},
				{0, 1, 1},
				{0, 1, 1},
				{1, 1, 1},
			},
			expected: [][]int{
				{0, 1, 1},
				{0, 0, 0},
				{0, 0, 0},
				{0, 1, 1},
			},
		},
		{
			name: "Large matrix with zeros",
			input: [][]int{
				{1, 2, 3, 4, 5},
				{6, 0, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 0, 20},
				{21, 22, 23, 24, 25},
			},
			expected: [][]int{
				{1, 0, 3, 0, 5},
				{0, 0, 0, 0, 0},
				{11, 0, 13, 0, 15},
				{0, 0, 0, 0, 0},
				{21, 0, 23, 0, 25},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of the input to avoid modifying the test case
			input := make([][]int, len(tt.input))
			for i := range tt.input {
				input[i] = make([]int, len(tt.input[i]))
				copy(input[i], tt.input[i])
			}

			SetZeroes(input)

			if !matrixEqual(input, tt.expected) {
				t.Errorf("SetZeroes() = %v, expected %v", input, tt.expected)
			}
		})
	}
}

func TestSetZeroesWithExtraSpace(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{
			name: "Example 1",
			input: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			expected: [][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
		{
			name: "Example 2",
			input: [][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			},
			expected: [][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
		{
			name: "Single element zero",
			input: [][]int{
				{0},
			},
			expected: [][]int{
				{0},
			},
		},
		{
			name: "Large matrix with zeros",
			input: [][]int{
				{1, 2, 3, 4, 5},
				{6, 0, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 0, 20},
				{21, 22, 23, 24, 25},
			},
			expected: [][]int{
				{1, 0, 3, 0, 5},
				{0, 0, 0, 0, 0},
				{11, 0, 13, 0, 15},
				{0, 0, 0, 0, 0},
				{21, 0, 23, 0, 25},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of the input to avoid modifying the test case
			input := make([][]int, len(tt.input))
			for i := range tt.input {
				input[i] = make([]int, len(tt.input[i]))
				copy(input[i], tt.input[i])
			}

			SetZeroesWithExtraSpace(input)

			if !matrixEqual(input, tt.expected) {
				t.Errorf("SetZeroesWithExtraSpace() = %v, expected %v", input, tt.expected)
			}
		})
	}
}

func BenchmarkSetZeroes(b *testing.B) {
	// Create a large matrix for benchmarking
	matrix := make([][]int, 200)
	for i := range matrix {
		matrix[i] = make([]int, 200)
		for j := range matrix[i] {
			// Place zeros at random positions (about 5% zeros)
			if i%20 == 0 || j%20 == 0 {
				matrix[i][j] = 0
			} else {
				matrix[i][j] = i*200 + j + 1
			}
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a copy for each benchmark iteration
		testMatrix := make([][]int, len(matrix))
		for i := range matrix {
			testMatrix[i] = make([]int, len(matrix[i]))
			copy(testMatrix[i], matrix[i])
		}
		SetZeroes(testMatrix)
	}
}

func BenchmarkSetZeroesWithExtraSpace(b *testing.B) {
	// Create a large matrix for benchmarking
	matrix := make([][]int, 200)
	for i := range matrix {
		matrix[i] = make([]int, 200)
		for j := range matrix[i] {
			// Place zeros at random positions (about 5% zeros)
			if i%20 == 0 || j%20 == 0 {
				matrix[i][j] = 0
			} else {
				matrix[i][j] = i*200 + j + 1
			}
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a copy for each benchmark iteration
		testMatrix := make([][]int, len(matrix))
		for i := range matrix {
			testMatrix[i] = make([]int, len(matrix[i]))
			copy(testMatrix[i], matrix[i])
		}
		SetZeroesWithExtraSpace(testMatrix)
	}
}

// Helper function to compare matrices
func matrixEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
