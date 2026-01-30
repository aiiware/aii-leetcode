package arrays

import (
	"fmt"
	"testing"
)

func TestReverseBits(t *testing.T) {
	tests := []struct {
		name     string
		input    uint32
		expected uint32
	}{
		// Test cases from LeetCode examples
		{
			name:     "Example 1: 43261596 -> 964176192",
			input:    43261596, // 00000010100101000001111010011100
			expected: 964176192, // 00111001011110000010100101000000
		},
		{
			name:     "Example 2: 4294967293 -> 3221225471",
			input:    4294967293, // 11111111111111111111111111111101
			expected: 3221225471, // 10111111111111111111111111111111
		},
		// Edge cases
		{
			name:     "Zero",
			input:    0, // 00000000000000000000000000000000
			expected: 0, // 00000000000000000000000000000000
		},
		{
			name:     "All ones",
			input:    4294967295, // 11111111111111111111111111111111
			expected: 4294967295, // 11111111111111111111111111111111
		},
		{
			name:     "Power of two (1)",
			input:    1, // 00000000000000000000000000000001
			expected: 2147483648, // 10000000000000000000000000000000
		},
		{
			name:     "Power of two (2^31)",
			input:    2147483648, // 10000000000000000000000000000000
			expected: 1, // 00000000000000000000000000000001
		},
		{
			name:     "Alternating pattern 1010...",
			input:    2863311530, // 10101010101010101010101010101010
			expected: 1431655765, // 01010101010101010101010101010101
		},
		{
			name:     "Alternating pattern 0101...",
			input:    1431655765, // 01010101010101010101010101010101
			expected: 2863311530, // 10101010101010101010101010101010
		},
		// Additional test cases
		{
			name:     "Single bit at position 16",
			input:    65536, // 00000000000000010000000000000000
			expected: 32768, // 00000000000000001000000000000000
		},
		{
			name:     "Single bit at position 8",
			input:    256, // 00000000000000000000000100000000
			expected: 8388608, // 00000000100000000000000000000000
		},
		{
			name:     "Single bit at position 0",
			input:    1, // 00000000000000000000000000000001
			expected: 2147483648, // 10000000000000000000000000000000
		},
		{
			name:     "Single bit at position 31",
			input:    2147483648, // 10000000000000000000000000000000
			expected: 1, // 00000000000000000000000000000001
		},
		{
			name:     "Max 32-bit unsigned integer",
			input:    4294967295, // 11111111111111111111111111111111
			expected: 4294967295, // 11111111111111111111111111111111
		},
		{
			name:     "Pattern: 11001100...",
			input:    3435973836, // 11001100110011001100110011001100
			expected: 858993459,  // 00110011001100110011001100110011
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reverseBits(tt.input)
			if result != tt.expected {
				t.Errorf("reverseBits(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestReverseBits1(t *testing.T) {
	tests := []struct {
		input    uint32
		expected uint32
	}{
		{43261596, 964176192},
		{4294967293, 3221225471},
		{0, 0},
		{4294967295, 4294967295},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			result := reverseBits1(tt.input)
			if result != tt.expected {
				t.Errorf("reverseBits1(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestReverseBits2(t *testing.T) {
	tests := []struct {
		input    uint32
		expected uint32
	}{
		{43261596, 964176192},
		{4294967293, 3221225471},
		{0, 0},
		{4294967295, 4294967295},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			result := reverseBits2(tt.input)
			if result != tt.expected {
				t.Errorf("reverseBits2(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestReverseBits3(t *testing.T) {
	tests := []struct {
		input    uint32
		expected uint32
	}{
		{43261596, 964176192},
		{4294967293, 3221225471},
		{0, 0},
		{4294967295, 4294967295},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
			result := reverseBits3(tt.input)
			if result != tt.expected {
				t.Errorf("reverseBits3(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestReverseBitsAllImplementationsMatch(t *testing.T) {
	testCases := []uint32{
		0,
		1,
		2,
		255,
		256,
		65535,
		65536,
		16777215,
		16777216,
		43261596,
		4294967293,
		4294967295,
		1431655765,
		2863311530,
		3435973836,
	}

	for _, input := range testCases {
		t.Run(fmt.Sprintf("Input %v", input), func(t *testing.T) {
			result1 := reverseBits1(input)
			result2 := reverseBits2(input)
			result3 := reverseBits3(input)

			if result1 != result2 || result2 != result3 {
				t.Errorf("Implementations don't match for input %v: reverseBits1=%v, reverseBits2=%v, reverseBits3=%v",
					input, result1, result2, result3)
			}
		})
	}
}

// Benchmark tests for performance comparison
func BenchmarkReverseBits1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseBits1(43261596)
	}
}

func BenchmarkReverseBits2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseBits2(43261596)
	}
}

func BenchmarkReverseBits3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseBits3(43261596)
	}
}

func BenchmarkReverseBits(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseBits(43261596)
	}
}

// Helper function to print binary representation (for debugging)
func printBinary(n uint32) string {
	return fmt.Sprintf("%032b", n)
}

func TestPrintBinary(t *testing.T) {
	// This test is just to verify our binary printing works correctly
	cases := []struct {
		input    uint32
		expected string
	}{
		{0, "00000000000000000000000000000000"},
		{1, "00000000000000000000000000000001"},
		{4294967295, "11111111111111111111111111111111"},
	}

	for _, tc := range cases {
		result := printBinary(tc.input)
		if result != tc.expected {
			t.Errorf("printBinary(%v) = %v, want %v", tc.input, result, tc.expected)
		}
	}
}