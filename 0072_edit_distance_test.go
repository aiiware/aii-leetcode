package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDistance(t *testing.T) {
	tests := []struct {
		name   string
		word1  string
		word2  string
		expected int
	}{
		{
			name:     "Example 1: horse -> ros",
			word1:    "horse",
			word2:    "ros",
			expected: 3,
		},
		{
			name:     "Example 2: intention -> execution",
			word1:    "intention",
			word2:    "execution",
			expected: 5,
		},
		{
			name:     "Empty to empty",
			word1:    "",
			word2:    "",
			expected: 0,
		},
		{
			name:     "Empty to non-empty",
			word1:    "",
			word2:    "abc",
			expected: 3,
		},
		{
			name:     "Non-empty to empty",
			word1:    "abc",
			word2:    "",
			expected: 3,
		},
		{
			name:     "Same strings",
			word1:    "kitten",
			word2:    "kitten",
			expected: 0,
		},
		{
			name:     "One character difference",
			word1:    "cat",
			word2:    "bat",
			expected: 1,
		},
		{
			name:     "All characters different",
			word1:    "abc",
			word2:    "def",
			expected: 3,
		},
		{
			name:     "Insert at beginning",
			word1:    "bc",
			word2:    "abc",
			expected: 1,
		},
		{
			name:     "Insert at end",
			word1:    "ab",
			word2:    "abc",
			expected: 1,
		},
		{
			name:     "Delete from beginning",
			word1:    "abc",
			word2:    "bc",
			expected: 1,
		},
		{
			name:     "Delete from end",
			word1:    "abc",
			word2:    "ab",
			expected: 1,
		},
		{
			name:     "Replace all characters",
			word1:    "aaa",
			word2:    "bbb",
			expected: 3,
		},
		{
			name:     "Complex case 1",
			word1:    "sunday",
			word2:    "saturday",
			expected: 3,
		},
		{
			name:     "Complex case 2",
			word1:    "algorithm",
			word2:    "altruistic",
			expected: 6,
		},
		{
			name:     "Long strings with common prefix",
			word1:    "abcdefghijklmnopqrstuvwxyz",
			word2:    "abcdefghijklmnopqrstuvwxy",
			expected: 1,
		},
		{
			name:     "Long strings with common suffix",
			word1:    "abcdefghijklmnopqrstuvwxyz",
			word2:    "bcdefghijklmnopqrstuvwxyz",
			expected: 1,
		},
		{
			name:     "Strings with repeated characters",
			word1:    "aaaaa",
			word2:    "aaa",
			expected: 2,
		},
		{
			name:     "Strings with pattern",
			word1:    "abcabcabc",
			word2:    "abc",
			expected: 6,
		},
		{
			name:     "Case sensitive (all lowercase in constraints)",
			word1:    "hello",
			word2:    "HELLO",
			expected: 5,
		},
		{
			name:     "Special characters (not in constraints but good test)",
			word1:    "hello123",
			word2:    "hello",
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinDistance(tt.word1, tt.word2)
			assert.Equal(t, tt.expected, result,
				"MinDistance(%q, %q) = %d, expected %d",
				tt.word1, tt.word2, result, tt.expected)
		})
	}
}

