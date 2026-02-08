package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfitWithCooldown(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected int
	}{
		{
			name:     "Example 1",
			prices:   []int{1, 2, 3, 0, 2},
			expected: 3,
		},
		{
			name:     "Example 2",
			prices:   []int{2, 1, 2, 0, 1},
			expected: 1, // Fixed: correct answer is 1, not 2 (cooldown prevents buying at day 3)
		},
		{
			name:     "Single price",
			prices:   []int{5},
			expected: 0,
		},
		{
			name:     "Empty prices",
			prices:   []int{},
			expected: 0,
		},
		{
			name:     "No profit possible",
			prices:   []int{5, 4, 3, 2, 1},
			expected: 0,
		},
		{
			name:     "Profit with cooldown",
			prices:   []int{1, 2, 4, 0, 5},
			expected: 6,
		},
		{
			name:     "Large numbers",
			prices:   []int{2, 5, 1, 3, 7, 4},
			expected: 7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxProfitWithCooldown(tt.prices)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkMaxProfitWithCooldown(b *testing.B) {
	prices := []int{1, 2, 3, 0, 2, 5, 1, 3, 7, 4, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxProfitWithCooldown(prices)
	}
}
