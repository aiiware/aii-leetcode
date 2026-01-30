package math

import (
	"fmt"
	"testing"
)

func TestCompareVersion(t *testing.T) {
	tests := []struct {
		version1 string
		version2 string
		expect   int
	}{
		{
			version1: "1.2",
			version2: "1.10",
			expect:   -1,
		},
		{
			version1: "1.01",
			version2: "1.001",
			expect:   0,
		},
		{
			version1: "1.0",
			version2: "1.0.0.0",
			expect:   0,
		},
		{
			version1: "0.1",
			version2: "1.1",
			expect:   -1,
		},
		{
			version1: "1.0.1",
			version2: "1",
			expect:   1,
		},
		{
			version1: "7.5.2.4",
			version2: "7.5.3",
			expect:   -1,
		},
		{
			version1: "1.0.0",
			version2: "1.0",
			expect:   0,
		},
		{
			version1: "2.0",
			version2: "2.0.0",
			expect:   0,
		},
		{
			version1: "1.2.3.4.5.6.7.8.9.10",
			version2: "1.2.3.4.5.6.7.8.9.9",
			expect:   1,
		},
		{
			version1: "3.0.4.10",
			version2: "3.0.4.2",
			expect:   1,
		},
		{
			version1: "1",
			version2: "1.0.0.0.0.0.0.0.0.0.0",
			expect:   0,
		},
		{
			version1: "1.0.0.0.0.0.0.0.0.0.1",
			version2: "1.0.0.0.0.0.0.0.0.0.0",
			expect:   1,
		},
		{
			version1: "100.200.300",
			version2: "100.200.300",
			expect:   0,
		},
		{
			version1: "1.0.0-alpha",
			version2: "1.0.0-beta",
			expect:   0, // Note: This test case is invalid per constraints, but included for robustness
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("Test case %d: %s vs %s", i, tt.version1, tt.version2), func(t *testing.T) {
			result := compareVersion(tt.version1, tt.version2)
			if result != tt.expect {
				t.Errorf("compareVersion(%q, %q) = %d, want %d", tt.version1, tt.version2, result, tt.expect)
			}
		})
	}
}

func BenchmarkCompareVersion(b *testing.B) {
	version1 := "1.2.3.4.5.6.7.8.9.10"
	version2 := "1.2.3.4.5.6.7.8.9.9"
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		compareVersion(version1, version2)
	}
}

func BenchmarkCompareVersionLong(b *testing.B) {
	// Create long version strings
	version1 := "1"
	version2 := "1"
	for i := 0; i < 50; i++ {
		version1 += ".0"
		version2 += ".0"
	}
	version2 += ".1" // Make version2 slightly longer
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		compareVersion(version1, version2)
	}
}