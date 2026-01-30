package math

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindConsecutiveNumbers(t *testing.T) {
	tests := []struct {
		name     string
		logs     []ConsecutiveLog
		expected []int
	}{
		{
			name: "Example 1: Basic case with consecutive 1s",
			logs: []ConsecutiveLog{
				{ID: 1, Num: 1},
				{ID: 2, Num: 1},
				{ID: 3, Num: 1},
				{ID: 4, Num: 2},
				{ID: 5, Num: 1},
			},
			expected: []int{1},
		},
		{
			name: "Multiple consecutive numbers",
			logs: []ConsecutiveLog{
				{ID: 1, Num: 1},
				{ID: 2, Num: 1},
				{ID: 3, Num: 1},
				{ID: 4, Num: 2},
				{ID: 5, Num: 2},
				{ID: 6, Num: 2},
				{ID: 7, Num: 3},
			},
			expected: []int{1, 2},
		},
		{
			name: "No consecutive numbers",
			logs: []ConsecutiveLog{
				{ID: 1, Num: 1},
				{ID: 2, Num: 2},
				{ID: 3, Num: 3},
				{ID: 4, Num: 1},
				{ID: 5, Num: 2},
			},
			expected: []int{},
		},
		{
			name: "Empty logs",
			logs: []ConsecutiveLog{},
			expected: []int{},
		},
		{
			name: "Single log",
			logs: []ConsecutiveLog{
				{ID: 1, Num: 1},
			},
			expected: []int{},
		},
		{
			name: "Two logs",
			logs: []ConsecutiveLog{
				{ID: 1, Num: 1},
				{ID: 2, Num: 1},
			},
			expected: []int{},
		},
		{
			name: "Exactly three consecutive",
			logs: []ConsecutiveLog{
				{ID: 1, Num: 5},
				{ID: 2, Num: 5},
				{ID: 3, Num: 5},
			},
			expected: []int{5},
		},
		{
			name: "Multiple instances of same consecutive number",
			logs: []ConsecutiveLog{
				{ID: 1, Num: 1},
				{ID: 2, Num: 1},
				{ID: 3, Num: 1},
				{ID: 4, Num: 2},
				{ID: 5, Num: 1},
				{ID: 6, Num: 1},
				{ID: 7, Num: 1},
			},
			expected: []int{1},
		},
		{
			name: "Negative numbers",
			logs: []ConsecutiveLog{
				{ID: 1, Num: -1},
				{ID: 2, Num: -1},
				{ID: 3, Num: -1},
				{ID: 4, Num: 0},
			},
			expected: []int{-1},
		},
		{
			name: "Zero values",
			logs: []ConsecutiveLog{
				{ID: 1, Num: 0},
				{ID: 2, Num: 0},
				{ID: 3, Num: 0},
				{ID: 4, Num: 1},
			},
			expected: []int{0},
		},
		{
			name: "Large numbers",
			logs: []ConsecutiveLog{
				{ID: 1, Num: 1000000},
				{ID: 2, Num: 1000000},
				{ID: 3, Num: 1000000},
				{ID: 4, Num: 2000000},
			},
			expected: []int{1000000},
		},
		{
			name: "Long sequence with multiple groups",
			logs: []ConsecutiveLog{
				{ID: 1, Num: 1},
				{ID: 2, Num: 1},
				{ID: 3, Num: 1},
				{ID: 4, Num: 2},
				{ID: 5, Num: 3},
				{ID: 6, Num: 3},
				{ID: 7, Num: 3},
				{ID: 8, Num: 4},
				{ID: 9, Num: 5},
				{ID: 10, Num: 5},
				{ID: 11, Num: 5},
			},
			expected: []int{1, 3, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindConsecutiveNumbers(tt.logs)
			
			// Sort both slices for comparison (order doesn't matter)
			slices.Sort(result)
			expected := make([]int, len(tt.expected))
			copy(expected, tt.expected)
			slices.Sort(expected)
			
			assert.Equal(t, expected, result,
				"FindConsecutiveNumbers(%v) = %v, want %v",
				tt.logs, result, expected)
		})
	}
}

