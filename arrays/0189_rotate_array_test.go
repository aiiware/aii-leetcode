package arrays

import (
	"reflect"
	"testing"
)

func TestRotateArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			k:        3,
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name:     "Example 2",
			nums:     []int{-1, -100, 3, 99},
			k:        2,
			expected: []int{3, 99, -1, -100},
		},
		{
			name:     "Single element",
			nums:     []int{1},
			k:        5,
			expected: []int{1},
		},
		{
			name:     "Empty array",
			nums:     []int{},
			k:        3,
			expected: []int{},
		},
		{
			name:     "k = 0",
			nums:     []int{1, 2, 3},
			k:        0,
			expected: []int{1, 2, 3},
		},
		{
			name:     "k equals array length",
			nums:     []int{1, 2, 3, 4, 5},
			k:        5,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "k greater than array length",
			nums:     []int{1, 2, 3, 4, 5},
			k:        8,
			expected: []int{3, 4, 5, 1, 2},
		},
		{
			name:     "Negative numbers",
			nums:     []int{-10, -5, 0, 5, 10},
			k:        2,
			expected: []int{5, 10, -10, -5, 0},
		},
		{
			name:     "Large k with small array",
			nums:     []int{1, 2},
			k:        3,
			expected: []int{2, 1},
		},
		{
			name:     "All same elements",
			nums:     []int{7, 7, 7, 7, 7},
			k:        3,
			expected: []int{7, 7, 7, 7, 7},
		},
		{
			name:     "k = 1 with large array",
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			k:        1,
			expected: []int{10, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:     "k = n-1",
			nums:     []int{1, 2, 3, 4, 5},
			k:        4,
			expected: []int{2, 3, 4, 5, 1},
		},
	}

	// Test all solution functions
	solutionFuncs := []struct {
		name string
		f    func([]int, int)
	}{
		{"rotate", rotate},
		{"rotateExtraArray", rotateExtraArray},
		{"rotateCyclic", rotateCyclic},
		{"rotateBuiltin", rotateBuiltin},
		{"rotateBruteForce", rotateBruteForce},
		{"rotateBlockSwap", rotateBlockSwap},
		{"rotateGCD", rotateGCD},
		{"rotateRecursive", rotateRecursive},
	}

	for _, sf := range solutionFuncs {
		t.Run(sf.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					// Create a copy of the input array
					nums := make([]int, len(tt.nums))
					copy(nums, tt.nums)

					// Apply the rotation
					sf.f(nums, tt.k)

					// Check the result
					if !reflect.DeepEqual(nums, tt.expected) {
						t.Errorf("%s(%v, %d) = %v, expected %v",
							sf.name, tt.nums, tt.k, nums, tt.expected)
					}
				})
			}
		})
	}
}

func TestRotateEdgeCases(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
	}{
		{
			name: "Very large k",
			nums: []int{1, 2, 3, 4, 5},
			k:    1000000,
		},
		{
			name: "k = 0 with empty array",
			nums: []int{},
			k:    0,
		},
		{
			name: "Single element with k = 1",
			nums: []int{42},
			k:    1,
		},
		{
			name: "Two elements with k = 1",
			nums: []int{1, 2},
			k:    1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test all solutions
			solutions := []func([]int, int){
				rotate,
				rotateExtraArray,
				rotateCyclic,
				rotateBuiltin,
				rotateBruteForce,
				rotateBlockSwap,
				rotateGCD,
				rotateRecursive,
			}

			// For each solution, verify it doesn't panic
			for i, f := range solutions {
				nums := make([]int, len(tt.nums))
				copy(nums, tt.nums)

				// This should not panic
				f(nums, tt.k)

				// Verify the array still has the same length
				if len(nums) != len(tt.nums) {
					t.Errorf("Solution %d changed array length from %d to %d",
						i, len(tt.nums), len(nums))
				}
			}
		})
	}
}

