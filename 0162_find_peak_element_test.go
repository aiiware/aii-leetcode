package leetcode

import (
	"fmt"
	"testing"
)

func TestFindPeakElement(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		expect int
	}{
		// LeetCode examples
		{
			name:   "Example 1: Peak at index 2",
			nums:   []int{1, 2, 3, 1},
			expect: 2,
		},
		{
			name:   "Example 2: Multiple peaks, return any",
			nums:   []int{1, 2, 1, 3, 5, 6, 4},
			expect: 5, // Could also be 1
		},

		// Edge cases
		{
			name:   "Single element",
			nums:   []int{1},
			expect: 0,
		},
		{
			name:   "Two elements ascending",
			nums:   []int{1, 2},
			expect: 1,
		},
		{
			name:   "Two elements descending",
			nums:   []int{2, 1},
			expect: 0,
		},
		{
			name:   "Three elements peak in middle",
			nums:   []int{1, 3, 2},
			expect: 1,
		},
		{
			name:   "Three elements peak at start",
			nums:   []int{3, 2, 1},
			expect: 0,
		},
		{
			name:   "Three elements peak at end",
			nums:   []int{1, 2, 3},
			expect: 2,
		},

		// Multiple peaks scenarios
		{
			name:   "Multiple peaks 1",
			nums:   []int{1, 2, 1, 2, 1},
			expect: 1, // Could also be 3
		},
		{
			name:   "Multiple peaks 2",
			nums:   []int{1, 3, 2, 4, 3, 5, 4},
			expect: 1, // Could also be 3 or 5
		},

		// Note: LeetCode problem states: nums[i] != nums[i+1] for all valid i
		// So we shouldn't have equal adjacent elements in valid test cases

		// Large arrays
		{
			name:   "Large array ascending",
			nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expect: 9,
		},
		{
			name:   "Large array descending",
			nums:   []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			expect: 0,
		},
		{
			name:   "Large array with peak in middle",
			nums:   []int{1, 2, 3, 4, 5, 6, 5, 4, 3, 2, 1},
			expect: 5,
		},

		// Negative numbers
		{
			name:   "Negative numbers peak",
			nums:   []int{-5, -4, -3, -2, -1},
			expect: 4,
		},
		{
			name:   "Mixed positive and negative",
			nums:   []int{-3, -2, -1, 0, 1, 2, 1, 0, -1},
			expect: 5,
		},

		// Zero values
		{
			name:   "Zeros with peak",
			nums:   []int{0, 1, 0},
			expect: 1,
		},

		// Boundary peaks
		{
			name: "Peak at first element",
			nums: []int{5, 4, 3, 2, 1},
			expect: 0,
		},
		{
			name:   "Peak at last element",
			nums:   []int{1, 2, 3, 4, 5},
			expect: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findPeakElement(tt.nums)
			
			// Verify that the result is a valid peak
			if !isValidPeak(tt.nums, result) {
				t.Errorf("findPeakElement(%v) = %d, which is not a valid peak", tt.nums, result)
			}
			
			// Note: We don't check exact index match because problem allows returning any peak
			// But we can log if it matches expected for reference
			if result != tt.expect {
				t.Logf("findPeakElement(%v) = %d (expected %d, but any valid peak is acceptable)", 
					tt.nums, result, tt.expect)
			}
		})
	}
}

// Helper function to check if an index is a valid peak
// According to LeetCode: A peak element is an element that is strictly greater than its neighbors.
func isValidPeak(nums []int, index int) bool {
	if len(nums) == 0 {
		return index == -1
	}
	
	if index < 0 || index >= len(nums) {
		return false
	}
	
	// Check left neighbor - must be strictly greater if neighbor exists
	leftValid := index == 0 || nums[index] > nums[index-1]
	
	// Check right neighbor - must be strictly greater if neighbor exists
	rightValid := index == len(nums)-1 || nums[index] > nums[index+1]
	
	return leftValid && rightValid
}

