package dp

import (
	"testing"
)

func TestRob(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3, 1},
			expected: 4,
		},
		{
			name:     "Example 2",
			nums:     []int{2, 7, 9, 3, 1},
			expected: 12,
		},
		{
			name:     "Single house",
			nums:     []int{5},
			expected: 5,
		},
		{
			name:     "Two houses - rob first",
			nums:     []int{2, 1},
			expected: 2,
		},
		{
			name:     "Two houses - rob second",
			nums:     []int{1, 2},
			expected: 2,
		},
		{
			name:     "Three houses",
			nums:     []int{2, 1, 1, 2},
			expected: 4,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "Large values",
			nums:     []int{100, 50, 400, 200, 100},
			expected: 600, // Rob houses 0, 2, 4: 100 + 400 + 100 = 600
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Rob(tt.nums)
			if result != tt.expected {
				t.Errorf("Rob(%v) = %d, expected %d", tt.nums, result, tt.expected)
			}

			// Also test DP array implementation
			resultDP := RobDPArray(tt.nums)
			if resultDP != tt.expected {
				t.Errorf("RobDPArray(%v) = %d, expected %d", tt.nums, resultDP, tt.expected)
			}

			// Also test recursive implementation
			resultRecursive := RobRecursive(tt.nums)
			if resultRecursive != tt.expected {
				t.Errorf("RobRecursive(%v) = %d, expected %d", tt.nums, resultRecursive, tt.expected)
			}
		})
	}
}

func TestRobEdgeCases(t *testing.T) {
	// Test with maximum constraints
	nums := make([]int, 100)
	for i := range nums {
		nums[i] = 400
	}
	// For alternating houses: sum of every other house
	expected := 50 * 400 // 50 houses * 400 each
	result := Rob(nums)
	if result != expected {
		t.Errorf("Rob(large array) = %d, expected %d", result, expected)
	}
}

func BenchmarkRob(b *testing.B) {
	nums := []int{2, 7, 9, 3, 1, 5, 8, 4, 6, 2, 7, 9, 3, 1, 5, 8, 4, 6}
	for i := 0; i < b.N; i++ {
		Rob(nums)
	}
}

func BenchmarkRobDPArray(b *testing.B) {
	nums := []int{2, 7, 9, 3, 1, 5, 8, 4, 6, 2, 7, 9, 3, 1, 5, 8, 4, 6}
	for i := 0; i < b.N; i++ {
		RobDPArray(nums)
	}
}

func BenchmarkRobRecursive(b *testing.B) {
	nums := []int{2, 7, 9, 3, 1, 5, 8, 4, 6, 2, 7, 9, 3, 1, 5, 8, 4, 6}
	for i := 0; i < b.N; i++ {
		RobRecursive(nums)
	}
}