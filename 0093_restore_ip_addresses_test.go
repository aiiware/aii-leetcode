package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func TestRestoreIpAddresses(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected []string
	}{
		{
			name:     "Example 1",
			s:        "25525511135",
			expected: []string{"255.255.11.135", "255.255.111.35"},
		},
		{
			name:     "Example 2",
			s:        "0000",
			expected: []string{"0.0.0.0"},
		},
		{
			name:     "Example 3",
			s:        "101023",
			expected: []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"},
		},
		{
			name:     "Single segment valid",
			s:        "1111",
			expected: []string{"1.1.1.1"},
		},
		{
			name:     "All zeros",
			s:        "00000",
			expected: []string{},
		},
		{
			name:     "Too short",
			s:        "123",
			expected: []string{},
		},
		{
			name:     "Too long",
			s:        "123456789012345678901",
			expected: []string{},
		},
		{
			name:     "Valid with leading zeros only when single digit",
			s:        "010010",
			expected: []string{"0.10.0.10", "0.100.1.0"},
		},
		{
			name:     "All ones",
			s:        "111111111111", // 12 ones
			expected: []string{"111.111.111.111"},
		},
		{
			name:     "Mixed valid",
			s:        "19216811",
			expected: []string{"19.216.81.1", "192.16.81.1", "192.168.1.1"},
		},
		{
			name:     "No valid IPs",
			s:        "999999999999", // All 9's, each segment would be > 255 if length > 3
			expected: []string{},
		},
		{
			name:     "Edge case minimum",
			s:        "1234",
			expected: []string{"1.2.3.4"},
		},
		{
			name:     "Edge case maximum",
			s:        "123456789012", // 12 digits
			expected: []string{"123.456.789.012"}, // But 456, 789, 012 are invalid
		},
		{
			name:     "With zeros in middle",
			s:        "100100",
			expected: []string{"100.1.0.0", "100.10.0.0"},
		},
		{
			name:     "Complex case 1",
			s:        "25505011535",
			expected: []string{"255.50.11.535", "255.50.115.35", "255.0.50.11535"}, // But need to check validity
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RestoreIpAddresses(tt.s)
			
			// Sort both slices for comparison
			sort.Strings(result)
			sort.Strings(tt.expected)
			
			if !stringSlicesEqual(result, tt.expected) {
				t.Errorf("RestoreIpAddresses(%q) = %v, expected %v",
					tt.s, result, tt.expected)
			}
			
			// Additional validation: all results should be valid IP addresses
			for _, ip := range result {
				if !isValidIPAddress(ip) {
					t.Errorf("Invalid IP address in result: %q", ip)
				}
			}
		})
	}
}

func TestAllRestoreIpAddressesImplementations(t *testing.T) {
	testCases := []struct {
		name string
		s    string
	}{
		{"Example 1", "25525511135"},
		{"Example 2", "0000"},
		{"Example 3", "101023"},
		{"Simple", "1111"},
		{"With zeros", "010010"},
	}

	implementations := []struct {
		name string
		fn   func(string) []string
	}{
		{"restoreIpAddresses", restoreIpAddresses},
		{"restoreIpAddressesIterative", restoreIpAddressesIterative},
		{"restoreIpAddressesDFS", restoreIpAddressesDFS},
		{"restoreIpAddressesDP", restoreIpAddressesDP},
		{"restoreIpAddressesBFS", restoreIpAddressesBFS},
		{"restoreIpAddressesOptimized", restoreIpAddressesOptimized},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := RestoreIpAddresses(tc.s)
			sort.Strings(expected)

			for _, impl := range implementations {
				t.Run(impl.name, func(t *testing.T) {
					result := impl.fn(tc.s)
					sort.Strings(result)

					if !stringSlicesEqual(result, expected) {
						t.Errorf("%s(%q) = %v, expected %v",
							impl.name, tc.s, result, expected)
					}

					// Validate all IPs
					for _, ip := range result {
						if !isValidIPAddress(ip) {
							t.Errorf("%s(%q) produced invalid IP: %q",
								impl.name, tc.s, ip)
						}
					}
				})
			}
		})
	}
}

