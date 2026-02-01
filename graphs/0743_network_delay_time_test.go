package graphs

import (
	"testing"
)

func TestSolve0743(t *testing.T) {
	tests := []struct {
		name     string
		times    [][]int
		n        int
		k        int
		expected int
	}{
		{
			name:     "Example 1",
			times:    [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}},
			n:        4,
			k:        2,
			expected: 2,
		},
		{
			name:     "Example 2",
			times:    [][]int{{1, 2, 1}},
			n:        2,
			k:        1,
			expected: 1,
		},
		{
			name:     "Example 3",
			times:    [][]int{{1, 2, 1}},
			n:        2,
			k:        2,
			expected: -1,
		},
		{
			name:     "Single node",
			times:    [][]int{},
			n:        1,
			k:        1,
			expected: 0,
		},
		{
			name:     "Star network",
			times:    [][]int{{1, 2, 2}, {1, 3, 4}, {1, 4, 1}},
			n:        4,
			k:        1,
			expected: 4,
		},
		{
			name:     "Unreachable node",
			times:    [][]int{{1, 2, 1}, {3, 4, 1}},
			n:        4,
			k:        1,
			expected: -1,
		},
		{
			name:     "Multiple paths",
			times:    [][]int{{1, 2, 1}, {1, 3, 4}, {2, 3, 2}, {3, 4, 1}},
			n:        4,
			k:        1,
			expected: 4, // Path: 1->2->3->4 = 1+2+1 = 4
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Solve0743(tt.times, tt.n, tt.k)

			if result != tt.expected {
				t.Errorf("Solve0743(%v, %d, %d) = %d, expected %d", tt.times, tt.n, tt.k, result, tt.expected)
			}
		})
	}
}