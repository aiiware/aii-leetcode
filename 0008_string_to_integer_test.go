package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "Example 1",
			input:    "42",
			expected: 42,
		},
		{
			name:     "Example 2",
			input:    "   -42",
			expected: -42,
		},
		{
			name:     "Example 3",
			input:    "4193 with words",
			expected: 4193,
		},
		{
			name:     "Example 4",
			input:    "words and 987",
			expected: 0,
		},
		{
			name:     "Example 5",
			input:    "-91283472332",
			expected: -2147483648, // Underflow
		},
		{
			name:     "Empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "Only whitespace",
			input:    "     ",
			expected: 0,
		},
		{
			name:     "Positive sign",
			input:    "+42",
			expected: 42,
		},
		{
			name:     "Multiple signs",
			input:    "+-12",
			expected: 0,
		},
		{
			name:     "Leading zeros",
			input:    "  0000000000012345678",
			expected: 12345678,
		},
		{
			name:     "Max int32",
			input:    "2147483647",
			expected: 2147483647,
		},
		{
			name:     "Just above max int32",
			input:    "2147483648",
			expected: 2147483647, // Overflow
		},
		{
			name:     "Min int32",
			input:    "-2147483648",
			expected: -2147483648,
		},
		{
			name:     "Just below min int32",
			input:    "-2147483649",
			expected: -2147483648, // Underflow
		},
		{
			name:     "Large overflow",
			input:    "999999999999999999",
			expected: 2147483647,
		},
		{
			name:     "Large underflow",
			input:    "-999999999999999999",
			expected: -2147483648,
		},
		{
			name:     "Decimal number",
			input:    "3.14159",
			expected: 3,
		},
		{
			name:     "Mixed characters",
			input:    "  -0012a42",
			expected: -12,
		},
		{
			name:     "Zero",
			input:    "0",
			expected: 0,
		},
		{
			name:     "Negative zero",
			input:    "-0",
			expected: 0,
		},
		{
			name:     "Positive zero",
			input:    "+0",
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MyAtoi(tt.input)
			assert.Equal(t, tt.expected, result, "String to integer conversion should match")
		})
	}
}

func BenchmarkMyAtoi(b *testing.B) {
	benchmarks := []struct {
		name  string
		input string
	}{
		{
			name:  "Simple positive",
			input: "12345",
		},
		{
			name:  "With whitespace",
			input: "   12345",
		},
		{
			name:  "Negative",
			input: "-12345",
		},
		{
			name:  "Large number",
			input: "2147483647",
		},
		{
			name:  "With trailing text",
			input: "12345abc",
		},
		{
			name:  "Overflow case",
			input: "999999999999999999",
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				MyAtoi(bm.input)
			}
		})
	}
}