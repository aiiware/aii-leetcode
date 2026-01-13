package leetcode

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLadders(t *testing.T) {
	tests := []struct {
		name       string
		beginWord  string
		endWord    string
		wordList   []string
		expected   [][]string
		shouldSort bool // Whether to sort results for comparison
	}{
		{
			name:      "Example 1",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			expected: [][]string{
				{"hit", "hot", "dot", "dog", "cog"},
				{"hit", "hot", "lot", "log", "cog"},
			},
			shouldSort: true,
		},
		{
			name:      "Example 2",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			expected:  [][]string{},
		},
		{
			name:      "Single transformation",
			beginWord: "hit",
			endWord:   "hot",
			wordList:  []string{"hot"},
			expected: [][]string{
				{"hit", "hot"},
			},
		},
		{
			name:      "No transformation possible",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cot"},
			expected:  [][]string{},
		},
		{
			name:      "Begin word equals end word",
			beginWord: "hit",
			endWord:   "hit",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			expected: [][]string{
				{"hit"},
			},
		},
		{
			name:      "Multiple shortest paths",
			beginWord: "red",
			endWord:   "tax",
			wordList:  []string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"},
			expected: [][]string{
				{"red", "ted", "tad", "tax"},
				{"red", "ted", "tex", "tax"},
				{"red", "rex", "tex", "tax"},
			},
			shouldSort: true,
		},
		{
			name:      "Word length 1",
			beginWord: "a",
			endWord:   "c",
			wordList:  []string{"a", "b", "c"},
			expected: [][]string{
				{"a", "c"},
			},
		},
		{
			name:      "Empty word list",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{},
			expected:  [][]string{},
		},
		{
			name:      "Begin word in word list",
			beginWord: "hot",
			endWord:   "dog",
			wordList:  []string{"hot", "dot", "dog"},
			expected: [][]string{
				{"hot", "dot", "dog"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindLadders(tt.beginWord, tt.endWord, tt.wordList)

			if tt.shouldSort {
				// Sort both expected and result for consistent comparison
				sortStringSlices(tt.expected)
				sortStringSlices(result)
			}

			assert.Equal(t, tt.expected, result,
				"FindLadders(%q, %q, %v) = %v, expected %v",
				tt.beginWord, tt.endWord, tt.wordList, result, tt.expected)
		})
	}
}

func TestFindLadders_EdgeCases(t *testing.T) {
	t.Run("Large word list", func(t *testing.T) {
		// Create a word list with many similar words
		wordList := []string{}
		for i := 0; i < 100; i++ {
			wordList = append(wordList, "word"+string(rune('a'+(i%26))))
		}

		beginWord := "worda"
		endWord := "wordz"

		// Add endWord to wordList
		wordList = append(wordList, endWord)

		result := FindLadders(beginWord, endWord, wordList)
		// Just verify it doesn't panic and returns something
		assert.NotNil(t, result)
	})

	t.Run("All words same length", func(t *testing.T) {
		beginWord := "abcde"
		endWord := "vwxyz"
		wordList := []string{
			"abcdf", "abcef", "abdef", "acdef", "bcdef",
			"vwxya", "vwxyb", "vwxyc", "vwxyd",
		}

		result := FindLadders(beginWord, endWord, wordList)
		assert.Empty(t, result, "Should return empty when no path exists")
	})

	t.Run("Direct transformation exists", func(t *testing.T) {
		beginWord := "cold"
		endWord := "cord"
		wordList := []string{"cord", "card", "ward", "warm", "cold"}

		result := FindLadders(beginWord, endWord, wordList)
		assert.Equal(t, [][]string{{"cold", "cord"}}, result)
	})
}

func BenchmarkFindLadders(b *testing.B) {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindLadders(beginWord, endWord, wordList)
	}
}

// Helper function to sort slices of string slices
func sortStringSlices(slices [][]string) {
	sort.Slice(slices, func(i, j int) bool {
		// Compare first by length
		if len(slices[i]) != len(slices[j]) {
			return len(slices[i]) < len(slices[j])
		}
		// Then lexicographically
		for k := 0; k < len(slices[i]); k++ {
			if slices[i][k] != slices[j][k] {
				return slices[i][k] < slices[j][k]
			}
		}
		return false
	})
}