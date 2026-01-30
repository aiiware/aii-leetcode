package strings

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Example 1",
			input:    "the sky is blue",
			expected: "blue is sky the",
		},
		{
			name:     "Example 2",
			input:    "  hello world  ",
			expected: "world hello",
		},
		{
			name:     "Example 3",
			input:    "a good   example",
			expected: "example good a",
		},
		{
			name:     "Single word",
			input:    "hello",
			expected: "hello",
		},
		{
			name:     "Multiple spaces between words",
			input:    "  Bob    Loves  Alice   ",
			expected: "Alice Loves Bob",
		},
		{
			name:     "Leading and trailing spaces",
			input:    "   leading and trailing spaces   ",
			expected: "spaces trailing and leading",
		},
		{
			name:     "All spaces",
			input:    "     ",
			expected: "",
		},
		{
			name:     "Mixed case",
			input:    "Hello World from Go",
			expected: "Go from World Hello",
		},
		{
			name:     "With numbers",
			input:    "123 abc 456 def",
			expected: "def 456 abc 123",
		},
		{
			name:     "Special characters",
			input:    "hello@world.com is email",
			expected: "email is hello@world.com",
		},
		{
			name:     "Long sentence",
			input:    "The quick brown fox jumps over the lazy dog",
			expected: "dog lazy the over jumps fox brown quick The",
		},
		{
			name:     "Tab characters",
			input:    "hello\tworld",
			expected: "world hello",
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Single character words",
			input:    "a b c d e",
			expected: "e d c b a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReverseWords(tt.input)
			if result != tt.expected {
				t.Errorf("ReverseWords(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestReverseWordsSimple(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Example 1",
			input:    "the sky is blue",
			expected: "blue is sky the",
		},
		{
			name:     "Example 2",
			input:    "  hello world  ",
			expected: "world hello",
		},
		{
			name:     "Example 3",
			input:    "a good   example",
			expected: "example good a",
		},
		{
			name:     "Single word",
			input:    "hello",
			expected: "hello",
		},
		{
			name:     "Multiple spaces",
			input:    "  Bob    Loves  Alice   ",
			expected: "Alice Loves Bob",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReverseWordsSimple(tt.input)
			if result != tt.expected {
				t.Errorf("ReverseWordsSimple(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

func BenchmarkReverseWords(b *testing.B) {
	testCases := []string{
		"the sky is blue",
		"  hello world  ",
		"a good   example",
		"The quick brown fox jumps over the lazy dog",
		"  Bob    Loves  Alice   ",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, s := range testCases {
			ReverseWords(s)
		}
	}
}

func BenchmarkReverseWordsSimple(b *testing.B) {
	testCases := []string{
		"the sky is blue",
		"  hello world  ",
		"a good   example",
		"The quick brown fox jumps over the lazy dog",
		"  Bob    Loves  Alice   ",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, s := range testCases {
			ReverseWordsSimple(s)
		}
	}
}

func TestReverseWordsEdgeCases(t *testing.T) {
	t.Run("All spaces", func(t *testing.T) {
		result := ReverseWords("     ")
		if result != "" {
			t.Errorf("ReverseWords(all spaces) = %q, expected empty string", result)
		}
	})

	t.Run("Single space", func(t *testing.T) {
		result := ReverseWords(" ")
		if result != "" {
			t.Errorf("ReverseWords(single space) = %q, expected empty string", result)
		}
	})

	t.Run("Mixed whitespace", func(t *testing.T) {
		result := ReverseWords("  \t\n  hello  \t\n  world  \t\n  ")
		// Note: Our implementation only handles spaces, not all whitespace
		// This test shows the limitation
		if result != "world hello" {
			t.Errorf("ReverseWords(mixed whitespace) = %q, expected 'world hello'", result)
		}
	})

	t.Run("Very long string", func(t *testing.T) {
		// Create a long string with many words
		longStr := ""
		for i := 0; i < 1000; i++ {
			if i > 0 {
				longStr += " "
			}
			longStr += "word"
		}
		
		result := ReverseWords(longStr)
		// The result should be the same since all words are identical
		if result != longStr {
			t.Errorf("ReverseWords(long string) length mismatch: got %d, expected %d", len(result), len(longStr))
		}
	})
}

func TestReverseWordsBothImplementationsMatch(t *testing.T) {
	testCases := []string{
		"the sky is blue",
		"  hello world  ",
		"a good   example",
		"hello",
		"",
		" ",
		"     ",
		"a b c d e",
		"  test  test  test  ",
	}

	for _, s := range testCases {
		result1 := ReverseWords(s)
		result2 := ReverseWordsSimple(s)
		
		if result1 != result2 {
			t.Errorf("Implementations differ for %q: ReverseWords=%q, ReverseWordsSimple=%q", s, result1, result2)
		}
	}
}