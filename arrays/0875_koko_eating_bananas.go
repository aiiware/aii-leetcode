package arrays

import "math"

// MinEatingSpeed solves LeetCode problem 875: Koko Eating Bananas
// Difficulty: Medium
// Tags: Array, Binary Search
//
// Koko loves to eat bananas. There are n piles of bananas, the i-th pile has piles[i] bananas.
// The guards have gone and will come back in h hours.
//
// Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile
// of bananas and eats k bananas from that pile. If the pile has less than k bananas, she
// eats all of them instead and will not eat any more bananas during this hour.
//
// Koko likes to eat slowly but still wants to finish eating all the bananas before
// the guards return.
//
// Return the minimum integer k such that she can eat all the bananas within h hours.
//
// Time complexity: O(n log(max(piles))), Space complexity: O(1)
func MinEatingSpeed(piles []int, h int) int {
	n := len(piles)
	if n == 0 {
		return 0
	}

	// Find the maximum pile size (upper bound for k)
	maxPile := 0
	for _, pile := range piles {
		if pile > maxPile {
			maxPile = pile
		}
	}

	// If hours equals number of piles, need speed at least max pile
	if h == n {
		return maxPile
	}

	// Binary search for minimum k
	left, right := 1, maxPile

	for left < right {
		mid := left + (right-left)/2

		if canEatAll(piles, h, mid) {
			// Can eat all with this speed, try slower speed
			right = mid
		} else {
			// Cannot eat all with this speed, need faster speed
			left = mid + 1
		}
	}

	return left
}

// Helper function to check if Koko can eat all bananas within h hours at speed k
func canEatAll(piles []int, h int, k int) bool {
	hoursNeeded := 0

	for _, pile := range piles {
		// Calculate hours needed for this pile
		// Ceiling division: ceil(pile / k)
		hoursNeeded += int(math.Ceil(float64(pile) / float64(k)))

		// Early exit if already exceeds h
		if hoursNeeded > h {
			return false
		}
	}

	return hoursNeeded <= h
}

// MinEatingSpeedOptimized is an optimized version with integer ceiling division
func MinEatingSpeedOptimized(piles []int, h int) int {
	n := len(piles)
	if n == 0 {
		return 0
	}

	// Find bounds
	maxPile := 0
	for _, pile := range piles {
		if pile > maxPile {
			maxPile = pile
		}
	}

	// Special cases
	if h == n {
		return maxPile
	}

	// Calculate total bananas for lower bound optimization
	totalBananas := 0
	for _, pile := range piles {
		totalBananas += pile
	}

	// Lower bound: ceil(totalBananas / h)
	// Upper bound: maxPile
	left := (totalBananas + h - 1) / h // Ceiling division
	right := maxPile

	// If left bound is already valid, we can start there
	if left > right {
		left = right
	}

	// Binary search
	for left < right {
		mid := left + (right-left)/2

		if canEatAllOptimized(piles, h, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

// Optimized version with integer ceiling division
func canEatAllOptimized(piles []int, h int, k int) bool {
	hoursNeeded := 0

	for _, pile := range piles {
		// Integer ceiling division: (pile + k - 1) / k
		hoursNeeded += (pile + k - 1) / k

		// Early exit
		if hoursNeeded > h {
			return false
		}
	}

	return true
}