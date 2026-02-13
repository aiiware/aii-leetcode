package sliding_window

import (
	"testing"
)

func TestCheckInclusion(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		expected bool
	}{
		{
			name:     "Example 1 from LeetCode",
			s1:       "ab",
			s2:       "eidbaooo",
			expected: true,
		},
		{
			name:     "Example 2 from LeetCode",
			s1:       "ab",
			s2:       "eidboaoo",
			expected: false,
		},
		{
			name:     "s1 equals s2",
			s1:       "abc",
			s2:       "abc",
			expected: true,
		},
		{
			name:     "s1 longer than s2",
			s1:       "abcd",
			s2:       "abc",
			expected: false,
		},
		{
			name:     "Permutation at beginning",
			s1:       "abc",
			s2:       "bcaxyz",
			expected: true,
		},
		{
			name:     "Permutation at end",
			s1:       "abc",
			s2:       "xyzcab",
			expected: true,
		},
		{
			name:     "No permutation",
			s1:       "abc",
			s2:       "defghij",
			expected: false,
		},
		{
			name:     "Single character s1",
			s1:       "a",
			s2:       "abcdef",
			expected: true,
		},
		{
			name:     "Single character s1 not found",
			s1:       "z",
			s2:       "abcdef",
			expected: false,
		},
		{
			name:     "Empty s1",
			s1:       "",
			s2:       "abc",
			expected: true, // Empty string is a permutation of any string
		},
		{
			name:     "Empty s2",
			s1:       "abc",
			s2:       "",
			expected: false,
		},
		{
			name:     "Both empty",
			s1:       "",
			s2:       "",
			expected: true,
		},
		{
			name:     "Duplicate characters in s1",
			s1:       "aab",
			s2:       "eidbaaoo",
			expected: true, // "baa" contains permutation "aab"
		},
		{
			name:     "All same characters",
			s1:       "aaa",
			s2:       "bbaaaabb",
			expected: true,
		},
		{
			name:     "Overlapping windows",
			s1:       "ab",
			s2:       "ababab",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := checkInclusion(tt.s1, tt.s2)
			if result != tt.expected {
				t.Errorf("checkInclusion(%q, %q) = %v, expected %v", tt.s1, tt.s2, result, tt.expected)
			}

			// Also test optimized version for comparison
			optimizedResult := checkInclusionOptimized(tt.s1, tt.s2)
			if result != optimizedResult {
				t.Errorf("Standard and optimized versions don't match: standard=%v, optimized=%v", result, optimizedResult)
			}
		})
	}
}

func TestCheckInclusionOptimized(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		expected bool
	}{
		{
			name:     "Example 1 from LeetCode",
			s1:       "ab",
			s2:       "eidbaooo",
			expected: true,
		},
		{
			name:     "Example 2 from LeetCode",
			s1:       "ab",
			s2:       "eidboaoo",
			expected: false,
		},
		{
			name:     "s1 equals s2",
			s1:       "abc",
			s2:       "abc",
			expected: true,
		},
		{
			name:     "s1 longer than s2",
			s1:       "abcd",
			s2:       "abc",
			expected: false,
		},
		{
			name:     "Permutation at beginning",
			s1:       "abc",
			s2:       "bcaxyz",
			expected: true,
		},
		{
			name:     "Permutation at end",
			s1:       "abc",
			s2:       "xyzcab",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := checkInclusionOptimized(tt.s1, tt.s2)
			if result != tt.expected {
				t.Errorf("checkInclusionOptimized(%q, %q) = %v, expected %v", tt.s1, tt.s2, result, tt.expected)
			}
		})
	}
}

func BenchmarkCheckInclusion(b *testing.B) {
	// Create test cases
	testCases := []struct {
		name string
		s1   string
		s2   string
	}{
		{
			name: "Small strings",
			s1:   "ab",
			s2:   "eidbaooo",
		},
		{
			name: "Medium strings",
			s1:   "abc",
			s2:   "xyzabcqwertyuiopasdfghjklzxcvbnm",
		},
		{
			name: "Large s2 with small s1",
			s2: func() string {
				var result string
				for i := 0; i < 10000; i++ {
					result += string('a' + byte(i%26))
				}
				return result
			}(),
			s1: "xyz",
		},
		{
			name: "Both strings large",
			s1: func() string {
				var result string
				for i := 0; i < 100; i++ {
					result += string('a' + byte(i%26))
				}
				return result
			}(),
			s2: func() string {
				var result string
				for i := 0; i < 5000; i++ {
					result += string('a' + byte(i%26))
				}
				return result
			}(),
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				checkInclusion(tc.s1, tc.s2)
			}
		})
	}
}

func BenchmarkCheckInclusionOptimized(b *testing.B) {
	// Create test cases
	testCases := []struct {
		name string
		s1   string
		s2   string
	}{
		{
			name: "Small strings",
			s1:   "ab",
			s2:   "eidbaooo",
		},
		{
			name: "Medium strings",
			s1:   "abc",
			s2:   "xyzabcqwertyuiopasdfghjklzxcvbnm",
		},
		{
			name: "Large s2 with small s1",
			s2: func() string {
				var result string
				for i := 0; i < 10000; i++ {
					result += string('a' + byte(i%26))
				}
				return result
			}(),
			s1: "xyz",
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				checkInclusionOptimized(tc.s1, tc.s2)
			}
		})
	}
}