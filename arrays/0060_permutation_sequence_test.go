package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPermutation(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		k        int
		expected string
	}{
		{
			name:     "Example 1: n=3, k=3",
			n:        3,
			k:        3,
			expected: "213",
		},
		{
			name:     "Example 2: n=4, k=9",
			n:        4,
			k:        9,
			expected: "2314",
		},
		{
			name:     "Example 3: n=3, k=1",
			n:        3,
			k:        1,
			expected: "123",
		},
		{
			name:     "n=1, k=1",
			n:        1,
			k:        1,
			expected: "1",
		},
		{
			name:     "n=2, k=1",
			n:        2,
			k:        1,
			expected: "12",
		},
		{
			name:     "n=2, k=2",
			n:        2,
			k:        2,
			expected: "21",
		},
		{
			name:     "n=3, k=6 (last permutation)",
			n:        3,
			k:        6,
			expected: "321",
		},
		{
			name:     "n=4, k=1 (first permutation)",
			n:        4,
			k:        1,
			expected: "1234",
		},
		{
			name:     "n=4, k=24 (last permutation)",
			n:        4,
			k:        24,
			expected: "4321",
		},
		{
			name:     "n=5, k=16",
			n:        5,
			k:        16,
			expected: "14352",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetPermutation(tt.n, tt.k)
			assert.Equal(t, tt.expected, result,
				"GetPermutation(%d, %d) = %s, expected %s",
				tt.n, tt.k, result, tt.expected)
		})
	}
}

func TestGetPermutation_EdgeCases(t *testing.T) {
	t.Run("n=0 returns empty string", func(t *testing.T) {
		result := GetPermutation(0, 1)
		assert.Empty(t, result)
	})

	t.Run("k=0 returns empty string", func(t *testing.T) {
		result := GetPermutation(3, 0)
		assert.Empty(t, result)
	})

	t.Run("k larger than n! returns empty string", func(t *testing.T) {
		result := GetPermutation(3, 7) // 3! = 6, so k=7 is invalid
		assert.Empty(t, result)
	})

	t.Run("n=9, k=362880 (max valid)", func(t *testing.T) {
		result := GetPermutation(9, 362880) // 9! = 362880
		assert.Equal(t, "987654321", result)
	})

	t.Run("n=9, k=1 (min valid)", func(t *testing.T) {
		result := GetPermutation(9, 1)
		assert.Equal(t, "123456789", result)
	})
}

func TestGetPermutation_Consistency(t *testing.T) {
	// Test that all permutations are generated correctly
	n := 4
	totalPermutations := 24 // 4! = 24

	// Generate all permutations and verify they're unique
	seen := make(map[string]bool)
	for k := 1; k <= totalPermutations; k++ {
		perm := GetPermutation(n, k)
		assert.False(t, seen[perm], "Duplicate permutation %s for k=%d", perm, k)
		seen[perm] = true
		assert.Len(t, perm, n, "Permutation should have length %d", n)
	}
	assert.Len(t, seen, totalPermutations, "Should generate all %d permutations", totalPermutations)
}

func BenchmarkGetPermutation(b *testing.B) {
	testCases := []struct {
		name string
		n    int
		k    int
	}{
		{"n=3,k=3", 3, 3},
		{"n=4,k=9", 4, 9},
		{"n=5,k=50", 5, 50},
		{"n=6,k=200", 6, 200},
		{"n=7,k=1000", 7, 1000},
		{"n=8,k=20000", 8, 20000},
		{"n=9,k=100000", 9, 100000},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				GetPermutation(tc.n, tc.k)
			}
		})
	}
}
