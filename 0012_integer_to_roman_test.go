package leetcode

import (
	"testing"
)

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		num      int
		expected string
	}{
		{1, "I"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{9, "IX"},
		{27, "XXVII"},
		{48, "XLVIII"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
		{3999, "MMMCMXCIX"},
		{2, "II"},
		{12, "XII"},
		{50, "L"},
		{100, "C"},
		{500, "D"},
		{1000, "M"},
		{444, "CDXLIV"},
		{888, "DCCCLXXXVIII"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := IntToRoman(tt.num)
			if result != tt.expected {
				t.Errorf("IntToRoman(%d) = %s, expected %s", tt.num, result, tt.expected)
			}
		})
	}
}

func BenchmarkIntToRoman(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntToRoman(1994)
	}
}

func BenchmarkIntToRomanLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntToRoman(3999)
	}
}
