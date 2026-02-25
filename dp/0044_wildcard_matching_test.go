package dp

import (
	"testing"
)

func TestIsMatchWildcard(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		p        string
		expected bool
	}{
		// LeetCode examples
		{"leetcode example 1", "aa", "a", false},
		{"leetcode example 2", "aa", "*", true},
		{"leetcode example 3", "cb", "?a", false},
		{"leetcode example 4", "adceb", "*a*b", true},
		{"leetcode example 5", "acdcb", "a*c?b", false},
		
		// Basic cases
		{"empty string and pattern", "", "", true},
		{"empty string with non-empty pattern", "", "a", false},
		{"non-empty string with empty pattern", "a", "", false},
		{"empty string with star pattern", "", "*", true},
		{"empty string with multiple stars", "", "***", true},
		{"empty string with question mark", "", "?", false},
		
		// Single character matches
		{"single char match", "a", "a", true},
		{"single char mismatch", "a", "b", false},
		{"question mark matches any char", "a", "?", true},
		{"question mark matches any char 2", "b", "?", true},
		
		// Star operator cases
		{"star matches empty sequence", "a", "a*", true},
		{"star matches one char", "a", "*a", true},
		{"star matches multiple chars", "ab", "*", true},
		{"star in middle matches", "abc", "a*c", true},
		{"multiple stars", "abc", "a**c", true},
		{"star at beginning", "abc", "*bc", true},
		{"star at end", "abc", "ab*", true},
		{"star matches everything", "anything", "*", true},
		
		// Question mark cases
		{"multiple question marks", "abc", "a?c", true},
		{"question mark mismatch", "abd", "a?c", false},
		{"question mark with star", "abc", "a?*", true},
		{"question mark with star 2", "abc", "*?c", true},
		
		// Complex patterns
		{"complex pattern 1", "mississippi", "m*iss*", true},
		{"complex pattern 2", "mississippi", "m*iss*pi", true},
		{"complex pattern 3", "mississippi", "m*iss*pp*", true},
		{"complex pattern 4", "hello world", "h*lo*w?rld", true},
		{"complex pattern 5", "hello world", "h*lo*w?rld?", false},
		
		// Edge cases with multiple stars
		{"consecutive stars", "abc", "a***c", true},
		{"stars and question marks", "abcde", "a?c*e", true},
		{"stars and question marks 2", "abcde", "a?c*", true},
		
		// Long strings
		{"long string match", "aaaaaaaaaa", "a*a", true},
		{"long string with stars", "aaaaaaaaaa", "*a*", true},
		{"long string mismatch", "aaaaaaaaaa", "a*b", false},
		
		// Special cases
		{"pattern longer than string", "a", "a*b*c*", false}, // Fixed: "a" doesn't match "a*b*c*" because after 'a' we need 'b'
		{"all question marks", "abc", "???", true},
		{"all question marks mismatch", "ab", "???", false},
		{"mixed pattern 1", "axbxcxdxe", "a*b?c*d?e", true},
		{"mixed pattern 2", "axbxcxdxe", "a*b?c*d?f", false},
		
		// Real-world like patterns
		{"file pattern 1", "document.txt", "*.txt", true},
		{"file pattern 2", "image.jpg", "*.jpg", true},
		{"file pattern 3", "test_file.go", "test_*.go", true},
		{"file pattern 4", "main_test.go", "*_test.go", true},
		{"file pattern 5", "main.go", "*_test.go", false},
		
		// Additional test cases from discussion
		{"star matches zero chars", "ab", "a*b", true},
		{"star matches multiple chars middle", "axxxxxxb", "a*b", true},
		{"question mark at end", "abc", "ab?", true},
		{"question mark at beginning", "abc", "?bc", true},
		{"pattern with only stars", "anything", "*****", true},
		{"pattern stars and chars", "abc", "*a*b*c*", true},
		
		// Additional edge cases
		{"multiple letters with stars", "ab", "a*b", true},
		{"multiple letters with stars 2", "ac", "a*b*c", false}, // "ac" doesn't have 'b'
		{"multiple letters with stars 3", "abc", "a*b*c", true},
		{"star then char", "ab", "*b", true},
		{"star then char 2", "ac", "*b", false},
		{"char then star then char", "abc", "a*c", true},
		{"char then star then char 2", "ac", "a*c", true},
		{"char then star then char 3", "ab", "a*c", false},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isMatchWildcard(tt.s, tt.p)
			if result != tt.expected {
				t.Errorf("isMatchWildcard(%q, %q) = %v, expected %v", tt.s, tt.p, result, tt.expected)
			}
		})
	}
}

