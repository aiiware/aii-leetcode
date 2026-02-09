package arrays

import (
	"reflect"
	"testing"
)

func TestCombine(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		k        int
		expected [][]int
	}{
		{
			name:     "Example 1 from LeetCode",
			n:        4,
			k:        2,
			expected: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			name:     "Example 2 from LeetCode",
			n:        1,
			k:        1,
			expected: [][]int{{1}},
		},
		{
			name:     "n=3, k=1",
			n:        3,
			k:        1,
			expected: [][]int{{1}, {2}, {3}},
		},
		{
			name:     "n=3, k=3",
			n:        3,
			k:        3,
			expected: [][]int{{1, 2, 3}},
		},
		{
			name: "n=5, k=2",
			n:    5,
			k:    2,
			expected: [][]int{
				{1, 2}, {1, 3}, {1, 4}, {1, 5},
				{2, 3}, {2, 4}, {2, 5},
				{3, 4}, {3, 5},
				{4, 5},
			},
		},
		{
			name:     "n=4, k=4",
			n:        4,
			k:        4,
			expected: [][]int{{1, 2, 3, 4}},
		},
		{
			name:     "n=4, k=0 (edge case)",
			n:        4,
			k:        0,
			expected: [][]int{},
		},
		{
			name:     "n=0, k=0 (edge case)",
			n:        0,
			k:        0,
			expected: [][]int{},
		},
		{
			name:     "k > n (edge case)",
			n:        2,
			k:        3,
			expected: [][]int{},
		},
		{
			name: "n=5, k=3",
			n:    5,
			k:    3,
			expected: [][]int{
				{1, 2, 3}, {1, 2, 4}, {1, 2, 5},
				{1, 3, 4}, {1, 3, 5}, {1, 4, 5},
				{2, 3, 4}, {2, 3, 5}, {2, 4, 5},
				{3, 4, 5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Combine(tt.n, tt.k)
			if !equalCombinations(result, tt.expected) {
				t.Errorf("Combine(%d, %d) = %v, expected %v", tt.n, tt.k, result, tt.expected)
			}
		})
	}
}

func TestCombineOptimized(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		k        int
		expected [][]int
	}{
		{
			name:     "Example 1 from LeetCode",
			n:        4,
			k:        2,
			expected: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			name:     "Example 2 from LeetCode",
			n:        1,
			k:        1,
			expected: [][]int{{1}},
		},
		{
			name:     "n=3, k=1",
			n:        3,
			k:        1,
			expected: [][]int{{1}, {2}, {3}},
		},
		{
			name:     "n=3, k=3",
			n:        3,
			k:        3,
			expected: [][]int{{1, 2, 3}},
		},
		{
			name: "n=5, k=2",
			n:    5,
			k:    2,
			expected: [][]int{
				{1, 2}, {1, 3}, {1, 4}, {1, 5},
				{2, 3}, {2, 4}, {2, 5},
				{3, 4}, {3, 5},
				{4, 5},
			},
		},
		{
			name:     "n=4, k=4",
			n:        4,
			k:        4,
			expected: [][]int{{1, 2, 3, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CombineOptimized(tt.n, tt.k)
			if !equalCombinations(result, tt.expected) {
				t.Errorf("CombineOptimized(%d, %d) = %v, expected %v", tt.n, tt.k, result, tt.expected)
			}
		})
	}
}

func TestCombineIterative(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		k        int
		expected [][]int
	}{
		{
			name:     "Example 1 from LeetCode",
			n:        4,
			k:        2,
			expected: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			name:     "Example 2 from LeetCode",
			n:        1,
			k:        1,
			expected: [][]int{{1}},
		},
		{
			name:     "n=3, k=1",
			n:        3,
			k:        1,
			expected: [][]int{{1}, {2}, {3}},
		},
		{
			name:     "n=3, k=3",
			n:        3,
			k:        3,
			expected: [][]int{{1, 2, 3}},
		},
		{
			name: "n=5, k=2",
			n:    5,
			k:    2,
			expected: [][]int{
				{1, 2}, {1, 3}, {1, 4}, {1, 5},
				{2, 3}, {2, 4}, {2, 5},
				{3, 4}, {3, 5},
				{4, 5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CombineIterative(tt.n, tt.k)
			if !equalCombinations(result, tt.expected) {
				t.Errorf("CombineIterative(%d, %d) = %v, expected %v", tt.n, tt.k, result, tt.expected)
			}
		})
	}
}

