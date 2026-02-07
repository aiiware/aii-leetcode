package arrays

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			name:     "Example 2",
			nums:     []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		},
		{
			name:     "Single element",
			nums:     []int{5},
			expected: []int{1},
		},
		{
			name:     "Two elements",
			nums:     []int{2, 3},
			expected: []int{3, 2},
		},
		{
			name:     "With zeros",
			nums:     []int{0, 0},
			expected: []int{0, 0},
		},
		{
			name:     "All ones",
			nums:     []int{1, 1, 1, 1},
			expected: []int{1, 1, 1, 1},
		},
		{
			name:     "Negative numbers",
			nums:     []int{-2, -3, -4},
			expected: []int{12, 8, 6},
		},
		{
			name:     "Mixed positive and negative",
			nums:     []int{2, -1, 3, -2},
			expected: []int{6, -12, 4, -6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ProductExceptSelf(tt.nums)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ProductExceptSelf(%v) = %v, expected %v", tt.nums, result, tt.expected)
			}
		})
	}
}

func BenchmarkProductExceptSelf(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		ProductExceptSelf(nums)
	}
}