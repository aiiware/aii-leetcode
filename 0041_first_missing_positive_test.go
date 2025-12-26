package leetcode

import (
    "testing"
)

func TestFirstMissingPositive(t *testing.T) {
    tests := []struct {
        nums     []int
        expected int
    }{
        {[]int{1, 2, 0}, 3},
        {[]int{3, 4, -1, 1}, 2},
        {[]int{7, 8, 9, 11, 12}, 1},
        {[]int{1}, 2},
        {[]int{}, 1},
        {[]int{1, 1}, 2},
        {[]int{2, 2}, 1},
        {[]int{1, 2, 3, 4, 5}, 6},
        {[]int{1, 1000}, 2},
        {[]int{-5, -3, -1, 0}, 1},
    }
    
    for i, test := range tests {
        // Test in-place swap approach
        numsCopy := make([]int, len(test.nums))
        copy(numsCopy, test.nums)
        result := FirstMissingPositive(numsCopy)
        
        if result != test.expected {
            t.Errorf("Test %d (in-place) failed: nums=%v, got %d, expected %d", 
                i, test.nums, result, test.expected)
        }
        
        // Test hash approach
        numsCopy = make([]int, len(test.nums))
        copy(numsCopy, test.nums)
        result = FirstMissingPositiveHash(numsCopy)
        
        if result != test.expected {
            t.Errorf("Test %d (hash) failed: nums=%v, got %d, expected %d", 
                i, test.nums, result, test.expected)
        }
    }
}

func BenchmarkFirstMissingPositive(b *testing.B) {
    nums := []int{3, 4, -1, 1, 2, 5, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        numsCopy := make([]int, len(nums))
        copy(numsCopy, nums)
        FirstMissingPositive(numsCopy)
    }
}

func BenchmarkFirstMissingPositiveHash(b *testing.B) {
    nums := []int{3, 4, -1, 1, 2, 5, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        numsCopy := make([]int, len(nums))
        copy(numsCopy, nums)
        FirstMissingPositiveHash(numsCopy)
    }
}