package arrays

import (
	"fmt"
	"testing"
    "leetcode/utils"
)

func TestSubsetsWithDup(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 2},
			expected: [][]int{
				{},
				{1},
				{2},
				{1, 2},
				{2, 2},
				{1, 2, 2},
			},
		},
		{
			name: "Example 2",
			nums: []int{0},
			expected: [][]int{
				{},
				{0},
			},
		},
		{
			name: "Empty array",
			nums: []int{},
			expected: [][]int{
				{},
			},
		},
		{
			name: "All duplicates",
			nums: []int{1, 1, 1},
			expected: [][]int{
				{},
				{1},
				{1, 1},
				{1, 1, 1},
			},
		},
		{
			name: "Multiple duplicates",
			nums: []int{1, 2, 2, 3},
			expected: [][]int{
				{},
				{1},
				{2},
				{3},
				{1, 2},
				{1, 3},
				{2, 2},
				{2, 3},
				{1, 2, 2},
				{1, 2, 3},
				{2, 2, 3},
				{1, 2, 2, 3},
			},
		},
		{
			name: "Negative numbers with duplicates",
			nums: []int{-1, 0, 1, -1},
			expected: [][]int{
				{},
				{-1},
				{0},
				{1},
				{-1, -1},
				{-1, 0},
				{-1, 1},
				{0, 1},
				{-1, -1, 0},
				{-1, -1, 1},
				{-1, 0, 1},
				{-1, -1, 0, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SubsetsWithDup(tt.nums)
			
			// Sort both result and expected for comparison
			utils.SortSubsets(result)
			utils.SortSubsets(tt.expected)
			
			if !utils.SubsetsEqual(result, tt.expected) {
				t.Errorf("SubsetsWithDup(%v) = %v, expected %v", 
					tt.nums, result, tt.expected)
			}
			
			// Additional check: no duplicates in result
			if utils.HasDuplicateSubsets(result) {
				t.Errorf("SubsetsWithDup(%v) contains duplicate subsets: %v", 
					tt.nums, result)
			}
			
			// Check that all subsets are valid
			for _, subset := range result {
				if !utils.IsSubset(subset, tt.nums) {
					t.Errorf("Subset %v is not a valid subset of %v", subset, tt.nums)
				}
			}
		})
	}
}

func TestAllSubsetsWithDupImplementations(t *testing.T) {
	testCases := []struct {
		name string
		nums []int
	}{
		{"Example 1", []int{1, 2, 2}},
		{"Single element", []int{0}},
		{"All duplicates", []int{1, 1, 1}},
		{"Mixed", []int{4, 4, 4, 1, 4}},
		{"Negative", []int{-1, 0, 1, -1}},
	}

	implementations := []struct {
		name string
		fn   func([]int) [][]int
	}{
		{"subsetsWithDup", subsetsWithDup},
		{"subsetsWithDupIterative", subsetsWithDupIterative},
		{"subsetsWithDupBitmask", subsetsWithDupBitmask},
		{"subsetsWithDupDFS", subsetsWithDupDFS},
		{"subsetsWithDupBFS", subsetsWithDupBFS},
		{"subsetsWithDupOptimized", subsetsWithDupOptimized},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := SubsetsWithDup(tc.nums)
			utils.SortSubsets(expected)

			for _, impl := range implementations {
				t.Run(impl.name, func(t *testing.T) {
					result := impl.fn(tc.nums)
					utils.SortSubsets(result)

					if !utils.SubsetsEqual(result, expected) {
						t.Errorf("%s(%v) = %v, expected %v",
							impl.name, tc.nums, result, expected)
					}

					// Check for duplicates
					if utils.HasDuplicateSubsets(result) {
						t.Errorf("%s(%v) contains duplicate subsets",
							impl.name, tc.nums)
					}
				})
			}
		})
	}
}

