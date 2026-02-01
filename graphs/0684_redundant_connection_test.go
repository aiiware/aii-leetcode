package graphs

import (
	"reflect"
	"testing"
)

func TestSolve0684(t *testing.T) {
	tests := []struct {
		name     string
		edges    [][]int
		expected []int
	}{
		{
			name:     "Example 1",
			edges:    [][]int{{1, 2}, {1, 3}, {2, 3}},
			expected: []int{2, 3},
		},
		{
			name:     "Example 2",
			edges:    [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}},
			expected: []int{1, 4},
		},
		{
			name:     "Simple triangle",
			edges:    [][]int{{1, 2}, {2, 3}, {3, 1}},
			expected: []int{3, 1},
		},
		{
			name:     "Multiple cycles, return last",
			edges:    [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 1}, {1, 5}},
			expected: []int{4, 1},
		},
		{
			name:     "Star with redundant edge",
			edges:    [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}},
			expected: []int{2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Solve0684(tt.edges)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Solve0684(%v) = %v, expected %v", tt.edges, result, tt.expected)
			}
		})
	}
}