package arrays

// MergeSortedArray merges two sorted arrays into one sorted array
// Difficulty: Easy
// Tags: Array, Two Pointers
//
// Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.
// The number of elements initialized in nums1 and nums2 are m and n respectively.
// You may assume that nums1 has enough space (size that is equal to m + n) to hold additional elements from nums2.
//
// Time complexity: O(m+n), Space complexity: O(1)
func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	// Start from the end of both arrays
	p1 := m - 1
	p2 := n - 1
	p := m + n - 1

	// Merge from the back
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
		p--
	}

	// If there are remaining elements in nums2, copy them
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p2--
		p--
	}
}
