package advanced_algorithms

import (
	"testing"
)

func TestNewSegmentTree(t *testing.T) {
	// Test with empty array
	st1 := NewSegmentTree([]int{})
	if st1.Size() != 0 {
		t.Errorf("Expected size 0 for empty array, got %d", st1.Size())
	}

	// Test with single element
	st2 := NewSegmentTree([]int{5})
	if st2.Size() != 1 {
		t.Errorf("Expected size 1, got %d", st2.Size())
	}
	if st2.Query(0, 0) != 5 {
		t.Errorf("Expected query(0,0) = 5, got %d", st2.Query(0, 0))
	}

	// Test with multiple elements
	nums := []int{1, 3, 5, 7, 9, 11}
	st3 := NewSegmentTree(nums)
	if st3.Size() != len(nums) {
		t.Errorf("Expected size %d, got %d", len(nums), st3.Size())
	}
}

func TestSegmentTree_Query(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9, 11}
	st := NewSegmentTree(nums)

	// Test full range
	expectedFull := 1 + 3 + 5 + 7 + 9 + 11
	if result := st.Query(0, 5); result != expectedFull {
		t.Errorf("Query(0,5) expected %d, got %d", expectedFull, result)
	}

	// Test partial range
	expectedPartial := 3 + 5 + 7
	if result := st.Query(1, 3); result != expectedPartial {
		t.Errorf("Query(1,3) expected %d, got %d", expectedPartial, result)
	}

	// Test single element
	if result := st.Query(2, 2); result != 5 {
		t.Errorf("Query(2,2) expected 5, got %d", result)
	}

	// Test invalid ranges
	if result := st.Query(-1, 2); result != 0 {
		t.Errorf("Query(-1,2) expected 0 for invalid range, got %d", result)
	}
	if result := st.Query(2, 10); result != 0 {
		t.Errorf("Query(2,10) expected 0 for invalid range, got %d", result)
	}
	if result := st.Query(3, 2); result != 0 {
		t.Errorf("Query(3,2) expected 0 for reversed range, got %d", result)
	}
}

func TestSegmentTree_Update(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9, 11}
	st := NewSegmentTree(nums)

	// Update single element
	st.Update(2, 10) // Change 5 to 10

	// Verify update
	if result := st.Query(2, 2); result != 10 {
		t.Errorf("After Update(2,10), Query(2,2) expected 10, got %d", result)
	}

	// Verify range query reflects update
	expected := 3 + 10 + 7
	if result := st.Query(1, 3); result != expected {
		t.Errorf("After Update(2,10), Query(1,3) expected %d, got %d", expected, result)
	}

	// Test invalid update
	st.Update(-1, 100) // Should do nothing
	st.Update(10, 100) // Should do nothing
}

func TestSegmentTree_RangeUpdate(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9, 11}
	st := NewSegmentTree(nums)

	// Add 2 to elements 1 through 3
	st.RangeUpdate(1, 3, 2)

	// Verify updates
	expected := []int{1, 5, 7, 9, 9, 11}
	for i, exp := range expected {
		if result := st.Query(i, i); result != exp {
			t.Errorf("After RangeUpdate(1,3,2), Query(%d,%d) expected %d, got %d", i, i, exp, result)
		}
	}

	// Test invalid range update
	st.RangeUpdate(-1, 2, 5) // Should do nothing
	st.RangeUpdate(2, 10, 5) // Should do nothing
}

func TestSegmentTree_GetArray(t *testing.T) {
	nums := []int{1, 3, 5, 7, 9, 11}
	st := NewSegmentTree(nums)

	arr := st.GetArray()
	if len(arr) != len(nums) {
		t.Errorf("GetArray() expected length %d, got %d", len(nums), len(arr))
	}

	for i, exp := range nums {
		if arr[i] != exp {
			t.Errorf("GetArray()[%d] expected %d, got %d", i, exp, arr[i])
		}
	}

	// Update and get array again
	st.Update(2, 10)
	arr = st.GetArray()
	expected := []int{1, 3, 10, 7, 9, 11}
	for i, exp := range expected {
		if arr[i] != exp {
			t.Errorf("After update, GetArray()[%d] expected %d, got %d", i, exp, arr[i])
		}
	}
}

func TestSegmentTree_ComplexOperations(t *testing.T) {
	// Test a sequence of operations
	st := NewSegmentTree([]int{0, 0, 0, 0, 0})

	// Multiple updates
	st.Update(0, 1)
	st.Update(1, 2)
	st.Update(2, 3)
	st.Update(3, 4)
	st.Update(4, 5)

	// Verify cumulative sum
	if sum := st.Query(0, 4); sum != 15 {
		t.Errorf("After updates, Query(0,4) expected 15, got %d", sum)
	}

	// Range update
	st.RangeUpdate(1, 3, 10) // Add 10 to indices 1, 2, 3

	// Verify after range update
	expected := []int{1, 12, 13, 14, 5}
	for i, exp := range expected {
		if result := st.Query(i, i); result != exp {
			t.Errorf("After RangeUpdate, Query(%d,%d) expected %d, got %d", i, i, exp, result)
		}
	}

	// Final sum
	if sum := st.Query(0, 4); sum != 45 {
		t.Errorf("Final Query(0,4) expected 45, got %d", sum)
	}
}

func BenchmarkSegmentTree_Query(b *testing.B) {
	// Create a large segment tree
	size := 10000
	nums := make([]int, size)
	for i := range nums {
		nums[i] = i
	}
	st := NewSegmentTree(nums)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Query random ranges
		st.Query(i%size, (i+100)%size)
	}
}

func BenchmarkSegmentTree_Update(b *testing.B) {
	size := 10000
	nums := make([]int, size)
	st := NewSegmentTree(nums)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		st.Update(i%size, i)
	}
}