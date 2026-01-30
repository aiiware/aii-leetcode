package arrays

import (
	"testing"
)

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		name     string
		tokens   []string
		expected int
	}{
		{
			name:     "Example 1",
			tokens:   []string{"2", "1", "+", "3", "*"},
			expected: 9,
		},
		{
			name:     "Example 2",
			tokens:   []string{"4", "13", "5", "/", "+"},
			expected: 6,
		},
		{
			name:     "Example 3",
			tokens:   []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			expected: 22,
		},
		{
			name:     "Single number",
			tokens:   []string{"42"},
			expected: 42,
		},
		{
			name:     "Simple addition",
			tokens:   []string{"3", "4", "+"},
			expected: 7,
		},
		{
			name:     "Simple subtraction",
			tokens:   []string{"10", "3", "-"},
			expected: 7,
		},
		{
			name:     "Simple multiplication",
			tokens:   []string{"6", "7", "*"},
			expected: 42,
		},
		{
			name:     "Simple division",
			tokens:   []string{"15", "3", "/"},
			expected: 5,
		},
		{
			name:     "Division truncates toward zero",
			tokens:   []string{"7", "3", "/"},
			expected: 2,
		},
		{
			name:     "Negative division",
			tokens:   []string{"-7", "3", "/"},
			expected: -2,
		},
		{
			name:     "Complex expression 1",
			tokens:   []string{"3", "4", "+", "2", "*", "1", "+"},
			expected: 15, // ((3+4)*2)+1 = 15
		},
		{
			name:     "Complex expression 2",
			tokens:   []string{"4", "2", "/", "3", "*", "5", "-"},
			expected: 1, // ((4/2)*3)-5 = 1
		},
		{
			name:     "Multiple operations",
			tokens:   []string{"1", "2", "+", "3", "+", "4", "+", "5", "+"},
			expected: 15,
		},
		{
			name:     "Nested operations",
			tokens:   []string{"9", "3", "/", "2", "*", "6", "+"},
			expected: 12, // ((9/3)*2)+6 = 12
		},
		{
			name:     "Negative numbers",
			tokens:   []string{"-3", "-4", "*", "2", "/"},
			expected: 6, // ((-3)*(-4))/2 = 6
		},
		{
			name:     "Large numbers",
			tokens:   []string{"100", "200", "+", "2", "/"},
			expected: 150, // (100+200)/2 = 150
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EvalRPN(tt.tokens)
			if result != tt.expected {
				t.Errorf("EvalRPN(%v) = %d, expected %d", tt.tokens, result, tt.expected)
			}
		})
	}
}

func BenchmarkEvalRPN(b *testing.B) {
	// Create a complex expression for benchmarking
	tokens := []string{
		"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+",
		"2", "*", "3", "/", "4", "+", "5", "-", "6", "*", "7", "/", "8", "+",
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EvalRPN(tokens)
	}
}

func TestEdgeCases(t *testing.T) {
	t.Run("Division by positive number truncates toward zero", func(t *testing.T) {
		result := EvalRPN([]string{"7", "3", "/"})
		if result != 2 {
			t.Errorf("7/3 should truncate to 2, got %d", result)
		}
	})
	
	t.Run("Division by negative number truncates toward zero", func(t *testing.T) {
		result := EvalRPN([]string{"-7", "3", "/"})
		if result != -2 {
			t.Errorf("-7/3 should truncate to -2, got %d", result)
		}
	})
	
	t.Run("Division of negative by negative", func(t *testing.T) {
		result := EvalRPN([]string{"-7", "-3", "/"})
		if result != 2 {
			t.Errorf("-7/-3 should truncate to 2, got %d", result)
		}
	})
	
	t.Run("Large multiplication", func(t *testing.T) {
		result := EvalRPN([]string{"1000", "1000", "*"})
		if result != 1000000 {
			t.Errorf("1000*1000 should be 1000000, got %d", result)
		}
	})
}