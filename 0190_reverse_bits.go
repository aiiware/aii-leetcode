package leetcode

/*
190. Reverse Bits

Reverse bits of a given 32 bits unsigned integer.

Note:
- Note that in some languages, such as Java, there is no unsigned integer type. In this case, both input and output will be given as signed integer types and should not affect your implementation, as the internal binary representation of the integer is the same whether it is signed or unsigned.
- In Java, the compiler represents the signed integers using 2's complement notation. Therefore, in Example 2 above, the input represents the signed integer -3 and the output represents the signed integer -1073741825.

Example 1:
Input: n = 00000010100101000001111010011100
Output:    00111001011110000010100101000000
Explanation: The input binary string 00000010100101000001111010011100 represents the unsigned integer 43261596, so return 964176192 which its binary representation is 00111001011110000010100101000000.

Example 2:
Input: n = 11111111111111111111111111111101
Output:    10111111111111111111111111111111
Explanation: The input binary string 11111111111111111111111111111101 represents the unsigned integer 4294967293, so return 3221225471 which its binary representation is 10111111111111111111111111111111.

Constraints:
- The input must be a binary string of length 32

Follow up: If this function is called many times, how would you optimize it?

Difficulty: Easy
Tags: Bit Manipulation, Divide and Conquer
Companies: Apple, Meta, Nvidia, Google, Qualcomm, Amazon, Airbnb
*/

// reverseBits1: Basic bit manipulation approach
// Time: O(32) = O(1), Space: O(1)
func reverseBits1(num uint32) uint32 {
	var result uint32 = 0
	for i := 0; i < 32; i++ {
		// Shift result left to make room for next bit
		result <<= 1
		// Add the least significant bit of num to result
		result |= num & 1
		// Shift num right to process next bit
		num >>= 1
	}
	return result
}

// reverseBits2: Optimized using bit operations without loop
// Time: O(1), Space: O(1)
func reverseBits2(num uint32) uint32 {
	// Swap odd and even bits
	num = ((num & 0xaaaaaaaa) >> 1) | ((num & 0x55555555) << 1)
	// Swap consecutive pairs
	num = ((num & 0xcccccccc) >> 2) | ((num & 0x33333333) << 2)
	// Swap nibbles
	num = ((num & 0xf0f0f0f0) >> 4) | ((num & 0x0f0f0f0f) << 4)
	// Swap bytes
	num = ((num & 0xff00ff00) >> 8) | ((num & 0x00ff00ff) << 8)
	// Swap 16-bit halves
	num = (num >> 16) | (num << 16)
	return num
}

// reverseBits3: Using lookup table for optimization (follow-up answer)
// This approach is optimal when the function is called many times
// Time: O(1), Space: O(256) for lookup table
func reverseBits3(num uint32) uint32 {
	// Precompute reverse of all 8-bit numbers (0-255)
	// This table can be computed once and reused
	var revByte [256]uint32
	for i := 0; i < 256; i++ {
		revByte[i] = reverseByte(uint32(i))
	}

	// Reverse each byte and combine
	return (revByte[num&0xff] << 24) |
		(revByte[(num>>8)&0xff] << 16) |
		(revByte[(num>>16)&0xff] << 8) |
		revByte[num>>24]
}

// reverseByte: Helper function to reverse an 8-bit number
func reverseByte(b uint32) uint32 {
	// Reverse bits in a byte using divide and conquer
	b = ((b & 0xf0) >> 4) | ((b & 0x0f) << 4)
	b = ((b & 0xcc) >> 2) | ((b & 0x33) << 2)
	b = ((b & 0xaa) >> 1) | ((b & 0x55) << 1)
	return b
}

// reverseBits: Main function that uses the basic approach
// This is the function that matches LeetCode's expected signature
func reverseBits(num uint32) uint32 {
	return reverseBits1(num)
}