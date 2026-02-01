package benchmarks

import (
	"leetcode/dp"
	"testing"
)

// BenchmarkClimbStairs benchmarks the Climbing Stairs solution
func BenchmarkClimbStairs(b *testing.B) {
	n := 30 // Reasonable size for benchmarking

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dp.ClimbStairs(n)
	}
}

// BenchmarkClimbStairsSmall benchmarks with small input
func BenchmarkClimbStairsSmall(b *testing.B) {
	n := 10

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dp.ClimbStairs(n)
	}
}

// BenchmarkClimbStairsLarge benchmarks with larger input
func BenchmarkClimbStairsLarge(b *testing.B) {
	n := 1000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dp.ClimbStairs(n)
	}
}

// BenchmarkClimbStairsVeryLarge benchmarks with very large input
func BenchmarkClimbStairsVeryLarge(b *testing.B) {
	n := 10000

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dp.ClimbStairs(n)
	}
}