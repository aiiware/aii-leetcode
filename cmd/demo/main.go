package main

import (
	"fmt"
	"leetcode"
)

func main() {
	fmt.Println("=== LeetCode Solutions Demo (Problems 0001-0020) ===")
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

	// Add Two Numbers examples
	fmt.Println("\n2. Add Two Numbers (Problem 0002)")
	fmt.Println("----------------------------------")

	// Example 1 from LeetCode: 342 + 465 = 807
	l1a := leetcode.NewListFromSlice([]int{2, 4, 3})  // 342
	l2a := leetcode.NewListFromSlice([]int{5, 6, 4})  // 465
	resultA := leetcode.AddTwoNumbers(l1a, l2a)       // 807
	fmt.Printf("Example 1: 342 + 465 = 807\n")
	fmt.Printf("  Result: [2, 4, 3] + [5, 6, 4] = %v (807)\n\n", resultA.ToSlice())

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

	// Container With Most Water examples
	fmt.Println("\n11. Container With Most Water (Problem 0011)")
	fmt.Println("---------------------------------------------")

	// Example 1 from LeetCode
	height11a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	area1 := leetcode.MaxArea(height11a)
	fmt.Printf("Example 1: height = %v\n", height11a)
	fmt.Printf("  Maximum area: %d\n\n", area1)

	// Example 2 from LeetCode
	height11b := []int{1, 1}
	area2 := leetcode.MaxArea(height11b)
	fmt.Printf("Example 2: height = %v\n", height11b)
	fmt.Printf("  Maximum area: %d\n\n", area2)

	// Integer to Roman examples
	fmt.Println("\n12. Integer to Roman (Problem 0012)")
	fmt.Println("-------------------------------------")

	// Example 1
	num12a := 3
	roman1 := leetcode.IntToRoman(num12a)
	fmt.Printf("Example 1: num = %d\n", num12a)
	fmt.Printf("  Roman: %s\n\n", roman1)

	// Example 2
	num12b := 58
	roman2 := leetcode.IntToRoman(num12b)
	fmt.Printf("Example 2: num = %d\n", num12b)
	fmt.Printf("  Roman: %s\n\n", roman2)

	// Example 3
	num12c := 1994
	roman3 := leetcode.IntToRoman(num12c)
	fmt.Printf("Example 3: num = %d\n", num12c)
	fmt.Printf("  Roman: %s\n\n", roman3)

	// Roman to Integer examples
	fmt.Println("\n13. Roman to Integer (Problem 0013)")
	fmt.Println("-------------------------------------")

	// Example 1
	roman13a := "III"
	intVal1 := leetcode.RomanToInt(roman13a)
	fmt.Printf("Example 1: roman = %q\n", roman13a)
	fmt.Printf("  Integer: %d\n\n", intVal1)

	// Example 2
	roman13b := "LVIII"
	intVal2 := leetcode.RomanToInt(roman13b)
	fmt.Printf("Example 2: roman = %q\n", roman13b)
	fmt.Printf("  Integer: %d\n\n", intVal2)

	// Example 3
	roman13c := "MCMXCIV"
	intVal3 := leetcode.RomanToInt(roman13c)
	fmt.Printf("Example 3: roman = %q\n", roman13c)
	fmt.Printf("  Integer: %d\n\n", intVal3)

	// Longest Common Prefix examples
	fmt.Println("\n14. Longest Common Prefix (Problem 0014)")
	fmt.Println("------------------------------------------")

	// Example 1
	strs14a := []string{"flower", "flow", "flight"}
	prefix1 := leetcode.LongestCommonPrefix(strs14a)
	fmt.Printf("Example 1: strs = %v\n", strs14a)
	fmt.Printf("  Longest common prefix: %q\n\n", prefix1)

	// Example 2
	strs14b := []string{"dog", "racecar", "car"}
	prefix2 := leetcode.LongestCommonPrefix(strs14b)
	fmt.Printf("Example 2: strs = %v\n", strs14b)
	fmt.Printf("  Longest common prefix: %q\n\n", prefix2)

	// 3Sum examples
	fmt.Println("\n15. 3Sum (Problem 0015)")
	fmt.Println("------------------------")

	// Example 1
	nums15a := []int{-1, 0, 1, 2, -1, -4}
	triplets1 := leetcode.ThreeSum(nums15a)
	fmt.Printf("Example 1: nums = %v\n", nums15a)
	fmt.Printf("  Triplets: %v\n\n", triplets1)

	// Example 2
	nums15b := []int{0}
	triplets2 := leetcode.ThreeSum(nums15b)
	fmt.Printf("Example 2: nums = %v\n", nums15b)
	fmt.Printf("  Triplets: %v\n\n", triplets2)

	// 3Sum Closest examples
	fmt.Println("\n16. 3Sum Closest (Problem 0016)")
	fmt.Println("----------------------------------")

	// Example 1
	nums16a := []int{-1, 2, 1, -4}
	target16a := 1
	closest1 := leetcode.ThreeSumClosest(nums16a, target16a)
	fmt.Printf("Example 1: nums = %v, target = %d\n", nums16a, target16a)
	fmt.Printf("  Closest sum: %d\n\n", closest1)

	// Example 2
	nums16b := []int{0, 0, 0}
	target16b := 1
	closest2 := leetcode.ThreeSumClosest(nums16b, target16b)
	fmt.Printf("Example 2: nums = %v, target = %d\n", nums16b, target16b)
	fmt.Printf("  Closest sum: %d\n\n", closest2)

	// Letter Combinations examples
	fmt.Println("\n17. Letter Combinations of a Phone Number (Problem 0017)")
	fmt.Println("---------------------------------------------------------")

	// Example 1
	digits17a := "23"
	combinations1 := leetcode.LetterCombinations(digits17a)
	fmt.Printf("Example 1: digits = %q\n", digits17a)
	fmt.Printf("  Combinations: %v\n\n", combinations1)

	// Example 2
	digits17b := ""
	combinations2 := leetcode.LetterCombinations(digits17b)
	fmt.Printf("Example 2: digits = %q\n", digits17b)
	fmt.Printf("  Combinations: %v\n\n", combinations2)

	// 4Sum examples
	fmt.Println("\n18. 4Sum (Problem 0018)")
	fmt.Println("------------------------")

	// Example 1
	nums18a := []int{1, 0, -1, 0, -2, 2}
	target18a := 0
	quadruplets1 := leetcode.FourSum(nums18a, target18a)
	fmt.Printf("Example 1: nums = %v, target = %d\n", nums18a, target18a)
	fmt.Printf("  Quadruplets: %v\n\n", quadruplets1)

	// Example 2
	nums18b := []int{2, 2, 2, 2, 2}
	target18b := 8
	quadruplets2 := leetcode.FourSum(nums18b, target18b)
	fmt.Printf("Example 2: nums = %v, target = %d\n", nums18b, target18b)
	fmt.Printf("  Quadruplets: %v\n\n", quadruplets2)

	// Remove Nth Node From End of List examples
	fmt.Println("\n19. Remove Nth Node From End of List (Problem 0019)")
	fmt.Println("-----------------------------------------------------")

	// Example 1
	head19a := leetcode.NewListFromSlice([]int{1, 2, 3, 4, 5})
	result19a := leetcode.RemoveNthFromEnd(head19a, 2)
	fmt.Printf("Example 1: head = [1, 2, 3, 4, 5], n = 2\n")
	fmt.Printf("  Result: %v\n\n", result19a.ToSlice())

	// Example 2
	head19b := leetcode.NewListFromSlice([]int{1})
	result19b := leetcode.RemoveNthFromEnd(head19b, 1)
	fmt.Printf("Example 2: head = [1], n = 1\n")
	fmt.Printf("  Result: %v\n\n", result19b.ToSlice())

	// Valid Parentheses examples
	fmt.Println("\n20. Valid Parentheses (Problem 0020)")
	fmt.Println("-------------------------------------")

	// Example 1
	s20a := "()"
	valid1 := leetcode.IsValid(s20a)
	fmt.Printf("Example 1: s = %q\n", s20a)
	fmt.Printf("  Valid: %v\n\n", valid1)

	// Example 2
	s20b := "()[]{}"
	valid2 := leetcode.IsValid(s20b)
	fmt.Printf("Example 2: s = %q\n", s20b)
	fmt.Printf("  Valid: %v\n\n", valid2)

	// Example 3
	s20c := "(]"
	valid3 := leetcode.IsValid(s20c)
	fmt.Printf("Example 3: s = %q\n", s20c)
	fmt.Printf("  Valid: %v\n\n", valid3)

	fmt.Println("=== Demo Complete ===")
	fmt.Println("\nAll LeetCode problems 0001-0020 implemented successfully!")
}
