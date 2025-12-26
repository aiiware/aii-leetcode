package leetcode

import (
    "testing"
)

func TestIsMatchWildcard(t *testing.T) {
    tests := []struct {
        s        string
        p        string
        expected bool
    }{
        {"aa", "a", false},
        {"aa", "*", true},
        {"cb", "?a", false},
        {"adceb", "*a*b", true},
        {"acdcb", "a*c?b", false},
        {"", "", true},
        {"", "*", true},
        {"", "?", false},
        {"a", "", false},
        {"hello", "h*llo", true},
        {"hello", "h?llo", true},
        {"hello", "h*", true},
        {"hello", "*o", true},
        {"hello", "h*l?o", true},
        {"hello", "h*l??", true}, // Fixed: should match "hello"
        {"abcd", "a*b*c*d", true},
        {"abcd", "a*b*c*d*", true},
        {"mississippi", "m??*ss*?i*pi", false},
        {"mississippi", "m*iss*iss*", true},
    }
    
    for i, test := range tests {
        // Test DP approach
        result := IsMatchWildcard(test.s, test.p)
        if result != test.expected {
            t.Errorf("Test %d (DP) failed: s=%q, p=%q, got %v, expected %v", 
                i, test.s, test.p, result, test.expected)
        }
        
        // Test greedy approach
        result = IsMatchWildcardGreedy(test.s, test.p)
        if result != test.expected {
            t.Errorf("Test %d (greedy) failed: s=%q, p=%q, got %v, expected %v", 
                i, test.s, test.p, result, test.expected)
        }
        
        // Test recursive approach
        result = IsMatchWildcardRecursive(test.s, test.p)
        if result != test.expected {
            t.Errorf("Test %d (recursive) failed: s=%q, p=%q, got %v, expected %v", 
                i, test.s, test.p, result, test.expected)
        }
    }
}

func BenchmarkIsMatchWildcard(b *testing.B) {
    s := "mississippi"
    p := "m*iss*iss*"
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        IsMatchWildcard(s, p)
    }
}

func BenchmarkIsMatchWildcardGreedy(b *testing.B) {
    s := "mississippi"
    p := "m*iss*iss*"
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        IsMatchWildcardGreedy(s, p)
    }
}

func BenchmarkIsMatchWildcardRecursive(b *testing.B) {
    s := "mississippi"
    p := "m*iss*iss*"
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        IsMatchWildcardRecursive(s, p)
    }
}