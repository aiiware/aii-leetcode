// 0031 - Next Permutation
// https://leetcode.com/problems/next-permutation/
// Medium - Array, Two Pointers

package leetcode

// NextPermutation rearranges numbers into the lexicographically next greater permutation.
// Time Complexity: O(n) where n is length of nums
// Space Complexity: O(1) in-place modification
func NextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	// Step 1: Find the first decreasing element from the right
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	// Step 2: If we found a decreasing element
	if i >= 0 {
		// Find the smallest element to the right that is greater than nums[i]
		j := n - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		// Swap nums[i] and nums[j]
		nums[i], nums[j] = nums[j], nums[i]
	}

	// Step 3: Reverse the suffix (elements after i)
	reverse(nums, i+1, n-1)
}

// reverse reverses elements in nums from start to end (inclusive)
func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// NextPermutationBruteForce is a brute force approach for comparison
func NextPermutationBruteForce(nums []int) {
	// This is not efficient but shows the concept
	// In practice, use the optimized algorithm above
	n := len(nums)
	
	// Find the next permutation by checking all permutations
	// This is O(n!) and not practical for large n
	// Implemented here only for educational purposes
	
	// Find the first position where we can make a change
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	
	if i >= 0 {
		// Find the next larger element to swap with
		j := n - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	
	// Reverse the suffix
	left, right := i+1, n-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}