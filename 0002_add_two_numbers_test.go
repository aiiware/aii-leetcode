package leetcode

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name     string
		l1       []int
		l2       []int
		expected []int
	}{
		// Example from LeetCode
		{
			name:     "Example 1",
			l1:       []int{2, 4, 3},
			l2:       []int{5, 6, 4},
			expected: []int{7, 0, 8},
		},
		// Additional test cases
		{
			name:     "Different lengths",
			l1:       []int{9, 9, 9, 9, 9, 9, 9},
			l2:       []int{9, 9, 9, 9},
			expected: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
		{
			name:     "Zero plus zero",
			l1:       []int{0},
			l2:       []int{0},
			expected: []int{0},
		},
		{
			name:     "Carry at the end",
			l1:       []int{9, 9},
			l2:       []int{1},
			expected: []int{0, 0, 1},
		},
		{
			name:     "One empty list (should handle nil)",
			l1:       []int{1, 2, 3},
			l2:       []int{},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Large numbers",
			l1:       []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			l2:       []int{5, 6, 4},
			expected: []int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := NewListFromSlice(tt.l1)
			l2 := NewListFromSlice(tt.l2)
			expected := NewListFromSlice(tt.expected)

			result := AddTwoNumbers(l1, l2)

			if !result.Equal(expected) {
				t.Errorf("AddTwoNumbers(%v, %v) = %v, expected %v",
					tt.l1, tt.l2, result.ToSlice(), tt.expected)
			}
		})
	}
}

func TestAddTwoNumbersEdgeCases(t *testing.T) {
	// Test with nil lists
	t.Run("Both nil lists", func(t *testing.T) {
		result := AddTwoNumbers(nil, nil)
		if result != nil {
			t.Errorf("AddTwoNumbers(nil, nil) = %v, expected nil", result)
		}
	})

	// Test single digit with carry
	t.Run("Single digit with carry", func(t *testing.T) {
		l1 := NewListFromSlice([]int{5})
		l2 := NewListFromSlice([]int{5})
		result := AddTwoNumbers(l1, l2)
		expected := NewListFromSlice([]int{0, 1})
		if !result.Equal(expected) {
			t.Errorf("AddTwoNumbers([5], [5]) = %v, expected [0, 1]", result.ToSlice())
		}
	})
}

func BenchmarkAddTwoNumbers(b *testing.B) {
	// Create large lists for benchmarking
	l1 := NewListFromSlice([]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9})
	l2 := NewListFromSlice([]int{1})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AddTwoNumbers(l1, l2)
	}
}

func BenchmarkAddTwoNumbersLarge(b *testing.B) {
	// Create even larger lists
	l1Vals := make([]int, 1000)
	l2Vals := make([]int, 1000)
	for i := range l1Vals {
		l1Vals[i] = 9
		l2Vals[i] = 1
	}
	l1 := NewListFromSlice(l1Vals)
	l2 := NewListFromSlice(l2Vals)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AddTwoNumbers(l1, l2)
	}
}