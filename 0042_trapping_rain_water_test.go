package leetcode

import (
    "testing"
)

func TestTrap(t *testing.T) {
    tests := []struct {
        height   []int
        expected int
    }{
        {[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
        {[]int{4, 2, 0, 3, 2, 5}, 9},
        {[]int{}, 0},
        {[]int{1}, 0},
        {[]int{1, 2}, 0},
        {[]int{1, 0, 2}, 1},
        {[]int{2, 0, 2}, 2},
        {[]int{3, 0, 0, 2, 0, 4}, 10},
        {[]int{1, 2, 3, 4, 5}, 0},
        {[]int{5, 4, 3, 2, 1}, 0},
        {[]int{5, 4, 3, 2, 1, 2, 3, 4, 5}, 16},
    }
    
    for i, test := range tests {
        // Test two pointers approach
        result := Trap(test.height)
        if result != test.expected {
            t.Errorf("Test %d (two pointers) failed: height=%v, got %d, expected %d", 
                i, test.height, result, test.expected)
        }
        
        // Test DP approach
        result = TrapDP(test.height)
        if result != test.expected {
            t.Errorf("Test %d (DP) failed: height=%v, got %d, expected %d", 
                i, test.height, result, test.expected)
        }
        
        // Test stack approach
        result = TrapStack(test.height)
        if result != test.expected {
            t.Errorf("Test %d (stack) failed: height=%v, got %d, expected %d", 
                i, test.height, result, test.expected)
        }
    }
}

func BenchmarkTrap(b *testing.B) {
    height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1, 0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        Trap(height)
    }
}

func BenchmarkTrapDP(b *testing.B) {
    height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1, 0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        TrapDP(height)
    }
}

func BenchmarkTrapStack(b *testing.B) {
    height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1, 0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        TrapStack(height)
    }
}