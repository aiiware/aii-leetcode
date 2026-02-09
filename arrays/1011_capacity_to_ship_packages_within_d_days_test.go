package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShipWithinDays(t *testing.T) {
	tests := []struct {
		name     string
		weights  []int
		days     int
		expected int
	}{
		{
			name:     "Example 1",
			weights:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			days:     5,
			expected: 15,
		},
		{
			name:     "Example 2",
			weights:  []int{3, 2, 2, 4, 1, 4},
			days:     3,
			expected: 6,
		},
		{
			name:     "Example 3",
			weights:  []int{1, 2, 3, 1, 1},
			days:     4,
			expected: 3,
		},
		{
			name:     "Single package",
			weights:  []int{10},
			days:     1,
			expected: 10,
		},
		{
			name:     "Single package, multiple days",
			weights:  []int{10},
			days:     5,
			expected: 10,
		},
		{
			name:     "All packages same weight",
			weights:  []int{5, 5, 5, 5, 5},
			days:     2,
			expected: 15, // Need to ship 3 packages (15) one day, 2 packages (10) next day
		},
		{
			name:     "Large weights",
			weights:  []int{100, 200, 300, 400, 500},
			days:     2,
			expected: 900, // [100,200,300,400]=1000 vs [100,200,300]=600 + [400,500]=900
		},
		{
			name:     "Many days available",
			weights:  []int{1, 2, 3, 4, 5},
			days:     5,
			expected: 5, // Can ship each package separately
		},
		{
			name:     "Few days available",
			weights:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			days:     2,
			expected: 28, // With capacity 28: Day1: 1-7=28, Day2: 8-10=27
		},
		{
			name:     "Empty weights",
			weights:  []int{},
			days:     5,
			expected: 0,
		},
		{
			name:     "One day only",
			weights:  []int{1, 2, 3, 4, 5},
			days:     1,
			expected: 15, // Sum of all weights
		},
		{
			name:     "Days equals number of packages",
			weights:  []int{1, 2, 3, 4, 5},
			days:     5,
			expected: 5, // Max weight
		},
		{
			name:     "Complex case 1",
			weights:  []int{10, 20, 30, 40, 50, 60, 70, 80, 90},
			days:     3,
			expected: 170, // [10-80]=360/3=120, but need 170
		},
		{
			name:     "Complex case 2",
			weights:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			days:     4,
			expected: 36, // Minimum capacity that works
		},
		{
			name:     "Large numbers",
			weights:  []int{1000, 2000, 3000, 4000, 5000},
			days:     2,
			expected: 9000, // [1000,2000,3000,4000]=10000 vs [1000,2000,3000]=6000 + [4000,5000]=9000
		},
		{
			name:     "Weights in descending order",
			weights:  []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			days:     3,
			expected: 21, // Minimum capacity that works
		},
		{
			name:     "All weights 1",
			weights:  []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			days:     3,
			expected: 4, // 10 weights / 3 days = ceil(10/3) = 4
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ShipWithinDays(tt.weights, tt.days)
			assert.Equal(t, tt.expected, result,
				"ShipWithinDays(%v, %d) = %d, expected %d",
				tt.weights, tt.days, result, tt.expected)
		})
	}
}

func TestShipWithinDaysOptimized(t *testing.T) {
	tests := []struct {
		name     string
		weights  []int
		days     int
		expected int
	}{
		{
			name:     "Example 1",
			weights:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			days:     5,
			expected: 15,
		},
		{
			name:     "Example 2",
			weights:  []int{3, 2, 2, 4, 1, 4},
			days:     3,
			expected: 6,
		},
		{
			name:     "Single package",
			weights:  []int{10},
			days:     1,
			expected: 10,
		},
		{
			name:     "One day only",
			weights:  []int{1, 2, 3, 4, 5},
			days:     1,
			expected: 15,
		},
		{
			name:     "Days equals number of packages",
			weights:  []int{1, 2, 3, 4, 5},
			days:     5,
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ShipWithinDaysOptimized(tt.weights, tt.days)
			assert.Equal(t, tt.expected, result,
				"ShipWithinDaysOptimized(%v, %d) = %d, expected %d",
				tt.weights, tt.days, result, tt.expected)
		})
	}
}

func TestShipWithinDays_Consistency(t *testing.T) {
	// Test that both implementations give the same result
	testCases := []struct {
		weights []int
		days    int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5},
		{[]int{3, 2, 2, 4, 1, 4}, 3},
		{[]int{1, 2, 3, 1, 1}, 4},
		{[]int{10}, 1},
		{[]int{5, 5, 5, 5, 5}, 2},
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{}, 5},
		{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 3},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result1 := ShipWithinDays(tc.weights, tc.days)
			result2 := ShipWithinDaysOptimized(tc.weights, tc.days)
			assert.Equal(t, result1, result2,
				"Both implementations should give same result for weights=%v, days=%d: %d vs %d",
				tc.weights, tc.days, result1, result2)
		})
	}
}

