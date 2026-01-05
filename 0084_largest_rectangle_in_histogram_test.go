package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		name     string
		heights  []int
		expected int
	}{
		{
			name:     "Example 1 from LeetCode",
			heights:  []int{2, 1, 5, 6, 2, 3},
			expected: 10,
		},
		{
			name:     "Example 2 from LeetCode",
			heights:  []int{2, 4},
			expected: 4,
		},
		{
			name:     "Single bar",
			heights:  []int{5},
			expected: 5,
		},
		{
			name:     "Empty array",
			heights:  []int{},
			expected: 0,
		},
		{
			name:     "All increasing",
			heights:  []int{1, 2, 3, 4, 5},
			expected: 9, // 3 * 3 = 9 (height 3 can extend from index 2 to 4)
		},
		{
			name:     "All decreasing",
			heights:  []int{5, 4, 3, 2, 1},
			expected: 9, // 3 * 3 = 9 (height 3 can extend from index 0 to 2)
		},
		{
			name:     "All same height",
			heights:  []int{3, 3, 3, 3, 3},
			expected: 15, // 3 * 5 = 15
		},
		{
			name:     "Mountain shape",
			heights:  []int{1, 3, 5, 3, 1},
			expected: 9, // 3 * 3 = 9 (height 3 can extend from index 1 to 3)
		},
		{
			name:     "Valley shape",
			heights:  []int{5, 3, 1, 3, 5},
			expected: 6, // 3 * 2 = 6 (height 3 at index 3 can extend to index 4)
		},
		{
			name:     "Zero heights in middle",
			heights:  []int{2, 0, 2},
			expected: 2, // Two bars of height 2
		},
		{
			name:     "Zero at beginning",
			heights:  []int{0, 2, 3, 4},
			expected: 6, // 3 * 2 = 6 (height 3 can extend from index 2 to 3)
		},
		{
			name:     "Zero at end",
			heights:  []int{4, 3, 2, 0},
			expected: 6, // 3 * 2 = 6 (height 3 can extend from index 0 to 1)
		},
		{
			name:     "Complex case 1",
			heights:  []int{6, 2, 5, 4, 5, 1, 6},
			expected: 12, // 4 * 3 = 12 (height 4 can extend from index 2 to 4)
		},
		{
			name:     "Complex case 2",
			heights:  []int{3, 5, 1, 7, 5, 9},
			expected: 15, // 5 * 3 = 15 (height 5 can extend from index 1 to 3)
		},
		{
			name:     "Large numbers",
			heights:  []int{10000, 10000, 10000},
			expected: 30000,
		},
		{
			name:     "Alternating heights",
			heights:  []int{1, 2, 1, 2, 1, 2, 1},
			expected: 7, // 1 * 7 = 7 (height 1 can extend entire array)
		},
		{
			name:     "Sawtooth pattern",
			heights:  []int{1, 3, 2, 4, 3, 5, 4},
			expected: 12, // 2 * 6 = 12 (height 2 can extend from index 1 to 6)
		},
		{
			name:     "Two peaks",
			heights:  []int{1, 2, 3, 2, 1, 2, 3, 2, 1},
			expected: 9, // 3 * 3 = 9 (height 3 can extend from index 2 to 4 or 6 to 8)
		},
		{
			name:     "Plateau in middle",
			heights:  []int{1, 2, 2, 2, 1},
			expected: 6, // 2 * 3 = 6 (height 2 plateau)
		},
		{
			name:     "Random test 1",
			heights:  []int{4, 2, 0, 3, 2, 5},
			expected: 6, // 3 * 2 = 6 or 2 * 3 = 6
		},
		{
			name:     "Random test 2",
			heights:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: 30, // 6 * 5 = 30 (height 6 can extend from index 5 to 9)
		},
		{
			name:     "Descending then ascending",
			heights:  []int{5, 4, 3, 2, 1, 2, 3, 4, 5},
			expected: 9, // 3 * 3 = 9 (multiple possibilities)
		},
		{
			name:     "Single tall bar in middle",
			heights:  []int{1, 1, 8, 1, 1},
			expected: 8, // 8 * 1 = 8
		},
		{
			name:     "Wide low bar",
			heights:  []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			expected: 10, // 1 * 10 = 10
		},
		{
			name:     "Edge case: all zeros",
			heights:  []int{0, 0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "Edge case: one zero",
			heights:  []int{0},
			expected: 0,
		},
		{
			name:     "Large array with pattern",
			heights:  []int{2, 1, 2, 1, 2, 1, 2, 1, 2, 1},
			expected: 10, // 2 * 5 = 10 (height 2 can extend every other position)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LargestRectangleArea(tt.heights)
			assert.Equal(t, tt.expected, result,
				"LargestRectangleArea(%v) = %d, expected %d",
				tt.heights, result, tt.expected)
		})
	}
}

