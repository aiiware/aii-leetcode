package arrays

// FirstBadVersion solves LeetCode problem 278: First Bad Version
// Difficulty: Easy
// Tags: Binary Search, Interactive
//
// You are a product manager and currently leading a team to develop a new product.
// Unfortunately, the latest version of your product fails the quality check.
// Since each version is developed based on the previous version, all the versions
// after a bad version are also bad.
//
// Suppose you have n versions [1, 2, ..., n] and you want to find out the first
// bad one, which causes all the following ones to be bad.
//
// You are given an API bool isBadVersion(version) which returns whether version
// is bad. Implement a function to find the first bad version. You should minimize
// the number of calls to the API.
//
// Time complexity: O(log n), Space complexity: O(1)
func FirstBadVersion(n int, isBadVersion func(int) bool) int {
	// Edge case: if n is 0 or negative (shouldn't happen per problem constraints)
	if n <= 0 {
		return 0
	}

	left, right := 1, n

	// Binary search to find the first bad version
	for left < right {
		// Calculate mid point (prevents overflow for large n)
		mid := left + (right-left)/2

		if isBadVersion(mid) {
			// Mid is bad, so first bad version is at mid or before
			right = mid
		} else {
			// Mid is good, so first bad version is after mid
			left = mid + 1
		}
	}

	// When left == right, we've found the first bad version
	return left
}