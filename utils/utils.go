package utils

import (
	"fmt"
	"sort"
)

// ListNode represents a node in a singly-linked list.
// This is the standard definition used in LeetCode problems.
type ListNode struct {
	Val  int
	Next *ListNode
}

// NewListNode creates a new ListNode with the given value.
func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

// NewListFromSlice creates a linked list from a slice of integers.
// Returns the head of the linked list.
func NewListFromSlice(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{Val: vals[0]}
	current := head
	for i := 1; i < len(vals); i++ {
		current.Next = &ListNode{Val: vals[i]}
		current = current.Next
	}
	return head
}

// ToSlice converts a linked list to a slice of integers.
func (l *ListNode) ToSlice() []int {
	result := []int{} // Return empty slice, not nil
	current := l
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}

// Equal compares two linked lists for equality.
func (l *ListNode) Equal(other *ListNode) bool {
	a, b := l, other
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return a == nil && b == nil
}

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeFromSlice creates a binary tree from a level-order slice representation.
// nil values in the slice represent missing nodes.
// This handles the compact LeetCode representation where nil nodes don't have their children represented.
func NewTreeFromSlice(vals []*int) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *vals[0]}
	
	// Use a queue to process nodes in level order
	queue := []*TreeNode{root}
	
	// We'll use i to track our position in the vals array
	i := 1
	
	for len(queue) > 0 && i < len(vals) {
		// Get the next node to process
		node := queue[0]
		queue = queue[1:]
		
		// Process left child
		if i < len(vals) {
			if vals[i] != nil {
				node.Left = &TreeNode{Val: *vals[i]}
				queue = append(queue, node.Left)
			}
			// If vals[i] is nil, we don't add anything to the queue
			i++
		}
		
		// Process right child
		if i < len(vals) {
			if vals[i] != nil {
				node.Right = &TreeNode{Val: *vals[i]}
				queue = append(queue, node.Right)
			}
			// If vals[i] is nil, we don't add anything to the queue
			i++
		}
	}

	return root
}

// ToSlice converts a binary tree to a level-order slice representation.
func (t *TreeNode) ToSlice() []*int {
	if t == nil {
		return []*int{}
	}

	var result []*int
	queue := []*TreeNode{t}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, nil)
			continue
		}

		result = append(result, &node.Val)
		queue = append(queue, node.Left, node.Right)
	}

	// Remove trailing nil values
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	return result
}

// Equal compares two trees for equality
func (t *TreeNode) Equal(other *TreeNode) bool {
	if t == nil && other == nil {
		return true
	}
	if t == nil || other == nil {
		return false
	}
	if t.Val != other.Val {
		return false
	}
	return t.Left.Equal(other.Left) && t.Right.Equal(other.Right)
}

// Helper functions for testing

// CloneTree creates a deep copy of a binary tree
func CloneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	return &TreeNode{
		Val:   root.Val,
		Left:  CloneTree(root.Left),
		Right: CloneTree(root.Right),
	}
}

// CreateCompleteTree creates a complete binary tree with n nodes
func CreateCompleteTree(n int) *TreeNode {
	if n <= 0 {
		return nil
	}
	
	// Create nodes
	nodes := make([]*TreeNode, n)
	for i := 0; i < n; i++ {
		nodes[i] = &TreeNode{Val: i + 1}
	}
	
	// Connect parent-child relationships
	for i := 0; i < n; i++ {
		leftIdx := 2*i + 1
		if leftIdx < n {
			nodes[i].Left = nodes[leftIdx]
		}
		
		rightIdx := 2*i + 2
		if rightIdx < n {
			nodes[i].Right = nodes[rightIdx]
		}
	}
	
	return nodes[0]
}

// CreateRightSkewedTree creates a right-skewed tree with n nodes
func CreateRightSkewedTree(n int) *TreeNode {
	if n <= 0 {
		return nil
	}
	
	root := &TreeNode{Val: 1}
	current := root
	
	for i := 2; i <= n; i++ {
		current.Right = &TreeNode{Val: i}
		current = current.Right
	}
	
	return root
}

// CreateLeftSkewedTree creates a left-skewed tree with n nodes
func CreateLeftSkewedTree(n int) *TreeNode {
	if n <= 0 {
		return nil
	}
	
	root := &TreeNode{Val: 1}
	current := root
	
	for i := 2; i <= n; i++ {
		current.Left = &TreeNode{Val: i}
		current = current.Left
	}
	
	return root
}

