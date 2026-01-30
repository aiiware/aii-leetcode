package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSumII(t *testing.T) {
	tests := []struct {
		numbers []int
		target  int
		expect  []int
	}{
		{
			numbers: []int{2, 7, 11, 15},
			target:  9,
			expect:  []int{1, 2},
		},
		{
			numbers: []int{2, 3, 4},
			target:  6,
			expect:  []int{1, 3},
		},
		{
			numbers: []int{-1, 0},
			target:  -1,
			expect:  []int{1, 2},
		},
		{
			numbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target:  17,
			// Both [7,10] and [8,9] sum to 17. Our two-pointer algorithm finds [7,10] first
			expect: []int{7, 10},
		},
		{
			numbers: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target:  19,
			expect:  []int{9, 10},
		},
		{
			numbers: []int{0, 0, 3, 4},
			target:  0,
			expect:  []int{1, 2},
		},
		{
			numbers: []int{-5, -3, -1, 0, 2, 4, 6},
			target:  1,
			// Multiple solutions: [-5,6]=[1,7], [-3,4]=[2,6], [-1,2]=[3,5]
			// Our two-pointer algorithm finds [-5,6] first
			expect: []int{1, 7},
		},
		{
			numbers: []int{1, 3, 4, 5, 7, 10, 11},
			target:  9,
			expect:  []int{3, 4},
		},
		{
			numbers: []int{5, 25, 75},
			target:  100,
			expect:  []int{2, 3},
		},
		{
			numbers: []int{1, 2, 3, 4, 4, 9, 56, 90},
			target:  8,
			expect:  []int{4, 5},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			result := twoSum(tt.numbers, tt.target)
			if !reflect.DeepEqual(result, tt.expect) {
				t.Errorf("twoSum(%v, %d) = %v, want %v", tt.numbers, tt.target, result, tt.expect)
			}
		})
	}
}

func BenchmarkTwoSumII(b *testing.B) {
	numbers := []int{2, 7, 11, 15}
	target := 9
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		twoSum(numbers, target)
	}
}

func BenchmarkTwoSumIILarge(b *testing.B) {
	// Create a large sorted array
	numbers := make([]int, 10000)
	for i := range numbers {
		numbers[i] = i * 2
	}
	target := 19998 // Sum of last two elements
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		twoSum(numbers, target)
	}
}

func BenchmarkTwoSumIIWorstCase(b *testing.B) {
	// Worst case: target is sum of first and last elements
	numbers := make([]int, 10000)
	for i := range numbers {
		numbers[i] = i
	}
	target := 9999 // Sum of first (0) and last (9999) elements
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		twoSum(numbers, target)
	}
}