package arrays

import (
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "Example 1 from LeetCode",
			nums:     []int{1, 2, 3},
			expected: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			name:     "Example 2 from LeetCode",
			nums:     []int{0},
			expected: [][]int{{}, {0}},
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: [][]int{{}},
		},
		{
			name:     "Single element",
			nums:     []int{5},
			expected: [][]int{{}, {5}},
		},
		{
			name:     "Two elements",
			nums:     []int{1, 2},
			expected: [][]int{{}, {1}, {2}, {1, 2}},
		},
		{
			name: "Four elements",
			nums: []int{1, 2, 3, 4},
			expected: [][]int{
				{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3},
				{4}, {1, 4}, {2, 4}, {1, 2, 4}, {3, 4}, {1, 3, 4}, {2, 3, 4}, {1, 2, 3, 4},
			},
		},
		{
			name: "Negative numbers",
			nums: []int{-1, 0, 1},
			expected: [][]int{
				{}, {-1}, {0}, {-1, 0}, {1}, {-1, 1}, {0, 1}, {-1, 0, 1},
			},
		},
		{
			name: "Large range",
			nums: []int{10, 20, 30},
			expected: [][]int{
				{}, {10}, {20}, {10, 20}, {30}, {10, 30}, {20, 30}, {10, 20, 30},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Subsets(tt.nums)
			if !equalSubsets(result, tt.expected) {
				t.Errorf("Subsets(%v) = %v, expected %v", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestSubsetsIterative(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "Example 1 from LeetCode",
			nums:     []int{1, 2, 3},
			expected: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			name:     "Example 2 from LeetCode",
			nums:     []int{0},
			expected: [][]int{{}, {0}},
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: [][]int{{}},
		},
		{
			name:     "Two elements",
			nums:     []int{1, 2},
			expected: [][]int{{}, {1}, {2}, {1, 2}},
		},
		{
			name: "Negative numbers",
			nums: []int{-1, 0, 1},
			expected: [][]int{
				{}, {-1}, {0}, {-1, 0}, {1}, {-1, 1}, {0, 1}, {-1, 0, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SubsetsIterative(tt.nums)
			if !equalSubsets(result, tt.expected) {
				t.Errorf("SubsetsIterative(%v) = %v, expected %v", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestSubsetsBitMask(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "Example 1 from LeetCode",
			nums:     []int{1, 2, 3},
			expected: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			name:     "Example 2 from LeetCode",
			nums:     []int{0},
			expected: [][]int{{}, {0}},
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: [][]int{{}},
		},
		{
			name:     "Two elements",
			nums:     []int{1, 2},
			expected: [][]int{{}, {1}, {2}, {1, 2}},
		},
		{
			name: "Negative numbers",
			nums: []int{-1, 0, 1},
			expected: [][]int{
				{}, {-1}, {0}, {-1, 0}, {1}, {-1, 1}, {0, 1}, {-1, 0, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SubsetsBitMask(tt.nums)
			if !equalSubsets(result, tt.expected) {
				t.Errorf("SubsetsBitMask(%v) = %v, expected %v", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestSubsetsDFS(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "Example 1 from LeetCode",
			nums:     []int{1, 2, 3},
			expected: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			name:     "Example 2 from LeetCode",
			nums:     []int{0},
			expected: [][]int{{}, {0}},
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: [][]int{{}},
		},
		{
			name:     "Two elements",
			nums:     []int{1, 2},
			expected: [][]int{{}, {1}, {2}, {1, 2}},
		},
		{
			name: "Negative numbers",
			nums: []int{-1, 0, 1},
			expected: [][]int{
				{}, {-1}, {0}, {-1, 0}, {1}, {-1, 1}, {0, 1}, {-1, 0, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SubsetsDFS(tt.nums)
			if !equalSubsets(result, tt.expected) {
				t.Errorf("SubsetsDFS(%v) = %v, expected %v", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestSubsetsLexicographic(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "Example 1 from LeetCode",
			nums:     []int{1, 2, 3},
			expected: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			name:     "Example 2 from LeetCode",
			nums:     []int{0},
			expected: [][]int{{}, {0}},
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: [][]int{{}},
		},
		{
			name:     "Two elements",
			nums:     []int{1, 2},
			expected: [][]int{{}, {1}, {2}, {1, 2}},
		},
		{
			name: "Negative numbers",
			nums: []int{-1, 0, 1},
			expected: [][]int{
				{}, {-1}, {0}, {-1, 0}, {1}, {-1, 1}, {0, 1}, {-1, 0, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SubsetsLexicographic(tt.nums)
			if !equalSubsets(result, tt.expected) {
				t.Errorf("SubsetsLexicographic(%v) = %v, expected %v", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestSubsetsWithDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "Unique elements - Example 1",
			nums:     []int{1, 2, 3},
			expected: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
		{
			name:     "Unique elements - Single",
			nums:     []int{0},
			expected: [][]int{{}, {0}},
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: [][]int{{}},
		},
		{
			name:     "Two unique elements",
			nums:     []int{1, 2},
			expected: [][]int{{}, {1}, {2}, {1, 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SubsetsWithDuplicates(tt.nums)
			if !equalSubsets(result, tt.expected) {
				t.Errorf("SubsetsWithDuplicates(%v) = %v, expected %v", tt.nums, result, tt.expected)
			}
		})
	}
}

func BenchmarkSubsets(b *testing.B) {
	testCases := []struct {
		name string
		nums []int
	}{
		{
			name: "Small (3 elements)",
			nums: []int{1, 2, 3},
		},
		{
			name: "Medium (6 elements)",
			nums: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "Large (10 elements)",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name: "Maximum (10 elements with negatives)",
			nums: []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4},
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Subsets(tc.nums)
			}
		})
	}
}

func BenchmarkSubsetsIterative(b *testing.B) {
	testCases := []struct {
		name string
		nums []int
	}{
		{
			name: "Small (3 elements)",
			nums: []int{1, 2, 3},
		},
		{
			name: "Medium (6 elements)",
			nums: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "Large (10 elements)",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SubsetsIterative(tc.nums)
			}
		})
	}
}

func BenchmarkSubsetsBitMask(b *testing.B) {
	testCases := []struct {
		name string
		nums []int
	}{
		{
			name: "Small (3 elements)",
			nums: []int{1, 2, 3},
		},
		{
			name: "Medium (6 elements)",
			nums: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "Large (10 elements)",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SubsetsBitMask(tc.nums)
			}
		})
	}
}

func BenchmarkSubsetsDFS(b *testing.B) {
	testCases := []struct {
		name string
		nums []int
	}{
		{
			name: "Small (3 elements)",
			nums: []int{1, 2, 3},
		},
		{
			name: "Medium (6 elements)",
			nums: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "Large (10 elements)",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SubsetsDFS(tc.nums)
			}
		})
	}
}

func BenchmarkSubsetsLexicographic(b *testing.B) {
	testCases := []struct {
		name string
		nums []int
	}{
		{
			name: "Small (3 elements)",
			nums: []int{1, 2, 3},
		},
		{
			name: "Medium (6 elements)",
			nums: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name: "Large (10 elements)",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SubsetsLexicographic(tc.nums)
			}
		})
	}
}

// Helper function to compare subsets (order of subsets doesn't matter)
func equalSubsets(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	// Create maps to count occurrences of each subset
	mapA := make(map[string]int)
	mapB := make(map[string]int)

	for _, subset := range a {
		// Sort and create key (simplified - we'll just create a string representation)
		key := ""
		// For empty subset
		if len(subset) == 0 {
			key = "[]"
		} else {
			// Create a simple string representation
			for _, num := range subset {
				key += string(rune(num)) + ","
			}
		}
		mapA[key]++
	}

	for _, subset := range b {
		key := ""
		if len(subset) == 0 {
			key = "[]"
		} else {
			for _, num := range subset {
				key += string(rune(num)) + ","
			}
		}
		mapB[key]++
	}

	return reflect.DeepEqual(mapA, mapB)
}
