package leetcode

import (
	"fmt"
	"testing"
)

func TestLargestNumber(t *testing.T) {
	tests := []struct {
		input    []int
		expected string
	}{
		// Test cases from LeetCode examples
		{[]int{10, 2}, "210"},
		{[]int{3, 30, 34, 5, 9}, "9534330"},
		{[]int{0, 0}, "0"},

		// Additional edge cases
		{[]int{1}, "1"},
		{[]int{0}, "0"},
		{[]int{1, 0, 0}, "100"},
		{[]int{10, 1}, "110"},
		{[]int{1, 10}, "110"},
		{[]int{9, 99, 999}, "999999"},
		{[]int{999, 99, 9}, "999999"},
		{[]int{121, 12}, "12121"},
		{[]int{12, 121}, "12121"},
		{[]int{3, 30, 34}, "34330"},
		{[]int{34, 3, 30}, "34330"},
		{[]int{432, 43243}, "43243432"},
		{[]int{43243, 432}, "43243432"},

		// Large numbers
		{[]int{999999999, 999999998, 999999997}, "999999999999999998999999997"},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}, "9876543210"},

		// All zeros
		{[]int{0, 0, 0, 0, 0}, "0"},

		// Single digit numbers
		{[]int{5, 4, 3, 2, 1}, "54321"},
		{[]int{1, 2, 3, 4, 5}, "54321"},

		// Mixed numbers
		{[]int{824, 938, 1399, 5607, 6973, 5703, 9609, 4398, 8247}, "9609938824824769735703560743981399"},
		{[]int{128, 12}, "12812"},
		{[]int{12, 128}, "12812"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			result := largestNumber(tt.input)
			if result != tt.expected {
				t.Errorf("largestNumber(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// Benchmark tests for performance
func BenchmarkLargestNumber(b *testing.B) {
	testCases := [][]int{
		{3, 30, 34, 5, 9},
		{10, 2},
		{0, 0},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		{824, 938, 1399, 5607, 6973, 5703, 9609, 4398, 8247},
	}

	for i := 0; i < b.N; i++ {
		for _, nums := range testCases {
			largestNumber(nums)
		}
	}
}

func TestLargestNumberEdgeCases(t *testing.T) {
	// Test with very large numbers
	t.Run("Large numbers", func(t *testing.T) {
		nums := []int{999999999, 888888888, 777777777}
		result := largestNumber(nums)
		expected := "999999999888888888777777777"
		if result != expected {
			t.Errorf("largestNumber(%v) = %v, want %v", nums, result, expected)
		}
	})

	// Test with numbers having same prefix
	t.Run("Same prefix numbers", func(t *testing.T) {
		nums := []int{12, 121, 1212}
		result := largestNumber(nums)
		// 121212121 > 12121212 > 121212
		expected := "121212121"
		if result != expected {
			t.Errorf("largestNumber(%v) = %v, want %v", nums, result, expected)
		}
	})

	// Test with single element
	t.Run("Single element", func(t *testing.T) {
		nums := []int{42}
		result := largestNumber(nums)
		expected := "42"
		if result != expected {
			t.Errorf("largestNumber(%v) = %v, want %v", nums, result, expected)
		}
	})

	// Test with all same numbers
	t.Run("All same numbers", func(t *testing.T) {
		nums := []int{7, 7, 7, 7}
		result := largestNumber(nums)
		expected := "7777"
		if result != expected {
			t.Errorf("largestNumber(%v) = %v, want %v", nums, result, expected)
		}
	})
}