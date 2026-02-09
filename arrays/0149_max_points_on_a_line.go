package arrays

// MaxPoints solves LeetCode problem 0149: Max Points on a Line
// Difficulty: Hard
// Tags: Array, Hash Table, Math, Geometry
//
// Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane,
// return the maximum number of points that lie on the same straight line.
//
// Example 1:
// Input: points = [[1,1],[2,2],[3,3]]
// Output: 3
//
// Example 2:
// Input: points = [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
// Output: 4
//
// Constraints:
// - 1 <= points.length <= 300
// - points[i].length == 2
// - -10^4 <= xi, yi <= 10^4
// - All the points are unique.
//
// Time complexity: O(n^2), Space complexity: O(n) for storing slope counts
func MaxPoints(points [][]int) int {
	n := len(points)
	if n <= 2 {
		return n
	}

	maxCount := 1

	for i := 0; i < n; i++ {
		// Map to store slope frequencies
		slopeCount := make(map[[2]int]int)
		// Count duplicate points
		duplicate := 1
		// Current max for point i
		currentMax := 0

		for j := i + 1; j < n; j++ {
			dx := points[j][0] - points[i][0]
			dy := points[j][1] - points[i][1]

			// Check for duplicate points
			if dx == 0 && dy == 0 {
				duplicate++
				continue
			}

			// Reduce the slope to simplest form
			g := gcd(dx, dy)
			dx /= g
			dy /= g

			// Normalize to canonical form: ensure consistent direction
			// Make dx positive, or if dx == 0, make dy positive
			if dx < 0 || (dx == 0 && dy < 0) {
				dx = -dx
				dy = -dy
			}

			// Use a tuple to represent the slope
			slope := [2]int{dx, dy}
			slopeCount[slope]++

			// Update current max
			if slopeCount[slope] > currentMax {
				currentMax = slopeCount[slope]
			}
		}

		// Update global max
		if currentMax+duplicate > maxCount {
			maxCount = currentMax + duplicate
		}
	}

	return maxCount
}

// MaxPointsWithFloat is an alternative implementation using float64 for slopes
// Note: This approach may have floating-point precision issues
func MaxPointsWithFloat(points [][]int) int {
	n := len(points)
	if n <= 2 {
		return n
	}

	maxCount := 1

	for i := 0; i < n; i++ {
		slopeCount := make(map[float64]int)
		duplicate := 1
		currentMax := 0

		for j := i + 1; j < n; j++ {
			dx := points[j][0] - points[i][0]
			dy := points[j][1] - points[i][1]

			// Check for duplicate points
			if dx == 0 && dy == 0 {
				duplicate++
				continue
			}

			// Handle vertical lines (infinite slope)
			var slope float64
			if dx == 0 {
				slope = float64(1<<63 - 1) // Represent infinity
			} else {
				slope = float64(dy) / float64(dx)
			}

			slopeCount[slope]++
			if slopeCount[slope] > currentMax {
				currentMax = slopeCount[slope]
			}
		}

		if currentMax+duplicate > maxCount {
			maxCount = currentMax + duplicate
		}
	}

	return maxCount
}

// gcd computes the greatest common divisor using Euclidean algorithm
func gcd(a, b int) int {
	// Make sure a and b are non-negative
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	for b != 0 {
		a, b = b, a%b
	}
	return a
}
