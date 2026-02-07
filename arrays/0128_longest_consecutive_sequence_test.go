package arrays

import (
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{100, 4, 200, 1, 3, 2},
			expected: 4,
		},
		{
			name:     "Example 2",
			nums:     []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			expected: 9,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "Single element",
			nums:     []int{5},
			expected: 1,
		},
		{
			name:     "No consecutive sequence",
			nums:     []int{1, 3, 5, 7, 9},
			expected: 1,
		},
		{
			name:     "All duplicates",
			nums:     []int{1, 1, 1, 1},
			expected: 1,
		},
		{
			name:     "Negative numbers",
			nums:     []int{-1, -2, -3, 0, 1, 2},
			expected: 6,
		},
		{
			name:     "Large gap in middle",
			nums:     []int{1, 2, 3, 100, 101, 102, 103},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LongestConsecutive(tt.nums)
			if result != tt.expected {
				t.Errorf("LongestConsecutive(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}
		})
	}
}

func BenchmarkLongestConsecutive(b *testing.B) {
	nums := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1, 100, 101, 102, 103, 200, 201, 202}
	for i := 0; i < b.N; i++ {
		LongestConsecutive(nums)
	}
}