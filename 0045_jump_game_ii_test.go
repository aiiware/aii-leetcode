package leetcode

import (
    "testing"
)

func TestJump(t *testing.T) {
    tests := []struct {
        nums     []int
        expected int
    }{
        {[]int{2, 3, 1, 1, 4}, 2},
        {[]int{2, 3, 0, 1, 4}, 2},
        {[]int{1, 2, 3}, 2},
        {[]int{1}, 0},
        {[]int{1, 1, 1, 1}, 3},
        {[]int{4, 1, 1, 3, 1, 1, 1}, 2},
        {[]int{1, 2, 1, 1, 1}, 3},
        {[]int{7, 0, 9, 6, 9, 6, 1, 7, 9, 0, 1, 2, 9, 0, 3}, 2},
        {[]int{0}, 0},
        {[]int{1, 1}, 1},
        {[]int{3, 2, 1}, 1},
        {[]int{2, 0, 2, 0, 1}, 2},
    }
    
    for i, test := range tests {
        // Test greedy approach
        result := Jump(test.nums)
        if result != test.expected {
            t.Errorf("Test %d (greedy) failed: nums=%v, got %d, expected %d", 
                i, test.nums, result, test.expected)
        }
        
        // Test DP approach
        result = JumpDP(test.nums)
        if result != test.expected {
            t.Errorf("Test %d (DP) failed: nums=%v, got %d, expected %d", 
                i, test.nums, result, test.expected)
        }
        
        // Test BFS approach
        result = JumpBFS(test.nums)
        if result != test.expected {
            t.Errorf("Test %d (BFS) failed: nums=%v, got %d, expected %d", 
                i, test.nums, result, test.expected)
        }
    }
}

func BenchmarkJump(b *testing.B) {
    nums := []int{2, 3, 1, 1, 4, 2, 1, 3, 2, 1, 4, 2, 3, 1, 1, 4}
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        Jump(nums)
    }
}

func BenchmarkJumpDP(b *testing.B) {
    nums := []int{2, 3, 1, 1, 4, 2, 1, 3, 2, 1, 4, 2, 3, 1, 1, 4}
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        JumpDP(nums)
    }
}

func BenchmarkJumpBFS(b *testing.B) {
    nums := []int{2, 3, 1, 1, 4, 2, 1, 3, 2, 1, 4, 2, 3, 1, 1, 4}
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        JumpBFS(nums)
    }
}