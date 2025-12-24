package leetcode

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		list1    []int
		list2    []int
		expected []int
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
		{[]int{}, []int{}, []int{}},
		{[]int{}, []int{0}, []int{0}},
		{[]int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 3}, []int{}, []int{1, 2, 3}},
		{[]int{5}, []int{1, 2, 4}, []int{1, 2, 4, 5}},
	}

	for i, test := range tests {
		list1 := NewListFromSlice(test.list1)
		list2 := NewListFromSlice(test.list2)

		// Test iterative version
		result := MergeTwoLists(list1, list2)
		expectedList := NewListFromSlice(test.expected)

		if !result.Equal(expectedList) {
			t.Errorf("Test %d (iterative) failed: got %v, expected %v",
				i, result.ToSlice(), test.expected)
		}

		// Test recursive version
		list1 = NewListFromSlice(test.list1)
		list2 = NewListFromSlice(test.list2)
		result = MergeTwoListsRecursive(list1, list2)

		if !result.Equal(expectedList) {
			t.Errorf("Test %d (recursive) failed: got %v, expected %v",
				i, result.ToSlice(), test.expected)
		}
	}
}

func BenchmarkMergeTwoLists(b *testing.B) {
	list1 := NewListFromSlice([]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19})
	list2 := NewListFromSlice([]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeTwoLists(list1, list2)
		// Reset lists for next iteration
		list1 = NewListFromSlice([]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19})
		list2 = NewListFromSlice([]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20})
	}
}

func BenchmarkMergeTwoListsRecursive(b *testing.B) {
	list1 := NewListFromSlice([]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19})
	list2 := NewListFromSlice([]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeTwoListsRecursive(list1, list2)
		// Reset lists for next iteration
		list1 = NewListFromSlice([]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19})
		list2 = NewListFromSlice([]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20})
	}
}