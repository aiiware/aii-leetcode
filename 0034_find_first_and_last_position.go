// 0034 - Find First and Last Position of Element in Sorted Array
// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
// Medium - Array, Binary Search

package leetcode

// SearchRange finds the starting and ending position of a given target value.
// Time Complexity: O(log n) binary search
// Space Complexity: O(1)
func SearchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	first := findFirst(nums, target)
	if first == -1 {
		return []int{-1, -1}
	}

	last := findLast(nums, target)
	return []int{first, last}
}

// findFirst finds the first occurrence of target
func findFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1
	first := -1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			first = mid
			right = mid - 1 // Continue searching in left half
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return first
}

// findLast finds the last occurrence of target
func findLast(nums []int, target int) int {
	left, right := 0, len(nums)-1
	last := -1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			last = mid
			left = mid + 1 // Continue searching in right half
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return last
}

// SearchRangeSinglePass finds range with single binary search
func SearchRangeSinglePass(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	// Find left boundary
	left := binarySearchLeft(nums, target)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}

	// Find right boundary
	right := binarySearchRight(nums, target)
	return []int{left, right}
}

// binarySearchLeft finds leftmost position where target could be inserted
func binarySearchLeft(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

// binarySearchRight finds rightmost position where target could be inserted
func binarySearchRight(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left - 1
}