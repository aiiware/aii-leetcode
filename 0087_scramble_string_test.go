package leetcode

import (
	"fmt"
	"testing"
)

func TestIsScramble(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		expected bool
	}{
		{
			name:     "Example 1",
			s1:       "great",
			s2:       "rgeat",
			expected: true,
		},
		{
			name:     "Example 2",
			s1:       "abcde",
			s2:       "caebd",
			expected: false,
		},
		{
			name:     "Example 3",
			s1:       "a",
			s2:       "a",
			expected: true,
		},
		{
			name:     "Empty strings",
			s1:       "",
			s2:       "",
			expected: true,
		},
		{
			name:     "Different lengths",
			s1:       "ab",
			s2:       "abc",
			expected: false,
		},
		{
			name:     "Simple scramble",
			s1:       "ab",
			s2:       "ba",
			expected: true,
		},
		{
			name:     "Not scramble",
			s1:       "abcd",
			s2:       "bdca",
			expected: false,
		},
		{
			name:     "Complex scramble 1",
			s1:       "abc",
			s2:       "bca",
			expected: true,
		},
		{
			name:     "Complex scramble 2",
			s1:       "abcd",
			s2:       "bcad",
			expected: true,
		},
		{
			name:     "All same characters",
			s1:       "aaaa",
			s2:       "aaaa",
			expected: true,
		},
		{
			name:     "Anagram but not scramble",
			s1:       "abcd",
			s2:       "badc",
			expected: true,
		},
		{
			name:     "Longer example 1",
			s1:       "abcdefghijklmnopqrstuvwxyz",
			s2:       "zyxwvutsrqponmlkjihgfedcba",
			expected: true,
		},
		{
			name:     "Same string",
			s1:       "leetcode",
			s2:       "leetcode",
			expected: true,
		},
		{
			name:     "Single character different",
			s1:       "a",
			s2:       "b",
			expected: false,
		},
		{
			name:     "Two characters swapped",
			s1:       "ab",
			s2:       "ba",
			expected: true,
		},
		{
			name:     "Three characters scramble",
			s1:       "abc",
			s2:       "acb",
			expected: true,
		},
		{
			name:     "Four characters complex",
			s1:       "abcd",
			s2:       "acbd",
			expected: true,
		},
		{
			name:     "LeetCode test case 1",
			s1:       "abb",
			s2:       "bba",
			expected: true,
		},
		{
			name:     "LeetCode test case 2",
			s1:       "abcdbdacbdac",
			s2:       "bdacabcdbdac",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsScramble(tt.s1, tt.s2)
			if result != tt.expected {
				t.Errorf("IsScramble(%q, %q) = %v, expected %v",
					tt.s1, tt.s2, result, tt.expected)
			}
		})
	}
}

func TestAllIsScrambleImplementations(t *testing.T) {
	testCases := []struct {
		name string
		s1   string
		s2   string
	}{
		{"Example 1", "great", "rgeat"},
		{"Example 2", "abcde", "caebd"},
		{"Simple", "a", "a"},
		{"Swap", "ab", "ba"},
		{"Complex", "abcd", "acbd"},
		{"Long", "abcdefgh", "hgfedcba"},
	}

	implementations := []struct {
		name string
		fn   func(string, string) bool
	}{
		{"isScramble", isScramble},
		{"isScrambleDP", isScrambleDP},
		{"isScrambleOptimized", isScrambleOptimized},
		{"isScrambleIterative", isScrambleIterative},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := IsScramble(tc.s1, tc.s2)

			for _, impl := range implementations {
				t.Run(impl.name, func(t *testing.T) {
					result := impl.fn(tc.s1, tc.s2)
					if result != expected {
						t.Errorf("%s(%q, %q) = %v, expected %v",
							impl.name, tc.s1, tc.s2, result, expected)
					}
				})
			}
		})
	}
}

func TestIsScrambleEdgeCases(t *testing.T) {
	t.Run("Empty strings", func(t *testing.T) {
		if !IsScramble("", "") {
			t.Error("Empty strings should be scramble of each other")
		}
	})

	t.Run("Single character same", func(t *testing.T) {
		if !IsScramble("a", "a") {
			t.Error("Single same character should be scramble")
		}
	})

	t.Run("Single character different", func(t *testing.T) {
		if IsScramble("a", "b") {
			t.Error("Different single characters should not be scramble")
		}
	})

	t.Run("Different lengths", func(t *testing.T) {
		if IsScramble("ab", "abc") {
			t.Error("Strings of different lengths should not be scramble")
		}
	})

	t.Run("Same string", func(t *testing.T) {
		if !IsScramble("leetcode", "leetcode") {
			t.Error("Same string should be scramble")
		}
	})

	t.Run("Anagram check", func(t *testing.T) {
		// Test that scramble strings are always anagrams
		s1 := "abcdefghijklmnopqrstuvwxyz"
		s2 := "zyxwvutsrqponmlkjihgfedcba"
		if IsScramble(s1, s2) {
			// Verify they are anagrams
			count1 := make([]int, 26)
			count2 := make([]int, 26)
			for i := 0; i < len(s1); i++ {
				count1[s1[i]-'a']++
				count2[s2[i]-'a']++
			}
			for i := 0; i < 26; i++ {
				if count1[i] != count2[i] {
					t.Errorf("Scramble strings should be anagrams, but character counts differ for %c", 'a'+i)
				}
			}
		}
	})
}

