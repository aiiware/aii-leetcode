package arrays

import (
	"sort"
	"testing"
)

func TestCombinationSum3(t *testing.T) {
	tests := []struct {
		name     string
		k        int
		n        int
		expected [][]int
	}{
		{
			name: "k=3, n=7",
			k:    3,
			n:    7,
			expected: [][]int{
				{1, 2, 4},
			},
		},
		{
			name: "k=3, n=9",
			k:    3,
			n:    9,
			expected: [][]int{
				{1, 2, 6},
				{1, 3, 5},
				{2, 3, 4},
			},
		},
		{
			name:     "k=2, n=18",
			k:        2,
			n:        18,
			expected: [][]int{},
		},
		{
			name: "k=1, n=5",
			k:    1,
			n:    5,
			expected: [][]int{
				{5},
			},
		},
		{
			name: "k=9, n=45",
			k:    9,
			n:    45,
			expected: [][]int{
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CombinationSum3(tt.k, tt.n)

			// Sort results for comparison
			sort.Slice(result, func(i, j int) bool {
				for k := range result[i] {
					if result[i][k] != result[j][k] {
						return result[i][k] < result[j][k]
					}
				}
				return false
			})

			if len(result) != len(tt.expected) {
				t.Errorf("CombinationSum3(%d, %d) = %v, want %v", tt.k, tt.n, result, tt.expected)
				return
			}

			for i := range result {
				if len(result[i]) != len(tt.expected[i]) {
					t.Errorf("CombinationSum3(%d, %d)[%d] = %v, want %v", tt.k, tt.n, i, result[i], tt.expected[i])
					continue
				}
				for j := range result[i] {
					if result[i][j] != tt.expected[i][j] {
						t.Errorf("CombinationSum3(%d, %d)[%d][%d] = %v, want %v", tt.k, tt.n, i, j, result[i][j], tt.expected[i][j])
					}
				}
			}
		})
	}
}

func BenchmarkCombinationSum3(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CombinationSum3(3, 9)
	}
}
