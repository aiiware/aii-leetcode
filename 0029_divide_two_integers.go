package leetcode

import "math"

// Divide solves LeetCode problem 0029: Divide Two Integers
// Difficulty: Medium
// Tags: Math, Bit Manipulation
//
// Given two integers dividend and divisor, divide two integers without using
// multiplication, division, and mod operator. The integer division should
// truncate toward zero, which means losing its fractional part.
//
// Time complexity: O(log n), Space complexity: O(1)
func Divide(dividend int, divisor int) int {
	// Handle edge cases
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend == math.MinInt32 {
			return math.MaxInt32 // Overflow case
		}
		return -dividend
	}

	// Determine sign of result
	negative := (dividend < 0) != (divisor < 0)

	// Work with positive numbers
	a := absInt(dividend)
	b := absInt(divisor)

	// Handle the case where divisor is greater than dividend
	if b > a {
		return 0
	}

	// Use bit manipulation for division
	quotient := 0

	// Find the largest multiple of divisor that is <= dividend
	for a >= b {
		shift := 0
		// Double the divisor until it's larger than the remaining dividend
		for a >= (b << shift) {
			shift++
		}
		// We went one shift too far, so back up
		shift--
		// Add 2^shift to quotient
		quotient += 1 << shift
		// Subtract b * 2^shift from a
		a -= b << shift
	}

	// Apply sign
	if negative {
		quotient = -quotient
	}

	// Handle overflow
	if quotient > math.MaxInt32 {
		return math.MaxInt32
	}
	if quotient < math.MinInt32 {
		return math.MinInt32
	}

	return quotient
}

// DivideSimple is a simpler implementation using subtraction
func DivideSimple(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend == math.MinInt32 {
			return math.MaxInt32
		}
		return -dividend
	}

	negative := (dividend < 0) != (divisor < 0)
	a := absInt(dividend)
	b := absInt(divisor)

	quotient := 0
	for a >= b {
		a -= b
		quotient++
	}

	if negative {
		quotient = -quotient
	}

	if quotient > math.MaxInt32 {
		return math.MaxInt32
	}
	if quotient < math.MinInt32 {
		return math.MinInt32
	}

	return quotient
}

// absInt returns the absolute value of an integer
func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}