package leetcode

/*
165. Compare Version Numbers

Given two version strings, version1 and version2, compare them.

A version string consists of revisions separated by dots '.'. The value of the revision is its integer conversion ignoring leading zeros.

To compare version strings, compare their revision values in left-to-right order. If one of the version strings has fewer revisions, treat the missing revision values as 0.

Return the following:
- If version1 < version2, return -1.
- If version1 > version2, return 1.
- Otherwise, return 0.

Example 1:
Input: version1 = "1.2", version2 = "1.10"
Output: -1
Explanation: version1's second revision is "2" and version2's second revision is "10": 2 < 10, so version1 < version2.

Example 2:
Input: version1 = "1.01", version2 = "1.001"
Output: 0
Explanation: Ignoring leading zeros, both "01" and "001" represent the same integer "1".

Example 3:
Input: version1 = "1.0", version2 = "1.0.0.0"
Output: 0
Explanation: version1 has 1 revision, while version2 has 4. The missing revisions in version1 are treated as "0".

Constraints:
- 1 <= version1.length, version2.length <= 500
- version1 and version2 only contain digits and '.'.
- version1 and version2 are valid version numbers.
- All the given revision numbers are in the range [0, 2^31 - 1].

Difficulty: Medium
Tags: Two Pointers, String
Companies: Amazon, Apple, Arista Networks, Google, Microsoft, Square
*/

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	// Split versions into revision arrays
	revs1 := strings.Split(version1, ".")
	revs2 := strings.Split(version2, ".")
	
	// Get the maximum length
	maxLen := len(revs1)
	if len(revs2) > maxLen {
		maxLen = len(revs2)
	}
	
	// Compare revisions one by one
	for i := 0; i < maxLen; i++ {
		// Get revision values, default to 0 if revision doesn't exist
		rev1 := 0
		if i < len(revs1) {
			rev1, _ = strconv.Atoi(revs1[i])
		}
		
		rev2 := 0
		if i < len(revs2) {
			rev2, _ = strconv.Atoi(revs2[i])
		}
		
		// Compare revisions
		if rev1 < rev2 {
			return -1
		} else if rev1 > rev2 {
			return 1
		}
		// If equal, continue to next revision
	}
	
	// All revisions are equal
	return 0
}