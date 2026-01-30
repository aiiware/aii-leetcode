package design

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	t.Run("Basic operations", func(t *testing.T) {
		lru := Constructor(2)

		// Put (1,1)
		lru.Put(1, 1)
		
		// Put (2,2)
		lru.Put(2, 2)
		
		// Get 1 should return 1
		if val := lru.Get(1); val != 1 {
			t.Errorf("Get(1) = %d, expected 1", val)
		}
		
		// Put (3,3) should evict key 2
		lru.Put(3, 3)
		
		// Get 2 should return -1 (evicted)
		if val := lru.Get(2); val != -1 {
			t.Errorf("Get(2) = %d, expected -1 (evicted)", val)
		}
		
		// Put (4,4) should evict key 1
		lru.Put(4, 4)
		
		// Get 1 should return -1 (evicted)
		if val := lru.Get(1); val != -1 {
			t.Errorf("Get(1) = %d, expected -1 (evicted)", val)
		}
		
		// Get 3 should return 3
		if val := lru.Get(3); val != 3 {
			t.Errorf("Get(3) = %d, expected 3", val)
		}
		
		// Get 4 should return 4
		if val := lru.Get(4); val != 4 {
			t.Errorf("Get(4) = %d, expected 4", val)
		}
	})

	t.Run("Example from LeetCode", func(t *testing.T) {
		lru := Constructor(2)
		
		operations := []struct {
			op       string
			args     []int
			expected int
		}{
			{"put", []int{1, 1}, 0}, // null
			{"put", []int{2, 2}, 0}, // null
			{"get", []int{1}, 1},
			{"put", []int{3, 3}, 0}, // null
			{"get", []int{2}, -1},
			{"put", []int{4, 4}, 0}, // null
			{"get", []int{1}, -1},
			{"get", []int{3}, 3},
			{"get", []int{4}, 4},
		}
		
		for _, op := range operations {
			switch op.op {
			case "put":
				lru.Put(op.args[0], op.args[1])
			case "get":
				result := lru.Get(op.args[0])
				if result != op.expected {
					t.Errorf("Get(%d) = %d, expected %d", op.args[0], result, op.expected)
				}
			}
		}
	})

	t.Run("Update existing key", func(t *testing.T) {
		lru := Constructor(2)
		
		lru.Put(1, 1)
		lru.Put(2, 2)
		
		// Update key 1
		lru.Put(1, 10)
		
		// Get 1 should return updated value
		if val := lru.Get(1); val != 10 {
			t.Errorf("Get(1) = %d, expected 10 after update", val)
		}
		
		// Get 2 should still work
		if val := lru.Get(2); val != 2 {
			t.Errorf("Get(2) = %d, expected 2", val)
		}
	})

	t.Run("Capacity 1", func(t *testing.T) {
		lru := Constructor(1)
		
		lru.Put(1, 1)
		
		if val := lru.Get(1); val != 1 {
			t.Errorf("Get(1) = %d, expected 1", val)
		}
		
		lru.Put(2, 2)
		
		if val := lru.Get(1); val != -1 {
			t.Errorf("Get(1) = %d, expected -1 (evicted)", val)
		}
		
		if val := lru.Get(2); val != 2 {
			t.Errorf("Get(2) = %d, expected 2", val)
		}
	})

	t.Run("Large capacity", func(t *testing.T) {
		capacity := 1000
		lru := Constructor(capacity)
		
		// Fill cache
		for i := 0; i < capacity; i++ {
			lru.Put(i, i*10)
		}
		
		// All should be accessible
		for i := 0; i < capacity; i++ {
			if val := lru.Get(i); val != i*10 {
				t.Errorf("Get(%d) = %d, expected %d", i, val, i*10)
			}
		}
		
		// Add one more to evict the first one (0)
		lru.Put(capacity, capacity*10)
		
		if val := lru.Get(0); val != -1 {
			t.Errorf("Get(0) = %d, expected -1 (evicted)", val)
		}
		
		// Others should still be accessible
		for i := 1; i <= capacity; i++ {
			expected := i * 10
			if val := lru.Get(i); val != expected {
				t.Errorf("Get(%d) = %d, expected %d", i, val, expected)
			}
		}
	})

	t.Run("Get updates LRU order", func(t *testing.T) {
		lru := Constructor(3)
		
		lru.Put(1, 1)
		lru.Put(2, 2)
		lru.Put(3, 3)
		
		// Access key 1 to make it most recently used
		lru.Get(1)
		
		// Add key 4, should evict key 2 (least recently used)
		lru.Put(4, 4)
		
		// Key 2 should be evicted
		if val := lru.Get(2); val != -1 {
			t.Errorf("Get(2) = %d, expected -1 (evicted)", val)
		}
		
		// Keys 1, 3, 4 should be present
		if val := lru.Get(1); val != 1 {
			t.Errorf("Get(1) = %d, expected 1", val)
		}
		if val := lru.Get(3); val != 3 {
			t.Errorf("Get(3) = %d, expected 3", val)
		}
		if val := lru.Get(4); val != 4 {
			t.Errorf("Get(4) = %d, expected 4", val)
		}
	})
}

