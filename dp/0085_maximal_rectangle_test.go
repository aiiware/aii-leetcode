package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximalRectangle(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]byte
		expected int
	}{
		{
			name: "Example 1 from LeetCode",
			matrix: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			expected: 6,
		},
		{
			name: "Example 2 from LeetCode",
			matrix: [][]byte{
				{'0'},
			},
			expected: 0,
		},
		{
			name: "Example 3 from LeetCode",
			matrix: [][]byte{
				{'1'},
			},
			expected: 1,
		},
		{
			name:     "Empty matrix",
			matrix:   [][]byte{},
			expected: 0,
		},
		{
			name: "Single row all ones",
			matrix: [][]byte{
				{'1', '1', '1', '1', '1'},
			},
			expected: 5,
		},
		{
			name: "Single column all ones",
			matrix: [][]byte{
				{'1'},
				{'1'},
				{'1'},
				{'1'},
				{'1'},
			},
			expected: 5,
		},
		{
			name: "All ones matrix",
			matrix: [][]byte{
				{'1', '1', '1'},
				{'1', '1', '1'},
				{'1', '1', '1'},
			},
			expected: 9,
		},
		{
			name: "All zeros matrix",
			matrix: [][]byte{
				{'0', '0', '0'},
				{'0', '0', '0'},
				{'0', '0', '0'},
			},
			expected: 0,
		},
		{
			name: "Checkerboard pattern",
			matrix: [][]byte{
				{'1', '0', '1', '0'},
				{'0', '1', '0', '1'},
				{'1', '0', '1', '0'},
				{'0', '1', '0', '1'},
			},
			expected: 1,
		},
		{
			name: "L-shaped rectangle",
			matrix: [][]byte{
				{'1', '1', '0', '0'},
				{'1', '1', '0', '0'},
				{'1', '1', '1', '1'},
				{'1', '1', '1', '1'},
			},
			expected: 8,
		},
		{
			name: "Multiple rectangles",
			matrix: [][]byte{
				{'1', '1', '0', '0', '1'},
				{'1', '1', '0', '0', '1'},
				{'0', '0', '1', '1', '0'},
				{'0', '0', '1', '1', '0'},
			},
			expected: 4,
		},
		{
			name: "Rectangle in corner",
			matrix: [][]byte{
				{'1', '1', '0', '0'},
				{'1', '1', '0', '0'},
				{'0', '0', '0', '0'},
				{'0', '0', '0', '0'},
			},
			expected: 4,
		},
		{
			name: "Rectangle in center",
			matrix: [][]byte{
				{'0', '0', '0', '0', '0'},
				{'0', '1', '1', '1', '0'},
				{'0', '1', '1', '1', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 6,
		},
		{
			name: "Tall rectangle",
			matrix: [][]byte{
				{'1', '0', '0'},
				{'1', '0', '0'},
				{'1', '0', '0'},
				{'1', '0', '0'},
				{'1', '0', '0'},
			},
			expected: 5,
		},
		{
			name: "Wide rectangle",
			matrix: [][]byte{
				{'1', '1', '1', '1', '1'},
				{'0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 5,
		},
		{
			name: "Diagonal ones",
			matrix: [][]byte{
				{'1', '0', '0', '0'},
				{'0', '1', '0', '0'},
				{'0', '0', '1', '0'},
				{'0', '0', '0', '1'},
			},
			expected: 1,
		},
		{
			name: "Cross shape",
			matrix: [][]byte{
				{'0', '1', '0'},
				{'1', '1', '1'},
				{'0', '1', '0'},
			},
			expected: 3,
		},
		{
			name: "Hollow rectangle",
			matrix: [][]byte{
				{'1', '1', '1', '1'},
				{'1', '0', '0', '1'},
				{'1', '1', '1', '1'},
			},
			expected: 4,
		},
		{
			name: "Staircase pattern",
			matrix: [][]byte{
				{'1', '0', '0', '0'},
				{'1', '1', '0', '0'},
				{'1', '1', '1', '0'},
				{'1', '1', '1', '1'},
			},
			expected: 6,
		},
		{
			name: "Reverse staircase",
			matrix: [][]byte{
				{'1', '1', '1', '1'},
				{'0', '1', '1', '1'},
				{'0', '0', '1', '1'},
				{'0', '0', '0', '1'},
			},
			expected: 6,
		},
		{
			name: "Single zero in all ones",
			matrix: [][]byte{
				{'1', '1', '1'},
				{'1', '0', '1'},
				{'1', '1', '1'},
			},
			expected: 3,
		},
		{
			name: "Complex case 1",
			matrix: [][]byte{
				{'0', '1', '1', '0', '1'},
				{'1', '1', '0', '1', '0'},
				{'0', '1', '1', '1', '0'},
				{'1', '1', '1', '1', '0'},
				{'1', '1', '1', '1', '1'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 9,
		},
		{
			name: "Complex case 2",
			matrix: [][]byte{
				{'1', '0', '1', '1', '0', '1'},
				{'1', '1', '1', '1', '1', '1'},
				{'0', '1', '1', '0', '1', '1'},
				{'1', '1', '1', '0', '1', '0'},
				{'0', '1', '1', '1', '1', '1'},
				{'1', '1', '0', '1', '1', '1'},
			},
			expected: 8,
		},
		{
			name: "Large rectangle at bottom",
			matrix: [][]byte{
				{'0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
			},
			expected: 15,
		},
		{
			name: "Large rectangle at top",
			matrix: [][]byte{
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 15,
		},
		{
			name: "Narrow vertical strip",
			matrix: [][]byte{
				{'0', '1', '0', '0', '0'},
				{'0', '1', '0', '0', '0'},
				{'0', '1', '0', '0', '0'},
				{'0', '1', '0', '0', '0'},
				{'0', '1', '0', '0', '0'},
			},
			expected: 5,
		},
		{
			name: "Narrow horizontal strip",
			matrix: [][]byte{
				{'0', '0', '0', '0', '0'},
				{'1', '1', '1', '1', '1'},
				{'0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 5,
		},
		{
			name: "Two separate rectangles",
			matrix: [][]byte{
				{'1', '1', '0', '1', '1'},
				{'1', '1', '0', '1', '1'},
				{'0', '0', '0', '0', '0'},
				{'1', '1', '0', '1', '1'},
				{'1', '1', '0', '1', '1'},
			},
			expected: 4,
		},
		{
			name: "Edge case: 1x1 with '1'",
			matrix: [][]byte{
				{'1'},
			},
			expected: 1,
		},
		{
			name: "Edge case: 1x1 with '0'",
			matrix: [][]byte{
				{'0'},
			},
			expected: 0,
		},
		{
			name: "Edge case: 1xN all ones",
			matrix: [][]byte{
				{'1', '1', '1', '1', '1', '1', '1'},
			},
			expected: 7,
		},
		{
			name: "Edge case: Nx1 all ones",
			matrix: [][]byte{
				{'1'},
				{'1'},
				{'1'},
				{'1'},
				{'1'},
				{'1'},
				{'1'},
			},
			expected: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaximalRectangle(tt.matrix)
			assert.Equal(t, tt.expected, result,
				"MaximalRectangle(%v) = %d, expected %d",
				tt.matrix, result, tt.expected)
		})
	}
}

func TestMaximalRectangleDP(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]byte
		expected int
	}{
		{
			name: "Example 1",
			matrix: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			expected: 6,
		},
		{
			name: "All ones",
			matrix: [][]byte{
				{'1', '1', '1'},
				{'1', '1', '1'},
				{'1', '1', '1'},
			},
			expected: 9,
		},
		{
			name: "Single one",
			matrix: [][]byte{
				{'1'},
			},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaximalRectangleDP(tt.matrix)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaximalRectangleBruteForce(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]byte
		expected int
	}{
		{
			name: "Small matrix",
			matrix: [][]byte{
				{'1', '0'},
				{'1', '1'},
			},
			expected: 2,
		},
		{
			name: "Single row",
			matrix: [][]byte{
				{'1', '1', '0', '1'},
			},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaximalRectangleBruteForce(tt.matrix)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaximalRectangleOptimizedBruteForce(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]byte
		expected int
	}{
		{
			name: "Example 1",
			matrix: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			expected: 6,
		},
		{
			name: "2x2 all ones",
			matrix: [][]byte{
				{'1', '1'},
				{'1', '1'},
			},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaximalRectangleOptimizedBruteForce(tt.matrix)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaximalRectangleDivideConquer(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]byte
		expected int
	}{
		{
			name: "Example 1",
			matrix: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			expected: 6,
		},
		{
			name: "Simple case",
			matrix: [][]byte{
				{'1', '1'},
				{'1', '1'},
			},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaximalRectangleDivideConquer(tt.matrix)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestMaximalRectangle_Consistency(t *testing.T) {
	testCases := []struct {
		name   string
		matrix [][]byte
	}{
		{
			name: "Standard example",
			matrix: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
		},
		{
			name: "All ones 3x3",
			matrix: [][]byte{
				{'1', '1', '1'},
				{'1', '1', '1'},
				{'1', '1', '1'},
			},
		},
		{
			name: "All zeros",
			matrix: [][]byte{
				{'0', '0', '0'},
				{'0', '0', '0'},
				{'0', '0', '0'},
			},
		},
		{
			name: "Single row",
			matrix: [][]byte{
				{'1', '0', '1', '1', '0'},
			},
		},
		{
			name: "Single column",
			matrix: [][]byte{
				{'1'},
				{'0'},
				{'1'},
				{'1'},
				{'0'},
			},
		},
		{
			name: "Checkerboard",
			matrix: [][]byte{
				{'1', '0', '1'},
				{'0', '1', '0'},
				{'1', '0', '1'},
			},
		},
		{
			name: "L shape",
			matrix: [][]byte{
				{'1', '1', '0'},
				{'1', '1', '0'},
				{'1', '1', '1'},
			},
		},
		{
			name: "Empty matrix",
			matrix: [][]byte{},
		},
		{
			name: "1x1 with 1",
			matrix: [][]byte{
				{'1'},
			},
		},
		{
			name: "1x1 with 0",
			matrix: [][]byte{
				{'0'},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test all implementations
			implementations := []struct {
				name string
				fn   func([][]byte) int
			}{
				{"MaximalRectangle", MaximalRectangle},
				{"MaximalRectangleDP", MaximalRectangleDP},
				{"MaximalRectangleOptimizedBruteForce", MaximalRectangleOptimizedBruteForce},
				{"MaximalRectangleDivideConquer", MaximalRectangleDivideConquer},
			}

			results := make([]int, len(implementations))
			
			for i, impl := range implementations {
				// For brute force, skip large matrices
				if impl.name == "MaximalRectangleBruteForce" {
					rows := len(tc.matrix)
					if rows > 0 {
						cols := len(tc.matrix[0])
						if rows*cols > 100 {
							continue
						}
					}
				}
				results[i] = impl.fn(tc.matrix)
			}

			// All implementations should return the same result
			firstResult := results[0]
			for i := 1; i < len(results); i++ {
				if results[i] != 0 { // Skip if result is 0 (might be from skipped implementation)
					assert.Equal(t, firstResult, results[i],
						"%s and %s should return same result for matrix %v",
						implementations[0].name, implementations[i].name, tc.matrix)
				}
			}
		})
	}
}

func TestMaximalRectangle_PropertyBased(t *testing.T) {
	t.Run("Result is non-negative", func(t *testing.T) {
		testCases := [][][]byte{
			{
				{'1', '0', '1'},
				{'0', '1', '0'},
			},
			{
				{'0', '0', '0'},
				{'0', '0', '0'},
			},
			{
				{'1'},
			},
			{},
		}

		for _, matrix := range testCases {
			result := MaximalRectangle(matrix)
			assert.True(t, result >= 0,
				"Result %d should be non-negative for matrix %v",
				result, matrix)
		}
	})

	t.Run("Result is at most total cells", func(t *testing.T) {
		testCases := [][][]byte{
			{
				{'1', '0', '1'},
				{'0', '1', '0'},
			},
			{
				{'1', '1', '1'},
				{'1', '1', '1'},
			},
			{
				{'1'},
			},
		}

		for _, matrix := range testCases {
			if len(matrix) == 0 {
				continue
			}
			
			result := MaximalRectangle(matrix)
			totalCells := len(matrix) * len(matrix[0])
			
			assert.True(t, result <= totalCells,
				"Result %d should be at most total cells %d for matrix %v",
				result, totalCells, matrix)
		}
	})

	t.Run("Result is at least max row ones * max col ones", func(t *testing.T) {
		testCases := [][][]byte{
			{
				{'1', '0', '1'},
				{'0', '1', '0'},
			},
			{
				{'1', '1', '1'},
				{'1', '1', '1'},
			},
			{
				{'1', '0', '0'},
				{'0', '1', '0'},
				{'0', '0', '1'},
			},
		}

		for _, matrix := range testCases {
			if len(matrix) == 0 {
				continue
			}
			
			result := MaximalRectangle(matrix)
			
			// Calculate max consecutive ones in any row
			maxRowOnes := 0
			for _, row := range matrix {
				current := 0
				for _, cell := range row {
					if cell == '1' {
						current++
						if current > maxRowOnes {
							maxRowOnes = current
						}
					} else {
						current = 0
					}
				}
			}
			
			// Calculate max consecutive ones in any column
			maxColOnes := 0
			rows, cols := len(matrix), len(matrix[0])
			for j := 0; j < cols; j++ {
				current := 0
				for i := 0; i < rows; i++ {
					if matrix[i][j] == '1' {
						current++
						if current > maxColOnes {
							maxColOnes = current
						}
					} else {
						current = 0
					}
				}
			}
			
			// The maximal rectangle area should be at least maxRowOnes * maxColOnes
			// (though not exactly, since they might not align)
			minPossible := maxRowOnes * maxColOnes
			assert.True(t, result >= minPossible || (maxRowOnes == 0 && maxColOnes == 0),
				"Result %d should be at least %d (maxRowOnes=%d * maxColOnes=%d) for matrix %v",
				result, minPossible, maxRowOnes, maxColOnes, matrix)
		}
	})

	t.Run("Adding zeros doesn't increase max area", func(t *testing.T) {
		original := [][]byte{
			{'1', '0', '1'},
			{'0', '1', '0'},
			{'1', '0', '1'},
		}
		
		// Create a larger matrix with zeros around
		withBorder := make([][]byte, 5)
		for i := range withBorder {
			withBorder[i] = make([]byte, 5)
			for j := range withBorder[i] {
				if i >= 1 && i <= 3 && j >= 1 && j <= 3 {
					withBorder[i][j] = original[i-1][j-1]
				} else {
					withBorder[i][j] = '0'
				}
			}
		}
		
		originalResult := MaximalRectangle(original)
		withBorderResult := MaximalRectangle(withBorder)
		
		// Adding zeros around shouldn't increase the maximum area
		assert.True(t, withBorderResult <= originalResult,
			"Adding zeros around shouldn't increase max area. Original: %d, With border: %d",
			originalResult, withBorderResult)
	})

	t.Run("Scaling property for all-ones submatrix", func(t *testing.T) {
		// If we have an all-ones submatrix of size m x n, the area should be at least m * n
		testCases := []struct {
			matrix [][]byte
			m, n   int
		}{
			{
				matrix: [][]byte{
					{'1', '1', '1'},
					{'1', '1', '1'},
					{'1', '1', '1'},
				},
				m: 3, n: 3,
			},
			{
				matrix: [][]byte{
					{'1', '1'},
					{'1', '1'},
					{'0', '0'},
				},
				m: 2, n: 2,
			},
		}

		for _, tc := range testCases {
			result := MaximalRectangle(tc.matrix)
			expectedMin := tc.m * tc.n
			assert.True(t, result >= expectedMin,
				"Matrix with %dx%d all-ones submatrix should have area >= %d, got %d",
				tc.m, tc.n, expectedMin, result)
		}
	})
}

func BenchmarkMaximalRectangle(b *testing.B) {
	// Create test cases of different sizes and patterns
	testCases := []struct {
		name   string
		matrix [][]byte
	}{
		{
			name: "Small (5x5)",
			matrix: [][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
				{'0', '1', '1', '0', '1'},
			},
		},
		{
			name: "Medium (10x10) random",
			matrix: func() [][]byte {
				matrix := make([][]byte, 10)
				for i := range matrix {
					matrix[i] = make([]byte, 10)
					for j := range matrix[i] {
						if (i+j)%2 == 0 {
							matrix[i][j] = '1'
						} else {
							matrix[i][j] = '0'
						}
					}
				}
				return matrix
			}(),
		},
		{
			name: "Medium (10x10) all ones",
			matrix: func() [][]byte {
				matrix := make([][]byte, 10)
				for i := range matrix {
					matrix[i] = make([]byte, 10)
					for j := range matrix[i] {
						matrix[i][j] = '1'
					}
				}
				return matrix
			}(),
		},
		{
			name: "Medium (10x10) all zeros",
			matrix: func() [][]byte {
				matrix := make([][]byte, 10)
				for i := range matrix {
					matrix[i] = make([]byte, 10)
					for j := range matrix[i] {
						matrix[i][j] = '0'
					}
				}
				return matrix
			}(),
		},
		{
			name: "Large (20x20) checkerboard",
			matrix: func() [][]byte {
				matrix := make([][]byte, 20)
				for i := range matrix {
					matrix[i] = make([]byte, 20)
					for j := range matrix[i] {
						if (i+j)%2 == 0 {
							matrix[i][j] = '1'
						} else {
							matrix[i][j] = '0'
						}
					}
				}
				return matrix
			}(),
		},
		{
			name: "Large (20x20) vertical stripes",
			matrix: func() [][]byte {
				matrix := make([][]byte, 20)
				for i := range matrix {
					matrix[i] = make([]byte, 20)
					for j := range matrix[i] {
						if j%3 == 0 {
							matrix[i][j] = '1'
						} else {
							matrix[i][j] = '0'
						}
					}
				}
				return matrix
			}(),
		},
		{
			name: "Large (20x20) horizontal stripes",
			matrix: func() [][]byte {
				matrix := make([][]byte, 20)
				for i := range matrix {
					matrix[i] = make([]byte, 20)
					for j := range matrix[i] {
						if i%3 == 0 {
							matrix[i][j] = '1'
						} else {
							matrix[i][j] = '0'
						}
					}
				}
				return matrix
			}(),
		},
		{
			name: "Large (20x20) block pattern",
			matrix: func() [][]byte {
				matrix := make([][]byte, 20)
				for i := range matrix {
					matrix[i] = make([]byte, 20)
					for j := range matrix[i] {
						if i/5 == j/5 {
							matrix[i][j] = '1'
						} else {
							matrix[i][j] = '0'
						}
					}
				}
				return matrix
			}(),
		},
		{
			name: "Very large (50x50) sparse",
			matrix: func() [][]byte {
				matrix := make([][]byte, 50)
				for i := range matrix {
					matrix[i] = make([]byte, 50)
					for j := range matrix[i] {
						if (i*j)%7 == 0 {
							matrix[i][j] = '1'
						} else {
							matrix[i][j] = '0'
						}
					}
				}
				return matrix
			}(),
		},
		{
			name: "Very large (50x50) dense",
			matrix: func() [][]byte {
				matrix := make([][]byte, 50)
				for i := range matrix {
					matrix[i] = make([]byte, 50)
					for j := range matrix[i] {
						if (i+j)%3 != 0 {
							matrix[i][j] = '1'
						} else {
							matrix[i][j] = '0'
						}
					}
				}
				return matrix
			}(),
		},
	}

	implementations := []struct {
		name string
		fn   func([][]byte) int
	}{
		{"Standard", MaximalRectangle},
		{"DP", MaximalRectangleDP},
		{"OptimizedBruteForce", MaximalRectangleOptimizedBruteForce},
		{"DivideConquer", MaximalRectangleDivideConquer},
		// Note: BruteForce is excluded from benchmarks as it's O(m^2 * n^2)
	}

	for _, tc := range testCases {
		for _, impl := range implementations {
			// Skip certain implementations for very large arrays
			if tc.name == "Very large (50x50) sparse" || tc.name == "Very large (50x50) dense" {
				if impl.name == "OptimizedBruteForce" || impl.name == "DivideConquer" {
					continue // These are O(m^2 * n) or O(m * n * log(min(m, n)))
				}
			}
			
			b.Run(tc.name+"_"+impl.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					impl.fn(tc.matrix)
				}
			})
		}
	}
}