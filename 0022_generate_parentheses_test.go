package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		n        int
		expected []string
	}{
		{
			1,
			[]string{"()"},
		},
		{
			2,
			[]string{"(())", "()()"},
		},
		{
			3,
			[]string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			0,
			[]string{""},
		},
	}

	for i, test := range tests {
		// Test backtracking approach
		result := GenerateParenthesis(test.n)
		sort.Strings(result)
		sort.Strings(test.expected)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Test %d (backtracking) failed: got %v, expected %v",
				i, result, test.expected)
		}

		// Test DP approach
		result = GenerateParenthesisDP(test.n)
		sort.Strings(result)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Test %d (DP) failed: got %v, expected %v",
				i, result, test.expected)
		}
	}
}

func BenchmarkGenerateParenthesis(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateParenthesis(8)
	}
}

func BenchmarkGenerateParenthesisDP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateParenthesisDP(8)
	}
}