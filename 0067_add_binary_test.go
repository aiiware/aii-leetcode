package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddBinary(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		expected string
	}{
		{
			name:     "Example 1: 11 + 1",
			a:        "11",
			b:        "1",
			expected: "100",
		},
		{
			name:     "Example 2: 1010 + 1011",
			a:        "1010",
			b:        "1011",
			expected: "10101",
		},
		{
			name:     "0 + 0",
			a:        "0",
			b:        "0",
			expected: "0",
		},
		{
			name:     "1 + 0",
			a:        "1",
			b:        "0",
			expected: "1",
		},
		{
			name:     "0 + 1",
			a:        "0",
			b:        "1",
			expected: "1",
		},
		{
			name:     "1 + 1",
			a:        "1",
			b:        "1",
			expected: "10",
		},
		{
			name:     "11 + 11",
			a:        "11",
			b:        "11",
			expected: "110",
		},
		{
			name:     "111 + 1",
			a:        "111",
			b:        "1",
			expected: "1000",
		},
		{
			name:     "101 + 110",
			a:        "101",
			b:        "110",
			expected: "1011",
		},
		{
			name:     "Large numbers: 111111 + 1",
			a:        "111111",
			b:        "1",
			expected: "1000000",
		},
		{
			name:     "Different lengths: 1 + 111",
			a:        "1",
			b:        "111",
			expected: "1000",
		},
		{
			name:     "Different lengths: 111 + 1",
			a:        "111",
			b:        "1",
			expected: "1000",
		},
		{
			name:     "All ones: 1111 + 1111",
			a:        "1111",
			b:        "1111",
			expected: "11110",
		},
		{
			name:     "With leading zeros in input",
			a:        "00101",
			b:        "00011",
			expected: "1000", // 5 + 3 = 8 = 1000
		},
		{
			name:     "Empty string a",
			a:        "",
			b:        "101",
			expected: "101",
		},
		{
			name:     "Empty string b",
			a:        "101",
			b:        "",
			expected: "101",
		},
		{
			name:     "Both empty",
			a:        "",
			b:        "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AddBinary(tt.a, tt.b)
			assert.Equal(t, tt.expected, result,
				"AddBinary(%q, %q) = %q, expected %q",
				tt.a, tt.b, result, tt.expected)
		})
	}
}

func TestAddBinary_EdgeCases(t *testing.T) {
	t.Run("Very large binary numbers", func(t *testing.T) {
		// Create two 1000-bit numbers of all 1s
		a := ""
		b := ""
		for i := 0; i < 1000; i++ {
			a += "1"
			b += "1"
		}
		// Result should be 1000 ones followed by a zero
		// (111...111 + 111...111 = 111...1110)
		expected := ""
		for i := 0; i < 1000; i++ {
			expected += "1"
		}
		expected += "0"

		result := AddBinary(a, b)
		assert.Equal(t, expected, result)
	})

	t.Run("Invalid binary digit returns empty string", func(t *testing.T) {
		assert.Empty(t, AddBinary("1012", "110"))
		assert.Empty(t, AddBinary("101", "11a"))
		assert.Empty(t, AddBinary("abc", "def"))
	})

	t.Run("Single digit with carry", func(t *testing.T) {
		assert.Equal(t, "10", AddBinary("1", "1"))
		assert.Equal(t, "11", AddBinary("1", "10"))
		assert.Equal(t, "100", AddBinary("11", "1"))
	})

	t.Run("Multiple carries", func(t *testing.T) {
		assert.Equal(t, "1000", AddBinary("111", "1"))
		assert.Equal(t, "10100", AddBinary("1111", "101"))
		assert.Equal(t, "11110", AddBinary("1111", "1111"))
	})

	t.Run("Binary addition is commutative", func(t *testing.T) {
		testPairs := [][2]string{
			{"101", "110"},
			{"1111", "1"},
			{"1001", "1010"},
			{"1111111111", "1"},
		}

		for _, pair := range testPairs {
			a, b := pair[0], pair[1]
			result1 := AddBinary(a, b)
			result2 := AddBinary(b, a)
			assert.Equal(t, result1, result2,
				"AddBinary(%q, %q) = %q should equal AddBinary(%q, %q) = %q",
				a, b, result1, b, a, result2)
		}
	})
}

func TestAddBinary_PropertyBased(t *testing.T) {
	// Test that AddBinary matches integer addition
	testCases := []struct {
		a int
		b int
	}{
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 1},
		{5, 3},    // 101 + 011 = 1000 (8)
		{10, 5},   // 1010 + 0101 = 1111 (15)
		{15, 15},  // 1111 + 1111 = 11110 (30)
		{255, 1},  // 11111111 + 1 = 100000000 (256)
		{123, 456}, // Large numbers
	}

	for _, tc := range testCases {
		t.Run("Property test", func(t *testing.T) {
			aBin := intToBinary(tc.a)
			bBin := intToBinary(tc.b)
			
			result := AddBinary(aBin, bBin)
			resultInt := binaryToInt(result)
			
			assert.Equal(t, tc.a+tc.b, resultInt,
				"AddBinary(%q (%d), %q (%d)) = %q (%d), expected %d",
				aBin, tc.a, bBin, tc.b, result, resultInt, tc.a+tc.b)
		})
	}

	// Test that adding zero returns the same number
	t.Run("Add zero identity", func(t *testing.T) {
		testBinaries := []string{"0", "1", "10", "11", "101", "1101", "111111"}
		for _, bin := range testBinaries {
			assert.Equal(t, bin, AddBinary(bin, "0"))
			assert.Equal(t, bin, AddBinary("0", bin))
		}
	})

	// Test that adding to all ones produces power of two
	t.Run("Add to all ones", func(t *testing.T) {
		for n := 1; n <= 10; n++ {
			allOnes := ""
			for i := 0; i < n; i++ {
				allOnes += "1"
			}
			
			result := AddBinary(allOnes, "1")
			
			// Result should be 1 followed by n zeros
			expected := "1"
			for i := 0; i < n; i++ {
				expected += "0"
			}
			
			assert.Equal(t, expected, result,
				"AddBinary(%q, \"1\") = %q, expected %q",
				allOnes, result, expected)
		}
	})
}

// Helper functions for property-based testing
func intToBinary(n int) string {
	if n == 0 {
		return "0"
	}
	
	var binary string
	for n > 0 {
		if n%2 == 1 {
			binary = "1" + binary
		} else {
			binary = "0" + binary
		}
		n /= 2
	}
	return binary
}

func binaryToInt(binary string) int {
	result := 0
	for _, ch := range binary {
		result = result * 2
		if ch == '1' {
			result += 1
		}
	}
	return result
}

func BenchmarkAddBinary(b *testing.B) {
	testCases := []struct {
		name string
		a    string
		b    string
	}{
		{"Small", "101", "110"},
		{"Medium", "111111", "1"},
		{"Large", func() string {
			s := ""
			for i := 0; i < 1000; i++ {
				s += "1"
			}
			return s
		}(), "1"},
		{"Very large", func() string {
			s := ""
			for i := 0; i < 10000; i++ {
				s += "1"
			}
			return s
		}(), func() string {
			s := ""
			for i := 0; i < 10000; i++ {
				s += "1"
			}
			return s
		}()},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				AddBinary(tc.a, tc.b)
			}
		})
	}
}