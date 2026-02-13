package sliding_window

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		p        string
		expected []int
	}{
		{
			name:     "Example 1 from LeetCode",
			s:        "cbaebabacd",
			p:        "abc",
			expected: []int{0, 6},
		},
		{
			name:     "Example 2 from LeetCode",
			s:        "abab",
			p:        "ab",
			expected: []int{0, 1, 2},
		},
		{
			name:     "No anagrams found",
			s:        "abcdefg",
			p:        "xyz",
			expected: []int{},
		},
		{
			name:     "p longer than s",
			s:        "abc",
			p:        "abcd",
			expected: []int{},
		},
		{
			name:     "Single character p",
			s:        "aaaa",
			p:        "a",
			expected: []int{0, 1, 2, 3},
		},
		{
			name:     "Multiple same anagrams",
			s:        "ababab",
			p:        "ab",
			expected: []int{0, 1, 2, 3, 4},
		},
		{
			name:     "Empty strings",
			s:        "",
			p:        "",
			expected: []int{},
		},
		{
			name:     "Empty p",
			s:        "abc",
			p:        "",
			expected: []int{},
		},
		{
			name:     "Empty s",
			s:        "",
			p:        "abc",
			expected: []int{},
		},
		{
			name:     "All characters same",
			s:        "aaaaaa",
			p:        "aaa",
			expected: []int{0, 1, 2, 3},
		},
		{
			name:     "p equals s",
			s:        "abc",
			p:        "abc",
			expected: []int{0},
		},
		{
			name:     "Overlapping anagrams",
			s:        "abacbabc",
			p:        "abc",
			expected: []int{1, 2, 3, 5},
		},
		{
			name:     "Case with duplicate letters in p",
			s:        "baa",
			p:        "aa",
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findAnagrams(tt.s, tt.p)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("findAnagrams(%q, %q) = %v, expected %v", tt.s, tt.p, result, tt.expected)
			}

			// Also test map version for comparison
			mapResult := findAnagramsMap(tt.s, tt.p)
			if !reflect.DeepEqual(result, mapResult) {
				t.Errorf("Array and map versions don't match: array=%v, map=%v", result, mapResult)
			}
		})
	}
}

func TestFindAnagramsMap(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		p        string
		expected []int
	}{
		{
			name:     "Example 1 from LeetCode",
			s:        "cbaebabacd",
			p:        "abc",
			expected: []int{0, 6},
		},
		{
			name:     "Example 2 from LeetCode",
			s:        "abab",
			p:        "ab",
			expected: []int{0, 1, 2},
		},
		{
			name:     "No anagrams found",
			s:        "abcdefg",
			p:        "xyz",
			expected: []int{},
		},
		{
			name:     "p longer than s",
			s:        "abc",
			p:        "abcd",
			expected: []int{},
		},
		{
			name:     "Single character p",
			s:        "aaaa",
			p:        "a",
			expected: []int{0, 1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findAnagramsMap(tt.s, tt.p)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("findAnagramsMap(%q, %q) = %v, expected %v", tt.s, tt.p, result, tt.expected)
			}
		})
	}
}

func BenchmarkFindAnagrams(b *testing.B) {
	// Create test cases
	testCases := []struct {
		name string
		s    string
		p    string
	}{
		{
			name: "Small strings",
			s:    "cbaebabacd",
			p:    "abc",
		},
		{
			name: "Medium strings",
			s:    "abababababababababababab",
			p:    "ab",
		},
		{
			name: "Large s with small p",
			s: func() string {
				var result string
				for i := 0; i < 10000; i++ {
					result += string('a' + byte(i%26))
				}
				return result
			}(),
			p: "abc",
		},
		{
			name: "Both strings large",
			s: func() string {
				var result string
				for i := 0; i < 5000; i++ {
					result += string('a' + byte(i%26))
				}
				return result
			}(),
			p: func() string {
				var result string
				for i := 0; i < 100; i++ {
					result += string('a' + byte(i%26))
				}
				return result
			}(),
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				findAnagrams(tc.s, tc.p)
			}
		})
	}
}

func BenchmarkFindAnagramsMap(b *testing.B) {
	// Create test cases
	testCases := []struct {
		name string
		s    string
		p    string
	}{
		{
			name: "Small strings",
			s:    "cbaebabacd",
			p:    "abc",
		},
		{
			name: "Medium strings",
			s:    "abababababababababababab",
			p:    "ab",
		},
		{
			name: "Large s with small p",
			s: func() string {
				var result string
				for i := 0; i < 10000; i++ {
					result += string('a' + byte(i%26))
				}
				return result
			}(),
			p: "abc",
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				findAnagramsMap(tc.s, tc.p)
			}
		})
	}
}