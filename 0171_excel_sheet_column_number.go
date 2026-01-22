package leetcode

/*
171. Excel Sheet Column Number

Given a string columnTitle that represents the column title as appears in an Excel sheet, return its corresponding column number.

For example:
A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28 
...

Example 1:
Input: columnTitle = "A"
Output: 1

Example 2:
Input: columnTitle = "AB"
Output: 28

Example 3:
Input: columnTitle = "ZY"
Output: 701

Constraints:
- 1 <= columnTitle.length <= 7
- columnTitle consists only of uppercase English letters.
- columnTitle is in the range ["A", "FXSHRXW"].

Difficulty: Easy
Tags: Math, String
Companies: Microsoft, Google, Amazon, Facebook, Apple
*/

// titleToNumber converts an Excel column title to its corresponding column number.
func titleToNumber(columnTitle string) int {
	result := 0
	for i := 0; i < len(columnTitle); i++ {
		// Convert character to number: 'A' = 1, 'B' = 2, ..., 'Z' = 26
		digit := int(columnTitle[i] - 'A' + 1)
		// Multiply previous result by 26 and add current digit
		result = result*26 + digit
	}
	return result
}