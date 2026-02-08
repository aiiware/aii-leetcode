package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegerBreak(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "n = 2",
			n:        2,
			expected: 1,
		},
		{
			name:     "n = 3",
			n:        3,
			expected: 2,
		},
		{
			name:     "n = 4",
			n:        4,
			expected: 4,
		},
		{
			name:     "n = 5",
			n:        5,
			expected: 6,
		},
		{
			name:     "n = 6",
			n:        6,
			expected: 9,
		},
		{
			name:     "n = 7",
			n:        7,
			expected: 12,
		},
		{
			name:     "n = 8",
			n:        8,
			expected: 18,
		},
		{
			name:     "n = 9",
			n:        9,
			expected: 27,
		},
		{
			name:     "n = 10",
			n:        10,
			expected: 36,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IntegerBreak(tt.n)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func BenchmarkIntegerBreak(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IntegerBreak(10)
	}
}