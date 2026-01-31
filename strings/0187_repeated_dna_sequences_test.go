package strings

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindRepeatedDnaSequences(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Example 1",
			input:    "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT",
			expected: []string{"AAAAACCCCC", "CCCCCAAAAA"},
		},
		{
			name:     "Example 2",
			input:    "AAAAAAAAAAAAA",
			expected: []string{"AAAAAAAAAA"},
		},
		{
			name:     "No repeats",
			input:    "ACGTACGTAC",
			expected: []string{},
		},
		{
			name:     "Short string",
			input:    "ACGT",
			expected: []string{},
		},
		{
			name:     "Exactly 10 characters, no repeats",
			input:    "ACGTACGTAC",
			expected: []string{},
		},
		{
			name:     "Exactly 10 characters, all same",
			input:    "AAAAAAAAAA",
			expected: []string{}, // A 10-char string has only one 10-letter substring, so it can't occur more than once
		},
		{
			name:     "Multiple repeats",
			input:    "ACGTACGTACGTACGTACGT",
			expected: []string{"ACGTACGTAC", "CGTACGTACG", "GTACGTACGT", "TACGTACGTA"},
		},
		{
			name:     "Overlapping repeats",
			input:    "AAAAAAAAAAA", // 11 A's
			expected: []string{"AAAAAAAAAA"},
		},
		{
			name:     "Mixed repeats",
			input:    "ACGTACGTACACGTACGTAC",
			expected: []string{"ACGTACGTAC"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindRepeatedDnaSequences(tt.input)
			
			// Sort both slices for comparison
			sort.Strings(result)
			sort.Strings(tt.expected)
			
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("FindRepeatedDnaSequences(%q) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFindRepeatedDnaSequencesAllImplementations(t *testing.T) {
	testCases := []struct {
		name  string
		input string
	}{
		{"Example 1", "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"},
		{"Example 2", "AAAAAAAAAAAAA"},
		{"No repeats", "ACGTACGTAC"},
		{"Short string", "ACGT"},
		{"Exactly 10 characters", "ACGTACGTAC"},
		{"All same 10 chars", "AAAAAAAAAA"},
		{"Multiple repeats", "ACGTACGTACGTACGTACGT"},
		{"Overlapping repeats", "AAAAAAAAAAA"},
		{"Mixed repeats", "ACGTACGTACACGTACGTAC"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Test all implementations
			result1 := FindRepeatedDnaSequences(tc.input)
			result2 := findRepeatedDnaSequencesTwoSets(tc.input)
			result3 := findRepeatedDnaSequencesBitManipulation(tc.input)
			result4 := findRepeatedDnaSequencesRollingHash(tc.input)

			// Sort all results for comparison
			sort.Strings(result1)
			sort.Strings(result2)
			sort.Strings(result3)
			sort.Strings(result4)

			// All implementations should return the same result
			if !reflect.DeepEqual(result1, result2) {
				t.Errorf("Implementation mismatch: HashMap=%v, TwoSets=%v", result1, result2)
			}
			if !reflect.DeepEqual(result1, result3) {
				t.Errorf("Implementation mismatch: HashMap=%v, BitManipulation=%v", result1, result3)
			}
			if !reflect.DeepEqual(result1, result4) {
				t.Errorf("Implementation mismatch: HashMap=%v, RollingHash=%v", result1, result4)
			}
		})
	}
}

func TestIsValidDnaString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"Valid DNA", "ACGT", true},
		{"Valid DNA long", "ACGTACGTACGT", true},
		{"Invalid character", "ACGTX", false},
		{"Lowercase", "acgt", false},
		{"Mixed case", "AcGT", false},
		{"Numbers", "ACGT123", false},
		{"Special chars", "ACGT!@#", false},
		{"Spaces", "AC GT", false},
		{"Empty string", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidDnaString(tt.input)
			if result != tt.expected {
				t.Errorf("isValidDnaString(%q) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFindRepeatedDnaSequencesEdgeCases(t *testing.T) {
	t.Run("Very long string with many repeats", func(t *testing.T) {
		// Create a pattern that repeats every 10 characters
		pattern := "ACGTACGTAC"
		longStr := ""
		for i := 0; i < 100; i++ {
			longStr += pattern
		}
		result := FindRepeatedDnaSequences(longStr)
		// Should find many overlapping repeats
		// The pattern itself should be in the result
		foundPattern := false
		for _, seq := range result {
			if seq == pattern {
				foundPattern = true
				break
			}
		}
		if !foundPattern {
			t.Errorf("Expected to find pattern %s in result, got %v", pattern, result)
		}
		// There should be many repeats (at least 10)
		if len(result) < 10 {
			t.Errorf("Expected many repeats in long repeated string, got only %d: %v", len(result), result)
		}
	})

	t.Run("String with all possible 10-mers", func(t *testing.T) {
		// This would be a very long test, so we'll just test the concept
		// with a shorter version
		shortStr := "ACGTACGTACGT"
		result := FindRepeatedDnaSequences(shortStr)
		// In a 12-character string, we can have at most 3 overlapping 10-mers
		// and they would all be different, so result should be empty
		if len(result) > 0 {
			t.Errorf("Expected no repeats in 12-char string, got %v", result)
		}
	})

	t.Run("Maximum length string", func(t *testing.T) {
		// Create a string of length 100,000 (max constraint)
		// We'll use a smaller version for testing
		maxStr := ""
		for i := 0; i < 1000; i++ {
			maxStr += "ACGT"
		}
		result := FindRepeatedDnaSequences(maxStr)
		// Should not panic and should return reasonable results
		if result == nil {
			t.Error("Result should not be nil")
		}
	})
}

func BenchmarkFindRepeatedDnaSequencesHashMap(b *testing.B) {
	// Create a test string
	testStr := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindRepeatedDnaSequences(testStr)
	}
}

func BenchmarkFindRepeatedDnaSequencesTwoSets(b *testing.B) {
	testStr := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findRepeatedDnaSequencesTwoSets(testStr)
	}
}

func BenchmarkFindRepeatedDnaSequencesBitManipulation(b *testing.B) {
	testStr := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findRepeatedDnaSequencesBitManipulation(testStr)
	}
}

func BenchmarkFindRepeatedDnaSequencesRollingHash(b *testing.B) {
	testStr := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		findRepeatedDnaSequencesRollingHash(testStr)
	}
}

func BenchmarkFindRepeatedDnaSequencesLongString(b *testing.B) {
	// Create a longer string for benchmarking
	longStr := ""
	for i := 0; i < 100; i++ {
		longStr += "ACGTACGTAC"
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindRepeatedDnaSequences(longStr)
	}
}