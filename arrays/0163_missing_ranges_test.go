package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindMissingRanges(t *testing.T) {
	tests := []struct {
		nums   []int
		lower  int
		upper  int
		expect []string
	}{
		{
			nums:   []int{0, 1, 3, 50, 75},
			lower:  0,
			upper:  99,
			expect: []string{"2", "4->49", "51->74", "76->99"},
		},
		{
			nums:   []int{-1},
			lower:  -1,
			upper:  -1,
			expect: []string{},
		},
		{
			nums:   []int{},
			lower:  1,
			upper:  1,
			expect: []string{"1"},
		},
		{
			nums:   []int{},
			lower:  -3,
			upper:  -1,
			expect: []string{"-3->-1"},
		},
		{
			nums:   []int{-1},
			lower:  -2,
			upper:  -1,
			expect: []string{"-2"},
		},
		{
			nums:   []int{-1},
			lower:  -1,
			upper:  0,
			expect: []string{"0"},
		},
		{
			nums:   []int{1, 3, 5, 7},
			lower:  0,
			upper:  10,
			expect: []string{"0", "2", "4", "6", "8->10"},
		},
		{
			nums:   []int{0, 1, 2, 3, 7},
			lower:  0,
			upper:  7,
			expect: []string{"4->6"},
		},
		{
			nums:   []int{},
			lower:  1,
			upper:  100,
			expect: []string{"1->100"},
		},
		{
			nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			lower:  1,
			upper:  10,
			expect: []string{},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			result := findMissingRanges(tt.nums, tt.lower, tt.upper)
			if !reflect.DeepEqual(result, tt.expect) {
				t.Errorf("findMissingRanges(%v, %d, %d) = %v, want %v",
					tt.nums, tt.lower, tt.upper, result, tt.expect)
			}
		})
	}
}

func BenchmarkFindMissingRanges(b *testing.B) {
	nums := []int{0, 1, 3, 50, 75}
	lower := 0
	upper := 99
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findMissingRanges(nums, lower, upper)
	}
}