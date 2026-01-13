package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfitIII(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected int
	}{
		{
			name:     "Example 1",
			prices:   []int{3, 3, 5, 0, 0, 3, 1, 4},
			expected: 6,
		},
		{
			name:     "Example 2",
			prices:   []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "Example 3",
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
			name:     "One big transaction is best",
			prices:   []int{1, 5, 3, 8},
			expected: 7, // Buy at 1, sell at 8 = 7
		},
		{
			name:     "Two transactions better than one",
			prices:   []int{1, 5, 3, 8, 2, 10},
			expected: 12, // (5-1) + (10-2) = 4 + 8 = 12
		},
		{
			name:     "All same prices",
			prices:   []int{5, 5, 5, 5, 5},
			expected: 0,
		},
		{
			name:     "Empty prices",
			prices:   []int{},
			expected: 0,
		},
		{
			name:     "Three peaks",
			prices:   []int{1, 4, 2, 5, 3, 6},
			expected: 8, // (4-1) + (6-3) = 3 + 5 = 8
		},
		{
			name:     "Complex pattern",
			prices:   []int{2, 1, 4, 5, 2, 9, 7},
			expected: 11, // (5-1) + (9-2) = 4 + 7 = 11
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxProfitIII(tt.prices)
			assert.Equal(t, tt.expected, result,
				"MaxProfitIII(%v) = %d, expected %d",
				tt.prices, result, tt.expected)
		})
	}
}

func TestMaxProfitIII_EdgeCases(t *testing.T) {
	t.Run("Very large input with single transaction", func(t *testing.T) {
		// Create a large array where best is one transaction
		prices := make([]int, 100000)
		for i := range prices {
			prices[i] = i // Increasing prices
		}
		// Best single transaction: buy at 0, sell at 99999 = 99999
		result := MaxProfitIII(prices)
		assert.Equal(t, 99999, result)
	})

	t.Run("Very large input with two transactions", func(t *testing.T) {
		// Create a large array where two transactions are better
		prices := make([]int, 100000)
		for i := range prices {
			if i < 50000 {
				prices[i] = i
			} else {
				prices[i] = 100000 - i
			}
		}
		// Best: (49999-0) + (49999-0) = 49999 + 49999 = 99998
		result := MaxProfitIII(prices)
		assert.Equal(t, 99998, result)
	})

	t.Run("All zeros", func(t *testing.T) {
		prices := make([]int, 1000)
		result := MaxProfitIII(prices)
		assert.Equal(t, 0, result)
	})

	t.Run("Maximum price values", func(t *testing.T) {
		prices := make([]int, 1000)
		for i := range prices {
			prices[i] = 100000
		}
		result := MaxProfitIII(prices)
		assert.Equal(t, 0, result)
	})
}

func BenchmarkMaxProfitIII(b *testing.B) {
	// Create a large slice for benchmarking
	prices := make([]int, 100000)
	for i := range prices {
		prices[i] = i % 100000 // Mix of values
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxProfitIII(prices)
	}
}