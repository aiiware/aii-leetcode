package arrays

// FindDuplicate solves LeetCode problem 0287: Find the Duplicate Number
// Difficulty: Medium
// Tags: Array, Two Pointers, Binary Search, Bit Manipulation
//
// Given an array of integers nums containing n + 1 integers where each integer
// is in the range [1, n] inclusive.
// There is only one repeated number in nums, return this repeated number.
// You must solve the problem without modifying the array nums and using only
// constant extra space.
//
// Constraints:
// - 1 <= n <= 10^5
// - nums.length == n + 1
// - 1 <= nums[i] <= n
// - All the integers in nums appear only once except for precisely one integer
//   which appears two or more times.
//
// Time complexity: O(n), Space complexity: O(1) using Floyd's algorithm
func FindDuplicate(nums []int) int {
	// Floyd's Tortoise and Hare algorithm (Cycle Detection)
	// Treat the array as a linked list where nums[i] points to nums[nums[i]]
	// Since there's a duplicate, there must be a cycle
	
	// Phase 1: Find intersection point
	slow := nums[0]
	fast := nums[0]
	
	// Move slow by 1 step, fast by 2 steps
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	
	// Phase 2: Find the entrance to the cycle (duplicate number)
	// Reset slow to start, keep fast at meeting point
	// Move both by 1 step until they meet again
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	
	return slow
}

// FindDuplicateBinarySearch solves the problem using binary search
// Time complexity: O(n log n), Space complexity: O(1)
// This approach counts numbers <= mid for each binary search step
func FindDuplicateBinarySearch(nums []int) int {
	// Binary search on the range [1, n] (not on the array indices)
	left, right := 1, len(nums)-1
	
	for left < right {
		mid := left + (right-left)/2
		
		// Count how many numbers are <= mid
		count := 0
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}
		
		// If count > mid, duplicate is in [left, mid]
		// Otherwise, duplicate is in [mid+1, right]
		if count > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	
	return left
}

// FindDuplicateBitManipulation solves the problem using bit manipulation
// Time complexity: O(n log n), Space complexity: O(1)
// This approach compares bits at each position
func FindDuplicateBitManipulation(nums []int) int {
	duplicate := 0
	n := len(nums) - 1
	
	// For each bit position
	for bit := 0; bit < 32; bit++ {
		mask := 1 << bit
		baseCount := 0  // Count of 1s in numbers 1..n
		numCount := 0   // Count of 1s in nums array
		
		// Count 1s in numbers from 1 to n
		for i := 1; i <= n; i++ {
			if i&mask != 0 {
				baseCount++
			}
		}
		
		// Count 1s in nums array
		for _, num := range nums {
			if num&mask != 0 {
				numCount++
			}
		}
		
		// If numCount > baseCount, this bit is set in duplicate
		if numCount > baseCount {
			duplicate |= mask
		}
	}
	
	return duplicate
}

// FindDuplicateHashSet solves the problem using a hash set (violates space constraint)
// Time complexity: O(n), Space complexity: O(n)
// Included for comparison and understanding
func FindDuplicateHashSet(nums []int) int {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return num
		}
		seen[num] = true
	}
	return -1 // Should never reach here given problem constraints
}