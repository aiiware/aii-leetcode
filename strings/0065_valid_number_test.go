package strings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNumber(t *testing.T) {
	// Test cases from LeetCode and additional cases
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		// Valid numbers
		{"Valid integer: 0", "0", true},
		{"Valid integer: 2", "2", true},
		{"Valid integer with sign: +3", "+3", true},
		{"Valid integer with sign: -1", "-1", true},
		{"Valid decimal: 0.1", "0.1", true},
		{"Valid decimal: 3.", "3.", true},
		{"Valid decimal: .1", ".1", true},
		{"Valid decimal with sign: +.8", "+.8", true},
		{"Valid decimal with sign: -.9", "-.9", true},
		{"Valid scientific: 2e10", "2e10", true},
		{"Valid scientific: -90E3", "-90E3", true},
		{"Valid scientific: 3e+7", "3e+7", true},
		{"Valid scientific: +6e-1", "+6e-1", true},
		{"Valid scientific: 53.5e93", "53.5e93", true},
		{"Valid scientific: -123.456e789", "-123.456e789", true},
		{"Valid: 005047e+6", "005047e+6", true},
		{"Valid: 46.e3", "46.e3", true},
		{"Valid: .2e81", ".2e81", true},
		{"Valid: +.8", "+.8", true},
		{"Valid: 3.", "3.", true},
		{"Valid: 3.e1", "3.e1", true},

		// Invalid numbers
		{"Empty string", "", false},
		{"Only whitespace", "   ", false},
		{"Just letter: abc", "abc", false},
		{"Just letter e", "e", false},
		{"Just letter E", "E", false},
		{"Just dot", ".", false},
		{"Just sign: +", "+", false},
		{"Just sign: -", "-", false},
		{"Two dots: 1..1", "1..1", false},
		{"Dot after e: 1e.1", "1e.1", false},
		{"Dot in exponent: 1e1.1", "1e1.1", false},
		{"Two e's: 1e1e1", "1e1e1", false},
		{"Sign in wrong place: 12e+5.4", "12e+5.4", false},
		{"Invalid character: 1a", "1a", false},
		{"Invalid character: 1 1", "1 1", false},
		{"Invalid: e9", "e9", false},
		{"Invalid: .", ".", false},
		{"Invalid: .e1", ".e1", false},
		{"Invalid: 4e+", "4e+", false},
		{"Invalid: +E3", "+E3", false},
		{"Invalid: -e3", "-e3", false},
		{"Invalid: 95a54e53", "95a54e53", false},
		{"Invalid: --6", "--6", false},
		{"Invalid: -+3", "-+3", false},
		{"Invalid: 99e2.5", "99e2.5", false},
		{"Invalid: 1e", "1e", false},
		{"Invalid: " + "e", "e", false},

		// With whitespace (should be trimmed)
		{"With leading space: ' 0.1 '", " 0.1 ", true},
		{"With trailing space: '2e10 '", "2e10 ", true},
		{"With spaces: '  -90E3  '", "  -90E3  ", true},
		{"Invalid with middle space: '1 1'", "1 1", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsNumber(tt.input)
			assert.Equal(t, tt.expected, result,
				"IsNumber(%q) = %v, expected %v",
				tt.input, result, tt.expected)
		})
	}
}

