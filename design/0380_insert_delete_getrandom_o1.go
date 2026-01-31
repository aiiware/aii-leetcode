package design

import (
	"math/rand"
)

// RandomizedSet implements the RandomizedSet data structure
// that supports insert, remove, and getRandom in average O(1) time
type RandomizedSet struct {
	// Map from value to index in the slice
	valueToIndex map[int]int
	// Slice storing all values
	values []int
}

// Constructor initializes your data structure
func Constructor() RandomizedSet {
	return RandomizedSet{
		valueToIndex: make(map[int]int),
		values:       []int{},
	}
}

// Insert inserts a value to the set. Returns true if the set did not already contain the specified element
func (rs *RandomizedSet) Insert(val int) bool {
	if _, exists := rs.valueToIndex[val]; exists {
		return false
	}
	
	// Add to slice and store index in map
	rs.valueToIndex[val] = len(rs.values)
	rs.values = append(rs.values, val)
	return true
}

// Remove removes a value from the set. Returns true if the set contained the specified element
func (rs *RandomizedSet) Remove(val int) bool {
	index, exists := rs.valueToIndex[val]
	if !exists {
		return false
	}
	
	// Get last element
	lastIndex := len(rs.values) - 1
	lastVal := rs.values[lastIndex]
	
	// Move last element to the position of the element to delete
	rs.values[index] = lastVal
	rs.valueToIndex[lastVal] = index
	
	// Remove the last element
	rs.values = rs.values[:lastIndex]
	delete(rs.valueToIndex, val)
	
	return true
}

// GetRandom returns a random element from the current set of elements
func (rs *RandomizedSet) GetRandom() int {
	if len(rs.values) == 0 {
		return -1 // Or panic, depending on requirements
	}
	randomIndex := rand.Intn(len(rs.values))
	return rs.values[randomIndex]
}

// RandomizedSetOptimized is an optimized version with pre-allocated slice
type RandomizedSetOptimized struct {
	valueToIndex map[int]int
	values       []int
	size         int // Track actual size vs capacity
}

// ConstructorOptimized creates an optimized RandomizedSet
func ConstructorOptimized() RandomizedSetOptimized {
	return RandomizedSetOptimized{
		valueToIndex: make(map[int]int),
		values:       make([]int, 0, 1000), // Pre-allocate capacity
		size:         0,
	}
}

// Insert inserts a value
func (rs *RandomizedSetOptimized) Insert(val int) bool {
	if _, exists := rs.valueToIndex[val]; exists {
		return false
	}
	
	// Ensure slice has capacity
	if rs.size >= len(rs.values) {
		rs.values = append(rs.values, val)
	} else {
		rs.values[rs.size] = val
	}
	
	rs.valueToIndex[val] = rs.size
	rs.size++
	return true
}

// Remove removes a value
func (rs *RandomizedSetOptimized) Remove(val int) bool {
	index, exists := rs.valueToIndex[val]
	if !exists {
		return false
	}
	
	// Get last element
	lastIndex := rs.size - 1
	lastVal := rs.values[lastIndex]
	
	// Swap with last element
	rs.values[index] = lastVal
	rs.valueToIndex[lastVal] = index
	
	// Remove from map and decrement size
	delete(rs.valueToIndex, val)
	rs.size--
	
	return true
}

// GetRandom returns a random element
func (rs *RandomizedSetOptimized) GetRandom() int {
	if rs.size == 0 {
		return -1
	}
	randomIndex := rand.Intn(rs.size)
	return rs.values[randomIndex]
}