func TestCombineDP(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		k        int
		expected [][]int
	}{
		{
			name:     "Example 1 from LeetCode",
			n:        4,
			k:        2,
			expected: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			name:     "Example 2 from LeetCode",
			n:        1,
			k:        1,
			expected: [][]int{{1}},
		},
		{
			name:     "n=3, k=1",
			n:        3,
			k:        1,
			expected: [][]int{{1}, {2}, {3}},
		},
		{
			name:     "n=3, k=3",
			n:        3,
			k:        3,
			expected: [][]int{{1, 2, 3}},
		},
		{
			name: "n=5, k=2",
			n:    5,
			k:    2,
			expected: [][]int{
				{1, 2}, {1, 3}, {1, 4}, {1, 5},
				{2, 3}, {2, 4}, {2, 5},
				{3, 4}, {3, 5},
				{4, 5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CombineDP(tt.n, tt.k)
			if !equalCombinations(result, tt.expected) {
				t.Errorf("CombineDP(%d, %d) = %v, expected %v", tt.n, tt.k, result, tt.expected)
			}
		})
	}
}

func TestCombineMath(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		k        int
		expected [][]int
	}{
		{
			name:     "Example 1 from LeetCode",
			n:        4,
			k:        2,
			expected: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			name:     "Example 2 from LeetCode",
			n:        1,
			k:        1,
			expected: [][]int{{1}},
		},
		{
			name:     "n=3, k=1",
			n:        3,
			k:        1,
			expected: [][]int{{1}, {2}, {3}},
		},
		{
			name:     "n=3, k=3",
			n:        3,
			k:        3,
			expected: [][]int{{1, 2, 3}},
		},
		{
			name: "n=5, k=2",
			n:    5,
			k:    2,
			expected: [][]int{
				{1, 2}, {1, 3}, {1, 4}, {1, 5},
				{2, 3}, {2, 4}, {2, 5},
				{3, 4}, {3, 5},
				{4, 5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CombineMath(tt.n, tt.k)
			if !equalCombinations(result, tt.expected) {
				t.Errorf("CombineMath(%d, %d) = %v, expected %v", tt.n, tt.k, result, tt.expected)
			}
		})
	}
}

func BenchmarkCombine(b *testing.B) {
	testCases := []struct {
		name string
		n    int
		k    int
	}{
		{
			name: "Small (n=5, k=2)",
			n:    5,
			k:    2,
		},
		{
			name: "Medium (n=10, k=3)",
			n:    10,
			k:    3,
		},
		{
			name: "Large (n=15, k=5)",
			n:    15,
			k:    5,
		},
		{
			name: "Maximum (n=20, k=10)",
			n:    20,
			k:    10,
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Combine(tc.n, tc.k)
			}
		})
	}
}

func BenchmarkCombineOptimized(b *testing.B) {
	testCases := []struct {
		name string
		n    int
		k    int
	}{
		{
			name: "Small (n=5, k=2)",
			n:    5,
			k:    2,
		},
		{
			name: "Medium (n=10, k=3)",
			n:    10,
			k:    3,
		},
		{
			name: "Large (n=15, k=5)",
			n:    15,
			k:    5,
		},
		{
			name: "Maximum (n=20, k=10)",
			n:    20,
			k:    10,
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CombineOptimized(tc.n, tc.k)
			}
		})
	}
}

func BenchmarkCombineIterative(b *testing.B) {
	testCases := []struct {
		name string
		n    int
		k    int
	}{
		{
			name: "Small (n=5, k=2)",
			n:    5,
			k:    2,
		},
		{
			name: "Medium (n=10, k=3)",
			n:    10,
			k:    3,
		},
		{
			name: "Large (n=15, k=5)",
			n:    15,
			k:    5,
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CombineIterative(tc.n, tc.k)
			}
		})
	}
}

func BenchmarkCombineDP(b *testing.B) {
	testCases := []struct {
		name string
		n    int
		k    int
	}{
		{
			name: "Small (n=5, k=2)",
			n:    5,
			k:    2,
		},
		{
			name: "Medium (n=10, k=3)",
			n:    10,
			k:    3,
		},
		{
			name: "Large (n=15, k=5)",
			n:    15,
			k:    5,
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CombineDP(tc.n, tc.k)
			}
		})
	}
}

func BenchmarkCombineMath(b *testing.B) {
	testCases := []struct {
		name string
		n    int
		k    int
	}{
		{
			name: "Small (n=5, k=2)",
			n:    5,
			k:    2,
		},
		{
			name: "Medium (n=10, k=3)",
			n:    10,
			k:    3,
		},
		{
			name: "Large (n=15, k=5)",
			n:    15,
			k:    5,
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CombineMath(tc.n, tc.k)
			}
		})
	}
}

// Helper function to compare combinations (order of combinations doesn't matter)
func equalCombinations(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	// Create maps to count occurrences of each combination
	mapA := make(map[string]int)
	mapB := make(map[string]int)

	for _, comb := range a {
		key := ""
		for _, num := range comb {
			key += string(rune(num)) + ","
		}
		mapA[key]++
	}

	for _, comb := range b {
		key := ""
		for _, num := range comb {
			key += string(rune(num)) + ","
		}
		mapB[key]++
	}

	return reflect.DeepEqual(mapA, mapB)
}
