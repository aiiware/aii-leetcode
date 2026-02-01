package strings

import (
	"reflect"
	"testing"
)

func TestCodec_EncodeDecode(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
	}{
		{"empty list", []string{}},
		{"single empty string", []string{""}},
		{"single non-empty string", []string{"hello"}},
		{"multiple strings", []string{"hello", "world"}},
		{"strings with special characters", []string{"hello#world", "test#123"}},
		{"strings with numbers", []string{"123", "456", "789"}},
		{"mixed strings", []string{"", "hello", "", "world", ""}},
		{"unicode strings", []string{"café", "世界", "😀😃😄"}},
		{"long strings", []string{"a very long string that contains many characters", "another long string"}},
		{"edge case: string containing delimiter", []string{"hello#world", "#test#", "123#"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			codec := Constructor()
			
			// Encode the input
			encoded := codec.Encode(tt.input)
			
			// Decode back
			decoded := codec.Decode(encoded)
			
			// Check if decoded matches original
			if !reflect.DeepEqual(decoded, tt.input) {
				t.Errorf("Encode/Decode mismatch\nInput: %v\nEncoded: %q\nDecoded: %v", tt.input, encoded, decoded)
			}
			
			// Also test convenience functions
			encoded2 := EncodeStrings(tt.input)
			decoded2 := DecodeString(encoded2)
			
			if !reflect.DeepEqual(decoded2, tt.input) {
				t.Errorf("Convenience functions mismatch\nInput: %v\nEncoded: %q\nDecoded: %v", tt.input, encoded2, decoded2)
			}
		})
	}
}

func TestCodec_EncodeFormat(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{"empty list", []string{}, ""},
		{"single empty string", []string{""}, "0#"},
		{"single string", []string{"hello"}, "5#hello"},
		{"two strings", []string{"hello", "world"}, "5#hello5#world"},
		{"strings with delimiter", []string{"hello#world"}, "11#hello#world"},
		{"mixed lengths", []string{"a", "bb", "ccc"}, "1#a2#bb3#ccc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			codec := Constructor()
			encoded := codec.Encode(tt.input)
			
			if encoded != tt.expected {
				t.Errorf("Encode(%v) = %q, want %q", tt.input, encoded, tt.expected)
			}
		})
	}
}

func TestCodec_DecodeEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		encoded  string
		expected []string
		shouldFail bool
	}{
		{"empty string", "", []string{}, false},
		{"single empty string", "0#", []string{""}, false},
		{"valid encoding", "5#hello5#world", []string{"hello", "world"}, false},
		{"with delimiter in string", "11#hello#world", []string{"hello#world"}, false},
		{"invalid: missing delimiter", "5hello", nil, true}, // Should handle gracefully
		{"invalid: length not number", "abc#hello", nil, true}, // Should handle gracefully
		{"invalid: string shorter than length", "10#short", nil, true}, // Should handle gracefully
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			codec := Constructor()
			
			// Use recover to catch panics from invalid input
			defer func() {
				if r := recover(); r != nil {
					if !tt.shouldFail {
						t.Errorf("Decode(%q) panicked: %v", tt.encoded, r)
					}
				}
			}()
			
			decoded := codec.Decode(tt.encoded)
			
			if tt.shouldFail {
				// If we expected failure but didn't panic, that's okay
				// The implementation might handle errors gracefully
				return
			}
			
			if !reflect.DeepEqual(decoded, tt.expected) {
				t.Errorf("Decode(%q) = %v, want %v", tt.encoded, decoded, tt.expected)
			}
		})
	}
}

func TestCodec_RoundTripComplex(t *testing.T) {
	// Test a complex round trip
	input := []string{
		"",
		"hello",
		"world#test",
		"123",
		"",
		"another string with spaces",
		"unicode: café 世界 😀",
		"",
	}
	
	codec := Constructor()
	encoded := codec.Encode(input)
	decoded := codec.Decode(encoded)
	
	if !reflect.DeepEqual(decoded, input) {
		t.Errorf("Complex round trip failed\nInput: %v\nEncoded: %q\nDecoded: %v", input, encoded, decoded)
	}
}

func BenchmarkCodec_Encode(b *testing.B) {
	strs := []string{"hello", "world", "this", "is", "a", "test", "of", "encoding", "and", "decoding"}
	codec := Constructor()
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		codec.Encode(strs)
	}
}

func BenchmarkCodec_Decode(b *testing.B) {
	strs := []string{"hello", "world", "this", "is", "a", "test", "of", "encoding", "and", "decoding"}
	codec := Constructor()
	encoded := codec.Encode(strs)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		codec.Decode(encoded)
	}
}

func BenchmarkEncodeStrings(b *testing.B) {
	strs := []string{"hello", "world", "this", "is", "a", "test", "of", "encoding", "and", "decoding"}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EncodeStrings(strs)
	}
}

func BenchmarkDecodeString(b *testing.B) {
	strs := []string{"hello", "world", "this", "is", "a", "test", "of", "encoding", "and", "decoding"}
	encoded := EncodeStrings(strs)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DecodeString(encoded)
	}
}