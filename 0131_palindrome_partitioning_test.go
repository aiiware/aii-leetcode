package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestPartition(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected [][]string
	}{
		{
			name: "Example 1",
			s:    "aab",
			expected: [][]string{
				{"a", "a", "b"},
				{"aa", "b"},
			},
		},
		{
			name: "Example 2",
			s:    "a",
			expected: [][]string{
				{"a"},
			},
		},
		{
			name: "Empty string",
			s:    "",
			expected: [][]string{
				{},
			},
		},
		{
			name: "Single character",
			s:    "b",
			expected: [][]string{
				{"b"},
			},
		},
		{
			name: "All same characters",
			s:    "aaa",
			expected: [][]string{
				{"a", "a", "a"},
				{"a", "aa"},
				{"aa", "a"},
				{"aaa"},
			},
		},
		{
			name: "Palindrome string",
			s:    "aba",
			expected: [][]string{
				{"a", "b", "a"},
				{"aba"},
			},
		},
		{
			name: "No palindrome partitions except single chars",
			s:    "abc",
			expected: [][]string{
				{"a", "b", "c"},
			},
		},
		{
			name: "Mixed palindrome possibilities",
			s:    "aabb",
			expected: [][]string{
				{"a", "a", "b", "b"},
				{"a", "a", "bb"},
				{"aa", "b", "b"},
				{"aa", "bb"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Partition(tt.s)
			
			// Sort both result and expected for comparison
			sortResult := func(slices [][]string) {
				for i := range slices {
					sort.Strings(slices[i])
				}
				sort.Slice(slices, func(i, j int) bool {
					if len(slices[i]) != len(slices[j]) {
						return len(slices[i]) < len(slices[j])
					}
					for k := 0; k < len(slices[i]); k++ {
						if slices[i][k] != slices[j][k] {
							return slices[i][k] < slices[j][k]
						}
					}
					return false
				})
			}
			
			sortResult(result)
			sortResult(tt.expected)
			
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Partition(%q) = %v, expected %v", tt.s, result, tt.expected)
			}
		})
	}
}

func BenchmarkPartition(b *testing.B) {
	testCases := []struct {
		name string
		s    string
	}{
		{"Short string", "aab"},
		{"Medium string", "aabbcc"},
		{"Longer string", "racecarlevel"},
		{"All same chars", "aaaaaaaa"},
		{"No palindromes", "abcdefgh"},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Partition(tc.s)
			}
		})
	}
}