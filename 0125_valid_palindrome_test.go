package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected bool
	}{
		{
			name:     "Example 1",
			s:        "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "Example 2",
			s:        "race a car",
			expected: false,
		},
		{
			name:     "Example 3",
			s:        " ",
			expected: true,
		},
		{
			name:     "Empty string",
			s:        "",
			expected: true,
		},
		{
			name:     "Single character",
			s:        "a",
			expected: true,
		},
		{
			name:     "Simple palindrome",
			s:        "racecar",
			expected: true,
		},
		{
			name:     "Not a palindrome",
			s:        "hello",
			expected: false,
		},
		{
			name:     "Palindrome with numbers",
			s:        "A1b2b1a",
			expected: true,
		},
		{
			name:     "Not palindrome with numbers",
			s:        "A1b2c1a",
			expected: false,
		},
		{
			name:     "Palindrome with only non-alphanumeric",
			s:        "!!!@@@###",
			expected: true, // After removing non-alphanumeric, empty string is palindrome
		},
		{
			name:     "Mixed case palindrome",
			s:        "Aa",
			expected: true,
		},
		{
			name:     "Palindrome with punctuation",
			s:        "Able was I, ere I saw Elba!",
			expected: true,
		},
		{
			name:     "Palindrome with underscores",
			s:        "ab_a",
			expected: true, // "aba" after removing underscore
		},
		{
			name:     "Zero as character",
			s:        "0P",
			expected: false, // "0p" is not palindrome
		},
		{
			name:     "Long palindrome",
			s:        "a" + string(make([]rune, 100000)) + "a",
			expected: true,
		},
		{
			name:     "Palindrome with spaces only",
			s:        "   ",
			expected: true,
		},
		{
			name:     "Complex mixed",
			s:        "No 'x' in Nixon",
			expected: true, // "noxinnixon" is palindrome
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidPalindrome(tt.s)
			assert.Equal(t, tt.expected, result,
				"IsValidPalindrome(%q) = %v, expected %v",
				tt.s, result, tt.expected)
		})
	}
}

func TestIsValidPalindrome_EdgeCases(t *testing.T) {
	t.Run("Very long palindrome", func(t *testing.T) {
		// Create a very long palindrome
		half := make([]byte, 50000)
		for i := range half {
			half[i] = byte('a' + (i % 26))
		}

		// Create palindrome by concatenating half + reverse(half)
		s := string(half)
		// Reverse the half
		reversed := make([]byte, len(half))
		for i := range half {
			reversed[len(half)-1-i] = half[i]
		}
		s += string(reversed)

		result := IsValidPalindrome(s)
		assert.True(t, result, "Very long palindrome should return true")
	})

	t.Run("Very long non-palindrome", func(t *testing.T) {
		// Create a very long string that's not a palindrome
		s := make([]byte, 100000)
		for i := range s {
			s[i] = byte('a' + (i % 26))
		}
		// Make sure it's not a palindrome by changing the last character
		if s[len(s)-1] == s[0] {
			s[len(s)-1] = 'z'
		}

		result := IsValidPalindrome(string(s))
		assert.False(t, result, "Very long non-palindrome should return false")
	})

	t.Run("String with only non-alphanumeric at ends", func(t *testing.T) {
		s := "!!!abc!!!"
		result := IsValidPalindrome(s)
		assert.False(t, result, "!!!abc!!! should NOT be palindrome (abc is NOT palindrome)")
	})

	t.Run("String with mixed non-alphanumeric", func(t *testing.T) {
		s := "a!!!b!!!c!!!b!!!a"
		result := IsValidPalindrome(s)
		assert.True(t, result, "Should handle mixed non-alphanumeric characters")
	})
}

func BenchmarkIsValidPalindrome(b *testing.B) {
	// Create a long string for benchmarking
	s := make([]byte, 100000)
	for i := range s {
		s[i] = byte('a' + (i % 26))
	}
	// Make it a palindrome
	half := s[:len(s)/2]
	for i := range half {
		s[len(s)-1-i] = half[i]
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsValidPalindrome(string(s))
	}
}