package leetcode

import (
	"strconv"
	"strings"
)

// Problem 0093: Restore IP Addresses
//
// A valid IP address consists of exactly four integers separated by single dots. 
// Each integer is between 0 and 255 (inclusive) and cannot have leading zeros.
//
// For example:
// - "0.1.2.201" and "192.168.1.1" are valid IP addresses.
// - "0.011.255.245", "192.168.1.312" and "192.168@1.1" are invalid IP addresses.
//
// Given a string s containing only digits, return all possible valid IP addresses 
// that can be formed by inserting dots into s. You are not allowed to reorder or 
// remove any digits in s. You may return the valid IP addresses in any order.
//
// Example 1:
// Input: s = "25525511135"
// Output: ["255.255.11.135","255.255.111.35"]
//
// Example 2:
// Input: s = "0000"
// Output: ["0.0.0.0"]
//
// Example 3:
// Input: s = "101023"
// Output: ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
//
// Constraints:
// - 1 <= s.length <= 20
// - s consists of digits only.

// restoreIpAddresses is the main solution function using backtracking.
// Time complexity: O(3^4) = O(1) since depth is fixed at 4, Space complexity: O(1)
func restoreIpAddresses(s string) []string {
	var result []string
	var current []string
	
	var backtrack func(int)
	backtrack = func(start int) {
		// If we have 4 segments and used all characters, add to result
		if len(current) == 4 {
			if start == len(s) {
				result = append(result, strings.Join(current, "."))
			}
			return
		}
		
		// Try segments of length 1, 2, and 3
		for length := 1; length <= 3; length++ {
			if start+length > len(s) {
				break
			}
			
			segment := s[start : start+length]
			
			// Check if segment is valid
			if isValidIPSegment(segment) {
				current = append(current, segment)
				backtrack(start + length)
				current = current[:len(current)-1] // backtrack
			}
		}
	}
	
	backtrack(0)
	return result
}

// isValidIPSegment checks if a string segment is valid for an IP address
func isValidIPSegment(segment string) bool {
	// Check length
	if len(segment) == 0 || len(segment) > 3 {
		return false
	}
	
	// Check for leading zero
	if len(segment) > 1 && segment[0] == '0' {
		return false
	}
	
	// Check numeric value
	val, err := strconv.Atoi(segment)
	if err != nil {
		return false
	}
	
	return val >= 0 && val <= 255
}

// restoreIpAddressesIterative uses iterative approach with nested loops.
func restoreIpAddressesIterative(s string) []string {
	var result []string
	n := len(s)
	
	// Try all possible positions for the three dots
	for i := 1; i <= 3 && i <= n-3; i++ {
		for j := i + 1; j <= i+3 && j <= n-2; j++ {
			for k := j + 1; k <= j+3 && k <= n-1; k++ {
				// Extract four segments
				seg1 := s[0:i]
				seg2 := s[i:j]
				seg3 := s[j:k]
				seg4 := s[k:]
				
				// Check if all segments are valid
				if isValidIPSegment(seg1) && isValidIPSegment(seg2) &&
					isValidIPSegment(seg3) && isValidIPSegment(seg4) {
					result = append(result, seg1+"."+seg2+"."+seg3+"."+seg4)
				}
			}
		}
	}
	
	return result
}

// restoreIpAddressesDFS uses DFS approach.
func restoreIpAddressesDFS(s string) []string {
	var result []string
	
	var dfs func(int, int, string)
	dfs = func(start, segmentNum, current string) {
		// If we have 4 segments and used all characters
		if segmentNum == 4 {
			if start == len(s) {
				result = append(result, current[1:]) // Remove leading dot
			}
			return
		}
		
		// Try segments of length 1, 2, and 3
		for length := 1; length <= 3; length++ {
			if start+length > len(s) {
				break
			}
			
			segment := s[start : start+length]
			if isValidIPSegment(segment) {
				dfs(start+length, segmentNum+1, current+"."+segment)
			}
		}
	}
	
	dfs(0, 0, "")
	return result
}

// restoreIpAddressesDP uses dynamic programming (memoization).
func restoreIpAddressesDP(s string) []string {
	n := len(s)
	if n < 4 || n > 12 {
		return []string{}
	}
	
	memo := make(map[[2]int][]string)
	return dp(s, 0, 0, memo)
}

func dp(s string, start, segments int, memo map[[2]int][]string) []string {
	key := [2]int{start, segments}
	if val, exists := memo[key]; exists {
		return val
	}
	
	var result []string
	
	// Base case: if we need 4 segments and we're at the end
	if segments == 4 {
		if start == len(s) {
			result = append(result, "")
		}
		memo[key] = result
		return result
	}
	
	// Try segments of length 1, 2, and 3
	for length := 1; length <= 3; length++ {
		if start+length > len(s) {
			break
		}
		
		segment := s[start : start+length]
		if isValidIPSegment(segment) {
			subResults := dp(s, start+length, segments+1, memo)
			for _, sub := range subResults {
				if sub == "" {
					result = append(result, segment)
				} else {
					result = append(result, segment+"."+sub)
				}
			}
		}
	}
	
	memo[key] = result
	return result
}

// restoreIpAddressesBFS uses BFS approach.
func restoreIpAddressesBFS(s string) []string {
	var result []string
	
	// Queue elements: (start index, segments built, current string)
	queue := [][3]interface{}{{0, 0, ""}}
	
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		
		start := front[0].(int)
		segments := front[1].(int)
		current := front[2].(string)
		
		// If we have 4 segments and used all characters
		if segments == 4 {
			if start == len(s) {
				result = append(result, current[1:]) // Remove leading dot
			}
			continue
		}
		
		// Try segments of length 1, 2, and 3
		for length := 1; length <= 3; length++ {
			if start+length > len(s) {
				break
			}
			
			segment := s[start : start+length]
			if isValidIPSegment(segment) {
				queue = append(queue, [3]interface{}{
					start + length,
					segments + 1,
					current + "." + segment,
				})
			}
		}
	}
	
	return result
}

// restoreIpAddressesOptimized is an optimized version.
func restoreIpAddressesOptimized(s string) []string {
	n := len(s)
	if n < 4 || n > 12 {
		return []string{}
	}
	
	var result []string
	var path []string
	
	var backtrack func(int, int)
	backtrack = func(start, segments int) {
		// Prune: if remaining characters too many or too few for remaining segments
		remaining := n - start
		remainingSegments := 4 - segments
		if remaining < remainingSegments || remaining > remainingSegments*3 {
			return
		}
		
		if segments == 4 && start == n {
			result = append(result, strings.Join(path, "."))
			return
		}
		
		// Try segments of length 1, 2, and 3
		for length := 1; length <= 3 && start+length <= n; length++ {
			segment := s[start : start+length]
			if isValidIPSegment(segment) {
				path = append(path, segment)
				backtrack(start+length, segments+1)
				path = path[:len(path)-1]
			}
		}
	}
	
	backtrack(0, 0)
	return result
}

// RestoreIpAddresses is the public interface function.
// It uses the optimized backtracking solution by default.
func RestoreIpAddresses(s string) []string {
	return restoreIpAddressesOptimized(s)
}