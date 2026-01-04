package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateRight(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		k        int
		expected []int
	}{
		{
			name:     "Example 1: [1,2,3,4,5], k=2",
			input:    []int{1, 2, 3, 4, 5},
			k:        2,
			expected: []int{4, 5, 1, 2, 3},
		},
		{
			name:     "Example 2: [0,1,2], k=4",
			input:    []int{0, 1, 2},
			k:        4,
			expected: []int{2, 0, 1},
		},
		{
			name:     "Empty list",
			input:    []int{},
			k:        3,
			expected: []int{},
		},
		{
			name:     "Single element, k=5",
			input:    []int{1},
			k:        5,
			expected: []int{1},
		},
		{
			name:     "Two elements, k=1",
			input:    []int{1, 2},
			k:        1,
			expected: []int{2, 1},
		},
		{
			name:     "Two elements, k=2",
			input:    []int{1, 2},
			k:        2,
			expected: []int{1, 2},
		},
		{
			name:     "Three elements, k=0",
			input:    []int{1, 2, 3},
			k:        0,
			expected: []int{1, 2, 3},
		},
		{
			name:     "Three elements, k=3",
			input:    []int{1, 2, 3},
			k:        3,
			expected: []int{1, 2, 3},
		},
		{
			name:     "Four elements, k=1",
			input:    []int{1, 2, 3, 4},
			k:        1,
			expected: []int{4, 1, 2, 3},
		},
		{
			name:     "Four elements, k=3",
			input:    []int{1, 2, 3, 4},
			k:        3,
			expected: []int{2, 3, 4, 1},
		},
		{
			name:     "Five elements, k=7 (k > length)",
			input:    []int{1, 2, 3, 4, 5},
			k:        7,
			expected: []int{4, 5, 1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := NewListFromSlice(tt.input)
			result := RotateRight(head, tt.k)
			assert.Equal(t, tt.expected, result.ToSlice(),
				"RotateRight(%v, %d) = %v, expected %v",
				tt.input, tt.k, result.ToSlice(), tt.expected)
		})
	}
}

func TestRotateRight_EdgeCases(t *testing.T) {
	t.Run("k equals length of list", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		head := NewListFromSlice(input)
		result := RotateRight(head, 5)
		assert.Equal(t, input, result.ToSlice(),
			"Rotating by length should return original list")
	})

	t.Run("k is multiple of length", func(t *testing.T) {
		input := []int{1, 2, 3}
		head := NewListFromSlice(input)
		result := RotateRight(head, 6) // 2 * length
		assert.Equal(t, input, result.ToSlice(),
			"Rotating by multiple of length should return original list")
	})

	t.Run("large k with small list", func(t *testing.T) {
		input := []int{1, 2}
		head := NewListFromSlice(input)
		result := RotateRight(head, 1001)
		// 1001 % 2 = 1, so should rotate by 1
		assert.Equal(t, []int{2, 1}, result.ToSlice())
	})

	t.Run("nil head", func(t *testing.T) {
		result := RotateRight(nil, 5)
		assert.Nil(t, result)
	})
}

func TestRotateRight_Consistency(t *testing.T) {
	// Test that rotating multiple times produces same result as rotating once with combined k
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	head1 := NewListFromSlice(input)
	head2 := NewListFromSlice(input)

	// Rotate head1 by 3, then by 4 (total 7)
	head1 = RotateRight(head1, 3)
	head1 = RotateRight(head1, 4)

	// Rotate head2 by 7 directly
	head2 = RotateRight(head2, 7)

	assert.Equal(t, head1.ToSlice(), head2.ToSlice(),
		"Rotating 3 then 4 should equal rotating 7")
}

func BenchmarkRotateRight(b *testing.B) {
	// Create a linked list with 1000 elements
	vals := make([]int, 1000)
	for i := range vals {
		vals[i] = i
	}

	testCases := []struct {
		name string
		k    int
	}{
		{"k=1", 1},
		{"k=10", 10},
		{"k=100", 100},
		{"k=500", 500},
		{"k=999", 999},
		{"k=1000", 1000},
		{"k=1500", 1500}, // k > length
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Clone the list for each iteration
				cloneHead := NewListFromSlice(vals)
				RotateRight(cloneHead, tc.k)
			}
		})
	}
}