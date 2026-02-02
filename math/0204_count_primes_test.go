package math

import (
	"testing"
)

func TestCountPrimes(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{0, 0},      // No primes less than 0
		{1, 0},      // No primes less than 1
		{2, 0},      // No primes less than 2 (2 is not less than 2)
		{3, 1},      // Primes less than 3: 2
		{10, 4},     // Primes less than 10: 2, 3, 5, 7
		{20, 8},     // Primes less than 20: 2, 3, 5, 7, 11, 13, 17, 19
		{30, 10},    // Primes less than 30: 2, 3, 5, 7, 11, 13, 17, 19, 23, 29
		{100, 25},   // Classic test case
		{1000, 168}, // Known value: π(1000) = 168
	}

	for _, test := range tests {
		result := CountPrimes(test.n)
		if result != test.expected {
			t.Errorf("CountPrimes(%d) = %d, expected %d", test.n, result, test.expected)
		}

		// Test optimized version
		result2 := CountPrimesOptimized(test.n)
		if result2 != test.expected {
			t.Errorf("CountPrimesOptimized(%d) = %d, expected %d", test.n, result2, test.expected)
		}

		// Test simple version
		result3 := CountPrimesSimple(test.n)
		if result3 != test.expected {
			t.Errorf("CountPrimesSimple(%d) = %d, expected %d", test.n, result3, test.expected)
		}
	}
}

func TestCountPrimes_LargeNumbers(t *testing.T) {
	// Test with larger numbers (but keep it reasonable for test speed)
	tests := []struct {
		n        int
		expected int
	}{
		{5000, 669},    // π(5000) = 669
		{10000, 1229},  // π(10000) = 1229
		{50000, 5133},  // π(50000) = 5133
	}

	for _, test := range tests {
		result := CountPrimes(test.n)
		if result != test.expected {
			t.Errorf("CountPrimes(%d) = %d, expected %d", test.n, result, test.expected)
		}
	}
}

func TestCountPrimes_EdgeCases(t *testing.T) {
	// Edge cases and negative numbers
	tests := []struct {
		n        int
		expected int
	}{
		{-10, 0},   // Negative input
		{-1, 0},    // Negative input
		{0, 0},     // Zero
		{1, 0},     // One
		{2, 0},     // Two (boundary case)
		{3, 1},     // Three
	}

	for _, test := range tests {
		result := CountPrimes(test.n)
		if result != test.expected {
			t.Errorf("CountPrimes(%d) = %d, expected %d", test.n, result, test.expected)
		}
	}
}

func BenchmarkCountPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountPrimes(1000000)
	}
}

func BenchmarkCountPrimesOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountPrimesOptimized(1000000)
	}
}

func BenchmarkCountPrimesSimple(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountPrimesSimple(1000000)
	}
}

// Helper function to verify prime counting for small ranges
func TestPrimeVerification(t *testing.T) {
	// Manually verify first few primes
	primesLessThan20 := []int{2, 3, 5, 7, 11, 13, 17, 19}
	
	for i := 0; i <= 20; i++ {
		count := CountPrimes(i)
		// Count how many primes in primesLessThan20 are less than i
		expected := 0
		for _, p := range primesLessThan20 {
			if p < i {
				expected++
			}
		}
		
		if count != expected {
			t.Errorf("CountPrimes(%d) = %d, but expected %d based on manual count", i, count, expected)
		}
	}
}