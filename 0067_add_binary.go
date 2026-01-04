package leetcode

// AddBinary solves LeetCode problem 0067: Add Binary
// Difficulty: Easy
// Tags: Math, String, Bit Manipulation
//
// Given two binary strings a and b, return their sum as a binary string.
//
// Example 1:
// Input: a = "11", b = "1"
// Output: "100"
//
// Example 2:
// Input: a = "1010", b = "1011"
// Output: "10101"
//
// Time complexity: O(max(m,n)), Space complexity: O(max(m,n))
func AddBinary(a string, b string) string {
	// Handle empty strings
	if a == "" {
		return b
	}
	if b == "" {
		return a
	}

	// Convert strings to rune slices for easier indexing
	runesA := []rune(a)
	runesB := []rune(b)

	// Start from the least significant digit (end of strings)
	i, j := len(runesA)-1, len(runesB)-1
	carry := 0
	result := make([]rune, 0, maxIntBinary(len(runesA), len(runesB))+1)

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry

		if i >= 0 {
			if runesA[i] == '1' {
				sum += 1
			} else if runesA[i] != '0' {
				// Invalid binary digit
				return ""
			}
			i--
		}

		if j >= 0 {
			if runesB[j] == '1' {
				sum += 1
			} else if runesB[j] != '0' {
				// Invalid binary digit
				return ""
			}
			j--
		}

		// Append current digit (0 or 1)
		result = append(result, rune('0'+sum%2))
		// Update carry for next digit
		carry = sum / 2
	}

	// Reverse the result since we built it from least to most significant
	for k := 0; k < len(result)/2; k++ {
		result[k], result[len(result)-1-k] = result[len(result)-1-k], result[k]
	}

	return string(result)
}

// Helper function
func maxIntBinary(a, b int) int {
	if a > b {
		return a
	}
	return b
}