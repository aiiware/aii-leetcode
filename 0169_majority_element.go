package leetcode

import "leetcode/utils"

// 0169. Majority Element
// https://leetcode.com/problems/majority-element
// 
// Problem Description:
// Given an array nums of size n, return the majority element.
// The majority element is the element that appears more than ⌊n / 2⌋ times.
// You may assume that the array is non-empty and the majority element always exists.
//
// Categories: Array, Hash Table, Divide and Conquer, Sorting, Counting
// Difficulty: Easy
// 
// Companies that have asked this problem:
// Amazon, Google, Meta, Apple, Adobe, Microsoft, Uber, Oracle, Bloomberg, Yahoo, Zenefits
//
// Example 1:
// Input: nums = [3,2,3]
// Output: 3
//
// Example 2:
// Input: nums = [2,2,1,1,1,2,2]
// Output: 2

// majorityElement is the main solution function
func majorityElement(nums []int) int {
	// Solution 1: Boyer-Moore Voting Algorithm (most efficient)
	return majorityElementBoyerMoore(nums)
}

// ===== Solution 1: Boyer-Moore Voting Algorithm =====
// Time complexity: O(n)
// Space complexity: O(1)
// This is the optimal solution for this problem

func majorityElementBoyerMoore(nums []int) int {
	candidate := 0
	count := 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}

		if num == candidate {
			count++
		} else {
			count--
		}
	}

	// Since problem guarantees majority element exists, we don't need to verify
	// But for completeness, we could add verification:
	// count = 0
	// for _, num := range nums {
	//     if num == candidate {
	//         count++
	//     }
	// }
	// if count > len(nums)/2 {
	//     return candidate
	// }
	// return -1 // Should not happen

	return candidate
}

// ===== Solution 2: Hash Map (Frequency Count) =====
// Time complexity: O(n)
// Space complexity: O(n)
// Straightforward approach using hash map

func majorityElementHashMap(nums []int) int {
	freq := make(map[int]int)
	n := len(nums)

	for _, num := range nums {
		freq[num]++
		if freq[num] > n/2 {
			return num
		}
	}

	// Should not reach here since majority element always exists
	return -1
}

// ===== Solution 3: Sorting =====
// Time complexity: O(n log n)
// Space complexity: O(1) or O(n) depending on sorting implementation
// After sorting, the majority element will be at index n/2

func majorityElementSorting(nums []int) int {
	// Create a copy to avoid modifying input
	sorted := make([]int, len(nums))
	copy(sorted, nums)
	
	// Sort the array
	utils.SortInts(sorted)
	
	// Majority element will be at the middle
	return sorted[len(sorted)/2]
}

// ===== Solution 4: Divide and Conquer =====
// Time complexity: O(n log n)
// Space complexity: O(log n) due to recursion
// Recursively divide the array and count majority elements in halves

func majorityElementDivideConquer(nums []int) int {
	return majorityElementHelper(nums, 0, len(nums)-1)
}

func majorityElementHelper(nums []int, left, right int) int {
	// Base case: single element
	if left == right {
		return nums[left]
	}

	// Divide
	mid := left + (right-left)/2
	leftMajority := majorityElementHelper(nums, left, mid)
	rightMajority := majorityElementHelper(nums, mid+1, right)

	// If both halves have same majority, return it
	if leftMajority == rightMajority {
		return leftMajority
	}

	// Otherwise, count which one appears more in the combined range
	leftCount := countInRange(nums, leftMajority, left, right)
	rightCount := countInRange(nums, rightMajority, left, right)

	if leftCount > rightCount {
		return leftMajority
	}
	return rightMajority
}

func countInRange(nums []int, target int, left, right int) int {
	count := 0
	for i := left; i <= right; i++ {
		if nums[i] == target {
			count++
		}
	}
	return count
}

// ===== Solution 5: Bit Manipulation =====
// Time complexity: O(32 * n) = O(n) for 32-bit integers
// Space complexity: O(1)
// Count bits position by position

func majorityElementBitManipulation(nums []int) int {
	majority := 0
	n := len(nums)

	// Consider each bit position (for 32-bit integers)
	for bit := 0; bit < 32; bit++ {
		count := 0
		mask := 1 << bit

		// Count numbers with this bit set
		for _, num := range nums {
			if num&mask != 0 {
				count++
			}
		}

		// If more than half have this bit set, set it in result
		if count > n/2 {
			majority |= mask
		}
	}

	return majority
}

// ===== Solution 6: Randomization =====
// Time complexity: O(∞) worst case, O(n) expected
// Space complexity: O(1)
// Randomly pick an element and check if it's majority

func majorityElementRandom(nums []int) int {
	n := len(nums)
	
	// In practice, we would use a random number generator
	// For deterministic testing, we'll just iterate
	for i := 0; i < 20; i++ { // Try up to 20 times
		candidate := nums[i%n]
		count := 0
		
		for _, num := range nums {
			if num == candidate {
				count++
			}
		}
		
		if count > n/2 {
			return candidate
		}
	}
	
	// Should not reach here
	return -1
}

// ===== Solution 7: Moore's Algorithm with Verification =====
// Time complexity: O(n)
// Space complexity: O(1)
// Same as Boyer-Moore but with explicit verification

func majorityElementWithVerification(nums []int) int {
	// Phase 1: Find candidate
	candidate := 0
	count := 0
	
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	
	// Phase 2: Verify candidate
	count = 0
	for _, num := range nums {
		if num == candidate {
			count++
		}
	}
	
	if count > len(nums)/2 {
		return candidate
	}
	
	// According to problem, this should never happen
	return -1
}

// ===== Solution 8: Using Built-in Sort =====
// Time complexity: O(n log n)
// Space complexity: O(n)
// Simple one-liner using sort

func majorityElementBuiltinSort(nums []int) int {
	// Note: This modifies the input array
	// utils.SortInts(nums)
	// return nums[len(nums)/2]
	
	// To avoid modifying input, create a copy
	copyNums := make([]int, len(nums))
	copy(copyNums, nums)
	utils.SortInts(copyNums)
	return copyNums[len(copyNums)/2]
}