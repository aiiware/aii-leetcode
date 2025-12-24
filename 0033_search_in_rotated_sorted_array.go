// 0033 - Search in Rotated Sorted Array
// https://leetcode.com/problems/search-in-rotated-sorted-array/
// Medium - Array, Binary Search

package leetcode

// SearchInRotatedSortedArray searches for target in rotated sorted array.
// Time Complexity: O(log n) binary search
// Space Complexity: O(1)
func SearchInRotatedSortedArray(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		// Check which side is sorted
		if nums[left] <= nums[mid] {
			// Left side is sorted
			if nums[left] <= target && target < nums[mid] {
				// Target is in left sorted side
				right = mid - 1
			} else {
				// Target is in right side
				left = mid + 1
			}
		} else {
			// Right side is sorted
			if nums[mid] < target && target <= nums[right] {
				// Target is in right sorted side
				left = mid + 1
			} else {
				// Target is in left side
				right = mid - 1
			}
		}
	}

	return -1
}

// SearchInRotatedSortedArrayTwoPass finds pivot first, then searches
func SearchInRotatedSortedArrayTwoPass(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	// Find the pivot (smallest element)
	pivot := findPivot(nums)

	// Binary search in the appropriate segment
	if pivot == 0 {
		return binarySearch(nums, target, 0, len(nums)-1)
	}

	if target >= nums[0] && target <= nums[pivot-1] {
		return binarySearch(nums, target, 0, pivot-1)
	}
	return binarySearch(nums, target, pivot, len(nums)-1)
}

// findPivot finds the index of the smallest element (rotation point)
func findPivot(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			// Pivot is in right half
			left = mid + 1
		} else {
			// Pivot is in left half (including mid)
			right = mid
		}
	}

	return left
}

// binarySearch performs standard binary search
func binarySearch(nums []int, target, left, right int) int {
	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}