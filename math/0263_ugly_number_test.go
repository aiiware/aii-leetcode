package math

import (
	"testing"
)

func TestIsUgly(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected bool
	}{
		{
			name:     "Example 1",
			n:        6,
			expected: true,
		},
		{
			name:     "Example 2",
			n:        1,
			expected: true,
		},
		{
			name:     "Example 3",
			n:        14,
			expected: false,
		},
		{
			name:     "Zero",
			n:        0,
			expected: false,
		},
		{
			name:     "Negative number",
			n:        -6,
			expected: false,
		},
		{
			name:     "Power of 2",
			n:        8,
			expected: true,
		},
		{
			name:     "Power of 3",
			n:        27,
			expected: true,
		},
		{
			name:     "Power of 5",
			n:        125,
			expected: true,
		},
		{
			name:     "Product of 2, 3, 5",
			n:        30,
			expected: true, // 2 * 3 * 5
		},
		{
			name:     "Contains prime factor 7",
			n:        21,
			expected: false, // 3 * 7
		},
		{
			name:     "Contains prime factor 11",
			n:        22,
			expected: false, // 2 * 11
		},
		{
			name:     "Large ugly number",
			n:        2 * 2 * 2 * 3 * 3 * 5 * 5, // 1800
			expected: true,
		},
		{
			name:     "Large non-ugly number",
			n:        2 * 3 * 7 * 11, // 462
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsUgly(tt.n)
			if result != tt.expected {
				t.Errorf("IsUgly(%d) = %v, expected %v", tt.n, result, tt.expected)
			}
		})
	}
}

func BenchmarkIsUgly(b *testing.B) {
	numbers := []int{6, 1, 14, 0, 30, 21, 1800, 462}
	
	for _, num := range numbers {
		b.Run("IsUgly", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsUgly(num)
			}
		})
	}
}