func TestFindConsecutiveNumbersOptimized(t *testing.T) {
	tests := []struct {
		name     string
		logs     []ConsecutiveLog
		expected []int
	}{
		{
			name: "Example 1: Basic case with consecutive 1s",
			logs: []ConsecutiveLog{
				{ID: 1, Num: 1},
				{ID: 2, Num: 1},
				{ID: 3, Num: 1},
				{ID: 4, Num: 2},
				{ID: 5, Num: 1},
			},
			expected: []int{1},
		},
		{
			name: "Non-sequential IDs (unsorted input)",
			logs: []ConsecutiveLog{
				{ID: 3, Num: 1},
				{ID: 1, Num: 1},
				{ID: 2, Num: 1},
				{ID: 4, Num: 2},
				{ID: 5, Num: 1},
			},
			expected: []int{1},
		},
		{
			name: "Gaps in IDs",
			logs: []ConsecutiveLog{
				{ID: 1, Num: 1},
				{ID: 2, Num: 1},
				{ID: 5, Num: 1}, // Gap here
				{ID: 6, Num: 2},
			},
			expected: []int{},
		},
		{
			name: "No consecutive numbers",
			logs: []ConsecutiveLog{
				{ID: 1, Num: 1},
				{ID: 2, Num: 2},
				{ID: 3, Num: 3},
			},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindConsecutiveNumbersOptimized(tt.logs)
			
			// Sort both slices for comparison
			slices.Sort(result)
			expected := make([]int, len(tt.expected))
			copy(expected, tt.expected)
			slices.Sort(expected)
			
			assert.Equal(t, expected, result,
				"FindConsecutiveNumbersOptimized(%v) = %v, want %v",
				tt.logs, result, expected)
		})
	}
}

func TestFindConsecutiveNumbersStreamingApproach(t *testing.T) {
	tests := []struct {
		name     string
		logs     []ConsecutiveLog
		expected []int
	}{
		{
			name: "Example 1: Basic case with consecutive 1s",
			logs: []ConsecutiveLog{
				{ID: 1, Num: 1},
				{ID: 2, Num: 1},
				{ID: 3, Num: 1},
				{ID: 4, Num: 2},
				{ID: 5, Num: 1},
			},
			expected: []int{1},
		},
		{
			name: "Multiple consecutive numbers",
			logs: []ConsecutiveLog{
				{ID: 1, Num: 1},
				{ID: 2, Num: 1},
				{ID: 3, Num: 1},
				{ID: 4, Num: 2},
				{ID: 5, Num: 2},
				{ID: 6, Num: 2},
			},
			expected: []int{1, 2},
		},
		{
			name: "No consecutive numbers",
			logs: []ConsecutiveLog{
				{ID: 1, Num: 1},
				{ID: 2, Num: 2},
				{ID: 3, Num: 3},
			},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindConsecutiveNumbersStreamingApproach(tt.logs)
			
			// Sort both slices for comparison
			slices.Sort(result)
			expected := make([]int, len(tt.expected))
			copy(expected, tt.expected)
			slices.Sort(expected)
			
			assert.Equal(t, expected, result,
				"FindConsecutiveNumbersStreamingApproach(%v) = %v, want %v",
				tt.logs, result, expected)
		})
	}
}

func BenchmarkFindConsecutiveNumbers(b *testing.B) {
	// Create a large dataset with some consecutive patterns
	logs := make([]ConsecutiveLog, 10000)
	for i := 0; i < 10000; i++ {
		logs[i] = ConsecutiveLog{
			ID:  i + 1,
			Num: (i % 100), // Creates many repeating patterns
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindConsecutiveNumbers(logs)
	}
}

func BenchmarkFindConsecutiveNumbersOptimized(b *testing.B) {
	logs := make([]ConsecutiveLog, 10000)
	for i := 0; i < 10000; i++ {
		logs[i] = ConsecutiveLog{
			ID:  i + 1,
			Num: (i % 100),
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindConsecutiveNumbersOptimized(logs)
	}
}

func BenchmarkFindConsecutiveNumbersStreamingApproach(b *testing.B) {
	logs := make([]ConsecutiveLog, 10000)
	for i := 0; i < 10000; i++ {
		logs[i] = ConsecutiveLog{
			ID:  i + 1,
			Num: (i % 100),
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindConsecutiveNumbersStreamingApproach(logs)
	}
}
