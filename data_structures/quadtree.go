package quadtree

import (
	"fmt"
)

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

// Contains checks if a point is within these bounds
func (b *Bounds) Contains(p Point) bool {
	return p.X >= b.X && p.X <= b.X+b.Width &&
		p.Y >= b.Y && p.Y <= b.Y+b.Height
}

// Intersects checks if two bounds intersect
func (b *Bounds) Intersects(other Bounds) bool {
	return !(b.X+b.Width < other.X ||
		other.X+other.Width < b.X ||
		b.Y+b.Height < other.Y ||
		other.Y+other.Height < b.Y)
}

// QuadTreeNode represents a node in the QuadTree
type QuadTreeNode struct {
	bounds    Bounds    // The region this node covers
	capacity  int       // Maximum points before subdivision
	points    []Point   // Points stored in this node
	depth     int       // Depth of this node in the tree
	maxDepth  int       // Maximum depth allowed
	divided   bool      // Whether this node has been subdivided
	northWest *QuadTreeNode
	northEast *QuadTreeNode
	southWest *QuadTreeNode
	southEast *QuadTreeNode
}

// QuadTree represents a spatial partitioning data structure
type QuadTree struct {
	root *QuadTreeNode
}

// NewQuadTree creates a new QuadTree with the given bounds and capacity
func NewQuadTree(bounds Bounds, capacity int, maxDepth int) *QuadTree {
	return &QuadTree{
		root: &QuadTreeNode{
			bounds:   bounds,
			capacity: capacity,
			points:   make([]Point, 0),
			depth:    0,
			maxDepth: maxDepth,
			divided:  false,
		},
	}
}

// subdivide splits a node into four child nodes
func (n *QuadTreeNode) subdivide() {
	if n.divided || n.depth >= n.maxDepth {
		return
	}

	halfWidth := n.bounds.Width / 2
	halfHeight := n.bounds.Height / 2
	x := n.bounds.X
	y := n.bounds.Y

	n.northWest = &QuadTreeNode{
		bounds:   Bounds{X: x, Y: y, Width: halfWidth, Height: halfHeight},
		capacity: n.capacity,
		points:   make([]Point, 0),
		depth:    n.depth + 1,
		maxDepth: n.maxDepth,
		divided:  false,
	}

	n.northEast = &QuadTreeNode{
		bounds:   Bounds{X: x + halfWidth, Y: y, Width: halfWidth, Height: halfHeight},
		capacity: n.capacity,
		points:   make([]Point, 0),
		depth:    n.depth + 1,
		maxDepth: n.maxDepth,
		divided:  false,
	}

	n.southWest = &QuadTreeNode{
		bounds:   Bounds{X: x, Y: y + halfHeight, Width: halfWidth, Height: halfHeight},
		capacity: n.capacity,
		points:   make([]Point, 0),
		depth:    n.depth + 1,
		maxDepth: n.maxDepth,
		divided:  false,
	}

	n.southEast = &QuadTreeNode{
		bounds:   Bounds{X: x + halfWidth, Y: y + halfHeight, Width: halfWidth, Height: halfHeight},
		capacity: n.capacity,
		points:   make([]Point, 0),
		depth:    n.depth + 1,
		maxDepth: n.maxDepth,
		divided:  false,
	}

	n.divided = true
}

// Insert adds a point to the QuadTree
func (qt *QuadTree) Insert(p Point) bool {
	return qt.root.insert(p)
}

func (n *QuadTreeNode) insert(p Point) bool {
	// Check if point is within bounds
	if !n.bounds.Contains(p) {
		return false
	}

	// If node has capacity and is not divided, add point here
	if len(n.points) < n.capacity && !n.divided {
		n.points = append(n.points, p)
		return true
	}

	// If at max depth, just add the point (exceed capacity)
	if n.depth >= n.maxDepth {
		n.points = append(n.points, p)
		return true
	}

	// If not divided yet and at capacity, subdivide
	if !n.divided {
		n.subdivide()
		
		// Redistribute existing points to children
		for _, point := range n.points {
			n.insertToChild(point)
		}
		n.points = nil // Clear points from this node
	}

	// Insert new point to appropriate child
	return n.insertToChild(p)
}

// insertToChild inserts a point into the appropriate child node
func (n *QuadTreeNode) insertToChild(p Point) bool {
	if n.northWest != nil && n.northWest.bounds.Contains(p) {
		return n.northWest.insert(p)
	}
	if n.northEast != nil && n.northEast.bounds.Contains(p) {
		return n.northEast.insert(p)
	}
	if n.southWest != nil && n.southWest.bounds.Contains(p) {
		return n.southWest.insert(p)
	}
	if n.southEast != nil && n.southEast.bounds.Contains(p) {
		return n.southEast.insert(p)
	}
	return false
}

// Query returns all points within the given bounds
func (qt *QuadTree) Query(bounds Bounds) []Point {
	points := make([]Point, 0)
	qt.root.query(bounds, &points)
	return points
}

func (n *QuadTreeNode) query(bounds Bounds, points *[]Point) {
	// If bounds don't intersect, return
	if !n.bounds.Intersects(bounds) {
		return
	}

	// Add points from this node if any
	for _, p := range n.points {
		if bounds.Contains(p) {
			*points = append(*points, p)
		}
	}

	// Query children if divided
	if n.divided {
		n.northWest.query(bounds, points)
		n.northEast.query(bounds, points)
		n.southWest.query(bounds, points)
		n.southEast.query(bounds, points)
	}
}

// GetAllPoints returns all points in the QuadTree
func (qt *QuadTree) GetAllPoints() []Point {
	points := make([]Point, 0)
	qt.root.getAllPoints(&points)
	return points
}

func (n *QuadTreeNode) getAllPoints(points *[]Point) {
	// Add points from this node
	*points = append(*points, n.points...)

	// Get points from children if divided
	if n.divided {
		n.northWest.getAllPoints(points)
		n.northEast.getAllPoints(points)
		n.southWest.getAllPoints(points)
		n.southEast.getAllPoints(points)
	}
}

// Count returns the total number of points in the QuadTree
func (qt *QuadTree) Count() int {
	return qt.root.count()
}

func (n *QuadTreeNode) count() int {
	count := len(n.points)
	if n.divided {
		count += n.northWest.count()
		count += n.northEast.count()
		count += n.southWest.count()
		count += n.southEast.count()
	}
	return count
}

// String returns a string representation of the QuadTree
func (qt *QuadTree) String() string {
	return qt.root.string("")
}

func (n *QuadTreeNode) string(indent string) string {
	result := fmt.Sprintf("%sNode at depth %d: bounds={X:%.1f, Y:%.1f, W:%.1f, H:%.1f}, points=%d\n",
		indent, n.depth, n.bounds.X, n.bounds.Y, n.bounds.Width, n.bounds.Height, len(n.points))
	
	if n.divided {
		result += n.northWest.string(indent + "  ")
		result += n.northEast.string(indent + "  ")
		result += n.southWest.string(indent + "  ")
		result += n.southEast.string(indent + "  ")
	}
	
	return result
}