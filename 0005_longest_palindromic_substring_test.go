package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Example 1",
			input:    "babad",
			expected: "bab", // "aba" is also valid
		},
		{
			name:     "Example 2",
			input:    "cbbd",
			expected: "bb",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Single character",
			input:    "a",
			expected: "a",
		},
		{
			name:     "All same characters",
			input:    "aaaaa",
			expected: "aaaaa",
		},
		{
			name:     "Palindrome at beginning",
			input:    "racecarxyz",
			expected: "racecar",
		},
		{
			name:     "Palindrome at end",
			input:    "xyzracecar",
			expected: "racecar",
		},
		{
			name:     "Multiple palindromes",
			input:    "abacdfgdcaba",
			expected: "aba", // "aba" at beginning or "aba" at end
		},
		{
			name:     "Even length palindrome",
			input:    "abba",
			expected: "abba",
		},
		{
			name:     "Odd length palindrome",
			input:    "abcba",
			expected: "abcba",
		},
		{
			name:     "No palindrome longer than 1",
			input:    "abc",
			expected: "a", // any single character
		},
		{
			name:     "Long palindrome",
			input:    "forgeeksskeegfor",
			expected: "geeksskeeg",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LongestPalindrome(tt.input)
			// For some cases, there might be multiple valid answers
			if tt.name == "Example 1" {
				// "babad" can return "bab" or "aba"
				assert.True(t, result == "bab" || result == "aba",
					"Expected 'bab' or 'aba', got %s", result)
			} else if tt.name == "Multiple palindromes" {
				// "abacdfgdcaba" can return "aba" (beginning) or "aba" (end)
				assert.Equal(t, 3, len(result), "Palindrome should be length 3")
				assert.Equal(t, result[0], result[2], "Should be palindrome")
			} else if tt.name == "No palindrome longer than 1" {
				assert.Equal(t, 1, len(result), "Should return a single character")
			} else {
				assert.Equal(t, tt.expected, result, "Longest palindrome should match")
			}
		})
	}
}

func TestLongestPalindromeDP(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Example 1",
			input:    "babad",
			expected: "bab", // DP version returns "bab" or "aba"
		},
		{
			name:     "Example 2",
			input:    "cbbd",
			expected: "bb",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Single character",
			input:    "a",
			expected: "a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestPalindromeDP(tt.input)
			if tt.name == "Example 1" {
				// "babad" can return "bab" or "aba"
				assert.True(t, result == "bab" || result == "aba",
					"Expected 'bab' or 'aba', got %s", result)
			} else {
				assert.Equal(t, tt.expected, result, "DP version should match")
			}
		})
	}
}

func BenchmarkLongestPalindrome(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{
			name:  "Short string",
			input: "babad",
		},
		{
			name:  "Medium string",
			input: "forgeeksskeegfor",
		},
		{
			name:  "Long string",
			input: "a" + makeString('b', 100) + "a",
		},
		{
			name:  "All same characters",
			input: makeString('a', 200),
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LongestPalindrome(bm.input)
			}
		})
	}
}

func BenchmarkLongestPalindromeDP(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{
			name:  "Short string",
			input: "babad",
		},
		{
			name:  "Medium string",
			input: "forgeeksskeegfor",
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				longestPalindromeDP(bm.input)
			}
		})
	}
}

// Helper function to create a string of repeated characters
func makeString(ch rune, length int) string {
	result := make([]rune, length)
	for i := range result {
		result[i] = ch
	}
	return string(result)
}