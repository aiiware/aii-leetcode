package main

import (
	"fmt"
	"leetcode"
)

func main() {
	fmt.Println("=== LeetCode Solutions Demo (Problems 0001-0010) ===")
	fmt.Println()

	// Two Sum examples
	fmt.Println("1. Two Sum (Problem 0001)")
	fmt.Println("---------------------------")
	
	// Example 1 from LeetCode
	nums1 := []int{2, 7, 11, 15}
	target1 := 9
	result1 := leetcode.TwoSum(nums1, target1)
	fmt.Printf("Example 1: nums = %v, target = %d\n", nums1, target1)
	fmt.Printf("  Result: indices %v → %d + %d = %d\n\n", 
		result1, nums1[result1[0]], nums1[result1[1]], target1)

	// Example 2 from LeetCode
	nums2 := []int{3, 2, 4}
	target2 := 6
	result2 := leetcode.TwoSum(nums2, target2)
	fmt.Printf("Example 2: nums = %v, target = %d\n", nums2, target2)
	fmt.Printf("  Result: indices %v → %d + %d = %d\n\n", 
		result2, nums2[result2[0]], nums2[result2[1]], target2)

	// Example 3 from LeetCode
	nums3 := []int{3, 3}
	target3 := 6
	result3 := leetcode.TwoSum(nums3, target3)
	fmt.Printf("Example 3: nums = %v, target = %d\n", nums3, target3)
	fmt.Printf("  Result: indices %v → %d + %d = %d\n\n", 
		result3, nums3[result3[0]], nums3[result3[1]], target3)

	// Add Two Numbers examples
	fmt.Println("\n2. Add Two Numbers (Problem 0002)")
	fmt.Println("----------------------------------")

	// Example 1 from LeetCode: 342 + 465 = 807
	l1a := leetcode.NewListFromSlice([]int{2, 4, 3})  // 342
	l2a := leetcode.NewListFromSlice([]int{5, 6, 4})  // 465
	resultA := leetcode.AddTwoNumbers(l1a, l2a)       // 807
	fmt.Printf("Example 1: 342 + 465 = 807\n")
	fmt.Printf("  Input: l1 = %v (342), l2 = %v (465)\n", 
		l1a.ToSlice(), l2a.ToSlice())
	fmt.Printf("  Result: %v (807)\n\n", resultA.ToSlice())

	// Example 2: 9999999 + 9999 = 10009998
	l1b := leetcode.NewListFromSlice([]int{9, 9, 9, 9, 9, 9, 9})  // 9999999
	l2b := leetcode.NewListFromSlice([]int{9, 9, 9, 9})           // 9999
	resultB := leetcode.AddTwoNumbers(l1b, l2b)                   // 10009998
	fmt.Printf("Example 2: 9999999 + 9999 = 10009998\n")
	fmt.Printf("  Input: l1 = %v (9999999), l2 = %v (9999)\n", 
		l1b.ToSlice(), l2b.ToSlice())
	fmt.Printf("  Result: %v (10009998)\n\n", resultB.ToSlice())

	// Longest Substring Without Repeating Characters examples
	fmt.Println("\n3. Longest Substring Without Repeating Characters (Problem 0003)")
	fmt.Println("------------------------------------------------------------------")

	// Example 1 from LeetCode: "abcabcbb" → 3 ("abc")
	str1 := "abcabcbb"
	resultStr1 := leetcode.LengthOfLongestSubstring(str1)
	fmt.Printf("Example 1: s = %q\n", str1)
	fmt.Printf("  Result: %d (substring: \"abc\")\n\n", resultStr1)

	// Example 2 from LeetCode: "bbbbb" → 1 ("b")
	str2 := "bbbbb"
	resultStr2 := leetcode.LengthOfLongestSubstring(str2)
	fmt.Printf("Example 2: s = %q\n", str2)
	fmt.Printf("  Result: %d (substring: \"b\")\n\n", resultStr2)

	// Example 3 from LeetCode: "pwwkew" → 3 ("wke")
	str3 := "pwwkew"
	resultStr3 := leetcode.LengthOfLongestSubstring(str3)
	fmt.Printf("Example 3: s = %q\n", str3)
	fmt.Printf("  Result: %d (substring: \"wke\")\n\n", resultStr3)

	// Median of Two Sorted Arrays examples
	fmt.Println("\n4. Median of Two Sorted Arrays (Problem 0004)")
	fmt.Println("-----------------------------------------------")

	// Example 1 from LeetCode
	nums4a1 := []int{1, 3}
	nums4a2 := []int{2}
	median1 := leetcode.FindMedianSortedArrays(nums4a1, nums4a2)
	fmt.Printf("Example 1: nums1 = %v, nums2 = %v\n", nums4a1, nums4a2)
	fmt.Printf("  Median: %.1f\n\n", median1)

	// Example 2 from LeetCode
	nums4b1 := []int{1, 2}
	nums4b2 := []int{3, 4}
	median2 := leetcode.FindMedianSortedArrays(nums4b1, nums4b2)
	fmt.Printf("Example 2: nums1 = %v, nums2 = %v\n", nums4b1, nums4b2)
	fmt.Printf("  Median: %.1f\n\n", median2)

	// Additional example
	nums4c1 := []int{0, 0}
	nums4c2 := []int{0, 0}
	median3 := leetcode.FindMedianSortedArrays(nums4c1, nums4c2)
	fmt.Printf("Example 3: nums1 = %v, nums2 = %v\n", nums4c1, nums4c2)
	fmt.Printf("  Median: %.1f\n\n", median3)

	// Longest Palindromic Substring examples
	fmt.Println("\n5. Longest Palindromic Substring (Problem 0005)")
	fmt.Println("-------------------------------------------------")

	// Example 1 from LeetCode
	str5a := "babad"
	palindrome1 := leetcode.LongestPalindrome(str5a)
	fmt.Printf("Example 1: s = %q\n", str5a)
	fmt.Printf("  Longest palindrome: %q (length: %d)\n\n", palindrome1, len(palindrome1))

	// Example 2 from LeetCode
	str5b := "cbbd"
	palindrome2 := leetcode.LongestPalindrome(str5b)
	fmt.Printf("Example 2: s = %q\n", str5b)
	fmt.Printf("  Longest palindrome: %q (length: %d)\n\n", palindrome2, len(palindrome2))

	// Additional examples
	str5c := "a"
	palindrome3 := leetcode.LongestPalindrome(str5c)
	fmt.Printf("Example 3: s = %q\n", str5c)
	fmt.Printf("  Longest palindrome: %q (length: %d)\n\n", palindrome3, len(palindrome3))

	str5d := "ac"
	palindrome4 := leetcode.LongestPalindrome(str5d)
	fmt.Printf("Example 4: s = %q\n", str5d)
	fmt.Printf("  Longest palindrome: %q (length: %d)\n\n", palindrome4, len(palindrome4))

	// Zigzag Conversion examples
	fmt.Println("\n6. Zigzag Conversion (Problem 0006)")
	fmt.Println("-------------------------------------")

	// Example 1 from LeetCode
	str6a := "PAYPALISHIRING"
	rows6a := 3
	zigzag1 := leetcode.Convert(str6a, rows6a)
	fmt.Printf("Example 1: s = %q, numRows = %d\n", str6a, rows6a)
	fmt.Printf("  Zigzag conversion: %q\n\n", zigzag1)

	// Example 2 from LeetCode
	rows6b := 4
	zigzag2 := leetcode.Convert(str6a, rows6b)
	fmt.Printf("Example 2: s = %q, numRows = %d\n", str6a, rows6b)
	fmt.Printf("  Zigzag conversion: %q\n\n", zigzag2)

	// Additional example
	str6c := "ABCDEFGHIJKLMNOP"
	rows6c := 5
	zigzag3 := leetcode.Convert(str6c, rows6c)
	fmt.Printf("Example 3: s = %q, numRows = %d\n", str6c, rows6c)
	fmt.Printf("  Zigzag conversion: %q\n\n", zigzag3)

	// Reverse Integer examples
	fmt.Println("\n7. Reverse Integer (Problem 0007)")
	fmt.Println("-----------------------------------")

	// Example 1 from LeetCode
	num7a := 123
	reversed1 := leetcode.Reverse(num7a)
	fmt.Printf("Example 1: x = %d\n", num7a)
	fmt.Printf("  Reversed: %d\n\n", reversed1)

	// Example 2 from LeetCode
	num7b := -123
	reversed2 := leetcode.Reverse(num7b)
	fmt.Printf("Example 2: x = %d\n", num7b)
	fmt.Printf("  Reversed: %d\n\n", reversed2)

	// Example 3 from LeetCode
	num7c := 120
	reversed3 := leetcode.Reverse(num7c)
	fmt.Printf("Example 3: x = %d\n", num7c)
	fmt.Printf("  Reversed: %d\n\n", reversed3)

	// Overflow example
	num7d := 1534236469
	reversed4 := leetcode.Reverse(num7d)
	fmt.Printf("Example 4: x = %d (overflow case)\n", num7d)
	fmt.Printf("  Reversed: %d (returns 0 on overflow)\n\n", reversed4)

	// String to Integer (atoi) examples
	fmt.Println("\n8. String to Integer (atoi) (Problem 0008)")
	fmt.Println("--------------------------------------------")

	// Example 1 from LeetCode
	str8a := "42"
	atoi1 := leetcode.MyAtoi(str8a)
	fmt.Printf("Example 1: s = %q\n", str8a)
	fmt.Printf("  Integer: %d\n\n", atoi1)

	// Example 2 from LeetCode
	str8b := "   -42"
	atoi2 := leetcode.MyAtoi(str8b)
	fmt.Printf("Example 2: s = %q\n", str8b)
	fmt.Printf("  Integer: %d\n\n", atoi2)

	// Example 3 from LeetCode
	str8c := "4193 with words"
	atoi3 := leetcode.MyAtoi(str8c)
	fmt.Printf("Example 3: s = %q\n", str8c)
	fmt.Printf("  Integer: %d\n\n", atoi3)

	// Example 4 from LeetCode
	str8d := "words and 987"
	atoi4 := leetcode.MyAtoi(str8d)
	fmt.Printf("Example 4: s = %q\n", str8d)
	fmt.Printf("  Integer: %d\n\n", atoi4)

	// Example 5 from LeetCode
	str8e := "-91283472332"
	atoi5 := leetcode.MyAtoi(str8e)
	fmt.Printf("Example 5: s = %q (underflow)\n", str8e)
	fmt.Printf("  Integer: %d (clamped to INT_MIN)\n\n", atoi5)

	// Palindrome Number examples
	fmt.Println("\n9. Palindrome Number (Problem 0009)")
	fmt.Println("-------------------------------------")

	// Example 1 from LeetCode
	num9a := 121
	palNum1 := leetcode.IsPalindrome(num9a)
	fmt.Printf("Example 1: x = %d\n", num9a)
	fmt.Printf("  Is palindrome: %v\n\n", palNum1)

	// Example 2 from LeetCode
	num9b := -121
	palNum2 := leetcode.IsPalindrome(num9b)
	fmt.Printf("Example 2: x = %d\n", num9b)
	fmt.Printf("  Is palindrome: %v\n\n", palNum2)

	// Example 3 from LeetCode
	num9c := 10
	palNum3 := leetcode.IsPalindrome(num9c)
	fmt.Printf("Example 3: x = %d\n", num9c)
	fmt.Printf("  Is palindrome: %v\n\n", palNum3)

	// Additional examples
	num9d := 12321
	palNum4 := leetcode.IsPalindrome(num9d)
	fmt.Printf("Example 4: x = %d\n", num9d)
	fmt.Printf("  Is palindrome: %v\n\n", palNum4)

	num9e := 0
	palNum5 := leetcode.IsPalindrome(num9e)
	fmt.Printf("Example 5: x = %d\n", num9e)
	fmt.Printf("  Is palindrome: %v\n\n", palNum5)

	// Regular Expression Matching examples
	fmt.Println("\n10. Regular Expression Matching (Problem 0010)")
	fmt.Println("------------------------------------------------")

	// Example 1 from LeetCode
	str10a := "aa"
	pattern10a := "a"
	match1 := leetcode.IsMatch(str10a, pattern10a)
	fmt.Printf("Example 1: s = %q, p = %q\n", str10a, pattern10a)
	fmt.Printf("  Matches: %v\n\n", match1)

	// Example 2 from LeetCode
	pattern10b := "a*"
	match2 := leetcode.IsMatch(str10a, pattern10b)
	fmt.Printf("Example 2: s = %q, p = %q\n", str10a, pattern10b)
	fmt.Printf("  Matches: %v\n\n", match2)

	// Example 3 from LeetCode
	str10c := "ab"
	pattern10c := ".*"
	match3 := leetcode.IsMatch(str10c, pattern10c)
	fmt.Printf("Example 3: s = %q, p = %q\n", str10c, pattern10c)
	fmt.Printf("  Matches: %v\n\n", match3)

	// Example 4 from LeetCode
	str10d := "aab"
	pattern10d := "c*a*b"
	match4 := leetcode.IsMatch(str10d, pattern10d)
	fmt.Printf("Example 4: s = %q, p = %q\n", str10d, pattern10d)
	fmt.Printf("  Matches: %v\n\n", match4)

	// Example 5 from LeetCode
	str10e := "mississippi"
	pattern10e := "mis*is*p*."
	match5 := leetcode.IsMatch(str10e, pattern10e)
	fmt.Printf("Example 5: s = %q, p = %q\n", str10e, pattern10e)
	fmt.Printf("  Matches: %v\n\n", match5)

	// Additional example
	str10f := "aaa"
	pattern10f := "ab*a*c*a"
	match6 := leetcode.IsMatch(str10f, pattern10f)
	fmt.Printf("Example 6: s = %q, p = %q\n", str10f, pattern10f)
	fmt.Printf("  Matches: %v\n\n", match6)

	fmt.Println("=== Demo Complete ===")
	fmt.Println("\nAll LeetCode problems 0001-0010 implemented successfully!")
}