func TestFindPeakElementBinarySearch(t *testing.T) {
	// Test the specific binary search implementation
	tests := []struct {
		name   string
		nums   []int
	}{
		{
			name:   "Binary search: standard case",
			nums:   []int{1, 2, 3, 1},
		},
		{
			name:   "Binary search: multiple peaks",
			nums:   []int{1, 2, 1, 3, 5, 6, 4},
		},
		{
			name:   "Binary search: single element",
			nums:   []int{5},
		},
		{
			name:   "Binary search: two elements",
			nums:   []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findPeakElementBinarySearch(tt.nums)
			
			if !isValidPeak(tt.nums, result) {
				t.Errorf("findPeakElementBinarySearch(%v) = %d, which is not a valid peak", tt.nums, result)
			}
		})
	}
}

func TestFindPeakElementLinear(t *testing.T) {
	// Test the linear scan implementation
	tests := []struct {
		name   string
		nums   []int
	}{
		{
			name:   "Linear: ascending",
			nums:   []int{1, 2, 3, 4, 5},
		},
		{
			name:   "Linear: descending",
			nums:   []int{5, 4, 3, 2, 1},
		},
		{
			name:   "Linear: peak in middle",
			nums:   []int{1, 2, 3, 2, 1},
		},
		{
			name:   "Linear: empty array",
			nums:   []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findPeakElementLinear(tt.nums)
			
			if !isValidPeak(tt.nums, result) {
				t.Errorf("findPeakElementLinear(%v) = %d, which is not a valid peak", tt.nums, result)
			}
		})
	}
}

func TestFindPeakElementBinarySearchExplicit(t *testing.T) {
	// Test the explicit binary search implementation
	tests := []struct {
		name   string
		nums   []int
	}{
		{
			name:   "Explicit binary search: standard",
			nums:   []int{1, 2, 3, 1},
		},
		{
			name:   "Explicit binary search: with negatives",
			nums:   []int{-3, -2, -1, 0, 1, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findPeakElementBinarySearchExplicit(tt.nums)
			
			if !isValidPeak(tt.nums, result) {
				t.Errorf("findPeakElementBinarySearchExplicit(%v) = %d, which is not a valid peak", tt.nums, result)
			}
		})
	}
}

func TestFindPeakElementDivideConquer(t *testing.T) {
	// Test the divide and conquer implementation
	tests := []struct {
		name   string
		nums   []int
	}{
		{
			name:   "Divide conquer: simple peak",
			nums:   []int{1, 2, 3, 1},
		},
		{
			name:   "Divide conquer: multiple peaks",
			nums:   []int{1, 2, 1, 3, 5, 6, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findPeakElementDivideConquer(tt.nums)
			
			if !isValidPeak(tt.nums, result) {
				t.Errorf("findPeakElementDivideConquer(%v) = %d, which is not a valid peak", tt.nums, result)
			}
		})
	}
}

func TestFindPeakElementIterative(t *testing.T) {
	// Test the iterative binary search implementation
	tests := []struct {
		name   string
		nums   []int
	}{
		{
			name:   "Iterative: standard",
			nums:   []int{1, 2, 3, 1},
		},
		{
			name:   "Iterative: edge peaks",
			nums:   []int{5, 4, 3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findPeakElementIterative(tt.nums)
			
			if !isValidPeak(tt.nums, result) {
				t.Errorf("findPeakElementIterative(%v) = %d, which is not a valid peak", tt.nums, result)
			}
		})
	}
}

func TestGetValue(t *testing.T) {
	// Test the helper function
	nums := []int{1, 2, 3, 4, 5}
	
	tests := []struct {
		index  int
		expect int
	}{
		{-1, -1 << 31}, // negative infinity for out of bounds
		{0, 1},
		{2, 3},
		{4, 5},
		{5, -1 << 31}, // out of bounds
		{10, -1 << 31}, // out of bounds
	}
	
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Index %d", tt.index), func(t *testing.T) {
			result := getValue(nums, tt.index)
			if result != tt.expect {
				t.Errorf("getValue(%v, %d) = %d, want %d", nums, tt.index, result, tt.expect)
			}
		})
	}
}

