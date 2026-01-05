package leetcode

import (
	"fmt"
	"testing"
)

func TestGrayCode(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		validate func([]int) bool
	}{
		{
			name: "n = 1",
			n:    1,
			validate: func(seq []int) bool {
				expected := []int{0, 1}
				return isValidGrayCode(seq, 1) && slicesEqual(seq, expected)
			},
		},
		{
			name: "n = 2",
			n:    2,
			validate: func(seq []int) bool {
				return isValidGrayCode(seq, 2)
			},
		},
		{
			name: "n = 3",
			n:    3,
			validate: func(seq []int) bool {
				return isValidGrayCode(seq, 3)
			},
		},
		{
			name: "n = 4",
			n:    4,
			validate: func(seq []int) bool {
				return isValidGrayCode(seq, 4)
			},
		},
		{
			name: "n = 5",
			n:    5,
			validate: func(seq []int) bool {
				return isValidGrayCode(seq, 5)
			},
		},
		{
			name: "n = 6",
			n:    6,
			validate: func(seq []int) bool {
				return isValidGrayCode(seq, 6)
			},
		},
		{
			name: "n = 7",
			n:    7,
			validate: func(seq []int) bool {
				return isValidGrayCode(seq, 7)
			},
		},
		{
			name: "n = 8",
			n:    8,
			validate: func(seq []int) bool {
				return isValidGrayCode(seq, 8)
			},
		},
		{
			name: "n = 0",
			n:    0,
			validate: func(seq []int) bool {
				// For n=0, the sequence should be [0]
				return len(seq) == 1 && seq[0] == 0
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GrayCode(tt.n)
			if !tt.validate(result) {
				t.Errorf("GrayCode(%d) = %v is not a valid Gray code sequence", tt.n, result)
			}
		})
	}
}

func TestAllGrayCodeImplementations(t *testing.T) {
	testCases := []struct {
		name string
		n    int
	}{
		{"n=1", 1},
		{"n=2", 2},
		{"n=3", 3},
		{"n=4", 4},
		{"n=5", 5},
	}

	implementations := []struct {
		name string
		fn   func(int) []int
	}{
		{"grayCode", grayCode},
		{"grayCodeFormula", grayCodeFormula},
		{"grayCodeRecursive", grayCodeRecursive},
		{"grayCodeIterative", grayCodeIterative},
		{"grayCodeIterative2", grayCodeIterative2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := GrayCode(tc.n)

			for _, impl := range implementations {
				t.Run(impl.name, func(t *testing.T) {
					result := impl.fn(tc.n)
					if !isValidGrayCode(result, tc.n) {
						t.Errorf("%s(%d) = %v is not a valid Gray code sequence",
							impl.name, tc.n, result)
					}

					// Check that it's a permutation of 0..2^n-1
					if !isPermutation(result, 1<<tc.n) {
						t.Errorf("%s(%d) = %v is not a permutation of 0..%d",
							impl.name, tc.n, result, (1<<tc.n)-1)
					}
				})
			}
		})
	}
}

func TestGrayCodeEdgeCases(t *testing.T) {
	t.Run("n = 0", func(t *testing.T) {
		result := GrayCode(0)
		expected := []int{0}
		if !slicesEqual(result, expected) {
			t.Errorf("GrayCode(0) = %v, expected %v", result, expected)
		}
	})

	t.Run("n = 1", func(t *testing.T) {
		result := GrayCode(1)
		if !isValidGrayCode(result, 1) {
			t.Errorf("GrayCode(1) = %v is not valid", result)
		}
		// Should be [0, 1] or [1, 0] (but our implementation returns [0, 1])
		if result[0] != 0 || result[1] != 1 {
			t.Errorf("GrayCode(1) = %v, expected [0, 1]", result)
		}
	})

	t.Run("n = 16 (maximum)", func(t *testing.T) {
		// This is the maximum allowed by constraints
		result := GrayCode(16)
		if !isValidGrayCode(result, 16) {
			t.Errorf("GrayCode(16) is not valid")
		}
		if len(result) != 1<<16 {
			t.Errorf("GrayCode(16) length = %d, expected %d", len(result), 1<<16)
		}
	})

	t.Run("Backtracking implementation", func(t *testing.T) {
		// Test the backtracking implementation separately since it's slower
		for n := 1; n <= 4; n++ {
			t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
				result := grayCodeBacktracking(n)
				if !isValidGrayCode(result, n) {
					t.Errorf("grayCodeBacktracking(%d) = %v is not valid", n, result)
				}
			})
		}
	})
}

