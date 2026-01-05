package leetcode

// Problem 0088: Merge Sorted Array
//
// You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, 
// and two integers m and n, representing the number of elements in nums1 and nums2 respectively.
//
// Merge nums1 and nums2 into a single array sorted in non-decreasing order.
//
// The final sorted array should not be returned by the function, but instead be stored 
// inside the array nums1. To accommodate this, nums1 has a length of m + n, where the 
// first m elements denote the elements that should be merged, and the last n elements 
// are set to 0 and should be ignored. nums2 has a length of n.
//
// Example 1:
// Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// Output: [1,2,2,3,5,6]
// Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
// The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.
//
// Example 2:
// Input: nums1 = [1], m = 1, nums2 = [], n = 0
// Output: [1]
// Explanation: The arrays we are merging are [1] and [].
// The result of the merge is [1].
//
// Example 3:
// Input: nums1 = [0], m = 0, nums2 = [1], n = 1
// Output: [1]
// Explanation: The arrays we are merging are [] and [1].
// The result of the merge is [1].
//
// Constraints:
// - nums1.length == m + n
// - nums2.length == n
// - 0 <= m, n <= 200
// - 1 <= m + n <= 200
// - -10^9 <= nums1[i], nums2[j] <= 10^9

// mergeSortedArray is the main solution function using three pointers from the end.
// Time complexity: O(m + n), Space complexity: O(1)
func mergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	// Start from the end of both arrays
	i, j, k := m-1, n-1, m+n-1

	// Merge from the end to avoid overwriting nums1 elements
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	// If there are remaining elements in nums2, copy them
	// No need to copy remaining nums1 elements as they're already in place
	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}

// mergeSortedArrayTwoPointers is an alternative solution using two pointers from start.
// This requires O(m) extra space to preserve nums1 elements.
func mergeSortedArrayTwoPointers(nums1 []int, m int, nums2 []int, n int) {
	// Make a copy of nums1's valid elements
	nums1Copy := make([]int, m)
	copy(nums1Copy, nums1[:m])

	i, j, k := 0, 0, 0

	// Merge nums1Copy and nums2 into nums1
	for i < m && j < n {
		if nums1Copy[i] <= nums2[j] {
			nums1[k] = nums1Copy[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
		k++
	}

	// Copy remaining elements from nums1Copy
	for i < m {
		nums1[k] = nums1Copy[i]
		i++
		k++
	}

	// Copy remaining elements from nums2
	for j < n {
		nums1[k] = nums2[j]
		j++
		k++
	}
}

// mergeSortedArraySimple is a simple solution using built-in sort.
// Not optimal but easy to understand.
func mergeSortedArraySimple(nums1 []int, m int, nums2 []int, n int) {
	// Copy nums2 into the end of nums1
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}

	// Sort the entire array
	// In a real implementation, we'd use a proper sort function
	// For simplicity, we'll use bubble sort (not efficient for large arrays)
	for i := 0; i < m+n-1; i++ {
		for j := 0; j < m+n-i-1; j++ {
			if nums1[j] > nums1[j+1] {
				nums1[j], nums1[j+1] = nums1[j+1], nums1[j]
			}
		}
	}
}

// mergeSortedArrayInPlace is another in-place solution.
func mergeSortedArrayInPlace(nums1 []int, m int, nums2 []int, n int) {
	// If nums2 is empty, nothing to do
	if n == 0 {
		return
	}

	// If nums1 has no elements, just copy nums2
	if m == 0 {
		copy(nums1, nums2)
		return
	}

	// Shift nums1 elements to the end to make space
	for i := m - 1; i >= 0; i-- {
		nums1[i+n] = nums1[i]
	}

	// Now merge
	i, j, k := n, 0, 0 // i starts at n (shifted nums1), j for nums2, k for result position

	for i < m+n && j < n {
		if nums1[i] <= nums2[j] {
			nums1[k] = nums1[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
		k++
	}

	// Copy remaining elements from shifted nums1
	for i < m+n {
		nums1[k] = nums1[i]
		i++
		k++
	}

	// Copy remaining elements from nums2
	for j < n {
		nums1[k] = nums2[j]
		j++
		k++
	}
}

// mergeSortedArrayRecursive is a recursive solution (not recommended for large arrays).
func mergeSortedArrayRecursive(nums1 []int, m int, nums2 []int, n int) {
	// Base cases
	if n == 0 {
		return
	}
	if m == 0 {
		copy(nums1, nums2)
		return
	}

	// Compare last elements
	if nums1[m-1] > nums2[n-1] {
		// Move nums1[m-1] to the end
		nums1[m+n-1] = nums1[m-1]
		// Recursively merge the rest
		mergeSortedArrayRecursive(nums1, m-1, nums2, n)
	} else {
		// Move nums2[n-1] to the end
		nums1[m+n-1] = nums2[n-1]
		// Recursively merge the rest
		mergeSortedArrayRecursive(nums1, m, nums2, n-1)
	}
}

// MergeSortedArray is the public interface function.
// It uses the optimal three-pointer solution by default.
func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	mergeSortedArray(nums1, m, nums2, n)
}