func TestShipWithinDays_EdgeCases(t *testing.T) {
	t.Run("Empty weights array", func(t *testing.T) {
		result := ShipWithinDays([]int{}, 5)
		assert.Equal(t, 0, result)
	})

	t.Run("Zero days (invalid per problem)", func(t *testing.T) {
		// Problem guarantees days >= 1, but test robustness
		result := ShipWithinDays([]int{1, 2, 3}, 0)
		// With 0 days, we need infinite capacity
		// Our implementation returns totalWeight when days=0
		assert.Equal(t, 6, result) // total weight
	})

	t.Run("Negative weights (shouldn't happen)", func(t *testing.T) {
		// Test robustness with invalid input
		result := ShipWithinDays([]int{-1, 2, 3}, 2)
		// Our implementation doesn't validate weights, so it will process them
		// This is acceptable since problem guarantees positive weights
		assert.Greater(t, result, 0)
	})

	t.Run("Very large weights", func(t *testing.T) {
		weights := make([]int, 1000)
		for i := range weights {
			weights[i] = 1000000
		}
		result := ShipWithinDays(weights, 10)
		// Each day needs to ship 100 packages of 1,000,000 each
		// 1000 packages / 10 days = 100 packages/day
		expected := 100 * 1000000
		assert.Equal(t, expected, result)
	})

	t.Run("Days greater than number of packages", func(t *testing.T) {
		result := ShipWithinDays([]int{10, 20, 30}, 5)
		// With 5 days for 3 packages, we can ship each separately
		// So capacity needs to be at least max weight
		assert.Equal(t, 30, result)
	})

	t.Run("Single heavy package with light packages", func(t *testing.T) {
		result := ShipWithinDays([]int{100, 1, 1, 1, 1, 1}, 2)
		// Need capacity at least 100 (for the heavy package)
		// Best partition: [100] and [1,1,1,1,1]=5
		assert.Equal(t, 100, result)
	})
}

func TestCanShipWithinDays(t *testing.T) {
	tests := []struct {
		name     string
		weights  []int
		days     int
		capacity int
		expected bool
	}{
		{
			name:     "Can ship with exact capacity",
			weights:  []int{1, 2, 3, 4, 5},
			days:     3,
			capacity: 6,
			expected: true, // [1,2,3]=6, [4]=4, [5]=5
		},
		{
			name:     "Cannot ship with insufficient capacity",
			weights:  []int{1, 2, 3, 4, 5},
			days:     2,
			capacity: 7,
			expected: false, // Need at least 8: [1,2,3,4]=10 vs [1,2,3]=6 + [4,5]=9
		},
		{
			name:     "Capacity too small for single package",
			weights:  []int{10, 1, 1},
			days:     3,
			capacity: 5,
			expected: false, // Can't ship the 10
		},
		{
			name:     "Exactly enough days",
			weights:  []int{1, 1, 1, 1, 1},
			days:     5,
			capacity: 1,
			expected: true,
		},
		{
			name:     "More days than needed",
			weights:  []int{1, 1, 1, 1, 1},
			days:     10,
			capacity: 1,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canShipWithinDays(tt.weights, tt.days, tt.capacity)
			assert.Equal(t, tt.expected, result,
				"canShipWithinDays(%v, %d, %d) = %v, expected %v",
				tt.weights, tt.days, tt.capacity, result, tt.expected)
		})
	}
}

func BenchmarkShipWithinDays(b *testing.B) {
	// Create test data
	weights := make([]int, 1000)
	for i := range weights {
		weights[i] = (i % 100) + 1 // Weights from 1 to 100
	}
	days := 10

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ShipWithinDays(weights, days)
	}
}

func BenchmarkShipWithinDaysOptimized(b *testing.B) {
	weights := make([]int, 1000)
	for i := range weights {
		weights[i] = (i % 100) + 1
	}
	days := 10

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ShipWithinDaysOptimized(weights, days)
	}
}

func BenchmarkShipWithinDays_WorstCase(b *testing.B) {
	// Worst case: binary search goes through many iterations
	weights := make([]int, 5000)
	for i := range weights {
		weights[i] = 1 // All weights 1
	}
	days := 1 // Forces capacity = sum(weights) = 5000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ShipWithinDays(weights, days)
	}
}

func BenchmarkShipWithinDays_BestCase(b *testing.B) {
	// Best case: days equals number of packages
	weights := make([]int, 1000)
	for i := range weights {
		weights[i] = i + 1
	}
	days := 1000 // Each package gets its own day

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ShipWithinDays(weights, days)
	}
}