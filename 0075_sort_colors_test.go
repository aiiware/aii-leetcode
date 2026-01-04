package leetcode

import (
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Example 1 from LeetCode",
			nums:     []int{2, 0, 2, 1, 1, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "Example 2 from LeetCode",
			nums:     []int{2, 0, 1},
			expected: []int{0, 1, 2},
		},
		{
			name:     "Already sorted",
			nums:     []int{0, 0, 1, 1, 2, 2},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "Reverse sorted",
			nums:     []int{2, 2, 1, 1, 0, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0, 0, 0},
			expected: []int{0, 0, 0, 0, 0},
		},
		{
			name:     "All ones",
			nums:     []int{1, 1, 1, 1, 1},
			expected: []int{1, 1, 1, 1, 1},
		},
		{
			name:     "All twos",
			nums:     []int{2, 2, 2, 2, 2},
			expected: []int{2, 2, 2, 2, 2},
		},
		{
			name:     "Single element - 0",
			nums:     []int{0},
			expected: []int{0},
		},
		{
			name:     "Single element - 1",
			nums:     []int{1},
			expected: []int{1},
		},
		{
			name:     "Single element - 2",
			nums:     []int{2},
			expected: []int{2},
		},
		{
			name:     "Mixed with only 0s and 1s",
			nums:     []int{1, 0, 1, 0, 1, 0},
			expected: []int{0, 0, 0, 1, 1, 1},
		},
		{
			name:     "Mixed with only 0s and 2s",
			nums:     []int{2, 0, 2, 0, 2, 0},
			expected: []int{0, 0, 0, 2, 2, 2},
		},
		{
			name:     "Mixed with only 1s and 2s",
			nums:     []int{2, 1, 2, 1, 2, 1},
			expected: []int{1, 1, 1, 2, 2, 2},
		},
		{
			name:     "Complex pattern 1",
			nums:     []int{1, 2, 0, 1, 2, 0, 1, 2, 0},
			expected: []int{0, 0, 0, 1, 1, 1, 2, 2, 2},
		},
		{
			name:     "Complex pattern 2",
			nums:     []int{2, 1, 0, 2, 1, 0, 2, 1, 0},
			expected: []int{0, 0, 0, 1, 1, 1, 2, 2, 2},
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: []int{},
		},
		{
			name:     "Large array with random distribution",
			nums:     []int{2, 0, 1, 2, 1, 0, 2, 0, 1, 2, 0, 1, 0, 2, 1},
			expected: []int{0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of the input to avoid modifying the test case
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			
			SortColors(nums)
			
			// Check if sorted correctly
			if !equalSlices(nums, tt.expected) {
				t.Errorf("SortColors() = %v, expected %v", nums, tt.expected)
			}
		})
	}
}

func TestSortColorsCounting(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Example 1 from LeetCode",
			nums:     []int{2, 0, 2, 1, 1, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "Example 2 from LeetCode",
			nums:     []int{2, 0, 1},
			expected: []int{0, 1, 2},
		},
		{
			name:     "Already sorted",
			nums:     []int{0, 0, 1, 1, 2, 2},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0, 0, 0},
			expected: []int{0, 0, 0, 0, 0},
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: []int{},
		},
		{
			name:     "Large array",
			nums:     []int{2, 0, 1, 2, 1, 0, 2, 0, 1, 2, 0, 1, 0, 2, 1},
			expected: []int{0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of the input to avoid modifying the test case
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			
			SortColorsCounting(nums)
			
			// Check if sorted correctly
			if !equalSlices(nums, tt.expected) {
				t.Errorf("SortColorsCounting() = %v, expected %v", nums, tt.expected)
			}
		})
	}
}