// CreateSymmetricTree creates a symmetric binary tree with specified number of levels
// levels=1: single node, levels=2: 3 nodes, levels=3: 7 nodes, etc.
// The tree will have symmetric values: for 3 levels: [1, 2, 2, 3, 4, 4, 3]
func CreateSymmetricTree(levels int) *TreeNode {
	if levels <= 0 {
		return nil
	}
	
	// Actually, let's use a simpler approach: build the tree level by level
	// with symmetric values
	
	// We'll use a queue to build the tree level by level
	if levels == 1 {
		return &TreeNode{Val: 1}
	}
	
	// Create root
	root := &TreeNode{Val: 1}
	
	// We'll use a slice to track nodes at each level
	currentLevel := []*TreeNode{root}
	nextVal := 2  // Next value to assign
	
	for level := 1; level < levels; level++ {
		nextLevel := make([]*TreeNode, 0, 1<<level)
		
		// For a symmetric tree, values at this level should be symmetric
		// Number of nodes at this level: 2^level
		nodesThisLevel := 1 << level
		
		// Create values for this level
		// First half: increasing values
		// Second half: mirror of first half
		half := nodesThisLevel / 2
		values := make([]int, nodesThisLevel)
		
		for i := 0; i < half; i++ {
			values[i] = nextVal
			nextVal++
		}
		
		// Mirror the values for the second half
		for i := 0; i < half; i++ {
			values[nodesThisLevel-1-i] = values[i]
		}
		
		// Create nodes for this level
		for i := 0; i < nodesThisLevel; i++ {
			nextLevel = append(nextLevel, &TreeNode{Val: values[i]})
		}
		
		// Connect parent-child relationships
		for i := 0; i < len(currentLevel); i++ {
			parent := currentLevel[i]
			leftIdx := 2 * i
			rightIdx := 2*i + 1
			
			if leftIdx < len(nextLevel) {
				parent.Left = nextLevel[leftIdx]
			}
			if rightIdx < len(nextLevel) {
				parent.Right = nextLevel[rightIdx]
			}
		}
		
		// Move to next level
		currentLevel = nextLevel
	}
	
	return root
}

// Basic utility functions

