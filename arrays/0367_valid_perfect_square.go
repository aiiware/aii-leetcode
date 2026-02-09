package arrays

// IsPerfectSquare solves LeetCode problem 367: Valid Perfect Square
// Difficulty: Easy
// Tags: Math, Binary Search
//
// Given a positive integer num, return true if num is a perfect square,
// otherwise return false.
//
// A perfect square is an integer that is the square of an integer.
// In other words, it is the product of some integer with itself.
//
// You must not use any built-in library function, such as sqrt.
//
// Time complexity: O(log num), Space complexity: O(1)
func IsPerfectSquare(num int) bool {
	// Edge cases
	if num < 0 {
		return false
	}
	if num == 0 || num == 1 {
		return true // 0 = 0², 1 = 1²
	}

	// Binary search for the square root
	left, right := 1, num

	for left <= right {
		// Calculate mid point (prevents overflow)
		mid := left + (right-left)/2

		// Check if mid² == num
		// Use 64-bit to prevent overflow when squaring
		square := int64(mid) * int64(mid)

		if square == int64(num) {
			return true
		} else if square < int64(num) {
			// Too small, search right half
			left = mid + 1
		} else {
			// Too large, search left half
			right = mid - 1
		}
	}

	return false
}

// IsPerfectSquareOptimized solves the same problem with optimization for large numbers
// This version stops early when mid² exceeds num and uses integer division to avoid overflow
func IsPerfectSquareOptimized(num int) bool {
	// Edge cases
	if num < 0 {
		return false
	}
	if num == 0 || num == 1 {
		return true
	}

	// Binary search with early exit
	left, right := 1, num

	for left <= right {
		mid := left + (right-left)/2

		// Use division to check if mid is the square root
		// This avoids overflow from multiplication
		quotient := num / mid
		remainder := num % mid

		if quotient == mid && remainder == 0 {
			return true
		} else if quotient > mid {
			// mid is too small (mid² < num)
			left = mid + 1
		} else {
			// mid is too large (mid² > num) or mid is not exact divisor
			right = mid - 1
		}
	}

	return false
}