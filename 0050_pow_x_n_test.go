package leetcode

import (
    "math"
    "testing"
)

func TestMyPow(t *testing.T) {
    tests := []struct {
        x        float64
        n        int
        expected float64
    }{
        {2.0, 10, 1024.0},
        {2.1, 3, 9.261},
        {2.0, -2, 0.25},
        {2.0, 0, 1.0},
        {0.0, 10, 0.0},
        {0.0, -2, math.Inf(1)}, // 1/0 = infinity
        {1.0, 100, 1.0},
        {-1.0, 2, 1.0},
        {-1.0, 3, -1.0},
        {0.00001, 2147483647, 0.0}, // Very small number to large power
        {1.00001, 100000, 2.7182682372}, // Approximates e, but not exactly
        {2.0, -2147483648, 0.0}, // Edge case: minimum int
    }
    
    tolerance := 1e-9
    
    for i, test := range tests {
        // Test recursive approach
        result := MyPow(test.x, test.n)
        if !almostEqual(result, test.expected, tolerance) {
            t.Errorf("Test %d (recursive) failed: %.5f^%d = %.10f, expected %.10f", 
                i, test.x, test.n, result, test.expected)
        }
        
        // Test iterative approach
        result = MyPowIterative(test.x, test.n)
        if !almostEqual(result, test.expected, tolerance) {
            t.Errorf("Test %d (iterative) failed: %.5f^%d = %.10f, expected %.10f", 
                i, test.x, test.n, result, test.expected)
        }
        
        // Test brute force (only for small n)
        if test.n >= -100 && test.n <= 100 {
            result = MyPowBruteForce(test.x, test.n)
            if !almostEqual(result, test.expected, tolerance) {
                t.Errorf("Test %d (brute force) failed: %.5f^%d = %.10f, expected %.10f", 
                    i, test.x, test.n, result, test.expected)
            }
        }
    }
}

// Helper function to compare floating point numbers
func almostEqual(a, b, tolerance float64) bool {
    if math.IsInf(a, 0) && math.IsInf(b, 0) {
        return math.Signbit(a) == math.Signbit(b)
    }
    if math.IsNaN(a) && math.IsNaN(b) {
        return true
    }
    return math.Abs(a-b) <= tolerance
}

func BenchmarkMyPow(b *testing.B) {
    x := 2.0
    n := 1000000
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        MyPow(x, n)
    }
}

func BenchmarkMyPowIterative(b *testing.B) {
    x := 2.0
    n := 1000000
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        MyPowIterative(x, n)
    }
}

func BenchmarkMyPowBruteForce(b *testing.B) {
    x := 2.0
    n := 1000 // Smaller n for brute force
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        MyPowBruteForce(x, n)
    }
}