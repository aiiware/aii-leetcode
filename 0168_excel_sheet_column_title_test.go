package leetcode

import (
	"fmt"
	"testing"
)

func TestConvertToTitle(t *testing.T) {
	tests := []struct {
		columnNumber int
		expect       string
	}{
		{
			columnNumber: 1,
			expect:       "A",
		},
		{
			columnNumber: 28,
			expect:       "AB",
		},
		{
			columnNumber: 701,
			expect:       "ZY",
		},
		{
			columnNumber: 26,
			expect:       "Z",
		},
		{
			columnNumber: 27,
			expect:       "AA",
		},
		{
			columnNumber: 52,
			expect:       "AZ",
		},
		{
			columnNumber: 53,
			expect:       "BA",
		},
		{
			columnNumber: 702,
			expect:       "ZZ",
		},
		{
			columnNumber: 703,
			expect:       "AAA",
		},
		{
			columnNumber: 18278,
			expect:       "ZZZ",
		},
		{
			columnNumber: 18279,
			expect:       "AAAA",
		},
		{
			columnNumber: 2147483647,
			expect:       "FXSHRXW",
		},
		{
			columnNumber: 2,
			expect:       "B",
		},
		{
			columnNumber: 3,
			expect:       "C",
		},
		{
			columnNumber: 4,
			expect:       "D",
		},
		{
			columnNumber: 5,
			expect:       "E",
		},
		{
			columnNumber: 6,
			expect:       "F",
		},
		{
			columnNumber: 7,
			expect:       "G",
		},
		{
			columnNumber: 8,
			expect:       "H",
		},
		{
			columnNumber: 9,
			expect:       "I",
		},
		{
			columnNumber: 10,
			expect:       "J",
		},
		{
			columnNumber: 11,
			expect:       "K",
		},
		{
			columnNumber: 12,
			expect:       "L",
		},
		{
			columnNumber: 13,
			expect:       "M",
		},
		{
			columnNumber: 14,
			expect:       "N",
		},
		{
			columnNumber: 15,
			expect:       "O",
		},
		{
			columnNumber: 16,
			expect:       "P",
		},
		{
			columnNumber: 17,
			expect:       "Q",
		},
		{
			columnNumber: 18,
			expect:       "R",
		},
		{
			columnNumber: 19,
			expect:       "S",
		},
		{
			columnNumber: 20,
			expect:       "T",
		},
		{
			columnNumber: 21,
			expect:       "U",
		},
		{
			columnNumber: 22,
			expect:       "V",
		},
		{
			columnNumber: 23,
			expect:       "W",
		},
		{
			columnNumber: 24,
			expect:       "X",
		},
		{
			columnNumber: 25,
			expect:       "Y",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test case %d: %d", i, tt.columnNumber), func(t *testing.T) {
			result := convertToTitle(tt.columnNumber)
			if result != tt.expect {
				t.Errorf("convertToTitle(%d) = %q, want %q", tt.columnNumber, result, tt.expect)
			}
		})
	}
}

func BenchmarkConvertToTitle(b *testing.B) {
	columnNumber := 2147483647
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		convertToTitle(columnNumber)
	}
}

func BenchmarkConvertToTitleSmall(b *testing.B) {
	columnNumber := 1
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		convertToTitle(columnNumber)
	}
}

func BenchmarkConvertToTitleMedium(b *testing.B) {
	columnNumber := 18278 // ZZZ
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		convertToTitle(columnNumber)
	}
}