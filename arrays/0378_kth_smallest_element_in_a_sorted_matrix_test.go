package arrays

import "testing"

func TestKthSmallest(t *testing.T) {
	// Test cases for kthSmallest function
	t.Run("Example 1", func(t *testing.T) {
		matrix := [][]int{
			{1, 5, 9},
			{10, 11, 13},
			{12, 13, 15},
		}
		k := 8
		result := kthSmallest(matrix, k)
		// Expected: 13
		if result != 13 {
			t.Errorf("Expected 13, got %d", result)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		matrix := [][]int{{-5}}
		k := 1
		result := kthSmallest(matrix, k)
		// Expected: -5
		if result != -5 {
			t.Errorf("Expected -5, got %d", result)
		}
	})
}

func TestKthSmallestBinarySearch(t *testing.T) {
	// Test cases for kthSmallestBinarySearch function
	t.Run("Example 1", func(t *testing.T) {
		matrix := [][]int{
			{1, 5, 9},
			{10, 11, 13},
			{12, 13, 15},
		}
		k := 8
		result := kthSmallestBinarySearch(matrix, k)
		// Expected: 13
		if result != 13 {
			t.Errorf("Expected 13, got %d", result)
		}
	})

	t.Run("Example 2", func(t *testing.T) {
		matrix := [][]int{{-5}}
		k := 1
		result := kthSmallestBinarySearch(matrix, k)
		// Expected: -5
		if result != -5 {
			t.Errorf("Expected -5, got %d", result)
		}
	})
}

func TestItem(t *testing.T) {
	// Test Item struct for kthSmallest
	item := Item{
		value: 10,
		row:   0,
		col:   0,
	}
	if item.value != 10 {
		t.Errorf("Expected value 10, got %d", item.value)
	}
	if item.row != 0 {
		t.Errorf("Expected row 0, got %d", item.row)
	}
	if item.col != 0 {
		t.Errorf("Expected col 0, got %d", item.col)
	}
}
