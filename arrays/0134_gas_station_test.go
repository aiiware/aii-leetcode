package arrays

import (
	"testing"
)

func TestCanCompleteCircuit(t *testing.T) {
	tests := []struct {
		name     string
		gas      []int
		cost     []int
		expected int
	}{
		{
			name:     "Example 1",
			gas:      []int{1, 2, 3, 4, 5},
			cost:     []int{3, 4, 5, 1, 2},
			expected: 3,
		},
		{
			name:     "Example 2",
			gas:      []int{2, 3, 4},
			cost:     []int{3, 4, 3},
			expected: -1,
		},
		{
			name:     "Single station possible",
			gas:      []int{5},
			cost:     []int{4},
			expected: 0,
		},
		{
			name:     "Single station impossible",
			gas:      []int{3},
			cost:     []int{4},
			expected: -1,
		},
		{
			name:     "All stations possible starting at 0",
			gas:      []int{5, 1, 2, 3, 4},
			cost:     []int{4, 4, 1, 5, 1},
			expected: 4,
		},
		{
			name:     "Multiple possible starts returns first",
			gas:      []int{3, 1, 1},
			cost:     []int{1, 2, 2},
			expected: 0,
		},
		{
			name:     "Large circuit possible",
			gas:      []int{5, 8, 2, 8},
			cost:     []int{6, 5, 6, 6},
			expected: 3,
		},
		{
			name:     "Just enough gas",
			gas:      []int{2, 3, 4},
			cost:     []int{3, 4, 2},
			expected: 2,
		},
		{
			name:     "All zero cost",
			gas:      []int{1, 2, 3, 4, 5},
			cost:     []int{0, 0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "Large values",
			gas:      []int{10000, 10000, 10000, 10000},
			cost:     []int{10000, 10000, 10000, 10000},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CanCompleteCircuit(tt.gas, tt.cost)
			if result != tt.expected {
				t.Errorf("CanCompleteCircuit(%v, %v) = %d, expected %d",
					tt.gas, tt.cost, result, tt.expected)
			}
		})
	}
}

func BenchmarkCanCompleteCircuit(b *testing.B) {
	// Create test cases of different sizes
	testCases := []struct {
		name string
		gas  []int
		cost []int
	}{
		{
			name: "Small circuit",
			gas:  []int{1, 2, 3, 4, 5},
			cost: []int{3, 4, 5, 1, 2},
		},
		{
			name: "Medium circuit",
			gas:  []int{5, 8, 2, 8, 3, 1, 7, 4, 2, 6},
			cost: []int{6, 5, 6, 6, 2, 3, 5, 4, 3, 5},
		},
		{
			name: "Large circuit",
			gas:  make([]int, 1000),
			cost: make([]int, 1000),
		},
		{
			name: "Very large circuit",
			gas:  make([]int, 10000),
			cost: make([]int, 10000),
		},
	}

	// Initialize large arrays
	for i := range testCases[2].gas {
		testCases[2].gas[i] = (i % 10) + 1
		testCases[2].cost[i] = ((i + 1) % 10) + 1
	}
	for i := range testCases[3].gas {
		testCases[3].gas[i] = (i % 100) + 1
		testCases[3].cost[i] = ((i + 1) % 100) + 1
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				CanCompleteCircuit(tc.gas, tc.cost)
			}
		})
	}
}
