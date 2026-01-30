package arrays

import (
	"fmt"
	"testing"
)

func TestMaximumGap(t *testing.T) {
	tests := []struct {
		nums   []int
		expect int
	}{
		{
			nums:   []int{3, 6, 9, 1},
			expect: 3,
		},
		{
			nums:   []int{10},
			expect: 0,
		},
		{
			nums:   []int{1, 10, 5},
			expect: 5,
		},
		{
			nums:   []int{1, 2, 3, 4, 5},
			expect: 1,
		},
		{
			nums:   []int{100, 3, 2, 1},
			expect: 97,
		},
		{
			nums:   []int{1, 1000000000},
			expect: 999999999,
		},
		{
			nums:   []int{1, 1, 1, 1},
			expect: 0,
		},
		{
			nums:   []int{1, 3, 100},
			expect: 97,
		},
		{
			nums:   []int{},
			expect: 0,
		},
		{
			nums:   []int{15252, 16764, 27963, 7817, 26155, 20757, 3478, 22602, 20404, 6739, 16790, 10588, 16521, 6644, 20880, 15632, 27078, 25463, 20124, 15728, 30042, 16604, 17223, 4388, 23646, 32683, 23688, 12439, 30630, 3895, 7926, 22101, 32406, 21540, 31799, 3768, 26679, 21799, 23740},
			expect: 2901,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			result := maximumGap(tt.nums)
			if result != tt.expect {
				t.Errorf("maximumGap(%v) = %d, want %d", tt.nums, result, tt.expect)
			}
		})
	}
}

func BenchmarkMaximumGap(b *testing.B) {
	nums := []int{3, 6, 9, 1}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maximumGap(nums)
	}
}

func BenchmarkMaximumGapLarge(b *testing.B) {
	// Create a large array for benchmarking
	nums := make([]int, 10000)
	for i := range nums {
		nums[i] = i * 100
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		maximumGap(nums)
	}
}