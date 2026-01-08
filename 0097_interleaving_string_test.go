package leetcode

import (
	"fmt"
	"testing"
)

func TestIsInterleave(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		s3       string
		expected bool
	}{
		{
			name:     "Example 1",
			s1:       "aabcc",
			s2:       "dbbca",
			s3:       "aadbbcbcac",
			expected: true,
		},
		{
			name:     "Example 2",
			s1:       "aabcc",
			s2:       "dbbca",
			s3:       "aadbbbaccc",
			expected: false,
		},
		{
			name:     "Example 3",
			s1:       "",
			s2:       "",
			s3:       "",
			expected: true,
		},
		{
			name:     "All empty",
			s1:       "",
			s2:       "",
			s3:       "",
			expected: true,
		},
		{
			name:     "s1 empty, s2 equals s3",
			s1:       "",
			s2:       "abc",
			s3:       "abc",
			expected: true,
		},
		{
			name:     "s2 empty, s1 equals s3",
			s1:       "abc",
			s2:       "",
			s3:       "abc",
			expected: true,
		},
		{
			name:     "Simple interleave",
			s1:       "ab",
			s2:       "cd",
			s3:       "acbd",
			expected: true,
		},
		{
			name:     "Simple not interleave",
			s1:       "ab",
			s2:       "cd",
			s3:       "adcb",
			expected: false,
		},
		{
			name:     "Length mismatch",
			s1:       "abc",
			s2:       "def",
			s3:       "abcdefg", // too long
			expected: false,
		},
		{
			name:     "Length mismatch 2",
			s1:       "abc",
			s2:       "def",
			s3:       "abde", // too short
			expected: false,
		},
		{
			name:     "Same characters different order",
			s1:       "aabc",
			s2:       "abad",
			s3:       "aabadabc",
			expected: true,
		},
		{
			name:     "Complex true case",
			s1:       "xxy",
			s2:       "xxz",
			s3:       "xxxyxz",
			expected: true,
		},
		{
			name:     "Complex false case",
			s1:       "xxy",
			s2:       "xxz",
			s3:       "xxxyzx",
			expected: false,
		},
		{
			name:     "All same characters",
			s1:       "aaa",
			s2:       "aaa",
			s3:       "aaaaaa",
			expected: true,
		},
		{
			name:     "Single characters",
			s1:       "a",
			s2:       "b",
			s3:       "ab",
			expected: true,
		},
		{
			name:     "Single characters reversed",
			s1:       "a",
			s2:       "b",
			s3:       "ba",
			expected: true,
		},
		{
			name:     "Single character mismatch",
			s1:       "a",
			s2:       "b",
			s3:       "ac",
			expected: false,
		},
		{
			name:     "Long strings true",
			s1:       "abcdefghij",
			s2:       "klmnopqrst",
			s3:       "akblcmdneofpgqhrisjt",
			expected: true,
		},
		{
			name:     "With duplicates",
			s1:       "aabbcc",
			s2:       "ddeeff",
			s3:       "aabbccddeeff",
			expected: true,
		},
		{
			name:     "Interleaved duplicates",
			s1:       "aabb",
			s2:       "ccdd",
			s3:       "acbdacbd",
			expected: false, // Fixed: this is not a valid interleaving
		},
		{
			name:     "LeetCode test case 1",
			s1:       "db",
			s2:       "b",
			s3:       "cbb",
			expected: false,
		},
		{
			name:     "LeetCode test case 2",
			s1:       "a",
			s2:       "b",
			s3:       "a",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsInterleave(tt.s1, tt.s2, tt.s3)
			if result != tt.expected {
				t.Errorf("IsInterleave(%q, %q, %q) = %v, expected %v",
					tt.s1, tt.s2, tt.s3, result, tt.expected)
			}
		})
	}
}

func TestAllIsInterleaveImplementations(t *testing.T) {
	testCases := []struct {
		name string
		s1   string
		s2   string
		s3   string
	}{
		{
			name: "Example 1",
			s1:   "aabcc",
			s2:   "dbbca",
			s3:   "aadbbcbcac",
		},
		{
			name: "Example 2",
			s1:   "aabcc",
			s2:   "dbbca",
			s3:   "aadbbbaccc",
		},
		{
			name: "Simple true",
			s1:   "ab",
			s2:   "cd",
			s3:   "acbd",
		},
		{
			name: "Simple false",
			s1:   "ab",
			s2:   "cd",
			s3:   "adcb",
		},
		{
			name: "Empty strings",
			s1:   "",
			s2:   "",
			s3:   "",
		},
		{
			name: "Single characters",
			s1:   "a",
			s2:   "b",
			s3:   "ab",
		},
	}

	implementations := []struct {
		name string
		fn   func(string, string, string) bool
	}{
		{"isInterleave", isInterleave},
		{"isInterleaveOptimized", isInterleaveOptimized},
		{"isInterleaveDFS", isInterleaveDFS},
		{"isInterleaveBFS", isInterleaveBFS},
		{"isInterleaveDP2", isInterleaveDP2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := IsInterleave(tc.s1, tc.s2, tc.s3)

			for _, impl := range implementations {
				t.Run(impl.name, func(t *testing.T) {
					result := impl.fn(tc.s1, tc.s2, tc.s3)
					if result != expected {
						t.Errorf("%s(%q, %q, %q) = %v, expected %v",
							impl.name, tc.s1, tc.s2, tc.s3, result, expected)
					}
				})
			}
		})
	}
}

