package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfitII(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected int
	}{
		{
			name:     "Example 1",
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 7,
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
			name:     "Multiple peaks and valleys",
			prices:   []int{3, 2, 6, 5, 0, 3},
			expected: 7, // (6-2) + (3-0) = 4 + 3 = 7
		},
		{
			name:     "All same prices",
			prices:   []int{5, 5, 5, 5, 5},
			expected: 0,
		},
		{
			name:     "Large fluctuations",
			prices:   []int{1, 5, 3, 8, 2, 10},
			expected: 17, // (5-1)=4, (8-3)=5, (10-2)=8, total=17
		},
		{
			name:     "Empty prices",
			prices:   []int{},
			expected: 0,
		},
		{
			name:     "Single valley then peak",
			prices:   []int{10, 9, 8, 7, 6, 7, 8, 9, 10},
			expected: 4, // (7-6) + (8-7) + (9-8) + (10-9) = 1+1+1+1 = 4
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxProfitII(tt.prices)
			assert.Equal(t, tt.expected, result,
				"MaxProfitII(%v) = %d, expected %d",
				tt.prices, result, tt.expected)
		})
	}
}

func TestMaxProfitII_EdgeCases(t *testing.T) {
	t.Run("Very large input", func(t *testing.T) {
		// Create a large array with alternating pattern
		prices := make([]int, 30000)
		for i := range prices {
			if i%2 == 0 {
				prices[i] = 100
			} else {
				prices[i] = 101
			}
		}

		result := MaxProfitII(prices)
		// Each even->odd transition gives profit of 1
		// With 30000 elements, indices 0-29999
		// Transitions: 0→1, 1→2, 2→3, ..., 29998→29999
		// Even→odd transitions: 0→1, 2→3, ..., 29998→29999
		// That's 15000 transitions (0, 2, 4, ..., 29998)
		expected := 15000
		assert.Equal(t, expected, result)
	})

	t.Run("All zeros", func(t *testing.T) {
		prices := make([]int, 1000)
		result := MaxProfitII(prices)
		assert.Equal(t, 0, result)
	})

	t.Run("Maximum constraint values", func(t *testing.T) {
		prices := make([]int, 30000)
		for i := range prices {
			prices[i] = 10000 - i%100 // Decreasing pattern with small fluctuations
		}
		// Just ensure it runs without error
		result := MaxProfitII(prices)
		assert.True(t, result >= 0)
	})
}

func BenchmarkMaxProfitII(b *testing.B) {
	// Create a large slice for benchmarking
	prices := make([]int, 30000)
	for i := range prices {
		prices[i] = i % 10000 // Mix of values
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MaxProfitII(prices)
	}
}