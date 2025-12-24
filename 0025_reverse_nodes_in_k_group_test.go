package leetcode

import (
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		input    []int
		k        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 1, []int{1, 2, 3, 4, 5}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 3, []int{3, 2, 1, 6, 5, 4, 7, 8}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, 4, []int{4, 3, 2, 1, 8, 7, 6, 5}},
		{[]int{}, 1, []int{}},
	}

	for i, test := range tests {
		list := NewListFromSlice(test.input)

		// Test recursive version
		result := ReverseKGroup(list, test.k)
		expectedList := NewListFromSlice(test.expected)

		if !result.Equal(expectedList) {
			t.Errorf("Test %d (recursive) failed: got %v, expected %v",
				i, result.ToSlice(), test.expected)
		}

		// Test iterative version
		list = NewListFromSlice(test.input)
		result = ReverseKGroupIterative(list, test.k)

		if !result.Equal(expectedList) {
			t.Errorf("Test %d (iterative) failed: got %v, expected %v",
				i, result.ToSlice(), test.expected)
		}
	}
}

func BenchmarkReverseKGroup(b *testing.B) {
	list := NewListFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReverseKGroup(list, 3)
		// Reset list for next iteration
		list = NewListFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	}
}

func BenchmarkReverseKGroupIterative(b *testing.B) {
	list := NewListFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ReverseKGroupIterative(list, 3)
		// Reset list for next iteration
		list = NewListFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20})
	}
}