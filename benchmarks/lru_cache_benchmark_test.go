package benchmarks

import (
	"leetcode/design"
	"testing"
)

// BenchmarkLRUCache benchmarks the LRU Cache implementation
func BenchmarkLRUCache(b *testing.B) {
	capacity := 1000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache := design.ConstructorLRUCache(capacity)
		
		// Perform operations
		for j := 0; j < capacity; j++ {
			cache.Put(j, j*10)
		}
		
		for j := 0; j < capacity; j++ {
			cache.Get(j)
		}
		
		// Evict some items
		for j := capacity; j < capacity*2; j++ {
			cache.Put(j, j*10)
		}
	}
}

// BenchmarkLRUCacheSmall benchmarks with small cache
func BenchmarkLRUCacheSmall(b *testing.B) {
	capacity := 10

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache := design.ConstructorLRUCache(capacity)
		
		// Mix of operations
		for j := 0; j < 100; j++ {
			cache.Put(j%capacity, j)
			if j%3 == 0 {
				cache.Get(j % capacity)
			}
		}
	}
}

// BenchmarkLRUCacheGet benchmarks Get operations
func BenchmarkLRUCacheGet(b *testing.B) {
	capacity := 100
	cache := design.ConstructorLRUCache(capacity)
	
	// Fill cache
	for i := 0; i < capacity; i++ {
		cache.Put(i, i*10)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Get(i % capacity)
	}
}

// BenchmarkLRUCachePut benchmarks Put operations
func BenchmarkLRUCachePut(b *testing.B) {
	capacity := 100

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache := design.ConstructorLRUCache(capacity)
		cache.Put(i%capacity, i)
	}
}