// Min returns the minimum of two integers
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max returns the maximum of two integers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Abs returns the absolute value of an integer
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// MinFloat64 returns the minimum of two float64 values
func MinFloat64(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// MaxFloat64 returns the maximum of two float64 values
func MaxFloat64(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// MinOf returns the minimum value among the given integers
func MinOf(values ...int) int {
	if len(values) == 0 {
		return 0
	}
	min := values[0]
	for _, v := range values[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

// MaxOf returns the maximum value among the given integers
func MaxOf(values ...int) int {
	if len(values) == 0 {
		return 0
	}
	max := values[0]
	for _, v := range values[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// Helper functions for LeetCode solutions and tests

// SlicesEqual compares two integer slices for equality
func SlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// CountBits counts the number of 1 bits in an integer
func CountBits(x int) int {
	count := 0
	for x > 0 {
		count += x & 1
		x >>= 1
	}
	return count
}

// IsValidGrayCode checks if a sequence is a valid n-bit Gray code
func IsValidGrayCode(seq []int, n int) bool {
	size := 1 << n

	// Check length
	if len(seq) != size {
		return false
	}

	// Check range and uniqueness
	seen := make([]bool, size)
	for i, num := range seq {
		if num < 0 || num >= size {
			return false
		}
		if seen[num] {
			return false
		}
		seen[num] = true

		// Check adjacent difference (except for last element)
		if i > 0 {
			diff := seq[i-1] ^ num
			if CountBits(diff) != 1 {
				return false
			}
		}
	}

	// Check first and last difference
	firstLastDiff := seq[0] ^ seq[size-1]
	return CountBits(firstLastDiff) == 1
}

// IsPermutation checks if a sequence is a permutation of 0..n-1
func IsPermutation(seq []int, n int) bool {
	if len(seq) != n {
		return false
	}

	seen := make([]bool, n)
	for _, num := range seq {
		if num < 0 || num >= n {
			return false
		}
		if seen[num] {
			return false
		}
		seen[num] = true
	}

	// All numbers should be seen
	for i := 0; i < n; i++ {
		if !seen[i] {
			return false
		}
	}
	return true
}

// IntsEqual compares two integer slices for equality (alias for SlicesEqual)
func IntsEqual(a, b []int) bool {
	return SlicesEqual(a, b)
}

// StringsEqual compares two string slices for equality
func StringsEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// BoolsEqual compares two boolean slices for equality
func BoolsEqual(a, b []bool) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// MatrixEqual compares two 2D integer slices for equality
func MatrixEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !SlicesEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}

// MakeRange creates a slice of integers from start to end (inclusive)
func MakeRange(start, end int) []int {
	if start > end {
		return []int{}
	}
	result := make([]int, end-start+1)
	for i := range result {
		result[i] = start + i
	}
	return result
}

// Repeat creates a string by repeating a character n times
func Repeat(char string, n int) string {
	if n <= 0 {
		return ""
	}
	result := ""
	for i := 0; i < n; i++ {
		result += char
	}
	return result
}

// IntPtr returns a pointer to an integer
func IntPtr(x int) *int {
	return &x
}

// StringSlicesEqual compares two string slices for equality (alias for StringsEqual)
func StringSlicesEqual(a, b []string) bool {
	return StringsEqual(a, b)
}

// IntsToList creates a linked list from a slice of integers
func IntsToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	
	head := &ListNode{Val: nums[0]}
	current := head
	
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}
	
	return head
}

// CopyList creates a deep copy of a linked list
func CopyList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	
	// Create a map to store mapping from original nodes to new nodes
	nodeMap := make(map[*ListNode]*ListNode)
	
	// First pass: create all new nodes
	current := head
	for current != nil {
		nodeMap[current] = &ListNode{Val: current.Val}
		current = current.Next
	}
	
	// Second pass: set up next pointers
	current = head
	for current != nil {
		if current.Next != nil {
			nodeMap[current].Next = nodeMap[current.Next]
		}
		current = current.Next
	}
	
	return nodeMap[head]
}

// PrintList prints a linked list in a readable format
func PrintList(head *ListNode) string {
	if head == nil {
		return "nil"
	}
	
	result := ""
	current := head
	for current != nil {
		result += fmt.Sprintf("%d", current.Val)
		if current.Next != nil {
			result += " -> "
		}
		current = current.Next
	}
	return result
}

// Helper functions for subsets testing

// SortSubsets sorts a slice of subsets for comparison
func SortSubsets(subsets [][]int) {
	// Sort each subset
	for _, subset := range subsets {
		sort.Ints(subset)
	}
	
	// Sort the list of subsets
	sort.Slice(subsets, func(i, j int) bool {
		// First by length
		if len(subsets[i]) != len(subsets[j]) {
			return len(subsets[i]) < len(subsets[j])
		}
		// Then lexicographically
		for k := 0; k < len(subsets[i]); k++ {
			if subsets[i][k] != subsets[j][k] {
				return subsets[i][k] < subsets[j][k]
			}
		}
		return false
	})
}

// SubsetsEqual compares two sets of subsets for equality
func SubsetsEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	
	// Sort both sets
	SortSubsets(a)
	SortSubsets(b)
	
	// Compare each subset
	for i := range a {
		if !SlicesEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}

// HasDuplicateSubsets checks if a set of subsets contains duplicates
func HasDuplicateSubsets(subsets [][]int) bool {
	seen := make(map[string]bool)
	
	for _, subset := range subsets {
		// Sort the subset for consistent key
		sorted := make([]int, len(subset))
		copy(sorted, subset)
		sort.Ints(sorted)
		
		// Create a string key
		key := ""
		for _, num := range sorted {
			key += string(rune(num + '0')) + ","
		}
		
		if seen[key] {
			return true
		}
		seen[key] = true
	}
	return false
}

// IsSubset checks if slice b is a subset of slice a (all elements of b appear in a)
func IsSubset(b, a []int) bool {
	// Count occurrences in a
	countA := make(map[int]int)
	for _, num := range a {
		countA[num]++
	}
	
	// Check if all elements of b are in a with sufficient count
	for _, num := range b {
		if countA[num] == 0 {
			return false
		}
		countA[num]--
	}
	return true
}