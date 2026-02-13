# Data Structures

This directory contains implementations of various data structures used in LeetCode solutions and as standalone utilities.

## Available Data Structures

### 1. QuadTree
A spatial partitioning data structure for efficient 2D point storage and range queries.

**File**: `quadtree.go`

#### Overview
A QuadTree is a tree data structure used to partition a 2D space by recursively subdividing it into four quadrants. It's particularly useful for spatial indexing, range queries, collision detection, and nearest neighbor searches.

#### Key Features
- **Efficient Insertion**: O(log n) average case for well-distributed points
- **Fast Range Queries**: Quickly find all points within a rectangular region
- **Automatic Subdivision**: Automatically subdivides when node capacity is exceeded
- **Depth Limiting**: Prevents infinite recursion with configurable max depth
- **Bounds Checking**: Ensures points are within the tree's bounds

#### Data Structures

```go
// Point represents a 2D point with x and y coordinates
type Point struct {
    X float64
    Y float64
}

// Bounds represents a rectangular region in 2D space
type Bounds struct {
    X      float64 // x-coordinate of top-left corner
    Y      float64 // y-coordinate of top-left corner
    Width  float64 // width of the bounds
    Height float64 // height of the bounds
}

// QuadTree represents a spatial partitioning data structure
type QuadTree struct {
    root *QuadTreeNode
}
```

#### API Reference

##### Constructor
```go
// NewQuadTree creates a new QuadTree with the given bounds and capacity
func NewQuadTree(bounds Bounds, capacity int, maxDepth int) *QuadTree
```
- `bounds`: The rectangular region covered by the QuadTree
- `capacity`: Maximum points per node before subdivision occurs
- `maxDepth`: Maximum recursion depth to prevent infinite subdivision

##### Core Methods
```go
// Insert adds a point to the QuadTree
// Returns true if insertion was successful, false if point is outside bounds
func (qt *QuadTree) Insert(p Point) bool

// Query returns all points within the given bounds
func (qt *QuadTree) Query(bounds Bounds) []Point

// GetAllPoints returns all points in the QuadTree
func (qt *QuadTree) GetAllPoints() []Point

// Count returns the total number of points in the QuadTree
func (qt *QuadTree) Count() int

// String returns a string representation of the QuadTree structure
func (qt *QuadTree) String() string
```

##### Bounds Methods
```go
// Contains checks if a point is within these bounds
func (b *Bounds) Contains(p Point) bool

// Intersects checks if two bounds intersect
func (b *Bounds) Intersects(other Bounds) bool
```

#### Usage Example

```go
package main

import (
    "fmt"
    "leetcode/data_structures"
)

func main() {
    // Create a QuadTree covering area from (0,0) to (100,100)
    bounds := data_structures.Bounds{X: 0, Y: 0, Width: 100, Height: 100}
    qt := data_structures.NewQuadTree(bounds, 4, 5) // capacity 4, max depth 5

    // Insert points
    points := []data_structures.Point{
        {X: 10, Y: 10},
        {X: 20, Y: 20},
        {X: 30, Y: 30},
        {X: 40, Y: 40},
        {X: 50, Y: 50}, // This triggers subdivision
    }

    for _, p := range points {
        qt.Insert(p)
    }

    // Query a region
    queryBounds := data_structures.Bounds{X: 0, Y: 0, Width: 40, Height: 40}
    queriedPoints := qt.Query(queryBounds)
    fmt.Printf("Points in region: %d\n", len(queriedPoints))

    // Get all points
    allPoints := qt.GetAllPoints()
    fmt.Printf("Total points: %d\n", len(allPoints))

    // Count points
    count := qt.Count()
    fmt.Printf("Count: %d\n", count)
}
```

#### Performance Characteristics

| Operation | Average Case | Worst Case |
|-----------|--------------|------------|
| Insert    | O(log n)     | O(n)       |
| Query     | O(log n + k) | O(n)       |
| Space     | O(n)         | O(n)       |

Where:
- `n` = total number of points
- `k` = number of points returned by query

The worst case occurs when points are poorly distributed (e.g., all points clustered in one small region), causing deep recursion without effective partitioning.

#### Use Cases

1. **Spatial Indexing**: Efficiently store and retrieve points in 2D space
2. **Range Queries**: Find all points within a rectangular area
3. **Collision Detection**: In games or simulations, quickly find objects near each other
4. **Geographic Information Systems**: Store and query geographic data
5. **Image Processing**: Query pixel regions efficiently

#### Configuration Tips

1. **Capacity**: Lower values lead to more subdivisions and deeper trees, which can improve query performance for dense datasets but increase memory usage.
2. **Max Depth**: Prevents infinite recursion. Set based on expected point distribution and tree size.
3. **Bounds**: Should cover the entire expected coordinate space.

#### Testing
Comprehensive tests are available in `quadtree_test.go`, covering:
- Basic operations (insert, query, count)
- Subdivision behavior
- Bounds checking
- Edge cases
- Performance characteristics

#### Demo Program
A complete demonstration program is available at `cmd/quadtree_demo/main.go` showing:
- Basic usage examples
- Subdivision demonstrations
- Performance comparisons
- Bounds checking examples
- Tree structure visualization

#### Related LeetCode Problems
QuadTrees are useful for problems involving:
- Spatial data organization
- Range sum queries
- Nearest neighbor searches
- Collision detection
- Geographic data processing

## Directory Structure
```
data_structures/
├── quadtree.go              # QuadTree implementation
├── quadtree_test.go         # QuadTree tests
├── 0705_design_hashset.go   # HashSet implementation (LeetCode 705)
├── 0706_design_hashmap.go   # HashMap implementation (LeetCode 706)
├── 0707_design_linked_list.go # Linked list implementation (LeetCode 707)
└── README.md               # This file
```

## Adding New Data Structures

When adding a new data structure to this directory:

1. **Naming Convention**: Use descriptive names (e.g., `segment_tree.go`, `trie.go`)
2. **Documentation**: Add comprehensive doc comments and examples
3. **Tests**: Include thorough test coverage
4. **Benchmarks**: Consider adding benchmark tests for performance-critical structures
5. **Update README**: Add documentation here with:
   - Overview and use cases
   - API reference
   - Usage examples
   - Performance characteristics

## Performance Considerations

Data structures in this directory are optimized for:
- **Clarity**: Readable implementations for learning
- **Correctness**: Thoroughly tested edge cases
- **Practical Performance**: Reasonable efficiency for LeetCode-style problems

For production use with extremely large datasets, consider:
- Using specialized libraries
- Adding additional optimizations
- Implementing concurrent versions