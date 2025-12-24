package leetcode

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
		modified []int
	}{
		{[]int{1, 1, 2}, 2, []int{1, 2}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
		{[]int{}, 0, []int{}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
		{[]int{1, 1, 1, 1}, 1, []int{1}},
	}

	for i, test := range tests {
		// Create a copy of the input for modification
		nums := make([]int, len(test.input))
		copy(nums, test.input)

		// Test the function
		result := RemoveDuplicates(nums)

		if result != test.expected {
			t.Errorf("Test %d failed: got count %d, expected %d",
				i, result, test.expected)
		}

		// Check the modified array
		if !reflect.DeepEqual(nums[:result], test.modified) {
			t.Errorf("Test %d failed: got array %v, expected %v",
				i, nums[:result], test.modified)
		}
	}
}

func TestRemoveDuplicatesGeneral(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 1, 2}, []int{1, 2}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 1, 1, 1}, []int{1}},
	}

	for i, test := range tests {
		// Create a copy of the input for modification
		nums := make([]int, len(test.input))
		copy(nums, test.input)

		// Test the function
		result := RemoveDuplicatesGeneral(nums)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Test %d failed: got %v, expected %v",
				i, result, test.expected)
		}
	}
}

func BenchmarkRemoveDuplicates(b *testing.B) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a fresh copy for each iteration
		testNums := make([]int, len(nums))
		copy(testNums, nums)
		RemoveDuplicates(testNums)
	}
}

func BenchmarkRemoveDuplicatesGeneral(b *testing.B) {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a fresh copy for each iteration
		testNums := make([]int, len(nums))
		copy(testNums, nums)
		RemoveDuplicatesGeneral(testNums)
	}
}