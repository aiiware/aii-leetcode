package leetcode

import (
	"testing"
)

func TestMinCut(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Example 1",
			s:        "aab",
			expected: 1,
		},
		{
			name:     "Example 2",
			s:        "a",
			expected: 0,
		},
		{
			name:     "Example 3",
			s:        "ab",
			expected: 1,
		},
		{
			name:     "Empty string",
			s:        "",
			expected: 0,
		},
		{
			name:     "Single character",
			s:        "b",
			expected: 0,
		},
		{
			name:     "Palindrome string",
			s:        "aba",
			expected: 0,
		},
		{
			name:     "Long palindrome",
			s:        "racecar",
			expected: 0,
		},
		{
			name:     "All same characters",
			s:        "aaaa",
			expected: 0,
		},
		{
			name:     "No palindromes except single chars",
			s:        "abcde",
			expected: 4,
		},
		{
			name:     "Mixed case 1",
			s:        "aabb",
			expected: 1,
		},
		{
			name:     "Mixed case 2",
			s:        "leet",
			expected: 2,
		},
		{
			name:     "Complex case",
			s:        "ababbbabbababa",
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinCut(tt.s)
			if result != tt.expected {
				t.Errorf("MinCut(%q) = %d, expected %d", tt.s, result, tt.expected)
			}
		})
	}
}

func BenchmarkMinCut(b *testing.B) {
	testCases := []struct {
		name string
		s    string
	}{
		{"Short string", "aab"},
		{"Medium string", "aabbcc"},
		{"Palindrome", "racecarlevel"},
		{"All same chars", "aaaaaaaa"},
		{"No palindromes", "abcdefghijklmnop"},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MinCut(tc.s)
			}
		})
	}
}