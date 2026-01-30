package arrays

import (
	"strconv"
)

/*
166. Fraction to Recurring Decimal

Given two integers representing the numerator and denominator of a fraction,
return the fraction in string format.

If the fractional part is repeating, enclose the repeating part in parentheses.

If multiple answers are possible, return any of them.

It is guaranteed that the length of the answer string is less than 10^4 for all given inputs.

Example 1:
Input: numerator = 1, denominator = 2
Output: "0.5"

Example 2:
Input: numerator = 2, denominator = 1
Output: "2"

Example 3:
Input: numerator = 4, denominator = 333
Output: "0.(012)"

Constraints:
- -2^31 <= numerator, denominator <= 2^31 - 1
- denominator != 0

Difficulty: Medium
Tags: Hash Table, Math, String
Companies: Airbnb, Amazon, Cohesity, Facebook, Goldman Sachs, Google, IXL, Rubrik, Tencent, Uber
*/

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	
	// Handle sign
	var result string
	if (numerator < 0) != (denominator < 0) {
		result += "-"
	}
	
	// Convert to positive numbers using int64 to handle edge cases
	num := int64(numerator)
	den := int64(denominator)
	
	// Use absolute values
	if num < 0 {
		num = -num
	}
	if den < 0 {
		den = -den
	}
	
	// Integer part
	result += strconv.FormatInt(num/den, 10)
	num %= den
	
	// If no fractional part
	if num == 0 {
		return result
	}
	
	// Fractional part
	result += "."
	
	// Map to store remainder positions
	remainderMap := make(map[int64]int)
	
	for num != 0 {
		// If we've seen this remainder before, we have a repeating decimal
		if pos, exists := remainderMap[num]; exists {
			// Insert parentheses around the repeating part
			result = result[:pos] + "(" + result[pos:] + ")"
			break
		}
		
		// Store the current remainder position
		remainderMap[num] = len(result)
		
		// Multiply remainder by 10 for next digit
		num *= 10
		result += strconv.FormatInt(num/den, 10)
		num %= den
	}
	
	return result
}