func TestIsInterleaveEdgeCases(t *testing.T) {
	t.Run("All empty strings", func(t *testing.T) {
		if !IsInterleave("", "", "") {
			t.Error("IsInterleave(\"\", \"\", \"\") should be true")
		}
	})

	t.Run("s3 empty but s1 and s2 not empty", func(t *testing.T) {
		if IsInterleave("a", "b", "") {
			t.Error("IsInterleave(\"a\", \"b\", \"\") should be false")
		}
	})

	t.Run("s1 and s2 empty but s3 not empty", func(t *testing.T) {
		if IsInterleave("", "", "a") {
			t.Error("IsInterleave(\"\", \"\", \"a\") should be false")
		}
	})

	t.Run("Length mismatch", func(t *testing.T) {
		testCases := []struct {
			s1, s2, s3 string
		}{
			{"abc", "def", "abcdefg"}, // s3 too long
			{"abc", "def", "abde"},    // s3 too short
			{"", "abc", "ab"},         // s3 shorter than s2
			{"abc", "", "ab"},         // s3 shorter than s1
		}
		for _, tc := range testCases {
			if IsInterleave(tc.s1, tc.s2, tc.s3) {
				t.Errorf("IsInterleave(%q, %q, %q) should be false (length mismatch)",
					tc.s1, tc.s2, tc.s3)
			}
		}
	})

	t.Run("Single character strings", func(t *testing.T) {
		testCases := []struct {
			s1, s2, s3 string
			expected   bool
		}{
			{"a", "b", "ab", true},
			{"a", "b", "ba", true},
			{"a", "b", "aa", false},
			{"a", "b", "bb", false},
			{"a", "b", "c", false},
			{"a", "", "a", true},
			{"", "b", "b", true},
		}
		for _, tc := range testCases {
			result := IsInterleave(tc.s1, tc.s2, tc.s3)
			if result != tc.expected {
				t.Errorf("IsInterleave(%q, %q, %q) = %v, expected %v",
					tc.s1, tc.s2, tc.s3, result, tc.expected)
			}
		}
	})

	t.Run("Large strings", func(t *testing.T) {
		// Create large strings
		s1 := ""
		s2 := ""
		s3 := ""
		for i := 0; i < 50; i++ {
			s1 += "a"
			s2 += "b"
			s3 += "ab"
		}
		
		if !IsInterleave(s1, s2, s3) {
			t.Error("Large interleaved strings should return true")
		}
		
		// Make s3 invalid
		s3 = s3[:len(s3)-1] + "c"
		if IsInterleave(s1, s2, s3) {
			t.Error("Invalid large string should return false")
		}
	})

	t.Run("Character frequency mismatch", func(t *testing.T) {
		// Even if lengths match, character frequencies must match
		s1 := "aabb"
		s2 := "ccdd"
		s3 := "aaccbbdd" // Valid interleaving
		if !IsInterleave(s1, s2, s3) {
			t.Error("Valid interleaving with matching frequencies should return true")
		}
		
		s3 = "aaccbbde" // 'e' not in s1 or s2
		if IsInterleave(s1, s2, s3) {
			t.Error("Invalid character should return false")
		}
		
		s3 = "aaccbbddd" // Extra 'd'
		if IsInterleave(s1, s2, s3) {
			t.Error("Extra character should return false")
		}
	})
}

func TestIsInterleaveProperties(t *testing.T) {
	// Property-based testing
	implementations := []struct {
		name string
		fn   func(string, string, string) bool
	}{
		{"isInterleave", isInterleave},
		{"isInterleaveOptimized", isInterleaveOptimized},
		{"isInterleaveDFS", isInterleaveDFS},
		{"isInterleaveBFS", isInterleaveBFS},
		{"isInterleaveDP2", isInterleaveDP2},
	}

	testCases := []struct {
		name string
		s1   string
		s2   string
		s3   string
	}{
		{"Simple true", "ab", "cd", "acbd"},
		{"Simple false", "ab", "cd", "adcb"},
		{"Empty", "", "", ""},
		{"s1 empty", "", "abc", "abc"},
		{"s2 empty", "abc", "", "abc"},
		{"All same", "aaa", "aaa", "aaaaaa"},
		{"Alternating", "abc", "def", "adbecf"},
	}

	for _, impl := range implementations {
		t.Run(impl.name+"_properties", func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					result := impl.fn(tc.s1, tc.s2, tc.s3)

					// Property 1: If result is true, lengths must match
					if result && len(tc.s1)+len(tc.s2) != len(tc.s3) {
						t.Errorf("True result but length mismatch: len(%q)+len(%q)=%d, len(%q)=%d",
							tc.s1, tc.s2, len(tc.s1)+len(tc.s2), tc.s3, len(tc.s3))
					}

					// Property 2: If result is true, character multiset must match
					if result {
						count1 := countChars(tc.s1)
						count2 := countChars(tc.s2)
						count3 := countChars(tc.s3)
						
						for char, count := range count3 {
							expected := count1[char] + count2[char]
							if count != expected {
								t.Errorf("Character %c count mismatch: s3 has %d, s1+s2 has %d",
									char, count, expected)
							}
						}
					}

					// Property 3: If s3 is concatenation of s1 and s2 (in either order), result should be true
					if tc.s3 == tc.s1+tc.s2 || tc.s3 == tc.s2+tc.s1 {
						if !result {
							t.Errorf("s3 is concatenation of s1 and s2, should return true")
						}
					}
				})
			}
		})
	}
}

