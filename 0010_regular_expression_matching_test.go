package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMatch(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		p        string
		expected bool
	}{
		{
			name:     "Example 1",
			s:        "aa",
			p:        "a",
			expected: false,
		},
		{
			name:     "Example 2",
			s:        "aa",
			p:        "a*",
			expected: true,
		},
		{
			name:     "Example 3",
			s:        "ab",
			p:        ".*",
			expected: true,
		},
		{
			name:     "Example 4",
			s:        "aab",
			p:        "c*a*b",
			expected: true,
		},
		{
			name:     "Example 5",
			s:        "mississippi",
			p:        "mis*is*p*.",
			expected: false,
		},
		{
			name:     "Empty string and pattern",
			s:        "",
			p:        "",
			expected: true,
		},
		{
			name:     "Empty string with pattern",
			s:        "",
			p:        "a*",
			expected: true,
		},
		{
			name:     "Empty string with multiple stars",
			s:        "",
			p:        "a*b*c*",
			expected: true,
		},
		{
			name:     "Non-empty string with empty pattern",
			s:        "abc",
			p:        "",
			expected: false,
		},
		{
			name:     "Single character match",
			s:        "a",
			p:        "a",
			expected: true,
		},
		{
			name:     "Single character mismatch",
			s:        "a",
			p:        "b",
			expected: false,
		},
		{
			name:     "Dot matches any character",
			s:        "a",
			p:        ".",
			expected: true,
		},
		{
			name:     "Dot star matches everything",
			s:        "abcdef",
			p:        ".*",
			expected: true,
		},
		{
			name:     "Multiple stars",
			s:        "aaa",
			p:        "a*a",
			expected: true,
		},
		{
			name:     "Complex pattern 1",
			s:        "ab",
			p:        ".*c",
			expected: false,
		},
		{
			name:     "Complex pattern 2",
			s:        "aaa",
			p:        "ab*a*c*a",
			expected: true,
		},
		{
			name:     "Complex pattern 3",
			s:        "ab",
			p:        ".*..",
			expected: true,
		},
		{
			name:     "Zero or more with dot",
			s:        "abc",
			p:        ".*c",
			expected: true,
		},
		{
			name:     "Zero or more mismatch",
			s:        "abcd",
			p:        "d*",
			expected: false,
		},
		{
			name:     "Multiple dots and stars",
			s:        "xxy",
			p:        "x.*y",
			expected: true,
		},
		{
			name:     "Pattern longer than string",
			s:        "a",
			p:        "ab*",
			expected: true,
		},
		{
			name:     "Consecutive stars - invalid pattern",
			s:        "aaa",
			p:        "a**a",
			expected: false, // a**a is invalid pattern (star at position 2 has no preceding char)
		},
		{
			name:     "Star at beginning - invalid pattern",
			s:        "aa",
			p:        "*a",
			expected: false, // * at beginning is invalid (no preceding character)
		},
		{
			name:     "Complex real world example",
			s:        "mississippi",
			p:        "mis*is*ip*.",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsMatch(tt.s, tt.p)
			assert.Equal(t, tt.expected, result, "IsMatch(%q, %q) should be %v", tt.s, tt.p, tt.expected)
		})
	}
}

func TestIsMatchRecursive(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		p        string
		expected bool
	}{
		{
			name:     "Example 1",
			s:        "aa",
			p:        "a",
			expected: false,
		},
		{
			name:     "Example 2",
			s:        "aa",
			p:        "a*",
			expected: true,
		},
		{
			name:     "Example 3",
			s:        "ab",
			p:        ".*",
			expected: true,
		},
		{
			name:     "Complex pattern",
			s:        "aab",
			p:        "c*a*b",
			expected: true,
		},
		{
			name:     "Dot star matches everything",
			s:        "abcdef",
			p:        ".*",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsMatchRecursive(tt.s, tt.p)
			assert.Equal(t, tt.expected, result, "Recursive version should match")
		})
	}
}

func BenchmarkIsMatch(b *testing.B) {
	benchmarks := []struct {
		name string
		s    string
		p    string
	}{
		{
			name: "Simple match",
			s:    "aa",
			p:    "a*",
		},
		{
			name: "Dot star",
			s:    "abcdefghij",
			p:    ".*",
		},
		{
			name: "Complex pattern",
			s:    "aab",
			p:    "c*a*b",
		},
		{
			name: "Long string simple pattern",
			s:    makeRegexString('a', 100),
			p:    "a*a*a*",
		},
		{
			name: "Medium complexity",
			s:    "mississippi",
			p:    "mis*is*ip*.",
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsMatch(bm.s, bm.p)
			}
		})
	}
}

func BenchmarkIsMatchRecursive(b *testing.B) {
	benchmarks := []struct {
		name string
		s    string
		p    string
	}{
		{
			name: "Simple match",
			s:    "aa",
			p:    "a*",
		},
		{
			name: "Dot star",
			s:    "abcdef",
			p:    ".*",
		},
		{
			name: "Complex pattern",
			s:    "aab",
			p:    "c*a*b",
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsMatchRecursive(bm.s, bm.p)
			}
		})
	}
}

// Helper function to create a string of repeated characters
func makeRegexString(ch rune, length int) string {
	result := make([]rune, length)
	for i := range result {
		result[i] = ch
	}
	return string(result)
}