package leetcode

import (
	"fmt"
	"testing"
)

func TestNumDecodings(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Example 1",
			s:        "12",
			expected: 2,
		},
		{
			name:     "Example 2",
			s:        "226",
			expected: 3,
		},
		{
			name:     "Example 3",
			s:        "06",
			expected: 0,
		},
		{
			name:     "Empty string",
			s:        "",
			expected: 0,
		},
		{
			name:     "Single digit valid",
			s:        "9",
			expected: 1,
		},
		{
			name:     "Single digit zero",
			s:        "0",
			expected: 0,
		},
		{
			name:     "All ones",
			s:        "111",
			expected: 3, // "1,1,1", "11,1", "1,11"
		},
		{
			name:     "With zero in middle",
			s:        "101",
			expected: 1, // Only "10,1" works
		},
		{
			name:     "Large number",
			s:        "12345",
			expected: 3, // "1,2,3,4,5", "12,3,4,5", "1,23,4,5"
		},
		{
			name:     "Ends with zero",
			s:        "10",
			expected: 1, // Only "10" works
		},
		{
			name:     "Double zero",
			s:        "100",
			expected: 0, // Cannot decode
		},
		{
			name:     "Complex case 1",
			s:        "1212",
			expected: 5,
		},
		{
			name:     "Complex case 2",
			s:        "12121",
			expected: 8,
		},
		{
			name:     "Maximum valid two-digit",
			s:        "26",
			expected: 2, // "2,6" and "26"
		},
		{
			name:     "Over maximum two-digit",
			s:        "27",
			expected: 1, // Only "2,7" works
		},
		{
			name:     "Leading zero",
			s:        "012",
			expected: 0,
		},
		{
			name:     "All twos",
			s:        "222",
			expected: 3, // "2,2,2", "22,2", "2,22"
		},
		{
			name:     "With zero after valid",
			s:        "110",
			expected: 1, // Only "1,10" works
		},
		{
			name:     "Large test case",
			s:        "111111111111111111111111111111111111111111111",
			expected: 1836311903, // Fibonacci-like sequence
		},
		{
			name:     "Zero in valid position",
			s:        "2101",
			expected: 1, // Only "2,10,1" works
		},
		{
			name:     "Multiple zeros",
			s:        "1001",
			expected: 0,
		},
		{
			name:     "Valid with zeros",
			s:        "1201234",
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NumDecodings(tt.s)
			if result != tt.expected {
				t.Errorf("NumDecodings(%q) = %d, expected %d",
					tt.s, result, tt.expected)
			}
		})
	}
}

func TestAllNumDecodingsImplementations(t *testing.T) {
	testCases := []struct {
		name string
		s    string
	}{
		{"Example 1", "12"},
		{"Example 2", "226"},
		{"Single digit", "9"},
		{"With zero", "101"},
		{"Complex", "1212"},
		{"Large", "111111"},
		{"Edge case", "0"},
		{"Edge case 2", "10"},
	}

	implementations := []struct {
		name string
		fn   func(string) int
	}{
		{"numDecodings", numDecodings},
		{"numDecodingsOptimized", numDecodingsOptimized},
		{"numDecodingsDFS", numDecodingsDFS},
		{"numDecodingsIterative", numDecodingsIterative},
		{"numDecodingsDP2", numDecodingsDP2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := NumDecodings(tc.s)

			for _, impl := range implementations {
				t.Run(impl.name, func(t *testing.T) {
					result := impl.fn(tc.s)
					if result != expected {
						t.Errorf("%s(%q) = %d, expected %d",
							impl.name, tc.s, result, expected)
					}
				})
			}
		})
	}
}

func TestNumDecodingsEdgeCases(t *testing.T) {
	t.Run("Empty string", func(t *testing.T) {
		if NumDecodings("") != 0 {
			t.Error("Empty string should return 0")
		}
	})

	t.Run("Single character '0'", func(t *testing.T) {
		if NumDecodings("0") != 0 {
			t.Error("'0' should return 0")
		}
	})

	t.Run("Single character '1'-'9'", func(t *testing.T) {
		for i := 1; i <= 9; i++ {
			s := fmt.Sprintf("%d", i)
			if NumDecodings(s) != 1 {
				t.Errorf("'%s' should return 1, got %d", s, NumDecodings(s))
			}
		}
	})

	t.Run("String starting with '0'", func(t *testing.T) {
		testCases := []string{"0", "01", "012", "001"}
		for _, s := range testCases {
			if NumDecodings(s) != 0 {
				t.Errorf("'%s' should return 0, got %d", s, NumDecodings(s))
			}
		}
	})

	t.Run("String containing '00'", func(t *testing.T) {
		testCases := []string{"100", "200", "300", "1001"}
		for _, s := range testCases {
			if NumDecodings(s) != 0 {
				t.Errorf("'%s' should return 0, got %d", s, NumDecodings(s))
			}
		}
	})

	t.Run("Valid two-digit numbers", func(t *testing.T) {
		// Test all valid two-digit combinations (10-26)
		for i := 10; i <= 26; i++ {
			s := fmt.Sprintf("%d", i)
			result := NumDecodings(s)
			if result != 2 {
				t.Errorf("'%s' should return 2, got %d", s, result)
			}
		}
	})

	t.Run("Invalid two-digit numbers", func(t *testing.T) {
		// Test some invalid two-digit combinations
		testCases := []string{"27", "30", "40", "50", "99"}
		for _, s := range testCases {
			result := NumDecodings(s)
			if result != 1 { // Can only decode as two single digits
				t.Errorf("'%s' should return 1, got %d", s, result)
			}
		}
	})

	t.Run("Large n with maximum answer", func(t *testing.T) {
		// Create a string of all '1's (Fibonacci sequence)
		s := ""
		for i := 0; i < 45; i++ {
			s += "1"
		}
		result := NumDecodings(s)
		// The 45th Fibonacci number is 1836311903
		if result != 1836311903 {
			t.Errorf("String of 45 '1's should return 1836311903, got %d", result)
		}
	})
}

