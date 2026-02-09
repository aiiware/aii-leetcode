package arrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestPalindromicSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "Example 1",
			s:    "babad",
			want: "bab", // could also be "aba"
		},
		{
			name: "Example 2",
			s:    "cbbd",
			want: "bb",
		},
		{
			name: "Example 3",
			s:    "a",
			want: "a",
		},
		{
			name: "Empty string",
			s:    "",
			want: "",
		},
		{
			name: "All same characters",
			s:    "aaaa",
			want: "aaaa",
		},
		{
			name: "Palindrome at beginning",
			s:    "racecarb",
			want: "racecar",
		},
		{
			name: "Palindrome at end",
			s:    "baracecar",
			want: "racecar",
		},
		{
			name: "Multiple palindromes",
			s:    "abacdfgdcaba",
			want: "aba", // first palindrome
		},
		{
			name: "Long palindrome",
			s:    "forgeeksskeegfor",
			want: "geeksskeeg",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LongestPalindromicSubstring(tt.s)
			// For "babad" test case, accept either "bab" or "aba"
			if tt.s == "babad" {
				if result != "bab" && result != "aba" {
					t.Errorf("LongestPalindromicSubstring(%q) = %q, want %q or \"aba\"", tt.s, result, tt.want)
				}
			} else {
				assert.Equal(t, tt.want, result)
			}
		})
	}
}

func BenchmarkLongestPalindromicSubstring(b *testing.B) {
	testCases := []string{
		"babad",
		"cbbd",
		"forgeeksskeegfor",
		"abacdfgdcaba",
		"racecarb",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, s := range testCases {
			LongestPalindromicSubstring(s)
		}
	}
}
