package leetcode

// Problem 0089: Gray Code
//
// An n-bit gray code sequence is a sequence of 2^n integers where:
// 1. Every integer is in the inclusive range [0, 2^n - 1],
// 2. The first integer is 0,
// 3. An integer appears no more than once in the sequence,
// 4. The binary representation of every pair of adjacent integers differs by exactly one bit, and
// 5. The binary representation of the first and last integers differs by exactly one bit.
//
// Given an integer n, return any valid n-bit gray code sequence.
//
// Example 1:
// Input: n = 2
// Output: [0,1,3,2]
// Explanation:
// The binary representation of [0,1,3,2] is [00,01,11,10].
// - 00 and 01 differ by one bit
// - 01 and 11 differ by one bit
// - 11 and 10 differ by one bit
// - 10 and 00 differ by one bit
// [0,2,3,1] is also a valid gray code sequence, whose binary representation is [00,10,11,01].
//
// Example 2:
// Input: n = 1
// Output: [0,1]
// Explanation:
// The binary representation of [0,1] is [0,1].
// - 0 and 1 differ by one bit
//
// Constraints:
// - 1 <= n <= 16

// grayCode is the main solution function using the reflect-and-prefix method.
// Time complexity: O(2^n), Space complexity: O(2^n)
func grayCode(n int) []int {
	if n <= 0 {
		return []int{0}
	}

	// Start with n = 1 case
	result := []int{0, 1}

	// Build up for larger n
	for i := 2; i <= n; i++ {
		// Reflect and add prefix (1 << (i-1))
		length := len(result)
		for j := length - 1; j >= 0; j-- {
			result = append(result, result[j]|(1<<(i-1)))
		}
	}

	return result
}

// grayCodeFormula uses the direct formula: G(i) = i ^ (i >> 1)
// This generates the standard binary reflected Gray code.
func grayCodeFormula(n int) []int {
	size := 1 << n // 2^n
	result := make([]int, size)

	for i := 0; i < size; i++ {
		result[i] = i ^ (i >> 1)
	}

	return result
}

// grayCodeRecursive uses recursion to generate Gray code.
func grayCodeRecursive(n int) []int {
	if n == 0 {
		return []int{0}
	}
	if n == 1 {
		return []int{0, 1}
	}

	// Get Gray code for n-1
	prev := grayCodeRecursive(n - 1)
	result := make([]int, len(prev)*2)

	// Copy the previous sequence
	copy(result, prev)

	// Add reflected sequence with MSB set
	msb := 1 << (n - 1)
	for i := 0; i < len(prev); i++ {
		result[len(prev)+i] = prev[len(prev)-1-i] | msb
	}

	return result
}

// grayCodeIterative uses iterative approach with bit manipulation.
func grayCodeIterative(n int) []int {
	size := 1 << n
	result := make([]int, size)

	// Generate using the formula iteratively
	for i := 0; i < size; i++ {
		result[i] = i ^ (i >> 1)
	}

	return result
}

// grayCodeBacktracking uses backtracking to find a valid Gray code sequence.
func grayCodeBacktracking(n int) []int {
	size := 1 << n
	result := make([]int, size)
	visited := make([]bool, size)

	// Start with 0
	result[0] = 0
	visited[0] = true

	if backtrackGrayCode(n, result, visited, 1) {
		return result
	}

	// Should always find a solution for valid n
	return []int{0}
}

func backtrackGrayCode(n int, result []int, visited []bool, pos int) bool {
	if pos == 1<<n {
		// Check if last and first differ by one bit
		last := result[pos-1]
		first := result[0]
		return countBits(last^first) == 1
	}

	last := result[pos-1]
	// Try all possible next numbers
	for i := 0; i < (1 << n); i++ {
		if !visited[i] && countBits(last^i) == 1 {
			result[pos] = i
			visited[i] = true

			if backtrackGrayCode(n, result, visited, pos+1) {
				return true
			}

			// Backtrack
			visited[i] = false
		}
	}

	return false
}

// countBits counts the number of 1 bits in an integer
func countBits(x int) int {
	count := 0
	for x > 0 {
		count += x & 1
		x >>= 1
	}
	return count
}

// grayCodeIterative2 uses another iterative method.
func grayCodeIterative2(n int) []int {
	if n == 0 {
		return []int{0}
	}

	result := []int{0, 1}
	for i := 1; i < n; i++ {
		msb := 1 << i
		// Append reversed sequence with MSB set
		for j := len(result) - 1; j >= 0; j-- {
			result = append(result, result[j]|msb)
		}
	}

	return result
}

// GrayCode is the public interface function.
// It uses the formula method by default as it's the most efficient.
func GrayCode(n int) []int {
	return grayCodeFormula(n)
}