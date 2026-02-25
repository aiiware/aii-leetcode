package dp

import (
	"testing"
)

func TestMinimumDeleteSum(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		expected int
	}{
		// LeetCode examples
		{
			name:     "Example 1: sea and eat",
			s1:       "sea",
			s2:       "eat",
			expected: 231, // Delete 's'(115) from "sea" and 't'(116) from "eat" = 231
		},
		{
			name:     "Example 2: delete and leet",
			s1:       "delete",
			s2:       "leet",
			expected: 403, // Delete "dee" from "delete" (100+101+101=302) and 'e' from "leet" (101) = 403
		},
		
		// Edge cases
		{
			name:     "Empty strings",
			s1:       "",
			s2:       "",
			expected: 0,
		},
		{
			name:     "First string empty",
			s1:       "",
			s2:       "abc",
			expected: 294, // ASCII sum of 'a'(97) + 'b'(98) + 'c'(99) = 294
		},
		{
			name:     "Second string empty",
			s1:       "abc",
			s2:       "",
			expected: 294, // ASCII sum of 'a'(97) + 'b'(98) + 'c'(99) = 294
		},
		{
			name:     "Identical strings",
			s1:       "hello",
			s2:       "hello",
			expected: 0, // No deletions needed
		},
		{
			name:     "Single character match",
			s1:       "a",
			s2:       "a",
			expected: 0,
		},
		{
			name:     "Single character mismatch",
			s1:       "a",
			s2:       "b",
			expected: 195, // Delete 'a'(97) + 'b'(98) = 195
		},
		
		// More complex cases - I'll compute these correctly
		{
			name:     "Partial match at beginning",
			s1:       "abcde",
			s2:       "abfgh",
			expected: 609, // Keep "ab", delete 'c'(99)+'d'(100)+'e'(101)=300 from s1, delete 'f'(102)+'g'(103)+'h'(104)=309 from s2, total=609
		},
		{
			name:     "Partial match at end",
			s1:       "xyzabc",
			s2:       "defabc",
			expected: 666, // Keep "abc", delete 'x'(120)+'y'(121)+'z'(122)=363 from s1, delete 'd'(100)+'e'(101)+'f'(102)=303 from s2, total=666
		},
		{
			name:     "Interleaved matches",
			s1:       "ace",
			s2:       "bdf",
			expected: 597, // No matches, delete all: 'a'(97)+'c'(99)+'e'(101)=297 + 'b'(98)+'d'(100)+'f'(102)=300 = 597
		},
		{
			name:     "All characters different",
			s1:       "abc",
			s2:       "def",
			expected: 597, // Delete all: 'a'(97)+'b'(98)+'c'(99)=294 + 'd'(100)+'e'(101)+'f'(102)=303 = 597
		},
		{
			name:     "One string subset of another",
			s1:       "abc",
			s2:       "abcdef",
			expected: 303, // Keep "abc", delete 'd'(100)+'e'(101)+'f'(102)=303 from s2
		},
		{
			name:     "Reverse strings",
			s1:       "abc",
			s2:       "cba",
			expected: 390, // Keep 'c', delete 'a'(97)+'b'(98)=195 from s1, delete 'b'(98)+'a'(97)=195 from s2, total=390
		},
		{
			name:     "Case sensitivity test",
			s1:       "Hello",
			s2:       "hello",
			expected: 176, // Keep "ello", delete 'H'(72) from s1, delete 'h'(104) from s2 = 72+104=176
		},
		{
			name:     "Numbers in strings",
			s1:       "a1b2",
			s2:       "a2b1",
			expected: 198, // Keep "a" and "b", delete '1'(49)+'2'(50)=99 from s1, delete '2'(50)+'1'(49)=99 from s2, total=198
		},
		{
			name:     "Special characters",
			s1:       "a!@#",
			s2:       "b!@#",
			expected: 195, // Keep "!@#", delete 'a'(97) from s1, delete 'b'(98) from s2 = 195
		},
		{
			name:     "Longer test case 1",
			s1:       "algorithm",
			s2:       "altruistic",
			expected: 979, // Let the algorithm compute this (verified)
		},
		{
			name:     "Longer test case 2",
			s1:       "dynamic",
			s2:       "programming",
			expected: 1306, // Let the algorithm compute this (verified)
		},
		{
			name:     "Unicode characters (ASCII only in problem)",
			s1:       "abc",
			s2:       "def",
			expected: 597, // Already covered
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minimumDeleteSum(tt.s1, tt.s2)
			if result != tt.expected {
				t.Errorf("minimumDeleteSum(%q, %q) = %d, expected %d", tt.s1, tt.s2, result, tt.expected)
			}
		})
	}
}

func TestMinimumDeleteSumOptimized(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		expected int
	}{
		// LeetCode examples
		{
			name:     "Example 1: sea and eat",
			s1:       "sea",
			s2:       "eat",
			expected: 231,
		},
		{
			name:     "Example 2: delete and leet",
			s1:       "delete",
			s2:       "leet",
			expected: 403,
		},
		
		// Edge cases
		{
			name:     "Empty strings",
			s1:       "",
			s2:       "",
			expected: 0,
		},
		{
			name:     "First string empty",
			s1:       "",
			s2:       "abc",
			expected: 294,
		},
		{
			name:     "Second string empty",
			s1:       "abc",
			s2:       "",
			expected: 294,
		},
		{
			name:     "Identical strings",
			s1:       "hello",
			s2:       "hello",
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minimumDeleteSumOptimized(tt.s1, tt.s2)
			if result != tt.expected {
				t.Errorf("minimumDeleteSumOptimized(%q, %q) = %d, expected %d", tt.s1, tt.s2, result, tt.expected)
			}
		})
	}
}

func TestMinimumDeleteSumConsistency(t *testing.T) {
	// Test that both implementations produce the same results
	testCases := []struct {
		s1 string
		s2 string
	}{
		{"sea", "eat"},
		{"delete", "leet"},
		{"", ""},
		{"", "abc"},
		{"abc", ""},
		{"hello", "hello"},
		{"a", "b"},
		{"abc", "def"},
		{"algorithm", "altruistic"},
		{"dynamic", "programming"},
		{"abcde", "abfgh"},
		{"xyzabc", "defabc"},
		{"Hello", "hello"},
		{"a1b2", "a2b1"},
		{"abc", "cba"},
	}

	for _, tc := range testCases {
		t.Run(tc.s1+"_"+tc.s2, func(t *testing.T) {
			result1 := minimumDeleteSum(tc.s1, tc.s2)
			result2 := minimumDeleteSumOptimized(tc.s1, tc.s2)
			
			if result1 != result2 {
				t.Errorf("Inconsistent results for (%q, %q): regular=%d, optimized=%d", 
					tc.s1, tc.s2, result1, result2)
			}
		})
	}
}

func BenchmarkMinimumDeleteSum(b *testing.B) {
	testCases := []struct {
		name string
		s1   string
		s2   string
	}{
		{"Small", "sea", "eat"},
		{"Medium", "delete", "leet"},
		{"Large", "algorithm", "altruistic"},
		{"XLarge", "dynamicprogrammingexample", "exampleofdynamicprogramming"},
	}

	for _, tc := range testCases {
		b.Run("Regular_"+tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				minimumDeleteSum(tc.s1, tc.s2)
			}
		})
		
		b.Run("Optimized_"+tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				minimumDeleteSumOptimized(tc.s1, tc.s2)
			}
		})
	}
}