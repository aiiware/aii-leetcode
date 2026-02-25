package dp

import (
	"testing"
)

func TestMinFallingPathSum(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected int
	}{
		{
			name: "Example 1",
			matrix: [][]int{
				{2, 1, 3},
				{6, 5, 4},
				{7, 8, 9},
			},
			expected: 13,
		},
		{
			name: "Example 2",
			matrix: [][]int{
				{-19, 57},
				{-40, -5},
			},
			expected: -59,
		},
		{
			name: "Example 3 - Single element",
			matrix: [][]int{
				{-48},
			},
			expected: -48,
		},
		{
			name: "2x2 matrix all positive",
			matrix: [][]int{
				{1, 2},
				{3, 4},
			},
			expected: 4, // Path: 1 → 3 = 4
		},
		{
			name: "3x3 matrix with negative values",
			matrix: [][]int{
				{1, -2, 3},
				{4, -5, 6},
				{7, -8, 9},
			},
			expected: -15, // Calculated: path 1→-2→-5→-8 = -15? Let's trace: row0: [1,-2,3], row1: [5,-7,4], row2: [0,-15,2], min = -15
		},
		{
			name: "All zeros",
			matrix: [][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			expected: 0,
		},
		{
			name: "All same positive values",
			matrix: [][]int{
				{5, 5, 5},
				{5, 5, 5},
				{5, 5, 5},
			},
			expected: 15, // Any path: 5 + 5 + 5 = 15
		},
		{
			name: "Path through edges",
			matrix: [][]int{
				{10, 1, 10},
				{10, 10, 10},
				{10, 10, 10},
			},
			expected: 21, // Path: 1 → 10 → 10 = 21
		},
		{
			name: "4x4 matrix",
			matrix: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			expected: 28, // Calculated: path 1→5→9→13 = 28
		},
		{
			name: "Empty matrix",
			matrix: [][]int{},
			expected: 0,
		},
		{
			name: "Single row matrix",
			matrix: [][]int{
				{1, 2, 3, 4, 5},
			},
			expected: 1, // Minimum of first row
		},
		{
			name: "Matrix with minimum path not starting at minimum of first row",
			matrix: [][]int{
				{5, 1, 5},
				{5, 1, 5},
				{5, 1, 5},
			},
			expected: 3, // Path: 1 → 1 → 1 = 3
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test the main function
			result := MinFallingPathSum(tt.matrix)
			if result != tt.expected {
				t.Errorf("MinFallingPathSum() = %d, expected %d", result, tt.expected)
			}

			// Test DP implementation
			if len(tt.matrix) > 0 {
				resultDP := minFallingPathSumDP(tt.matrix)
				if resultDP != tt.expected {
					t.Errorf("minFallingPathSumDP() = %d, expected %d", resultDP, tt.expected)
				}

				// Test optimized implementation
				resultOpt := minFallingPathSumOptimized(tt.matrix)
				if resultOpt != tt.expected {
					t.Errorf("minFallingPathSumOptimized() = %d, expected %d", resultOpt, tt.expected)
				}

				// Test in-place implementation (need to copy matrix first)
				if len(tt.matrix) > 0 {
					matrixCopy := make([][]int, len(tt.matrix))
					for i := range tt.matrix {
						matrixCopy[i] = make([]int, len(tt.matrix[i]))
						copy(matrixCopy[i], tt.matrix[i])
					}
					resultInPlace := minFallingPathSumInPlace(matrixCopy)
					if resultInPlace != tt.expected {
						t.Errorf("minFallingPathSumInPlace() = %d, expected %d", resultInPlace, tt.expected)
					}
				}
			}
		})
	}
}

func TestMinFallingPathSumDP(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected int
	}{
		{
			name: "Basic 3x3",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: 12, // Path: 1 → 4 → 7 = 12
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minFallingPathSumDP(tt.matrix)
			if result != tt.expected {
				t.Errorf("minFallingPathSumDP() = %d, expected %d", result, tt.expected)
			}
		})
	}
}

func TestMinFallingPathSumOptimized(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected int
	}{
		{
			name: "Basic 3x3",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: 12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minFallingPathSumOptimized(tt.matrix)
			if result != tt.expected {
				t.Errorf("minFallingPathSumOptimized() = %d, expected %d", result, tt.expected)
			}
		})
	}
}

func TestMinFallingPathSumInPlace(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected int
	}{
		{
			name: "Basic 3x3",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: 12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Copy matrix since in-place modifies it
			matrixCopy := make([][]int, len(tt.matrix))
			for i := range tt.matrix {
				matrixCopy[i] = make([]int, len(tt.matrix[i]))
				copy(matrixCopy[i], tt.matrix[i])
			}
			
			result := minFallingPathSumInPlace(matrixCopy)
			if result != tt.expected {
				t.Errorf("minFallingPathSumInPlace() = %d, expected %d", result, tt.expected)
			}
		})
	}
}

func BenchmarkMinFallingPathSum(b *testing.B) {
	matrix := [][]int{
		{2, 1, 3, 4, 5},
		{6, 5, 4, 3, 2},
		{7, 8, 9, 1, 2},
		{3, 4, 5, 6, 7},
		{8, 9, 1, 2, 3},
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinFallingPathSum(matrix)
	}
}

func BenchmarkMinFallingPathSumDP(b *testing.B) {
	matrix := [][]int{
		{2, 1, 3, 4, 5},
		{6, 5, 4, 3, 2},
		{7, 8, 9, 1, 2},
		{3, 4, 5, 6, 7},
		{8, 9, 1, 2, 3},
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		minFallingPathSumDP(matrix)
	}
}

func BenchmarkMinFallingPathSumOptimized(b *testing.B) {
	matrix := [][]int{
		{2, 1, 3, 4, 5},
		{6, 5, 4, 3, 2},
		{7, 8, 9, 1, 2},
		{3, 4, 5, 6, 7},
		{8, 9, 1, 2, 3},
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		minFallingPathSumOptimized(matrix)
	}
}

func BenchmarkMinFallingPathSumInPlace(b *testing.B) {
	// Create fresh matrix for each iteration since it modifies in-place
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		matrix := [][]int{
			{2, 1, 3, 4, 5},
			{6, 5, 4, 3, 2},
			{7, 8, 9, 1, 2},
			{3, 4, 5, 6, 7},
			{8, 9, 1, 2, 3},
		}
		minFallingPathSumInPlace(matrix)
	}
}