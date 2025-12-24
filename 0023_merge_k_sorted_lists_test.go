package leetcode

import (
	"testing"
)

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		lists    [][]int
		expected []int
	}{
		{
			[][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			[]int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			[][]int{},
			[]int{},
		},
		{
			[][]int{{}},
			[]int{},
		},
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			[][]int{{1}, {0}},
			[]int{0, 1},
		},
	}

	for i, test := range tests {
		// Convert slices to lists
		lists := make([]*ListNode, len(test.lists))
		for j, arr := range test.lists {
			lists[j] = NewListFromSlice(arr)
		}

		// Test heap approach
		result := MergeKLists(lists)
		expectedList := NewListFromSlice(test.expected)

		if !result.Equal(expectedList) {
			t.Errorf("Test %d (heap) failed: got %v, expected %v",
				i, result.ToSlice(), test.expected)
		}

		// Test divide and conquer approach
		lists = make([]*ListNode, len(test.lists))
		for j, arr := range test.lists {
			lists[j] = NewListFromSlice(arr)
		}
		result = MergeKListsDivideConquer(lists)

		if !result.Equal(expectedList) {
			t.Errorf("Test %d (divide & conquer) failed: got %v, expected %v",
				i, result.ToSlice(), test.expected)
		}
	}
}

func BenchmarkMergeKListsHeap(b *testing.B) {
	lists := []*ListNode{
		NewListFromSlice([]int{1, 4, 7, 10, 13}),
		NewListFromSlice([]int{2, 5, 8, 11, 14}),
		NewListFromSlice([]int{3, 6, 9, 12, 15}),
		NewListFromSlice([]int{16, 19, 22, 25, 28}),
		NewListFromSlice([]int{17, 20, 23, 26, 29}),
		NewListFromSlice([]int{18, 21, 24, 27, 30}),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeKLists(lists)
		// Reset lists
		lists = []*ListNode{
			NewListFromSlice([]int{1, 4, 7, 10, 13}),
			NewListFromSlice([]int{2, 5, 8, 11, 14}),
			NewListFromSlice([]int{3, 6, 9, 12, 15}),
			NewListFromSlice([]int{16, 19, 22, 25, 28}),
			NewListFromSlice([]int{17, 20, 23, 26, 29}),
			NewListFromSlice([]int{18, 21, 24, 27, 30}),
		}
	}
}

func BenchmarkMergeKListsDivideConquer(b *testing.B) {
	lists := []*ListNode{
		NewListFromSlice([]int{1, 4, 7, 10, 13}),
		NewListFromSlice([]int{2, 5, 8, 11, 14}),
		NewListFromSlice([]int{3, 6, 9, 12, 15}),
		NewListFromSlice([]int{16, 19, 22, 25, 28}),
		NewListFromSlice([]int{17, 20, 23, 26, 29}),
		NewListFromSlice([]int{18, 21, 24, 27, 30}),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeKListsDivideConquer(lists)
		// Reset lists
		lists = []*ListNode{
			NewListFromSlice([]int{1, 4, 7, 10, 13}),
			NewListFromSlice([]int{2, 5, 8, 11, 14}),
			NewListFromSlice([]int{3, 6, 9, 12, 15}),
			NewListFromSlice([]int{16, 19, 22, 25, 28}),
			NewListFromSlice([]int{17, 20, 23, 26, 29}),
			NewListFromSlice([]int{18, 21, 24, 27, 30}),
		}
	}
}