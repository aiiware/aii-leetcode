package dp

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordBreakII(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		expected []string
	}{
		{
			name:     "Example 1",
			s:        "catsanddog",
			wordDict: []string{"cat", "cats", "and", "sand", "dog"},
			expected: []string{"cats and dog", "cat sand dog"},
		},
		{
			name:     "Example 2",
			s:        "pineapplepenapple",
			wordDict: []string{"apple", "pen", "applepen", "pine", "pineapple"},
			expected: []string{"pine apple pen apple", "pineapple pen apple", "pine applepen apple"},
		},
		{
			name:     "Example 3",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			expected: []string{},
		},
		{
			name:     "Single word match",
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			expected: []string{"leet code"},
		},
		{
			name:     "Multiple matches",
			s:        "aaaa",
			wordDict: []string{"a", "aa", "aaa", "aaaa"},
			expected: []string{
				"a a a a",
				"a a aa",
				"a aa a",
				"a aaa",
				"aa a a",
				"aa aa",
				"aaa a",
				"aaaa",
			},
		},
		{
			name:     "No match",
			s:        "leetcode",
			wordDict: []string{"leet", "cod"},
			expected: []string{},
		},
		{
			name:     "Empty string",
			s:        "",
			wordDict: []string{"a", "b", "c"},
			expected: []string{""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := wordBreakII(tt.s, tt.wordDict)
			
			// Sort both slices for comparison
			sort.Strings(result)
			sort.Strings(tt.expected)
			
			assert.Equal(t, tt.expected, result, "Should return correct sentences")
		})
	}
}

func TestWordBreakII_EdgeCases(t *testing.T) {
	t.Run("Long string with many combinations", func(t *testing.T) {
		s := "aaaaaaaaaaaaaaaaaaaa"
		wordDict := []string{"a", "aa", "aaa", "aaaa", "aaaaa"}
		result := wordBreakII(s, wordDict)
		// Just check that we get some results (not checking all combinations)
		assert.True(t, len(result) > 0, "Should return some sentences")
	})

	t.Run("Dictionary with overlapping words", func(t *testing.T) {
		s := "applepenapple"
		wordDict := []string{"apple", "pen", "applepen"}
		result := wordBreakII(s, wordDict)
		expected := []string{"apple pen apple", "applepen apple"}
		sort.Strings(result)
		sort.Strings(expected)
		assert.Equal(t, expected, result, "Should handle overlapping words")
	})

	t.Run("Word longer than string", func(t *testing.T) {
		s := "cat"
		wordDict := []string{"cats", "dog", "mouse"}
		result := wordBreakII(s, wordDict)
		assert.Equal(t, []string{}, result, "Should return empty when no word matches")
	})
}

func BenchmarkWordBreakII(b *testing.B) {
	s := "pineapplepenapple"
	wordDict := []string{"apple", "pen", "applepen", "pine", "pineapple"}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wordBreakII(s, wordDict)
	}
}

func BenchmarkWordBreakII_Large(b *testing.B) {
	s := "aaaaaaaaaaaaaaaaaaaa"
	wordDict := []string{"a", "aa", "aaa", "aaaa", "aaaaa"}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wordBreakII(s, wordDict)
	}
}