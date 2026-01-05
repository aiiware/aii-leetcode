package leetcode

import (
	"fmt"
	"testing"
)

func TestMergeSortedArray(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{
			name:     "Example 1",
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 5, 6},
			n:        3,
			expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:     "Example 2",
			nums1:    []int{1},
			m:        1,
			nums2:    []int{},
			n:        0,
			expected: []int{1},
		},
		{
			name:     "Example 3",
			nums1:    []int{0},
			m:        0,
			nums2:    []int{1},
			n:        1,
			expected: []int{1},
		},
		{
			name:     "Empty both arrays",
			nums1:    []int{},
			m:        0,
			nums2:    []int{},
			n:        0,
			expected: []int{},
		},
		{
			name:     "nums1 empty, nums2 has elements",
			nums1:    []int{0, 0, 0},
			m:        0,
			nums2:    []int{1, 2, 3},
			n:        3,
			expected: []int{1, 2, 3},
		},
		{
			name:     "nums2 empty, nums1 has elements",
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{},
			n:        0,
			expected: []int{1, 2, 3},
		},
		{
			name:     "All elements in nums1 smaller",
			nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			nums2:    []int{4, 5, 6},
			n:        3,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "All elements in nums2 smaller",
			nums1:    []int{4, 5, 6, 0, 0, 0},
			m:        3,
			nums2:    []int{1, 2, 3},
			n:        3,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "Interleaved elements",
			nums1:    []int{1, 3, 5, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 4, 6},
			n:        3,
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "Duplicate elements",
			nums1:    []int{1, 2, 2, 0, 0, 0},
			m:        3,
			nums2:    []int{2, 3, 4},
			n:        3,
			expected: []int{1, 2, 2, 2, 3, 4},
		},
		{
			name:     "Single element arrays",
			nums1:    []int{2, 0},
			m:        1,
			nums2:    []int{1},
			n:        1,
			expected: []int{1, 2},
		},
		{
			name:     "Negative numbers",
			nums1:    []int{-3, -1, 0, 0, 0},
			m:        2,
			nums2:    []int{-2, 0, 2},
			n:        3,
			expected: []int{-3, -2, -1, 0, 2},
		},
		{
			name:     "Large numbers",
			nums1:    []int{100, 200, 300, 0, 0},
			m:        3,
			nums2:    []int{150, 250},
			n:        2,
			expected: []int{100, 150, 200, 250, 300},
		},
		{
			name:     "Already merged",
			nums1:    []int{1, 2, 3, 4, 5, 0, 0, 0},
			m:        5,
			nums2:    []int{6, 7, 8},
			n:        3,
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name:     "Reverse order merge",
			nums1:    []int{8, 9, 10, 0, 0, 0},
			m:        3,
			nums2:    []int{5, 6, 7},
			n:        3,
			expected: []int{5, 6, 7, 8, 9, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of nums1 to avoid modifying the test case
			nums1Copy := make([]int, len(tt.nums1))
			copy(nums1Copy, tt.nums1)

			MergeSortedArray(nums1Copy, tt.m, tt.nums2, tt.n)

			if !slicesEqual(nums1Copy, tt.expected) {
				t.Errorf("MergeSortedArray(%v, %d, %v, %d) = %v, expected %v",
					tt.nums1, tt.m, tt.nums2, tt.n, nums1Copy, tt.expected)
			}
		})
	}
}

func TestAllMergeSortedArrayImplementations(t *testing.T) {
	testCases := []struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
	}{
		{
			name:  "Example 1",
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
		},
		{
			name:  "Empty nums2",
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{},
			n:     0,
		},
		{
			name:  "Empty nums1",
			nums1: []int{0, 0, 0},
			m:     0,
			nums2: []int{1, 2, 3},
			n:     3,
		},
		{
			name:  "Interleaved",
			nums1: []int{1, 3, 5, 0, 0, 0},
			m:     3,
			nums2: []int{2, 4, 6},
			n:     3,
		},
	}

	implementations := []struct {
		name string
		fn   func([]int, int, []int, int)
	}{
		{"mergeSortedArray", mergeSortedArray},
		{"mergeSortedArrayTwoPointers", mergeSortedArrayTwoPointers},
		{"mergeSortedArraySimple", mergeSortedArraySimple},
		{"mergeSortedArrayInPlace", mergeSortedArrayInPlace},
		{"mergeSortedArrayRecursive", mergeSortedArrayRecursive},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Get expected result using the default implementation
			nums1Expected := make([]int, len(tc.nums1))
			copy(nums1Expected, tc.nums1)
			MergeSortedArray(nums1Expected, tc.m, tc.nums2, tc.n)

			for _, impl := range implementations {
				t.Run(impl.name, func(t *testing.T) {
					// Create a fresh copy for each implementation
					nums1Copy := make([]int, len(tc.nums1))
					copy(nums1Copy, tc.nums1)

					impl.fn(nums1Copy, tc.m, tc.nums2, tc.n)

					if !slicesEqual(nums1Copy, nums1Expected) {
						t.Errorf("%s(%v, %d, %v, %d) = %v, expected %v",
							impl.name, tc.nums1, tc.m, tc.nums2, tc.n, nums1Copy, nums1Expected)
					}
				})
			}
		})
	}
}