func TestIsNumber_EdgeCases(t *testing.T) {
	t.Run("Very large exponent", func(t *testing.T) {
		assert.True(t, IsNumber("1e999999"))
		assert.True(t, IsNumber("-1e-999999"))
	})

	t.Run("Multiple zeros", func(t *testing.T) {
		assert.True(t, IsNumber("000"))
		assert.True(t, IsNumber("00.00"))
		assert.True(t, IsNumber("00e00"))
		assert.True(t, IsNumber("00.00e00"))
	})

	t.Run("Decimal with no digits before dot but with exponent", func(t *testing.T) {
		assert.True(t, IsNumber(".1e1"))
		assert.False(t, IsNumber(".e1"))
	})

	t.Run("Integer with exponent but no digits after e", func(t *testing.T) {
		assert.False(t, IsNumber("1e"))
		assert.False(t, IsNumber("1E"))
		assert.False(t, IsNumber("1e+"))
		assert.False(t, IsNumber("1e-"))
	})

	t.Run("Sign in middle", func(t *testing.T) {
		assert.False(t, IsNumber("1+2"))
		assert.False(t, IsNumber("1-2"))
		assert.False(t, IsNumber("1e2+3"))
	})

	t.Run("Multiple decimal points", func(t *testing.T) {
		assert.False(t, IsNumber("1.2.3"))
		assert.False(t, IsNumber("1..2"))
		assert.False(t, IsNumber("..1"))
	})

	t.Run("Scientific notation with decimal in mantissa", func(t *testing.T) {
		assert.True(t, IsNumber("1.2e3"))
		assert.True(t, IsNumber(".2e3"))
		assert.True(t, IsNumber("2.e3"))
	})

	t.Run("Scientific notation with decimal in exponent", func(t *testing.T) {
		assert.False(t, IsNumber("1e2.3"))
		assert.False(t, IsNumber("1e.3"))
		assert.False(t, IsNumber("1e3."))
	})
}

func TestIsNumber_PropertyBased(t *testing.T) {
	// Test that valid patterns remain valid with additional valid characters
	validPatterns := []string{
		"0", "1", "123",
		"0.1", "1.0", "123.456",
		".1", "1.",
		"1e2", "1E2", "1e+2", "1e-2",
		"1.2e3", ".2e3", "2.e3",
		"+1", "-1", "+1.2", "-1.2", "+1e2", "-1e2",
	}

	for _, pattern := range validPatterns {
		t.Run("Valid pattern: "+pattern, func(t *testing.T) {
			assert.True(t, IsNumber(pattern), "Pattern %q should be valid", pattern)
		})
	}

	// Test that adding invalid characters makes valid patterns invalid
	// Note: space is NOT invalid at beginning/end (it gets trimmed), but IS invalid in middle
	// However, for property-based testing, we're adding suffixes/prefixes, so space would be invalid
	// if it appears immediately after/before the number without being trimmed.
	// Actually, let's remove space from invalidSuffixes since the implementation trims it.
	// Also, "." is valid as a suffix for some numbers (e.g., "3." is valid).
	invalidSuffixes := []string{"a", "e", "E", "+", "-", "x", "!"}
	for _, pattern := range validPatterns {
		for _, suffix := range invalidSuffixes {
			invalid := pattern + suffix
			t.Run("Invalid with suffix: "+invalid, func(t *testing.T) {
				assert.False(t, IsNumber(invalid), "Pattern %q should be invalid", invalid)
			})
		}
	}

	// Test that adding invalid prefixes makes valid patterns invalid
	// Note: space, "+", "-" at beginning are valid (trimmed or sign)
	// "." at beginning is valid for some patterns (e.g., ".1")
	// So we need a different set for prefixes
	invalidPrefixes := []string{"a", "e", "E", "x", "!"}
	for _, pattern := range validPatterns {
		for _, prefix := range invalidPrefixes {
			invalid := prefix + pattern
			t.Run("Invalid with prefix: "+invalid, func(t *testing.T) {
				assert.False(t, IsNumber(invalid), "Pattern %q should be invalid", invalid)
			})
		}
	}
}

func BenchmarkIsNumber(b *testing.B) {
	testCases := []struct {
		name  string
		input string
	}{
		{"Simple integer", "123"},
		{"Simple decimal", "123.456"},
		{"Scientific notation", "1.23e45"},
		{"Complex scientific", "-123.456e-789"},
		{"With signs", "+.8e-3"},
		{"Long number", "123456789012345678901234567890.123456789012345678901234567890e123456789012345678901234567890"},
		{"Invalid: letters", "abc"},
		{"Invalid: multiple dots", "1.2.3"},
		{"Invalid: e at end", "1e"},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsNumber(tc.input)
			}
		})
	}
}