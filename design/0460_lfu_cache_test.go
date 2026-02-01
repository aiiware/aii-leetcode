package design

import (
	"testing"
)

func TestLFUCache(t *testing.T) {
	t.Run("Basic operations", func(t *testing.T) {
		capacity := 2
		cache := ConstructorLFU(capacity)
		
		// Test put and get
		cache.Put(1, 1)
		cache.Put(2, 2)
		
		if val := cache.Get(1); val != 1 {
			t.Errorf("Get(1) = %v, want 1", val)
		}
		
		// Test update frequency
		if val := cache.Get(2); val != 2 {
			t.Errorf("Get(2) = %v, want 2", val)
		}
		
		// Test eviction (key 1 should be evicted since it's less frequently used)
		cache.Put(3, 3) // This should evict key 1
		
		if val := cache.Get(1); val != -1 {
			t.Errorf("Get(1) after eviction = %v, want -1", val)
		}
		if val := cache.Get(3); val != 3 {
			t.Errorf("Get(3) = %v, want 3", val)
		}
		if val := cache.Get(2); val != 2 {
			t.Errorf("Get(2) = %v, want 2", val)
		}
	})
	
	t.Run("LFU with same frequency uses LRU", func(t *testing.T) {
		capacity := 2
		cache := ConstructorLFU(capacity)
		
		cache.Put(1, 1)
		cache.Put(2, 2)
		
		// Both keys have frequency 1
		// Access key 1 to make it more recently used
		cache.Get(1)
		
		// Add new key, should evict key 2 (least recently used among freq 1)
		cache.Put(3, 3)
		
		if val := cache.Get(2); val != -1 {
			t.Errorf("Get(2) = %v, want -1 (should be evicted)", val)
		}
		if val := cache.Get(1); val != 1 {
			t.Errorf("Get(1) = %v, want 1", val)
		}
		if val := cache.Get(3); val != 3 {
			t.Errorf("Get(3) = %v, want 3", val)
		}
	})
	
	t.Run("Update existing key", func(t *testing.T) {
		capacity := 2
		cache := ConstructorLFU(capacity)
		
		cache.Put(1, 1)
		cache.Put(2, 2)
		cache.Put(1, 10) // Update value and increase frequency
		
		if val := cache.Get(1); val != 10 {
			t.Errorf("Get(1) after update = %v, want 10", val)
		}
		
		// Add new key, should evict key 2 (frequency 1) not key 1 (frequency 2)
		cache.Put(3, 3)
		
		if val := cache.Get(2); val != -1 {
			t.Errorf("Get(2) = %v, want -1 (should be evicted)", val)
		}
		if val := cache.Get(1); val != 10 {
			t.Errorf("Get(1) = %v, want 10", val)
		}
		if val := cache.Get(3); val != 3 {
			t.Errorf("Get(3) = %v, want 3", val)
		}
	})
	
	t.Run("Zero capacity", func(t *testing.T) {
		cache := ConstructorLFU(0)
		
		cache.Put(1, 1)
		if val := cache.Get(1); val != -1 {
			t.Errorf("Get(1) with zero capacity = %v, want -1", val)
		}
	})
	
	t.Run("Complex frequency updates", func(t *testing.T) {
		capacity := 3
		cache := ConstructorLFU(capacity)
		
		// Setup: 1:1, 2:1, 3:1
		cache.Put(1, 1)
		cache.Put(2, 2)
		cache.Put(3, 3)
		
		// Access pattern: 1 (freq 2), 2 (freq 2), 3 (freq 1)
		cache.Get(1)
		cache.Get(2)
		
		// Add new key, should evict key 3 (lowest frequency)
		cache.Put(4, 4)
		
		if val := cache.Get(3); val != -1 {
			t.Errorf("Get(3) = %v, want -1 (should be evicted)", val)
		}
		if val := cache.Get(1); val != 1 {
			t.Errorf("Get(1) = %v, want 1", val)
		}
		if val := cache.Get(2); val != 2 {
			t.Errorf("Get(2) = %v, want 2", val)
		}
		if val := cache.Get(4); val != 4 {
			t.Errorf("Get(4) = %v, want 4", val)
		}
	})
}