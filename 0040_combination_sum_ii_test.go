package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		expected   [][]int
	}{
		{
			name:       "Example 1",
			candidates: []int{10, 1, 2, 7, 6, 1, 5},
			target:     8,
			expected: [][]int{
				{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6},
			},
		},
		{
			name:       "Example 2",
			candidates: []int{2, 5, 2, 1, 2},
			target:     5,
			expected: [][]int{
				{1, 2, 2},
				{5},
			},
		},
		{
			name:       "Empty candidates",
			candidates: []int{},
			target:     7,
			expected:   [][]int{},
		},
		{
			name:       "Target zero",
			candidates: []int{1, 2, 3},
			target:     0,
			expected:   [][]int{{}},
		},
		{
			name:       "No solution",
			candidates: []int{2, 4, 6},
			target:     5,
			expected:   [][]int{},
		},
		{
			name:       "Single candidate matches target",
			candidates: []int{5},
			target:     5,
			expected:   [][]int{{5}},
		},
		{
			name:       "All candidates same",
			candidates: []int{2, 2, 2, 2},
			target:     4,
			expected:   [][]int{{2, 2}},
		},
		{
			name:       "Large target with duplicates",
			candidates: []int{1, 1, 2, 5, 6, 7, 10},
			target:     8,
			expected: [][]int{
				{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6},
			},
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
			result := CombinationSum2(test.candidates, test.target)
			sortResults(result)
			
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("CombinationSum2(%v, %d) = %v, expected %v", 
					test.candidates, test.target, result, expected)
			}

			// Test DP approach
			result = CombinationSum2DP(test.candidates, test.target)
			sortResults(result)
			
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("CombinationSum2DP(%v, %d) = %v, expected %v", 
					test.candidates, test.target, result, expected)
			}

			// Test iterative approach
			result = CombinationSum2Iterative(test.candidates, test.target)
			sortResults(result)
			
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("CombinationSum2Iterative(%v, %d) = %v, expected %v", 
					test.candidates, test.target, result, expected)
			}
		})
	}
}

func BenchmarkCombinationSum2(b *testing.B) {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CombinationSum2(candidates, target)
	}
}

func BenchmarkCombinationSum2DP(b *testing.B) {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CombinationSum2DP(candidates, target)
	}
}

func BenchmarkCombinationSum2Iterative(b *testing.B) {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CombinationSum2Iterative(candidates, target)
	}
}