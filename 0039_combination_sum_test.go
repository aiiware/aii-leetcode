package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			name:       "Example 1",
			candidates: []int{2, 3, 6, 7},
			target:     7,
			expected:   [][]int{{2, 2, 3}, {7}},
		},
		{
			name:       "Example 2",
			candidates: []int{2, 3, 5},
			target:     8,
			expected:   [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name:       "Example 3",
			candidates: []int{2},
			target:     1,
			expected:   [][]int{},
		},
		{
			name:       "Empty candidates",
			candidates: []int{},
			target:     7,
			expected:   [][]int{},
		},
		{
			name:       "Target zero",
			candidates: []int{2, 3, 5},
			target:     0,
			expected:   [][]int{{}},
		},
		{
			name:       "Single candidate matches target",
			candidates: []int{5},
			target:     5,
			expected:   [][]int{{5}},
		},
		{
			name:       "Multiple of same candidate",
			candidates: []int{2, 4},
			target:     8,
			expected:   [][]int{{2, 2, 2, 2}, {2, 2, 4}, {4, 4}},
		},
		{
			name:       "Large target",
			candidates: []int{2, 3},
			target:     10,
			expected:   [][]int{{2, 2, 2, 2, 2}, {2, 2, 3, 3}},
		},
	}

	// Helper function to sort results for comparison
	sortResults := func(results [][]int) {
		for _, r := range results {
			sort.Ints(r)
		}
		sort.Slice(results, func(i, j int) bool {
			if len(results[i]) != len(results[j]) {
				return len(results[i]) < len(results[j])
			}
			for k := 0; k < len(results[i]); k++ {
				if results[i][k] != results[j][k] {
					return results[i][k] < results[j][k]
				}
			}
			return false
		})
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Create a copy of expected for sorting
			expected := make([][]int, len(test.expected))
			for i, comb := range test.expected {
				expected[i] = make([]int, len(comb))
				copy(expected[i], comb)
			}
			sortResults(expected)
			
			// Test backtracking approach
			result := CombinationSum(test.candidates, test.target)
			sortResults(result)
			
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("CombinationSum(%v, %d) = %v, expected %v", 
					test.candidates, test.target, result, expected)
			}

			// Test DP approach
			result = CombinationSumDP(test.candidates, test.target)
			sortResults(result)
			
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("CombinationSumDP(%v, %d) = %v, expected %v", 
					test.candidates, test.target, result, expected)
			}

			// Test iterative approach
			result = CombinationSumIterative(test.candidates, test.target)
			sortResults(result)
			
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("CombinationSumIterative(%v, %d) = %v, expected %v", 
					test.candidates, test.target, result, expected)
			}
		})
	}
}

func BenchmarkCombinationSum(b *testing.B) {
	candidates := []int{2, 3, 5, 7}
	target := 15
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CombinationSum(candidates, target)
	}
}

func BenchmarkCombinationSumDP(b *testing.B) {
	candidates := []int{2, 3, 5, 7}
	target := 15
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CombinationSumDP(candidates, target)
	}
}

func BenchmarkCombinationSumIterative(b *testing.B) {
	candidates := []int{2, 3, 5, 7}
	target := 15
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CombinationSumIterative(candidates, target)
	}
}