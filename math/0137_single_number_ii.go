package math

/*
# 0137 - Single Number II
## Problem Description
Given an integer array nums where every element appears three times except for one, which appears exactly once. Find the single element and return it.

You must implement a solution with a linear runtime complexity and use only constant extra space.

## Examples
Example 1:
Input: nums = [2,2,3,2]
Output: 3

Example 2:
Input: nums = [0,1,0,1,0,1,99]
Output: 99

## Constraints
- 1 <= nums.length <= 3 * 10^4
- -2^31 <= nums[i] <= 2^31 - 1
- Each element in nums appears exactly three times except for one element which appears once.

## Solution Approach
This problem can be solved using bit manipulation:
- Count the number of 1s at each bit position across all numbers
- For each bit position, if the count % 3 == 1, that bit is set in the single number
- This works because numbers appearing 3 times will contribute 3 * bit value, which mod 3 is 0

Alternative approach using bitwise operations:
- Use two variables to represent bits that have appeared once, twice, or three times
- This approach uses constant space and linear time

Time Complexity: O(N) where N is the length of nums
Space Complexity: O(1)
*/

// SingleNumberII returns the element that appears only once (others appear 3 times)
func SingleNumberII(nums []int) int {
	// Method 1: Count bits approach (32-bit integers)
	var result int32 = 0
	
	// Check each bit position (0 to 31)
	for i := 0; i < 32; i++ {
		count := 0
		bitMask := int32(1 << i)
		
		// Count how many numbers have this bit set
		for _, num := range nums {
			if int32(num)&bitMask != 0 {
				count++
			}
		}
		
		// If count % 3 == 1, this bit is set in the single number
		if count%3 == 1 {
			result |= bitMask
		}
	}
	
	return int(result)
}

// SingleNumberIIOptimized is an optimized version using bitwise operations
func SingleNumberIIOptimized(nums []int) int {
	// Method 2: Using two variables to track bits
	// ones: bits that have appeared once
	// twos: bits that have appeared twice
	// When a bit appears three times, it gets cleared from both ones and twos
	ones, twos := 0, 0
	
	for _, num := range nums {
		// Update ones:
		// If num has a bit set that's already in ones, it means this is the second time
		// we're seeing it, so we should remove it from ones (using XOR with num).
		// But we also need to consider bits that are in twos (appeared twice already).
		// The formula: ones = (ones ^ num) & ^twos
		// This means: XOR with num to toggle the bit, then clear bits that are in twos
		ones = (ones ^ num) & ^twos
		
		// Update twos:
		// If num has a bit set that's already in ones (after update), it means this is
		// the second time we're seeing it, so we should add it to twos.
		// The formula: twos = (twos ^ num) & ^ones
		// This means: XOR with num to toggle the bit, then clear bits that are in ones
		twos = (twos ^ num) & ^ones
	}
	
	// At the end, ones contains the single number
	// (bits that appeared exactly once, not three times)
	return ones
}