func TestSubsetsWithDupEdgeCases(t *testing.T) {
	t.Run("Empty array", func(t *testing.T) {
		result := SubsetsWithDup([]int{})
		expected := [][]int{{}}
		if !utils.SubsetsEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Single element", func(t *testing.T) {
		result := SubsetsWithDup([]int{5})
		expected := [][]int{{}, {5}}
		utils.SortSubsets(result)
		utils.SortSubsets(expected)
		if !utils.SubsetsEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("All same elements", func(t *testing.T) {
		result := SubsetsWithDup([]int{2, 2, 2, 2})
		// Should have 5 subsets: [], [2], [2,2], [2,2,2], [2,2,2,2]
		if len(result) != 5 {
			t.Errorf("Expected 5 subsets, got %d: %v", len(result), result)
		}
		
		// Check all subsets are valid
		for _, subset := range result {
			for _, num := range subset {
				if num != 2 {
					t.Errorf("Invalid element %d in subset %v", num, subset)
				}
			}
		}
		
		// Check no duplicates
		if utils.HasDuplicateSubsets(result) {
			t.Errorf("Contains duplicate subsets: %v", result)
		}
	})

	t.Run("Maximum size (n=10)", func(t *testing.T) {
		// Test with maximum allowed size
		nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		result := SubsetsWithDup(nums)
		
		// Without duplicates, should have 2^10 = 1024 subsets
		if len(result) != 1024 {
			t.Errorf("Expected 1024 subsets for n=10 without duplicates, got %d", len(result))
		}
		
		// Check no duplicates
		if utils.HasDuplicateSubsets(result) {
			t.Errorf("Contains duplicate subsets")
		}
	})

	t.Run("With many duplicates", func(t *testing.T) {
		nums := []int{1, 1, 2, 2, 2, 3}
		result := SubsetsWithDup(nums)
		
		// Should have fewer than 2^6 = 64 subsets due to duplicates
		if len(result) >= 64 {
			t.Errorf("Should have fewer than 64 subsets due to duplicates, got %d", len(result))
		}
		
		// Check all subsets are valid
		for _, subset := range result {
			if !utils.IsSubset(subset, nums) {
				t.Errorf("Invalid subset %v for nums %v", subset, nums)
			}
		}
		
		// Check no duplicates
		if utils.HasDuplicateSubsets(result) {
			t.Errorf("Contains duplicate subsets: %v", result)
		}
	})
}

func TestSubsetsWithDupProperties(t *testing.T) {
	// Property-based testing
	implementations := []struct {
		name string
		fn   func([]int) [][]int
	}{
		{"subsetsWithDup", subsetsWithDup},
		{"subsetsWithDupIterative", subsetsWithDupIterative},
		{"subsetsWithDupDFS", subsetsWithDupDFS},
		{"subsetsWithDupBFS", subsetsWithDupBFS},
		{"subsetsWithDupOptimized", subsetsWithDupOptimized},
	}

	testArrays := [][]int{
		{},
		{1},
		{1, 2},
		{1, 1},
		{1, 2, 3},
		{1, 1, 2},
		{1, 2, 2},
		{1, 1, 2, 2},
	}

	for _, impl := range implementations {
		t.Run(impl.name+"_properties", func(t *testing.T) {
			for _, nums := range testArrays {
				t.Run(fmt.Sprintf("nums=%v", nums), func(t *testing.T) {
					result := impl.fn(nums)

					// Property 1: Should contain empty subset
					hasEmpty := false
					for _, subset := range result {
						if len(subset) == 0 {
							hasEmpty = true
							break
						}
					}
					if !hasEmpty {
						t.Errorf("Result should contain empty subset")
					}

					// Property 2: No duplicate subsets
					if utils.HasDuplicateSubsets(result) {
						t.Errorf("Contains duplicate subsets: %v", result)
					}

					// Property 3: All subsets are valid
					for _, subset := range result {
						if !utils.IsSubset(subset, nums) {
							t.Errorf("Invalid subset %v for nums %v", subset, nums)
						}
					}

					// Property 4: Number of subsets should be <= 2^n
					n := len(nums)
					if len(result) > 1<<n {
						t.Errorf("Too many subsets: %d > 2^%d = %d", 
							len(result), n, 1<<n)
					}

					// Property 5: Should contain all unique elements as single-element subsets
					uniqueElements := make(map[int]bool)
					for _, num := range nums {
						uniqueElements[num] = true
					}
					for element := range uniqueElements {
						found := false
						for _, subset := range result {
							if len(subset) == 1 && subset[0] == element {
								found = true
								break
							}
						}
						if !found {
							t.Errorf("Missing single-element subset {%d}", element)
						}
					}
				})
			}
		})
	}
}

func BenchmarkSubsetsWithDup(b *testing.B) {
	// Test cases of different sizes and duplicate patterns
	testCases := []struct {
		name string
		nums []int
	}{
		{"Small no duplicates", []int{1, 2, 3}},
		{"Small with duplicates", []int{1, 2, 2}},
		{"Medium no duplicates", []int{1, 2, 3, 4, 5, 6}},
		{"Medium with duplicates", []int{1, 1, 2, 2, 3, 3}},
		{"Large with many duplicates", []int{1, 1, 1, 2, 2, 2, 3, 3, 3}},
	}

	implementations := []struct {
		name string
		fn   func([]int) [][]int
	}{
		{"subsetsWithDup", subsetsWithDup},
		{"subsetsWithDupIterative", subsetsWithDupIterative},
		{"subsetsWithDupBitmask", subsetsWithDupBitmask},
		{"subsetsWithDupDFS", subsetsWithDupDFS},
		{"subsetsWithDupBFS", subsetsWithDupBFS},
		{"subsetsWithDupOptimized", subsetsWithDupOptimized},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for _, impl := range implementations {
				b.Run(impl.name, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						impl.fn(tc.nums)
					}
				})
			}
		})
	}
}

func BenchmarkSubsetsWithDupWorstCase(b *testing.B) {
	// Worst case: all elements are unique (maximum number of subsets)
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // n=10, 2^10=1024 subsets

	b.ResetTimer()
	
	b.Run("subsetsWithDup", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			subsetsWithDup(nums)
		}
	})
	
	b.Run("subsetsWithDupOptimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			subsetsWithDupOptimized(nums)
		}
	})
	
	b.Run("subsetsWithDupIterative", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			subsetsWithDupIterative(nums)
		}
	})
}