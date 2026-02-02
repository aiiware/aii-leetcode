package math

// 0201 - Bitwise AND of Numbers Range (Medium)
// https://leetcode.com/problems/bitwise-and-of-numbers-range/

// RangeBitwiseAnd returns the bitwise AND of all numbers in the range [left, right]
// Time Complexity: O(1) - we just shift bits until left == right
// Space Complexity: O(1)
func RangeBitwiseAnd(left int, right int) int {
	// The key insight is that the AND of a range will have bits set only where
	// all numbers in the range have that bit set. As we move through the range,
	// the lower bits will change, so we need to find the common prefix of left and right.
	
	shift := 0
	// Keep shifting right until left and right are equal
	for left < right {
		left >>= 1
		right >>= 1
		shift++
	}
	// Shift back left to restore the common prefix bits
	return left << shift
}

// Alternative approach using bit manipulation
func RangeBitwiseAnd2(left int, right int) int {
	// Keep removing the last bit from right until right <= left
	for left < right {
		// Remove the last set bit from right
		right = right & (right - 1)
	}
	return right & left
}