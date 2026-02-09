package arrays

/*
191. Number of 1 Bits

Write a function that takes the binary representation of an unsigned integer and returns the number of '1' bits it has (also known as the Hamming weight).

Note:
- Note that in some languages, such as Java, there is no unsigned integer type. In this case, the input will be given as a signed integer type. It should not affect your implementation, as the integer's internal binary representation is the same, whether it is signed or unsigned.
- In Java, the compiler represents the signed integers using 2's complement notation. Therefore, in Example 3, the input represents the signed integer -3.

Example 1:
Input: n = 00000000000000000000000000001011
Output: 3
Explanation: The input binary string 00000000000000000000000000001011 has a total of three '1' bits.

Example 2:
Input: n = 00000000000000000000000010000000
Output: 1
Explanation: The input binary string 00000000000000000000000010000000 has a total of one '1' bit.

Example 3:
Input: n = 11111111111111111111111111111101
Output: 31
Explanation: The input binary string 11111111111111111111111111111101 has a total of thirty one '1' bits.

Constraints:
- The input must be a binary string of length 32.

Follow up: If this function is called many times, how would you optimize it?

Difficulty: Easy
Tags: Bit Manipulation, Divide and Conquer
Companies: Apple, Microsoft, Amazon, Google, Facebook, Adobe, Bloomberg, Uber
*/

// hammingWeight1: Basic approach using bit shifting
// Time: O(32) = O(1), Space: O(1)
func hammingWeight1(num uint32) int {
	count := 0
	for num != 0 {
		// Check if the least significant bit is 1
		count += int(num & 1)
		// Right shift to process next bit
		num >>= 1
	}
	return count
}

// hammingWeight2: Brian Kernighan's algorithm
// This algorithm clears the least significant 1-bit in each iteration
// Time: O(k) where k is number of 1 bits, Space: O(1)
func hammingWeight2(num uint32) int {
	count := 0
	for num != 0 {
		// Clear the least significant 1-bit
		num &= num - 1
		count++
	}
	return count
}

// hammingWeight3: Using built-in bits.OnesCount32
// Time: O(1), Space: O(1)
// This uses hardware instructions when available
func hammingWeight3(num uint32) int {
	// Go's math/bits package provides optimized implementation
	// This would be: return bits.OnesCount32(num)
	// But we'll implement our own for educational purposes

	// Divide and conquer approach
	// Count bits in parallel
	num = (num & 0x55555555) + ((num >> 1) & 0x55555555)  // count bits in pairs
	num = (num & 0x33333333) + ((num >> 2) & 0x33333333)  // count bits in 4-bit groups
	num = (num & 0x0f0f0f0f) + ((num >> 4) & 0x0f0f0f0f)  // count bits in 8-bit groups
	num = (num & 0x00ff00ff) + ((num >> 8) & 0x00ff00ff)  // count bits in 16-bit groups
	num = (num & 0x0000ffff) + ((num >> 16) & 0x0000ffff) // count bits in 32-bit number
	return int(num)
}

// hammingWeight4: Lookup table approach (optimal for repeated calls)
// Time: O(1), Space: O(256) for lookup table
func hammingWeight4(num uint32) int {
	// Precompute Hamming weight for all 8-bit numbers (0-255)
	// This table can be computed once and reused
	var bitCount [256]int
	for i := 0; i < 256; i++ {
		bitCount[i] = bitCount[i>>1] + (i & 1)
	}

	// Count bits in each byte using lookup table
	return bitCount[num&0xff] +
		bitCount[(num>>8)&0xff] +
		bitCount[(num>>16)&0xff] +
		bitCount[num>>24]
}

// hammingWeight: Main function that uses Brian Kernighan's algorithm
// This is the function that matches LeetCode's expected signature
func hammingWeight(num uint32) int {
	return hammingWeight2(num)
}
