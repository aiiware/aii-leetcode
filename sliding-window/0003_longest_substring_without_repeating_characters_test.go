package sliding_window

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Example 1",
			input:    "abcabcbb",
			expected: 3,
		},
		{
			name:     "Example 2",
			input:    "bbbbb",
			expected: 1,
		},
		{
			name:     "Example 3",
			input:    "pwwkew",
			expected: 3,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "Single character",
			input:    "a",
			expected: 1,
		},
		{
			name:     "Two different characters",
			input:    "ab",
			expected: 2,
		},
		{
			name:     "All same characters",
			input:    "aaaaaa",
			expected: 1,
		},
		{
			name:     "Complex case",
			input:    "abcabcbb",
			expected: 3,
		},
		{
			name:     "With spaces",
			input:    "a b c",
			expected: 3, // "a b" or "b c" - space appears twice so max is 3
		},
		{
			name:     "All unique with spaces",
			input:    "ab cde",
			expected: 6, // "ab cde" - all 6 chars are unique
		},
		{
			name:     "Unicode characters",
			input:    "αβγαβγ",
			expected: 3,
		},
		{
			name:     "Long repeated pattern",
			input:    "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
			expected: 26,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LengthOfLongestSubstring(tt.input)
			assert.Equal(t, tt.expected, result,
				"LengthOfLongestSubstring(%q) = %d, expected %d",
				tt.input, result, tt.expected)
		})
	}
}

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	// Create a large test string for benchmarking
	testString := "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LengthOfLongestSubstring(testString)
	}
}