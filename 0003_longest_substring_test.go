package leetcode

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		// LeetCode examples
		{"Example 1: abcabcbb", "abcabcbb", 3}, // "abc" or "bca" or "cab"
		{"Example 2: bbbbb", "bbbbb", 1},       // "b"
		{"Example 3: pwwkew", "pwwkew", 3},     // "wke"
		
		// Edge cases
		{"Empty string", "", 0},
		{"Single character", "a", 1},
		{"Two different characters", "ab", 2},
		{"Two same characters", "aa", 1},
		
		// Complex cases
		{"All unique", "abcdefghijklmnopqrstuvwxyz", 26},
		{"Repeating pattern", "abcabcabc", 3},
		{"Mixed case", "aAbBcC", 6}, // Different ASCII values
		{"With spaces", "a b c d", 3}, // Space repeats: "a b c" (a, space, b) = 3
		
		// Special characters
		{"Special chars", "!@#$%^&*()", 10},
		{"Unicode test", "🎉🎊🎈", 3}, // Emojis (multi-byte) - all different
		
		// Substring at beginning
		{"Beginning substring", "abcdefff", 6}, // "abcdef" (6 chars)
		{"Ending substring", "fffabcde", 6},    // "fabcde" (6 chars)
		{"Middle substring", "ffabcdeff", 6},   // "fabcde" (6 chars)
		
		// LeetCode additional test cases
		{"dvdf", "dvdf", 3}, // "vdf"
		{"anviaj", "anviaj", 5}, // "nviaj"
		{"tmmzuxt", "tmmzuxt", 5}, // "mzuxt"
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LengthOfLongestSubstringRune(tt.input)
			if result != tt.expected {
				t.Errorf("LengthOfLongestSubstringRune(%q) = %d, expected %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestLengthOfLongestSubstringOptimized(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		// ASCII-only tests (optimized version works best for these)
		{"Example 1: abcabcbb", "abcabcbb", 3},
		{"Example 2: bbbbb", "bbbbb", 1},
		{"Example 3: pwwkew", "pwwkew", 3},
		{"Empty string", "", 0},
		{"Single character", "a", 1},
		{"All unique lowercase", "abcdefghijklmnopqrstuvwxyz", 26},
		{"Mixed ASCII", "aAbBcC123!@#", 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LengthOfLongestSubstringOptimized(tt.input)
			if result != tt.expected {
				t.Errorf("LengthOfLongestSubstringOptimized(%q) = %d, expected %d", tt.input, result, tt.expected)
			}
		})
	}
}

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	testCases := []struct {
		name  string
		input string
	}{
		{"Short (10 chars)", "abcdeabcde"},
		{"Medium (50 chars)", "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"},
		{"Long (200 chars)", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat."},
		{"All same char (100 chars)", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
		{"Pattern (100 chars)", "abcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabcabc"},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LengthOfLongestSubstringRune(tc.input)
			}
		})
	}
}

func BenchmarkLengthOfLongestSubstringOptimized(b *testing.B) {
	testCases := []struct {
		name  string
		input string
	}{
		{"Short ASCII", "abcdeabcde"},
		{"Medium ASCII", "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"},
		{"Long ASCII", "The quick brown fox jumps over the lazy dog. THE QUICK BROWN FOX JUMPS OVER THE LAZY DOG. 1234567890!@#$%^&*()"},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LengthOfLongestSubstringOptimized(tc.input)
			}
		})
	}
}

func TestBothImplementationsMatch(t *testing.T) {
	testStrings := []string{
		"",
		"a",
		"aa",
		"ab",
		"abc",
		"abcabcbb",
		"bbbbb",
		"pwwkew",
		"abcdefghijklmnopqrstuvwxyz",
		"aAbBcC",
		"1234567890",
	}

	for _, s := range testStrings {
		result1 := LengthOfLongestSubstringRune(s)
		result2 := LengthOfLongestSubstringOptimized(s)
		if result1 != result2 {
			t.Errorf("Implementations differ for %q: rune=%d, array=%d", s, result1, result2)
		}
	}
}