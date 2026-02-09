package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinEatingSpeed(t *testing.T) {
	tests := []struct {
		name     string
		piles    []int
		h        int
		expected int
	}{
		{
			name:     "Example 1",
			piles:    []int{3, 6, 7, 11},
			h:        8,
			expected: 4,
		},
		{
			name:     "Example 2",
			piles:    []int{30, 11, 23, 4, 20},
			h:        5,
			expected: 30,
		},
		{
			name:     "Example 3",
			piles:    []int{30, 11, 23, 4, 20},
			h:        6,
			expected: 23,
		},
		{
			name:     "Single pile",
			piles:    []int{10},
			h:        1,
			expected: 10,
		},
		{
			name:     "Single pile with more time",
			piles:    []int{10},
			h:        5,
			expected: 2, // ceil(10/5) = 2
		},
		{
			name:     "All piles same size",
			piles:    []int{5, 5, 5, 5, 5},
			h:        5,
			expected: 5, // One pile per hour
		},
		{
			name:     "All piles same size with more time",
			piles:    []int{5, 5, 5, 5, 5},
			h:        10,
			expected: 3, // ceil(5/2) = 3 per hour, but need ceil(25/10) = 3
		},
		{
			name:     "Empty piles",
			piles:    []int{},
			h:        5,
			expected: 0,
		},
		{
			name:     "Exactly enough hours",
			piles:    []int{1, 2, 3, 4, 5},
			h:        5,
			expected: 5, // Need to eat largest pile in one hour
		},
		{
			name:     "More hours than piles",
			piles:    []int{1, 2, 3, 4, 5},
			h:        10,
			expected: 2, // Speed 1: ceil(1/1)+ceil(2/1)+ceil(3/1)+ceil(4/1)+ceil(5/1)=15 > 10
			// Speed 2: ceil(1/2)+ceil(2/2)+ceil(3/2)+ceil(4/2)+ceil(5/2)=1+1+2+2+3=9 ≤ 10
		},
		{
			name:     "Large numbers",
			piles:    []int{1000000000, 1000000000, 1000000000},
			h:        3,
			expected: 1000000000, // One pile per hour
		},
		{
			name:     "Large numbers with more time",
			piles:    []int{1000000000, 1000000000, 1000000000},
			h:        6,
			expected: 500000000, // ceil(1e9/2) = 500M
		},
		{
			name:     "Mixed sizes",
			piles:    []int{1, 10, 100, 1000},
			h:        10,
			expected: 143, // Speed 100: ceil(1/100)+ceil(10/100)+ceil(100/100)+ceil(1000/100)=1+1+1+10=13 > 10
			// Speed 143: ceil(1/143)+ceil(10/143)+ceil(100/143)+ceil(1000/143)=1+1+1+7=10 ≤ 10
		},
		{
			name:     "Complex case 1",
			piles:    []int{5, 8, 12, 15, 20},
			h:        7,
			expected: 12, // Optimal speed
		},
		{
			name:     "Complex case 2",
			piles:    []int{2, 4, 6, 8, 10},
			h:        10,
			expected: 4, // Speed 2: ceil(2/2)+ceil(4/2)+ceil(6/2)+ceil(8/2)+ceil(10/2)=1+2+3+4+5=15 > 10
			// Speed 4: ceil(2/4)+ceil(4/4)+ceil(6/4)+ceil(8/4)+ceil(10/4)=1+1+2+2+3=9 ≤ 10
		},
		{
			name:     "One very large pile",
			piles:    []int{100, 1, 1, 1, 1},
			h:        5,
			expected: 100, // Speed 34: ceil(100/34)+ceil(1/34)*4=3+4=7 > 5
			// Speed 100: ceil(100/100)+ceil(1/100)*4=1+4=5 ≤ 5
		},
		{
			name:     "Minimum speed needed",
			piles:    []int{1, 1, 1, 1, 1},
			h:        5,
			expected: 1,
		},
		{
			name:     "Just enough time",
			piles:    []int{3, 6, 7, 11},
			h:        4,
			expected: 11, // Need to eat each pile in one hour
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinEatingSpeed(tt.piles, tt.h)
			assert.Equal(t, tt.expected, result,
				"MinEatingSpeed(%v, %d) = %d, expected %d",
				tt.piles, tt.h, result, tt.expected)
		})
	}
}

func TestMinEatingSpeedOptimized(t *testing.T) {
	tests := []struct {
		name     string
		piles    []int
		h        int
		expected int
	}{
		{
			name:     "Example 1",
			piles:    []int{3, 6, 7, 11},
			h:        8,
			expected: 4,
		},
		{
			name:     "Example 2",
			piles:    []int{30, 11, 23, 4, 20},
			h:        5,
			expected: 30,
		},
		{
			name:     "Single pile",
			piles:    []int{10},
			h:        1,
			expected: 10,
		},
		{
			name:     "More hours than piles",
			piles:    []int{1, 2, 3, 4, 5},
			h:        10,
			expected: 2, // Speed 1: 15 hours > 10, Speed 2: 9 hours ≤ 10
		},
		{
			name:     "Large numbers",
			piles:    []int{1000000000, 1000000000, 1000000000},
			h:        6,
			expected: 500000000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinEatingSpeedOptimized(tt.piles, tt.h)
			assert.Equal(t, tt.expected, result,
				"MinEatingSpeedOptimized(%v, %d) = %d, expected %d",
				tt.piles, tt.h, result, tt.expected)
		})
	}
}

