package arrays

// RemoveDuplicatesII solves LeetCode problem 0080: Remove Duplicates from Sorted Array II
// Difficulty: Medium
// Tags: Array, Two Pointers
//
// Given an integer array nums sorted in non-decreasing order, remove some duplicates in-place
// such that each unique element appears at most twice. The relative order of the elements
// should be kept the same.
//
// Since it is impossible to change the length of the array in some languages, you must instead
// have the result be placed in the first part of the array nums. More formally, if there are
// k elements after removing the duplicates, then the first k elements of nums should hold the
// final result. It does not matter what you leave beyond the first k elements.
//
// Return k after placing the final result in the first k slots of nums.
//
// Do not allocate extra space for another array. You must do this by modifying the input array
// in-place with O(1) extra memory.
//
// Example 1:
// Input: nums = [1,1,1,2,2,3]
// Output: 5, nums = [1,1,2,2,3,_]
// Explanation: Your function should return k = 5, with the first five elements of nums being
// 1, 1, 2, 2, 3 respectively. It does not matter what you leave beyond the returned k.
//
// Example 2:
// Input: nums = [0,0,1,1,1,1,2,3,3]
// Output: 7, nums = [0,0,1,1,2,3,3,_,_]
// Explanation: Your function should return k = 7, with the first seven elements of nums being
// 0, 0, 1, 1, 2, 3, 3 respectively. It does not matter what you leave beyond the returned k.
//
// Constraints:
// 1 <= nums.length <= 3 * 10^4
// -10^4 <= nums[i] <= 10^4
// nums is sorted in non-decreasing order.
//
// Time complexity: O(n)
// Space complexity: O(1)
func RemoveDuplicatesII(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	// Two-pointer approach
	// k tracks the position where next valid element should be placed
	k := 2

	for i := 2; i < len(nums); i++ {
		// We can include nums[i] if it's different from the element at position k-2
		// This ensures at most 2 duplicates
		if nums[i] != nums[k-2] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

// RemoveDuplicatesIISimple is a more explicit version of the same algorithm
func RemoveDuplicatesIISimple(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	// Count tracks how many times we've seen the current element
	count := 1
	k := 1 // position to write next valid element

	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			count++
		} else {
			count = 1
		}

		// Write element if we haven't seen it more than twice
		if count <= 2 {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

// RemoveDuplicatesIIBruteForce is a straightforward but less efficient approach
// It's included for educational purposes
func RemoveDuplicatesIIBruteForce(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	// Use a slice to track the result (violates O(1) space but shows the logic)
	result := make([]int, 0, len(nums))
	count := 1
	result = append(result, nums[0])

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			count++
		} else {
			count = 1
		}

		if count <= 2 {
			result = append(result, nums[i])
		}
	}

	// Copy result back to nums (this is the in-place modification)
	for i := 0; i < len(result); i++ {
		nums[i] = result[i]
	}

	return len(result)
}

// RemoveDuplicatesIIGeneric allows at most n duplicates (generalized version)
func RemoveDuplicatesIIGeneric(nums []int, maxDuplicates int) int {
	if maxDuplicates < 0 {
		return 0
	}

	if len(nums) <= maxDuplicates {
		return len(nums)
	}

	// Special case: maxDuplicates = 0 means keep only unique elements (remove all duplicates)
	if maxDuplicates == 0 {
		if len(nums) == 0 {
			return 0
		}

		k := 1 // position to write next unique element
		for i := 1; i < len(nums); i++ {
			if nums[i] != nums[k-1] {
				nums[k] = nums[i]
				k++
			}
		}
		return k
	}

	k := maxDuplicates // position to write next valid element

	for i := maxDuplicates; i < len(nums); i++ {
		// Check if current element is different from element at k-maxDuplicates
		if nums[i] != nums[k-maxDuplicates] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

// RemoveDuplicatesIIWithMap uses a frequency map (not optimal but shows alternative approach)
// Note: This doesn't fully satisfy O(1) space but works for understanding
func RemoveDuplicatesIIWithMap(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	// Track frequency of each element
	freq := make(map[int]int)
	k := 0

	for i := 0; i < len(nums); i++ {
		freq[nums[i]]++
		if freq[nums[i]] <= 2 {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

// RemoveDuplicatesIIEarlyExit optimizes by skipping large duplicate blocks
func RemoveDuplicatesIIEarlyExit(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	k := 2 // position to write

	for i := 2; i < n; i++ {
		// If current element is same as element at k-2, skip it
		if nums[i] == nums[k-2] {
			continue
		}

		// If we have a block of more than 2 duplicates, we can skip ahead
		// Find the next different element
		if i+1 < n && nums[i] == nums[i+1] && nums[i] == nums[k-2] {
			// Skip this entire duplicate block
			for i+1 < n && nums[i] == nums[i+1] {
				i++
			}
			continue
		}

		nums[k] = nums[i]
		k++
	}

	return k
}
