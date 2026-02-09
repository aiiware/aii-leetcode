package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSortedArray(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		// Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
		// Output: [1,2,2,3,5,6]
		nums1 := []int{1, 2, 3, 0, 0, 0}
		nums2 := []int{2, 5, 6}

		MergeSortedArray(nums1, 3, nums2, 3)

		expected := []int{1, 2, 2, 3, 5, 6}
		assert.Equal(t, expected, nums1)
	})

	t.Run("Example 2", func(t *testing.T) {
		// Input: nums1 = [1], m = 1, nums2 = [], n = 0
		// Output: [1]
		nums1 := []int{1}
		nums2 := []int{}

		MergeSortedArray(nums1, 1, nums2, 0)

		expected := []int{1}
		assert.Equal(t, expected, nums1)
	})

	t.Run("Example 3", func(t *testing.T) {
		// Input: nums1 = [0], m = 0, nums2 = [1], n = 1
		// Output: [1]
		nums1 := []int{0}
		nums2 := []int{1}

		MergeSortedArray(nums1, 0, nums2, 1)

		expected := []int{1}
		assert.Equal(t, expected, nums1)
	})

	t.Run("Empty nums1 (all zeros)", func(t *testing.T) {
		// Input: nums1 = [0,0,0], m = 0, nums2 = [1,2,3], n = 3
		// Output: [1,2,3]
		nums1 := []int{0, 0, 0}
		nums2 := []int{1, 2, 3}

		MergeSortedArray(nums1, 0, nums2, 3)

		expected := []int{1, 2, 3}
		assert.Equal(t, expected, nums1)
	})

	t.Run("nums1 with zeros", func(t *testing.T) {
		// Input: nums1 = [0,0,0,0,0], m = 0, nums2 = [1,2,3,4,5], n = 5
		// Output: [1,2,3,4,5]
		nums1 := []int{0, 0, 0, 0, 0}
		nums2 := []int{1, 2, 3, 4, 5}

		MergeSortedArray(nums1, 0, nums2, 5)

		expected := []int{1, 2, 3, 4, 5}
		assert.Equal(t, expected, nums1)
	})
}

func BenchmarkMergeSortedArray(b *testing.B) {
	// Create test data
	nums1 := make([]int, 1000)
	for i := 0; i < 500; i++ {
		nums1[i] = i * 2
	}
	nums2 := make([]int, 500)
	for i := 0; i < 500; i++ {
		nums2[i] = i*2 + 1
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		MergeSortedArray(nums1, 500, nums2, 500)
	}
}
