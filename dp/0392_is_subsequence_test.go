package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{
			name:     "Example 1",
			s:        "ace",
			t:        "abcde",
			expected: true,
		},
		{
			name:     "Example 2",
			s:        "aec",
			t:        "abcde",
			expected: false,
		},
		{
			name:     "Empty s",
			s:        "",
			t:        "abcde",
			expected: true,
		},
		{
			name:     "Empty t",
			s:        "abc",
			t:        "",
			expected: false,
		},
		{
			name:     "Both empty",
			s:        "",
			t:        "",
			expected: true,
		},
		{
			name:     "s equals t",
			s:        "abc",
			t:        "abc",
			expected: true,
		},
		{
			name:     "s longer than t",
			s:        "abcde",
			t:        "abc",
			expected: false,
		},
		{
			name:     "Complex case",
			s:        "aab",
			t:        "baabababab",
			expected: true,
		},
		{
			name:     "No match",
			s:        "abc",
			t:        "def",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsSubsequence(tt.s, tt.t)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkIsSubsequence(b *testing.B) {
	s := "ace"
	t := "abcdeabcdeabcdeabcdeabcdeabcdeabcdeabcdeabcde"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsSubsequence(s, t)
	}
}