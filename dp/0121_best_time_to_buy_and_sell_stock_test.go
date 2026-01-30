package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected int
	}{
		{
			name:     "Example 1",
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			name:     "Example 2",
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0,
		},
		{
			name:     "Single day",
			prices:   []int{5},
			expected: 0,
		},
		{
			name:     "Two days increasing",
			prices:   []int{1, 2},
			expected: 1,
		},
		{
			name:     "Two days decreasing",
			prices:   []int{2, 1},
			expected: 0,
		},
		{
			name:     "Multiple peaks",
			prices:   []int{3, 2, 6, 5, 0, 3},
			expected: 4,
		},
		{
			name:     "All same prices",
			prices:   []int{5, 5, 5, 5, 5},
			expected: 0,
		},
		{
			name:     "Increasing prices",
			prices:   []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "Decreasing then increasing",
			prices:   []int{10, 9, 8, 7, 6, 7, 8, 9, 10},
			expected: 4,
		},
		{
			name:     "Large profit at end",
			prices:   []int{3, 2, 1, 1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "Large profit in middle",
			prices:   []int{5, 4, 3, 2, 1, 2, 3, 4, 3, 2, 1},
			expected: 3,
		},
		{
			name:     "Empty prices",
			prices:   []int{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxProfit(tt.prices)
			assert.Equal(t, tt.expected, result,
				"MaxProfit(%v) = %d, expected %d",
				tt.prices, result, tt.expected)
		})
	}
}

func TestMaxProfit_EdgeCases(t *testing.T) {
	t.Run("Very large input", func(t *testing.T) {
		// Create a large array where profit should be 9999
		prices := make([]int, 100000)
		for i := range prices {
			// Create decreasing prices from 10000 down to 1
			prices[i] = 10000 - (i % 10000)
		}
		prices[0] = 1
		prices[99999] = 10000

		result := MaxProfit(prices)
		assert.Equal(t, 9999, result)
	})

	t.Run("All zeros", func(t *testing.T) {
		prices := make([]int, 1000)
		result := MaxProfit(prices)
		assert.Equal(t, 0, result)
	})
}

func BenchmarkMaxProfit(b *testing.B) {
	// Create a large slice for benchmarking
	prices := make([]int, 100000)
	for i := range prices {
		prices[i] = i % 10000 // Mix of values
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxProfit(prices)
	}
}