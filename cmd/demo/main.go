package main

import (
	"fmt"
	"leetcode"
)

func main() {
	fmt.Println("=== LeetCode Solutions Demo (Problems 0001-0050) ===")
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

	// Merge Two Sorted Lists examples
	fmt.Println("\n21. Merge Two Sorted Lists (Problem 0021)")
	fmt.Println("------------------------------------------")

	// Example 1
	list21a1 := leetcode.NewListFromSlice([]int{1, 2, 4})
	list21a2 := leetcode.NewListFromSlice([]int{1, 3, 4})
	merged1 := leetcode.MergeTwoLists(list21a1, list21a2)
	fmt.Printf("Example 1: list1 = [1, 2, 4], list2 = [1, 3, 4]\n")
	fmt.Printf("  Merged: %v\n\n", merged1.ToSlice())

	// Example 2
	list21b1 := leetcode.NewListFromSlice([]int{})
	list21b2 := leetcode.NewListFromSlice([]int{0})
	merged2 := leetcode.MergeTwoLists(list21b1, list21b2)
	fmt.Printf("Example 2: list1 = [], list2 = [0]\n")
	fmt.Printf("  Merged: %v\n\n", merged2.ToSlice())

	// Generate Parentheses examples
	fmt.Println("\n22. Generate Parentheses (Problem 0022)")
	fmt.Println("----------------------------------------")

	// Example 1
	n22a := 3
	parentheses1 := leetcode.GenerateParenthesis(n22a)
	fmt.Printf("Example 1: n = %d\n", n22a)
	fmt.Printf("  Generated: %v (length: %d)\n\n", len(parentheses1), len(parentheses1))

	// Example 2
	n22b := 1
	parentheses2 := leetcode.GenerateParenthesis(n22b)
	fmt.Printf("Example 2: n = %d\n", n22b)
	fmt.Printf("  Generated: %v\n\n", parentheses2)

	// Merge k Sorted Lists examples
	fmt.Println("\n23. Merge k Sorted Lists (Problem 0023)")
	fmt.Println("----------------------------------------")

	// Example 1
	lists23a := []*leetcode.ListNode{
		leetcode.NewListFromSlice([]int{1, 4, 5}),
		leetcode.NewListFromSlice([]int{1, 3, 4}),
		leetcode.NewListFromSlice([]int{2, 6}),
	}
	mergedK1 := leetcode.MergeKLists(lists23a)
	fmt.Printf("Example 1: lists = [[1,4,5], [1,3,4], [2,6]]\n")
	fmt.Printf("  Merged: %v\n\n", mergedK1.ToSlice())

	// Example 2
	lists23b := []*leetcode.ListNode{}
	mergedK2 := leetcode.MergeKLists(lists23b)
	fmt.Printf("Example 2: lists = []\n")
	fmt.Printf("  Merged: %v\n\n", mergedK2.ToSlice())

	// Swap Nodes in Pairs examples
	fmt.Println("\n24. Swap Nodes in Pairs (Problem 0024)")
	fmt.Println("----------------------------------------")

	// Example 1
	list24a := leetcode.NewListFromSlice([]int{1, 2, 3, 4})
	swapped1 := leetcode.SwapPairs(list24a)
	fmt.Printf("Example 1: head = [1, 2, 3, 4]\n")
	fmt.Printf("  Swapped: %v\n\n", swapped1.ToSlice())

	// Example 2
	list24b := leetcode.NewListFromSlice([]int{1})
	swapped2 := leetcode.SwapPairs(list24b)
	fmt.Printf("Example 2: head = [1]\n")
	fmt.Printf("  Swapped: %v\n\n", swapped2.ToSlice())

	// Reverse Nodes in k-Group examples
	fmt.Println("\n25. Reverse Nodes in k-Group (Problem 0025)")
	fmt.Println("--------------------------------------------")

	// Example 1
	list25a := leetcode.NewListFromSlice([]int{1, 2, 3, 4, 5})
	reversedK1 := leetcode.ReverseKGroup(list25a, 2)
	fmt.Printf("Example 1: head = [1, 2, 3, 4, 5], k = 2\n")
	fmt.Printf("  Reversed: %v\n\n", reversedK1.ToSlice())

	// Example 2
	list25b := leetcode.NewListFromSlice([]int{1, 2, 3, 4, 5})
	reversedK2 := leetcode.ReverseKGroup(list25b, 3)
	fmt.Printf("Example 2: head = [1, 2, 3, 4, 5], k = 3\n")
	fmt.Printf("  Reversed: %v\n\n", reversedK2.ToSlice())

	// Remove Duplicates from Sorted Array examples
	fmt.Println("\n26. Remove Duplicates from Sorted Array (Problem 0026)")
	fmt.Println("-------------------------------------------------------")

	// Example 1
	nums26a := []int{1, 1, 2}
	k26a := leetcode.RemoveDuplicates(nums26a)
	fmt.Printf("Example 1: nums = [1, 1, 2]\n")
	fmt.Printf("  Result: k = %d, nums = %v\n\n", k26a, nums26a[:k26a])

	// Example 2
	nums26b := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k26b := leetcode.RemoveDuplicates(nums26b)
	fmt.Printf("Example 2: nums = [0, 0, 1, 1, 1, 2, 2, 3, 3, 4]\n")
	fmt.Printf("  Result: k = %d, nums = %v\n\n", k26b, nums26b[:k26b])

	// Remove Element examples
	fmt.Println("\n27. Remove Element (Problem 0027)")
	fmt.Println("-----------------------------------")

	// Example 1
	nums27a := []int{3, 2, 2, 3}
	val27a := 3
	k27a := leetcode.RemoveElement(nums27a, val27a)
	fmt.Printf("Example 1: nums = [3, 2, 2, 3], val = %d\n", val27a)
	fmt.Printf("  Result: k = %d, nums = %v\n\n", k27a, nums27a[:k27a])

	// Example 2
	nums27b := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val27b := 2
	k27b := leetcode.RemoveElement(nums27b, val27b)
	fmt.Printf("Example 2: nums = [0, 1, 2, 2, 3, 0, 4, 2], val = %d\n", val27b)
	fmt.Printf("  Result: k = %d, nums = %v\n\n", k27b, nums27b[:k27b])

	// Find the Index of the First Occurrence in a String examples
	fmt.Println("\n28. Find the Index of the First Occurrence in a String (Problem 0028)")
	fmt.Println("------------------------------------------------------------------------")

	// Example 1
	haystack28a := "sadbutsad"
	needle28a := "sad"
	index28a := leetcode.StrStr(haystack28a, needle28a)
	fmt.Printf("Example 1: haystack = %q, needle = %q\n", haystack28a, needle28a)
	fmt.Printf("  Result: index = %d\n\n", index28a)

	// Example 2
	haystack28b := "leetcode"
	needle28b := "leeto"
	index28b := leetcode.StrStr(haystack28b, needle28b)
	fmt.Printf("Example 2: haystack = %q, needle = %q\n", haystack28b, needle28b)
	fmt.Printf("  Result: index = %d\n\n", index28b)

	// Divide Two Integers examples
	fmt.Println("\n29. Divide Two Integers (Problem 0029)")
	fmt.Println("---------------------------------------")

	// Example 1
	dividend29a := 10
	divisor29a := 3
	quotient29a := leetcode.Divide(dividend29a, divisor29a)
	fmt.Printf("Example 1: dividend = %d, divisor = %d\n", dividend29a, divisor29a)
	fmt.Printf("  Result: %d / %d = %d\n\n", dividend29a, divisor29a, quotient29a)

	// Example 2
	dividend29b := 7
	divisor29b := -3
	quotient29b := leetcode.Divide(dividend29b, divisor29b)
	fmt.Printf("Example 2: dividend = %d, divisor = %d\n", dividend29b, divisor29b)
	fmt.Printf("  Result: %d / %d = %d\n\n", dividend29b, divisor29b, quotient29b)

	// Substring with Concatenation of All Words examples
	fmt.Println("\n30. Substring with Concatenation of All Words (Problem 0030)")
	fmt.Println("--------------------------------------------------------------")

	// Example 1
	s30a := "barfoothefoobarman"
	words30a := []string{"foo", "bar"}
	indices30a := leetcode.FindSubstring(s30a, words30a)
	fmt.Printf("Example 1: s = %q, words = %v\n", s30a, words30a)
	fmt.Printf("  Result: indices = %v\n\n", indices30a)

	// Example 2
	s30b := "wordgoodgoodgoodbestword"
	words30b := []string{"word", "good", "best", "word"}
	indices30b := leetcode.FindSubstring(s30b, words30b)
	fmt.Printf("Example 2: s = %q, words = %v\n", s30b, words30b)
	fmt.Printf("  Result: indices = %v\n\n", indices30b)

	// Next Permutation examples
	fmt.Println("\n31. Next Permutation (Problem 0031)")
	fmt.Println("-------------------------------------")

	// Example 1
	nums31a := []int{1, 2, 3}
	leetcode.NextPermutation(nums31a)
	fmt.Printf("Example 1: nums = [1, 2, 3]\n")
	fmt.Printf("  Next permutation: %v\n\n", nums31a)

	// Example 2
	nums31b := []int{3, 2, 1}
	leetcode.NextPermutation(nums31b)
	fmt.Printf("Example 2: nums = [3, 2, 1]\n")
	fmt.Printf("  Next permutation: %v\n\n", nums31b)

	// Longest Valid Parentheses examples
	fmt.Println("\n32. Longest Valid Parentheses (Problem 0032)")
	fmt.Println("---------------------------------------------")

	// Example 1
	s32a := "(()"
	longest32a := leetcode.LongestValidParentheses(s32a)
	fmt.Printf("Example 1: s = %q\n", s32a)
	fmt.Printf("  Longest valid parentheses length: %d\n\n", longest32a)

	// Example 2
	s32b := ")()())"
	longest32b := leetcode.LongestValidParentheses(s32b)
	fmt.Printf("Example 2: s = %q\n", s32b)
	fmt.Printf("  Longest valid parentheses length: %d\n\n", longest32b)

	// Search in Rotated Sorted Array examples
	fmt.Println("\n33. Search in Rotated Sorted Array (Problem 0033)")
	fmt.Println("---------------------------------------------------")

	// Example 1
	nums33a := []int{4, 5, 6, 7, 0, 1, 2}
	target33a := 0
	result33a := leetcode.SearchInRotatedSortedArray(nums33a, target33a)
	fmt.Printf("Example 1: nums = %v, target = %d\n", nums33a, target33a)
	fmt.Printf("  Index: %d\n\n", result33a)

	// Example 2
	nums33b := []int{4, 5, 6, 7, 0, 1, 2}
	target33b := 3
	result33b := leetcode.SearchInRotatedSortedArray(nums33b, target33b)
	fmt.Printf("Example 2: nums = %v, target = %d\n", nums33b, target33b)
	fmt.Printf("  Index: %d\n\n", result33b)

	// Find First and Last Position examples
	fmt.Println("\n34. Find First and Last Position of Element in Sorted Array (Problem 0034)")
	fmt.Println("------------------------------------------------------------------------------------")

	// Example 1
	nums34a := []int{5, 7, 7, 8, 8, 10}
	target34a := 8
	result34a := leetcode.SearchRange(nums34a, target34a)
	fmt.Printf("Example 1: nums = %v, target = %d\n", nums34a, target34a)
	fmt.Printf("  Range: %v\n\n", result34a)

	// Example 2
	nums34b := []int{5, 7, 7, 8, 8, 10}
	target34b := 6
	result34b := leetcode.SearchRange(nums34b, target34b)
	fmt.Printf("Example 2: nums = %v, target = %d\n", nums34b, target34b)
	fmt.Printf("  Range: %v\n\n", result34b)

	// Search Insert Position examples
	fmt.Println("\n35. Search Insert Position (Problem 0035)")
	fmt.Println("------------------------------------------")

	// Example 1
	nums35a := []int{1, 3, 5, 6}
	target35a := 5
	result35a := leetcode.SearchInsert(nums35a, target35a)
	fmt.Printf("Example 1: nums = %v, target = %d\n", nums35a, target35a)
	fmt.Printf("  Index: %d\n\n", result35a)

	// Example 2
	nums35b := []int{1, 3, 5, 6}
	target35b := 7
	result35b := leetcode.SearchInsert(nums35b, target35b)
	fmt.Printf("Example 2: nums = %v, target = %d\n", nums35b, target35b)
	fmt.Printf("  Index: %d\n\n", result35b)

	// Valid Sudoku examples
	fmt.Println("\n36. Valid Sudoku (Problem 0036)")
	fmt.Println("---------------------------------")

	// Example 1
	board36a := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	valid36a := leetcode.IsValidSudoku(board36a)
	fmt.Printf("Example 1: Valid sudoku board\n")
	fmt.Printf("  Is valid: %v\n\n", valid36a)

	// Count and Say examples
	fmt.Println("\n38. Count and Say (Problem 0038)")
	fmt.Println("---------------------------------")

	// Example 1
	n38a := 1
	result38a := leetcode.CountAndSay(n38a)
	fmt.Printf("Example 1: n = %d\n", n38a)
	fmt.Printf("  Sequence: %q\n\n", result38a)

	// Example 2
	n38b := 4
	result38b := leetcode.CountAndSay(n38b)
	fmt.Printf("Example 2: n = %d\n", n38b)
	fmt.Printf("  Sequence: %q\n\n", result38b)

	// Combination Sum examples
	fmt.Println("\n39. Combination Sum (Problem 0039)")
	fmt.Println("-----------------------------------")

	// Example 1
	candidates39a := []int{2, 3, 6, 7}
	target39a := 7
	combos39a := leetcode.CombinationSum(candidates39a, target39a)
	fmt.Printf("Example 1: candidates = %v, target = %d\n", candidates39a, target39a)
	fmt.Printf("  Result: %v (count: %d)\n\n", len(combos39a), len(combos39a))

	// Example 2
	candidates39b := []int{2, 3, 5}
	target39b := 8
	combos39b := leetcode.CombinationSum(candidates39b, target39b)
	fmt.Printf("Example 2: candidates = %v, target = %d\n", candidates39b, target39b)
	fmt.Printf("  Result: %v (count: %d)\n\n", len(combos39b), len(combos39b))

	// Combination Sum II examples
	fmt.Println("\n40. Combination Sum II (Problem 0040)")
	fmt.Println("--------------------------------------")

	// Example 1
	candidates40a := []int{10, 1, 2, 7, 6, 1, 5}
	target40a := 8
	combos40a := leetcode.CombinationSum2(candidates40a, target40a)
	fmt.Printf("Example 1: candidates = %v, target = %d\n", candidates40a, target40a)
	fmt.Printf("  Result: %v (count: %d)\n\n", len(combos40a), len(combos40a))

	// Example 2
	candidates40b := []int{2, 5, 2, 1, 2}
	target40b := 5
	combos40b := leetcode.CombinationSum2(candidates40b, target40b)
	fmt.Printf("Example 2: candidates = %v, target = %d\n", candidates40b, target40b)
	fmt.Printf("  Result: %v (count: %d)\n\n", len(combos40b), len(combos40b))

	// First Missing Positive examples
	fmt.Println("\n41. First Missing Positive (Problem 0041)")
	fmt.Println("------------------------------------------")

	// Example 1
	nums41a := []int{1, 2, 0}
	result41a := leetcode.FirstMissingPositive(nums41a)
	fmt.Printf("Example 1: nums = %v\n", nums41a)
	fmt.Printf("  First missing positive: %d\n\n", result41a)

	// Example 2
	nums41b := []int{3, 4, -1, 1}
	result41b := leetcode.FirstMissingPositive(nums41b)
	fmt.Printf("Example 2: nums = %v\n", nums41b)
	fmt.Printf("  First missing positive: %d\n\n", result41b)

	// Trapping Rain Water examples
	fmt.Println("\n42. Trapping Rain Water (Problem 0042)")
	fmt.Println("---------------------------------------")

	// Example 1
	height42a := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	water42a := leetcode.Trap(height42a)
	fmt.Printf("Example 1: height = %v\n", height42a)
	fmt.Printf("  Water trapped: %d\n\n", water42a)

	// Example 2
	height42b := []int{4, 2, 0, 3, 2, 5}
	water42b := leetcode.Trap(height42b)
	fmt.Printf("Example 2: height = %v\n", height42b)
	fmt.Printf("  Water trapped: %d\n\n", water42b)

	// Multiply Strings examples
	fmt.Println("\n43. Multiply Strings (Problem 0043)")
	fmt.Println("------------------------------------")

	// Example 1
	num143a := "123"
	num243a := "456"
	product43a := leetcode.Multiply(num143a, num243a)
	fmt.Printf("Example 1: num1 = %q, num2 = %q\n", num143a, num243a)
	fmt.Printf("  Product: %q\n\n", product43a)

	// Example 2
	num143b := "0"
	num243b := "123"
	product43b := leetcode.Multiply(num143b, num243b)
	fmt.Printf("Example 2: num1 = %q, num2 = %q\n", num143b, num243b)
	fmt.Printf("  Product: %q\n\n", product43b)

	// Wildcard Matching examples
	fmt.Println("\n44. Wildcard Matching (Problem 0044)")
	fmt.Println("-------------------------------------")

	// Example 1
	s44a := "aa"
	p44a := "a"
	match44a := leetcode.IsMatchWildcard(s44a, p44a)
	fmt.Printf("Example 1: s = %q, p = %q\n", s44a, p44a)
	fmt.Printf("  Matches: %v\n\n", match44a)

	// Example 2
	s44b := "aa"
	p44b := "*"
	match44b := leetcode.IsMatchWildcard(s44b, p44b)
	fmt.Printf("Example 2: s = %q, p = %q\n", s44b, p44b)
	fmt.Printf("  Matches: %v\n\n", match44b)

	// Jump Game II examples
	fmt.Println("\n45. Jump Game II (Problem 0045)")
	fmt.Println("--------------------------------")

	// Example 1
	nums45a := []int{2, 3, 1, 1, 4}
	jumps45a := leetcode.Jump(nums45a)
	fmt.Printf("Example 1: nums = %v\n", nums45a)
	fmt.Printf("  Minimum jumps: %d\n\n", jumps45a)

	// Example 2
	nums45b := []int{2, 3, 0, 1, 4}
	jumps45b := leetcode.Jump(nums45b)
	fmt.Printf("Example 2: nums = %v\n", nums45b)
	fmt.Printf("  Minimum jumps: %d\n\n", jumps45b)

	// Permutations examples
	fmt.Println("\n46. Permutations (Problem 0046)")
	fmt.Println("--------------------------------")

	// Example 1
	nums46a := []int{1, 2, 3}
	perms46a := leetcode.Permute(nums46a)
	fmt.Printf("Example 1: nums = %v\n", nums46a)
	fmt.Printf("  Permutations count: %d\n\n", len(perms46a))

	// Example 2
	nums46b := []int{0, 1}
	perms46b := leetcode.Permute(nums46b)
	fmt.Printf("Example 2: nums = %v\n", nums46b)
	fmt.Printf("  Permutations count: %d\n\n", len(perms46b))

	// Permutations II examples
	fmt.Println("\n47. Permutations II (Problem 0047)")
	fmt.Println("-----------------------------------")

	// Example 1
	nums47a := []int{1, 1, 2}
	perms47a := leetcode.PermuteUnique(nums47a)
	fmt.Printf("Example 1: nums = %v\n", nums47a)
	fmt.Printf("  Unique permutations count: %d\n\n", len(perms47a))

	// Example 2
	nums47b := []int{1}
	perms47b := leetcode.PermuteUnique(nums47b)
	fmt.Printf("Example 2: nums = %v\n", nums47b)
	fmt.Printf("  Unique permutations count: %d\n\n", len(perms47b))

	// Rotate Image examples
	fmt.Println("\n48. Rotate Image (Problem 0048)")
	fmt.Println("--------------------------------")

	// Example 1
	matrix48a := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	leetcode.Rotate(matrix48a)
	fmt.Printf("Example 1: Rotated 3x3 matrix\n")
	fmt.Printf("  Result: [%d, %d, %d], [%d, %d, %d], [%d, %d, %d]\n\n",
		matrix48a[0][0], matrix48a[0][1], matrix48a[0][2],
		matrix48a[1][0], matrix48a[1][1], matrix48a[1][2],
		matrix48a[2][0], matrix48a[2][1], matrix48a[2][2])

	// Group Anagrams examples
	fmt.Println("\n49. Group Anagrams (Problem 0049)")
	fmt.Println("---------------------------------")

	// Example 1
	strs49a := []string{"eat", "tea", "ate", "eat", "tan", "nat", "bat"}
	groups49a := leetcode.GroupAnagrams(strs49a)
	fmt.Printf("Example 1: strs = %v\n", strs49a)
	fmt.Printf("  Groups count: %d\n\n", len(groups49a))

	// Example 2
	strs49b := []string{""}
	groups49b := leetcode.GroupAnagrams(strs49b)
	fmt.Printf("Example 2: strs = %v\n", strs49b)
	fmt.Printf("  Groups count: %d\n\n", len(groups49b))

	// Pow(x, n) examples
	fmt.Println("\n50. Pow(x, n) (Problem 0050)")
	fmt.Println("------------------------------")

	// Example 1
	x50a := 2.0
	n50a := 10
	result50a := leetcode.MyPow(x50a, n50a)
	fmt.Printf("Example 1: x = %f, n = %d\n", x50a, n50a)
	fmt.Printf("  Result: %.6f\n\n", result50a)

	// Example 2
	x50b := 2.1
	n50b := 3
	result50b := leetcode.MyPow(x50b, n50b)
	fmt.Printf("Example 2: x = %f, n = %d\n", x50b, n50b)
	fmt.Printf("  Result: %.6f\n\n", result50b)

	fmt.Println("=== Demo Complete ===")
	fmt.Println("\nAll LeetCode problems 0001-0050 implemented successfully!")
}
