package main

import (
	"fmt"
	"time"

	"leetcode"
)

func main() {
	fmt.Println("=== N-Queens II (Problem 52) Demo ===")
	fmt.Println("Returns only the count of distinct solutions")
	fmt.Println()

	// Test different board sizes
	fmt.Println("Number of solutions for different board sizes:")
	fmt.Println("n | TotalNQueens | TotalNQueensDFS | SolveNQueensCount")
	fmt.Println("--|--------------|-----------------|------------------")
	
	for n := 1; n <= 9; n++ {
		// Get counts using different implementations
		count1 := leetcode.TotalNQueens(n)
		count2 := leetcode.TotalNQueensDFS(n)
		count3 := leetcode.SolveNQueensCount(n)
		
		fmt.Printf("%d | %13d | %15d | %16d\n", n, count1, count2, count3)
		
		// Verify all implementations give same result
		if count1 != count2 || count1 != count3 {
			fmt.Printf("ERROR: Mismatch for n=%d: %d vs %d vs %d\n", n, count1, count2, count3)
		}
	}

	// Performance comparison
	fmt.Println("\n=== Performance Comparison (n=9) ===")
	fmt.Println("Implementation | Time per call (ns)")
	fmt.Println("---------------|-------------------")
	
	iterations := 10000
	n := 9
	
	// Test TotalNQueens (bit manipulation)
	start := time.Now()
	for i := 0; i < iterations; i++ {
		_ = leetcode.TotalNQueens(n)
	}
	time1 := time.Since(start).Nanoseconds() / int64(iterations)
	fmt.Printf("TotalNQueens    | %17d\n", time1)
	
	// Test TotalNQueensDFS
	start = time.Now()
	for i := 0; i < iterations; i++ {
		_ = leetcode.TotalNQueensDFS(n)
	}
	time2 := time.Since(start).Nanoseconds() / int64(iterations)
	fmt.Printf("TotalNQueensDFS | %17d\n", time2)
	
	// Test SolveNQueensCount (from problem 51)
	start = time.Now()
	for i := 0; i < iterations; i++ {
		_ = leetcode.SolveNQueensCount(n)
	}
	time3 := time.Since(start).Nanoseconds() / int64(iterations)
	fmt.Printf("SolveNQueensCount | %15d\n", time3)
	
	fmt.Printf("\nBit manipulation is %.1fx faster than DFS\n", float64(time2)/float64(time1))
	
	// Show known solution counts
	fmt.Println("\n=== Known Solution Counts ===")
	knownCounts := map[int]int{
		1:  1,
		2:  0,
		3:  0,
		4:  2,
		5:  10,
		6:  4,
		7:  40,
		8:  92,
		9:  352,
		10: 724,
		11: 2680,
		12: 14200,
	}
	
	fmt.Println("n  | Solutions")
	fmt.Println("---|----------")
	for n := 1; n <= 12; n++ {
		if count, exists := knownCounts[n]; exists {
			fmt.Printf("%2d | %8d\n", n, count)
		}
	}
	
	fmt.Println("\n=== Example: n=4 (2 solutions) ===")
	fmt.Println("The 4x4 board has 2 distinct solutions:")
	fmt.Println()
	fmt.Println("Solution 1:      Solution 2:")
	fmt.Println(". Q . .          . . Q .")
	fmt.Println(". . . Q          Q . . .")
	fmt.Println("Q . . .          . . . Q")
	fmt.Println(". . Q .          . Q . .")
}