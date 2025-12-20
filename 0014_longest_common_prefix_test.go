package leetcode

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected string
	}{
		{"Example 1", []string{"flower", "flow", "flight"}, "fl"},
		{"Example 2", []string{"dog", "racecar", "car"}, ""},
		{"Single string", []string{"hello"}, "hello"},
		{"Two identical strings", []string{"hello", "hello"}, "hello"},
		{"No common prefix", []string{"a", "b"}, ""},
		{"Empty array", []string{}, ""},
		{"All single char", []string{"a", "a", "a"}, "a"},
		{"Common full prefix", []string{"abc", "abc", "abc"}, "abc"},
		{"Common one char", []string{"abc", "abd", "abe"}, "ab"},
		{"One char different", []string{"ab", "ac", "ad"}, "a"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LongestCommonPrefix(tt.strs)
			if result != tt.expected {
				t.Errorf("LongestCommonPrefix(%v) = %q, expected %q", tt.strs, result, tt.expected)
			}
		})
	}
}

func TestLongestCommonPrefixVertical(t *testing.T) {
	tests := []struct {
		name     string
		strs     []string
		expected string
	}{
		{"Example 1", []string{"flower", "flow", "flight"}, "fl"},
		{"Example 2", []string{"dog", "racecar", "car"}, ""},
		{"Single string", []string{"hello"}, "hello"},
		{"Two identical strings", []string{"hello", "hello"}, "hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LongestCommonPrefixVertical(tt.strs)
			if result != tt.expected {
				t.Errorf("LongestCommonPrefixVertical(%v) = %q, expected %q", tt.strs, result, tt.expected)
			}
		})
	}
}

func BenchmarkLongestCommonPrefix(b *testing.B) {
	strs := []string{"flower", "flow", "flight"}
	for i := 0; i < b.N; i++ {
		LongestCommonPrefix(strs)
	}
}

func BenchmarkLongestCommonPrefixVertical(b *testing.B) {
	strs := []string{"flower", "flow", "flight"}
	for i := 0; i < b.N; i++ {
		LongestCommonPrefixVertical(strs)
	}
}
