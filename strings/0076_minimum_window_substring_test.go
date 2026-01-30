package strings

import (
	"testing"
)

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected string
	}{
		{
			name:     "Example 1 from LeetCode",
			s:        "ADOBECODEBANC",
			t:        "ABC",
			expected: "BANC",
		},
		{
			name:     "Example 2 from LeetCode",
			s:        "a",
			t:        "a",
			expected: "a",
		},
		{
			name:     "Example 3 from LeetCode",
			s:        "a",
			t:        "aa",
			expected: "",
		},
		{
			name:     "Empty s string",
			s:        "",
			t:        "a",
			expected: "",
		},
		{
			name:     "Empty t string",
			s:        "abc",
			t:        "",
			expected: "",
		},
		{
			name:     "Both strings empty",
			s:        "",
			t:        "",
			expected: "",
		},
		{
			name:     "s shorter than t",
			s:        "ab",
			t:        "abc",
			expected: "",
		},
		{
			name:     "Exact match",
			s:        "abc",
			t:        "abc",
			expected: "abc",
		},
		{
			name:     "Multiple possible windows, choose smallest",
			s:        "ADOBECODEBANCABC",
			t:        "ABC",
			expected: "CAB", // Both "ABC" and "CAB" are valid and same length
		},
		{
			name:     "t with duplicate characters",
			s:        "aa",
			t:        "aa",
			expected: "aa",
		},
		{
			name:     "t with duplicate characters not satisfied",
			s:        "a",
			t:        "aa",
			expected: "",
		},
		{
			name:     "Case sensitive",
			s:        "aBc",
			t:        "Bc",
			expected: "Bc",
		},
		{
			name:     "Window at beginning",
			s:        "ABCxxxx",
			t:        "ABC",
			expected: "ABC",
		},
		{
			name:     "Window at end",
			s:        "xxxxABC",
			t:        "ABC",
			expected: "ABC",
		},
		{
			name:     "t characters scattered",
			s:        "AXBYCZ",
			t:        "ABC",
			expected: "AXBYC",
		},
		{
			name:     "t characters in different order",
			s:        "CBA",
			t:        "ABC",
			expected: "CBA",
		},
		{
			name:     "Large s with small t",
			s:        "this is a test string with multiple a's and b's and c's",
			t:        "abc",
			expected: "b's and c", // Shorter than "a's and b's and c"
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinWindow(tt.s, tt.t)
			if result != tt.expected {
				t.Errorf("MinWindow(%q, %q) = %q, expected %q", tt.s, tt.t, result, tt.expected)
			}
		})
	}
}

func TestMinWindowOptimized(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected string
	}{
		{
			name:     "Example 1 from LeetCode",
			s:        "ADOBECODEBANC",
			t:        "ABC",
			expected: "BANC",
		},
		{
			name:     "Example 2 from LeetCode",
			s:        "a",
			t:        "a",
			expected: "a",
		},
		{
			name:     "Example 3 from LeetCode",
			s:        "a",
			t:        "aa",
			expected: "",
		},
		{
			name:     "Empty s string",
			s:        "",
			t:        "a",
			expected: "",
		},
		{
			name:     "Exact match",
			s:        "abc",
			t:        "abc",
			expected: "abc",
		},
		{
			name:     "Multiple possible windows, choose smallest",
			s:        "ADOBECODEBANCABC",
			t:        "ABC",
			expected: "CAB", // Both "ABC" and "CAB" are valid and same length
		},
		{
			name:     "t with duplicate characters",
			s:        "aa",
			t:        "aa",
			expected: "aa",
		},
		{
			name:     "Case sensitive",
			s:        "aBc",
			t:        "Bc",
			expected: "Bc",
		},
		{
			name:     "Window at beginning",
			s:        "ABCxxxx",
			t:        "ABC",
			expected: "ABC",
		},
		{
			name:     "Window at end",
			s:        "xxxxABC",
			t:        "ABC",
			expected: "ABC",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinWindowOptimized(tt.s, tt.t)
			if result != tt.expected {
				t.Errorf("MinWindowOptimized(%q, %q) = %q, expected %q", tt.s, tt.t, result, tt.expected)
			}
		})
	}
}

