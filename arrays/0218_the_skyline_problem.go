package arrays

import "sort"

// GetSkyline solves LeetCode problem 0218: The Skyline Problem
// Difficulty: Hard
// Tags: Array, Heap, Divide and Conquer, Line Sweep
//
// A city's skyline is the outer contour of the silhouette formed by all the buildings
// in that city when viewed from a distance. Given the locations and heights of all
// the buildings, return the skyline formed by these buildings collectively.
//
// Time complexity: O(n log n), Space complexity: O(n)
func GetSkyline(buildings [][]int) [][]int {
	if len(buildings) == 0 {
		return [][]int{}
	}

	// Convert buildings to events (start and end points)
	type event struct {
		pos     int
		height  int
		isStart bool
	}

	events := make([]event, 0, len(buildings)*2)
	for _, b := range buildings {
		events = append(events, event{b[0], b[2], true})  // start
		events = append(events, event{b[1], b[2], false}) // end
	}

	// Sort events by position
	// For same position: starts before ends
	// For starts: higher first
	// For ends: lower first
	sort.Slice(events, func(i, j int) bool {
		if events[i].pos != events[j].pos {
			return events[i].pos < events[j].pos
		}
		// Same position: starts before ends
		if events[i].isStart != events[j].isStart {
			return events[i].isStart
		}
		// Both starts: higher first
		if events[i].isStart {
			return events[i].height > events[j].height
		}
		// Both ends: lower first
		return events[i].height < events[j].height
	})

	// Use a map to track active building heights (height -> count)
	heights := make(map[int]int)
	var currentMaxHeight int

	result := make([][]int, 0)

	for i := 0; i < len(events); i++ {
		pos := events[i].pos
		height := events[i].height
		isStart := events[i].isStart

		if isStart {
			heights[height]++
			if height > currentMaxHeight {
				currentMaxHeight = height
				result = append(result, []int{pos, height})
			}
		} else {
			heights[height]--
			if heights[height] == 0 {
				delete(heights, height)
			}

			// Recalculate max height
			newMaxHeight := 0
			for h := range heights {
				if h > newMaxHeight {
					newMaxHeight = h
				}
			}

			// If height changed, add a point
			if newMaxHeight != currentMaxHeight {
				result = append(result, []int{pos, newMaxHeight})
				currentMaxHeight = newMaxHeight
			}
		}
	}

	return result
}