func TestGrayCodeProperties(t *testing.T) {
	// Property-based testing
	implementations := []struct {
		name string
		fn   func(int) []int
	}{
		{"grayCode", grayCode},
		{"grayCodeFormula", grayCodeFormula},
		{"grayCodeRecursive", grayCodeRecursive},
		{"grayCodeIterative", grayCodeIterative},
		{"grayCodeIterative2", grayCodeIterative2},
	}

	for _, impl := range implementations {
		t.Run(impl.name+"_properties", func(t *testing.T) {
			for n := 1; n <= 8; n++ {
				t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
					result := impl.fn(n)

					// Property 1: Length should be 2^n
					expectedLength := 1 << n
					if len(result) != expectedLength {
						t.Errorf("Length = %d, expected %d", len(result), expectedLength)
					}

					// Property 2: Should contain all numbers from 0 to 2^n-1
					seen := make([]bool, expectedLength)
					for _, num := range result {
						if num < 0 || num >= expectedLength {
							t.Errorf("Number %d out of range [0, %d]", num, expectedLength-1)
						}
						if seen[num] {
							t.Errorf("Duplicate number %d", num)
						}
						seen[num] = true
					}
					for i := 0; i < expectedLength; i++ {
						if !seen[i] {
							t.Errorf("Missing number %d", i)
						}
					}

					// Property 3: Adjacent numbers should differ by exactly one bit
					for i := 0; i < len(result)-1; i++ {
						diff := result[i] ^ result[i+1]
						if countBits(diff) != 1 {
							t.Errorf("Adjacent numbers %d and %d differ by %d bits (diff=%d)",
								result[i], result[i+1], countBits(diff), diff)
						}
					}

					// Property 4: First and last should differ by exactly one bit
					firstLastDiff := result[0] ^ result[len(result)-1]
					if countBits(firstLastDiff) != 1 {
						t.Errorf("First and last numbers differ by %d bits",
							countBits(firstLastDiff))
					}

					// Property 5: Sequence should start with 0
					if result[0] != 0 {
						t.Errorf("Sequence should start with 0, got %d", result[0])
					}
				})
			}
		})
	}
}

func BenchmarkGrayCode(b *testing.B) {
	// Test cases of different sizes
	testCases := []struct {
		name string
		n    int
	}{
		{"n=4", 4},
		{"n=8", 8},
		{"n=12", 12},
		{"n=16", 16},
	}

	implementations := []struct {
		name string
		fn   func(int) []int
	}{
		{"grayCode", grayCode},
		{"grayCodeFormula", grayCodeFormula},
		{"grayCodeRecursive", grayCodeRecursive},
		{"grayCodeIterative", grayCodeIterative},
		{"grayCodeIterative2", grayCodeIterative2},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for _, impl := range implementations {
				b.Run(impl.name, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						impl.fn(tc.n)
					}
				})
			}
		})
	}
}

func BenchmarkGrayCodeBacktracking(b *testing.B) {
	// Backtracking is much slower, only test small n
	for n := 1; n <= 4; n++ {
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				grayCodeBacktracking(n)
			}
		})
	}
}

// Helper functions

// isValidGrayCode checks if a sequence is a valid n-bit Gray code
func isValidGrayCode(seq []int, n int) bool {
	size := 1 << n

	// Check length
	if len(seq) != size {
		return false
	}

	// Check range and uniqueness
	seen := make([]bool, size)
	for i, num := range seq {
		if num < 0 || num >= size {
			return false
		}
		if seen[num] {
			return false
		}
		seen[num] = true

		// Check adjacent difference (except for last element)
		if i > 0 {
			diff := seq[i-1] ^ num
			if countBits(diff) != 1 {
				return false
			}
		}
	}

	// Check first and last difference
	firstLastDiff := seq[0] ^ seq[size-1]
	return countBits(firstLastDiff) == 1
}

// isPermutation checks if a sequence is a permutation of 0..n-1
func isPermutation(seq []int, n int) bool {
	if len(seq) != n {
		return false
	}

	seen := make([]bool, n)
	for _, num := range seq {
		if num < 0 || num >= n {
			return false
		}
		if seen[num] {
			return false
		}
		seen[num] = true
	}

	// All numbers should be seen
	for i := 0; i < n; i++ {
		if !seen[i] {
			return false
		}
	}
	return true
}

// countBits counts the number of 1 bits in an integer
func countBits(x int) int {
	count := 0
	for x > 0 {
		count += x & 1
		x >>= 1
	}
	return count
}

// slicesEqual compares two slices for equality
func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}