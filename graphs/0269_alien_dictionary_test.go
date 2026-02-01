package graphs

import (
	"testing"
)

func TestAlienOrder(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  string
	}{
		{
			name:  "Example 1",
			words: []string{"wrt", "wrf", "er", "ett", "rftt"},
			want:  "wertf",
		},
		{
			name:  "Example 2",
			words: []string{"z", "x"},
			want:  "zx",
		},
		{
			name:  "Example 3",
			words: []string{"z", "x", "z"},
			want:  "", // Cycle detected
		},
		{
			name:  "Empty input",
			words: []string{},
			want:  "",
		},
		{
			name:  "Single word",
			words: []string{"hello"},
			want:  "ehlo", // All unique characters in any order
		},
		{
			name:  "Invalid prefix",
			words: []string{"abc", "ab"},
			want:  "", // Invalid: "ab" is prefix of "abc"
		},
		{
			name:  "Multiple valid orders",
			words: []string{"baa", "abcd", "abca", "cab", "cad"},
			want:  "bdac",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := alienOrder(tt.words)
			if tt.name == "Single word" {
				// For single word, any permutation of unique characters is valid
				if len(got) != len(tt.want) {
					t.Errorf("alienOrder(%v) = %v (length %d), want length %d", tt.words, got, len(got), len(tt.want))
				}
				// Check that all characters are unique and match expected set
				expectedChars := make(map[byte]bool)
				for i := 0; i < len(tt.want); i++ {
					expectedChars[tt.want[i]] = true
				}
				gotChars := make(map[byte]bool)
				for i := 0; i < len(got); i++ {
					gotChars[got[i]] = true
				}
				if len(gotChars) != len(expectedChars) {
					t.Errorf("alienOrder(%v) = %v has %d unique chars, want %d unique chars", tt.words, got, len(gotChars), len(expectedChars))
				}
				for ch := range expectedChars {
					if !gotChars[ch] {
						t.Errorf("alienOrder(%v) = %v missing character %c", tt.words, got, ch)
					}
				}
			} else {
				if got != tt.want {
					t.Errorf("alienOrder(%v) = %v, want %v", tt.words, got, tt.want)
				}
			}
		})
	}
}

func TestAlienOrderEdgeCases(t *testing.T) {
	tests := []struct {
		name  string
		words []string
	}{
		{
			name:  "All same words",
			words: []string{"abc", "abc", "abc"},
		},
		{
			name:  "Single character words",
			words: []string{"a", "b", "c", "d"},
		},
		{
			name:  "Mixed lengths",
			words: []string{"a", "aa", "aaa", "aaaa"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := alienOrder(tt.words)
			// For these cases, we just want to ensure no panic
			// and that we get some valid ordering
			if len(got) > 0 {
				// Verify all characters are unique in result
				seen := make(map[byte]bool)
				for i := 0; i < len(got); i++ {
					if seen[got[i]] {
						t.Errorf("alienOrder(%v) = %v has duplicate character %c", tt.words, got, got[i])
					}
					seen[got[i]] = true
				}
			}
		})
	}
}