package leetcode

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		nums     []int
		val      int
		expected int
		modified []int
	}{
		{[]int{3, 2, 2, 3}, 3, 2, []int{2, 2}},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5, []int{0, 1, 3, 0, 4}},
		{[]int{}, 0, 0, []int{}},
		{[]int{1}, 1, 0, []int{}},
		{[]int{1, 2, 3, 4}, 5, 4, []int{1, 2, 3, 4}},
		{[]int{4, 4, 4, 4}, 4, 0, []int{}},
	}

	for i, test := range tests {
		// Create a copy of the input for modification
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)

		// Test the standard approach
		result := RemoveElement(nums, test.val)

		if result != test.expected {
			t.Errorf("Test %d (standard) failed: got count %d, expected %d",
				i, result, test.expected)
		}

		// Sort and compare the modified portion (order may differ)
		actual := nums[:result]
		expected := test.modified

		// Sort both slices for comparison
		sortInts(actual)
		sortInts(expected)

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Test %d (standard) failed: got array %v, expected %v",
				i, nums[:result], test.modified)
		}
	}
}

func TestRemoveElementTwoPointers(t *testing.T) {
	tests := []struct {
		nums     []int
		val      int
		expected int
		modified []int
	}{
		{[]int{3, 2, 2, 3}, 3, 2, []int{2, 2}},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5, []int{0, 1, 3, 0, 4}},
		{[]int{}, 0, 0, []int{}},
		{[]int{1}, 1, 0, []int{}},
		{[]int{1, 2, 3, 4}, 5, 4, []int{1, 2, 3, 4}},
		{[]int{4, 4, 4, 4}, 4, 0, []int{}},
	}

	for i, test := range tests {
		// Create a copy of the input for modification
		nums := make([]int, len(test.nums))
		copy(nums, test.nums)

		// Test the two pointers approach
		result := RemoveElementTwoPointers(nums, test.val)

		if result != test.expected {
			t.Errorf("Test %d (two pointers) failed: got count %d, expected %d",
				i, result, test.expected)
		}

		// Sort and compare the modified portion (order may differ)
		actual := nums[:result]
		expected := test.modified

		// Sort both slices for comparison
		sortInts(actual)
		sortInts(expected)

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Test %d (two pointers) failed: got array %v, expected %v",
				i, nums[:result], test.modified)
		}
	}
}

// Helper function to sort integers
func sortInts(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func BenchmarkRemoveElement(b *testing.B) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2, 5, 6, 7, 8, 9, 10, 2, 3, 4, 5, 6, 7}
	val := 2

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a fresh copy for each iteration
		testNums := make([]int, len(nums))
		copy(testNums, nums)
		RemoveElement(testNums, val)
	}
}

func BenchmarkRemoveElementTwoPointers(b *testing.B) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2, 5, 6, 7, 8, 9, 10, 2, 3, 4, 5, 6, 7}
	val := 2

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a fresh copy for each iteration
		testNums := make([]int, len(nums))
		copy(testNums, nums)
		RemoveElementTwoPointers(testNums, val)
	}
}