package data_structures

import (
	"testing"
)

func TestQuadTreeBasicOperations(t *testing.T) {
	// Create a QuadTree covering area from (0,0) to (100,100)
	bounds := Bounds{X: 0, Y: 0, Width: 100, Height: 100}
	qt := NewQuadTree(bounds, 4, 5) // capacity 4, max depth 5

	// Test 1: Insert points
	points := []Point{
		{X: 10, Y: 10},
		{X: 20, Y: 20},
		{X: 30, Y: 30},
		{X: 40, Y: 40},
		{X: 50, Y: 50}, // This should trigger subdivision
	}

	for i, p := range points {
		if !qt.Insert(p) {
			t.Errorf("Failed to insert point %d: %v", i, p)
		}
	}

	// Test 2: Count points
	count := qt.Count()
	if count != len(points) {
		t.Errorf("Expected count %d, got %d", len(points), count)
	}

	// Test 3: GetAllPoints
	allPoints := qt.GetAllPoints()
	if len(allPoints) != len(points) {
		t.Errorf("Expected %d points from GetAllPoints, got %d", len(points), len(allPoints))
	}

	// Test 4: Query points in a region
	queryBounds := Bounds{X: 0, Y: 0, Width: 30, Height: 30}
	queriedPoints := qt.Query(queryBounds)
	expectedInQuery := 3 // Points at (10,10), (20,20), (30,30)
	if len(queriedPoints) != expectedInQuery {
		t.Errorf("Expected %d points in query region, got %d", expectedInQuery, len(queriedPoints))
	}

	// Test 5: Query points outside region
	queryBounds2 := Bounds{X: 80, Y: 80, Width: 10, Height: 10}
	queriedPoints2 := qt.Query(queryBounds2)
	if len(queriedPoints2) != 0 {
		t.Errorf("Expected 0 points in query region, got %d", len(queriedPoints2))
	}
}

func TestQuadTreeSubdivision(t *testing.T) {
	bounds := Bounds{X: 0, Y: 0, Width: 100, Height: 100}
	qt := NewQuadTree(bounds, 2, 3) // Small capacity to force subdivision

	// Insert points that will cause subdivision
	points := []Point{
		{X: 10, Y: 10}, // NW
		{X: 10, Y: 60}, // SW
		{X: 60, Y: 10}, // NE
		{X: 60, Y: 60}, // SE
		{X: 15, Y: 15}, // NW - should go to child
	}

	for i, p := range points {
		if !qt.Insert(p) {
			t.Errorf("Failed to insert point %d: %v", i, p)
		}
	}

	// Verify all points are present
	count := qt.Count()
	if count != len(points) {
		t.Errorf("Expected count %d after subdivision, got %d", len(points), count)
	}

	// Query each quadrant
	nwBounds := Bounds{X: 0, Y: 0, Width: 50, Height: 50}
	nwPoints := qt.Query(nwBounds)
	if len(nwPoints) != 2 { // (10,10) and (15,15)
		t.Errorf("Expected 2 points in NW quadrant, got %d", len(nwPoints))
	}

	neBounds := Bounds{X: 50, Y: 0, Width: 50, Height: 50}
	nePoints := qt.Query(neBounds)
	if len(nePoints) != 1 { // (60,10)
		t.Errorf("Expected 1 point in NE quadrant, got %d", len(nePoints))
	}

	swBounds := Bounds{X: 0, Y: 50, Width: 50, Height: 50}
	swPoints := qt.Query(swBounds)
	if len(swPoints) != 1 { // (10,60)
		t.Errorf("Expected 1 point in SW quadrant, got %d", len(swPoints))
	}

	seBounds := Bounds{X: 50, Y: 50, Width: 50, Height: 50}
	sePoints := qt.Query(seBounds)
	if len(sePoints) != 1 { // (60,60)
		t.Errorf("Expected 1 point in SE quadrant, got %d", len(sePoints))
	}
}

func TestQuadTreeBoundsChecking(t *testing.T) {
	bounds := Bounds{X: 0, Y: 0, Width: 100, Height: 100}
	qt := NewQuadTree(bounds, 4, 5)

	// Test inserting point outside bounds
	pointOutside := Point{X: 150, Y: 150}
	if qt.Insert(pointOutside) {
		t.Errorf("Should not insert point outside bounds: %v", pointOutside)
	}

	// Test inserting point on boundary
	pointOnBoundary := Point{X: 100, Y: 100}
	if !qt.Insert(pointOnBoundary) {
		t.Errorf("Should insert point on boundary: %v", pointOnBoundary)
	}

	// Test query with no intersection
	queryBounds := Bounds{X: 200, Y: 200, Width: 50, Height: 50}
	points := qt.Query(queryBounds)
	if len(points) != 0 {
		t.Errorf("Expected 0 points for non-intersecting query, got %d", len(points))
	}
}

