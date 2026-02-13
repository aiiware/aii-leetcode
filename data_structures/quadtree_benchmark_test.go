package data_structures

import (
	"testing"
)

// BenchmarkQuadTreeInsert measures performance of inserting points into a QuadTree
func BenchmarkQuadTreeInsert(b *testing.B) {
	// Create a QuadTree covering a large area
	bounds := Bounds{X: 0, Y: 0, Width: 1000, Height: 1000}
	
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		qt := NewQuadTree(bounds, 10, 8)
		points := generateRandomPoints(1000, 1000, 1000)
		b.StartTimer()
		
		for _, p := range points {
			qt.Insert(p)
		}
	}
}

// BenchmarkQuadTreeQuerySmallRegion measures performance of querying a small region
func BenchmarkQuadTreeQuerySmallRegion(b *testing.B) {
	bounds := Bounds{X: 0, Y: 0, Width: 1000, Height: 1000}
	qt := NewQuadTree(bounds, 10, 8)
	
	// Insert many points
	points := generateRandomPoints(10000, 1000, 1000)
	for _, p := range points {
		qt.Insert(p)
	}
	
	// Query a small region
	queryBounds := Bounds{X: 400, Y: 400, Width: 20, Height: 20}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		qt.Query(queryBounds)
	}
}

// BenchmarkQuadTreeQueryLargeRegion measures performance of querying a large region
func BenchmarkQuadTreeQueryLargeRegion(b *testing.B) {
	bounds := Bounds{X: 0, Y: 0, Width: 1000, Height: 1000}
	qt := NewQuadTree(bounds, 10, 8)
	
	// Insert many points
	points := generateRandomPoints(10000, 1000, 1000)
	for _, p := range points {
		qt.Insert(p)
	}
	
	// Query a large region (half the space)
	queryBounds := Bounds{X: 0, Y: 0, Width: 500, Height: 500}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		qt.Query(queryBounds)
	}
}

// BenchmarkQuadTreeGetAllPoints measures performance of retrieving all points
func BenchmarkQuadTreeGetAllPoints(b *testing.B) {
	bounds := Bounds{X: 0, Y: 0, Width: 1000, Height: 1000}
	
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		qt := NewQuadTree(bounds, 10, 8)
		points := generateRandomPoints(1000, 1000, 1000)
		for _, p := range points {
			qt.Insert(p)
		}
		b.StartTimer()
		
		qt.GetAllPoints()
	}
}

// BenchmarkQuadTreeCount measures performance of counting all points
func BenchmarkQuadTreeCount(b *testing.B) {
	bounds := Bounds{X: 0, Y: 0, Width: 1000, Height: 1000}
	
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		qt := NewQuadTree(bounds, 10, 8)
		points := generateRandomPoints(1000, 1000, 1000)
		for _, p := range points {
			qt.Insert(p)
		}
		b.StartTimer()
		
		qt.Count()
	}
}

// BenchmarkQuadTreeWorstCaseInsert measures performance with worst-case distribution
// (all points clustered in one small area, causing deep trees)
func BenchmarkQuadTreeWorstCaseInsert(b *testing.B) {
	bounds := Bounds{X: 0, Y: 0, Width: 1000, Height: 1000}
	
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		qt := NewQuadTree(bounds, 4, 10) // Small capacity, high depth
		// All points clustered in a small region
		points := generateClusteredPoints(1000, 450, 450, 50, 50)
		b.StartTimer()
		
		for _, p := range points {
			qt.Insert(p)
		}
	}
}

// BenchmarkQuadTreeWorstCaseQuery measures query performance with worst-case distribution
func BenchmarkQuadTreeWorstCaseQuery(b *testing.B) {
	bounds := Bounds{X: 0, Y: 0, Width: 1000, Height: 1000}
	qt := NewQuadTree(bounds, 4, 10)
	
	// Insert clustered points
	points := generateClusteredPoints(10000, 450, 450, 50, 50)
	for _, p := range points {
		qt.Insert(p)
	}
	
	// Query the clustered region
	queryBounds := Bounds{X: 400, Y: 400, Width: 100, Height: 100}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		qt.Query(queryBounds)
	}
}

// BenchmarkLinearSearchComparison compares QuadTree query vs linear search
func BenchmarkLinearSearchComparison(b *testing.B) {
	// Generate test data
	allPoints := generateRandomPoints(10000, 1000, 1000)
	queryBounds := Bounds{X: 400, Y: 400, Width: 20, Height: 20}
	
	// Linear search benchmark
	b.Run("LinearSearch", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			result := make([]Point, 0)
			for _, p := range allPoints {
				if queryBounds.Contains(p) {
					result = append(result, p)
				}
			}
			_ = result
		}
	})
	
	// QuadTree benchmark
	b.Run("QuadTree", func(b *testing.B) {
		bounds := Bounds{X: 0, Y: 0, Width: 1000, Height: 1000}
		qt := NewQuadTree(bounds, 10, 8)
		for _, p := range allPoints {
			qt.Insert(p)
		}
		
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			qt.Query(queryBounds)
		}
	})
}

// BenchmarkQuadTreeMemoryUsage measures memory allocation patterns
func BenchmarkQuadTreeMemoryUsage(b *testing.B) {
	bounds := Bounds{X: 0, Y: 0, Width: 1000, Height: 1000}
	
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		qt := NewQuadTree(bounds, 10, 8)
		points := generateRandomPoints(1000, 1000, 1000)
		for _, p := range points {
			qt.Insert(p)
		}
		_ = qt.Count()
	}
}

// BenchmarkDifferentCapacities compares performance with different node capacities
func BenchmarkDifferentCapacities(b *testing.B) {
	bounds := Bounds{X: 0, Y: 0, Width: 1000, Height: 1000}
	pointCount := 10000
	points := generateRandomPoints(pointCount, 1000, 1000)
	queryBounds := Bounds{X: 400, Y: 400, Width: 20, Height: 20}
	
	capacities := []int{2, 4, 8, 16, 32}
	for _, capacity := range capacities {
		b.Run("Capacity_"+string(rune('0'+capacity)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				qt := NewQuadTree(bounds, capacity, 8)
				for _, p := range points {
					qt.Insert(p)
				}
				b.StartTimer()
				
				qt.Query(queryBounds)
			}
		})
	}
}

// Helper function to generate random points
func generateRandomPoints(count int, maxX, maxY float64) []Point {
	points := make([]Point, count)
	for i := 0; i < count; i++ {
		// Simple pseudo-random distribution
		x := float64((i*37)%int(maxX)) * 0.1
		y := float64((i*73)%int(maxY)) * 0.1
		points[i] = Point{X: x, Y: y}
	}
	return points
}

// Helper function to generate clustered points
func generateClusteredPoints(count int, centerX, centerY, spreadX, spreadY float64) []Point {
	points := make([]Point, count)
	for i := 0; i < count; i++ {
		// Points clustered around center with some spread
		x := centerX + float64((i*17)%int(spreadX)) - spreadX/2
		y := centerY + float64((i*29)%int(spreadY)) - spreadY/2
		points[i] = Point{X: x, Y: y}
	}
	return points
}