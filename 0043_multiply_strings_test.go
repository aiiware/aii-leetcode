package leetcode

import (
    "testing"
)

func TestMultiply(t *testing.T) {
    tests := []struct {
        num1     string
        num2     string
        expected string
    }{
        {"2", "3", "6"},
        {"123", "456", "56088"},
        {"0", "123", "0"},
        {"123", "0", "0"},
        {"999", "999", "998001"},
        {"1", "1", "1"},
        {"10", "10", "100"},
        {"99", "99", "9801"},
        {"123456789", "987654321", "121932631112635269"},
        {"0", "0", "0"},
        {"1000000000", "1000000000", "1000000000000000000"},
    }
    
    for i, test := range tests {
        // Test main implementation
        result := Multiply(test.num1, test.num2)
        if result != test.expected {
            t.Errorf("Test %d failed: %s * %s = %s, expected %s", 
                i, test.num1, test.num2, result, test.expected)
        }
        
        // Test simple implementation (only for small numbers)
        if len(test.num1) <= 9 && len(test.num2) <= 9 {
            result = MultiplySimple(test.num1, test.num2)
            if result != test.expected {
                t.Errorf("Test %d (simple) failed: %s * %s = %s, expected %s", 
                    i, test.num1, test.num2, result, test.expected)
            }
        }
    }
}

func BenchmarkMultiply(b *testing.B) {
    num1 := "12345678901234567890"
    num2 := "98765432109876543210"
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        Multiply(num1, num2)
    }
}

func BenchmarkMultiplySimple(b *testing.B) {
    num1 := "12345"
    num2 := "67890"
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        MultiplySimple(num1, num2)
    }
}