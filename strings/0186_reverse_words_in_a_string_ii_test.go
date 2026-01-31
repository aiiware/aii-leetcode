package strings

import (
	"testing"
)

func TestReverseWordsII(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:     "Example 1",
			input:    []byte("the sky is blue"),
			expected: []byte("blue is sky the"),
		},
		{
			name:     "Example 2",
			input:    []byte("a"),
			expected: []byte("a"),
		},
		{
			name:     "Two words",
			input:    []byte("hello world"),
			expected: []byte("world hello"),
		},
		{
			name:     "Three words",
			input:    []byte("one two three"),
			expected: []byte("three two one"),
		},
		{
			name:     "Multiple words with single letters",
			input:    []byte("a b c d e"),
			expected: []byte("e d c b a"),
		},
		{
			name:     "Mixed case",
			input:    []byte("Hello World from Go"),
			expected: []byte("Go from World Hello"),
		},
		{
			name:     "With numbers",
			input:    []byte("123 abc 456 def"),
			expected: []byte("def 456 abc 123"),
		},
		{
			name:     "Special characters in words",
			input:    []byte("hello@world.com is email"),
			expected: []byte("email is hello@world.com"),
		},
		{
			name:     "Long sentence",
			input:    []byte("The quick brown fox jumps over the lazy dog"),
			expected: []byte("dog lazy the over jumps fox brown quick The"),
		},
		{
			name:     "Single word long",
			input:    []byte("supercalifragilisticexpialidocious"),
			expected: []byte("supercalifragilisticexpialidocious"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of input since we'll modify it in-place
			input := make([]byte, len(tt.input))
			copy(input, tt.input)

			ReverseWordsII(input)

			// Compare results
			if string(input) != string(tt.expected) {
				t.Errorf("ReverseWordsII(%q) = %q, expected %q", tt.input, input, tt.expected)
			}
		})
	}
}

func TestReverseWordsIIAlternative(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:     "Example 1",
			input:    []byte("the sky is blue"),
			expected: []byte("blue is sky the"),
		},
		{
			name:     "Example 2",
			input:    []byte("a"),
			expected: []byte("a"),
		},
		{
			name:     "Two words",
			input:    []byte("hello world"),
			expected: []byte("world hello"),
		},
		{
			name:     "Three words",
			input:    []byte("one two three"),
			expected: []byte("three two one"),
		},
		{
			name:     "Multiple words with single letters",
			input:    []byte("a b c d e"),
			expected: []byte("e d c b a"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of input since we'll modify it in-place
			input := make([]byte, len(tt.input))
			copy(input, tt.input)

			ReverseWordsIIAlternative(input)

			// Compare results
			if string(input) != string(tt.expected) {
				t.Errorf("ReverseWordsIIAlternative(%q) = %q, expected %q", tt.input, input, tt.expected)
			}
		})
	}
}

