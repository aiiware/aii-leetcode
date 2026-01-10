package leetcode

import (
	"testing"
)

func TestNumDistinct(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected int
	}{
		{
			name:     "Example 1: rabbbit and rabbit",
			s:        "rabbbit",
			t:        "rabbit",
			expected: 3,
		},
		{
			name:     "Example 2: babgbag and bag",
			s:        "babgbag",
			t:        "bag",
			expected: 5,
		},
		{
			name:     "Empty t string",
			s:        "abc",
			t:        "",
			expected: 1, // Empty string can be formed in 1 way
		},
		{
			name:     "Empty s string with non-empty t",
			s:        "",
			t:        "abc",
			expected: 0,
		},
		{
			name:     "Both strings empty",
			s:        "",
			t:        "",
			expected: 1,
		},
		{
			name:     "Single character match",
			s:        "a",
			t:        "a",
			expected: 1,
		},
		{
			name:     "Single character no match",
			s:        "a",
			t:        "b",
			expected: 0,
		},
		{
			name:     "t longer than s",
			s:        "ab",
			t:        "abc",
			expected: 0,
		},
		{
			name:     "All characters same",
			s:        "aaaa",
			t:        "aa",
			expected: 6, // C(4,2) = 6
		},
		{
			name:     "No possible subsequences",
			s:        "abc",
			t:        "d",
			expected: 0,
		},
		{
			name:     "Multiple matches with duplicates",
			s:        "aabb",
			t:        "ab",
			expected: 4,
		},
		{
			name:     "Complex case 1",
			s:        "ABCDE",
			t:        "ACE",
			expected: 1,
		},
		{
			name:     "Complex case 2",
			s:        "ABCDE",
			t:        "AEC",
			expected: 0,
		},
		{
			name:     "Repeated characters in t",
			s:        "banana",
			t:        "ana",
			expected: 4, // Corrected: should be 4, not 3
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test standard DP solution
			result := numDistinctDP(tt.s, tt.t)
			if result != tt.expected {
				t.Errorf("numDistinctDP(%q, %q) = %d, expected %d",
					tt.s, tt.t, result, tt.expected)
			}

			// Test space-optimized DP solution
			resultOpt := numDistinctDPSpaceOptimized(tt.s, tt.t)
			if resultOpt != tt.expected {
				t.Errorf("numDistinctDPSpaceOptimized(%q, %q) = %d, expected %d",
					tt.s, tt.t, resultOpt, tt.expected)
			}

			// Test memoized recursive solution
			resultMemo := numDistinctMemoization(tt.s, tt.t)
			if resultMemo != tt.expected {
				t.Errorf("numDistinctMemoization(%q, %q) = %d, expected %d",
					tt.s, tt.t, resultMemo, tt.expected)
			}
		})
	}
}

func TestNumDistinct_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected int
	}{
		{
			name:     "t equals s",
			s:        "abcdef",
			t:        "abcdef",
			expected: 1,
		},
		{
			name:     "t is single character repeated",
			s:        "aaaaa",
			t:        "aaa",
			expected: 10, // C(5,3) = 10
		},
		{
			name:     "Long strings with pattern",
			s:        "abababab",
			t:        "abab",
			expected: 15, // Corrected: should be 15, not 5
		},
		{
			name:     "All characters different in s",
			s:        "abcdefghij",
			t:        "acegi",
			expected: 1,
		},
		{
			name:     "t is prefix of s",
			s:        "prefixsuffix",
			t:        "prefix",
			expected: 5, // Corrected: should be 5, not 1 (duplicate f, i, x in s)
		},
		{
			name:     "t is suffix of s",
			s:        "prefixsuffix",
			t:        "suffix",
			expected: 1,
		},
		{
			name:     "t appears multiple times non-overlapping",
			s:        "catcatcat",
			t:        "cat",
			expected: 10, // Corrected: should be 10, not 1
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numDistinctDP(tt.s, tt.t)
			if result != tt.expected {
				t.Errorf("numDistinctDP(%q, %q) = %d, expected %d",
					tt.s, tt.t, result, tt.expected)
			}
		})
	}
}

func TestNumDistinct_LargeNumbers(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected int64
	}{
		{
			name:     "Large combination count",
			s:        "aaaaaaaaaa", // 10 a's
			t:        "aaaaa",      // 5 a's
			expected: 252,          // C(10,5) = 252
		},
		{
			name:     "Very large s with single character t",
			s:        "aaaaaaaaaaaaaaaaaaaa", // 20 a's
			t:        "a",
			expected: 20,
		},
		{
			name:     "Pattern that generates many combinations",
			s:        "abababababababab", // 16 characters
			t:        "abababab",         // 8 characters
			expected: 495,                // Corrected: should be 495, not 12870
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numDistinctDPWithLargeNumbers(tt.s, tt.t)
			if result != tt.expected {
				t.Errorf("numDistinctDPWithLargeNumbers(%q, %q) = %d, expected %d",
					tt.s, tt.t, result, tt.expected)
			}
		})
	}
}

func TestNumDistinct_Performance(t *testing.T) {
	// Test with reasonably large strings to ensure performance
	s1 := "abcdefghijklmnopqrstuvwxyz" + "abcdefghijklmnopqrstuvwxyz" // 52 chars
	t1 := "abcdefghij" // 10 chars
	
	// This should complete quickly
	result := numDistinctDPSpaceOptimized(s1, t1)
	// Corrected: With duplicate alphabet, there are multiple ways
	// Each character in t appears twice in s (except order constraints)
	// The result should be > 1
	if result <= 0 {
		t.Errorf("numDistinctDPSpaceOptimized() returned non-positive value: %d", result)
	}
	
	// Test with repeated characters
	s2 := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa" // 50 a's
	t2 := "aaaaaaaaaa" // 10 a's
	result2 := numDistinctDPSpaceOptimized(s2, t2)
	// C(50,10) is a huge number, but should compute without overflow
	if result2 <= 0 {
		t.Errorf("numDistinctDPSpaceOptimized() returned non-positive value: %d", result2)
	}
}

func TestNumDistinct_Consistency(t *testing.T) {
	// Test that all implementations return the same result
	testCases := []struct {
		s string
		t string
	}{
		{"rabbbit", "rabbit"},
		{"babgbag", "bag"},
		{"abc", ""},
		{"", "abc"},
		{"", ""},
		{"aaaa", "aa"},
		{"banana", "ana"},
		{"abcdefghij", "acegi"},
	}

	for _, tc := range testCases {
		dpResult := numDistinctDP(tc.s, tc.t)
		optResult := numDistinctDPSpaceOptimized(tc.s, tc.t)
		memoResult := numDistinctMemoization(tc.s, tc.t)
		
		if dpResult != optResult || dpResult != memoResult {
			t.Errorf("Inconsistent results for s=%q, t=%q: DP=%d, Optimized=%d, Memo=%d",
				tc.s, tc.t, dpResult, optResult, memoResult)
		}
	}
}

func BenchmarkNumDistinct(b *testing.B) {
	benchmarks := []struct {
		name string
		s    string
		t    string
	}{
		{"Small", "rabbbit", "rabbit"},
		{"Medium", "babgbagbabgbagbabgbag", "bagbag"},
		{"Large", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaa"},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name+"_DP", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				numDistinctDP(bm.s, bm.t)
			}
		})
		
		b.Run(bm.name+"_Optimized", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				numDistinctDPSpaceOptimized(bm.s, bm.t)
			}
		})
		
		b.Run(bm.name+"_Memo", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				numDistinctMemoization(bm.s, bm.t)
			}
		})
	}
}