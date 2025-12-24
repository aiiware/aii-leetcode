package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	tests := []struct {
		s        string
		words    []string
		expected []int
	}{
		{
			"barfoothefoobarman",
			[]string{"foo", "bar"},
			[]int{0, 9},
		},
		{
			"wordgoodgoodgoodbestword",
			[]string{"word", "good", "best", "word"},
			[]int{},
		},
		{
			"barfoofoobarthefoobarman",
			[]string{"bar", "foo", "the"},
			[]int{6, 9, 12},
		},
		{
			"wordgoodgoodgoodbestword",
			[]string{"word", "good", "best", "good"},
			[]int{8},
		},
		{
			"",
			[]string{"foo", "bar"},
			[]int{},
		},
		{
			"barfoothefoobarman",
			[]string{},
			[]int{},
		},
		{
			"a",
			[]string{"a"},
			[]int{0},
		},
		{
			"aaaaaaaa",
			[]string{"aa", "aa", "aa"},
			[]int{0, 1, 2},
		},
	}

	for i, test := range tests {
		// Test basic approach
		result := FindSubstring(test.s, test.words)
		sort.Ints(result)
		sort.Ints(test.expected)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Test %d (basic) failed: got %v, expected %v",
				i, result, test.expected)
		}

		// Test optimized approach
		result = FindSubstringOptimized(test.s, test.words)
		sort.Ints(result)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Test %d (optimized) failed: got %v, expected %v",
				i, result, test.expected)
		}
	}
}

func BenchmarkFindSubstring(b *testing.B) {
	s := "barfoothefoobarmanbarfoothefoobarmanbarfoothefoobarman"
	words := []string{"foo", "bar"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindSubstring(s, words)
	}
}

func BenchmarkFindSubstringOptimized(b *testing.B) {
	s := "barfoothefoobarmanbarfoothefoobarmanbarfoothefoobarman"
	words := []string{"foo", "bar"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindSubstringOptimized(s, words)
	}
}