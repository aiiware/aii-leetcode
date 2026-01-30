package dp

import (
	"fmt"
	"testing"
)

func TestNumTrees(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "Example 1",
			n:        3,
			expected: 5,
		},
		{
			name:     "Example 2",
			n:        1,
			expected: 1,
		},
		{
			name:     "n = 0",
			n:        0,
			expected: 0,
		},
		{
			name:     "n = 2",
			n:        2,
			expected: 2,
		},
		{
			name:     "n = 4",
			n:        4,
			expected: 14,
		},
		{
			name:     "n = 5",
			n:        5,
			expected: 42,
		},
		{
			name:     "n = 6",
			n:        6,
			expected: 132,
		},
		{
			name:     "n = 7",
			n:        7,
			expected: 429,
		},
		{
			name:     "n = 8",
			n:        8,
			expected: 1430,
		},
		{
			name:     "n = 9",
			n:        9,
			expected: 4862,
		},
		{
			name:     "n = 10",
			n:        10,
			expected: 16796,
		},
		{
			name:     "n = 15",
			n:        15,
			expected: 9694845,
		},
		{
			name:     "n = 19 (maximum per constraints)",
			n:        19,
			expected: 1767263190,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NumTrees(tt.n)
			if result != tt.expected {
				t.Errorf("NumTrees(%d) = %d, expected %d",
					tt.n, result, tt.expected)
			}
		})
	}
}

func TestAllNumTreesImplementations(t *testing.T) {
	testCases := []struct {
		name string
		n    int
	}{
		{"n=0", 0},
		{"n=1", 1},
		{"n=2", 2},
		{"n=3", 3},
		{"n=4", 4},
		{"n=5", 5},
		{"n=6", 6},
		{"n=7", 7},
		{"n=8", 8},
	}

	implementations := []struct {
		name string
		fn   func(int) int
	}{
		{"numTrees", numTrees},
		{"numTreesCatalan", numTreesCatalan},
		{"numTreesRecursive", numTreesRecursive},
		{"numTreesIterative", numTreesIterative},
		{"numTreesOptimized", numTreesOptimized},
		{"numTreesDP2", numTreesDP2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := NumTrees(tc.n)

			for _, impl := range implementations {
				t.Run(impl.name, func(t *testing.T) {
					result := impl.fn(tc.n)
					if result != expected {
						t.Errorf("%s(%d) = %d, expected %d",
							impl.name, tc.n, result, expected)
					}
				})
			}
		})
	}
}

func TestNumTreesEdgeCases(t *testing.T) {
	t.Run("n = 0", func(t *testing.T) {
		result := NumTrees(0)
		if result != 0 {
			t.Errorf("NumTrees(0) = %d, expected 0", result)
		}
	})

	t.Run("n = 1", func(t *testing.T) {
		result := NumTrees(1)
		if result != 1 {
			t.Errorf("NumTrees(1) = %d, expected 1", result)
		}
	})

	t.Run("n = 19 (maximum)", func(t *testing.T) {
		result := NumTrees(19)
		// 19th Catalan number is 1767263190
		expected := 1767263190
		if result != expected {
			t.Errorf("NumTrees(19) = %d, expected %d", result, expected)
		}
	})

	t.Run("Negative n", func(t *testing.T) {
		// Our functions should handle negative input gracefully
		result := NumTrees(-1)
		if result != 0 {
			t.Errorf("NumTrees(-1) = %d, expected 0", result)
		}
	})

	t.Run("Large n beyond constraints", func(t *testing.T) {
		// Test beyond constraints (n=20) to ensure no overflow
		result := NumTrees(20)
		// 20th Catalan number is 6564120420
		expected := 6564120420
		if result != expected {
			t.Errorf("NumTrees(20) = %d, expected %d", result, expected)
		}
	})
}

