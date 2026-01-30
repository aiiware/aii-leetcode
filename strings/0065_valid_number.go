package strings

// IsNumber solves LeetCode problem 0065: Valid Number
// Difficulty: Hard
// Tags: String
//
// A valid number can be split up into these components (in order):
// 1. A decimal number or an integer.
// 2. (Optional) An 'e' or 'E', followed by an integer.
//
// A decimal number can be split up into these components (in order):
// 1. (Optional) A sign character ('+' or '-').
// 2. One of the following formats:
//    a. At least one digit, followed by a dot '.'.
//    b. At least one digit, followed by a dot '.', followed by at least one digit.
//    c. A dot '.', followed by at least one digit.
//
// An integer can be split up into these components (in order):
// 1. (Optional) A sign character ('+' or '-').
// 2. At least one digit.
//
// Time complexity: O(n), Space complexity: O(1)
func IsNumber(s string) bool {
	if len(s) == 0 {
		return false
	}

	// Remove leading and trailing whitespace
	// Note: In Go, strings are immutable, so we'll process with indices
	start, end := 0, len(s)-1
	for start <= end && s[start] == ' ' {
		start++
	}
	for end >= start && s[end] == ' ' {
		end--
	}
	if start > end {
		return false
	}

	// Process the trimmed string
	str := s[start : end+1]

	seenDigit := false
	seenDot := false
	seenE := false

	for i, ch := range str {
		if ch >= '0' && ch <= '9' {
			seenDigit = true
		} else if ch == '+' || ch == '-' {
			// Sign can only appear at beginning or right after 'e'/'E'
			if i > 0 && (str[i-1] != 'e' && str[i-1] != 'E') {
				return false
			}
		} else if ch == '.' {
			// Dot cannot appear after 'e'/'E' and cannot appear twice
			if seenDot || seenE {
				return false
			}
			seenDot = true
		} else if ch == 'e' || ch == 'E' {
			// 'e'/'E' cannot appear twice and must have digit before it
			if seenE || !seenDigit {
				return false
			}
			seenE = true
			seenDigit = false // Reset for the exponent part
		} else {
			// Any other character is invalid
			return false
		}
	}

	// Must end with a digit (unless it's a valid decimal like "3." which LeetCode accepts)
	// Actually, according to LeetCode, "3." is valid, "3.e1" is valid, etc.
	// So we need to check: if we saw 'e'/'E', must have digit after it
	if seenE {
		return seenDigit
	}
	
	// If no 'e'/'E', must have at least one digit
	// And if there's a dot, must have at least one digit somewhere
	return seenDigit
}