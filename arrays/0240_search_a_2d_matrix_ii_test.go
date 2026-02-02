package arrays

import (
	"testing"
)

func TestSearchMatrixII(t *testing.T) {
	// Test matrix from LeetCode example
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	tests := []struct {
		target   int
		expected bool
	}{
		// Test cases from LeetCode
		{target: 5, expected: true},
		{target: 20, expected: false},
		// Edge cases
		{target: 1, expected: true},   // Top-left corner
		{target: 30, expected: true},  // Bottom-right corner
		{target: 18, expected: true},  // Bottom-left corner
		{target: 15, expected: true},  // Top-right corner
		{target: 0, expected: false},  // Below minimum
		{target: 31, expected: false}, // Above maximum
		// Elements in the matrix
		{target: 9, expected: true},
		{target: 14, expected: true},
		{target: 23, expected: true},
		// Elements not in matrix
		{target: 25, expected: false},
		{target: 31, expected: false},
	}

	for i, test := range tests {
		// Test main implementation (top-right search)
		result := SearchMatrixII(matrix, test.target)
		if result != test.expected {
			t.Errorf("Test case %d failed (top-right): target=%d, expected=%v, got=%v",
				i+1, test.target, test.expected, result)
		}

		// Test binary search implementation
		binaryResult := SearchMatrixIIBinarySearch(matrix, test.target)
		if binaryResult != test.expected {
			t.Errorf("Test case %d failed (binary search): target=%d, expected=%v, got=%v",
				i+1, test.target, test.expected, binaryResult)
		}

		// Test divide and conquer implementation
		dcResult := SearchMatrixIIDivideConquer(matrix, test.target)
		if dcResult != test.expected {
			t.Errorf("Test case %d failed (divide & conquer): target=%d, expected=%v, got=%v",
				i+1, test.target, test.expected, dcResult)
		}

		// Ensure all implementations agree
		if result != binaryResult || result != dcResult {
			t.Errorf("Test case %d: Implementations disagree: top-right=%v, binary=%v, dc=%v",
				i+1, result, binaryResult, dcResult)
		}
	}
}

func TestSearchMatrixIIEdgeCases(t *testing.T) {
	// Test empty matrix
	t.Run("Empty matrix", func(t *testing.T) {
		emptyMatrix := [][]int{}
		if SearchMatrixII(emptyMatrix, 5) != false {
			t.Error("Expected false for empty matrix")
		}
	})

	// Test matrix with empty rows
	t.Run("Matrix with empty rows", func(t *testing.T) {
		matrix := [][]int{{}, {}, {}}
		if SearchMatrixII(matrix, 5) != false {
			t.Error("Expected false for matrix with empty rows")
		}
	})

	// Test 1x1 matrix
	t.Run("1x1 matrix", func(t *testing.T) {
		matrix := [][]int{{5}}
		if SearchMatrixII(matrix, 5) != true {
			t.Error("Expected true for target in 1x1 matrix")
		}
		if SearchMatrixII(matrix, 3) != false {
			t.Error("Expected false for target not in 1x1 matrix")
		}
	})

	// Test 1xN matrix (single row)
	t.Run("1xN matrix", func(t *testing.T) {
		matrix := [][]int{{1, 3, 5, 7, 9}}
		if SearchMatrixII(matrix, 5) != true {
			t.Error("Expected true for target in single row")
		}
		if SearchMatrixII(matrix, 4) != false {
			t.Error("Expected false for target not in single row")
		}
	})

	// Test Nx1 matrix (single column)
	t.Run("Nx1 matrix", func(t *testing.T) {
		matrix := [][]int{{1}, {3}, {5}, {7}, {9}}
		if SearchMatrixII(matrix, 5) != true {
			t.Error("Expected true for target in single column")
		}
		if SearchMatrixII(matrix, 4) != false {
			t.Error("Expected false for target not in single column")
		}
	})

	// Test with negative numbers
	t.Run("Matrix with negative numbers", func(t *testing.T) {
		matrix := [][]int{
			{-5, -3, -1},
			{-4, -2, 0},
			{-3, -1, 1},
		}
		if SearchMatrixII(matrix, -2) != true {
			t.Error("Expected true for negative target")
		}
		if SearchMatrixII(matrix, 2) != false {
			t.Error("Expected false for target not in matrix")
		}
	})
}

func BenchmarkSearchMatrixII(b *testing.B) {
	// Create a large matrix for benchmarking
	matrix := make([][]int, 100)
	for i := 0; i < 100; i++ {
		matrix[i] = make([]int, 100)
		for j := 0; j < 100; j++ {
			matrix[i][j] = i + j*10
		}
	}
	target := 550 // In the matrix

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SearchMatrixII(matrix, target)
	}
}

func BenchmarkSearchMatrixIIBinarySearch(b *testing.B) {
	matrix := make([][]int, 100)
	for i := 0; i < 100; i++ {
		matrix[i] = make([]int, 100)
		for j := 0; j < 100; j++ {
			matrix[i][j] = i + j*10
		}
	}
	target := 550

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SearchMatrixIIBinarySearch(matrix, target)
	}
}

func BenchmarkSearchMatrixIIDivideConquer(b *testing.B) {
	// Smaller matrix for divide & conquer (it's less efficient)
	matrix := make([][]int, 50)
	for i := 0; i < 50; i++ {
		matrix[i] = make([]int, 50)
		for j := 0; j < 50; j++ {
			matrix[i][j] = i + j*10
		}
	}
	target := 250

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SearchMatrixIIDivideConquer(matrix, target)
	}
}