func TestLRUCacheSimple(t *testing.T) {
	t.Run("Basic operations", func(t *testing.T) {
		lru := NewLRUCacheSimple(2)

		lru.Put(1, 1)
		lru.Put(2, 2)
		
		if val := lru.Get(1); val != 1 {
			t.Errorf("Get(1) = %d, expected 1", val)
		}
		
		lru.Put(3, 3)
		
		if val := lru.Get(2); val != -1 {
			t.Errorf("Get(2) = %d, expected -1 (evicted)", val)
		}
		
		lru.Put(4, 4)
		
		if val := lru.Get(1); val != -1 {
			t.Errorf("Get(1) = %d, expected -1 (evicted)", val)
		}
		
		if val := lru.Get(3); val != 3 {
			t.Errorf("Get(3) = %d, expected 3", val)
		}
		
		if val := lru.Get(4); val != 4 {
			t.Errorf("Get(4) = %d, expected 4", val)
		}
	})

	t.Run("Update existing key", func(t *testing.T) {
		lru := NewLRUCacheSimple(2)
		
		lru.Put(1, 1)
		lru.Put(2, 2)
		lru.Put(1, 10)
		
		if val := lru.Get(1); val != 10 {
			t.Errorf("Get(1) = %d, expected 10 after update", val)
		}
		
		if val := lru.Get(2); val != 2 {
			t.Errorf("Get(2) = %d, expected 2", val)
		}
	})
}

func BenchmarkLRUCache(b *testing.B) {
	// Benchmark for capacity 1000
	capacity := 1000
	
	b.Run("Standard implementation", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			lru := Constructor(capacity)
			
			// Fill cache
			for j := 0; j < capacity; j++ {
				lru.Put(j, j)
			}
			
			// Access all keys
			for j := 0; j < capacity; j++ {
				lru.Get(j)
			}
			
			// Update all keys
			for j := 0; j < capacity; j++ {
				lru.Put(j, j*2)
			}
			
			// Add more to cause evictions
			for j := capacity; j < capacity*2; j++ {
				lru.Put(j, j)
			}
		}
	})
	
	b.Run("Simple implementation", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			lru := NewLRUCacheSimple(capacity)
			
			// Fill cache
			for j := 0; j < capacity; j++ {
				lru.Put(j, j)
			}
			
			// Access all keys
			for j := 0; j < capacity; j++ {
				lru.Get(j)
			}
			
			// Update all keys
			for j := 0; j < capacity; j++ {
				lru.Put(j, j*2)
			}
			
			// Add more to cause evictions
			for j := capacity; j < capacity*2; j++ {
				lru.Put(j, j)
			}
		}
	})
}

func TestLRUCacheEdgeCases(t *testing.T) {
	t.Run("Zero capacity (should handle gracefully)", func(t *testing.T) {
		// Note: LeetCode constraints say capacity >= 1, but we should handle edge case
		lru := Constructor(0)
		
		// Put should not panic
		lru.Put(1, 1)
		
		// Get should return -1
		if val := lru.Get(1); val != -1 {
			t.Errorf("Get(1) with capacity 0 = %d, expected -1", val)
		}
	})

	t.Run("Negative key or value", func(t *testing.T) {
		lru := Constructor(2)
		
		// Put negative key and value
		lru.Put(-1, -10)
		
		if val := lru.Get(-1); val != -10 {
			t.Errorf("Get(-1) = %d, expected -10", val)
		}
		
		// Put another negative
		lru.Put(-2, -20)
		
		if val := lru.Get(-2); val != -20 {
			t.Errorf("Get(-2) = %d, expected -20", val)
		}
	})

	t.Run("Same key multiple times", func(t *testing.T) {
		lru := Constructor(3)
		
		// Put same key multiple times
		for i := 0; i < 10; i++ {
			lru.Put(1, i)
		}
		
		// Should have last value
		if val := lru.Get(1); val != 9 {
			t.Errorf("Get(1) after multiple puts = %d, expected 9", val)
		}
		
		// Cache should only have one entry
		lru.Put(2, 2)
		lru.Put(3, 3)
		
		// All should be accessible
		if val := lru.Get(1); val != 9 {
			t.Errorf("Get(1) = %d, expected 9", val)
		}
		if val := lru.Get(2); val != 2 {
			t.Errorf("Get(2) = %d, expected 2", val)
		}
		if val := lru.Get(3); val != 3 {
			t.Errorf("Get(3) = %d, expected 3", val)
		}
	})

	t.Run("Consecutive gets", func(t *testing.T) {
		lru := Constructor(2)
		
		lru.Put(1, 1)
		lru.Put(2, 2)
		
		// Multiple consecutive gets
		for i := 0; i < 10; i++ {
			if val := lru.Get(1); val != 1 {
				t.Errorf("Get(1) iteration %d = %d, expected 1", i, val)
			}
		}
		
		// Add new key, should evict key 2 (not key 1)
		lru.Put(3, 3)
		
		if val := lru.Get(2); val != -1 {
			t.Errorf("Get(2) = %d, expected -1 (evicted)", val)
		}
		
		if val := lru.Get(1); val != 1 {
			t.Errorf("Get(1) = %d, expected 1", val)
		}
		
		if val := lru.Get(3); val != 3 {
			t.Errorf("Get(3) = %d, expected 3", val)
		}
	})
}