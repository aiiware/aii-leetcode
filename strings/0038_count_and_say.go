package strings

/*
38. Count and Say
https://leetcode.com/problems/count-and-say/

The count-and-say sequence is a sequence of digit strings defined by the recursive formula:
- countAndSay(1) = "1"
- countAndSay(n) is the way you would "say" the digit string from countAndSay(n-1), which is then converted into a different digit string.

To determine how you "say" a digit string, split it into the minimal number of substrings such that each substring contains exactly one unique digit. Then for each substring, say the number of digits, then say the digit. Finally, concatenate every said digit.

For example, the saying and conversion for digit string "3322251":
- "3322251" -> two 3's, three 2's, one 5, and one 1 -> "23321511"

Given a positive integer n, return the nth term of the count-and-say sequence.

Example 1:
Input: n = 1
Output: "1"
Explanation: This is the base case.

Example 2:
Input: n = 4
Output: "1211"
Explanation:
countAndSay(1) = "1"
countAndSay(2) = say "1" = one 1 = "11"
countAndSay(3) = say "11" = two 1's = "21"
countAndSay(4) = say "21" = one 2 + one 1 = "12" + "11" = "1211"

Constraints:
- 1 <= n <= 30

Difficulty: Medium
Tags: String
Companies: Facebook, Amazon
*/

// CountAndSay generates the nth term of the count-and-say sequence.
// Time Complexity: O(2^n) approximately (exponential growth)
// Space Complexity: O(2^n) for storing the result
func CountAndSay(n int) string {
	if n <= 0 {
		return ""
	}
	
	result := "1"
	
	for i := 1; i < n; i++ {
		result = nextTerm(result)
	}
	
	return result
}

// nextTerm generates the next term in the count-and-say sequence
func nextTerm(s string) string {
	var result []byte
	n := len(s)
	
	for i := 0; i < n; {
		count := 1
		// Count consecutive same characters
		for i+count < n && s[i] == s[i+count] {
			count++
		}
		
		// Append count and character
		result = append(result, byte('0'+count))
		result = append(result, s[i])
		
		i += count
	}
	
	return string(result)
}

// CountAndSayRecursive uses recursion
func CountAndSayRecursive(n int) string {
	if n == 1 {
		return "1"
	}
	
	prev := CountAndSayRecursive(n - 1)
	return nextTerm(prev)
}

// CountAndSayIterativeOptimized uses string builder for efficiency
func CountAndSayIterativeOptimized(n int) string {
	if n <= 0 {
		return ""
	}
	
	current := []byte{'1'}
	
	for i := 1; i < n; i++ {
		var next []byte
		count := 1
		
		for j := 1; j < len(current); j++ {
			if current[j] == current[j-1] {
				count++
			} else {
				next = append(next, byte('0'+count))
				next = append(next, current[j-1])
				count = 1
			}
		}
		
		// Append the last group
		next = append(next, byte('0'+count))
		next = append(next, current[len(current)-1])
		
		current = next
	}
	
	return string(current)
}