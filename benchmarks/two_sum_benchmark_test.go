package benchmarks

import (
	"leetcode/arrays"
	"testing"
)

// BenchmarkTwoSum benchmarks the Two Sum solution
func BenchmarkTwoSum(b *testing.B) {
	// Test case: find two numbers that sum to 9
	nums := []int{2, 7, 11, 15, 3, 6, 8, 1, 4, 5, 9, 10, 12, 13, 14}
	target := 9

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arrays.TwoSum(nums, target)
	}
}

// BenchmarkTwoSumSmall benchmarks with small array
func BenchmarkTwoSumSmall(b *testing.B) {
	nums := []int{2, 7, 11, 15}
	target := 9

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arrays.TwoSum(nums, target)
	}
}

// BenchmarkTwoSumLargeArray benchmarks with larger input
func BenchmarkTwoSumLargeArray(b *testing.B) {
	// Create a large array
	size := 10000
	nums := make([]int, size)
	for i := 0; i < size; i++ {
		nums[i] = i
	}
	target := 19999 // Last two elements should sum to this

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arrays.TwoSum(nums, target)
	}
}

// BenchmarkTwoSumNoSolution benchmarks when no solution exists
func BenchmarkTwoSumNoSolution(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5}
	target := 20 // No two numbers sum to this

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arrays.TwoSum(nums, target)
	}
}

// BenchmarkTwoSumDuplicateValues benchmarks with duplicate values
func BenchmarkTwoSumDuplicateValues(b *testing.B) {
	nums := []int{3, 3, 4, 4, 5, 5, 6, 6}
	target := 7

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		arrays.TwoSum(nums, target)
	}
}