func TestMergeSortedArrayEdgeCases(t *testing.T) {
	t.Run("Both arrays empty", func(t *testing.T) {
		nums1 := []int{}
		nums2 := []int{}
		MergeSortedArray(nums1, 0, nums2, 0)
		if len(nums1) != 0 {
			t.Error("Empty arrays should remain empty")
		}
	})

	t.Run("Single element each", func(t *testing.T) {
		nums1 := []int{2, 0}
		nums2 := []int{1}
		MergeSortedArray(nums1, 1, nums2, 1)
		expected := []int{1, 2}
		if !slicesEqual(nums1, expected) {
			t.Errorf("Expected %v, got %v", expected, nums1)
		}
	})

	t.Run("nums1 already full", func(t *testing.T) {
		nums1 := []int{1, 2, 3}
		nums2 := []int{}
		MergeSortedArray(nums1, 3, nums2, 0)
		expected := []int{1, 2, 3}
		if !slicesEqual(nums1, expected) {
			t.Errorf("Expected %v, got %v", expected, nums1)
		}
	})

	t.Run("Large arrays", func(t *testing.T) {
		// Test with arrays at maximum size (200 elements)
		m, n := 100, 100
		nums1 := make([]int, m+n)
		nums2 := make([]int, n)

		// Fill nums1 with even numbers
		for i := 0; i < m; i++ {
			nums1[i] = i * 2
		}

		// Fill nums2 with odd numbers
		for i := 0; i < n; i++ {
			nums2[i] = i*2 + 1
		}

		MergeSortedArray(nums1, m, nums2, n)

		// Verify the result is sorted
		for i := 1; i < len(nums1); i++ {
			if nums1[i] < nums1[i-1] {
				t.Errorf("Array not sorted at index %d: %d < %d", i, nums1[i], nums1[i-1])
			}
		}

		// Verify all numbers 0-199 are present
		expectedSum := (199 * 200) / 2 // Sum of 0..199
		actualSum := 0
		for _, num := range nums1 {
			actualSum += num
		}
		if actualSum != expectedSum {
			t.Errorf("Sum mismatch: expected %d, got %d", expectedSum, actualSum)
		}
	})

	t.Run("All same values", func(t *testing.T) {
		nums1 := []int{5, 5, 5, 0, 0, 0}
		nums2 := []int{5, 5, 5}
		MergeSortedArray(nums1, 3, nums2, 3)
		expected := []int{5, 5, 5, 5, 5, 5}
		if !slicesEqual(nums1, expected) {
			t.Errorf("Expected %v, got %v", expected, nums1)
		}
	})
}

