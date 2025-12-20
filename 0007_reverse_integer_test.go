package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "Example 1",
			input:    123,
			expected: 321,
		},
		{
			name:     "Example 2",
			input:    -123,
			expected: -321,
		},
		{
			name:     "Example 3",
			input:    120,
			expected: 21,
		},
		{
			name:     "Zero",
			input:    0,
			expected: 0,
		},
		{
			name:     "Single digit positive",
			input:    5,
			expected: 5,
		},
		{
			name:     "Single digit negative",
			input:    -5,
			expected: -5,
		},
		{
			name:     "Ends with zero",
			input:    100,
			expected: 1,
		},
		{
			name:     "Large positive number",
			input:    123456789,
			expected: 987654321,
		},
		{
			name:     "Large negative number",
			input:    -123456789,
			expected: -987654321,
		},
		{
			name:     "Overflow positive",
			input:    1534236469,
			expected: 0, // Reverses to 9646324351 which overflows 32-bit
		},
		{
			name:     "Overflow negative",
			input:    -1534236469,
			expected: 0, // Reverses to -9646324351 which overflows 32-bit
		},
		{
			name:     "Max int32",
			input:    2147483647,
			expected: 0, // Reverses to 7463847412 which overflows
		},
		{
			name:     "Min int32",
			input:    -2147483648,
			expected: 0, // Reverses to -8463847412 which overflows
		},
		{
			name:     "Just below overflow",
			input:    1463847412,
			expected: 2147483641,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reverse(tt.input)
			assert.Equal(t, tt.expected, result, "Reversed integer should match")
		})
	}
}

func TestReverseInt64(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{
			name:     "Example 1",
			input:    123,
			expected: 321,
		},
		{
			name:     "Example 2",
			input:    -123,
			expected: -321,
		},
		{
			name:     "Overflow positive",
			input:    1534236469,
			expected: 0,
		},
		{
			name:     "Overflow negative",
			input:    -1534236469,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reverseInt64(tt.input)
			assert.Equal(t, tt.expected, result, "reverseInt64 version should match")
		})
	}
}

func BenchmarkReverse(b *testing.B) {
	benchmarks := []struct {
		name  string
		input int
	}{
		{
			name:  "Small positive",
			input: 12345,
		},
		{
			name:  "Small negative",
			input: -12345,
		},
		{
			name:  "Medium number",
			input: 123456789,
		},
		{
			name:  "Large number",
			input: 1234567890,
		},
		{
			name:  "Near overflow",
			input: 1463847412,
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Reverse(bm.input)
			}
		})
	}
}

func BenchmarkReverseInt64(b *testing.B) {
	benchmarks := []struct {
		name  string
		input int
	}{
		{
			name:  "Small positive",
			input: 12345,
		},
		{
			name:  "Small negative",
			input: -12345,
		},
		{
			name:  "Medium number",
			input: 123456789,
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				reverseInt64(bm.input)
			}
		})
	}
}