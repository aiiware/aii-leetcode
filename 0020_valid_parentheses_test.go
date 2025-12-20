package leetcode

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected bool
	}{
		{"Example 1", "()", true},
		{"Example 2", "()[]{}", true},
		{"Example 3", "(]", false},
		{"Example 4", "([)]", false},
		{"Example 5", "{[]}", true},
		{"Empty string", "", true},
		{"Single opening", "(", false},
		{"Single closing", ")", false},
		{"Multiple valid", "((()))", true},
		{"Nested valid", "{[()]}", true},
		{"Multiple types", "()[]{}", true},
		{"Wrong order", "({[}])", false},
		{"Extra closing", "())", false},
		{"Extra opening", "(()", false},
		{"Mixed wrong", "([{}]", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValid(tt.s)
			if result != tt.expected {
				t.Errorf("IsValid(%q) = %v, expected %v", tt.s, result, tt.expected)
			}
		})
	}
}

func BenchmarkIsValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValid("()")
	}
}

func BenchmarkIsValidComplex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValid("{[()]}")
	}
}

func BenchmarkIsValidLarge(b *testing.B) {
	// Create a large valid parentheses string
	s := ""
	for i := 0; i < 100; i++ {
		s += "()"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsValid(s)
	}
}
