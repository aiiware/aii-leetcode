package arrays

import (
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Example 1 - Hello World",
			input:    "Hello World",
			expected: 5,
		},
		{
			name:     "Example 2 - multiple spaces",
			input:    "   fly me   to   the moon  ",
			expected: 4,
		},
		{
			name:     "Example 3 - no trailing spaces",
			input:    "luffy is still joyboy",
			expected: 6,
		},
		{
			name:     "Single word",
			input:    "Hello",
			expected: 5,
		},
		{
			name:     "Single word with trailing spaces",
			input:    "Hello   ",
			expected: 5,
		},
		{
			name:     "Single word with leading spaces",
			input:    "   Hello",
			expected: 5,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "Only spaces",
			input:    "     ",
			expected: 0,
		},
		{
			name:     "Multiple spaces between words",
			input:    "a   b    c",
			expected: 1,
		},
		{
			name:     "Tab characters treated as part of word (LeetCode spec)",
			input:    "word\tword2",
			expected: 10, // "word\tword2" length - tabs are not spaces
		},
		{
			name:     "Mixed whitespace",
			input:    "word\nword2",
			expected: 10, // newline is not a space character
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LengthOfLastWord(tt.input)
			if result != tt.expected {
				t.Errorf("LengthOfLastWord(%q) = %d, expected %d", tt.input, result, tt.expected)
			}
		})
	}
}

func BenchmarkLengthOfLastWord(b *testing.B) {
	testString := "   fly me   to   the moon  "
	for i := 0; i < b.N; i++ {
		LengthOfLastWord(testString)
	}
}

func BenchmarkLengthOfLastWordLong(b *testing.B) {
	// Create a long string with many words
	testString := ""
	for i := 0; i < 1000; i++ {
		testString += "word "
	}
	testString += "lastword"

	for i := 0; i < b.N; i++ {
		LengthOfLastWord(testString)
	}
}