func TestMinWindowSimplified(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected string
	}{
		{
			name:     "Example 1 from LeetCode",
			s:        "ADOBECODEBANC",
			t:        "ABC",
			expected: "BANC",
		},
		{
			name:     "Example 2 from LeetCode",
			s:        "a",
			t:        "a",
			expected: "a",
		},
		{
			name:     "Example 3 from LeetCode",
			s:        "a",
			t:        "aa",
			expected: "",
		},
		{
			name:     "Empty s string",
			s:        "",
			t:        "a",
			expected: "",
		},
		{
			name:     "Exact match",
			s:        "abc",
			t:        "abc",
			expected: "abc",
		},
		{
			name:     "Multiple possible windows, choose smallest",
			s:        "ADOBECODEBANCABC",
			t:        "ABC",
			expected: "CAB", // Both "ABC" and "CAB" are valid and same length
		},
		{
			name:     "t with duplicate characters",
			s:        "aa",
			t:        "aa",
			expected: "aa",
		},
		{
			name:     "Case sensitive",
			s:        "aBc",
			t:        "Bc",
			expected: "Bc",
		},
		{
			name:     "Window at beginning",
			s:        "ABCxxxx",
			t:        "ABC",
			expected: "ABC",
		},
		{
			name:     "Window at end",
			s:        "xxxxABC",
			t:        "ABC",
			expected: "ABC",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MinWindowSimplified(tt.s, tt.t)
			if result != tt.expected {
				t.Errorf("MinWindowSimplified(%q, %q) = %q, expected %q", tt.s, tt.t, result, tt.expected)
			}
		})
	}
}

func BenchmarkMinWindow(b *testing.B) {
	testCases := []struct {
		name string
		s    string
		t    string
	}{
		{
			name: "Small strings",
			s:    "ADOBECODEBANC",
			t:    "ABC",
		},
		{
			name: "Medium strings",
			s:    "this is a test string with multiple a's and b's and c's and other characters to make it longer",
			t:    "abc",
		},
		{
			name: "Large s with small t",
			s: func() string {
				// Create a 1000 character string
				var result string
				for i := 0; i < 100; i++ {
					result += "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
				}
				return result
			}(),
			t: "xyz",
		},
		{
			name: "Both strings large",
			s: func() string {
				var result string
				for i := 0; i < 50; i++ {
					result += "abcdefghijklmnopqrstuvwxyz"
				}
				return result
			}(),
			t: func() string {
				var result string
				for i := 0; i < 10; i++ {
					result += "abcdefghij"
				}
				return result
			}(),
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MinWindow(tc.s, tc.t)
			}
		})
	}
}

func BenchmarkMinWindowOptimized(b *testing.B) {
	testCases := []struct {
		name string
		s    string
		t    string
	}{
		{
			name: "Small strings",
			s:    "ADOBECODEBANC",
			t:    "ABC",
		},
		{
			name: "Medium strings",
			s:    "this is a test string with multiple a's and b's and c's and other characters to make it longer",
			t:    "abc",
		},
		{
			name: "Large s with small t",
			s: func() string {
				var result string
				for i := 0; i < 100; i++ {
					result += "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
				}
				return result
			}(),
			t: "xyz",
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MinWindowOptimized(tc.s, tc.t)
			}
		})
	}
}

func BenchmarkMinWindowSimplified(b *testing.B) {
	testCases := []struct {
		name string
		s    string
		t    string
	}{
		{
			name: "Small strings",
			s:    "ADOBECODEBANC",
			t:    "ABC",
		},
		{
			name: "Medium strings",
			s:    "this is a test string with multiple a's and b's and c's and other characters to make it longer",
			t:    "abc",
		},
		{
			name: "Large s with small t",
			s: func() string {
				var result string
				for i := 0; i < 100; i++ {
					result += "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
				}
				return result
			}(),
			t: "xyz",
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MinWindowSimplified(tc.s, tc.t)
			}
		})
	}
}