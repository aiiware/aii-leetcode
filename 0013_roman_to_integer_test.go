package leetcode

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		roman    string
		expected int
	}{
		{"I", 1},
		{"II", 2},
		{"III", 3},
		{"IV", 4},
		{"V", 5},
		{"IX", 9},
		{"XII", 12},
		{"XXVII", 27},
		{"XLVIII", 48},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
		{"MMMCMXCIX", 3999},
		{"XLIV", 44},
		{"CDXLIV", 444},
		{"DCCCLXXXVIII", 888},
		{"M", 1000},
		{"CM", 900},
		{"XC", 90},
		{"IV", 4},
		{"CD", 400},
		{"XL", 40},
	}

	for _, tt := range tests {
		t.Run(tt.roman, func(t *testing.T) {
			result := RomanToInt(tt.roman)
			if result != tt.expected {
				t.Errorf("RomanToInt(%s) = %d, expected %d", tt.roman, result, tt.expected)
			}
		})
	}
}

func BenchmarkRomanToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RomanToInt("MCMXCIV")
	}
}

func BenchmarkRomanToIntLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RomanToInt("MMMCMXCIX")
	}
}
