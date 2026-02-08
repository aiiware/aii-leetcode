package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Example 1",
			s:        "Hello World",
			expected: 5,
		},
		{
			name:     "Example 2",
			s:        "   fly me   to   the moon  ",
			expected: 4,
		},
		{
			name:     "Example 3",
			s:        "luffy is still joyboy",
			expected: 6,
		},
		{
			name:     "Single word",
			s:        "hello",
			expected: 5,
		},
		{
			name:     "Only spaces",
			s:        "   ",
			expected: 0,
		},
		{
			name:     "Empty string",
			s:        "",
			expected: 0,
		},
		{
			name:     "Trailing spaces",
			s:        "a   ",
			expected: 1,
		},
		{
			name:     "Multiple words",
			s:        "the quick brown fox",
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LengthOfLastWord(tt.s)
			assert.Equal(t, tt.expected, result, "LengthOfLastWord(%q) = %d, expected %d", tt.s, result, tt.expected)
		})
	}
}

func BenchmarkLengthOfLastWord(b *testing.B) {
	s := "Hello World Hello World Hello World Hello World Hello World"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LengthOfLastWord(s)
	}
}