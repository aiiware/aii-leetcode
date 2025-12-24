package leetcode

import (
	"testing"
)

func TestSwapPairs(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, []int{2, 1, 4, 3}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3}, []int{2, 1, 3}},
		{[]int{1, 2, 3, 4, 5}, []int{2, 1, 4, 3, 5}},
		{[]int{1, 2, 3, 4, 5, 6}, []int{2, 1, 4, 3, 6, 5}},
	}

	for i, test := range tests {
		list := NewListFromSlice(test.input)

		// Test iterative version
		result := SwapPairs(list)
		expectedList := NewListFromSlice(test.expected)

		if !result.Equal(expectedList) {
			t.Errorf("Test %d (iterative) failed: got %v, expected %v",
				i, result.ToSlice(), test.expected)
		}

		// Test recursive version
		list = NewListFromSlice(test.input)
		result = SwapPairsRecursive(list)

		if !result.Equal(expectedList) {
			t.Errorf("Test %d (recursive) failed: got %v, expected %v",
				i, result.ToSlice(), test.expected)
		}
	}
}

func BenchmarkSwapPairs(b *testing.B) {
	list := NewListFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SwapPairs(list)
		// Reset list for next iteration
		list = NewListFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}

func BenchmarkSwapPairsRecursive(b *testing.B) {
	list := NewListFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SwapPairsRecursive(list)
		// Reset list for next iteration
		list = NewListFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	}
}