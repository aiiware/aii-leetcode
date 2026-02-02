package advanced_algorithms

// UnionFind implements the Disjoint Set Union (DSU) data structure
// with path compression and union by rank optimizations.
type UnionFind struct {
	parent []int
	rank   []int
	count  int // number of components
}

// NewUnionFind creates a new UnionFind structure for n elements.
func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
		rank:   make([]int, n),
		count:  n,
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.rank[i] = 1
	}
	return uf
}

// Find returns the root of the set containing element x.
// Uses path compression to flatten the tree.
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // path compression
	}
	return uf.parent[x]
}

// Union merges the sets containing elements x and y.
// Uses union by rank to keep the tree balanced.
func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	
	if rootX == rootY {
		return false // already in the same set
	}
	
	// Union by rank: attach smaller tree under larger tree
	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}
	
	uf.count--
	return true
}

// Connected returns true if x and y are in the same set.
func (uf *UnionFind) Connected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

// Count returns the number of disjoint sets.
func (uf *UnionFind) Count() int {
	return uf.count
}

// GetComponents returns a map from root to all elements in that component.
func (uf *UnionFind) GetComponents() map[int][]int {
	components := make(map[int][]int)
	for i := 0; i < len(uf.parent); i++ {
		root := uf.Find(i)
		components[root] = append(components[root], i)
	}
	return components
}

// Reset resets the UnionFind structure to its initial state.
func (uf *UnionFind) Reset() {
	for i := 0; i < len(uf.parent); i++ {
		uf.parent[i] = i
		uf.rank[i] = 1
	}
	uf.count = len(uf.parent)
}

// Example usage:
// uf := NewUnionFind(10)
// uf.Union(0, 1)
// uf.Union(1, 2)
// uf.Connected(0, 2) // true
// uf.Count() // 8 (10 total - 2 unions = 8 components)