func TestIsMatchWildcardOptimized(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		p        string
		expected bool
	}{
		// LeetCode examples
		{"leetcode example 1", "aa", "a", false},
		{"leetcode example 2", "aa", "*", true},
		{"leetcode example 3", "cb", "?a", false},
		{"leetcode example 4", "adceb", "*a*b", true},
		{"leetcode example 5", "acdcb", "a*c?b", false},
		
		// Basic cases
		{"empty string and pattern", "", "", true},
		{"empty string with star pattern", "", "*", true},
		{"empty string with multiple stars", "", "***", true},
		
		// Single character matches
		{"single char match", "a", "a", true},
		{"question mark matches any char", "a", "?", true},
		
		// Star operator cases
		{"star matches empty sequence", "a", "a*", true},
		{"star matches one char", "a", "*a", true},
		{"star matches multiple chars", "ab", "*", true},
		{"star in middle matches", "abc", "a*c", true},
		
		// Question mark cases
		{"multiple question marks", "abc", "a?c", true},
		{"question mark with star", "abc", "a?*", true},
		
		// Complex patterns
		{"complex pattern 1", "mississippi", "m*iss*", true},
		{"complex pattern 2", "mississippi", "m*iss*pi", true},
		
		// Edge cases
		{"consecutive stars", "abc", "a***c", true},
		{"long string match", "aaaaaaaaaa", "a*a", true},
		{"pattern longer than string", "a", "a*b*c*", false}, // Fixed
		
		// File patterns
		{"file pattern 1", "document.txt", "*.txt", true},
		{"file pattern 2", "main_test.go", "*_test.go", true},
		
		// Additional edge cases
		{"multiple letters with stars", "ab", "a*b", true},
		{"star then char", "ab", "*b", true},
		{"char then star then char", "abc", "a*c", true},
	}
	
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isMatchWildcardOptimized(tt.s, tt.p)
			if result != tt.expected {
				t.Errorf("isMatchWildcardOptimized(%q, %q) = %v, expected %v", tt.s, tt.p, result, tt.expected)
			}
		})
	}
}

func TestIsMatchWildcardConsistency(t *testing.T) {
	// Test that both implementations give the same results
	testCases := []struct {
		s string
		p string
	}{
		{"", ""},
		{"", "*"},
		{"", "***"},
		{"", "?"},
		{"a", "a"},
		{"a", "?"},
		{"a", "*"},
		{"aa", "a"},
		{"aa", "*"},
		{"cb", "?a"},
		{"adceb", "*a*b"},
		{"acdcb", "a*c?b"},
		{"abc", "a?c"},
		{"abc", "a*c"},
		{"mississippi", "m*iss*"},
		{"hello world", "h*lo*w?rld"},
		{"aaaaaaaaaa", "a*a"},
		{"document.txt", "*.txt"},
		{"main_test.go", "*_test.go"},
		{"axbxcxdxe", "a*b?c*d?e"},
		{"anything", "*****"},
		{"ab", "a*b"},
		{"abc", "a*b*c"},
		{"ac", "*b"},
	}
	
	for _, tc := range testCases {
		t.Run(tc.s+"_"+tc.p, func(t *testing.T) {
			result1 := isMatchWildcard(tc.s, tc.p)
			result2 := isMatchWildcardOptimized(tc.s, tc.p)
			
			if result1 != result2 {
				t.Errorf("Inconsistent results for s=%q, p=%q: isMatchWildcard=%v, isMatchWildcardOptimized=%v",
					tc.s, tc.p, result1, result2)
			}
		})
	}
}

func BenchmarkIsMatchWildcard(b *testing.B) {
	testCases := []struct {
		name string
		s    string
		p    string
	}{
		{"short_match", "aa", "*"},
		{"short_mismatch", "aa", "a"},
		{"medium_match", "adceb", "*a*b"},
		{"medium_mismatch", "acdcb", "a*c?b"},
		{"long_match", "aaaaaaaaaa", "*"},
		{"long_complex", "mississippi", "m*iss*pi"},
		{"file_pattern", "document_test_file.go", "*_test_*.go"},
	}
	
	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isMatchWildcard(tc.s, tc.p)
			}
		})
	}
}

func BenchmarkIsMatchWildcardOptimized(b *testing.B) {
	testCases := []struct {
		name string
		s    string
		p    string
	}{
		{"short_match", "aa", "*"},
		{"short_mismatch", "aa", "a"},
		{"medium_match", "adceb", "*a*b"},
		{"medium_mismatch", "acdcb", "a*c?b"},
		{"long_match", "aaaaaaaaaa", "*"},
		{"long_complex", "mississippi", "m*iss*pi"},
		{"file_pattern", "document_test_file.go", "*_test_*.go"},
	}
	
	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isMatchWildcardOptimized(tc.s, tc.p)
			}
		})
	}
}