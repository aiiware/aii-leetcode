package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLadderLength(t *testing.T) {
	tests := []struct {
		name      string
		beginWord string
		endWord   string
		wordList  []string
		expected  int
	}{
		{
			name:      "Example 1",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			expected:  5,
		},
		{
			name:      "Example 2",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			expected:  0,
		},
		{
			name:      "Single transformation",
			beginWord: "hit",
			endWord:   "hot",
			wordList:  []string{"hot"},
			expected:  2,
		},
		{
			name:      "No transformation possible",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cot"},
			expected:  0,
		},
		{
			name:      "Begin word equals end word",
			beginWord: "hit",
			endWord:   "hit",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			expected:  1,
		},
		{
			name:      "Multiple paths same length",
			beginWord: "red",
			endWord:   "tax",
			wordList:  []string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"},
			expected:  4,
		},
		{
			name:      "Word length 1",
			beginWord: "a",
			endWord:   "c",
			wordList:  []string{"a", "b", "c"},
			expected:  2,
		},
		{
			name:      "Empty word list",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{},
			expected:  0,
		},
		{
			name:      "Begin word in word list",
			beginWord: "hot",
			endWord:   "dog",
			wordList:  []string{"hot", "dot", "dog"},
			expected:  3,
		},
		{
			name:      "Direct transformation",
			beginWord: "cold",
			endWord:   "cord",
			wordList:  []string{"cord", "card", "ward", "warm", "cold"},
			expected:  2,
		},
		{
			name:      "Long chain",
			beginWord: "aaaaa",
			endWord:   "bbbbb",
			wordList: []string{
				"aaaab", "aaabb", "aabbb", "abbbb", "bbbbb",
				"baaaa", "bbaaa", "bbbaa", "bbbba",
			},
			expected: 6, // aaaaa -> aaaab -> aaabb -> aabbb -> abbbb -> bbbbb
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LadderLength(tt.beginWord, tt.endWord, tt.wordList)
			assert.Equal(t, tt.expected, result,
				"LadderLength(%q, %q, %v) = %d, expected %d",
				tt.beginWord, tt.endWord, tt.wordList, result, tt.expected)
		})
	}
}

func TestLadderLength_EdgeCases(t *testing.T) {
	t.Run("Large word list", func(t *testing.T) {
		// Create a large word list
		wordList := make([]string, 5000)
		for i := range wordList {
			wordList[i] = "word" + string(rune('a'+(i%26)))
		}

		beginWord := "worda"
		endWord := "wordz"
		wordList[0] = endWord // Ensure endWord is in list

		result := LadderLength(beginWord, endWord, wordList)
		// Just verify it doesn't panic
		assert.True(t, result >= 0)
	})

	t.Run("All words same", func(t *testing.T) {
		beginWord := "abc"
		endWord := "def"
		wordList := []string{"abc", "abd", "abe", "abf", "abg"}

		result := LadderLength(beginWord, endWord, wordList)
		assert.Equal(t, 0, result, "Should return 0 when endWord not in list")
	})

	t.Run("Maximum word length", func(t *testing.T) {
		beginWord := "aaaaaaaaaa"
		endWord := "bbbbbbbbbb"
		wordList := []string{
			"aaaaaaaaab",
			"aaaaaaaabb",
			"aaaaaaabbb",
			"aaaaaabbbb",
			"aaaaabbbbb",
			"aaaabbbbbb",
			"aaabbbbbbb",
			"aabbbbbbbb",
			"abbbbbbbbb",
			"bbbbbbbbbb",
		}

		result := LadderLength(beginWord, endWord, wordList)
		assert.Equal(t, 11, result) // 10 transformations + beginWord
	})
}

func BenchmarkLadderLength(b *testing.B) {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LadderLength(beginWord, endWord, wordList)
	}
}