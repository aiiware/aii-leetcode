package leetcode

// 154. Find Minimum in Rotated Sorted Array II
// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/

// findMinBinarySearchII is the optimal solution using binary search that handles duplicates
// Time Complexity: O(log n) average, O(n) worst case (when many duplicates)
// Space Complexity: O(1)
func findMinBinarySearchII(nums []int) int {
	// Edge case: single element
	if len(nums) == 1 {
		return nums[0]
	}

	left, right := 0, len(nums)-1

	// Binary search
	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			// Minimum is in the right half
			left = mid + 1
		} else if nums[mid] < nums[right] {
			// Minimum is in the left half (including mid)
			right = mid
		} else {
			// nums[mid] == nums[right], we can't decide which half
			// Move right pointer left by one to eliminate duplicate
			right--
		}
	}

	return nums[left]
}

// findMinLinearII is a simple O(n) solution for comparison
// Time Complexity: O(n), Space Complexity: O(1)
func findMinLinearII(nums []int) int {
	minVal := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < minVal {
			minVal = nums[i]
		}
	}
	return minVal
}

// findMinII is the main function that uses binary search (optimal solution)
func findMinII(nums []int) int {
	return findMinBinarySearchII(nums)
}

// findMinRecursiveII is a recursive binary search implementation that handles duplicates
// Time Complexity: O(log n) average, O(n) worst case
// Space Complexity: O(log n) due to recursion stack
func findMinRecursiveII(nums []int) int {
	return findMinRecursiveHelperII(nums, 0, len(nums)-1)
}

func findMinRecursiveHelperII(nums []int, left, right int) int {
	// Base cases
	if left == right {
		return nums[left]
	}
	if nums[left] < nums[right] {
		return nums[left]
	}

	mid := left + (right-left)/2

	if nums[mid] > nums[right] {
		// Minimum is in right half
		return findMinRecursiveHelperII(nums, mid+1, right)
	} else if nums[mid] < nums[right] {
		// Minimum is in left half (including mid)
		return findMinRecursiveHelperII(nums, left, mid)
	} else {
		// nums[mid] == nums[right], handle duplicates
		// Try both halves or move right pointer
		// We'll try moving right pointer left by one
		return findMinRecursiveHelperII(nums, left, right-1)
	}
}

// findMinEarlyExitII is an optimized version that checks for early exit conditions
// This can be faster in cases with many duplicates
func findMinEarlyExitII(nums []int) int {
	left, right := 0, len(nums)-1

	// Early exit: check if array is already sorted
	if nums[left] < nums[right] {
		return nums[left]
	}

	// Binary search with early exit for duplicates
	for left < right {
		// Early exit: if left and right are equal, we found the minimum
		if left == right {
			return nums[left]
		}

		mid := left + (right-left)/2

		// Check if mid is the minimum
		if mid > 0 && nums[mid] < nums[mid-1] {
			return nums[mid]
		}
		if mid < len(nums)-1 && nums[mid+1] < nums[mid] {
			return nums[mid+1]
		}

		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else {
			// Handle duplicates by checking if we can eliminate from left or right
			if nums[mid] == nums[left] {
				left++
			} else {
				right--
			}
		}
	}

	return nums[left]
}

// findMinTwoPointersII is another approach using two pointers from both ends
// This can be more efficient when there are many duplicates
func findMinTwoPointersII(nums []int) int {
	left, right := 0, len(nums)-1

	// Handle edge cases
	if len(nums) == 1 {
		return nums[0]
	}
	if nums[left] < nums[right] {
		return nums[left]
	}

	// Two pointers approach
	for left < right {
		// Skip duplicates from left
		for left < right && nums[left] == nums[left+1] {
			left++
		}
		// Skip duplicates from right
		for left < right && nums[right] == nums[right-1] {
			right--
		}

		// If pointers meet or cross
		if left >= right {
			return nums[left]
		}

		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}