package arrays

import (
	"fmt"
	"testing"
)

func TestValidPhoneNumbers(t *testing.T) {
	tests := []struct {
		name     string
		phones   []string
		expected []string
	}{
		{
			name: "valid formats",
			phones: []string{
				"(123) 456-7890",
				"123-456-7890",
			},
			expected: []string{
				"(123) 456-7890",
				"123-456-7890",
			},
		},
		{
			name: "invalid formats",
			phones: []string{
				"1234567890",
				"123-45-6789",
				"(123)456-7890",
				"abc-def-ghij",
			},
			expected: []string{},
		},
		{
			name: "mixed valid and invalid",
			phones: []string{
				"(111) 222-3333",
				"444-555-6666",
				"invalid",
				"(999) 888-7777",
			},
			expected: []string{
				"(111) 222-3333",
				"444-555-6666",
				"(999) 888-7777",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidPhoneNumbers(tt.phones)
			fmt.Printf("Result: %v\n", result)
			fmt.Printf("Expected: %v\n", tt.expected)
			if len(result) != len(tt.expected) {
				t.Errorf("ValidPhoneNumbers() = %v, want %v", result, tt.expected)
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("ValidPhoneNumbers()[%d] = %v, want %v", i, result[i], tt.expected[i])
				}
			}
		})
	}
}

func BenchmarkValidPhoneNumbers(b *testing.B) {
	phones := []string{
		"(123) 456-7890",
		"123-456-7890",
		"invalid",
		"(999) 888-7777",
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ValidPhoneNumbers(phones)
	}
}
