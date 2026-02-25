package dp

import (
	"testing"
)

func TestMinCostClimbingStairs(t *testing.T) {
	tests := []struct {
		name     string
		cost     []int
		expected int
	}{
		{
			name:     "Example 1",
			cost:     []int{10, 15, 20},
			expected: 15,
		},
		{
			name:     "Example 2",
			cost:     []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			expected: 6,
		},
		{
			name:     "Two steps",
			cost:     []int{10, 15},
			expected: 10, // Start from step 0 (cost 10) or step 1 (cost 15) -> min is 10
		},
		{
			name:     "Three steps ascending",
			cost:     []int{1, 2, 3},
			expected: 2, // Start from step 0: 1 + 3 = 4, start from step 1: 2 -> min is 2
		},
		{
			name:     "Three steps descending",
			cost:     []int{3, 2, 1},
			expected: 2, // Start from step 0: 3 + 1 = 4, start from step 1: 2 -> min is 2
		},
		{
			name:     "All zeros",
			cost:     []int{0, 0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "Single high cost in middle",
			cost:     []int{1, 100, 1, 1, 1},
			expected: 3, // Path: 0 -> 2 -> 4 -> top (1 + 1 + 1 = 3)
		},
		{
			name:     "Alternating high costs",
			cost:     []int{100, 1, 100, 1, 100},
			expected: 2, // Path: 1 -> 3 -> top (1 + 1 = 2)
		},
		{
			name:     "Minimum length array",
			cost:     []int{5, 10},
			expected: 5,
		},
		{
			name:     "Large array with pattern",
			cost:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: 25, // Path: 0 -> 2 -> 4 -> 6 -> 8 -> top (1 + 3 + 5 + 7 + 9 = 25)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minCostClimbingStairs(tt.cost)
			if result != tt.expected {
				t.Errorf("minCostClimbingStairs(%v) = %d, expected %d", tt.cost, result, tt.expected)
			}
		})
	}
}

func TestMinCostClimbingStairsDP(t *testing.T) {
	tests := []struct {
		name     string
		cost     []int
		expected int
	}{
		{
			name:     "Example 1",
			cost:     []int{10, 15, 20},
			expected: 15,
		},
		{
			name:     "Example 2",
			cost:     []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minCostClimbingStairsDP(tt.cost)
			if result != tt.expected {
				t.Errorf("minCostClimbingStairsDP(%v) = %d, expected %d", tt.cost, result, tt.expected)
			}
		})
	}
}

func TestMinCostClimbingStairsOptimized(t *testing.T) {
	tests := []struct {
		name     string
		cost     []int
		expected int
	}{
		{
			name:     "Example 1",
			cost:     []int{10, 15, 20},
			expected: 15,
		},
		{
			name:     "Example 2",
			cost:     []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minCostClimbingStairsOptimized(tt.cost)
			if result != tt.expected {
				t.Errorf("minCostClimbingStairsOptimized(%v) = %d, expected %d", tt.cost, result, tt.expected)
			}
		})
	}
}

func BenchmarkMinCostClimbingStairs(b *testing.B) {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	for i := 0; i < b.N; i++ {
		minCostClimbingStairs(cost)
	}
}

func BenchmarkMinCostClimbingStairsDP(b *testing.B) {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	for i := 0; i < b.N; i++ {
		minCostClimbingStairsDP(cost)
	}
}

func BenchmarkMinCostClimbingStairsOptimized(b *testing.B) {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	for i := 0; i < b.N; i++ {
		minCostClimbingStairsOptimized(cost)
	}
}