func TestNumTreesProperties(t *testing.T) {
	// Property-based testing
	implementations := []struct {
		name string
		fn   func(int) int
	}{
		{"numTrees", numTrees},
		{"numTreesCatalan", numTreesCatalan},
		{"numTreesRecursive", numTreesRecursive},
		{"numTreesIterative", numTreesIterative},
		{"numTreesOptimized", numTreesOptimized},
		{"numTreesDP2", numTreesDP2},
	}

	// Catalan numbers for n = 0..10
	catalanNumbers := []int{1, 1, 2, 5, 14, 42, 132, 429, 1430, 4862, 16796}

	for _, impl := range implementations {
		t.Run(impl.name+"_properties", func(t *testing.T) {
			// Test that our functions return 0 for n=0 (by design, not mathematically)
			if impl.fn(0) != 0 {
				t.Errorf("%s(0) = %d, expected 0", impl.name, impl.fn(0))
			}

			// Test first few Catalan numbers (for n >= 1)
			for n := 1; n <= 10; n++ {
				result := impl.fn(n)
				expected := catalanNumbers[n]
				if result != expected {
					t.Errorf("%s(%d) = %d, expected %d (C%d)",
						impl.name, n, result, expected, n)
				}
			}

			// Test recurrence relation: C(n+1) = sum(C(i)*C(n-i)) for i=0..n
			// Note: We need to use C(0)=1 for the recurrence, not impl.fn(0)=0
			for n := 1; n <= 8; n++ {
				// Compute C(n+1) using recurrence with C(0)=1
				sum := 0
				for i := 0; i <= n; i++ {
					ci := 1 // C(0) = 1
					if i > 0 {
						ci = impl.fn(i)
					}
					cnMinusI := 1 // C(0) = 1
					if n-i > 0 {
						cnMinusI = impl.fn(n - i)
					}
					sum += ci * cnMinusI
				}
				cNPlus1 := impl.fn(n + 1)
				if cNPlus1 != sum {
					t.Errorf("Recurrence failed for n=%d: C(%d)=%d, sum(C(i)*C(%d-i))=%d",
						n, n+1, cNPlus1, n, sum)
				}
			}

			// Test monotonic increase: C(n) < C(n+1) for n >= 1
			for n := 1; n <= 9; n++ {
				if impl.fn(n) >= impl.fn(n+1) {
					t.Errorf("Not monotonic: C(%d)=%d >= C(%d)=%d",
						n, impl.fn(n), n+1, impl.fn(n+1))
				}
			}
		})
	}
}

func TestNumTreesMatchesGenerateTrees(t *testing.T) {
	// Test that NumTrees matches the count from GenerateTrees (problem 95)
	// Only test small n because GenerateTrees is expensive
	for n := 1; n <= 5; n++ {
		t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
			treeCount := len(GenerateTrees(n))
			numberCount := NumTrees(n)
			if treeCount != numberCount {
				t.Errorf("Mismatch for n=%d: GenerateTrees count = %d, NumTrees = %d",
					n, treeCount, numberCount)
			}
		})
	}
}

func BenchmarkNumTrees(b *testing.B) {
	// Test cases up to n=19 (maximum per constraints)
	testCases := []struct {
		name string
		n    int
	}{
		{"n=1", 1},
		{"n=5", 5},
		{"n=10", 10},
		{"n=15", 15},
		{"n=19", 19},
		{"n=25", 25}, // Beyond constraints
	}

	implementations := []struct {
		name string
		fn   func(int) int
	}{
		{"numTrees", numTrees},
		{"numTreesCatalan", numTreesCatalan},
		{"numTreesRecursive", numTreesRecursive},
		{"numTreesIterative", numTreesIterative},
		{"numTreesOptimized", numTreesOptimized},
		{"numTreesDP2", numTreesDP2},
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

func BenchmarkNumTreesWorstCase(b *testing.B) {
	// n=19 is the worst case within constraints
	n := 19

	b.ResetTimer()

	b.Run("numTrees", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			numTrees(n)
		}
	})

	b.Run("numTreesCatalan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			numTreesCatalan(n)
		}
	})

	b.Run("numTreesOptimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			numTreesOptimized(n)
		}
	})

	b.Run("numTreesRecursive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			numTreesRecursive(n)
		}
	})
}

func BenchmarkNumTreesLarge(b *testing.B) {
	// Test beyond constraints
	n := 100

	b.ResetTimer()

	b.Run("numTreesCatalan", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			numTreesCatalan(n)
		}
	})

	b.Run("numTreesOptimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			numTreesOptimized(n)
		}
	})
}

// Test mathematical formula (may overflow for large n)
func TestNumTreesMath(t *testing.T) {
	// Only test small n to avoid overflow
	for n := 1; n <= 10; n++ {
		t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
			result := numTreesMath(n)
			expected := NumTrees(n)
			if result != expected {
				t.Errorf("numTreesMath(%d) = %d, expected %d", n, result, expected)
			}
		})
	}
}

// Helper function to compute Catalan numbers for verification
func catalanNumber(n int) int {
	if n < 0 {
		return 0
	}
	// Using the iterative formula to avoid overflow
	result := 1
	for i := 0; i < n; i++ {
		result = result * (2*n - i) / (i + 1)
	}
	return result / (n + 1)
}

func TestCatalanNumberHelper(t *testing.T) {
	// Test our helper function against known values
	tests := []struct {
		n        int
		expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 5},
		{4, 14},
		{5, 42},
		{6, 132},
		{7, 429},
		{8, 1430},
		{9, 4862},
		{10, 16796},
	}

	for _, tt := range tests {
		result := catalanNumber(tt.n)
		if result != tt.expected {
			t.Errorf("catalanNumber(%d) = %d, expected %d", tt.n, result, tt.expected)
		}
	}
}