func TestMinDistance_EdgeCases(t *testing.T) {
	t.Run("Very long strings", func(t *testing.T) {
		// Create two 500-character strings (max length per constraints)
		word1 := ""
		word2 := ""
		for i := 0; i < 500; i++ {
			word1 += "a"
			if i < 250 {
				word2 += "a"
			} else {
				word2 += "b"
			}
		}
		
		result := MinDistance(word1, word2)
		// Expected: 250 replacements (second half of word2 is 'b' instead of 'a')
		assert.Equal(t, 250, result)
	})

	t.Run("All insertions", func(t *testing.T) {
		word1 := ""
		word2 := "abcdefghijklmnopqrstuvwxyz"
		result := MinDistance(word1, word2)
		assert.Equal(t, 26, result)
	})

	t.Run("All deletions", func(t *testing.T) {
		word1 := "abcdefghijklmnopqrstuvwxyz"
		word2 := ""
		result := MinDistance(word1, word2)
		assert.Equal(t, 26, result)
	})

	t.Run("Palindrome transformation", func(t *testing.T) {
		word1 := "racecar"
		word2 := "raceecar" // Insert 'e' in middle
		result := MinDistance(word1, word2)
		assert.Equal(t, 1, result)
	})

	t.Run("Anagrams", func(t *testing.T) {
		word1 := "listen"
		word2 := "silent"
		result := MinDistance(word1, word2)
		// Minimum operations:
		// listen -> sisten (replace l with s)
		// sisten -> silen (delete t)
		// silen -> silent (insert t) - 3 ops
		// OR:
		// listen -> lisen (delete t)
		// lisen -> silen (replace l with s)
		// silen -> silent (insert t) - 3 ops
		// Correct answer is actually 2 by some definitions, but for this DP it's higher. Let's trace.
		// l->s (1), i->i(1), s->l(2), t->e(3), e->n(4), n->t(5) - no
		// DP table gives 4.
		assert.Equal(t, 4, result)
	})
}

func TestMinDistance_Properties(t *testing.T) {
	t.Run("Symmetry property", func(t *testing.T) {
		testCases := []struct {
			word1 string
			word2 string
		}{
			{"kitten", "sitting"},
			{"sunday", "saturday"},
			{"abc", "def"},
			{"", "hello"},
			{"algorithm", "altruistic"},
		}

		for _, tc := range testCases {
			t.Run(tc.word1+"<->"+tc.word2, func(t *testing.T) {
				d1 := MinDistance(tc.word1, tc.word2)
				d2 := MinDistance(tc.word2, tc.word1)
				assert.Equal(t, d1, d2, "Edit distance should be symmetric")
			})
		}
	})

	t.Run("Triangle inequality", func(t *testing.T) {
		// For any strings a, b, c: d(a,c) <= d(a,b) + d(b,c)
		testCases := []struct {
			a string
			b string
			c string
		}{
			{"cat", "bat", "rat"},
			{"hello", "help", "yelp"},
			{"algorithm", "logarithm", "arithmetic"},
			{"", "a", "ab"},
			{"abc", "abd", "abcd"},
		}

		for _, tc := range testCases {
			t.Run(tc.a+"->"+tc.b+"->"+tc.c, func(t *testing.T) {
				d_ab := MinDistance(tc.a, tc.b)
				d_bc := MinDistance(tc.b, tc.c)
				d_ac := MinDistance(tc.a, tc.c)
				
				assert.LessOrEqual(t, d_ac, d_ab+d_bc,
					"Triangle inequality should hold: d(%s,%s) <= d(%s,%s) + d(%s,%s)",
					tc.a, tc.c, tc.a, tc.b, tc.b, tc.c)
			})
		}
	})

	t.Run("Identity of indiscernibles", func(t *testing.T) {
		// d(a,b) = 0 if and only if a = b
		testCases := []struct {
			word1 string
			word2 string
		}{
			{"", ""},
			{"a", "a"},
			{"hello", "hello"},
			{"abcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxyz"},
		}

		for _, tc := range testCases {
			t.Run(tc.word1+"=="+tc.word2, func(t *testing.T) {
				d := MinDistance(tc.word1, tc.word2)
				assert.Equal(t, 0, d, "Distance should be 0 for identical strings")
			})
		}
	})
}

func BenchmarkMinDistance(b *testing.B) {
	testCases := []struct {
		name  string
		word1 string
		word2 string
	}{
		{"Short strings", "cat", "bat"},
		{"Medium strings", "sunday", "saturday"},
		{"Long similar strings", "abcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxy"},
		{"Long different strings", "aaaaaaaaaa", "bbbbbbbbbb"},
		{"Empty strings", "", ""},
		{"Empty to non-empty", "", "hello"},
		{"Non-empty to empty", "world", ""},
		{"Very long strings", 
			"abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
			"abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxy"},
		{"Complex transformation", "intention", "execution"},
		{"Anagrams", "listen", "silent"},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MinDistance(tc.word1, tc.word2)
			}
		})
	}
}