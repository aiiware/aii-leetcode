package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullJustify(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		maxWidth int
		expected []string
	}{
		{
			name:     "Empty words list",
			words:    []string{},
			maxWidth: 10,
			expected: []string{},
		},
		{
			name:     "Example 1",
			words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
			maxWidth: 16,
			expected: []string{
				"This    is    an",
				"example  of text",
				"justification.  ",
			},
		},
		{
			name:     "Example 2",
			words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
			maxWidth: 16,
			expected: []string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        ",
			},
		},
		{
			name:     "Single word",
			words:    []string{"Hello"},
			maxWidth: 10,
			expected: []string{"Hello     "},
		},
		{
			name:     "Multiple words exact fit",
			words:    []string{"a", "b", "c", "d"},
			maxWidth: 7,
			expected: []string{"a b c d"},
		},
		{
			name:     "Words with varying lengths",
			words:    []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			maxWidth: 20,
			expected: []string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  ",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FullJustify(tt.words, tt.maxWidth)
			assert.Equal(t, tt.expected, result,
				"FullJustify(%v, %d) = %v, expected %v",
				tt.words, tt.maxWidth, result, tt.expected)
		})
	}
}

func TestFullJustify_EdgeCases(t *testing.T) {
	t.Run("Single long word", func(t *testing.T) {
		words := []string{"Supercalifragilisticexpialidocious"}
		maxWidth := 50
		result := FullJustify(words, maxWidth)
		expected := []string{"Supercalifragilisticexpialidocious                "}
		assert.Equal(t, expected, result)
	})

	t.Run("All words same length", func(t *testing.T) {
		words := []string{"aaa", "bbb", "ccc", "ddd", "eee"}
		maxWidth := 15
		result := FullJustify(words, maxWidth)
		expected := []string{"aaa bbb ccc ddd", "eee            "}
		assert.Equal(t, expected, result)
	})

	t.Run("Max width 1", func(t *testing.T) {
		words := []string{"a", "b", "c"}
		maxWidth := 1
		result := FullJustify(words, maxWidth)
		expected := []string{"a", "b", "c"}
		assert.Equal(t, expected, result)
	})
}

func BenchmarkFullJustify(b *testing.B) {
	// Create a large set of words for benchmarking
	words := []string{
		"This", "is", "a", "benchmark", "test", "for", "text", "justification",
		"algorithm", "performance", "measurement", "in", "Go", "language",
		"implementation", "of", "LeetCode", "problem", "number", "68",
	}
	maxWidth := 50

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FullJustify(words, maxWidth)
	}
}