func TestMergeSortedArrayProperties(t *testing.T) {
	// Property-based testing
	implementations := []struct {
		name string
		fn   func([]int, int, []int, int)
	}{
		{"mergeSortedArray", mergeSortedArray},
		{"mergeSortedArrayTwoPointers", mergeSortedArrayTwoPointers},
		{"mergeSortedArrayInPlace", mergeSortedArrayInPlace},
	}

	for _, impl := range implementations {
		t.Run(impl.name+"_properties", func(t *testing.T) {
			// Test multiple random cases
			testCases := []struct {
				m, n int
			}{
				{0, 0},
				{1, 0},
				{0, 1},
				{3, 3},
				{5, 2},
				{2, 5},
				{10, 10},
			}

			for _, tc := range testCases {
				t.Run(fmt.Sprintf("m=%d,n=%d", tc.m, tc.n), func(t *testing.T) {
					// Generate sorted arrays
					nums1 := make([]int, tc.m+tc.n)
					nums2 := make([]int, tc.n)

					// Fill with sorted random data
					for i := 0; i < tc.m; i++ {
						nums1[i] = i * 2 // Even numbers
					}
					for i := 0; i < tc.n; i++ {
						nums2[i] = i*2 + 1 // Odd numbers
					}

					// Make a copy for the implementation
					nums1Copy := make([]int, len(nums1))
					copy(nums1Copy, nums1)

					impl.fn(nums1Copy, tc.m, nums2, tc.n)

					// Property 1: Result should be sorted
					for i := 1; i < len(nums1Copy); i++ {
						if nums1Copy[i] < nums1Copy[i-1] {
							t.Errorf("Result not sorted at index %d: %d < %d",
								i, nums1Copy[i], nums1Copy[i-1])
						}
					}

					// Property 2: Result should contain all elements from both arrays
					elementCount := make(map[int]int)
					for i := 0; i < tc.m; i++ {
						elementCount[nums1[i]]++
					}
					for i := 0; i < tc.n; i++ {
						elementCount[nums2[i]]++
					}
					for i := 0; i < len(nums1Copy); i++ {
						elementCount[nums1Copy[i]]--
					}

					for num, count := range elementCount {
						if count != 0 {
							t.Errorf("Element %d count mismatch: expected 0, got %d", num, count)
						}
					}

					// Property 3: Size should be m + n
					if len(nums1Copy) != tc.m+tc.n {
						t.Errorf("Size mismatch: expected %d, got %d",
							tc.m+tc.n, len(nums1Copy))
					}
				})
			}
		})
	}
}

func BenchmarkMergeSortedArray(b *testing.B) {
	// Test cases of different sizes
	testCases := []struct {
		name string
		m, n int
	}{
		{"Small", 5, 5},
		{"Medium", 50, 50},
		{"Large", 100, 100},
		{"Uneven", 30, 70},
	}

	implementations := []struct {
		name string
		fn   func([]int, int, []int, int)
	}{
		{"mergeSortedArray", mergeSortedArray},
		{"mergeSortedArrayTwoPointers", mergeSortedArrayTwoPointers},
		{"mergeSortedArraySimple", mergeSortedArraySimple},
		{"mergeSortedArrayInPlace", mergeSortedArrayInPlace},
		{"mergeSortedArrayRecursive", mergeSortedArrayRecursive},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			// Prepare test data
			nums1 := make([]int, tc.m+tc.n)
			nums2 := make([]int, tc.n)

			// Fill with sorted data
			for i := 0; i < tc.m; i++ {
				nums1[i] = i * 2
			}
			for i := 0; i < tc.n; i++ {
				nums2[i] = i*2 + 1
			}

			for _, impl := range implementations {
				b.Run(impl.name, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						// Create fresh copies for each iteration
						nums1Copy := make([]int, len(nums1))
						copy(nums1Copy, nums1)
						nums2Copy := make([]int, len(nums2))
						copy(nums2Copy, nums2)

						impl.fn(nums1Copy, tc.m, nums2Copy, tc.n)
					}
				})
			}
		})
	}
}

func BenchmarkMergeSortedArrayWorstCase(b *testing.B) {
	// Worst case: all elements need to be shifted
	m, n := 100, 100
	nums1 := make([]int, m+n)
	nums2 := make([]int, n)

	// nums1 has large numbers, nums2 has small numbers
	for i := 0; i < m; i++ {
		nums1[i] = i + 100
	}
	for i := 0; i < n; i++ {
		nums2[i] = i
	}

	b.ResetTimer()

	b.Run("mergeSortedArray", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			nums1Copy := make([]int, len(nums1))
			copy(nums1Copy, nums1)
			nums2Copy := make([]int, len(nums2))
			copy(nums2Copy, nums2)

			mergeSortedArray(nums1Copy, m, nums2Copy, n)
		}
	})

	b.Run("mergeSortedArrayTwoPointers", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			nums1Copy := make([]int, len(nums1))
			copy(nums1Copy, nums1)
			nums2Copy := make([]int, len(nums2))
			copy(nums2Copy, nums2)

			mergeSortedArrayTwoPointers(nums1Copy, m, nums2Copy, n)
		}
	})
}

// Helper function to compare slices
func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}