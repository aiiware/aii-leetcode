package dp

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		name     string
		text1    string
		text2    string
		expected int
	}{
		{
			name:     "Example 1",
			text1:    "abcde",
			text2:    "ace",
			expected: 3,
		},
		{
			name:     "Example 2",
			text1:    "abc",
			text2:    "abc",
			expected: 3,
		},
		{
			name:     "Example 3",
			text1:    "abc",
			text2:    "def",
			expected: 0,
		},
		{
			name:     "Empty strings",
			text1:    "",
			text2:    "",
			expected: 0,
		},
		{
			name:     "First string empty",
			text1:    "",
			text2:    "abc",
			expected: 0,
		},
		{
			name:     "Second string empty",
			text1:    "abc",
			text2:    "",
			expected: 0,
		},
		{
			name:     "Single character match",
			text1:    "a",
			text2:    "a",
			expected: 1,
		},
		{
			name:     "Single character no match",
			text1:    "a",
			text2:    "b",
			expected: 0,
		},
		{
			name:     "Subsequence not substring",
			text1:    "abcde",
			text2:    "ace",
			expected: 3,
		},
		{
			name:     "Repeated characters",
			text1:    "aaaa",
			text2:    "aa",
			expected: 2,
		},
		{
			name:     "Longer example",
			text1:    "abcdefghij",
			text2:    "acegi",
			expected: 5,
		},
		{
			name:     "Interleaved subsequence",
			text1:    "abcabcabc",
			text2:    "abc",
			expected: 3,
		},
		{
			name:     "Palindrome subsequence",
			text1:    "racecar",
			text2:    "racecar",
			expected: 7,
		},
		{
			name:     "Reverse strings",
			text1:    "abcde",
			text2:    "edcba",
			expected: 1,
		},
		{
			name:     "All same characters",
			text1:    "zzzzz",
			text2:    "zzzzz",
			expected: 5,
		},
		{
			name:     "Long strings with partial match",
			text1:    "abcdefghijklmnopqrstuvwxyz",
			text2:    "acegikmoqsuwy",
			expected: 13,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LongestCommonSubsequence(tt.text1, tt.text2)
			assert.Equal(t, tt.expected, result,
				"LongestCommonSubsequence(%q, %q) = %d, expected %d",
				tt.text1, tt.text2, result, tt.expected)
		})
	}
}

func TestLongestCommonSubsequenceOptimized(t *testing.T) {
	tests := []struct {
		name     string
		text1    string
		text2    string
		expected int
	}{
		{
			name:     "Example 1",
			text1:    "abcde",
			text2:    "ace",
			expected: 3,
		},
		{
			name:     "Example 2",
			text1:    "abc",
			text2:    "abc",
			expected: 3,
		},
		{
			name:     "Example 3",
			text1:    "abc",
			text2:    "def",
			expected: 0,
		},
		{
			name:     "First string shorter",
			text1:    "ace",
			text2:    "abcde",
			expected: 3,
		},
		{
			name:     "Second string shorter",
			text1:    "abcde",
			text2:    "ace",
			expected: 3,
		},
		{
			name:     "Empty strings",
			text1:    "",
			text2:    "",
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LongestCommonSubsequenceOptimized(tt.text1, tt.text2)
			assert.Equal(t, tt.expected, result,
				"LongestCommonSubsequenceOptimized(%q, %q) = %d, expected %d",
				tt.text1, tt.text2, result, tt.expected)
		})
	}
}

func TestLongestCommonSubsequenceRecursive(t *testing.T) {
	tests := []struct {
		name     string
		text1    string
		text2    string
		expected int
	}{
		{
			name:     "Example 1",
			text1:    "abcde",
			text2:    "ace",
			expected: 3,
		},
		{
			name:     "Example 2",
			text1:    "abc",
			text2:    "abc",
			expected: 3,
		},
		{
			name:     "Example 3",
			text1:    "abc",
			text2:    "def",
			expected: 0,
		},
		{
			name:     "Small strings",
			text1:    "ab",
			text2:    "ac",
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LongestCommonSubsequenceRecursive(tt.text1, tt.text2)
			assert.Equal(t, tt.expected, result,
				"LongestCommonSubsequenceRecursive(%q, %q) = %d, expected %d",
				tt.text1, tt.text2, result, tt.expected)
		})
	}
}

