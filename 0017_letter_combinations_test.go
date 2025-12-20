package leetcode

import (
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name     string
		digits   string
		expected []string
	}{
		{
			"Example 1",
			"23",
			[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			"Example 2",
			"",
			[]string{},
		},
		{
			"Single digit",
			"2",
			[]string{"a", "b", "c"},
		},
		{
			"Single digit 7",
			"7",
			[]string{"p", "q", "r", "s"},
		},
		{
			"Three digits",
			"234",
			[]string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi",
				"bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi",
				"cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LetterCombinations(tt.digits)
			if !stringSlicesEqual(result, tt.expected) {
				t.Errorf("LetterCombinations(%q) = %v, expected %v", tt.digits, result, tt.expected)
			}
		})
	}
}

func TestLetterCombinationsIterative(t *testing.T) {
	tests := []struct {
		name     string
		digits   string
		expected []string
	}{
		{
			"Example 1",
			"23",
			[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			"Single digit",
			"2",
			[]string{"a", "b", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LetterCombinationsIterative(tt.digits)
			if !stringSlicesEqual(result, tt.expected) {
				t.Errorf("LetterCombinationsIterative(%q) = %v, expected %v", tt.digits, result, tt.expected)
			}
		})
	}
}

// Helper function to compare string slices
func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func BenchmarkLetterCombinations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LetterCombinations("23")
	}
}

func BenchmarkLetterCombinationsIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LetterCombinationsIterative("23")
	}
}

func BenchmarkLetterCombinationsLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LetterCombinations("234567")
	}
}
