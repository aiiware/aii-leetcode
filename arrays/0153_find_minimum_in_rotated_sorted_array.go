package arrays

// 153. Find Minimum in Rotated Sorted Array
// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/

// findMinBinarySearch is the optimal O(log n) solution using binary search
// Time Complexity: O(log n), Space Complexity: O(1)
func findMinBinarySearch(nums []int) int {
	// Edge case: single element
	if len(nums) == 1 {
		return nums[0]
	}

	left, right := 0, len(nums)-1

	// If the array is not rotated (already sorted)
	if nums[left] < nums[right] {
		return nums[left]
	}

	// Binary search to find the pivot point (minimum)
	for left <= right {
		mid := left + (right-left)/2

		// Check if mid is the minimum
		// Minimum element is the one where the next element is smaller
		if mid > 0 && nums[mid] < nums[mid-1] {
			return nums[mid]
		}
		// Also check if mid+1 is the minimum
		if mid < len(nums)-1 && nums[mid+1] < nums[mid] {
			return nums[mid+1]
		}

		// Decide which half to search
		if nums[mid] > nums[0] {
			// Left half is sorted, minimum is in right half
			left = mid + 1
		} else {
			// Right half is sorted, minimum is in left half
			right = mid - 1
		}
	}

	// Should never reach here for valid input
	return nums[0]
}

// findMinLinear is a simple O(n) solution for comparison
// Time Complexity: O(n), Space Complexity: O(1)
func findMinLinear(nums []int) int {
	minVal := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < minVal {
			minVal = nums[i]
		}
	}
	return minVal
}

// findMin is the main function that uses binary search (optimal solution)
func findMin(nums []int) int {
	return findMinBinarySearch(nums)
}

// findMinAlt is an alternative implementation with cleaner binary search
// This is the most common implementation found in solutions
// Time Complexity: O(log n), Space Complexity: O(1)
func findMinAlt(nums []int) int {
	left, right := 0, len(nums)-1

	// Binary search
	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			// Minimum is in the right half
			left = mid + 1
		} else {
			// Minimum is in the left half (including mid)
			right = mid
		}
	}

	return nums[left]
}

// findMinRecursive is a recursive binary search implementation
// Time Complexity: O(log n), Space Complexity: O(log n) due to recursion stack
func findMinRecursive(nums []int) int {
	return findMinRecursiveHelper(nums, 0, len(nums)-1)
}

func findMinRecursiveHelper(nums []int, left, right int) int {
	// Base cases
	if left == right {
		return nums[left]
	}
	if nums[left] < nums[right] {
		return nums[left]
	}

	mid := left + (right-left)/2

	// Check if mid is the minimum
	if mid > left && nums[mid] < nums[mid-1] {
		return nums[mid]
	}
	if mid < right && nums[mid+1] < nums[mid] {
		return nums[mid+1]
	}

	// Recursive search
	if nums[mid] > nums[left] {
		// Minimum is in right half
		return findMinRecursiveHelper(nums, mid+1, right)
	} else {
		// Minimum is in left half
		return findMinRecursiveHelper(nums, left, mid-1)
	}
}

// findMinOnePass is another variation that finds minimum in one pass
// This is similar to findMinAlt but with different condition
func findMinOnePass(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		// Compare with the leftmost element
		if nums[mid] < nums[0] {
			// Minimum is in left half (including mid)
			right = mid
		} else {
			// Minimum is in right half
			left = mid + 1
		}
	}

	// After loop, left == right
	// Check if we found the minimum or if array wasn't rotated
	if nums[left] < nums[0] {
		return nums[left]
	}
	return nums[0]
}