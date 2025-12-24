// 0038 - Count and Say
// https://leetcode.com/problems/count-and-say/
// Medium - String

package leetcode

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