func TestSortColorsTwoPass(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Example 1 from LeetCode",
			nums:     []int{2, 0, 2, 1, 1, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "Example 2 from LeetCode",
			nums:     []int{2, 0, 1},
			expected: []int{0, 1, 2},
		},
		{
			name:     "Already sorted",
			nums:     []int{0, 0, 1, 1, 2, 2},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0, 0, 0},
			expected: []int{0, 0, 0, 0, 0},
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: []int{},
		},
		{
			name:     "Large array",
			nums:     []int{2, 0, 1, 2, 1, 0, 2, 0, 1, 2, 0, 1, 0, 2, 1},
			expected: []int{0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of the input to avoid modifying the test case
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)
			
			SortColorsTwoPass(nums)
			
			// Check if sorted correctly
			if !equalSlices(nums, tt.expected) {
				t.Errorf("SortColorsTwoPass() = %v, expected %v", nums, tt.expected)
			}
		})
	}
}

func BenchmarkSortColors(b *testing.B) {
	// Create test arrays of different sizes
	testCases := []struct {
		name string
		nums []int
	}{
		{
			name: "Small array (10 elements)",
			nums: []int{2, 0, 1, 2, 1, 0, 2, 0, 1, 2},
		},
		{
			name: "Medium array (100 elements)",
			nums: func() []int {
				arr := make([]int, 100)
				for i := range arr {
					arr[i] = i % 3 // 0, 1, 2 repeating
				}
				return arr
			}(),
		},
		{
			name: "Large array (1000 elements)",
			nums: func() []int {
				arr := make([]int, 1000)
				for i := range arr {
					arr[i] = (i * 7) % 3 // More varied distribution
				}
				return arr
			}(),
		},
		{
			name: "Very large array (10000 elements)",
			nums: func() []int {
				arr := make([]int, 10000)
				for i := range arr {
					arr[i] = (i * 13) % 3
				}
				return arr
			}(),
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			// Create a fresh copy for each iteration
			nums := make([]int, len(tc.nums))
			b.ResetTimer()
			
			for i := 0; i < b.N; i++ {
				copy(nums, tc.nums)
				SortColors(nums)
			}
		})
	}
}

func BenchmarkSortColorsCounting(b *testing.B) {
	// Create test arrays of different sizes
	testCases := []struct {
		name string
		nums []int
	}{
		{
			name: "Small array (10 elements)",
			nums: []int{2, 0, 1, 2, 1, 0, 2, 0, 1, 2},
		},
		{
			name: "Medium array (100 elements)",
			nums: func() []int {
				arr := make([]int, 100)
				for i := range arr {
					arr[i] = i % 3
				}
				return arr
			}(),
		},
		{
			name: "Large array (1000 elements)",
			nums: func() []int {
				arr := make([]int, 1000)
				for i := range arr {
					arr[i] = (i * 7) % 3
				}
				return arr
			}(),
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			// Create a fresh copy for each iteration
			nums := make([]int, len(tc.nums))
			b.ResetTimer()
			
			for i := 0; i < b.N; i++ {
				copy(nums, tc.nums)
				SortColorsCounting(nums)
			}
		})
	}
}

func BenchmarkSortColorsTwoPass(b *testing.B) {
	// Create test arrays of different sizes
	testCases := []struct {
		name string
		nums []int
	}{
		{
			name: "Small array (10 elements)",
			nums: []int{2, 0, 1, 2, 1, 0, 2, 0, 1, 2},
		},
		{
			name: "Medium array (100 elements)",
			nums: func() []int {
				arr := make([]int, 100)
				for i := range arr {
					arr[i] = i % 3
				}
				return arr
			}(),
		},
		{
			name: "Large array (1000 elements)",
			nums: func() []int {
				arr := make([]int, 1000)
				for i := range arr {
					arr[i] = (i * 7) % 3
				}
				return arr
			}(),
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			// Create a fresh copy for each iteration
			nums := make([]int, len(tc.nums))
			b.ResetTimer()
			
			for i := 0; i < b.N; i++ {
				copy(nums, tc.nums)
				SortColorsTwoPass(nums)
			}
		})
	}
}

// Helper function to compare slices
func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}