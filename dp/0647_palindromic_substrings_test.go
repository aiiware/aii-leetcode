package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSubstrings(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Example 1",
			s:        "abc",
			expected: 3,
		},
		{
			name:     "Example 2",
			s:        "aaa",
			expected: 6,
		},
		{
			name:     "Empty string",
			s:        "",
			expected: 0,
		},
		{
			name:     "Single character",
			s:        "a",
			expected: 1,
		},
		{
			name:     "Two same characters",
			s:        "aa",
			expected: 3,
		},
		{
			name:     "Palindrome",
			s:        "aba",
			expected: 4,
		},
		{
			name:     "Complex case",
			s:        "abcba",
			expected: 7,
		},
		{
			name:     "All same characters",
			s:        "aaaaa",
			expected: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountSubstrings(tt.s)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkCountSubstrings(b *testing.B) {
	s := "abcbaabcbaabcbaabcba"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountSubstrings(s)
	}
}