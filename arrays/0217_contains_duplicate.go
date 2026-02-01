package arrays

// ContainsDuplicate solves LeetCode problem 0217: Contains Duplicate
// Difficulty: Easy
// Tags: Array, Hash Table, Sorting
//
// Given an integer array nums, return true if any value appears at least twice
// in the array, and return false if every element is distinct.
//
// Time complexity: O(n), Space complexity: O(n)
func ContainsDuplicate(nums []int) bool {
	// Create a map to track seen numbers
	seen := make(map[int]bool)

	for _, num := range nums {
		// If we've seen this number before, we found a duplicate
		if seen[num] {
			return true
		}
		// Mark this number as seen
		seen[num] = true
	}

	// No duplicates found
	return false
}

// ContainsDuplicateSorting solves the same problem using sorting
// Time complexity: O(n log n), Space complexity: O(1) or O(n) depending on sort implementation
func ContainsDuplicateSorting(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	// Create a copy to avoid modifying the input
	sorted := make([]int, len(nums))
	copy(sorted, nums)

	// Sort the array (in practice, you'd use sort.Ints)
	// For demonstration, we'll implement a simple bubble sort
	for i := 0; i < len(sorted)-1; i++ {
		for j := 0; j < len(sorted)-i-1; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}

	// Check adjacent elements for duplicates
	for i := 0; i < len(sorted)-1; i++ {
		if sorted[i] == sorted[i+1] {
			return true
		}
	}

	return false
}