package leetcode

import (
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Example 1",
			s:        "(()",
			expected: 2,
		},
		{
			name:     "Example 2",
			s:        ")()())",
			expected: 4,
		},
		{
			name:     "Empty string",
			s:        "",
			expected: 0,
		},
		{
			name:     "Single character",
			s:        "(",
			expected: 0,
		},
		{
			name:     "All valid",
			s:        "()()()",
			expected: 6,
		},
		{
			name:     "Nested valid",
			s:        "((()))",
			expected: 6,
		},
		{
			name:     "Mixed",
			s:        "()(()",
			expected: 2,
		},
		{
			name:     "Complex case",
			s:        ")(()())()())",
			expected: 10, // Corrected: "(()())()()" is valid and length 10
		},
		{
			name:     "No valid",
			s:        "(((",
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Test stack approach
			result := LongestValidParentheses(test.s)
			if result != test.expected {
				t.Errorf("LongestValidParentheses(%q) = %d, expected %d", test.s, result, test.expected)
			}

			// Test DP approach
			result = LongestValidParenthesesDP(test.s)
			if result != test.expected {
				t.Errorf("LongestValidParenthesesDP(%q) = %d, expected %d", test.s, result, test.expected)
			}

			// Test two-pass approach
			result = LongestValidParenthesesTwoPass(test.s)
			if result != test.expected {
				t.Errorf("LongestValidParenthesesTwoPass(%q) = %d, expected %d", test.s, result, test.expected)
			}
		})
	}
}

func BenchmarkLongestValidParentheses(b *testing.B) {
	s := "((()())()(()))((()())()(()))((()())()(()))"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LongestValidParentheses(s)
	}
}

func BenchmarkLongestValidParenthesesDP(b *testing.B) {
	s := "((()())()(()))((()())()(()))((()())()(()))"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LongestValidParenthesesDP(s)
	}
}

func BenchmarkLongestValidParenthesesTwoPass(b *testing.B) {
	s := "((()())()(()))((()())()(()))((()())()(()))"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LongestValidParenthesesTwoPass(s)
	}
}