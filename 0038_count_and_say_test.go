package leetcode

import (
	"testing"
)

func TestCountAndSay(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected string
	}{
		{
			name:     "n = 1",
			n:        1,
			expected: "1",
		},
		{
			name:     "n = 2",
			n:        2,
			expected: "11",
		},
		{
			name:     "n = 3",
			n:        3,
			expected: "21",
		},
		{
			name:     "n = 4",
			n:        4,
			expected: "1211",
		},
		{
			name:     "n = 5",
			n:        5,
			expected: "111221",
		},
		{
			name:     "n = 6",
			n:        6,
			expected: "312211",
		},
		{
			name:     "n = 7",
			n:        7,
			expected: "13112221",
		},
		{
			name:     "n = 8",
			n:        8,
			expected: "1113213211",
		},
		{
			name:     "n = 9",
			n:        9,
			expected: "31131211131221",
		},
		{
			name:     "n = 10",
			n:        10,
			expected: "13211311123113112211",
		},
		{
			name:     "n = 0",
			n:        0,
			expected: "",
		},
		{
			name:     "n = -1",
			n:        -1,
			expected: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Test iterative approach
			result := CountAndSay(test.n)
			if result != test.expected {
				t.Errorf("CountAndSay(%d) = %q, expected %q", test.n, result, test.expected)
			}

			// Test recursive approach
			if test.n > 0 { // Recursive doesn't handle n <= 0
				result = CountAndSayRecursive(test.n)
				if result != test.expected {
					t.Errorf("CountAndSayRecursive(%d) = %q, expected %q", test.n, result, test.expected)
				}
			}

			// Test optimized iterative approach
			result = CountAndSayIterativeOptimized(test.n)
			if result != test.expected {
				t.Errorf("CountAndSayIterativeOptimized(%d) = %q, expected %q", test.n, result, test.expected)
			}
		})
	}
}

func BenchmarkCountAndSay(b *testing.B) {
	n := 15
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountAndSay(n)
	}
}

func BenchmarkCountAndSayRecursive(b *testing.B) {
	n := 15
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountAndSayRecursive(n)
	}
}

func BenchmarkCountAndSayIterativeOptimized(b *testing.B) {
	n := 15
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CountAndSayIterativeOptimized(n)
	}
}