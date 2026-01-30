package dp

import (
	"testing"
)

func TestWordBreak(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		expected bool
	}{
		{
			name:     "Example 1",
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			expected: true,
		},
		{
			name:     "Example 2",
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			expected: true,
		},
		{
			name:     "Example 3",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			expected: false,
		},
		{
			name:     "Empty string",
			s:        "",
			wordDict: []string{"a", "b", "c"},
			expected: true, // Empty string can always be segmented
		},
		{
			name:     "Single character in dict",
			s:        "a",
			wordDict: []string{"a"},
			expected: true,
		},
		{
			name:     "Single character not in dict",
			s:        "a",
			wordDict: []string{"b", "c"},
			expected: false,
		},
		{
			name:     "Multiple ways to segment",
			s:        "catsanddog",
			wordDict: []string{"cat", "cats", "and", "sand", "dog"},
			expected: true,
		},
		{
			name:     "Word reuse required",
			s:        "aaaaaaa",
			wordDict: []string{"aaaa", "aaa"},
			expected: true,
		},
		{
			name:     "Impossible segmentation",
			s:        "abcd",
			wordDict: []string{"a", "abc", "b", "cd"},
			expected: true, // "a" + "b" + "cd"
		},
		{
			name:     "Long string with overlap",
			s:        "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
			wordDict: []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"},
			expected: false, // Has 'b' at the end which can't be segmented
		},
		{
			name:     "Dictionary has empty string",
			s:        "abc",
			wordDict: []string{"", "a", "bc"},
			expected: true,
		},
		{
			name:     "All characters separate words",
			s:        "abcdef",
			wordDict: []string{"a", "b", "c", "d", "e", "f"},
			expected: true,
		},
		{
			name:     "Word longer than string",
			s:        "abc",
			wordDict: []string{"abcd", "efgh"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WordBreak(tt.s, tt.wordDict)
			if result != tt.expected {
				t.Errorf("WordBreak(%q, %v) = %v, expected %v", 
					tt.s, tt.wordDict, result, tt.expected)
			}
			
			// Also test the optimized version
			resultOpt := WordBreakOptimized(tt.s, tt.wordDict)
			if resultOpt != tt.expected {
				t.Errorf("WordBreakOptimized(%q, %v) = %v, expected %v", 
					tt.s, tt.wordDict, resultOpt, tt.expected)
			}
		})
	}
}

func BenchmarkWordBreak(b *testing.B) {
	// Create test cases of different sizes
	testCases := []struct {
		name     string
		s        string
		wordDict []string
	}{
		{
			name:     "Small string",
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
		},
		{
			name:     "Medium string",
			s:        "applepenapplepineapple",
			wordDict: []string{"apple", "pen", "pine", "pineapple"},
		},
		{
			name:     "Long repetitive string",
			s:        "aaaaaaaaaaaaaaaaaaaaaa", // 22 a's
			wordDict: []string{"a", "aa", "aaa", "aaaa", "aaaaa"},
		},
		{
			name:     "Very long string",
			s:        "catscatscatscatscatscatscatscatscatscats",
			wordDict: []string{"cat", "cats", "and", "sand", "dog"},
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name+"_standard", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				WordBreak(tc.s, tc.wordDict)
			}
		})
		
		b.Run(tc.name+"_optimized", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				WordBreakOptimized(tc.s, tc.wordDict)
			}
		})
	}
}