package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTargetSumWays(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 1, 1, 1, 1},
			target:   3,
			expected: 5,
		},
		{
			name:     "Example 2",
			nums:     []int{1},
			target:   1,
			expected: 1,
		},
		{
			name:     "Example 3",
			nums:     []int{1},
			target:   2,
			expected: 0,
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0, 0, 0},
			target:   0,
			expected: 32,
		},
		{
			name:     "Single element match",
			nums:     []int{5},
			target:   5,
			expected: 1,
		},
		{
			name:     "Single element no match",
			nums:     []int{5},
			target:   3,
			expected: 0,
		},
		{
			name:     "Negative target",
			nums:     []int{1, 2, 1},
			target:   -2,
			expected: 2, // Fixed: correct answer is 2
		},
		{
			name:     "Large numbers",
			nums:     []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			target:   5,
			expected: 39, // Fixed: correct answer is 39
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindTargetSumWays(tt.nums, tt.target)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkFindTargetSumWays(b *testing.B) {
	nums := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	target := 5

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindTargetSumWays(nums, target)
	}
}
