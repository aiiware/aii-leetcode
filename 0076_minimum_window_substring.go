package leetcode

import "math"

// MinWindow solves LeetCode problem 0076: Minimum Window Substring
func MinWindow(s string, t string) string {
	if len(t) > len(s) || len(t) == 0 {
		return ""
	}

	tFreq := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tFreq[t[i]]++
	}

	windowFreq := make(map[byte]int)
	required := len(tFreq)
	formed := 0
	left := 0
	minLength := math.MaxInt32
	minStart := 0

	for right := 0; right < len(s); right++ {
		char := s[right]
		windowFreq[char]++

		if _, ok := tFreq[char]; ok && windowFreq[char] == tFreq[char] {
			formed++
		}

		for left <= right && formed == required {
			currentLength := right - left + 1
			if currentLength < minLength {
				minLength = currentLength
				minStart = left
			}

			leftChar := s[left]
			windowFreq[leftChar]--
			if _, ok := tFreq[leftChar]; ok && windowFreq[leftChar] < tFreq[leftChar] {
				formed--
			}
			left++
		}
	}

	if minLength == math.MaxInt32 {
		return ""
	}

	return s[minStart : minStart+minLength]
}


// MinWindowOptimized calls the main MinWindow function.
func MinWindowOptimized(s string, t string) string {
	return MinWindow(s, t)
}

// MinWindowSimplified calls the main MinWindow function.
func MinWindowSimplified(s string, t string) string {
	return MinWindow(s, t)
}