// 0035 - Search Insert Position
// https://leetcode.com/problems/search-insert-position/
// Easy - Array, Binary Search

package leetcode

// SearchInsert finds the index where target should be inserted in sorted array.
// Time Complexity: O(log n) binary search
// Space Complexity: O(1)
func SearchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			// Find first occurrence by moving left
			for mid > 0 && nums[mid-1] == target {
				mid--
			}
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// If not found, left is the insertion position
	return left
}

// SearchInsertLinear is a linear search approach (for comparison)
func SearchInsertLinear(nums []int, target int) int {
	for i, num := range nums {
		if num >= target {
			return i
		}
	}
	return len(nums)
}

// SearchInsertRecursive is a recursive binary search approach
func SearchInsertRecursive(nums []int, target int) int {
	return binarySearchInsert(nums, target, 0, len(nums)-1)
}

func binarySearchInsert(nums []int, target, left, right int) int {
	if left > right {
		return left
	}

	mid := left + (right-left)/2

	if nums[mid] == target {
		// Find first occurrence by moving left
		for mid > 0 && nums[mid-1] == target {
			mid--
		}
		return mid
	} else if nums[mid] < target {
		return binarySearchInsert(nums, target, mid+1, right)
	} else {
		return binarySearchInsert(nums, target, left, mid-1)
	}
}