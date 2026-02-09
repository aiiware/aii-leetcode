package arrays

// SearchInRotatedSortedArrayII solves LeetCode problem 0081: Search in Rotated Sorted Array II
// Difficulty: Medium
// Tags: Array, Binary Search
//
// There is an integer array nums sorted in non-decreasing order (not necessarily with distinct values).
//
// Before being passed to your function, nums is rotated at an unknown pivot index k
// (0 <= k < nums.length) such that the resulting array is
// [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]].
//
// For example, [0,1,2,4,4,4,5,6,6,7] might be rotated at pivot index 5 and become
// [4,5,6,6,7,0,1,2,4,4].
//
// Given the array nums after the rotation and an integer target, return true if target
// is in nums, or false otherwise.
//
// You must decrease the overall operation steps as much as possible.
//
// Example 1:
// Input: nums = [2,5,6,0,0,1,2], target = 0
// Output: true
//
// Example 2:
// Input: nums = [2,5,6,0,0,1,2], target = 3
// Output: false
//
// Constraints:
// 1 <= nums.length <= 5000
// -10^4 <= nums[i] <= 10^4
// nums is guaranteed to be rotated at some pivot.
// -10^4 <= target <= 10^4
//
// Time complexity: O(n) in worst case (when many duplicates), O(log n) average
// Space complexity: O(1)
func SearchInRotatedSortedArrayII(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return true
		}

		// The tricky part: when nums[left] == nums[mid] == nums[right]
		// We can't decide which side is sorted, so we shrink the search space
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left++
			right--
			continue
		}

		// Check if left half is sorted
		if nums[left] <= nums[mid] {
			// Target is in the sorted left half
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// Right half is sorted
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return false
}

// SearchInRotatedSortedArrayIILinear is a simple linear search for comparison
func SearchInRotatedSortedArrayIILinear(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}
	return false
}

// SearchInRotatedSortedArrayIIFindPivot finds pivot first then searches
func SearchInRotatedSortedArrayIIFindPivot(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	// Find pivot (smallest element)
	pivot := findPivotWithDuplicates(nums)

	// Binary search in appropriate segment
	if pivot == 0 {
		// Array is not rotated or fully rotated
		return binarySearchWithDuplicates(nums, target)
	}

	// Search in left or right segment
	if target >= nums[0] {
		// Target is in left segment (from 0 to pivot-1)
		return binarySearchWithDuplicates(nums[:pivot], target)
	} else {
		// Target is in right segment (from pivot to end)
		return binarySearchWithDuplicates(nums[pivot:], target)
	}
}

// findPivotWithDuplicates finds the index of the smallest element
func findPivotWithDuplicates(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		// Handle duplicates
		if nums[mid] == nums[right] {
			right--
		} else if nums[mid] > nums[right] {
			// Pivot is in right half
			left = mid + 1
		} else {
			// Pivot is in left half (including mid)
			right = mid
		}
	}

	return left
}

// binarySearchWithDuplicates performs binary search with duplicate handling
func binarySearchWithDuplicates(nums []int, target int) bool {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return true
		}

		// Handle duplicates
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left++
			right--
			continue
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

// SearchInRotatedSortedArrayIIRecursive is a recursive implementation
func SearchInRotatedSortedArrayIIRecursive(nums []int, target int) bool {
	return searchRecursive(nums, target, 0, len(nums)-1)
}

func searchRecursive(nums []int, target, left, right int) bool {
	if left > right {
		return false
	}

	mid := left + (right-left)/2

	if nums[mid] == target {
		return true
	}

	// Handle duplicates
	if nums[left] == nums[mid] && nums[mid] == nums[right] {
		return searchRecursive(nums, target, left+1, right-1)
	}

	// Check which side is sorted
	if nums[left] <= nums[mid] {
		// Left side is sorted
		if nums[left] <= target && target < nums[mid] {
			return searchRecursive(nums, target, left, mid-1)
		}
		return searchRecursive(nums, target, mid+1, right)
	} else {
		// Right side is sorted
		if nums[mid] < target && target <= nums[right] {
			return searchRecursive(nums, target, mid+1, right)
		}
		return searchRecursive(nums, target, left, mid-1)
	}
}

// SearchInRotatedSortedArrayIIEarlyExit adds early exit optimizations
func SearchInRotatedSortedArrayIIEarlyExit(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	// Early exit: check first and last elements
	if nums[0] == target || nums[len(nums)-1] == target {
		return true
	}

	// If array is small, use linear search
	if len(nums) <= 10 {
		for _, num := range nums {
			if num == target {
				return true
			}
		}
		return false
	}

	return SearchInRotatedSortedArrayII(nums, target)
}

// SearchInRotatedSortedArrayIITwoPass uses two-phase approach
func SearchInRotatedSortedArrayIITwoPass(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}

	// Phase 1: Find rotation point
	rotation := findRotationPoint(nums)

	// Phase 2: Binary search in appropriate segment
	if rotation == 0 {
		// No rotation or full rotation
		return binarySearchBool(nums, target, 0, len(nums)-1)
	}

	// Determine which segment to search
	if target >= nums[0] {
		// Search in left segment
		return binarySearchBool(nums, target, 0, rotation-1)
	} else {
		// Search in right segment
		return binarySearchBool(nums, target, rotation, len(nums)-1)
	}
}

// findRotationPoint finds where the rotation occurs
func findRotationPoint(nums []int) int {
	left, right := 0, len(nums)-1

	// If not rotated or all elements same
	if nums[left] < nums[right] {
		return 0
	}

	for left < right {
		// Skip duplicates from left
		for left < right && nums[left] == nums[left+1] {
			left++
		}
		// Skip duplicates from right
		for left < right && nums[right] == nums[right-1] {
			right--
		}

		if left >= right {
			break
		}

		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			// Rotation is in right half
			left = mid + 1
		} else {
			// Rotation is in left half (including mid)
			right = mid
		}
	}

	return left
}

// binarySearchBool is a standard binary search that returns bool
func binarySearchBool(nums []int, target, left, right int) bool {
	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return true
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
