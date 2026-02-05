package testutils

import (
	"fmt"
	"leetcode/utils"
)

// AssertEqual compares two values and returns an error if they're not equal
func AssertEqual[T comparable](got, expected T, message string) error {
	if got != expected {
		return fmt.Errorf("%s: got %v, expected %v", message, got, expected)
	}
	return nil
}

// AssertTrue checks if a condition is true
func AssertTrue(condition bool, message string) error {
	if !condition {
		return fmt.Errorf("%s: condition is false", message)
	}
	return nil
}

// AssertFalse checks if a condition is false
func AssertFalse(condition bool, message string) error {
	if condition {
		return fmt.Errorf("%s: condition is true", message)
	}
	return nil
}

// AssertNil checks if a value is nil
func AssertNil(value interface{}, message string) error {
	if value != nil {
		return fmt.Errorf("%s: value is not nil", message)
	}
	return nil
}

// AssertNotNil checks if a value is not nil
func AssertNotNil(value interface{}, message string) error {
	if value == nil {
		return fmt.Errorf("%s: value is nil", message)
	}
	return nil
}

// CompareTrees compares two binary trees for equality
func CompareTrees(t1, t2 *utils.TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1.Val != t2.Val {
		return false
	}
	return CompareTrees(t1.Left, t2.Left) && CompareTrees(t1.Right, t2.Right)
}

// CompareLists compares two linked lists for equality
func CompareLists(l1, l2 *utils.ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

// SliceToSet converts a slice to a map for easy comparison
func SliceToSet[T comparable](slice []T) map[T]bool {
	set := make(map[T]bool)
	for _, v := range slice {
		set[v] = true
	}
	return set
}

// CompareSlicesUnordered compares two slices ignoring order
func CompareSlicesUnordered[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}
	set1 := SliceToSet(s1)
	set2 := SliceToSet(s2)
	if len(set1) != len(set2) {
		return false
	}
	for k := range set1 {
		if !set2[k] {
			return false
		}
	}
	return true
}

// CreateChainTree creates a chain tree (completely skewed) with n nodes
// starting from startVal. This is exported for use in other packages.
func CreateChainTree(n int, startVal int) *utils.TreeNode {
	if n <= 0 {
		return nil
	}
	
	root := &utils.TreeNode{Val: startVal}
	current := root
	for i := 1; i < n; i++ {
		current.Right = &utils.TreeNode{Val: startVal + i}
		current = current.Right
	}
	return root
}

// CreatePerfectBinaryTree creates a perfect binary tree of given height
// starting from startVal. This is exported for use in other packages.
func CreatePerfectBinaryTree(height int, startVal int) *utils.TreeNode {
	if height <= 0 {
		return nil
	}
	
	var build func(depth int, val *int) *utils.TreeNode
	build = func(depth int, val *int) *utils.TreeNode {
		if depth > height {
			return nil
		}
		
		node := &utils.TreeNode{Val: *val}
		*val++
		
		node.Left = build(depth+1, val)
		node.Right = build(depth+1, val)
		
		return node
	}
	
	val := startVal
	return build(1, &val)
}