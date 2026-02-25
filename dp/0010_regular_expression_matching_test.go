package dp

import (
	"testing"
)

func TestIsMatch(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		p        string
		expected bool
	}{
		// Basic cases
		{"empty string and pattern", "", "", true},
		{"empty string with non-empty pattern", "", "a", false},
		{"non-empty string with empty pattern", "a", "", false},
		
		// Single character matches
		{"single char match", "a", "a", true},
		{"single char mismatch", "a", "b", false},
		{"dot matches any char", "a", ".", true},
		{"dot matches any char 2", "b", ".", true},
		
		// Star operator cases
		{"star matches zero occurrences", "", "a*", true},
		{"star matches one occurrence", "a", "a*", true},
		{"star matches multiple occurrences", "aa", "a*", true},
		{"star with dot", "aaa", ".*", true},
		{"star with dot matches any sequence", "abc", ".*", true},
		
		// Complex patterns
		{"multiple stars", "aab", "c*a*b", true},
		{"star with mismatch", "ab", ".*c", false},
		{"star with preceding dot", "ab", ".*", true},
		{"star with preceding dot 2", "abc", ".*", true},
		
		// LeetCode examples
		{"leetcode example 1", "aa", "a", false},
		{"leetcode example 2", "aa", "a*", true},
		{"leetcode example 3", "ab", ".*", true},
		{"leetcode example 4", "aab", "c*a*b", true},
		{"leetcode example 5", "mississippi", "mis*is*p*.", false},
		
		// Edge cases
		{"multiple stars complex", "aaa", "a*a", true},
		{"star with dot complex", "ab", ".*c", false},
		{"long pattern", "aaaaaaaaaaaaab", "a*a*a*a*a*a*a*a*a*a*c", false},
		{"mixed pattern", "abcd", "a.*d", true},
		{"mixed pattern 2", "abcd", "a.*e", false},
		
		// Additional edge cases
		{"pattern ends with star", "a", "ab*", true},
		{"multiple dots", "abc", "a.c", true},
		{"multiple dots mismatch", "abd", "a.c", false},
		{"star after dot", "abc", "a.*c", true},
		{"star after dot mismatch", "abd", "a.*c", false},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isMatch(tt.s, tt.p)
			if result != tt.expected {
				t.Errorf("isMatch(%q, %q) = %v, expected %v", tt.s, tt.p, result, tt.expected)
			}
		})
	}
}

func TestIsMatchOptimized(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		p        string
		expected bool
	}{
		// Basic cases
		{"empty string and pattern", "", "", true},
		{"empty string with non-empty pattern", "", "a", false},
		{"non-empty string with empty pattern", "a", "", false},
		
		// Single character matches
		{"single char match", "a", "a", true},
		{"single char mismatch", "a", "b", false},
		{"dot matches any char", "a", ".", true},
		
		// Star operator cases
		{"star matches zero occurrences", "", "a*", true},
		{"star matches one occurrence", "a", "a*", true},
		{"star matches multiple occurrences", "aa", "a*", true},
		{"star with dot", "aaa", ".*", true},
		
		// Complex patterns
		{"multiple stars", "aab", "c*a*b", true},
		{"star with mismatch", "ab", ".*c", false},
		
		// LeetCode examples
		{"leetcode example 1", "aa", "a", false},
		{"leetcode example 2", "aa", "a*", true},
		{"leetcode example 3", "ab", ".*", true},
		{"leetcode example 4", "aab", "c*a*b", true},
		{"leetcode example 5", "mississippi", "mis*is*p*.", false},
		
		// Edge cases
		{"multiple stars complex", "aaa", "a*a", true},
		{"long pattern", "aaaaaaaaaaaaab", "a*a*a*a*a*a*a*a*a*a*c", false},
		{"pattern ends with star", "a", "ab*", true},
		{"multiple dots", "abc", "a.c", true},
		{"star after dot", "abc", "a.*c", true},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isMatchOptimized(tt.s, tt.p)
			if result != tt.expected {
				t.Errorf("isMatchOptimized(%q, %q) = %v, expected %v", tt.s, tt.p, result, tt.expected)
			}
		})
	}
}

func TestIsMatchConsistency(t *testing.T) {
	// Test that both implementations give the same results
	testCases := []struct {
		s string
		p string
	}{
		{"", ""},
		{"", "a*"},
		{"a", "a"},
		{"a", "."},
		{"aa", "a*"},
		{"ab", ".*"},
		{"aab", "c*a*b"},
		{"mississippi", "mis*is*p*."},
		{"aaa", "a*a"},
		{"abcd", "a.*d"},
		{"aaaaaaaaaaaaab", "a*a*a*a*a*a*a*a*a*a*c"},
		{"a", "ab*"},
		{"abc", "a.c"},
		{"abc", "a.*c"},
	}
	
	for _, tc := range testCases {
		t.Run(tc.s+"_"+tc.p, func(t *testing.T) {
			result1 := isMatch(tc.s, tc.p)
			result2 := isMatchOptimized(tc.s, tc.p)
			
			if result1 != result2 {
				t.Errorf("Inconsistent results for s=%q, p=%q: isMatch=%v, isMatchOptimized=%v",
					tc.s, tc.p, result1, result2)
			}
		})
	}
}

func BenchmarkIsMatch(b *testing.B) {
	testCases := []struct {
		name string
		s    string
		p    string
	}{
		{"short_match", "aa", "a*"},
		{"short_mismatch", "aa", "a"},
		{"medium_match", "aab", "c*a*b"},
		{"medium_mismatch", "mississippi", "mis*is*p*."},
		{"long_match", "aaaaaaaaaa", "a*"},
		{"long_complex", "aaaaaaaaaaaaab", "a*a*a*a*a*a*a*a*a*a*c"},
	}
	
	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isMatch(tc.s, tc.p)
			}
		})
	}
}

func BenchmarkIsMatchOptimized(b *testing.B) {
	testCases := []struct {
		name string
		s    string
		p    string
	}{
		{"short_match", "aa", "a*"},
		{"short_mismatch", "aa", "a"},
		{"medium_match", "aab", "c*a*b"},
		{"medium_mismatch", "mississippi", "mis*is*p*."},
		{"long_match", "aaaaaaaaaa", "a*"},
		{"long_complex", "aaaaaaaaaaaaab", "a*a*a*a*a*a*a*a*a*a*c"},
	}
	
	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isMatchOptimized(tc.s, tc.p)
			}
		})
	}
}