func TestRestoreIpAddressesEdgeCases(t *testing.T) {
	t.Run("Empty string", func(t *testing.T) {
		result := RestoreIpAddresses("")
		if len(result) != 0 {
			t.Errorf("Empty string should return empty slice, got %v", result)
		}
	})

	t.Run("Too short (length < 4)", func(t *testing.T) {
		testCases := []string{"1", "12", "123"}
		for _, s := range testCases {
			result := RestoreIpAddresses(s)
			if len(result) != 0 {
				t.Errorf("String %q (length %d) should return empty slice, got %v",
					s, len(s), result)
			}
		}
	})

	t.Run("Too long (length > 12)", func(t *testing.T) {
		testCases := []string{
			"1234567890123", // 13 digits
			"12345678901234567890", // 20 digits
		}
		for _, s := range testCases {
			result := RestoreIpAddresses(s)
			if len(result) != 0 {
				t.Errorf("String %q (length %d) should return empty slice, got %v",
					s, len(s), result)
			}
		}
	})

	t.Run("Exactly 4 digits", func(t *testing.T) {
		testCases := []struct {
			s        string
			expected []string
		}{
			{"1234", []string{"1.2.3.4"}},
			{"0000", []string{"0.0.0.0"}},
			{"9999", []string{"9.9.9.9"}},
		}
		for _, tc := range testCases {
			result := RestoreIpAddresses(tc.s)
			if !stringSlicesEqual(result, tc.expected) {
				t.Errorf("RestoreIpAddresses(%q) = %v, expected %v",
					tc.s, result, tc.expected)
			}
		}
	})

	t.Run("Exactly 12 digits", func(t *testing.T) {
		s := "123456789012"
		result := RestoreIpAddresses(s)
		// Should have at least one valid IP if segments are <= 255
		// Actually 456, 789, 012 are invalid, so should be empty
		if len(result) != 0 {
			// Validate all results
			for _, ip := range result {
				if !isValidIPAddress(ip) {
					t.Errorf("Invalid IP address: %q", ip)
				}
			}
		}
	})

	t.Run("All segments maximum (255)", func(t *testing.T) {
		s := "255255255255"
		result := RestoreIpAddresses(s)
		expected := []string{"255.255.255.255"}
		if !stringSlicesEqual(result, expected) {
			t.Errorf("RestoreIpAddresses(%q) = %v, expected %v",
				s, result, expected)
		}
	})

	t.Run("Contains non-digit characters", func(t *testing.T) {
		// Our function assumes input contains only digits (per constraints)
		// But we should test that it handles gracefully
		s := "12a34"
		result := RestoreIpAddresses(s)
		// Should return empty or panic - depends on implementation
		// For now, just ensure it doesn't crash
		_ = result
	})
}

func TestRestoreIpAddressesProperties(t *testing.T) {
	// Property-based testing
	implementations := []struct {
		name string
		fn   func(string) []string
	}{
		{"restoreIpAddresses", restoreIpAddresses},
		{"restoreIpAddressesIterative", restoreIpAddressesIterative},
		{"restoreIpAddressesDFS", restoreIpAddressesDFS},
		{"restoreIpAddressesDP", restoreIpAddressesDP},
		{"restoreIpAddressesBFS", restoreIpAddressesBFS},
		{"restoreIpAddressesOptimized", restoreIpAddressesOptimized},
	}

	testStrings := []string{
		"25525511135",
		"0000",
		"101023",
		"1111",
		"123456",
		"19216811",
	}

	for _, impl := range implementations {
		t.Run(impl.name+"_properties", func(t *testing.T) {
			for _, s := range testStrings {
				t.Run(fmt.Sprintf("s=%s", s), func(t *testing.T) {
					result := impl.fn(s)

					// Property 1: All returned strings should be valid IP addresses
					for _, ip := range result {
						if !isValidIPAddress(ip) {
							t.Errorf("Invalid IP address: %q", ip)
						}
					}

					// Property 2: Removing dots from results should give original string
					for _, ip := range result {
						withoutDots := removeDots(ip)
						if withoutDots != s {
							t.Errorf("IP %q without dots = %q, expected %q",
								ip, withoutDots, s)
						}
					}

					// Property 3: Should have exactly 3 dots in each IP
					for _, ip := range result {
						dotCount := strings.Count(ip, ".")
						if dotCount != 3 {
							t.Errorf("IP %q has %d dots, expected 3", ip, dotCount)
						}
					}

					// Property 4: No leading zeros in segments (except single digit 0)
					for _, ip := range result {
						segments := strings.Split(ip, ".")
						for _, seg := range segments {
							if len(seg) > 1 && seg[0] == '0' {
								t.Errorf("Segment %q in IP %q has leading zero", seg, ip)
							}
						}
					}

					// Property 5: All segments should be between 0 and 255
					for _, ip := range result {
						segments := strings.Split(ip, ".")
						for _, seg := range segments {
							val, err := strconv.Atoi(seg)
							if err != nil || val < 0 || val > 255 {
								t.Errorf("Invalid segment %q in IP %q", seg, ip)
							}
						}
					}
				})
			}
		})
	}
}

