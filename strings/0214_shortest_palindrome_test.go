package strings

import "testing"

func TestShortestPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "already palindrome",
			input:    "aacecaaa",
			expected: "aaacecaaa",
		},
		{
			name:     "single character",
			input:    "a",
			expected: "a",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "two characters",
			input:    "ab",
			expected: "bab",
		},
		{
			name:     "complex case",
			input:    "abcd",
			expected: "dcbabcd",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ShortestPalindrome(tt.input)
			if result != tt.expected {
				t.Errorf("ShortestPalindrome(%q) = %q, want %q", tt.input, result, tt.expected)
			}
			// Verify result is a palindrome
			if !isPalindrome(result) {
				t.Errorf("Result %q is not a palindrome", result)
			}
		})
	}
}

func isPalindrome(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func BenchmarkShortestPalindrome(b *testing.B) {
	input := "aacecaaa"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ShortestPalindrome(input)
	}
}
