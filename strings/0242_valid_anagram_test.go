package strings

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		// Basic cases
		{"empty strings", "", "", true},
		{"single character same", "a", "a", true},
		{"single character different", "a", "b", false},
		{"simple anagram", "anagram", "nagaram", true},
		{"not anagram", "rat", "car", false},
		{"different lengths", "abc", "abcd", false},
		{"same string", "hello", "hello", true},
		{"case sensitive", "Hello", "hello", false}, // 'H' vs 'h' different ASCII
		{"with spaces", "anagram", "nag aram", false}, // spaces not handled

		// Edge cases
		{"repeated characters", "aaa", "aaa", true},
		{"repeated characters different count", "aaab", "aabb", false},
		{"all same character", "zzzz", "zzzz", true},
		{"palindrome", "racecar", "racecar", true},
		{"anagram of palindrome", "aaccrre", "racecar", true},

		// Longer strings
		{"long anagram", "abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcba", true},
		{"long not anagram", "abcdefghijklmnopqrstuvwxyz", "zyxwvutsrqponmlkjihgfedcb", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAnagram(tt.s, tt.t); got != tt.want {
				t.Errorf("IsAnagram(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}

func TestIsAnagramSorting(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		// Basic cases
		{"empty strings", "", "", true},
		{"single character same", "a", "a", true},
		{"single character different", "a", "b", false},
		{"simple anagram", "anagram", "nagaram", true},
		{"not anagram", "rat", "car", false},
		{"different lengths", "abc", "abcd", false},
		{"same string", "hello", "hello", true},
		{"case sensitive", "Hello", "hello", false},
		{"with spaces", "anagram", "nag aram", false},

		// Edge cases
		{"repeated characters", "aaa", "aaa", true},
		{"repeated characters different count", "aaab", "aabb", false},
		{"all same character", "zzzz", "zzzz", true},
		{"palindrome", "racecar", "racecar", true},
		{"anagram of palindrome", "aaccrre", "racecar", true},

		// Unicode cases (sorting handles Unicode)
		{"unicode same", "café", "éfac", true},
		{"unicode different", "café", "cafe", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAnagramSorting(tt.s, tt.t); got != tt.want {
				t.Errorf("IsAnagramSorting(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}

func TestIsAnagramUnicode(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		// Basic cases
		{"empty strings", "", "", true},
		{"single character same", "a", "a", true},
		{"single character different", "a", "b", false},
		{"simple anagram", "anagram", "nagaram", true},
		{"not anagram", "rat", "car", false},
		{"different lengths", "abc", "abcd", false},
		{"same string", "hello", "hello", true},
		{"case sensitive", "Hello", "hello", false},

		// Unicode cases
		{"unicode anagram", "café", "éfac", true},
		{"unicode not anagram", "café", "cafe", false},
		{"emoji anagram", "😀😁😂", "😂😁😀", true},
		{"emoji not anagram", "😀😁😂", "😀😁😁", false},
		{"mixed unicode", "café😀", "😀éfac", true},
		{"chinese characters", "你好", "好你", true},
		{"chinese not anagram", "你好", "你好吗", false},

		// Edge cases
		{"repeated characters", "aaa", "aaa", true},
		{"repeated characters different count", "aaab", "aabb", false},
		{"all same character", "zzzz", "zzzz", true},
		{"palindrome", "racecar", "racecar", true},
		{"anagram of palindrome", "aaccrre", "racecar", true},

		// With spaces and punctuation
		{"with spaces", "anagram", "nag aram", false},
		{"with punctuation", "anagram!", "!nagaram", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAnagramUnicode(tt.s, tt.t); got != tt.want {
				t.Errorf("IsAnagramUnicode(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}

func BenchmarkIsAnagram(b *testing.B) {
	s := "abcdefghijklmnopqrstuvwxyz"
	t := "zyxwvutsrqponmlkjihgfedcba"
	for i := 0; i < b.N; i++ {
		IsAnagram(s, t)
	}
}

func BenchmarkIsAnagramSorting(b *testing.B) {
	s := "abcdefghijklmnopqrstuvwxyz"
	t := "zyxwvutsrqponmlkjihgfedcba"
	for i := 0; i < b.N; i++ {
		IsAnagramSorting(s, t)
	}
}

func BenchmarkIsAnagramUnicode(b *testing.B) {
	s := "abcdefghijklmnopqrstuvwxyz"
	t := "zyxwvutsrqponmlkjihgfedcba"
	for i := 0; i < b.N; i++ {
		IsAnagramUnicode(s, t)
	}
}