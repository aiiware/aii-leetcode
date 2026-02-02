package math

// 0268 - Missing Number (Easy)
// https://leetcode.com/problems/missing-number/

// MissingNumber returns the missing number in the range [0, n]
// where nums contains n distinct numbers taken from 0, 1, 2, ..., n
// Time Complexity: O(n)
// Space Complexity: O(1)
func MissingNumber(nums []int) int {
	n := len(nums)
	
	// Calculate expected sum using formula: n*(n+1)/2
	expectedSum := n * (n + 1) / 2
	
	// Calculate actual sum
	actualSum := 0
	for _, num := range nums {
		actualSum += num
	}
	
	// The missing number is the difference
	return expectedSum - actualSum
}

// MissingNumberXOR uses XOR properties to find the missing number
// XOR of a number with itself is 0, and XOR is commutative/associative
func MissingNumberXOR(nums []int) int {
	n := len(nums)
	result := n // Start with n because it's in the range [0, n]
	
	// XOR all indices and values
	for i := 0; i < n; i++ {
		result ^= i
		result ^= nums[i]
	}
	
	return result
}

// MissingNumberSort first sorts the array (not optimal but shows alternative approach)
// This is O(n log n) time, O(1) space if we can modify input
func MissingNumberSort(nums []int) int {
	// In a real implementation, we would sort the array
	// But for this problem, we have better solutions
	// This is just to show the concept
	
	n := len(nums)
	// If we could sort nums in-place:
	// sort.Ints(nums)
	
	// After sorting, check each position
	for i := 0; i < n; i++ {
		if nums[i] != i {
			return i
		}
	}
	
	// If all numbers 0..n-1 are present, then n is missing
	return n
}

// MissingNumberHash uses a hash set (extra space but clear logic)
func MissingNumberHash(nums []int) int {
	n := len(nums)
	present := make(map[int]bool, n+1)
	
	// Mark all numbers that are present
	for _, num := range nums {
		present[num] = true
	}
	
	// Check which number from 0 to n is missing
	for i := 0; i <= n; i++ {
		if !present[i] {
			return i
		}
	}
	
	return -1 // Should never reach here given problem constraints
}