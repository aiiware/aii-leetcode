package arrays

// 0219. Contains Duplicate II
// https://leetcode.com/problems/contains-duplicate-ii/
//
// Given an integer array nums and an integer k, return true if there are two
// distinct indices i and j in the array such that nums[i] == nums[j] and
// abs(i - j) <= k.

// containsNearbyDuplicate uses a sliding window with a hash map to track
// the most recent index of each number.
// Time complexity: O(n), Space complexity: O(min(n, k))
func containsNearbyDuplicate(nums []int, k int) bool {
	if k <= 0 {
		return false
	}

	// Map to store the most recent index of each number
	indexMap := make(map[int]int)

	for i, num := range nums {
		// Check if we've seen this number before within distance k
		if lastIndex, exists := indexMap[num]; exists {
			if i-lastIndex <= k {
				return true
			}
		}
		// Update the most recent index for this number
		indexMap[num] = i
	}

	return false
}

// containsNearbyDuplicateSlidingWindow uses a sliding window of size k+1
// implemented with a set to check for duplicates within the window.
// Time complexity: O(n), Space complexity: O(k)
func containsNearbyDuplicateSlidingWindow(nums []int, k int) bool {
	if k <= 0 {
		return false
	}

	// Set to store numbers in the current window
	windowSet := make(map[int]bool)

	for i, num := range nums {
		// If the number is already in the window, we found a duplicate
		if windowSet[num] {
			return true
		}

		// Add current number to the window
		windowSet[num] = true

		// If window size exceeds k, remove the oldest element
		if i >= k {
			windowSet[nums[i-k]] = false
			// Actually remove it from the map to save space
			delete(windowSet, nums[i-k])
		}
	}

	return false
}

// containsNearbyDuplicateBruteForce is a naive O(n*k) solution for comparison.
func containsNearbyDuplicateBruteForce(nums []int, k int) bool {
	if k <= 0 {
		return false
	}

	for i := 0; i < len(nums); i++ {
		// Check up to k elements ahead
		end := i + k
		if end >= len(nums) {
			end = len(nums) - 1
		}
		for j := i + 1; j <= end; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}
