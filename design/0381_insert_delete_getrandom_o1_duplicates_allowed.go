package design

import (
	"math/rand"
)

// RandomizedCollection implements the RandomizedCollection data structure
// that supports insert, remove, and getRandom in average O(1) time with duplicates allowed
type RandomizedCollection struct {
	// Map from value to set of indices in the slice
	valueToIndices map[int]map[int]bool
	// Slice storing all values
	values []int
}

// ConstructorRandomizedCollection initializes your data structure
func ConstructorRandomizedCollection() RandomizedCollection {
	return RandomizedCollection{
		valueToIndices: make(map[int]map[int]bool),
		values:         []int{},
	}
}

// Insert inserts a value to the collection. Returns true if the collection did not already contain the specified element
func (rc *RandomizedCollection) Insert(val int) bool {
	// Check if value already exists
	exists := rc.valueToIndices[val] != nil && len(rc.valueToIndices[val]) > 0
	
	// Add to slice
	index := len(rc.values)
	rc.values = append(rc.values, val)
	
	// Add index to map
	if rc.valueToIndices[val] == nil {
		rc.valueToIndices[val] = make(map[int]bool)
	}
	rc.valueToIndices[val][index] = true
	
	return !exists
}

// Remove removes a value from the collection. Returns true if the collection contained the specified element
func (rc *RandomizedCollection) Remove(val int) bool {
	indices, exists := rc.valueToIndices[val]
	if !exists || len(indices) == 0 {
		return false
	}
	
	// Get any index of this value
	var indexToRemove int
	for idx := range indices {
		indexToRemove = idx
		break
	}
	
	// Get last element
	lastIndex := len(rc.values) - 1
	lastVal := rc.values[lastIndex]
	
	// If we're not removing the last element, we need to swap
	if indexToRemove != lastIndex {
		// Move last element to the position of the element to delete
		rc.values[indexToRemove] = lastVal
		
		// Update indices for the moved element
		// Remove old index
		delete(rc.valueToIndices[lastVal], lastIndex)
		// Add new index
		rc.valueToIndices[lastVal][indexToRemove] = true
	}
	
	// Remove the last element from slice
	rc.values = rc.values[:lastIndex]
	
	// Remove index from map for the deleted value
	delete(rc.valueToIndices[val], indexToRemove)
	
	// Clean up empty map entry
	if len(rc.valueToIndices[val]) == 0 {
		delete(rc.valueToIndices, val)
	}
	
	return true
}

// GetRandom returns a random element from the current collection of elements
func (rc *RandomizedCollection) GetRandom() int {
	if len(rc.values) == 0 {
		return -1 // Or panic, depending on requirements
	}
	randomIndex := rand.Intn(len(rc.values))
	return rc.values[randomIndex]
}

// GetRandomWithWeight returns a random element with probability proportional to its frequency
func (rc *RandomizedCollection) GetRandomWithWeight() int {
	if len(rc.values) == 0 {
		return -1
	}
	randomIndex := rand.Intn(len(rc.values))
	return rc.values[randomIndex]
}

// Count returns the count of a specific value in the collection
func (rc *RandomizedCollection) Count(val int) int {
	if indices, exists := rc.valueToIndices[val]; exists {
		return len(indices)
	}
	return 0
}

// Size returns the total number of elements in the collection
func (rc *RandomizedCollection) Size() int {
	return len(rc.values)
}

// RandomizedCollectionOptimized is an optimized version with pre-allocated slice
type RandomizedCollectionOptimized struct {
	valueToIndices map[int]map[int]bool
	values         []int
	size           int // Track actual size vs capacity
}

// ConstructorRandomizedCollectionOptimized creates an optimized RandomizedCollection
func ConstructorRandomizedCollectionOptimized() RandomizedCollectionOptimized {
	return RandomizedCollectionOptimized{
		valueToIndices: make(map[int]map[int]bool),
		values:         make([]int, 0, 1000), // Pre-allocate capacity
		size:           0,
	}
}

// Insert inserts a value
func (rc *RandomizedCollectionOptimized) Insert(val int) bool {
	// Check if value already exists
	exists := rc.valueToIndices[val] != nil && len(rc.valueToIndices[val]) > 0
	
	// Ensure slice has capacity
	index := rc.size
	if rc.size >= len(rc.values) {
		rc.values = append(rc.values, val)
	} else {
		rc.values[index] = val
	}
	
	// Add index to map
	if rc.valueToIndices[val] == nil {
		rc.valueToIndices[val] = make(map[int]bool)
	}
	rc.valueToIndices[val][index] = true
	rc.size++
	
	return !exists
}

// Remove removes a value
func (rc *RandomizedCollectionOptimized) Remove(val int) bool {
	indices, exists := rc.valueToIndices[val]
	if !exists || len(indices) == 0 {
		return false
	}
	
	// Get any index of this value
	var indexToRemove int
	for idx := range indices {
		indexToRemove = idx
		break
	}
	
	// Get last element
	lastIndex := rc.size - 1
	lastVal := rc.values[lastIndex]
	
	// If we're not removing the last element, we need to swap
	if indexToRemove != lastIndex {
		// Move last element to the position of the element to delete
		rc.values[indexToRemove] = lastVal
		
		// Update indices for the moved element
		// Remove old index
		delete(rc.valueToIndices[lastVal], lastIndex)
		// Add new index
		rc.valueToIndices[lastVal][indexToRemove] = true
	}
	
	// Remove index from map for the deleted value
	delete(rc.valueToIndices[val], indexToRemove)
	
	// Clean up empty map entry
	if len(rc.valueToIndices[val]) == 0 {
		delete(rc.valueToIndices, val)
	}
	
	rc.size--
	return true
}

// GetRandom returns a random element
func (rc *RandomizedCollectionOptimized) GetRandom() int {
	if rc.size == 0 {
		return -1
	}
	randomIndex := rand.Intn(rc.size)
	return rc.values[randomIndex]
}