func TestRotateConsistency(t *testing.T) {
	// Test that all solutions produce the same result
	testCases := []struct {
		name string
		nums []int
		k    int
	}{
		{"Case 1", []int{1, 2, 3, 4, 5, 6, 7}, 3},
		{"Case 2", []int{-1, -100, 3, 99}, 2},
		{"Case 3", []int{1, 2, 3, 4, 5}, 13},
		{"Case 4", []int{10, 20, 30, 40, 50, 60}, 4},
		{"Case 5", []int{5, 4, 3, 2, 1}, 2},
	}

	solutions := []struct {
		name string
		f    func([]int, int)
	}{
		{"rotate", rotate},
		{"rotateExtraArray", rotateExtraArray},
		{"rotateCyclic", rotateCyclic},
		{"rotateBuiltin", rotateBuiltin},
		{"rotateBruteForce", rotateBruteForce},
		{"rotateBlockSwap", rotateBlockSwap},
		{"rotateGCD", rotateGCD},
		{"rotateRecursive", rotateRecursive},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			results := make([][]int, len(solutions))

			// Run each solution
			for i, sf := range solutions {
				nums := make([]int, len(tc.nums))
				copy(nums, tc.nums)
				sf.f(nums, tc.k)
				results[i] = nums
			}

			// Verify all results are the same
			for i := 1; i < len(results); i++ {
				if !reflect.DeepEqual(results[0], results[i]) {
					t.Errorf("Solution %s produced different result than %s: %v vs %v",
						solutions[i].name, solutions[0].name, results[i], results[0])
				}
			}
		})
	}
}

func BenchmarkRotateArray(b *testing.B) {
	benchmarks := []struct {
		name string
		f    func([]int, int)
	}{
		{"rotate", rotate},
		{"rotateExtraArray", rotateExtraArray},
		{"rotateCyclic", rotateCyclic},
		{"rotateBuiltin", rotateBuiltin},
		{"rotateBruteForce", rotateBruteForce},
		{"rotateBlockSwap", rotateBlockSwap},
		{"rotateGCD", rotateGCD},
		{"rotateRecursive", rotateRecursive},
	}

	// Test cases of different sizes
	testCases := []struct {
		name string
		nums []int
		k    int
	}{
		{"Small", make([]int, 100), 30},
		{"Medium", make([]int, 1000), 300},
		{"Large", make([]int, 10000), 3000},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for _, bm := range benchmarks {
				b.Run(bm.name, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						// Reset array for each iteration
						nums := make([]int, len(tc.nums))
						copy(nums, tc.nums)
						bm.f(nums, tc.k)
					}
				})
			}
		})
	}
}

func TestRotateSpecialCases(t *testing.T) {
	// Test special mathematical properties
	t.Run("Identity rotation", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		expected := []int{1, 2, 3, 4, 5}

		// Rotating by array length should give same array
		rotate(nums, len(nums))
		if !reflect.DeepEqual(nums, expected) {
			t.Errorf("Rotating by array length should give identity: got %v", nums)
		}

		// Rotating by 0 should give same array
		nums = []int{1, 2, 3, 4, 5}
		rotate(nums, 0)
		if !reflect.DeepEqual(nums, expected) {
			t.Errorf("Rotating by 0 should give identity: got %v", nums)
		}
	})

	t.Run("Double rotation property", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		k1, k2 := 3, 4

		// Copy for first approach
		nums1 := make([]int, len(nums))
		copy(nums1, nums)
		rotate(nums1, k1)
		rotate(nums1, k2)

		// Copy for second approach
		nums2 := make([]int, len(nums))
		copy(nums2, nums)
		rotate(nums2, k1+k2)

		if !reflect.DeepEqual(nums1, nums2) {
			t.Errorf("rotate(rotate(nums, %d), %d) != rotate(nums, %d): %v vs %v",
				k1, k2, k1+k2, nums1, nums2)
		}
	})

	t.Run("Reverse rotation", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		k := 3

		// Rotate forward by k
		numsForward := make([]int, len(nums))
		copy(numsForward, nums)
		rotate(numsForward, k)

		// Rotate backward by k (which is rotate by n-k forward)
		numsBackward := make([]int, len(nums))
		copy(numsBackward, numsForward)
		rotate(numsBackward, len(nums)-k)

		if !reflect.DeepEqual(numsBackward, nums) {
			t.Errorf("Rotating forward by %d then backward by %d should give original: got %v",
				k, len(nums)-k, numsBackward)
		}
	})
}
