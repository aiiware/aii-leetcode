package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{
			name:     "Example 1",
			input:    121,
			expected: true,
		},
		{
			name:     "Example 2",
			input:    -121,
			expected: false,
		},
		{
			name:     "Example 3",
			input:    10,
			expected: false,
		},
		{
			name:     "Zero",
			input:    0,
			expected: true,
		},
		{
			name:     "Single digit positive",
			input:    5,
			expected: true,
		},
		{
			name:     "Single digit negative",
			input:    -5,
			expected: false,
		},
		{
			name:     "Even digit palindrome",
			input:    1221,
			expected: true,
		},
		{
			name:     "Odd digit palindrome",
			input:    12321,
			expected: true,
		},
		{
			name:     "Not palindrome even digits",
			input:    1234,
			expected: false,
		},
		{
			name:     "Not palindrome odd digits",
			input:    12345,
			expected: false,
		},
		{
			name:     "Large palindrome",
			input:    123454321,
			expected: true,
		},
		{
			name:     "Large not palindrome",
			input:    123456789,
			expected: false,
		},
		{
			name:     "Ends with zero",
			input:    120,
			expected: false,
		},
		{
			name:     "Multiple zeros",
			input:    1001,
			expected: true,
		},
		{
			name:     "Max int32 palindrome",
			input:    2147447412,
			expected: true,
		},
		{
			name:     "Max int32 not palindrome",
			input:    2147483647,
			expected: false,
		},
		{
			name:     "Min int32",
			input:    -2147483648,
			expected: false,
		},
		{
			name:     "Palindrome with repeated digits",
			input:    9999,
			expected: true,
		},
		{
			name:     "Palindrome with different digits",
			input:    123321,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPalindrome(tt.input)
			assert.Equal(t, tt.expected, result, "Palindrome check should match")
		})
	}
}

func TestIsPalindromeString(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{
			name:     "Example 1",
			input:    121,
			expected: true,
		},
		{
			name:     "Example 2",
			input:    -121,
			expected: false,
		},
		{
			name:     "Large palindrome",
			input:    123454321,
			expected: true,
		},
		{
			name:     "Large not palindrome",
			input:    123456789,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isPalindromeString(tt.input)
			assert.Equal(t, tt.expected, result, "String version should match")
		})
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	benchmarks := []struct {
		name  string
		input int
	}{
		{
			name:  "Small palindrome",
			input: 121,
		},
		{
			name:  "Small not palindrome",
			input: 123,
		},
		{
			name:  "Medium palindrome",
			input: 12321,
		},
		{
			name:  "Medium not palindrome",
			input: 12345,
		},
		{
			name:  "Large palindrome",
			input: 123454321,
		},
		{
			name:  "Large not palindrome",
			input: 123456789,
		},
		{
			name:  "Very large palindrome",
			input: 12345678987654321,
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsPalindrome(bm.input)
			}
		})
	}
}

func BenchmarkIsPalindromeString(b *testing.B) {
	benchmarks := []struct {
		name  string
		input int
	}{
		{
			name:  "Small palindrome",
			input: 121,
		},
		{
			name:  "Small not palindrome",
			input: 123,
		},
		{
			name:  "Medium palindrome",
			input: 12321,
		},
		{
			name:  "Medium not palindrome",
			input: 12345,
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isPalindromeString(bm.input)
			}
		})
	}
}