func TestIsScrambleProperties(t *testing.T) {
	// Property-based testing
	implementations := []struct {
		name string
		fn   func(string, string) bool
	}{
		{"isScramble", isScramble},
		{"isScrambleDP", isScrambleDP},
		{"isScrambleOptimized", isScrambleOptimized},
	}

	for _, impl := range implementations {
		t.Run(impl.name+"_properties", func(t *testing.T) {
			// Property 1: Reflexivity - a string is always a scramble of itself
			testStrings := []string{"", "a", "ab", "abc", "abcd", "abcde"}
			for _, s := range testStrings {
				if !impl.fn(s, s) {
					t.Errorf("Reflexivity failed for %q: string should be scramble of itself", s)
				}
			}

			// Property 2: Symmetry - if s1 is scramble of s2, then s2 is scramble of s1
			pairs := [][2]string{
				{"great", "rgeat"},
				{"ab", "ba"},
				{"abc", "bca"},
			}
			for _, pair := range pairs {
				s1, s2 := pair[0], pair[1]
				result1 := impl.fn(s1, s2)
				result2 := impl.fn(s2, s1)
				if result1 != result2 {
					t.Errorf("Symmetry failed for (%q, %q): %v != %v",
						s1, s2, result1, result2)
				}
			}

			// Property 3: Character count preservation
			s1 := "abcdefghijklmnopqrstuvwxyz"
			s2 := "zyxwvutsrqponmlkjihgfedcba"
			if impl.fn(s1, s2) {
				// Count characters
				count1 := make([]int, 26)
				count2 := make([]int, 26)
				for i := 0; i < len(s1); i++ {
					count1[s1[i]-'a']++
					count2[s2[i]-'a']++
				}
				for i := 0; i < 26; i++ {
					if count1[i] != count2[i] {
						t.Errorf("Character count mismatch for %c: %d != %d",
							'a'+i, count1[i], count2[i])
					}
				}
			}
		})
	}
}

func BenchmarkIsScramble(b *testing.B) {
	// Test cases of increasing complexity
	testCases := []struct {
		name string
		s1   string
		s2   string
	}{
		{"Small", "ab", "ba"},
		{"Medium", "great", "rgeat"},
		{"Large", "abcdefghijklmnop", "ponmlkjihgfedcba"},
	}

	implementations := []struct {
		name string
		fn   func(string, string) bool
	}{
		{"isScramble", isScramble},
		{"isScrambleDP", isScrambleDP},
		{"isScrambleOptimized", isScrambleOptimized},
		{"isScrambleIterative", isScrambleIterative},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for _, impl := range implementations {
				b.Run(impl.name, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						impl.fn(tc.s1, tc.s2)
					}
				})
			}
		})
	}
}

func BenchmarkIsScrambleWorstCase(b *testing.B) {
	// Worst case: strings are anagrams but not scrambles
	s1 := "abcdabcdabcdabcd"
	s2 := "dcbadcbadcbadcba"

	b.ResetTimer()
	
	b.Run("isScramble", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isScramble(s1, s2)
		}
	})
	
	b.Run("isScrambleOptimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isScrambleOptimized(s1, s2)
		}
	})
}

// Helper function to generate all scrambles of a string
func generateScrambles(s string) []string {
	if len(s) <= 1 {
		return []string{s}
	}

	result := make(map[string]bool)
	for i := 1; i < len(s); i++ {
		left := s[:i]
		right := s[i:]

		leftScrambles := generateScrambles(left)
		rightScrambles := generateScrambles(right)

		// No swap
		for _, ls := range leftScrambles {
			for _, rs := range rightScrambles {
				result[ls+rs] = true
			}
		}

		// Swap
		for _, ls := range leftScrambles {
			for _, rs := range rightScrambles {
				result[rs+ls] = true
			}
		}
	}

	scrambles := make([]string, 0, len(result))
	for scramble := range result {
		scrambles = append(scrambles, scramble)
	}
	return scrambles
}

func TestIsScrambleComprehensive(t *testing.T) {
	// Test that our implementation correctly identifies all scrambles
	testStrings := []string{"ab", "abc", "abcd"}

	for _, s := range testStrings {
		t.Run(fmt.Sprintf("String %q", s), func(t *testing.T) {
			scrambles := generateScrambles(s)
			
			for _, scramble := range scrambles {
				if !IsScramble(s, scramble) {
					t.Errorf("Failed to recognize scramble: %q -> %q", s, scramble)
				}
			}

			// Also test some non-scrambles (anagrams that aren't scrambles)
			// This is harder to generate, but we can at least verify reflexivity
			if !IsScramble(s, s) {
				t.Errorf("Failed reflexivity for %q", s)
			}
		})
	}
}