package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Example 1",
			input:    []int{3, 0, 1},
			expected: 2,
		},
		{
			name:     "Example 2",
			input:    []int{0, 1},
			expected: 2,
		},
		{
			name:     "Example 3",
			input:    []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			expected: 8,
		},
		{
			name:     "Single element - missing 0",
			input:    []int{1},
			expected: 0,
		},
		{
			name:     "Single element - missing 1",
			input:    []int{0},
			expected: 1,
		},
		{
			name:     "Missing first number",
			input:    []int{1, 2, 3, 4},
			expected: 0,
		},
		{
			name:     "Missing last number",
			input:    []int{0, 1, 2, 3},
			expected: 4,
		},
		{
			name:     "Missing middle number",
			input:    []int{0, 1, 3, 4},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MissingNumber(tt.input)
			assert.Equal(t, tt.expected, result, "MissingNumber should find the missing number")
		})
	}
}

func TestMissingNumberXOR(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Example 1",
			input:    []int{3, 0, 1},
			expected: 2,
		},
		{
			name:     "Example 2",
			input:    []int{0, 1},
			expected: 2,
		},
		{
			name:     "Example 3",
			input:    []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			expected: 8,
		},
		{
			name:     "Single element",
			input:    []int{1},
			expected: 0,
		},
		{
			name:     "Missing last number",
			input:    []int{0, 1, 2, 3},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MissingNumberXOR(tt.input)
			assert.Equal(t, tt.expected, result, "MissingNumberXOR should find the missing number")
		})
	}
}

func TestMissingNumberSort(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Example 1",
			input:    []int{3, 0, 1},
			expected: 2,
		},
		{
			name:     "Example 2",
			input:    []int{0, 1},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy to avoid modifying the original
			nums := make([]int, len(tt.input))
			copy(nums, tt.input)
			
			result := MissingNumberSort(nums)
			assert.Equal(t, tt.expected, result, "MissingNumberSort should find the missing number")
		})
	}
}

func TestMissingNumberHash(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Example 1",
			input:    []int{3, 0, 1},
			expected: 2,
		},
		{
			name:     "Example 2",
			input:    []int{0, 1},
			expected: 2,
		},
		{
			name:     "Example 3",
			input:    []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			expected: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MissingNumberHash(tt.input)
			assert.Equal(t, tt.expected, result, "MissingNumberHash should find the missing number")
		})
	}
}

func TestMissingNumber_Consistency(t *testing.T) {
	// Test that all implementations give the same result
	testCases := [][]int{
		{3, 0, 1},
		{0, 1},
		{9, 6, 4, 2, 3, 5, 7, 0, 1},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, // Missing 10
	}

	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			result1 := MissingNumber(tc)
			result2 := MissingNumberXOR(tc)
			result3 := MissingNumberHash(tc)
			
			assert.Equal(t, result1, result2, "Standard and XOR methods should match")
			assert.Equal(t, result1, result3, "Standard and Hash methods should match")
		})
	}
}

func BenchmarkMissingNumber(b *testing.B) {
	benchmarks := []struct {
		name string
		size int
	}{
		{"Small", 10},
		{"Medium", 100},
		{"Large", 1000},
		{"Very Large", 10000},
	}

	for _, bm := range benchmarks {
		// Create test data
		nums := make([]int, bm.size)
		for i := 0; i < bm.size; i++ {
			nums[i] = i
		}
		// Remove the last number (which is bm.size)
		nums = nums[:bm.size-1]
		
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MissingNumber(nums)
			}
		})
	}
}

func BenchmarkMissingNumberXOR(b *testing.B) {
	nums := make([]int, 1000)
	for i := 0; i < 999; i++ {
		nums[i] = i
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MissingNumberXOR(nums)
	}
}

func BenchmarkMissingNumberHash(b *testing.B) {
	nums := make([]int, 1000)
	for i := 0; i < 999; i++ {
		nums[i] = i
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MissingNumberHash(nums)
	}
}
