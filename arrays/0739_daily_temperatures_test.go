package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		name         string
		temperatures []int
		expected     []int
	}{
		{
			name:         "Example 1",
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			expected:     []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			name:         "Example 2",
			temperatures: []int{30, 40, 50, 60},
			expected:     []int{1, 1, 1, 0},
		},
		{
			name:         "Example 3",
			temperatures: []int{30, 60, 90},
			expected:     []int{1, 1, 0},
		},
		{
			name:         "All decreasing temperatures",
			temperatures: []int{90, 80, 70, 60, 50},
			expected:     []int{0, 0, 0, 0, 0},
		},
		{
			name:         "All increasing temperatures",
			temperatures: []int{50, 60, 70, 80, 90},
			expected:     []int{1, 1, 1, 1, 0},
		},
		{
			name:         "Single temperature",
			temperatures: []int{100},
			expected:     []int{0},
		},
		{
			name:         "Empty array",
			temperatures: []int{},
			expected:     []int{},
		},
		{
			name:         "Same temperatures",
			temperatures: []int{70, 70, 70, 70},
			expected:     []int{0, 0, 0, 0},
		},
		{
			name:         "Mixed with same temperatures",
			temperatures: []int{70, 71, 70, 72},
			expected:     []int{1, 2, 1, 0},
		},
		{
			name:         "Large temperature jump",
			temperatures: []int{30, 31, 32, 33, 100},
			expected:     []int{1, 1, 1, 1, 0},
		},
		{
			name:         "Temperature drop then rise",
			temperatures: []int{80, 70, 60, 70, 80},
			expected:     []int{0, 3, 1, 1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DailyTemperatures(tt.temperatures)
			assert.Equal(t, tt.expected, result,
				"DailyTemperatures(%v) = %v, expected %v",
				tt.temperatures, result, tt.expected)
		})
	}
}

func TestDailyTemperatures_EdgeCases(t *testing.T) {
	t.Run("All zeros", func(t *testing.T) {
		result := DailyTemperatures([]int{0, 0, 0, 0})
		assert.Equal(t, []int{0, 0, 0, 0}, result)
	})

	t.Run("Negative temperatures", func(t *testing.T) {
		result := DailyTemperatures([]int{-10, -5, 0, 5})
		assert.Equal(t, []int{1, 1, 1, 0}, result)
	})

	t.Run("Alternating temperatures", func(t *testing.T) {
		result := DailyTemperatures([]int{10, 20, 10, 20, 10})
		assert.Equal(t, []int{1, 0, 1, 0, 0}, result)
	})

	t.Run("Large array", func(t *testing.T) {
		// Create array: 1, 2, 3, ..., 10000
		n := 10000
		temperatures := make([]int, n)
		for i := 0; i < n; i++ {
			temperatures[i] = i + 1
		}

		// Expected: 1, 1, 1, ..., 1, 0 (last element)
		expected := make([]int, n)
		for i := 0; i < n-1; i++ {
			expected[i] = 1
		}
		expected[n-1] = 0

		result := DailyTemperatures(temperatures)
		assert.Equal(t, expected, result)
	})
}

func BenchmarkDailyTemperatures(b *testing.B) {
	// Create a large array for benchmarking
	n := 100000
	temperatures := make([]int, n)
	for i := 0; i < n; i++ {
		temperatures[i] = i % 100 // Cycle temperatures between 0-99
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DailyTemperatures(temperatures)
	}
}