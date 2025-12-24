package leetcode

import (
	"testing"
)

func TestStrStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		expected int
	}{
		{"sadbutsad", "sad", 0},
		{"leetcode", "leeto", -1},
		{"hello", "ll", 2},
		{"", "", 0},
		{"", "a", -1},
		{"a", "", 0},
		{"aaa", "aaaa", -1},
		{"mississippi", "issip", 4},
		{"mississippi", "pi", 9},
		{"abc", "c", 2},
		{"abc", "abc", 0},
		{"abc", "abcd", -1},
	}

	for i, test := range tests {
		// Test brute force approach
		result := StrStr(test.haystack, test.needle)

		if result != test.expected {
			t.Errorf("Test %d (brute force) failed: got %d, expected %d",
				i, result, test.expected)
		}

		// Test KMP approach
		result = StrStrKMP(test.haystack, test.needle)

		if result != test.expected {
			t.Errorf("Test %d (KMP) failed: got %d, expected %d",
				i, result, test.expected)
		}
	}
}

func BenchmarkStrStr(b *testing.B) {
	haystack := "a" + repeat("b", 1000) + "a" + repeat("b", 1000) + "a"
	needle := "a" + repeat("b", 1000) + "a"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StrStr(haystack, needle)
	}
}

func BenchmarkStrStrKMP(b *testing.B) {
	haystack := "a" + repeat("b", 1000) + "a" + repeat("b", 1000) + "a"
	needle := "a" + repeat("b", 1000) + "a"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StrStrKMP(haystack, needle)
	}
}

// Helper function to repeat a string
func repeat(s string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}