package strings

import (
	"strconv"
	"strings"
)

// Codec provides methods to encode and decode strings
type Codec struct{}

// Constructor initializes a new Codec
func Constructor() Codec {
	return Codec{}
}

// Encode encodes a list of strings to a single string
// Format: length1#string1length2#string2...
func (c *Codec) Encode(strs []string) string {
	var encoded strings.Builder

	for _, s := range strs {
		// Write length followed by delimiter
		encoded.WriteString(strconv.Itoa(len(s)))
		encoded.WriteByte('#')
		// Write the string itself
		encoded.WriteString(s)
	}

	return encoded.String()
}

// Decode decodes a single string to a list of strings
func (c *Codec) Decode(s string) []string {
	if s == "" {
		return []string{}
	}
	
	var result []string
	i := 0
	n := len(s)

	for i < n {
		// Find the delimiter
		j := i
		for j < n && s[j] != '#' {
			j++
		}

		// If we reached the end without finding delimiter, break
		if j == n {
			break
		}

		// Parse the length
		length, err := strconv.Atoi(s[i:j])
		if err != nil {
			// If we can't parse the length, return what we have so far
			break
		}

		// Move past the delimiter
		i = j + 1

		// Check if we have enough characters for the string
		if i+length > n {
			// Not enough characters, return what we have so far
			break
		}

		// Extract the string
		result = append(result, s[i:i+length])

		// Move to next encoded string
		i += length
	}

	return result
}

// EncodeStrings is a convenience function that encodes strings
func EncodeStrings(strs []string) string {
	codec := Constructor()
	return codec.Encode(strs)
}

// DecodeString is a convenience function that decodes a string
func DecodeString(s string) []string {
	codec := Constructor()
	return codec.Decode(s)
}