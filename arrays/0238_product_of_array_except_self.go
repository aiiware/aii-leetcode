package arrays

// ProductExceptSelf solves LeetCode problem 0238: Product of Array Except Self
// Difficulty: Medium
// Tags: Array, Prefix Sum
//
// Given an integer array nums, return an array answer such that answer[i] is
// equal to the product of all the elements of nums except nums[i].
//
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit
// integer.
//
// You must write an algorithm that runs in O(n) time and without using the
// division operation.
//
// Time complexity: O(n), Space complexity: O(1) (excluding output array)
func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	// Initialize result array
	result := make([]int, n)

	// First pass: calculate prefix products
	// result[i] will contain product of all elements to the left of i
	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	// Second pass: calculate suffix products and multiply with prefix
	// Use a variable to keep track of product of elements to the right
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		result[i] = result[i] * suffix
		suffix *= nums[i]
	}

	return result
}