func TestLargestRectangleAreaOptimized(t *testing.T) {
	tests := []struct {
		name     string
		heights  []int
		expected int
	}{
		{
			name:     "Example 1",
			heights:  []int{2, 1, 5, 6, 2, 3},
			expected: 10,
		},
		{
			name:     "Example 2",
			heights:  []int{2, 4},
			expected: 4,
		},
		{
			name:     "All same",
			heights:  []int{3, 3, 3},
			expected: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LargestRectangleAreaOptimized(tt.heights)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestLargestRectangleAreaBruteForce(t *testing.T) {
	tests := []struct {
		name     string
		heights  []int
		expected int
	}{
		{
			name:     "Small test",
			heights:  []int{2, 1, 5},
			expected: 5,
		},
		{
			name:     "Single element",
			heights:  []int{5},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LargestRectangleAreaBruteForce(tt.heights)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestLargestRectangleAreaDivideConquer(t *testing.T) {
	tests := []struct {
		name     string
		heights  []int
		expected int
	}{
		{
			name:     "Example 1",
			heights:  []int{2, 1, 5, 6, 2, 3},
			expected: 10,
		},
		{
			name:     "Small array",
			heights:  []int{2, 4},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LargestRectangleAreaDivideConquer(tt.heights)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestLargestRectangleAreaDP(t *testing.T) {
	tests := []struct {
		name     string
		heights  []int
		expected int
	}{
		{
			name:     "Example 1",
			heights:  []int{2, 1, 5, 6, 2, 3},
			expected: 10,
		},
		{
			name:     "All increasing",
			heights:  []int{1, 2, 3, 4, 5},
			expected: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LargestRectangleAreaDP(tt.heights)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestLargestRectangleAreaSegmentTree(t *testing.T) {
	tests := []struct {
		name     string
		heights  []int
		expected int
	}{
		{
			name:     "Example 1",
			heights:  []int{2, 1, 5, 6, 2, 3},
			expected: 10,
		},
		{
			name:     "Simple case",
			heights:  []int{1, 2, 3},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LargestRectangleAreaSegmentTree(tt.heights)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestLargestRectangleArea_Consistency(t *testing.T) {
	testCases := []struct {
		name    string
		heights []int
	}{
		{
			name:    "Standard example",
			heights: []int{2, 1, 5, 6, 2, 3},
		},
		{
			name:    "All same",
			heights: []int{3, 3, 3, 3},
		},
		{
			name:    "Empty array",
			heights: []int{},
		},
		{
			name:    "Single element",
			heights: []int{5},
		},
		{
			name:    "Increasing",
			heights: []int{1, 2, 3, 4, 5},
		},
		{
			name:    "Decreasing",
			heights: []int{5, 4, 3, 2, 1},
		},
		{
			name:    "With zeros",
			heights: []int{0, 2, 0, 3, 0},
		},
		{
			name:    "Complex pattern",
			heights: []int{6, 2, 5, 4, 5, 1, 6},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test all implementations
			implementations := []struct {
				name string
				fn   func([]int) int
			}{
				{"LargestRectangleArea", LargestRectangleArea},
				{"LargestRectangleAreaOptimized", LargestRectangleAreaOptimized},
				{"LargestRectangleAreaBruteForce", LargestRectangleAreaBruteForce},
				{"LargestRectangleAreaDivideConquer", LargestRectangleAreaDivideConquer},
				{"LargestRectangleAreaDP", LargestRectangleAreaDP},
				{"LargestRectangleAreaSegmentTree", LargestRectangleAreaSegmentTree},
			}

			results := make([]int, len(implementations))
			
			for i, impl := range implementations {
				// For brute force, skip large arrays (it's O(n^2))
				if impl.name == "LargestRectangleAreaBruteForce" && len(tc.heights) > 100 {
					continue
				}
				results[i] = impl.fn(tc.heights)
			}

			// All implementations should return the same result
			// (except brute force which we might skip for large arrays)
			firstResult := results[0]
			for i := 1; i < len(results); i++ {
				if results[i] != 0 || len(tc.heights) <= 100 { // Skip if result is 0 and not from brute force
					assert.Equal(t, firstResult, results[i],
						"%s and %s should return same result for heights %v",
						implementations[0].name, implementations[i].name, tc.heights)
				}
			}
		})
	}
}

func TestLargestRectangleArea_PropertyBased(t *testing.T) {
	t.Run("Result is non-negative", func(t *testing.T) {
		testCases := [][]int{
			{2, 1, 5, 6, 2, 3},
			{0, 0, 0},
			{1, 2, 3},
			{},
			{5},
		}

		for _, heights := range testCases {
			result := LargestRectangleArea(heights)
			assert.True(t, result >= 0,
				"Result %d should be non-negative for heights %v",
				result, heights)
		}
	})

	t.Run("Result is at least max height", func(t *testing.T) {
		testCases := [][]int{
			{2, 1, 5, 6, 2, 3},
			{1, 2, 3},
			{5, 1, 1},
		}

		for _, heights := range testCases {
			if len(heights) == 0 {
				continue
			}
			
			result := LargestRectangleArea(heights)
			maxHeight := heights[0]
			for _, h := range heights {
				if h > maxHeight {
					maxHeight = h
				}
			}
			
			assert.True(t, result >= maxHeight,
				"Result %d should be at least max height %d for heights %v",
				result, maxHeight, heights)
		}
	})

	t.Run("Result is at most total area", func(t *testing.T) {
		testCases := [][]int{
			{2, 1, 5, 6, 2, 3},
			{1, 2, 3},
			{5, 1, 1},
		}

		for _, heights := range testCases {
			result := LargestRectangleArea(heights)
			totalArea := 0
			for _, h := range heights {
				totalArea += h
			}
			
			assert.True(t, result <= totalArea,
				"Result %d should be at most total area %d for heights %v",
				result, totalArea, heights)
		}
	})

	t.Run("Monotonic property for sorted arrays", func(t *testing.T) {
		// For sorted arrays (increasing or decreasing), the result should be
		// at least n * min(height) and at most n * max(height)
		testCases := [][]int{
			{1, 2, 3, 4, 5},    // increasing
			{5, 4, 3, 2, 1},    // decreasing
			{3, 3, 3, 3, 3},    // constant
		}

		for _, heights := range testCases {
			n := len(heights)
			if n == 0 {
				continue
			}
			
			result := LargestRectangleArea(heights)
			
			// Find min and max
			minHeight := heights[0]
			maxHeight := heights[0]
			for _, h := range heights {
				if h < minHeight {
					minHeight = h
				}
				if h > maxHeight {
					maxHeight = h
				}
			}
			
			assert.True(t, result >= n*minHeight,
				"For sorted array %v, result %d should be >= %d",
				heights, result, n*minHeight)
			
			assert.True(t, result <= n*maxHeight,
				"For sorted array %v, result %d should be <= %d",
				heights, result, n*maxHeight)
		}
	})

	t.Run("Adding zeros doesn't increase max area", func(t *testing.T) {
		original := []int{2, 1, 5, 6, 2, 3}
		withZeros := []int{0, 2, 1, 5, 6, 2, 3, 0}
		
		originalResult := LargestRectangleArea(original)
		withZerosResult := LargestRectangleArea(withZeros)
		
		// Adding zeros at boundaries shouldn't increase the maximum area
		// (it might actually be the same or less)
		assert.True(t, withZerosResult <= originalResult,
			"Adding zeros shouldn't increase max area. Original: %d, With zeros: %d",
			originalResult, withZerosResult)
	})

	t.Run("Scaling property", func(t *testing.T) {
		// If we scale all heights by factor k, the area should scale by k
		heights := []int{2, 1, 5, 6, 2, 3}
		scaleFactor := 2
		
		scaledHeights := make([]int, len(heights))
		for i, h := range heights {
			scaledHeights[i] = h * scaleFactor
		}
		
		originalResult := LargestRectangleArea(heights)
		scaledResult := LargestRectangleArea(scaledHeights)
		
		// Allow for small floating point errors if we were using floats
		// Since we're using integers, it should be exact
		assert.Equal(t, originalResult*scaleFactor, scaledResult,
			"Scaling heights by %d should scale area by %d. Original: %d, Scaled: %d",
			scaleFactor, scaleFactor, originalResult, scaledResult)
	})
}

func BenchmarkLargestRectangleArea(b *testing.B) {
	// Create test cases of different sizes and patterns
	testCases := []struct {
		name    string
		heights []int
	}{
		{
			name: "Small (10 elements)",
			heights: []int{2, 1, 5, 6, 2, 3, 4, 1, 2, 3},
		},
		{
			name: "Medium (100 elements) random",
			heights: func() []int {
				nums := make([]int, 100)
				for i := range nums {
					nums[i] = i % 10 // Creates pattern
				}
				return nums
			}(),
		},
		{
			name: "Medium (100 elements) increasing",
			heights: func() []int {
				nums := make([]int, 100)
				for i := range nums {
					nums[i] = i + 1
				}
				return nums
			}(),
		},
		{
			name: "Medium (100 elements) decreasing",
			heights: func() []int {
				nums := make([]int, 100)
				for i := range nums {
					nums[i] = 100 - i
				}
				return nums
			}(),
		},
		{
			name: "Large (1000 elements)",
			heights: func() []int {
				nums := make([]int, 1000)
				for i := range nums {
					nums[i] = (i * i) % 100 // Quadratic pattern
				}
				return nums
			}(),
		},
		{
			name: "Very large (10000 elements)",
			heights: func() []int {
				nums := make([]int, 10000)
				for i := range nums {
					nums[i] = i % 50
				}
				return nums
			}(),
		},
		{
			name: "All same (1000 elements)",
			heights: func() []int {
				nums := make([]int, 1000)
				for i := range nums {
					nums[i] = 7
				}
				return nums
			}(),
		},
		{
			name: "Sawtooth pattern (1000 elements)",
			heights: func() []int {
				nums := make([]int, 1000)
				for i := range nums {
					nums[i] = i % 10
				}
				return nums
			}(),
		},
		{
			name: "Mountain pattern (1000 elements)",
			heights: func() []int {
				nums := make([]int, 1000)
				for i := range nums {
					if i < 500 {
						nums[i] = i
					} else {
						nums[i] = 1000 - i
					}
				}
				return nums
			}(),
		},
		{
			name: "Valley pattern (1000 elements)",
			heights: func() []int {
				nums := make([]int, 1000)
				for i := range nums {
					if i < 500 {
						nums[i] = 500 - i
					} else {
						nums[i] = i - 500
					}
				}
				return nums
			}(),
		},
	}

	implementations := []struct {
		name string
		fn   func([]int) int
	}{
		{"Standard", LargestRectangleArea},
		{"Optimized", LargestRectangleAreaOptimized},
		{"DivideConquer", LargestRectangleAreaDivideConquer},
		{"DP", LargestRectangleAreaDP},
		{"SegmentTree", LargestRectangleAreaSegmentTree},
		// Note: BruteForce is excluded from benchmarks as it's O(n^2)
	}

	for _, tc := range testCases {
		for _, impl := range implementations {
			// Skip certain implementations for very large arrays
			if tc.name == "Very large (10000 elements)" {
				if impl.name == "DivideConquer" || impl.name == "SegmentTree" {
					continue // These are O(n log n) or worse
				}
			}
			
			b.Run(tc.name+"_"+impl.name, func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					impl.fn(tc.heights)
				}
			})
		}
	}
}