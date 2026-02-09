package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstBadVersion(t *testing.T) {
	tests := []struct {
		name           string
		n              int
		firstBad       int
		expected       int
		expectedCalls  int
		maxAllowedCalls int
	}{
		{
			name:           "Example 1: First version is bad",
			n:              5,
			firstBad:       1,
			expected:       1,
			expectedCalls:  3, // log2(5) ≈ 2.32 → 3 calls
			maxAllowedCalls: 3,
		},
		{
			name:           "Example 2: Last version is bad",
			n:              5,
			firstBad:       5,
			expected:       5,
			expectedCalls:  3,
			maxAllowedCalls: 3,
		},
		{
			name:           "Example 3: Middle version is bad",
			n:              5,
			firstBad:       3,
			expected:       3,
			expectedCalls:  3,
			maxAllowedCalls: 3,
		},
		{
			name:           "Single version, bad",
			n:              1,
			firstBad:       1,
			expected:       1,
			expectedCalls:  1,
			maxAllowedCalls: 1,
		},
		{
			name:           "Large n, first bad at beginning",
			n:              1000,
			firstBad:       1,
			expected:       1,
			expectedCalls:  10, // log2(1000) ≈ 9.97 → 10 calls
			maxAllowedCalls: 10,
		},
		{
			name:           "Large n, first bad at end",
			n:              1000,
			firstBad:       1000,
			expected:       1000,
			expectedCalls:  10,
			maxAllowedCalls: 10,
		},
		{
			name:           "Large n, first bad in middle",
			n:              1000,
			firstBad:       500,
			expected:       500,
			expectedCalls:  10,
			maxAllowedCalls: 10,
		},
		{
			name:           "Even number of versions",
			n:              10,
			firstBad:       4,
			expected:       4,
			expectedCalls:  4, // log2(10) ≈ 3.32 → 4 calls
			maxAllowedCalls: 4,
		},
		{
			name:           "Odd number of versions",
			n:              11,
			firstBad:       6,
			expected:       6,
			expectedCalls:  4, // log2(11) ≈ 3.46 → 4 calls
			maxAllowedCalls: 4,
		},
		{
			name:           "First bad version is 2",
			n:              10,
			firstBad:       2,
			expected:       2,
			expectedCalls:  4,
			maxAllowedCalls: 4,
		},
		{
			name:           "First bad version is n-1",
			n:              10,
			firstBad:       9,
			expected:       9,
			expectedCalls:  4,
			maxAllowedCalls: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			callCount := 0
			isBadVersion := func(version int) bool {
				callCount++
				return version >= tt.firstBad
			}

			result := FirstBadVersion(tt.n, isBadVersion)

			assert.Equal(t, tt.expected, result,
				"FirstBadVersion(%d, isBadVersion) = %d, expected %d",
				tt.n, result, tt.expected)

			// Verify we didn't exceed the expected number of API calls
			assert.LessOrEqual(t, callCount, tt.maxAllowedCalls,
				"Made %d API calls, expected at most %d", callCount, tt.maxAllowedCalls)

			// For binary search, call count should be O(log n)
			// We can also check it's close to expectedCalls (allow ±1 for edge cases)
			assert.GreaterOrEqual(t, callCount, tt.expectedCalls-1,
				"Made %d API calls, expected at least %d", callCount, tt.expectedCalls-1)
		})
	}
}

func TestFirstBadVersion_EdgeCases(t *testing.T) {
	t.Run("All versions are good (shouldn't happen per problem)", func(t *testing.T) {
		callCount := 0
		isBadVersion := func(version int) bool {
			callCount++
			return false // No bad versions
		}

		// According to problem, there's always at least one bad version
		// But we test the behavior anyway
		result := FirstBadVersion(10, isBadVersion)

		// When no version is bad, binary search will return the last version checked
		// which is n (10) because left will keep moving right
		assert.Equal(t, 10, result)
	})

	t.Run("All versions are bad", func(t *testing.T) {
		callCount := 0
		isBadVersion := func(version int) bool {
			callCount++
			return true // All versions are bad
		}

		result := FirstBadVersion(10, isBadVersion)

		// First bad version should be 1
		assert.Equal(t, 1, result)
		assert.LessOrEqual(t, callCount, 4) // log2(10) ≈ 3.32 → 4 calls
	})

	t.Run("n = 0 (edge case)", func(t *testing.T) {
		callCount := 0
		isBadVersion := func(version int) bool {
			callCount++
			return true
		}

		result := FirstBadVersion(0, isBadVersion)

		// Should return 0 for invalid input
		assert.Equal(t, 0, result)
		assert.Equal(t, 0, callCount) // No API calls for n = 0
	})

	t.Run("n = 2, first bad is 1", func(t *testing.T) {
		callCount := 0
		isBadVersion := func(version int) bool {
			callCount++
			return version >= 1
		}

		result := FirstBadVersion(2, isBadVersion)

		assert.Equal(t, 1, result)
		// Binary search: check version 1 (mid=1), it's bad, set right=1, loop ends
		// Only 1 API call needed
		assert.Equal(t, 1, callCount)
	})

	t.Run("n = 2, first bad is 2", func(t *testing.T) {
		callCount := 0
		isBadVersion := func(version int) bool {
			callCount++
			return version >= 2
		}

		result := FirstBadVersion(2, isBadVersion)

		assert.Equal(t, 2, result)
		// Binary search: check version 1 (mid=1), it's good, set left=2, loop ends
		// Only 1 API call needed (worst case would be 2 if we needed to check version 2)
		assert.Equal(t, 1, callCount)
	})
}

func BenchmarkFirstBadVersion(b *testing.B) {
	// Benchmark with large n
	n := 1000000
	firstBad := 750000

	isBadVersion := func(version int) bool {
		return version >= firstBad
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FirstBadVersion(n, isBadVersion)
	}
}

func BenchmarkFirstBadVersion_WorstCase(b *testing.B) {
	// Benchmark worst case: first bad version is at the end
	n := 1000000
	firstBad := n

	isBadVersion := func(version int) bool {
		return version >= firstBad
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FirstBadVersion(n, isBadVersion)
	}
}

func BenchmarkFirstBadVersion_BestCase(b *testing.B) {
	// Benchmark best case: first bad version is at the beginning
	n := 1000000
	firstBad := 1

	isBadVersion := func(version int) bool {
		return version >= firstBad
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FirstBadVersion(n, isBadVersion)
	}
}