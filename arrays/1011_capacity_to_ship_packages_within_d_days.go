package arrays

// ShipWithinDays solves LeetCode problem 1011: Capacity To Ship Packages Within D Days
// Difficulty: Medium
// Tags: Array, Binary Search
//
// A conveyor belt has packages that must be shipped from one port to another within D days.
//
// The i-th package on the conveyor belt has a weight of weights[i].
// Each day, we load the ship with packages on the conveyor belt (in the order given).
// We may not load more weight than the maximum weight capacity of the ship.
//
// Return the least weight capacity of the ship that will result in all the packages
// being shipped within D days.
//
// Time complexity: O(n log(sum(weights))), Space complexity: O(1)
func ShipWithinDays(weights []int, days int) int {
	n := len(weights)
	if n == 0 {
		return 0
	}

	// Find the minimum and maximum possible capacity
	// Min capacity must be at least the maximum weight (to ship the heaviest package)
	// Max capacity could be the sum of all weights (ship everything in one day)
	maxWeight := 0
	totalWeight := 0
	for _, w := range weights {
		if w > maxWeight {
			maxWeight = w
		}
		totalWeight += w
	}

	// Binary search for the minimum capacity
	left, right := maxWeight, totalWeight

	for left < right {
		mid := left + (right-left)/2

		if canShipWithinDays(weights, days, mid) {
			// Can ship with this capacity, try smaller capacity
			right = mid
		} else {
			// Cannot ship with this capacity, need larger capacity
			left = mid + 1
		}
	}

	return left
}

// Helper function to check if we can ship all packages within given days with given capacity
func canShipWithinDays(weights []int, days int, capacity int) bool {
	currentWeight := 0
	daysNeeded := 1 // Start with first day

	for _, w := range weights {
		// If a single package exceeds capacity, impossible
		if w > capacity {
			return false
		}

		// If adding this package exceeds capacity, start a new day
		if currentWeight+w > capacity {
			daysNeeded++
			currentWeight = w

			// If we've already exceeded the allowed days, return false
			if daysNeeded > days {
				return false
			}
		} else {
			// Add package to current day
			currentWeight += w
		}
	}

	return daysNeeded <= days
}

// ShipWithinDaysOptimized is an optimized version with early exit
func ShipWithinDaysOptimized(weights []int, days int) int {
	n := len(weights)
	if n == 0 {
		return 0
	}

	// Calculate bounds
	maxWeight := 0
	totalWeight := 0
	for _, w := range weights {
		if w > maxWeight {
			maxWeight = w
		}
		totalWeight += w
	}

	// If we have only one day, we need capacity equal to total weight
	if days == 1 {
		return totalWeight
	}

	// If we have at least n days, we need capacity at least maxWeight
	if days >= n {
		return maxWeight
	}

	// Binary search
	left, right := maxWeight, totalWeight

	for left < right {
		mid := left + (right-left)/2

		if canShipWithinDaysOptimized(weights, days, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

// Optimized version of canShipWithinDays with early exit
func canShipWithinDaysOptimized(weights []int, days int, capacity int) bool {
	currentWeight := 0
	daysNeeded := 1

	for _, w := range weights {
		// If a single package exceeds capacity, impossible
		if w > capacity {
			return false
		}

		if currentWeight+w > capacity {
			daysNeeded++
			if daysNeeded > days {
				return false
			}
			currentWeight = w
		} else {
			currentWeight += w
		}
	}

	return true
}