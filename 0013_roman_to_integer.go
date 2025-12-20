package leetcode

// RomanToInt solves LeetCode problem 0013: Roman to Integer
// Difficulty: Easy
// Tags: Hash Table, Math, String
//
// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
//
// For example, 2 is written as II in Roman numeral, just two ones added together.
// 12 is written as XII, which is simply X + II.
// The number 27 is written as XXVII, which is XX + V + II.
//
// Usually, smaller values before larger values indicate subtraction.
// For example, IV equals 4, IX equals 9, and XL equals 40.
//
// Given a roman numeral, convert it to an integer.
//
// Time complexity: O(n), Space complexity: O(1)
func RomanToInt(s string) int {
	// Create a map of Roman characters to their values
	romanValues := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	prevValue := 0

	// Traverse the string from right to left
	runes := []rune(s)
	for i := len(runes) - 1; i >= 0; i-- {
		currentValue := romanValues[runes[i]]

		// If current value is less than previous value, it's a subtraction case
		// e.g., in IX, I (1) comes before X (10), so we subtract
		if currentValue < prevValue {
			result -= currentValue
		} else {
			result += currentValue
		}

		prevValue = currentValue
	}

	return result
}
