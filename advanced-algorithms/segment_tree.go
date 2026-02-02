package advanced_algorithms

// SegmentTree implements a segment tree for range sum queries and updates.
// It supports O(log n) range queries and point updates.
type SegmentTree struct {
	n    int   // size of the original array
	tree []int // segment tree stored as array
}

// NewSegmentTree creates a new segment tree from the given array.
func NewSegmentTree(nums []int) *SegmentTree {
	n := len(nums)
	tree := make([]int, 4*n) // allocate 4*n space for the tree
	st := &SegmentTree{n: n, tree: tree}
	if n > 0 {
		st.build(nums, 1, 0, n-1)
	}
	return st
}

// build recursively builds the segment tree.
// node: current node index in tree array
// left, right: range of original array covered by this node
func (st *SegmentTree) build(nums []int, node, left, right int) {
	if left == right {
		st.tree[node] = nums[left]
		return
	}
	
	mid := left + (right-left)/2
	st.build(nums, node*2, left, mid)      // left child
	st.build(nums, node*2+1, mid+1, right) // right child
	
	// Combine results from children (sum in this case)
	st.tree[node] = st.tree[node*2] + st.tree[node*2+1]
}

// Query returns the sum of elements in range [l, r] (0-indexed).
func (st *SegmentTree) Query(l, r int) int {
	if l < 0 || r >= st.n || l > r {
		return 0
	}
	return st.query(1, 0, st.n-1, l, r)
}

// query recursively computes the sum for range [l, r].
func (st *SegmentTree) query(node, left, right, l, r int) int {
	// If current segment is completely outside query range
	if right < l || left > r {
		return 0
	}
	
	// If current segment is completely inside query range
	if left >= l && right <= r {
		return st.tree[node]
	}
	
	// Partial overlap, query both children
	mid := left + (right-left)/2
	leftSum := st.query(node*2, left, mid, l, r)
	rightSum := st.query(node*2+1, mid+1, right, l, r)
	
	return leftSum + rightSum
}

// Update updates the value at index idx to val (0-indexed).
func (st *SegmentTree) Update(idx, val int) {
	if idx < 0 || idx >= st.n {
		return
	}
	st.update(1, 0, st.n-1, idx, val)
}

// update recursively updates the tree.
func (st *SegmentTree) update(node, left, right, idx, val int) {
	if left == right {
		st.tree[node] = val
		return
	}
	
	mid := left + (right-left)/2
	if idx <= mid {
		st.update(node*2, left, mid, idx, val)
	} else {
		st.update(node*2+1, mid+1, right, idx, val)
	}
	
	// Update parent with new sum
	st.tree[node] = st.tree[node*2] + st.tree[node*2+1]
}

// RangeUpdate updates all elements in range [l, r] by adding delta.
// This is a naive implementation; for efficient range updates,
// consider implementing lazy propagation.
func (st *SegmentTree) RangeUpdate(l, r, delta int) {
	if l < 0 || r >= st.n || l > r {
		return
	}
	for i := l; i <= r; i++ {
		// Get current value and update
		// Note: This is O(n log n) - for O(log n) range updates,
		// implement lazy propagation
		current := st.Query(i, i)
		st.Update(i, current+delta)
	}
}

// GetArray returns the current array representation.
func (st *SegmentTree) GetArray() []int {
	result := make([]int, st.n)
	for i := 0; i < st.n; i++ {
		result[i] = st.Query(i, i)
	}
	return result
}

// Size returns the size of the original array.
func (st *SegmentTree) Size() int {
	return st.n
}

// Example usage:
// nums := []int{1, 3, 5, 7, 9, 11}
// st := NewSegmentTree(nums)
// st.Query(1, 3) // 3 + 5 + 7 = 15
// st.Update(2, 10) // update index 2 to 10
// st.Query(1, 3) // 3 + 10 + 7 = 20