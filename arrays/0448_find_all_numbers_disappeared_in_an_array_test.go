package arrays

import (
	"reflect"
	"testing"
)

func TestFindDisappearedNumbers(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{4, 3, 2, 7, 8, 2, 3, 1},
			expected: []int{5, 6},
		},
		{
			name:     "Example 2",
			nums:     []int{1, 1},
			expected: []int{2},
		},
		{
			name:     "All numbers present",
			nums:     []int{1, 2, 3, 4, 5},
			expected: []int{},
		},
		{
			name:     "Single element missing",
			nums:     []int{1, 1, 2, 2, 3, 3, 4, 4},
			expected: []int{5},
		},
		{
			name:     "Multiple missing at ends",
			nums:     []int{2, 2, 3, 3, 4, 4},
			expected: []int{1, 5, 6},
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: []int{},
		},
		{
			name:     "Single element array with 1",
			nums:     []int{1},
			expected: []int{},
		},
		{
			name:     "Single element array with 2",
			nums:     []int{2},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of nums since the function modifies the input
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)
			
			result := FindDisappearedNumbers(numsCopy)
			
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("FindDisappearedNumbers(%v) = %v, expected %v", tt.nums, result, tt.expected)
			}
		})
	}
}

func TestFindDisappearedNumbersEdgeCases(t *testing.T) {
	// Test with maximum constraints
	t.Run("Large array", func(t *testing.T) {
		// Create an array where all numbers from 1 to 1000 are present except 500
		n := 1000
		nums := make([]int, n)
		for i := 0; i < n; i++ {
			if i < 499 {
				nums[i] = i + 1
			} else {
				// Skip 500, so from index 499 onward, we use i+2
				nums[i] = i + 2
			}
		}
		
		expected := []int{500}
		result := FindDisappearedNumbers(nums)
		
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("FindDisappearedNumbers(large array) = %v, expected %v", result, expected)
		}
	})
	
	// Test that the function doesn't panic with nil input
	t.Run("Nil input", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("FindDisappearedNumbers panicked with nil input: %v", r)
			}
		}()
		
		result := FindDisappearedNumbers(nil)
		if result == nil || len(result) != 0 {
			t.Errorf("FindDisappearedNumbers(nil) = %v, expected []", result)
		}
	})
}