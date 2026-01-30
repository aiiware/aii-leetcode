package arrays

import (
	"fmt"
	"testing"
)

func TestFractionToDecimal(t *testing.T) {
	tests := []struct {
		numerator   int
		denominator int
		expect      string
	}{
		{
			numerator:   1,
			denominator: 2,
			expect:      "0.5",
		},
		{
			numerator:   2,
			denominator: 1,
			expect:      "2",
		},
		{
			numerator:   4,
			denominator: 333,
			expect:      "0.(012)",
		},
		{
			numerator:   1,
			denominator: 3,
			expect:      "0.(3)",
		},
		{
			numerator:   1,
			denominator: 6,
			expect:      "0.1(6)",
		},
		{
			numerator:   1,
			denominator: 7,
			expect:      "0.(142857)",
		},
		{
			numerator:   -1,
			denominator: 2,
			expect:      "-0.5",
		},
		{
			numerator:   1,
			denominator: -2,
			expect:      "-0.5",
		},
		{
			numerator:   -1,
			denominator: -2,
			expect:      "0.5",
		},
		{
			numerator:   0,
			denominator: 3,
			expect:      "0",
		},
		{
			numerator:   22,
			denominator: 7,
			expect:      "3.(142857)",
		},
		{
			numerator:   1,
			denominator: 333,
			expect:      "0.(003)",
		},
		{
			numerator:   1,
			denominator: 8,
			expect:      "0.125",
		},
		{
			numerator:   1,
			denominator: 90,
			expect:      "0.0(1)",
		},
		{
			numerator:   1,
			denominator: 99,
			expect:      "0.(01)",
		},
		{
			numerator:   1,
			denominator: 999,
			expect:      "0.(001)",
		},
		{
			numerator:   -2147483648,
			denominator: -1,
			expect:      "2147483648",
		},
		{
			numerator:   -2147483648,
			denominator: 1,
			expect:      "-2147483648",
		},
		{
			numerator:   500,
			denominator: 10,
			expect:      "50",
		},
		{
			numerator:   7,
			denominator: 12,
			expect:      "0.58(3)",
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test case %d: %d/%d", i, tt.numerator, tt.denominator), func(t *testing.T) {
			result := fractionToDecimal(tt.numerator, tt.denominator)
			if result != tt.expect {
				t.Errorf("fractionToDecimal(%d, %d) = %q, want %q", tt.numerator, tt.denominator, result, tt.expect)
			}
		})
	}
}

func BenchmarkFractionToDecimal(b *testing.B) {
	numerator := 1
	denominator := 7
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fractionToDecimal(numerator, denominator)
	}
}

func BenchmarkFractionToDecimalLarge(b *testing.B) {
	numerator := 1000000007
	denominator := 999983
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fractionToDecimal(numerator, denominator)
	}
}

func BenchmarkFractionToDecimalRepeating(b *testing.B) {
	numerator := 1
	denominator := 333
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fractionToDecimal(numerator, denominator)
	}
}