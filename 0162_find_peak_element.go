package leetcode

// 0162. Find Peak Element
// https://leetcode.com/problems/find-peak-element

// findPeakElement is the main solution function
func findPeakElement(nums []int) int {
	// Solution 1: Binary search
	return findPeakElementBinarySearch(nums)
}

// ===== Solution 1: Binary search =====
// Time complexity: O(log n)
// Space complexity: O(1)

func findPeakElementBinarySearch(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n == 1 {
		return 0
	}

	left, right := 0, n-1

	for left < right {
		mid := left + (right-left)/2

		// Compare mid with its right neighbor
		if nums[mid] > nums[mid+1] {
			// Peak is in the left half (including mid)
			right = mid
		} else {
			// Peak is in the right half
			left = mid + 1
		}
	}

	return left
}

// ===== Solution 2: Linear scan =====
// Time complexity: O(n)
// Space complexity: O(1)

func findPeakElementLinear(nums []int) int {
	n := len(nums)

	// Check edge cases
	if n == 0 {
		return -1
	}
	if n == 1 {
		return 0
	}

	// Check first element
	if nums[0] > nums[1] {
		return 0
	}

	// Check middle elements
	for i := 1; i < n-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		}
	}

	// Check last element
	if nums[n-1] > nums[n-2] {
		return n - 1
	}

	return -1 // Should not reach here for valid input
}

// ===== Solution 3: Binary search with explicit comparisons =====
// Time complexity: O(log n)
// Space complexity: O(1)

func findPeakElementBinarySearchExplicit(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	left, right := 0, n-1

	for left <= right {
		mid := left + (right-left)/2

		// Check if mid is a peak
		leftVal := getValue(nums, mid-1)
		rightVal := getValue(nums, mid+1)
		midVal := nums[mid]

		if midVal > leftVal && midVal > rightVal {
			return mid
		}

		// Move toward the higher neighbor
		if leftVal > midVal {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

// Helper function to get value with boundary checks
func getValue(nums []int, index int) int {
	if index < 0 || index >= len(nums) {
		return -1 << 31 // Return negative infinity for out of bounds
	}
	return nums[index]
}

// ===== Solution 4: Divide and conquer =====
// Time complexity: O(log n)
// Space complexity: O(log n) due to recursion

func findPeakElementDivideConquer(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	return findPeakHelper(nums, 0, n-1)
}

func findPeakHelper(nums []int, left, right int) int {
	if left == right {
		return left
	}

	mid := left + (right-left)/2

	// Check if mid is a peak
	if (mid == 0 || nums[mid] > nums[mid-1]) && (mid == len(nums)-1 || nums[mid] > nums[mid+1]) {
		return mid
	}

	// If left neighbor is greater, search left half
	if mid > 0 && nums[mid-1] > nums[mid] {
		return findPeakHelper(nums, left, mid-1)
	}

	// Otherwise search right half
	return findPeakHelper(nums, mid+1, right)
}

// ===== Solution 5: Iterative binary search with early exit =====
// Time complexity: O(log n)
// Space complexity: O(1)

func findPeakElementIterative(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	left, right := 0, n-1

	for left <= right {
		// Early exit for small ranges
		if right-left <= 1 {
			if nums[left] >= nums[right] {
				return left
			}
			return right
		}

		mid := left + (right-left)/2

		// Check if mid is a peak
		if (mid == 0 || nums[mid] > nums[mid-1]) && (mid == n-1 || nums[mid] > nums[mid+1]) {
			return mid
		}

		// Move toward the higher neighbor
		if mid > 0 && nums[mid-1] > nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}