func TestMinEatingSpeed_Consistency(t *testing.T) {
	// Test that both implementations give the same result
	testCases := []struct {
		piles []int
		h     int
	}{
		{[]int{3, 6, 7, 11}, 8},
		{[]int{30, 11, 23, 4, 20}, 5},
		{[]int{30, 11, 23, 4, 20}, 6},
		{[]int{10}, 1},
		{[]int{10}, 5},
		{[]int{5, 5, 5, 5, 5}, 5},
		{[]int{}, 5},
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{1, 2, 3, 4, 5}, 10},
		{[]int{1000000000, 1000000000, 1000000000}, 3},
		{[]int{1000000000, 1000000000, 1000000000}, 6},
		{[]int{1, 10, 100, 1000}, 10},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result1 := MinEatingSpeed(tc.piles, tc.h)
			result2 := MinEatingSpeedOptimized(tc.piles, tc.h)
			assert.Equal(t, result1, result2,
				"Both implementations should give same result for piles=%v, h=%d: %d vs %d",
				tc.piles, tc.h, result1, result2)
		})
	}
}

func TestMinEatingSpeed_EdgeCases(t *testing.T) {
	t.Run("Empty piles array", func(t *testing.T) {
		result := MinEatingSpeed([]int{}, 5)
		assert.Equal(t, 0, result)
	})

	t.Run("Zero hours (invalid per problem)", func(t *testing.T) {
		// Problem guarantees h >= piles.length, but test robustness
		result := MinEatingSpeed([]int{1, 2, 3}, 0)
		// With 0 hours, need infinite speed
		// Our implementation will return maxPile
		assert.Equal(t, 3, result) // max pile
	})

	t.Run("Hours less than number of piles (invalid per problem)", func(t *testing.T) {
		// Problem guarantees h >= piles.length
		result := MinEatingSpeed([]int{1, 2, 3}, 2)
		// Need to eat 3 piles in 2 hours, so speed must be at least max pile
		assert.Equal(t, 3, result)
	})

	t.Run("Very large piles with small h", func(t *testing.T) {
		piles := make([]int, 100)
		for i := range piles {
			piles[i] = 1000000
		}
		result := MinEatingSpeed(piles, 100)
		// Need to eat each pile in one hour
		assert.Equal(t, 1000000, result)
	})

	t.Run("Very large piles with large h", func(t *testing.T) {
		piles := make([]int, 100)
		for i := range piles {
			piles[i] = 1000000
		}
		result := MinEatingSpeed(piles, 200)
		// Can take 2 hours per pile
		assert.Equal(t, 500000, result) // ceil(1,000,000 / 2) = 500,000
	})

	t.Run("All piles size 1", func(t *testing.T) {
		piles := make([]int, 1000)
		for i := range piles {
			piles[i] = 1
		}
		result := MinEatingSpeed(piles, 1000)
		assert.Equal(t, 1, result)
	})

	t.Run("All piles size 1 with extra time", func(t *testing.T) {
		piles := make([]int, 1000)
		for i := range piles {
			piles[i] = 1
		}
		result := MinEatingSpeed(piles, 2000)
		// Can eat 2 piles per hour (0.5 hours per pile)
		// But speed must be integer, so need speed 1
		assert.Equal(t, 1, result)
	})
}

func TestCanEatAll(t *testing.T) {
	tests := []struct {
		name     string
		piles    []int
		h        int
		k        int
		expected bool
	}{
		{
			name:     "Can eat with speed 4 in 8 hours",
			piles:    []int{3, 6, 7, 11},
			h:        8,
			k:        4,
			expected: true,
		},
		{
			name:     "Cannot eat with speed 3 in 8 hours",
			piles:    []int{3, 6, 7, 11},
			h:        8,
			k:        3,
			expected: false,
		},
		{
			name:     "Exact fit",
			piles:    []int{5, 5, 5},
			h:        3,
			k:        5,
			expected: true,
		},
		{
			name:     "More time than needed",
			piles:    []int{5, 5, 5},
			h:        6,
			k:        3,
			expected: true, // ceil(5/3)=2 hours per pile, total 6 hours
		},
		{
			name:     "Speed too slow",
			piles:    []int{10, 10, 10},
			h:        5,
			k:        5,
			expected: false, // ceil(10/5)=2 hours per pile, total 6 hours > 5
		},
		{
			name:     "Speed exactly enough",
			piles:    []int{7, 7, 7},
			h:        6,
			k:        4,
			expected: true, // ceil(7/4)=2 hours per pile, total 6 hours
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := canEatAll(tt.piles, tt.h, tt.k)
			assert.Equal(t, tt.expected, result,
				"canEatAll(%v, %d, %d) = %v, expected %v",
				tt.piles, tt.h, tt.k, result, tt.expected)
		})
	}
}

func BenchmarkMinEatingSpeed(b *testing.B) {
	// Create test data
	piles := make([]int, 1000)
	for i := range piles {
		piles[i] = (i % 100) + 1 // Piles from 1 to 100
	}
	h := 1000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinEatingSpeed(piles, h)
	}
}

func BenchmarkMinEatingSpeedOptimized(b *testing.B) {
	piles := make([]int, 1000)
	for i := range piles {
		piles[i] = (i % 100) + 1
	}
	h := 1000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinEatingSpeedOptimized(piles, h)
	}
}

func BenchmarkMinEatingSpeed_WorstCase(b *testing.B) {
	// Worst case: binary search goes through many iterations
	piles := make([]int, 1000)
	for i := range piles {
		piles[i] = 1000000 // All large piles
	}
	h := 1000 // Exactly one hour per pile

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinEatingSpeed(piles, h)
	}
}

func BenchmarkMinEatingSpeed_BestCase(b *testing.B) {
	// Best case: h is large, so lower bound is close to answer
	piles := make([]int, 1000)
	for i := range piles {
		piles[i] = i + 1
	}
	h := 100000 // Lots of time, can eat slowly

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MinEatingSpeed(piles, h)
	}
}