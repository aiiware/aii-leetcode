package leetcode

import (
	"fmt"
	"testing"
)

func TestTitleToNumber(t *testing.T) {
	tests := []struct {
		columnTitle string
		expected    int
	}{
		// Basic single letters
		{"A", 1},
		{"B", 2},
		{"Z", 26},
		
		// Two-letter combinations
		{"AA", 27},  // 1*26 + 1 = 27
		{"AB", 28},  // 1*26 + 2 = 28
		{"AZ", 52},  // 1*26 + 26 = 52
		{"BA", 53},  // 2*26 + 1 = 53
		{"ZY", 701}, // 26*26 + 25 = 701
		{"ZZ", 702}, // 26*26 + 26 = 702
		
		// Three-letter combinations
		{"AAA", 703},     // 1*26*26 + 1*26 + 1 = 703
		{"AAB", 704},     // 1*26*26 + 1*26 + 2 = 704
		{"ABC", 731},     // 1*26*26 + 2*26 + 3 = 731
		{"XFD", 16384},   // 24*26*26 + 6*26 + 4 = 16384
		
		// Edge cases from constraints
		{"FXSHRXW", 2147483647}, // Maximum value for 7 letters
		
		// More examples
		{"C", 3},
		{"Y", 25},
		{"CA", 79},   // 3*26 + 1 = 79
		{"ZZZ", 18278}, // 26*26*26 + 26*26 + 26 = 18278
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test case %d: %s", i, tt.columnTitle), func(t *testing.T) {
			result := titleToNumber(tt.columnTitle)
			if result != tt.expected {
				t.Errorf("titleToNumber(%q) = %v, want %v", tt.columnTitle, result, tt.expected)
			}
		})
	}
}

// Benchmark tests for performance
func BenchmarkTitleToNumber(b *testing.B) {
	testCases := []string{"A", "AB", "ZY", "ABC", "FXSHRXW"}
	
	for _, tc := range testCases {
		b.Run(fmt.Sprintf("ColumnTitle_%s", tc), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				titleToNumber(tc)
			}
		})
	}
}

// Additional test for property-based testing
func TestTitleToNumberProperty(t *testing.T) {
	// Test that converting from number to title and back gives same result
	// (This would require implementing numberToTitle function)
	
	// For now, test some known properties:
	// 1. titleToNumber should be strictly increasing for lexicographically increasing strings
	testCases := []string{"A", "B", "Z", "AA", "AB", "AZ", "BA", "ZZ", "AAA"}
	
	for i := 1; i < len(testCases); i++ {
		prev := titleToNumber(testCases[i-1])
		curr := titleToNumber(testCases[i])
		if curr <= prev {
			t.Errorf("titleToNumber should be increasing: %q=%d <= %q=%d", 
				testCases[i], curr, testCases[i-1], prev)
		}
	}
}