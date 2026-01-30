package dp

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name     string
		numRows  int
		expected [][]int
	}{
		{
			name:    "0 rows",
			numRows: 0,
			expected: [][]int{},
		},
		{
			name:    "1 row",
			numRows: 1,
			expected: [][]int{
				{1},
			},
		},
		{
			name:    "2 rows",
			numRows: 2,
			expected: [][]int{
				{1},
				{1, 1},
			},
		},
		{
			name:    "3 rows",
			numRows: 3,
			expected: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
			},
		},
		{
			name:    "4 rows",
			numRows: 4,
			expected: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
			},
		},
		{
			name:    "5 rows",
			numRows: 5,
			expected: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
		{
			name:    "6 rows",
			numRows: 6,
			expected: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
				{1, 5, 10, 10, 5, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generate(tt.numRows)
			if !trianglesEqual(result, tt.expected) {
				t.Errorf("generate(%d) = %v, expected %v", tt.numRows, result, tt.expected)
			}
		})
	}
}

func TestGenerateOptimized(t *testing.T) {
	tests := []struct {
		name     string
		numRows  int
		expected [][]int
	}{
		{
			name:    "0 rows",
			numRows: 0,
			expected: [][]int{},
		},
		{
			name:    "1 row",
			numRows: 1,
			expected: [][]int{
				{1},
			},
		},
		{
			name:    "5 rows",
			numRows: 5,
			expected: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateOptimized(tt.numRows)
			if !trianglesEqual(result, tt.expected) {
				t.Errorf("generateOptimized(%d) = %v, expected %v", tt.numRows, result, tt.expected)
			}
		})
	}
}

func TestTrianglesEqual(t *testing.T) {
	tests := []struct {
		name     string
		t1       [][]int
		t2       [][]int
		expected bool
	}{
		{
			name:     "both empty",
			t1:       [][]int{},
			t2:       [][]int{},
			expected: true,
		},
		{
			name: "identical triangles",
			t1: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
			},
			t2: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
			},
			expected: true,
		},
		{
			name: "different triangles",
			t1: [][]int{
				{1},
				{1, 1},
			},
			t2: [][]int{
				{1},
				{1, 2},
			},
			expected: false,
		},
		{
			name: "different lengths",
			t1: [][]int{
				{1},
				{1, 1},
			},
			t2: [][]int{
				{1},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := trianglesEqual(tt.t1, tt.t2)
			if result != tt.expected {
				t.Errorf("trianglesEqual() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestPrintTriangle(t *testing.T) {
	tests := []struct {
		name     string
		triangle [][]int
		expected string
	}{
		{
			name:     "empty triangle",
			triangle: [][]int{},
			expected: "[]",
		},
		{
			name: "single row",
			triangle: [][]int{
				{1},
			},
			expected: "    1", // 4 spaces + "1"
		},
		{
			name: "three rows",
			triangle: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
			},
			expected: `        1
      1     1
    1     2     1`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := printTriangle(tt.triangle)
			if result != tt.expected {
				t.Errorf("printTriangle() = %q, expected %q", result, tt.expected)
			}
		})
	}
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generate(100)
	}
}

func BenchmarkGenerateOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateOptimized(100)
	}
}