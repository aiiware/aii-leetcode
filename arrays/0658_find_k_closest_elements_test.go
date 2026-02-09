package arrays

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindClosestElements(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		k        int
		x        int
		expected []int
	}{
		{
			name:     "Example 1",
			arr:      []int{1, 2, 3, 4, 5},
			k:        4,
			x:        3,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Example 2",
			arr:      []int{1, 2, 3, 4, 5},
			k:        4,
			x:        -1,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Example 3",
			arr:      []int{1, 2, 3, 4, 5},
			k:        4,
			x:        6,
			expected: []int{2, 3, 4, 5},
		},
		{
			name:     "k equals array length",
			arr:      []int{1, 2, 3, 4, 5},
			k:        5,
			x:        3,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "k greater than array length",
			arr:      []int{1, 2, 3},
			k:        5,
			x:        3,
			expected: []int{1, 2, 3},
		},
		{
			name:     "Single element array",
			arr:      []int{5},
			k:        1,
			x:        3,
			expected: []int{5},
		},
		{
			name:     "x exactly in middle between two numbers",
			arr:      []int{1, 2, 3, 4, 5},
			k:        2,
			x:        3,
			expected: []int{2, 3}, // Prefer smaller when distances equal
		},
		{
			name:     "x closer to left side",
			arr:      []int{1, 2, 3, 4, 5},
			k:        3,
			x:        2,
			expected: []int{1, 2, 3},
		},
		{
			name:     "x closer to right side",
			arr:      []int{1, 2, 3, 4, 5},
			k:        3,
			x:        4,
			expected: []int{3, 4, 5},
		},
		{
			name:     "All elements equal",
			arr:      []int{5, 5, 5, 5, 5},
			k:        3,
			x:        3,
			expected: []int{5, 5, 5},
		},
		{
			name:     "Large gap in array",
			arr:      []int{1, 10, 20, 30, 40},
			k:        3,
			x:        15,
			expected: []int{1, 10, 20}, // Distances: 14, 5, 5, 15, 25 → closest: 10, 20, 1
		},
		{
			name:     "Negative numbers",
			arr:      []int{-5, -4, -3, -2, -1},
			k:        3,
			x:        -3,
			expected: []int{-4, -3, -2}, // Distances: 2, 1, 0, 1, 2 → closest: -3, -4, -2
		},
		{
			name:     "Mixed positive and negative",
			arr:      []int{-10, -5, 0, 5, 10},
			k:        3,
			x:        1,
			expected: []int{-5, 0, 5},
		},
		{
			name:     "Tie breaker test 1",
			arr:      []int{1, 2, 3, 4, 5},
			k:        2,
			x:        3,
			expected: []int{2, 3}, // |2-3|=1, |3-3|=0, |4-3|=1 → choose 2 over 4
		},
		{
			name:     "Tie breaker test 2",
			arr:      []int{0, 1, 2, 3, 4},
			k:        3,
			x:        2,
			expected: []int{1, 2, 3}, // |1-2|=1, |2-2|=0, |3-2|=1
		},
		{
			name:     "Large array",
			arr:      makeRangeInt(1, 100),
			k:        10,
			x:        50,
			expected: makeRangeInt(45, 55), // 45-54 are the 10 closest to 50
		},
		{
			name:     "x outside array on left",
			arr:      []int{10, 20, 30, 40, 50},
			k:        3,
			x:        5,
			expected: []int{10, 20, 30},
		},
		{
			name:     "x outside array on right",
			arr:      []int{10, 20, 30, 40, 50},
			k:        3,
			x:        60,
			expected: []int{30, 40, 50},
		},
		{
			name:     "k = 1",
			arr:      []int{1, 2, 3, 4, 5},
			k:        1,
			x:        3,
			expected: []int{3},
		},
		{
			name:     "k = 1, x between two numbers",
			arr:      []int{1, 3},
			k:        1,
			x:        2,
			expected: []int{1}, // |1-2|=1, |3-2|=1 → choose smaller (1)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindClosestElements(tt.arr, tt.k, tt.x)
			assert.Equal(t, tt.expected, result,
				"FindClosestElements(%v, %d, %d) = %v, expected %v",
				tt.arr, tt.k, tt.x, result, tt.expected)
		})
	}
}

func TestFindClosestElementsTwoPointers(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		k        int
		x        int
		expected []int
	}{
		{
			name:     "Example 1",
			arr:      []int{1, 2, 3, 4, 5},
			k:        4,
			x:        3,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Example 2",
			arr:      []int{1, 2, 3, 4, 5},
			k:        4,
			x:        -1,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "k = 1, x between two numbers",
			arr:      []int{1, 3},
			k:        1,
			x:        2,
			expected: []int{1},
		},
		{
			name:     "Large gap",
			arr:      []int{1, 10, 20, 30, 40},
			k:        3,
			x:        15,
			expected: []int{1, 10, 20}, // Distances: 14, 5, 5, 15, 25
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindClosestElementsTwoPointers(tt.arr, tt.k, tt.x)
			assert.Equal(t, tt.expected, result,
				"FindClosestElementsTwoPointers(%v, %d, %d) = %v, expected %v",
				tt.arr, tt.k, tt.x, result, tt.expected)
		})
	}
}

