package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name     string
		m        int
		n        int
		expected int
	}{
		{
			name:     "Example 1: m=3, n=7",
			m:        3,
			n:        7,
			expected: 28,
		},
		{
			name:     "Example 2: m=3, n=2",
			m:        3,
			n:        2,
			expected: 3,
		},
		{
			name:     "m=1, n=1",
			m:        1,
			n:        1,
			expected: 1,
		},
		{
			name:     "m=1, n=10",
			m:        1,
			n:        10,
			expected: 1,
		},
		{
			name:     "m=10, n=1",
			m:        10,
			n:        1,
			expected: 1,
		},
		{
			name:     "m=2, n=2",
			m:        2,
			n:        2,
			expected: 2,
		},
		{
			name:     "m=2, n=3",
			m:        2,
			n:        3,
			expected: 3,
		},
		{
			name:     "m=3, n=3",
			m:        3,
			n:        3,
			expected: 6,
		},
		{
			name:     "m=4, n=4",
			m:        4,
			n:        4,
			expected: 20,
		},
		{
			name:     "m=5, n=5",
			m:        5,
			n:        5,
			expected: 70,
		},
		{
			name:     "m=7, n=3",
			m:        7,
			n:        3,
			expected: 28, // Same as m=3, n=7 (symmetric)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := UniquePaths(tt.m, tt.n)
			assert.Equal(t, tt.expected, result,
				"UniquePaths(%d, %d) = %d, expected %d",
				tt.m, tt.n, result, tt.expected)
		})
	}
}

func TestUniquePaths_EdgeCases(t *testing.T) {
	t.Run("m=0 returns 0", func(t *testing.T) {
		result := UniquePaths(0, 5)
		assert.Equal(t, 0, result)
	})

	t.Run("n=0 returns 0", func(t *testing.T) {
		result := UniquePaths(5, 0)
		assert.Equal(t, 0, result)
	})

	t.Run("both zero returns 0", func(t *testing.T) {
		result := UniquePaths(0, 0)
		assert.Equal(t, 0, result)
	})

	t.Run("negative m returns 0", func(t *testing.T) {
		result := UniquePaths(-1, 5)
		assert.Equal(t, 0, result)
	})

	t.Run("negative n returns 0", func(t *testing.T) {
		result := UniquePaths(5, -1)
		assert.Equal(t, 0, result)
	})

	t.Run("large grid m=10, n=10", func(t *testing.T) {
		result := UniquePaths(10, 10)
		// 10x10 grid has 48620 unique paths (known value)
		assert.Equal(t, 48620, result)
	})

	t.Run("very narrow grid m=1, n=100", func(t *testing.T) {
		result := UniquePaths(1, 100)
		assert.Equal(t, 1, result)
	})

	t.Run("very tall grid m=100, n=1", func(t *testing.T) {
		result := UniquePaths(100, 1)
		assert.Equal(t, 1, result)
	})
}

func TestUniquePaths_Symmetry(t *testing.T) {
	// Test that UniquePaths(m, n) == UniquePaths(n, m)
	testCases := []struct {
		m int
		n int
	}{
		{3, 7},
		{5, 5},
		{2, 8},
		{8, 2},
		{4, 6},
		{6, 4},
	}

	for _, tc := range testCases {
		result1 := UniquePaths(tc.m, tc.n)
		result2 := UniquePaths(tc.n, tc.m)
		assert.Equal(t, result1, result2,
			"UniquePaths(%d, %d) = %d should equal UniquePaths(%d, %d) = %d",
			tc.m, tc.n, result1, tc.n, tc.m, result2)
	}
}

func BenchmarkUniquePaths(b *testing.B) {
	testCases := []struct {
		name string
		m    int
		n    int
	}{
		{"small", 3, 7},
		{"medium", 10, 10},
		{"large", 20, 20},
		{"wide", 5, 50},
		{"tall", 50, 5},
		{"very large", 50, 50},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UniquePaths(tc.m, tc.n)
			}
		})
	}
}