func BenchmarkIsInterleave(b *testing.B) {
	// Test cases of different sizes
	testCases := []struct {
		name string
		s1   string
		s2   string
		s3   string
	}{
		{
			name: "Small",
			s1:   "aabcc",
			s2:   "dbbca",
			s3:   "aadbbcbcac",
		},
		{
			name: "Medium",
			s1:   "abcdefghij",
			s2:   "klmnopqrst",
			s3:   "akblcmdneofpgqhrisjt",
		},
		{
			name: "Large",
			s1:   repeat97("abc", 10),
			s2:   repeat97("def", 10),
			s3:   repeat97("adbecf", 10),
		},
		{
			name: "All same characters",
			s1:   repeat97("a", 20),
			s2:   repeat97("a", 20),
			s3:   repeat97("a", 40),
		},
	}

	implementations := []struct {
		name string
		fn   func(string, string, string) bool
	}{
		{"isInterleave", isInterleave},
		{"isInterleaveOptimized", isInterleaveOptimized},
		{"isInterleaveDFS", isInterleaveDFS},
		{"isInterleaveBFS", isInterleaveBFS},
		{"isInterleaveDP2", isInterleaveDP2},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for _, impl := range implementations {
				b.Run(impl.name, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						impl.fn(tc.s1, tc.s2, tc.s3)
					}
				})
			}
		})
	}
}

func BenchmarkIsInterleaveWorstCase(b *testing.B) {
	// Worst case: all characters are the same, many possible paths
	s1 := repeat97("a", 50)
	s2 := repeat97("a", 50)
	s3 := repeat97("a", 100)

	b.ResetTimer()

	b.Run("isInterleave", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isInterleave(s1, s2, s3)
		}
	})

	b.Run("isInterleaveOptimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isInterleaveOptimized(s1, s2, s3)
		}
	})

	b.Run("isInterleaveDFS", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isInterleaveDFS(s1, s2, s3)
		}
	})

	b.Run("isInterleaveBFS", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isInterleaveBFS(s1, s2, s3)
		}
	})
}

func BenchmarkIsInterleaveRecursive(b *testing.B) {
	// Test recursive separately (it's exponential)
	s1 := "abc"
	s2 := "def"
	s3 := "adbecf"

	b.ResetTimer()
	b.Run("isInterleaveRecursive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			isInterleaveRecursive(s1, s2, s3)
		}
	})
}

// Helper functions

func countChars(s string) map[byte]int {
	count := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}
	return count
}

func repeat97(s string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}

// Test that all implementations agree on random test cases
func TestIsInterleaveRandom(t *testing.T) {
	// Generate random test cases and verify all implementations agree
	implementations := []struct {
		name string
		fn   func(string, string, string) bool
	}{
		{"isInterleave", isInterleave},
		{"isInterleaveOptimized", isInterleaveOptimized},
		{"isInterleaveDFS", isInterleaveDFS},
		{"isInterleaveBFS", isInterleaveBFS},
		{"isInterleaveDP2", isInterleaveDP2},
		{"isInterleaveRecursive", isInterleaveRecursive},
	}

	// Simple test cases
	testCases := []struct {
		s1, s2, s3 string
	}{
		{"a", "b", "ab"},
		{"a", "b", "ba"},
		{"aa", "bb", "abab"},
		{"aa", "bb", "aabb"},
		{"aa", "bb", "bbaa"},
		{"aab", "aac", "aaabac"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("s1=%s,s2=%s,s3=%s", tc.s1, tc.s2, tc.s3), func(t *testing.T) {
			var firstResult bool
			var firstName string
			
			for i, impl := range implementations {
				result := impl.fn(tc.s1, tc.s2, tc.s3)
				if i == 0 {
					firstResult = result
					firstName = impl.name
				} else if result != firstResult {
					t.Errorf("Implementation mismatch: %s=%v, %s=%v",
						firstName, firstResult, impl.name, result)
				}
			}
		})
	}
}