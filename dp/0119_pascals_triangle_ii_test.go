package dp

import (
	"testing"
)

func TestGetRow(t *testing.T) {
	tests := []struct {
		name     string
		rowIndex int
		expected []int
	}{
		{
			name:     "row 0",
			rowIndex: 0,
			expected: []int{1},
		},
		{
			name:     "row 1",
			rowIndex: 1,
			expected: []int{1, 1},
		},
		{
			name:     "row 2",
			rowIndex: 2,
			expected: []int{1, 2, 1},
		},
		{
			name:     "row 3",
			rowIndex: 3,
			expected: []int{1, 3, 3, 1},
		},
		{
			name:     "row 4",
			rowIndex: 4,
			expected: []int{1, 4, 6, 4, 1},
		},
		{
			name:     "row 5",
			rowIndex: 5,
			expected: []int{1, 5, 10, 10, 5, 1},
		},
		{
			name:     "row 6",
			rowIndex: 6,
			expected: []int{1, 6, 15, 20, 15, 6, 1},
		},
		{
			name:     "row 10",
			rowIndex: 10,
			expected: []int{1, 10, 45, 120, 210, 252, 210, 120, 45, 10, 1},
		},
		{
			name:     "negative index",
			rowIndex: -1,
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getRow(tt.rowIndex)
			if !rowsEqual(result, tt.expected) {
				t.Errorf("getRow(%d) = %v, expected %v", tt.rowIndex, result, tt.expected)
			}
		})
	}
}

func TestGetRowTwoArrays(t *testing.T) {
	tests := []struct {
		name     string
		rowIndex int
		expected []int
	}{
		{
			name:     "row 0",
			rowIndex: 0,
			expected: []int{1},
		},
		{
			name:     "row 3",
			rowIndex: 3,
			expected: []int{1, 3, 3, 1},
		},
		{
			name:     "row 5",
			rowIndex: 5,
			expected: []int{1, 5, 10, 10, 5, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getRowTwoArrays(tt.rowIndex)
			if !rowsEqual(result, tt.expected) {
				t.Errorf("getRowTwoArrays(%d) = %v, expected %v", tt.rowIndex, result, tt.expected)
			}
		})
	}
}

func TestGetRowMath(t *testing.T) {
	tests := []struct {
		name     string
		rowIndex int
		expected []int
	}{
		{
			name:     "row 0",
			rowIndex: 0,
			expected: []int{1},
		},
		{
			name:     "row 3",
			rowIndex: 3,
			expected: []int{1, 3, 3, 1},
		},
		{
			name:     "row 5",
			rowIndex: 5,
			expected: []int{1, 5, 10, 10, 5, 1},
		},
		{
			name:     "row 10",
			rowIndex: 10,
			expected: []int{1, 10, 45, 120, 210, 252, 210, 120, 45, 10, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getRowMath(tt.rowIndex)
			if !rowsEqual(result, tt.expected) {
				t.Errorf("getRowMath(%d) = %v, expected %v", tt.rowIndex, result, tt.expected)
			}
		})
	}
}

func TestRowsEqual(t *testing.T) {
	tests := []struct {
		name     string
		r1       []int
		r2       []int
		expected bool
	}{
		{
			name:     "both empty",
			r1:       []int{},
			r2:       []int{},
			expected: true,
		},
		{
			name:     "identical rows",
			r1:       []int{1, 3, 3, 1},
			r2:       []int{1, 3, 3, 1},
			expected: true,
		},
		{
			name:     "different rows",
			r1:       []int{1, 2, 1},
			r2:       []int{1, 3, 1},
			expected: false,
		},
		{
			name:     "different lengths",
			r1:       []int{1, 1},
			r2:       []int{1},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rowsEqual(tt.r1, tt.r2)
			if result != tt.expected {
				t.Errorf("rowsEqual() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestGetRowsUpTo(t *testing.T) {
	tests := []struct {
		name     string
		rowIndex int
		expected [][]int
	}{
		{
			name:     "up to row 0",
			rowIndex: 0,
			expected: [][]int{
				{1},
			},
		},
		{
			name:     "up to row 2",
			rowIndex: 2,
			expected: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
			},
		},
		{
			name:     "up to row 3",
			rowIndex: 3,
			expected: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getRowsUpTo(tt.rowIndex)
			if len(result) != len(tt.expected) {
				t.Errorf("getRowsUpTo(%d) returned %d rows, expected %d", tt.rowIndex, len(result), len(tt.expected))
				return
			}
			for i := range result {
				if !rowsEqual(result[i], tt.expected[i]) {
					t.Errorf("row %d: got %v, expected %v", i, result[i], tt.expected[i])
				}
			}
		})
	}
}

func BenchmarkGetRow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getRow(100)
	}
}

func BenchmarkGetRowTwoArrays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getRowTwoArrays(100)
	}
}

func BenchmarkGetRowMath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getRowMath(100)
	}
}