func TestLongestCommonSubsequenceWithReconstruction(t *testing.T) {
	tests := []struct {
		name           string
		text1          string
		text2          string
		expectedLCS    string
		expectedLength int
	}{
		{
			name:           "Example 1",
			text1:          "abcde",
			text2:          "ace",
			expectedLCS:    "ace",
			expectedLength: 3,
		},
		{
			name:           "Example 2",
			text1:          "abc",
			text2:          "abc",
			expectedLCS:    "abc",
			expectedLength: 3,
		},
		{
			name:           "Example 3",
			text1:          "abc",
			text2:          "def",
			expectedLCS:    "",
			expectedLength: 0,
		},
		{
			name:           "Multiple possible LCS",
			text1:          "abcab",
			text2:          "acb",
			expectedLength: 3,
			// Note: Could be "acb" or "abc", depends on implementation
		},
		{
			name:           "Empty strings",
			text1:          "",
			text2:          "",
			expectedLCS:    "",
			expectedLength: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lcs, length := LongestCommonSubsequenceWithReconstruction(tt.text1, tt.text2)
			
			// Always check length
			assert.Equal(t, tt.expectedLength, length,
				"LongestCommonSubsequenceWithReconstruction(%q, %q) length = %d, expected %d",
				tt.text1, tt.text2, length, tt.expectedLength)
			
			// Check LCS string if specified
			if tt.expectedLCS != "" {
				assert.Equal(t, tt.expectedLCS, lcs,
					"LongestCommonSubsequenceWithReconstruction(%q, %q) LCS = %q, expected %q",
					tt.text1, tt.text2, lcs, tt.expectedLCS)
			}
			
			// Verify that the LCS is indeed a subsequence of both strings
			if length > 0 {
				assert.True(t, isSubsequence(lcs, tt.text1),
					"LCS %q is not a subsequence of %q", lcs, tt.text1)
				assert.True(t, isSubsequence(lcs, tt.text2),
					"LCS %q is not a subsequence of %q", lcs, tt.text2)
			}
		})
	}
}

func TestLCSAllImplementationsConsistent(t *testing.T) {
	testCases := []struct {
		text1 string
		text2 string
	}{
		{"abcde", "ace"},
		{"abc", "abc"},
		{"abc", "def"},
		{"", ""},
		{"a", "a"},
		{"a", "b"},
		{"aaaa", "aa"},
		{"abcdefghij", "acegi"},
		{"racecar", "racecar"},
		{"abcde", "edcba"},
	}

	for _, tc := range testCases {
		t.Run(tc.text1+"_"+tc.text2, func(t *testing.T) {
			dpResult := LongestCommonSubsequence(tc.text1, tc.text2)
			optResult := LongestCommonSubsequenceOptimized(tc.text1, tc.text2)
			recResult := LongestCommonSubsequenceRecursive(tc.text1, tc.text2)
			_, reconResult := LongestCommonSubsequenceWithReconstruction(tc.text1, tc.text2)
			
			// All implementations should return the same length
			assert.Equal(t, dpResult, optResult,
				"DP and optimized results differ for (%q, %q): %d vs %d",
				tc.text1, tc.text2, dpResult, optResult)
			assert.Equal(t, dpResult, recResult,
				"DP and recursive results differ for (%q, %q): %d vs %d",
				tc.text1, tc.text2, dpResult, recResult)
			assert.Equal(t, dpResult, reconResult,
				"DP and reconstruction results differ for (%q, %q): %d vs %d",
				tc.text1, tc.text2, dpResult, reconResult)
		})
	}
}

func TestLCSEdgeCases(t *testing.T) {
	t.Run("Very long strings", func(t *testing.T) {
		// Create strings of length 1000 (max constraint)
		text1 := ""
		text2 := ""
		for i := 0; i < 1000; i++ {
			text1 += string('a' + byte(i%26))
			text2 += string('a' + byte((i+13)%26))
		}
		
		result := LongestCommonSubsequence(text1, text2)
		// We don't know the exact expected value, but it should be computed
		assert.GreaterOrEqual(t, result, 0)
		assert.LessOrEqual(t, result, 1000)
	})
	
	t.Run("Same long string", func(t *testing.T) {
		text := ""
		for i := 0; i < 500; i++ {
			text += string('a' + byte(i%26))
		}
		
		result := LongestCommonSubsequence(text, text)
		assert.Equal(t, len(text), result)
	})
	
	t.Run("One character repeated", func(t *testing.T) {
		text1 := "aaaaaaaaaa" // 10 a's
		text2 := "aaaaa"      // 5 a's
		
		result := LongestCommonSubsequence(text1, text2)
		assert.Equal(t, 5, result)
	})
}

// Helper function to check if s is a subsequence of t
func isSubsequence(s, t string) bool {
	if len(s) == 0 {
		return true
	}
	
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	
	return i == len(s)
}

func BenchmarkLongestCommonSubsequence(b *testing.B) {
	// Create test strings of moderate size
	text1 := strings.Repeat("abcdefghijklmnopqrstuvwxyz", 10) // 260 characters
	text2 := strings.Repeat("acegikmoqsuwy", 20)              // 260 characters
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LongestCommonSubsequence(text1, text2)
	}
}

func BenchmarkLongestCommonSubsequenceOptimized(b *testing.B) {
	// Create test strings of moderate size
	text1 := strings.Repeat("abcdefghijklmnopqrstuvwxyz", 10) // 260 characters
	text2 := strings.Repeat("acegikmoqsuwy", 20)              // 260 characters
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LongestCommonSubsequenceOptimized(text1, text2)
	}
}

func BenchmarkLongestCommonSubsequenceRecursive(b *testing.B) {
	// Use smaller strings for recursive benchmark (it's slower)
	text1 := "abcdefghij"
	text2 := "acegi"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LongestCommonSubsequenceRecursive(text1, text2)
	}
}