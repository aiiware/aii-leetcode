package leetcode

// LetterCombinations solves LeetCode problem 0017: Letter Combinations of a Phone Number
// Difficulty: Medium
// Tags: String, Backtracking, Combinatorics
//
// Given a string containing digits from 2-9 inclusive, return all possible letter
// combinations that the number could represent. Return the answer in any order.
//
// A mapping of digit to letters (just like on the telephone buttons) is given below:
// 2 -> "abc"
// 3 -> "def"
// 4 -> "ghi"
// 5 -> "jkl"
// 6 -> "mno"
// 7 -> "pqrs"
// 8 -> "tuv"
// 9 -> "wxyz"
//
// Time complexity: O(4^n * n) where n is length of digits (4 is max letters per digit)
// Space complexity: O(4^n) for storing results
func LetterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	// Phone keypad mapping
	phoneMap := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	result := []string{}
	var backtrack func(index int, currentCombination string)

	backtrack = func(index int, currentCombination string) {
		// Base case: we've processed all digits
		if index == len(digits) {
			result = append(result, currentCombination)
			return
		}

		// Get letters for current digit
		digit := rune(digits[index])
		letters := phoneMap[digit]

		// Try each letter for this digit
		for _, letter := range letters {
			backtrack(index+1, currentCombination+string(letter))
		}
	}

	backtrack(0, "")
	return result
}

// Iterative approach
func LetterCombinationsIterative(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	phoneMap := map[rune]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	result := []string{""}

	for _, digit := range digits {
		letters := phoneMap[digit]
		tempResult := []string{}

		for _, combination := range result {
			for _, letter := range letters {
				tempResult = append(tempResult, combination+string(letter))
			}
		}

		result = tempResult
	}

	return result
}
