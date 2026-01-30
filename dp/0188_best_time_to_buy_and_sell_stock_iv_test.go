package dp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfitIV(t *testing.T) {
	tests := []struct {
		name     string
		k        int
		prices   []int
		expected int
	}{
		// Example test cases from LeetCode
		{
			name:     "Example 1",
			k:        2,
			prices:   []int{2, 4, 1},
			expected: 2,
		},
		{
			name:     "Example 2",
			k:        2,
			prices:   []int{3, 2, 6, 5, 0, 3},
			expected: 7,
		},
		{
			name:     "No transactions possible",
			k:        2,
			prices:   []int{5, 4, 3, 2, 1},
			expected: 0,
		},
		{
			name:     "Single transaction maximum",
			k:        1,
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			name:     "Unlimited transactions (k >= n/2)",
			k:        10,
			prices:   []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "Zero k",
			k:        0,
			prices:   []int{1, 2, 3, 4, 5},
			expected: 0,
		},
		{
			name:     "Single price",
			k:        2,
			prices:   []int{5},
			expected: 0,
		},
		{
			name:     "Empty prices",
			k:        2,
			prices:   []int{},
			expected: 0,
		},
		{
			name:     "Multiple peaks with k=2",
			k:        2,
			prices:   []int{3, 3, 5, 0, 0, 3, 1, 4},
			expected: 6,
		},
		{
			name:     "k=3 with 3 optimal transactions",
			k:        3,
			prices:   []int{2, 5, 7, 1, 4, 3, 1, 3},
			expected: 10, // Buy at 2 sell at 7 (5) + Buy at 1 sell at 4 (3) + Buy at 1 sell at 3 (2) = 10
		},
		{
			name:     "Large k with small price movements",
			k:        100,
			prices:   []int{1, 2, 1, 2, 1, 2, 1, 2},
			expected: 4, // 4 transactions of profit 1 each
		},
		{
			name:     "All same prices",
			k:        2,
			prices:   []int{5, 5, 5, 5, 5},
			expected: 0,
		},
		{
			name:     "k=1 with decreasing then increasing",
			k:        1,
			prices:   []int{10, 9, 8, 7, 6, 7, 8, 9, 10},
			expected: 4,
		},
		{
			name:     "k=2 with complex pattern",
			k:        2,
			prices:   []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0},
			expected: 13, // Buy at 1 sell at 7 (6) + Buy at 2 sell at 9 (7) = 13
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaxProfitIV(tt.k, tt.prices)
			assert.Equal(t, tt.expected, result,
				"MaxProfitIV(k=%d, prices=%v) = %d, expected %d",
				tt.k, tt.prices, result, tt.expected)
		})
	}
}

func TestMaxProfitIV_EdgeCases(t *testing.T) {
	t.Run("k equals maximum possible transactions", func(t *testing.T) {
		prices := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		k := len(prices) / 2 // 5 transactions maximum possible
		result := MaxProfitIV(k, prices)
		assert.Equal(t, 9, result) // Buy at 1, sell at 10 = 9
	})

	t.Run("k larger than maximum possible transactions", func(t *testing.T) {
		prices := []int{1, 2, 3, 4, 5}
		k := 10 // More than n/2
		result := MaxProfitIV(k, prices)
		assert.Equal(t, 4, result) // Unlimited transactions case
	})

	t.Run("Very large k with many prices", func(t *testing.T) {
		// Create prices with alternating pattern
		prices := make([]int, 1000)
		for i := range prices {
			if i%2 == 0 {
				prices[i] = 1
			} else {
				prices[i] = 2
			}
		}
		k := 500 // Exactly n/2
		result := MaxProfitIV(k, prices)
		// Each even->odd transition gives profit 1, there are 500 such transitions
		assert.Equal(t, 500, result)
	})

	t.Run("All zeros in prices", func(t *testing.T) {
		prices := make([]int, 100)
		k := 10
		result := MaxProfitIV(k, prices)
		assert.Equal(t, 0, result)
	})

	t.Run("Maximum constraint values", func(t *testing.T) {
		// Test with maximum constraints: k=100, prices length=1000
		prices := make([]int, 1000)
		for i := range prices {
			prices[i] = i % 1000 // Mix of values
		}
		k := 100
		result := MaxProfitIV(k, prices)
		// Should not panic and return a valid result
		assert.GreaterOrEqual(t, result, 0)
	})
}

