package main

import (
	"fmt"
	"leetcode/data_structures"
)

func main() {
	fmt.Println("=== QuadTree Spatial Data Structure Demo ===")
	fmt.Println()

	// Example 1: Basic QuadTree usage
	fmt.Println("1. Basic QuadTree Operations")
	fmt.Println("---------------------------")
	
	// Create a QuadTree covering area from (0,0) to (100,100)
	bounds := quadtree.Bounds{X: 0, Y: 0, Width: 100, Height: 100}
	qt := quadtree.NewQuadTree(bounds, 4, 5) // capacity 4, max depth 5
	
	// Insert some points
	points := []quadtree.Point{
		{X: 10, Y: 10},
		{X: 20, Y: 20},
		{X: 30, Y: 30},
		{X: 40, Y: 40},
		{X: 50, Y: 50},
		{X: 60, Y: 60},
		{X: 70, Y: 70},
		{X: 80, Y: 80},
	}
	
	for _, p := range points {
		qt.Insert(p)
	}
	
	fmt.Printf("Inserted %d points into QuadTree\n", qt.Count())
	
	// Query a region
	queryBounds := quadtree.Bounds{X: 0, Y: 0, Width: 40, Height: 40}
	queriedPoints := qt.Query(queryBounds)
	fmt.Printf("Points in region (0,0) to (40,40): %d points\n", len(queriedPoints))
	for _, p := range queriedPoints {
		fmt.Printf("  Point at (%.1f, %.1f)\n", p.X, p.Y)
	}
	
	// Example 2: QuadTree subdivision
	fmt.Println("\n2. QuadTree Subdivision Example")
	fmt.Println("------------------------------")
	
	// Create a QuadTree with small capacity to force subdivision
	qt2 := quadtree.NewQuadTree(bounds, 2, 4)
	
	// Insert points that will cause subdivision
	subdivisionPoints := []quadtree.Point{
		{X: 10, Y: 10}, // NW quadrant
		{X: 10, Y: 60}, // SW quadrant  
		{X: 60, Y: 10}, // NE quadrant
		{X: 60, Y: 60}, // SE quadrant
		{X: 15, Y: 15}, // NW quadrant - will go to child node
	}
	
	for _, p := range subdivisionPoints {
		qt2.Insert(p)
	}
	
	fmt.Printf("QuadTree with subdivision has %d points\n", qt2.Count())
	
	// Query each quadrant
	nwBounds := quadtree.Bounds{X: 0, Y: 0, Width: 50, Height: 50}
	nwPoints := qt2.Query(nwBounds)
	fmt.Printf("Points in NW quadrant: %d\n", len(nwPoints))
	
	neBounds := quadtree.Bounds{X: 50, Y: 0, Width: 50, Height: 50}
	nePoints := qt2.Query(neBounds)
	fmt.Printf("Points in NE quadrant: %d\n", len(nePoints))
	
	swBounds := quadtree.Bounds{X: 0, Y: 50, Width: 50, Height: 50}
	swPoints := qt2.Query(swBounds)
	fmt.Printf("Points in SW quadrant: %d\n", len(swPoints))
	
	seBounds := quadtree.Bounds{X: 50, Y: 50, Width: 50, Height: 50}
	sePoints := qt2.Query(seBounds)
	fmt.Printf("Points in SE quadrant: %d\n", len(sePoints))
	
	// Example 3: Performance demonstration
	fmt.Println("\n3. Performance Demonstration")
	fmt.Println("---------------------------")
	
	// Create a larger QuadTree
	largeBounds := quadtree.Bounds{X: 0, Y: 0, Width: 1000, Height: 1000}
	qt3 := quadtree.NewQuadTree(largeBounds, 10, 8)
	
	// Insert many points
	pointCount := 500
	fmt.Printf("Inserting %d points into large QuadTree...\n", pointCount)
	for i := 0; i < pointCount; i++ {
		// Distribute points somewhat randomly
		x := float64((i*37)%1000) * 0.1
		y := float64((i*73)%1000) * 0.1
		qt3.Insert(quadtree.Point{X: x, Y: y})
	}
	
	fmt.Printf("Total points in QuadTree: %d\n", qt3.Count())
	
	// Query a small region - this is efficient with QuadTree
	smallQuery := quadtree.Bounds{X: 450, Y: 450, Width: 10, Height: 10}
	pointsInSmallRegion := qt3.Query(smallQuery)
	fmt.Printf("Points in small region (450,450) to (460,460): %d points\n", len(pointsInSmallRegion))
	
	// Query a large region
	largeQuery := quadtree.Bounds{X: 0, Y: 0, Width: 500, Height: 500}
	pointsInLargeRegion := qt3.Query(largeQuery)
	fmt.Printf("Points in large region (0,0) to (500,500): %d points\n", len(pointsInLargeRegion))
	
	// Example 4: Bounds checking
	fmt.Println("\n4. Bounds Checking Examples")
	fmt.Println("--------------------------")
	
	// Try to insert point outside bounds
	outsidePoint := quadtree.Point{X: 150, Y: 150}
	inserted := qt.Insert(outsidePoint)
	fmt.Printf("Inserting point at (150,150) into bounds (0,0)-(100,100): %v\n", inserted)
	
	// Insert point on boundary
	boundaryPoint := quadtree.Point{X: 100, Y: 100}
	inserted = qt.Insert(boundaryPoint)
	fmt.Printf("Inserting point on boundary at (100,100): %v\n", inserted)
	
	// Query with no intersection
	noIntersectQuery := quadtree.Bounds{X: 200, Y: 200, Width: 50, Height: 50}
	noIntersectPoints := qt.Query(noIntersectQuery)
	fmt.Printf("Points in non-intersecting query region: %d\n", len(noIntersectPoints))
	
	// Example 5: String representation
	fmt.Println("\n5. QuadTree Structure")
	fmt.Println("--------------------")
	
	// Create a simple QuadTree to show structure
	simpleBounds := quadtree.Bounds{X: 0, Y: 0, Width: 100, Height: 100}
	simpleQt := quadtree.NewQuadTree(simpleBounds, 2, 3)
	
	simpleQt.Insert(quadtree.Point{X: 10, Y: 10})
	simpleQt.Insert(quadtree.Point{X: 60, Y: 60})
	simpleQt.Insert(quadtree.Point{X: 20, Y: 20}) // Will cause subdivision
	
	fmt.Println("QuadTree structure (indented by depth):")
	fmt.Print(simpleQt.String())
	
	fmt.Println("\n=== QuadTree Demo Complete ===")
	fmt.Println("\nQuadTree is useful for:")
	fmt.Println("- Spatial indexing of points")
	fmt.Println("- Efficient range queries")
	fmt.Println("- Collision detection in games")
	fmt.Println("- Geographic information systems")
	fmt.Println("- Image processing (region queries)")
}