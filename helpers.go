package leetcode

import (
	"fmt"
	"sort"
)

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

// sortSubsets sorts a slice of subsets for comparison
func sortSubsets(subsets [][]int) {
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

// subsetsEqual compares two sets of subsets for equality
func subsetsEqual(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	
	// Sort both sets
	sortSubsets(a)
	sortSubsets(b)
	
	// Compare each subset
	for i := range a {
		if !SlicesEqual(a[i], b[i]) {
			return false
		}
	}
	return true
}

// hasDuplicateSubsets checks if a set of subsets contains duplicates
func hasDuplicateSubsets(subsets [][]int) bool {
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

// isSubset checks if slice b is a subset of slice a (all elements of b appear in a)
func isSubset(b, a []int) bool {
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