package arrays

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			name:     "Example 2",
			nums:     []int{0},
			expected: []int{0},
		},
		{
			name:     "No zeros",
			nums:     []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0, 0},
			expected: []int{0, 0, 0, 0},
		},
		{
			name:     "Zeros at beginning",
			nums:     []int{0, 0, 1, 2, 3},
			expected: []int{1, 2, 3, 0, 0},
		},
		{
			name:     "Zeros at end",
			nums:     []int{1, 2, 3, 0, 0},
			expected: []int{1, 2, 3, 0, 0},
		},
		{
			name:     "Zeros in middle",
			nums:     []int{1, 0, 2, 0, 3},
			expected: []int{1, 2, 3, 0, 0},
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: []int{},
		},
		{
			name:     "Single non-zero",
			nums:     []int{42},
			expected: []int{42},
		},
		{
			name:     "Mixed with negative numbers",
			nums:     []int{0, -1, 0, 3, -12},
			expected: []int{-1, 3, -12, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of the input for modification
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)

			MoveZeroes(nums)

			if !reflect.DeepEqual(nums, tt.expected) {
				t.Errorf("MoveZeroes() = %v, expected %v", nums, tt.expected)
			}
		})
	}
}

func TestMoveZeroesSinglePass(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			name:     "Example 2",
			nums:     []int{0},
			expected: []int{0},
		},
		{
			name:     "No zeros",
			nums:     []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0, 0},
			expected: []int{0, 0, 0, 0},
		},
		{
			name:     "Zeros at beginning",
			nums:     []int{0, 0, 1, 2, 3},
			expected: []int{1, 2, 3, 0, 0},
		},
		{
			name:     "Zeros at end",
			nums:     []int{1, 2, 3, 0, 0},
			expected: []int{1, 2, 3, 0, 0},
		},
		{
			name:     "Zeros in middle",
			nums:     []int{1, 0, 2, 0, 3},
			expected: []int{1, 2, 3, 0, 0},
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: []int{},
		},
		{
			name:     "Single non-zero",
			nums:     []int{42},
			expected: []int{42},
		},
		{
			name:     "Mixed with negative numbers",
			nums:     []int{0, -1, 0, 3, -12},
			expected: []int{-1, 3, -12, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of the input for modification
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)

			MoveZeroesSinglePass(nums)

			if !reflect.DeepEqual(nums, tt.expected) {
				t.Errorf("MoveZeroesSinglePass() = %v, expected %v", nums, tt.expected)
			}
		})
	}
}

func TestMoveZeroesSnowball(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			name:     "Example 2",
			nums:     []int{0},
			expected: []int{0},
		},
		{
			name:     "No zeros",
			nums:     []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "All zeros",
			nums:     []int{0, 0, 0, 0},
			expected: []int{0, 0, 0, 0},
		},
		{
			name:     "Zeros at beginning",
			nums:     []int{0, 0, 1, 2, 3},
			expected: []int{1, 2, 3, 0, 0},
		},
		{
			name:     "Zeros at end",
			nums:     []int{1, 2, 3, 0, 0},
			expected: []int{1, 2, 3, 0, 0},
		},
		{
			name:     "Zeros in middle",
			nums:     []int{1, 0, 2, 0, 3},
			expected: []int{1, 2, 3, 0, 0},
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: []int{},
		},
		{
			name:     "Single non-zero",
			nums:     []int{42},
			expected: []int{42},
		},
		{
			name:     "Mixed with negative numbers",
			nums:     []int{0, -1, 0, 3, -12},
			expected: []int{-1, 3, -12, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of the input for modification
			nums := make([]int, len(tt.nums))
			copy(nums, tt.nums)

			MoveZeroesSnowball(nums)

			if !reflect.DeepEqual(nums, tt.expected) {
				t.Errorf("MoveZeroesSnowball() = %v, expected %v", nums, tt.expected)
			}
		})
	}
}

func BenchmarkMoveZeroes(b *testing.B) {
	nums := []int{0, 1, 0, 3, 12, 0, 0, 5, 6, 0, 7, 8, 0, 9, 10, 0, 11, 12, 0, 13}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a fresh copy for each iteration
		testNums := make([]int, len(nums))
		copy(testNums, nums)
		MoveZeroes(testNums)
	}
}

func BenchmarkMoveZeroesSinglePass(b *testing.B) {
	nums := []int{0, 1, 0, 3, 12, 0, 0, 5, 6, 0, 7, 8, 0, 9, 10, 0, 11, 12, 0, 13}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a fresh copy for each iteration
		testNums := make([]int, len(nums))
		copy(testNums, nums)
		MoveZeroesSinglePass(testNums)
	}
}

func BenchmarkMoveZeroesSnowball(b *testing.B) {
	nums := []int{0, 1, 0, 3, 12, 0, 0, 5, 6, 0, 7, 8, 0, 9, 10, 0, 11, 12, 0, 13}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a fresh copy for each iteration
		testNums := make([]int, len(nums))
		copy(testNums, nums)
		MoveZeroesSnowball(testNums)
	}
}