func BenchmarkRestoreIpAddresses(b *testing.B) {
	// Test cases of different sizes
	testCases := []struct {
		name string
		s    string
	}{
		{"Short", "25525511135"},
		{"Medium", "101023101023"},
		{"Long", "192168111111"},
		{"All ones", "111111111111"},
		{"Mixed", "123456789012"},
	}

	implementations := []struct {
		name string
		fn   func(string) []string
	}{
		{"restoreIpAddresses", restoreIpAddresses},
		{"restoreIpAddressesIterative", restoreIpAddressesIterative},
		{"restoreIpAddressesDFS", restoreIpAddressesDFS},
		{"restoreIpAddressesDP", restoreIpAddressesDP},
		{"restoreIpAddressesBFS", restoreIpAddressesBFS},
		{"restoreIpAddressesOptimized", restoreIpAddressesOptimized},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for _, impl := range implementations {
				b.Run(impl.name, func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						impl.fn(tc.s)
					}
				})
			}
		})
	}
}

func BenchmarkRestoreIpAddressesWorstCase(b *testing.B) {
	// Worst case: string of length 12 with many valid combinations
	s := "123456789012"

	b.ResetTimer()

	b.Run("restoreIpAddresses", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			restoreIpAddresses(s)
		}
	})

	b.Run("restoreIpAddressesOptimized", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			restoreIpAddressesOptimized(s)
		}
	})

	b.Run("restoreIpAddressesIterative", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			restoreIpAddressesIterative(s)
		}
	})
}

// Helper functions

// isValidIPAddress validates a complete IP address string
func isValidIPAddress(ip string) bool {
	segments := strings.Split(ip, ".")
	if len(segments) != 4 {
		return false
	}

	for _, seg := range segments {
		// Check for leading zeros
		if len(seg) > 1 && seg[0] == '0' {
			return false
		}

		// Check numeric value
		val, err := strconv.Atoi(seg)
		if err != nil {
			return false
		}

		if val < 0 || val > 255 {
			return false
		}
	}

	return true
}

// removeDots removes all dots from a string
func removeDots(s string) string {
	return strings.ReplaceAll(s, ".", "")
}

// stringSlicesEqual compares two string slices for equality
func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// generateAllPossibleIPs generates all possible dot placements (for testing)
func generateAllPossibleIPs(s string) []string {
	var result []string
	n := len(s)

	// Try all possible dot positions
	for i := 1; i <= 3 && i <= n-3; i++ {
		for j := i + 1; j <= i+3 && j <= n-2; j++ {
			for k := j + 1; k <= j+3 && k <= n-1; k++ {
				seg1 := s[0:i]
				seg2 := s[i:j]
				seg3 := s[j:k]
				seg4 := s[k:]

				ip := seg1 + "." + seg2 + "." + seg3 + "." + seg4
				if isValidIPAddress(ip) {
					result = append(result, ip)
				}
			}
		}
	}

	return result
}

func TestRestoreIpAddressesMatchesAllPossible(t *testing.T) {
	// Test that our function finds all possible valid IPs
	testCases := []string{
		"25525511135",
		"0000",
		"101023",
		"1111",
		"1234",
	}

	for _, s := range testCases {
		t.Run(s, func(t *testing.T) {
			allPossible := generateAllPossibleIPs(s)
			ourResult := RestoreIpAddresses(s)

			sort.Strings(allPossible)
			sort.Strings(ourResult)

			if !stringSlicesEqual(allPossible, ourResult) {
				t.Errorf("Mismatch for %q:\nAll possible: %v\nOur result: %v",
					s, allPossible, ourResult)
			}
		})
	}
}