package math

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlusOne(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Example 1: [1,2,3]",
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 4},
		},
		{
			name:     "Example 2: [4,3,2,1]",
			input:    []int{4, 3, 2, 1},
			expected: []int{4, 3, 2, 2},
		},
		{
			name:     "Example 3: [9]",
			input:    []int{9},
			expected: []int{1, 0},
		},
		{
			name:     "Single digit less than 9",
			input:    []int{5},
			expected: []int{6},
		},
		{
			name:     "Multiple 9s at end",
			input:    []int{1, 9, 9},
			expected: []int{2, 0, 0},
		},
		{
			name:     "All 9s",
			input:    []int{9, 9, 9},
			expected: []int{1, 0, 0, 0},
		},
		{
			name:     "No carry needed",
			input:    []int{1, 2, 4},
			expected: []int{1, 2, 5},
		},
		{
			name:     "Carry through multiple digits",
			input:    []int{1, 9, 9, 9},
			expected: []int{2, 0, 0, 0},
		},
		{
			name:     "Large number with carry",
			input:    []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 9},
			expected: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 0},
		},
		{
			name:     "Zero",
			input:    []int{0},
			expected: []int{1},
		},
		{
			name:     "Empty slice",
			input:    []int{},
			expected: []int{1},
		},
		{
			name:     "Very large number all 9s",
			input:    []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
			expected: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			name:     "Mixed with zeros",
			input:    []int{1, 0, 0, 0, 0},
			expected: []int{1, 0, 0, 0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PlusOne(tt.input)
			assert.Equal(t, tt.expected, result,
				"PlusOne(%v) = %v, expected %v",
				tt.input, result, tt.expected)
		})
	}
}

func TestPlusOne_EdgeCases(t *testing.T) {
	t.Run("Nil slice returns [1]", func(t *testing.T) {
		result := PlusOne(nil)
		assert.Equal(t, []int{1}, result)
	})

	t.Run("Single digit 0", func(t *testing.T) {
		result := PlusOne([]int{0})
		assert.Equal(t, []int{1}, result)
	})

	t.Run("Very large number with no carry", func(t *testing.T) {
		input := make([]int, 1000)
		input[999] = 5 // Last digit is 5
		expected := make([]int, 1000)
		copy(expected, input)
		expected[999] = 6

		result := PlusOne(input)
		assert.Equal(t, expected, result)
	})

	t.Run("Very large number all 9s", func(t *testing.T) {
		input := make([]int, 1000)
		for i := range input {
			input[i] = 9
		}
		expected := make([]int, 1001)
		expected[0] = 1
		// Rest are 0 by default

		result := PlusOne(input)
		assert.Equal(t, expected, result)
	})

	t.Run("Number with leading zeros (shouldn't happen per problem)", func(t *testing.T) {
		// According to problem, no leading zeros, but test anyway
		result := PlusOne([]int{0, 1, 2})
		assert.Equal(t, []int{0, 1, 3}, result)
	})
}

func TestPlusOne_PropertyBased(t *testing.T) {
	// Test that PlusOne is the inverse of converting to int, adding 1, and converting back
	testNumbers := []int{0, 1, 9, 10, 99, 100, 123, 999, 1000, 123456789, 999999999}

	for _, num := range testNumbers {
		t.Run("Property test for "+strconv.Itoa(num), func(t *testing.T) {
			// Convert number to digits
			digits := intToDigits(num)
			
			// Apply PlusOne
			result := PlusOne(digits)
			
			// Convert result back to int
			resultNum := digitsToInt(result)
			
			// Should equal original number + 1
			assert.Equal(t, num+1, resultNum,
				"PlusOne on %d (%v) should give %d, got %d (%v)",
				num, digits, num+1, resultNum, result)
		})
	}

	// Test that PlusOne doesn't modify the original slice when it needs to grow
	t.Run("Does not modify original when growing", func(t *testing.T) {
		original := []int{9, 9, 9}
		copy := make([]int, len(original))
		copyCopy := copy
		result := PlusOne(original)
		
		// Original should be unchanged
		assert.Equal(t, []int{9, 9, 9}, original)
		// Result should be new slice
		assert.Equal(t, []int{1, 0, 0, 0}, result)
		// Copy should be unchanged
		assert.Equal(t, copyCopy, copy)
	})

	t.Run("May modify original when not growing", func(t *testing.T) {
		original := []int{1, 2, 3}
		result := PlusOne(original)
		
		// When no new slice is needed, original may be modified
		// That's acceptable since problem doesn't specify
		assert.Equal(t, []int{1, 2, 4}, result)
	})
}

// Helper functions for property-based testing
func intToDigits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	
	var digits []int
	for n > 0 {
		digits = append([]int{n % 10}, digits...)
		n /= 10
	}
	return digits
}

func digitsToInt(digits []int) int {
	result := 0
	for _, digit := range digits {
		result = result*10 + digit
	}
	return result
}

func BenchmarkPlusOne(b *testing.B) {
	testCases := []struct {
		name   string
		digits []int
	}{
		{"Small no carry", []int{1, 2, 3}},
		{"Small with carry", []int{1, 2, 9}},
		{"Medium all 9s", make([]int, 100)},
		{"Large mixed", func() []int {
			digits := make([]int, 1000)
			for i := range digits {
				digits[i] = i % 10
			}
			return digits
		}()},
		{"Very large all 9s", func() []int {
			digits := make([]int, 10000)
			for i := range digits {
				digits[i] = 9
			}
			return digits
		}()},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				// Create a fresh copy for each iteration
				digits := make([]int, len(tc.digits))
				copy(digits, tc.digits)
				PlusOne(digits)
			}
		})
	}
}