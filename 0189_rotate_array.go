package leetcode

/*
189. Rotate Array

Companies: Amazon, Microsoft, Meta, Apple, Bloomberg, Adobe, Google, Cisco, Facebook, Goldman Sachs, Oracle, Paypal, Snapchat, Uber

Categories: Array, Math, Two Pointers

Difficulty: Medium

Description:
Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

Example 1:
Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]

Example 2:
Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]

Constraints:
- 1 <= nums.length <= 10^5
- -2^31 <= nums[i] <= 2^31 - 1
- 0 <= k <= 10^5

Follow up:
- Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
- Could you do it in-place with O(1) extra space?
*/

// Solution 1: Reverse Method (Optimal in-place solution)
// Time: O(n), Space: O(1)
// This is the most elegant solution using three reversals
func rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 || k == 0 {
		return
	}
	
	// Normalize k to avoid unnecessary rotations
	k = k % n
	
	// If k is 0 after normalization, no rotation needed
	if k == 0 {
		return
	}
	
	// Helper function to reverse a portion of the array
	reverse := func(start, end int) {
		for start < end {
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end--
		}
	}
	
	// Step 1: Reverse the entire array
	reverse(0, n-1)
	// Step 2: Reverse the first k elements
	reverse(0, k-1)
	// Step 3: Reverse the remaining n-k elements
	reverse(k, n-1)
}

// Solution 2: Using Extra Array
// Time: O(n), Space: O(n)
// Simple and intuitive but uses extra space
func rotateExtraArray(nums []int, k int) {
	n := len(nums)
	if n == 0 || k == 0 {
		return
	}
	
	k = k % n
	if k == 0 {
		return
	}
	
	// Create a copy of the array
	temp := make([]int, n)
	
	// Copy elements to their new positions
	for i := 0; i < n; i++ {
		temp[(i+k)%n] = nums[i]
	}
	
	// Copy back to original array
	copy(nums, temp)
}

// Solution 3: Cyclic Replacements
// Time: O(n), Space: O(1)
// Moves elements in cycles without extra array
func rotateCyclic(nums []int, k int) {
	n := len(nums)
	if n == 0 || k == 0 {
		return
	}
	
	k = k % n
	if k == 0 {
		return
	}
	
	count := 0 // Number of elements moved
	for start := 0; count < n; start++ {
		current := start
		prev := nums[start]
		
		for {
			next := (current + k) % n
			temp := nums[next]
			nums[next] = prev
			prev = temp
			current = next
			count++
			
			if start == current {
				break
			}
		}
	}
}

// Solution 4: Using Built-in Copy and Append
// Time: O(n), Space: O(n)
// Go-specific solution using slice operations
func rotateBuiltin(nums []int, k int) {
	n := len(nums)
	if n == 0 || k == 0 {
		return
	}
	
	k = k % n
	if k == 0 {
		return
	}
	
	// Get the last k elements
	lastK := make([]int, k)
	copy(lastK, nums[n-k:])
	
	// Shift the first n-k elements to the right
	copy(nums[k:], nums[:n-k])
	
	// Copy the last k elements to the beginning
	copy(nums[:k], lastK)
}

// Solution 5: Brute Force (Naive approach)
// Time: O(n*k), Space: O(1)
// Not efficient for large arrays but demonstrates the basic idea
func rotateBruteForce(nums []int, k int) {
	n := len(nums)
	if n == 0 || k == 0 {
		return
	}
	
	k = k % n
	if k == 0 {
		return
	}
	
	for i := 0; i < k; i++ {
		// Store the last element
		last := nums[n-1]
		
		// Shift all elements to the right by 1
		for j := n - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		
		// Put the last element at the beginning
		nums[0] = last
	}
}

// Solution 6: Using Two Pointers with Block Swap
// Time: O(n), Space: O(1)
// Another in-place solution using block swapping
func rotateBlockSwap(nums []int, k int) {
	n := len(nums)
	if n == 0 || k == 0 {
		return
	}
	
	k = k % n
	if k == 0 {
		return
	}

	// Convert right rotation to left rotation distance.
	d := n - k
	
	// Helper function to swap two blocks
	swap := func(start1, start2, length int) {
		for i := 0; i < length; i++ {
			nums[start1+i], nums[start2+i] = nums[start2+i], nums[start1+i]
		}
	}
	
	// If k is exactly half the array length
	if d == n-d {
		swap(0, n-d, d)
		return
	}
	
	// Handle general case
	i, j := d, n-d
	for i != j {
		if i < j {
			swap(d-i, d+j-i, i)
			j -= i
		} else {
			swap(d-i, d, j)
			i -= j
		}
	}
	
	swap(d-i, d, i)
}

// Solution 7: Using GCD for Cycle Detection
// Time: O(n), Space: O(1)
// Mathematical approach using greatest common divisor
func rotateGCD(nums []int, k int) {
	n := len(nums)
	if n == 0 || k == 0 {
		return
	}
	
	k = k % n
	if k == 0 {
		return
	}
	
	// Function to calculate GCD
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	
	cycles := gcd(n, k)
	
	for i := 0; i < cycles; i++ {
		current := i
		prev := nums[i]
		
		for {
			next := (current + k) % n
			temp := nums[next]
			nums[next] = prev
			prev = temp
			current = next
			
			if current == i {
				break
			}
		}
	}
}

// Solution 8: Using Recursive Approach
// Time: O(n), Space: O(n) for recursion stack during reversal
// Demonstrates a recursive version of the reverse method
func rotateRecursive(nums []int, k int) {
	n := len(nums)
	if n == 0 || k == 0 {
		return
	}
	
	k = k % n
	if k == 0 {
		return
	}
	
	var reverse func(start, end int)
	reverse = func(start, end int) {
		if start >= end {
			return
		}
		nums[start], nums[end] = nums[end], nums[start]
		reverse(start+1, end-1)
	}

	reverse(0, n-1)
	reverse(0, k-1)
	reverse(k, n-1)
}
