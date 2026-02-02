package linkedlists

import (
	"testing"
	"leetcode/utils"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "Example 1 - palindrome even length",
			input:    []int{1, 2, 2, 1},
			expected: true,
		},
		{
			name:     "Example 2 - not palindrome",
			input:    []int{1, 2},
			expected: false,
		},
		{
			name:     "Single node",
			input:    []int{1},
			expected: true,
		},
		{
			name:     "Empty list",
			input:    []int{},
			expected: true,
		},
		{
			name:     "Palindrome odd length",
			input:    []int{1, 2, 3, 2, 1},
			expected: true,
		},
		{
			name:     "Not palindrome odd length",
			input:    []int{1, 2, 3, 4, 1},
			expected: false,
		},
		{
			name:     "All same values",
			input:    []int{5, 5, 5, 5, 5},
			expected: true,
		},
		{
			name:     "Long palindrome",
			input:    []int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1},
			expected: true,
		},
		{
			name:     "Long not palindrome",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: false,
		},
		{
			name:     "Palindrome with negative numbers",
			input:    []int{-1, -2, -3, -2, -1},
			expected: true,
		},
		{
			name:     "Palindrome with zeros",
			input:    []int{0, 0, 0, 0},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := utils.NewListFromSlice(tt.input)
			result := IsPalindrome(input)

			if result != tt.expected {
				t.Errorf("IsPalindrome(%v) = %v, expected %v",
					tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsPalindromeStack(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "Palindrome even length",
			input:    []int{1, 2, 2, 1},
			expected: true,
		},
		{
			name:     "Not palindrome",
			input:    []int{1, 2},
			expected: false,
		},
		{
			name:     "Single node",
			input:    []int{1},
			expected: true,
		},
		{
			name:     "Empty list",
			input:    []int{},
			expected: true,
		},
		{
			name:     "Palindrome odd length",
			input:    []int{1, 2, 3, 2, 1},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := utils.NewListFromSlice(tt.input)
			result := IsPalindromeStack(input)

			if result != tt.expected {
				t.Errorf("IsPalindromeStack(%v) = %v, expected %v",
					tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsPalindromeRecursive(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "Palindrome even length",
			input:    []int{1, 2, 2, 1},
			expected: true,
		},
		{
			name:     "Not palindrome",
			input:    []int{1, 2},
			expected: false,
		},
		{
			name:     "Single node",
			input:    []int{1},
			expected: true,
		},
		{
			name:     "Empty list",
			input:    []int{},
			expected: true,
		},
		{
			name:     "Palindrome odd length",
			input:    []int{1, 2, 3, 2, 1},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := utils.NewListFromSlice(tt.input)
			result := IsPalindromeRecursive(input)

			if result != tt.expected {
				t.Errorf("IsPalindromeRecursive(%v) = %v, expected %v",
					tt.input, result, tt.expected)
			}
		})
	}
}

func TestAllMethodsAgree(t *testing.T) {
	testCases := []struct {
		name  string
		input []int
	}{
		{"Empty", []int{}},
		{"Single", []int{1}},
		{"Two nodes palindrome", []int{1, 1}},
		{"Two nodes not palindrome", []int{1, 2}},
		{"Three nodes palindrome", []int{1, 2, 1}},
		{"Three nodes not palindrome", []int{1, 2, 3}},
		{"Four nodes palindrome", []int{1, 2, 2, 1}},
		{"Four nodes not palindrome", []int{1, 2, 3, 4}},
		{"Five nodes palindrome", []int{1, 2, 3, 2, 1}},
		{"Five nodes not palindrome", []int{1, 2, 3, 4, 5}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			input := utils.NewListFromSlice(tc.input)
			
			result1 := IsPalindrome(input)
			result2 := IsPalindromeStack(input)
			result3 := IsPalindromeRecursive(input)
			
			if result1 != result2 || result2 != result3 {
				t.Errorf("Methods disagree: O(1) space=%v, stack=%v, recursive=%v",
					result1, result2, result3)
			}
		})
	}
}

func TestIsPalindromePreservesList(t *testing.T) {
	// Test that the list is restored after checking palindrome
	original := []int{1, 2, 3, 2, 1}
	input := utils.NewListFromSlice(original)
	
	// Check palindrome (should restore list)
	isPalindrome := IsPalindrome(input)
	
	if !isPalindrome {
		t.Errorf("Expected [1,2,3,2,1] to be palindrome")
	}
	
	// Verify list is restored
	restored := input.ToSlice()
	for i, val := range original {
		if restored[i] != val {
			t.Errorf("List not restored correctly at index %d: got %v, want %v",
				i, restored, original)
			break
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	// Create a large palindrome list
	vals := make([]int, 10000)
	for i := 0; i < 5000; i++ {
		vals[i] = i
		vals[9999-i] = i
	}
	input := utils.NewListFromSlice(vals)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsPalindrome(utils.CopyList(input))
	}
}

func BenchmarkIsPalindromeStack(b *testing.B) {
	vals := make([]int, 10000)
	for i := 0; i < 5000; i++ {
		vals[i] = i
		vals[9999-i] = i
	}
	input := utils.NewListFromSlice(vals)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsPalindromeStack(utils.CopyList(input))
	}
}

func BenchmarkIsPalindromeRecursive(b *testing.B) {
	// Use smaller list for recursive benchmark
	vals := make([]int, 1000)
	for i := 0; i < 500; i++ {
		vals[i] = i
		vals[999-i] = i
	}
	input := utils.NewListFromSlice(vals)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsPalindromeRecursive(utils.CopyList(input))
	}
}