func BenchmarkFindPeakElement(b *testing.B) {
	// Benchmark with various array sizes
	testCases := []struct {
		name string
		nums []int
	}{
		{"Small array", []int{1, 2, 3, 1}},
		{"Medium array", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
		{"Large array", make([]int, 1000)},
		{"Very large array", make([]int, 10000)},
	}
	
	// Initialize large arrays with a peak pattern
	for i := 0; i < 1000; i++ {
		testCases[2].nums[i] = i
		if i == 999 {
			testCases[2].nums[i] = 1000
		}
	}
	for i := 0; i < 10000; i++ {
		testCases[3].nums[i] = i
		if i == 9999 {
			testCases[3].nums[i] = 10000
		}
	}
	
	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				findPeakElement(tc.nums)
			}
		})
	}
}

func BenchmarkFindPeakElementBinarySearch(b *testing.B) {
	nums := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		nums[i] = i
	}
	nums[9999] = 10000 // Create a peak at the end
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findPeakElementBinarySearch(nums)
	}
}

func BenchmarkFindPeakElementLinear(b *testing.B) {
	nums := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = i
	}
	nums[999] = 1000
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findPeakElementLinear(nums)
	}
}

func BenchmarkFindPeakElementBinarySearchExplicit(b *testing.B) {
	nums := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		nums[i] = i
	}
	nums[9999] = 10000
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findPeakElementBinarySearchExplicit(nums)
	}
}

func BenchmarkFindPeakElementDivideConquer(b *testing.B) {
	nums := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		nums[i] = i
	}
	nums[9999] = 10000
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findPeakElementDivideConquer(nums)
	}
}

func BenchmarkFindPeakElementIterative(b *testing.B) {
	nums := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		nums[i] = i
	}
	nums[9999] = 10000
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findPeakElementIterative(nums)
	}
}

// Test that all implementations return valid peaks
func TestAllImplementationsReturnValidPeaks(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
	}{
		{"Standard case", []int{1, 2, 3, 1}},
		{"Multiple peaks", []int{1, 2, 1, 3, 5, 6, 4}},
		{"Single element", []int{5}},
		{"Two elements", []int{1, 2}},
		{"Large array", make([]int, 100)},
	}
	
	// Initialize large array
	for i := 0; i < 100; i++ {
		testCases[4].nums[i] = i
		if i == 99 {
			testCases[4].nums[i] = 100
		}
	}
	
	implementations := []struct {
		name string
		fn   func([]int) int
	}{
		{"findPeakElement", findPeakElement},
		{"findPeakElementBinarySearch", findPeakElementBinarySearch},
		{"findPeakElementLinear", findPeakElementLinear},
		{"findPeakElementBinarySearchExplicit", findPeakElementBinarySearchExplicit},
		{"findPeakElementDivideConquer", findPeakElementDivideConquer},
		{"findPeakElementIterative", findPeakElementIterative},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for _, impl := range implementations {
				t.Run(impl.name, func(t *testing.T) {
					result := impl.fn(tc.nums)
					if !isValidPeak(tc.nums, result) {
						t.Errorf("%s(%v) = %d, which is not a valid peak", impl.name, tc.nums, result)
					}
				})
			}
		})
	}
}

// Test empty array handling
func TestEmptyArray(t *testing.T) {
	empty := []int{}
	
	// All implementations should return -1 for empty array
	implementations := []struct {
		name string
		fn   func([]int) int
	}{
		{"findPeakElement", findPeakElement},
		{"findPeakElementBinarySearch", findPeakElementBinarySearch},
		{"findPeakElementLinear", findPeakElementLinear},
		{"findPeakElementBinarySearchExplicit", findPeakElementBinarySearchExplicit},
		{"findPeakElementDivideConquer", findPeakElementDivideConquer},
		{"findPeakElementIterative", findPeakElementIterative},
	}
	
	for _, impl := range implementations {
		t.Run(impl.name, func(t *testing.T) {
			result := impl.fn(empty)
			if result != -1 {
				t.Errorf("%s([]) = %d, want -1", impl.name, result)
			}
		})
	}
}