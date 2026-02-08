package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWiggleMaxLength(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 7, 4, 9, 2, 5},
			expected: 6,
		},
		{
			name:     "Example 2",
			nums:     []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8},
			expected: 7,
		},
		{
			name:     "Example 3",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: 2,
		},
		{
			name:     "Single element",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "All same elements",
			nums:     []int{1, 1, 1, 1},
			expected: 1,
		},
		{
			name:     "Two different elements",
			nums:     []int{1, 2},
			expected: 2,
		},
		{
			name:     "Complex case",
			nums:     []int{2, 1, 4, 1, 1, 3, 1, 5, 1, 2},
			expected: 9, // Fixed: correct answer is 9
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WiggleMaxLength(tt.nums)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkWiggleMaxLength(b *testing.B) {
	nums := []int{1, 7, 4, 9, 2, 5, 1, 17, 5, 10, 13, 15, 10, 5, 16, 8}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		WiggleMaxLength(nums)
	}
}