func TestQuadTreeMaxDepth(t *testing.T) {
	bounds := Bounds{X: 0, Y: 0, Width: 100, Height: 100}
	qt := NewQuadTree(bounds, 1, 2) // Very shallow tree

	// Insert many points - tree should stop subdividing at depth 2
	points := []Point{
		{X: 10, Y: 10},
		{X: 20, Y: 20},
		{X: 30, Y: 30},
		{X: 40, Y: 40},
		{X: 50, Y: 50},
		{X: 60, Y: 60},
		{X: 70, Y: 70},
		{X: 80, Y: 80},
	}

	for i, p := range points {
		if !qt.Insert(p) {
			t.Errorf("Failed to insert point %d: %v", i, p)
		}
	}

	// All points should be inserted (even if some nodes exceed capacity due to max depth)
	count := qt.Count()
	if count != len(points) {
		t.Errorf("Expected all %d points inserted despite max depth, got %d", len(points), count)
	}
}

func TestQuadTreeStringRepresentation(t *testing.T) {
	bounds := Bounds{X: 0, Y: 0, Width: 100, Height: 100}
	qt := NewQuadTree(bounds, 2, 3)

	// Insert some points
	qt.Insert(Point{X: 10, Y: 10})
	qt.Insert(Point{X: 60, Y: 60})

	str := qt.String()
	if str == "" {
		t.Error("String representation should not be empty")
	}

	// Basic check that string contains expected info
	if len(str) < 50 {
		t.Errorf("String representation seems too short: %s", str)
	}
}

func TestBoundsContains(t *testing.T) {
	bounds := Bounds{X: 0, Y: 0, Width: 100, Height: 100}

	testCases := []struct {
		point    Point
		expected bool
		desc     string
	}{
		{Point{X: 50, Y: 50}, true, "point inside bounds"},
		{Point{X: 0, Y: 0}, true, "point at top-left corner"},
		{Point{X: 100, Y: 100}, true, "point at bottom-right corner"},
		{Point{X: -10, Y: 50}, false, "point left of bounds"},
		{Point{X: 150, Y: 50}, false, "point right of bounds"},
		{Point{X: 50, Y: -10}, false, "point above bounds"},
		{Point{X: 50, Y: 150}, false, "point below bounds"},
	}

	for _, tc := range testCases {
		result := bounds.Contains(tc.point)
		if result != tc.expected {
			t.Errorf("%s: expected %v, got %v for point %v", tc.desc, tc.expected, result, tc.point)
		}
	}
}

func TestBoundsIntersects(t *testing.T) {
	bounds1 := Bounds{X: 0, Y: 0, Width: 100, Height: 100}

	testCases := []struct {
		bounds2  Bounds
		expected bool
		desc     string
	}{
		{Bounds{X: 20, Y: 20, Width: 50, Height: 50}, true, "bounds2 inside bounds1"},
		{Bounds{X: -50, Y: 20, Width: 100, Height: 50}, true, "bounds2 overlaps left edge"},
		{Bounds{X: 50, Y: -50, Width: 50, Height: 100}, true, "bounds2 overlaps top edge"},
		{Bounds{X: 80, Y: 20, Width: 50, Height: 50}, true, "bounds2 overlaps right edge"},
		{Bounds{X: 20, Y: 80, Width: 50, Height: 50}, true, "bounds2 overlaps bottom edge"},
		{Bounds{X: -100, Y: 20, Width: 50, Height: 50}, false, "bounds2 completely left"},
		{Bounds{X: 150, Y: 20, Width: 50, Height: 50}, false, "bounds2 completely right"},
		{Bounds{X: 20, Y: -100, Width: 50, Height: 50}, false, "bounds2 completely above"},
		{Bounds{X: 20, Y: 150, Width: 50, Height: 50}, false, "bounds2 completely below"},
		{Bounds{X: 0, Y: 0, Width: 100, Height: 100}, true, "bounds2 identical to bounds1"},
	}

	for _, tc := range testCases {
		result := bounds1.Intersects(tc.bounds2)
		if result != tc.expected {
			t.Errorf("%s: expected %v, got %v for bounds2 %v", tc.desc, tc.expected, result, tc.bounds2)
		}
	}
}

func TestQuadTreePerformance(t *testing.T) {
	// This test demonstrates the performance benefit of QuadTree
	// by inserting many points and querying a small region
	bounds := Bounds{X: 0, Y: 0, Width: 1000, Height: 1000}
	qt := NewQuadTree(bounds, 10, 8)

	// Insert 1000 random points
	pointCount := 1000
	for i := 0; i < pointCount; i++ {
		// Simple pseudo-random distribution
		x := float64((i*37)%1000) * 0.1
		y := float64((i*73)%1000) * 0.1
		qt.Insert(Point{X: x, Y: y})
	}

	// Verify all points inserted
	count := qt.Count()
	if count != pointCount {
		t.Errorf("Expected %d points, got %d", pointCount, count)
	}

	// Query a small region - should be much faster than checking all points
	queryBounds := Bounds{X: 400, Y: 400, Width: 20, Height: 20}
	pointsInRegion := qt.Query(queryBounds)

	// The exact number depends on the distribution, but should be small
	if len(pointsInRegion) > 50 {
		t.Errorf("Expected relatively few points in small region, got %d", len(pointsInRegion))
	}

	// Verify all returned points are actually in the query region
	for _, p := range pointsInRegion {
		if !queryBounds.Contains(p) {
			t.Errorf("Point %v returned by query but not in query bounds", p)
		}
	}
}