func TestReverseWordsIIString(t *testing.T) {
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
			input:    "a",
			expected: "a",
		},
		{
			name:     "Two words",
			input:    "hello world",
			expected: "world hello",
		},
		{
			name:     "Three words",
			input:    "one two three",
			expected: "three two one",
		},
		{
			name:     "Multiple words with single letters",
			input:    "a b c d e",
			expected: "e d c b a",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReverseWordsIIString(tt.input)

			// Compare results
			if result != tt.expected {
				t.Errorf("ReverseWordsIIString(%q) = %q, expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestReverseWordsIIBothImplementationsMatch(t *testing.T) {
	testCases := []string{
		"the sky is blue",
		"a",
		"hello world",
		"one two three",
		"a b c d e",
		"Hello World from Go",
		"123 abc 456 def",
		"",
		"single",
	}

	for _, s := range testCases {
		if s == "" {
			continue // Skip empty string as it's not valid input for this problem
		}

		// Test first implementation
		input1 := []byte(s)
		ReverseWordsII(input1)
		result1 := string(input1)

		// Test second implementation
		input2 := []byte(s)
		ReverseWordsIIAlternative(input2)
		result2 := string(input2)

		// Test helper function
		result3 := ReverseWordsIIString(s)

		// All three should match
		if result1 != result2 || result1 != result3 {
			t.Errorf("Implementations differ for %q: ReverseWordsII=%q, ReverseWordsIIAlternative=%q, ReverseWordsIIString=%q",
				s, result1, result2, result3)
		}
	}
}

func TestReverseRange(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		start    int
		end      int
		expected []byte
	}{
		{
			name:     "Reverse entire string",
			input:    []byte("hello"),
			start:    0,
			end:      4,
			expected: []byte("olleh"),
		},
		{
			name:     "Reverse middle portion",
			input:    []byte("abcdef"),
			start:    1,
			end:      4,
			expected: []byte("aedcbf"),
		},
		{
			name:     "Reverse single character",
			input:    []byte("hello"),
			start:    2,
			end:      2,
			expected: []byte("hello"),
		},
		{
			name:     "Reverse two characters",
			input:    []byte("hello"),
			start:    1,
			end:      2,
			expected: []byte("hlelo"),
		},
		{
			name:     "Reverse empty range",
			input:    []byte("hello"),
			start:    3,
			end:      2,
			expected: []byte("hello"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of input since we'll modify it in-place
			input := make([]byte, len(tt.input))
			copy(input, tt.input)

			reverseRange(input, tt.start, tt.end)

			// Compare results
			if string(input) != string(tt.expected) {
				t.Errorf("reverseRange(%q, %d, %d) = %q, expected %q",
					tt.input, tt.start, tt.end, input, tt.expected)
			}
		})
	}
}

func BenchmarkReverseWordsII(b *testing.B) {
	testCases := []string{
		"the sky is blue",
		"hello world",
		"one two three four five six seven eight nine ten",
		"The quick brown fox jumps over the lazy dog",
		"a b c d e f g h i j k l m n o p q r s t u v w x y z",
	}

	// Convert to byte slices for benchmarking
	byteCases := make([][]byte, len(testCases))
	for i, s := range testCases {
		byteCases[i] = []byte(s)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, s := range byteCases {
			// Make a copy since we modify in-place
			input := make([]byte, len(s))
			copy(input, s)
			ReverseWordsII(input)
		}
	}
}

func BenchmarkReverseWordsIIAlternative(b *testing.B) {
	testCases := []string{
		"the sky is blue",
		"hello world",
		"one two three four five six seven eight nine ten",
		"The quick brown fox jumps over the lazy dog",
		"a b c d e f g h i j k l m n o p q r s t u v w x y z",
	}

	// Convert to byte slices for benchmarking
	byteCases := make([][]byte, len(testCases))
	for i, s := range testCases {
		byteCases[i] = []byte(s)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, s := range byteCases {
			// Make a copy since we modify in-place
			input := make([]byte, len(s))
			copy(input, s)
			ReverseWordsIIAlternative(input)
		}
	}
}

func TestReverseWordsIIEdgeCases(t *testing.T) {
	t.Run("Already reversed", func(t *testing.T) {
		input := []byte("blue is sky the")
		expected := []byte("the sky is blue")
		ReverseWordsII(input)
		if string(input) != string(expected) {
			t.Errorf("ReverseWordsII(already reversed) = %q, expected %q", input, expected)
		}
	})

	t.Run("Palindrome words", func(t *testing.T) {
		input := []byte("racecar madam level")
		expected := []byte("level madam racecar")
		ReverseWordsII(input)
		if string(input) != string(expected) {
			t.Errorf("ReverseWordsII(palindrome words) = %q, expected %q", input, expected)
		}
	})

	t.Run("Very long word", func(t *testing.T) {
		// Create a very long word
		longWord := ""
		for i := 0; i < 1000; i++ {
			longWord += "a"
		}
		input := []byte(longWord + " b c")
		expected := []byte("c b " + longWord)
		ReverseWordsII(input)
		if string(input) != string(expected) {
			t.Errorf("ReverseWordsII(very long word) length mismatch")
		}
	})

	t.Run("Many short words", func(t *testing.T) {
		// Create many short words
		manyWords := ""
		for i := 0; i < 100; i++ {
			if i > 0 {
				manyWords += " "
			}
			manyWords += "w"
		}
		input := []byte(manyWords)
		// Reversing many identical words should give the same result
		ReverseWordsII(input)
		if string(input) != manyWords {
			t.Errorf("ReverseWordsII(many short words) = %q, expected %q", input, manyWords)
		}
	})
}