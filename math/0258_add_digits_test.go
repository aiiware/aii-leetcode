package math

import (
	"testing"
)

func TestAddDigits(t *testing.T) {
	tests := []struct {
		name     string
		num      int
		expected int
	}{
		{
			name:     "Example 1",
			num:      38,
			expected: 2,
		},
		{
			name:     "Example 2",
			num:      0,
			expected: 0,
		},
		{
			name:     "Single digit",
			num:      5,
			expected: 5,
		},
		{
			name:     "Multiple of 9",
			num:      18,
			expected: 9,
		},
		{
			name:     "Large number",
			num:      12345,
			expected: 6, // 1+2+3+4+5=15, 1+5=6
		},
		{
			name:     "Another multiple of 9",
			num:      999,
			expected: 9, // 9+9+9=27, 2+7=9
		},
		{
			name:     "Maximum constraint",
			num:      2147483647, // 2^31 - 1
			expected: 1,          // Sum of digits = 46, 4+6=10, 1+0=1
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AddDigits(tt.num)
			if result != tt.expected {
				t.Errorf("AddDigits(%d) = %d, expected %d", tt.num, result, tt.expected)
			}
			
			// Also test the loop version
			loopResult := AddDigitsLoop(tt.num)
			if loopResult != tt.expected {
				t.Errorf("AddDigitsLoop(%d) = %d, expected %d", tt.num, loopResult, tt.expected)
			}
		})
	}
}

func BenchmarkAddDigits(b *testing.B) {
	numbers := []int{38, 0, 12345, 999, 2147483647}
	
	for _, num := range numbers {
		b.Run("Formula", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				AddDigits(num)
			}
		})
		
		b.Run("Loop", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				AddDigitsLoop(num)
			}
		})
	}
}