func TestMaxProfitIV_PropertyBased(t *testing.T) {
	t.Run("Profit should be non-negative", func(t *testing.T) {
		// Random test: profit should never be negative
		testCases := []struct {
			k      int
			prices []int
		}{
			{k: 1, prices: []int{5, 4, 3, 2, 1}},
			{k: 2, prices: []int{1, 2, 3, 4, 5}},
			{k: 3, prices: []int{3, 2, 6, 5, 0, 3}},
			{k: 4, prices: []int{1, 3, 2, 4, 3, 5}},
		}

		for _, tc := range testCases {
			result := MaxProfitIV(tc.k, tc.prices)
			assert.GreaterOrEqual(t, result, 0,
				"MaxProfitIV(k=%d, prices=%v) = %d should be >= 0",
				tc.k, tc.prices, result)
		}
	})

	t.Run("Increasing k should not decrease profit", func(t *testing.T) {
		prices := []int{3, 2, 6, 5, 0, 3, 2, 4, 1, 5}
		prevProfit := -1

		for k := 0; k <= 5; k++ {
			profit := MaxProfitIV(k, prices)
			if prevProfit != -1 {
				assert.GreaterOrEqual(t, profit, prevProfit,
					"Profit with k=%d (%d) should be >= profit with k=%d (%d)",
					k, profit, k-1, prevProfit)
			}
			prevProfit = profit
		}
	})

	t.Run("Monotonic prices test", func(t *testing.T) {
		// For monotonic increasing prices, profit should be last - first for k=1
		prices := []int{1, 2, 3, 4, 5}
		result := MaxProfitIV(1, prices)
		assert.Equal(t, 4, result, "For monotonic increasing prices, profit should be last - first")

		// For monotonic decreasing prices, profit should be 0
		prices2 := []int{5, 4, 3, 2, 1}
		result2 := MaxProfitIV(1, prices2)
		assert.Equal(t, 0, result2, "For monotonic decreasing prices, profit should be 0")
	})
}

func BenchmarkMaxProfitIV(b *testing.B) {
	benchmarks := []struct {
		name   string
		k      int
		prices []int
	}{
		{
			name:   "Small input k=2",
			k:      2,
			prices: []int{3, 2, 6, 5, 0, 3, 2, 4, 1, 5},
		},
		{
			name:   "Medium input k=5",
			k:      5,
			prices: make([]int, 100),
		},
		{
			name:   "Large input k=50",
			k:      50,
			prices: make([]int, 1000),
		},
		{
			name:   "Unlimited transactions case",
			k:      1000,
			prices: make([]int, 100),
		},
	}

	// Initialize benchmark prices
	for i := range benchmarks[1].prices {
		benchmarks[1].prices[i] = i % 100
	}
	for i := range benchmarks[2].prices {
		benchmarks[2].prices[i] = i % 1000
	}
	for i := range benchmarks[3].prices {
		benchmarks[3].prices[i] = i % 100
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MaxProfitIV(bm.k, bm.prices)
			}
		})
	}
}

func BenchmarkMaxProfitIV_Comparison(b *testing.B) {
	// Compare performance for different k values with same prices
	prices := make([]int, 500)
	for i := range prices {
		prices[i] = i % 200
	}

	kValues := []int{1, 2, 5, 10, 25, 50, 100}
	for _, k := range kValues {
		b.Run(fmt.Sprintf("k=%d", k), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MaxProfitIV(k, prices)
			}
		})
	}
}