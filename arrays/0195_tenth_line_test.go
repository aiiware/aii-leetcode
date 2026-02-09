package arrays

import (
	"fmt"
	"testing"
)

func TestTenthLine(t *testing.T) {
	tests := []struct {
		name     string
		lines    []string
		expected string
	}{
		{
			name: "10th line exists",
			lines: []string{
				"Line 1",
				"Line 2",
				"Line 3",
				"Line 4",
				"Line 5",
				"Line 6",
				"Line 7",
				"Line 8",
				"Line 9",
				"Line 10",
			},
			expected: "Line 10",
		},
		{
			name:     "less than 10 lines",
			lines:    []string{"Line 1", "Line 2", "Line 3"},
			expected: "",
		},
		{
			name:     "empty lines",
			lines:    []string{},
			expected: "",
		},
		{
			name:     "exactly 10 lines",
			lines:    make([]string, 10),
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TenthLine(tt.lines)
			if result != tt.expected {
				t.Errorf("TenthLine() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func BenchmarkTenthLine(b *testing.B) {
	lines := make([]string, 100)
	for i := 0; i < 100; i++ {
		lines[i] = fmt.Sprintf("Line %d", i+1)
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TenthLine(lines)
	}
}
