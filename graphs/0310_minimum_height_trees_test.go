package graphs

import (
	"reflect"
	"testing"
)

func TestSolve0310(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		edges    [][]int
		expected []int
	}{
		{
			name:     "Example 1",
			n:        4,
			edges:    [][]int{{1, 0}, {1, 2}, {1, 3}},
			expected: []int{1},
		},
		{
			name:     "Example 2",
			n:        6,
			edges:    [][]int{{3, 0}, {3, 1}, {3, 2}, {3, 4}, {5, 4}},
			expected: []int{3, 4},
		},
		{
			name:     "Single node",
			n:        1,
			edges:    [][]int{},
			expected: []int{0},
		},
		{
			name:     "Two nodes",
			n:        2,
			edges:    [][]int{{0, 1}},
			expected: []int{0, 1},
		},
		{
			name:     "Line of 5 nodes",
			n:        5,
			edges:    [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
			expected: []int{2},
		},
		{
			name:     "Star shape",
			n:        7,
			edges:    [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}, {0, 6}},
			expected: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Solve0310(tt.n, tt.edges)

			// Sort both slices for comparison
			sortInts(result)
			sortInts(tt.expected)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Solve0310(%d, %v) = %v, expected %v", tt.n, tt.edges, result, tt.expected)
			}
		})
	}
}

func sortInts(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}