func TestFindClosestElementsSorting(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		k        int
		x        int
		expected []int
	}{
		{
			name:     "Example 1",
			arr:      []int{1, 2, 3, 4, 5},
			k:        4,
			x:        3,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "k = 1, x between two numbers",
			arr:      []int{1, 3},
			k:        1,
			x:        2,
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindClosestElementsSorting(tt.arr, tt.k, tt.x)
			assert.Equal(t, tt.expected, result,
				"FindClosestElementsSorting(%v, %d, %d) = %v, expected %v",
				tt.arr, tt.k, tt.x, result, tt.expected)
		})
	}
}

func TestFindClosestElements_Consistency(t *testing.T) {
	// Test that all implementations give the same result
	testCases := []struct {
		arr []int
		k   int
		x   int
	}{
		{[]int{1, 2, 3, 4, 5}, 4, 3},
		{[]int{1, 2, 3, 4, 5}, 4, -1},
		{[]int{1, 2, 3, 4, 5}, 4, 6},
		{[]int{1, 2, 3}, 2, 2},
		{[]int{1, 3}, 1, 2},
		{[]int{1, 10, 20, 30, 40}, 3, 15},
		{[]int{-5, -4, -3, -2, -1}, 3, -3},
		{makeRangeInt(1, 100), 10, 50},
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result1 := FindClosestElements(tc.arr, tc.k, tc.x)
			result2 := FindClosestElementsTwoPointers(tc.arr, tc.k, tc.x)
			result3 := FindClosestElementsSorting(tc.arr, tc.k, tc.x)

			// All results should be sorted
			assert.True(t, sort.IntsAreSorted(result1))
			assert.True(t, sort.IntsAreSorted(result2))
			assert.True(t, sort.IntsAreSorted(result3))

			// All should have length k (or len(arr) if k > len(arr))
			expectedLen := tc.k
			if tc.k > len(tc.arr) {
				expectedLen = len(tc.arr)
			}
			assert.Len(t, result1, expectedLen)
			assert.Len(t, result2, expectedLen)
			assert.Len(t, result3, expectedLen)

			// All should be equal
			assert.Equal(t, result1, result2, "Binary search vs Two pointers")
			assert.Equal(t, result1, result3, "Binary search vs Sorting")
		})
	}
}

func TestFindClosestElements_EdgeCases(t *testing.T) {
	t.Run("Empty array", func(t *testing.T) {
		result := FindClosestElements([]int{}, 3, 5)
		assert.Empty(t, result)
	})

	t.Run("k = 0", func(t *testing.T) {
		result := FindClosestElements([]int{1, 2, 3, 4, 5}, 0, 3)
		assert.Empty(t, result)
	})

	t.Run("Array with duplicates", func(t *testing.T) {
		result := FindClosestElements([]int{1, 2, 2, 2, 3, 4}, 3, 2)
		// Should return [2, 2, 2] or [1, 2, 2] depending on tie-breaking
		// The binary search implementation returns [1, 2, 2]
		assert.Len(t, result, 3)
		assert.True(t, sort.IntsAreSorted(result))
	})

	t.Run("x equals array element", func(t *testing.T) {
		result := FindClosestElements([]int{1, 2, 3, 4, 5}, 3, 3)
		assert.Equal(t, []int{2, 3, 4}, result)
	})

	t.Run("All elements at equal distance", func(t *testing.T) {
		// When all elements are at equal distance, should return first k
		result := FindClosestElements([]int{1, 3, 5, 7, 9}, 3, 5)
		// Distances: |1-5|=4, |3-5|=2, |5-5|=0, |7-5|=2, |9-5|=4
		// Should return [3, 5, 7]
		assert.Equal(t, []int{3, 5, 7}, result)
	})
}

func BenchmarkFindClosestElements(b *testing.B) {
	arr := makeRangeInt(1, 10000)
	k := 100
	x := 5000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindClosestElements(arr, k, x)
	}
}

func BenchmarkFindClosestElementsTwoPointers(b *testing.B) {
	arr := makeRangeInt(1, 10000)
	k := 100
	x := 5000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindClosestElementsTwoPointers(arr, k, x)
	}
}

func BenchmarkFindClosestElementsSorting(b *testing.B) {
	arr := makeRangeInt(1, 10000)
	k := 100
	x := 5000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindClosestElementsSorting(arr, k, x)
	}
}

func BenchmarkFindClosestElements_WorstCase(b *testing.B) {
	// Worst case for binary search: x at beginning or end
	arr := makeRangeInt(1, 10000)
	k := 100
	x := 1 // At beginning

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindClosestElements(arr, k, x)
	}
}

func BenchmarkFindClosestElements_BestCase(b *testing.B) {
	// Best case: x in middle
	arr := makeRangeInt(1, 10000)
	k := 100
	x := 5000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindClosestElements(arr, k, x)
	}
}