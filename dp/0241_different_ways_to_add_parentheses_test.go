package dp

import (
	"reflect"
	"sort"
	"testing"
)

func TestDiffWaysToCompute(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		expected   []int
	}{
		{
			name:       "Example 1",
			expression: "2-1-1",
			expected:   []int{0, 2},
		},
		{
			name:       "Example 2",
			expression: "2*3-4*5",
			expected:   []int{-34, -14, -10, -10, 10},
		},
		{
			name:       "Single number",
			expression: "5",
			expected:   []int{5},
		},
		{
			name:       "Simple addition",
			expression: "1+2",
			expected:   []int{3},
		},
		{
			name:       "Multiple operators",
			expression: "2*3+4",
			expected:   []int{10, 14}, // (2*(3+4))=14, ((2*3)+4)=10
		},
		{
			name:       "Complex expression",
			expression: "1+2*3",
			expected:   []int{7, 9}, // (1+(2*3))=7, ((1+2)*3)=9
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DiffWaysToCompute(tt.expression)
			
			// Sort both slices for comparison
			sort.Ints(result)
			sort.Ints(tt.expected)
			
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("DiffWaysToCompute(%q) = %v, expected %v", tt.expression, result, tt.expected)
			}
		})
	}
}

func BenchmarkDiffWaysToCompute(b *testing.B) {
	expressions := []string{
		"2-1-1",
		"2*3-4*5",
		"1+2*3+4",
		"1+2+3+4+5",
	}
	
	for _, expr := range expressions {
		b.Run(expr, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				DiffWaysToCompute(expr)
			}
		})
	}
}