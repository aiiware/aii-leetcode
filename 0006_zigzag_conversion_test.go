package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		numRows  int
		expected string
	}{
		{
			name:     "Example 1",
			input:    "PAYPALISHIRING",
			numRows:  3,
			expected: "PAHNAPLSIIGYIR",
		},
		{
			name:     "Example 2",
			input:    "PAYPALISHIRING",
			numRows:  4,
			expected: "PINALSIGYAHRPI",
		},
		{
			name:     "Single row",
			input:    "ABCDEF",
			numRows:  1,
			expected: "ABCDEF",
		},
		{
			name:     "Two rows",
			input:    "ABCDEF",
			numRows:  2,
			expected: "ACEBDF",
		},
		{
			name:     "More rows than characters",
			input:    "ABC",
			numRows:  5,
			expected: "ABC",
		},
		{
			name:     "Empty string",
			input:    "",
			numRows:  3,
			expected: "",
		},
		{
			name:     "Single character",
			input:    "A",
			numRows:  3,
			expected: "A",
		},
		{
			name:     "Three rows simple",
			input:    "ABCDEF",
			numRows:  3,
			expected: "AEBDFC",
		},
		{
			name:     "Four rows simple",
			input:    "ABCDEFGHIJ",
			numRows:  4,
			expected: "AGBFHCEIDJ",
		},
		{
			name:     "Five rows - corrected",
			input:    "ABCDEFGHIJKLMNOP",
			numRows:  5,
			expected: "AIBHJPCGKODFLNEM", // 16 characters
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Convert(tt.input, tt.numRows)
			assert.Equal(t, tt.expected, result, "Zigzag conversion should match")
		})
	}
}

func TestConvertOptimized(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		numRows  int
		expected string
	}{
		{
			name:     "Example 1",
			input:    "PAYPALISHIRING",
			numRows:  3,
			expected: "PAHNAPLSIIGYIR",
		},
		{
			name:     "Example 2",
			input:    "PAYPALISHIRING",
			numRows:  4,
			expected: "PINALSIGYAHRPI",
		},
		{
			name:     "Single row",
			input:    "ABCDEF",
			numRows:  1,
			expected: "ABCDEF",
		},
		{
			name:     "Two rows",
			input:    "ABCDEF",
			numRows:  2,
			expected: "ACEBDF",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertOptimized(tt.input, tt.numRows)
			assert.Equal(t, tt.expected, result, "Optimized version should match")
		})
	}
}

func BenchmarkConvert(b *testing.B) {
	benchmarks := []struct {
		name    string
		input   string
		numRows int
	}{
		{
			name:    "Short string 3 rows",
			input:   "PAYPALISHIRING",
			numRows: 3,
		},
		{
			name:    "Short string 4 rows",
			input:   "PAYPALISHIRING",
			numRows: 4,
		},
		{
			name:    "Medium string",
			input:   makeZigzagString('A', 100),
			numRows: 5,
		},
		{
			name:    "Long string",
			input:   makeZigzagString('A', 1000),
			numRows: 10,
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Convert(bm.input, bm.numRows)
			}
		})
	}
}

func BenchmarkConvertOptimized(b *testing.B) {
	benchmarks := []struct {
		name    string
		input   string
		numRows int
	}{
		{
			name:    "Short string 3 rows",
			input:   "PAYPALISHIRING",
			numRows: 3,
		},
		{
			name:    "Short string 4 rows",
			input:   "PAYPALISHIRING",
			numRows: 4,
		},
		{
			name:    "Medium string",
			input:   makeZigzagString('A', 100),
			numRows: 5,
		},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ConvertOptimized(bm.input, bm.numRows)
			}
		})
	}
}

// Helper function to create a string of repeated characters
func makeZigzagString(ch rune, length int) string {
	result := make([]rune, length)
	for i := range result {
		result[i] = ch
	}
	return string(result)
}