func TestNumDecodingsProperties(t *testing.T) {
	// Property-based testing
	implementations := []struct {
		name string
		fn   func(string) int
	}{
		{"numDecodings", numDecodings},
		{"numDecodingsOptimized", numDecodingsOptimized},
		{"numDecodingsDFS", numDecodingsDFS},
		{"numDecodingsIterative", numDecodingsIterative},
		{"numDecodingsDP2", numDecodingsDP2},
	}

	testStrings := []string{
		"1",
		"12",
		"123",
		"1234",
		"101",
		"110",
		"111",
		"121",
		"226",
		"999",
		"1001",
		"1010",
	}

	for _, impl := range implementations {
		t.Run(impl.name+"_properties", func(t *testing.T) {
			for _, s := range testStrings {
				t.Run(fmt.Sprintf("s=%s", s), func(t *testing.T) {
					result := impl.fn(s)

					// Property 1: Result should be non-negative
					if result < 0 {
						t.Errorf("Result should be non-negative, got %d", result)
					}

					// Property 2: If string contains "00", result should be 0
					contains00 := false
					for i := 0; i < len(s)-1; i++ {
						if s[i] == '0' && s[i+1] == '0' {
							contains00 = true
							break
						}
					}
					if contains00 && result != 0 {
						t.Errorf("String %q contains '00' but result is %d, should be 0", s, result)
					}

					// Property 3: If string starts with '0', result should be 0
					if len(s) > 0 && s[0] == '0' && result != 0 {
						t.Errorf("String %q starts with '0' but result is %d, should be 0", s, result)
					}

					// Property 4: Result should be <= Fibonacci(len(s))
					// The number of decodings follows a Fibonacci-like sequence
					maxFib := fibonacci(len(s) + 1)
					if result > maxFib {
						t.Errorf("Result %d exceeds maximum possible %d for length %d",
							result, maxFib, len(s))
					}
				})
			}
		})
	}
}

func BenchmarkNumDecodings(b *testing.B) {
	// Test cases of different sizes and patterns
	testCases := []struct {
		name string
		s    string
	}{
		{"Small", "121"},
		{"Medium", "121212"},
		{"Large", "1212121212"},
		{"With zeros", "1010101010"},
		{"All ones", "1111111111"},
		{"Mixed", "1234567890"},
	}

	implementations := []struct {
		name string
		fn   func(string) int
	}{
		{"numDecodings", numDecodings},
		{"numDecodingsOptimized", numDecodingsOptimized},
		{"numDecodingsDFS", numDecodingsDFS},
		{"numDecodingsIterative", numDecodingsIterative},
		{"numDecodingsDP2", numDecodingsDP2},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for _, impl := range implementations {
				b.Run(impl.name, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						impl.fn(tc.s)
					}
				})
			}
		})
	}
}

func BenchmarkNumDecodingsWorstCase(b *testing.B) {
	// Worst case: all '1's (maximum branching)
	s := ""
	for i := 0; i < 100; i++ {
		s += "1"
	}

	b.ResetTimer()

	b.Run("numDecodings", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			numDecodings(s)
		}
	})

	b.Run("numDecodingsOptimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			numDecodingsOptimized(s)
		}
	})

	b.Run("numDecodingsDFS", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			numDecodingsDFS(s)
		}
	})
}

func BenchmarkNumDecodingsRecursive(b *testing.B) {
	// Test recursive implementation separately (it's exponential)
	s := "1111" // Small enough for recursive

	b.ResetTimer()
	b.Run("numDecodingsRecursive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			numDecodingsRecursive(s)
		}
	})
}

// Helper functions

// fibonacci computes the nth Fibonacci number
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// generateAllDecodings generates all possible decodings (for testing)
func generateAllDecodings(s string) []string {
	var result []string
	var backtrack func(int, string)
	
	backtrack = func(index int, current string) {
		if index == len(s) {
			result = append(result, current)
			return
		}
		
		// Cannot start with '0'
		if s[index] == '0' {
			return
		}
		
		// Decode single digit
		digit := s[index] - '0'
		backtrack(index+1, current+string('A'+digit-1))
		
		// Decode two digits if possible
		if index+1 < len(s) {
			twoDigit := (s[index]-'0')*10 + (s[index+1]-'0')
			if twoDigit >= 10 && twoDigit <= 26 {
				backtrack(index+2, current+string('A'+twoDigit-1))
			}
		}
	}
	
	backtrack(0, "")
	return result
}

func TestNumDecodingsMatchesActualDecodings(t *testing.T) {
	// Test that our count matches the actual number of decodings
	testCases := []string{
		"1",
		"12",
		"123",
		"121",
		"226",
		"101",
		"111",
	}

	for _, s := range testCases {
		t.Run(s, func(t *testing.T) {
			actualDecodings := generateAllDecodings(s)
			expectedCount := len(actualDecodings)
			result := NumDecodings(s)
			
			if result != expectedCount {
				t.Errorf("NumDecodings(%q) = %d, but there are %d actual decodings: %v",
					s, result, expectedCount, actualDecodings)
			}
		})
	}
}