package leetcode

import (
	"testing"
)

func TestCandy(t *testing.T) {
	tests := []struct {
		name     string
		ratings  []int
		expected int
	}{
		{
			name:     "Example 1",
			ratings:  []int{1, 0, 2},
			expected: 5,
		},
		{
			name:     "Example 2",
			ratings:  []int{1, 2, 2},
			expected: 4,
		},
		{
			name:     "Empty ratings",
			ratings:  []int{},
			expected: 0,
		},
		{
			name:     "Single child",
			ratings:  []int{5},
			expected: 1,
		},
		{
			name:     "All same ratings",
			ratings:  []int{3, 3, 3, 3},
			expected: 4,
		},
		{
			name:     "Increasing ratings",
			ratings:  []int{1, 2, 3, 4, 5},
			expected: 15, // 1 + 2 + 3 + 4 + 5 = 15
		},
		{
			name:     "Decreasing ratings",
			ratings:  []int{5, 4, 3, 2, 1},
			expected: 15, // 5 + 4 + 3 + 2 + 1 = 15
		},
		{
			name:     "Mountain shape",
			ratings:  []int{1, 2, 3, 2, 1},
			expected: 9, // 1 + 2 + 3 + 2 + 1 = 9
		},
		{
			name:     "Valley shape",
			ratings:  []int{3, 2, 1, 2, 3},
			expected: 11, // 3 + 2 + 1 + 2 + 3 = 11
		},
		{
			name:     "Complex case 1",
			ratings:  []int{1, 3, 2, 2, 1},
			expected: 7, // 1 + 2 + 1 + 2 + 1 = 7
		},
		{
			name:     "Complex case 2",
			ratings:  []int{1, 2, 87, 87, 87, 2, 1},
			expected: 13, // 1 + 2 + 3 + 1 + 3 + 2 + 1 = 13
		},
		{
			name:     "Large ratings",
			ratings:  []int{10000, 20000, 30000, 20000, 10000},
			expected: 9, // Same pattern as mountain shape
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Candy(tt.ratings)
			if result != tt.expected {
				t.Errorf("Candy(%v) = %d, expected %d", tt.ratings, result, tt.expected)
			}
		})
	}
}

func BenchmarkCandy(b *testing.B) {
	// Create test cases of different sizes
	testCases := []struct {
		name    string
		ratings []int
	}{
		{
			name:    "Small array",
			ratings: []int{1, 0, 2},
		},
		{
			name:    "Medium array",
			ratings: []int{1, 2, 3, 4, 5, 4, 3, 2, 1},
		},
		{
			name:    "Large increasing array",
			ratings: make([]int, 1000),
		},
		{
			name:    "Very large random array",
			ratings: make([]int, 10000),
		},
	}

	// Initialize large arrays
	for i := range testCases[2].ratings {
		testCases[2].ratings[i] = i + 1
	}
	for i := range testCases[3].ratings {
		// Create a pattern: increasing then decreasing
		if i < 5000 {
			testCases[3].ratings[i] = i + 1
		} else {
			testCases[3].ratings[i] = 10000 - i
		}
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Candy(tc.ratings)
			}
		})
	}
}