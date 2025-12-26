package leetcode

import (
    "reflect"
    "sort"
    "testing"
)

// Helper function to sort permutations for comparison
func sortPermutations(perms [][]int) {
    // Sort each permutation
    for _, perm := range perms {
        sort.Ints(perm)
    }
    
    // Sort the list of permutations
    sort.Slice(perms, func(i, j int) bool {
        for k := 0; k < len(perms[i]) && k < len(perms[j]); k++ {
            if perms[i][k] != perms[j][k] {
                return perms[i][k] < perms[j][k]
            }
        }
        return len(perms[i]) < len(perms[j])
    })
}

func TestPermute(t *testing.T) {
    tests := []struct {
        nums     []int
        expected [][]int
    }{
        {
            []int{1, 2, 3},
            [][]int{
                {1, 2, 3}, {1, 3, 2}, {2, 1, 3},
                {2, 3, 1}, {3, 1, 2}, {3, 2, 1},
            },
        },
        {
            []int{0, 1},
            [][]int{{0, 1}, {1, 0}},
        },
        {
            []int{1},
            [][]int{{1}},
        },
        {
            []int{},
            [][]int{{}}, // Fixed: empty array has one permutation: empty permutation
        },
        {
            []int{1, 2},
            [][]int{{1, 2}, {2, 1}},
        },
    }
    
    for i, test := range tests {
        // Test backtracking approach
        result := Permute(test.nums)
        sortPermutations(result)
        sortPermutations(test.expected)
        
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("Test %d (backtracking) failed: nums=%v, got %v, expected %v", 
                i, test.nums, result, test.expected)
        }
        
        // Test Heap's algorithm (skip for empty array to avoid stack overflow)
        if len(test.nums) > 0 {
            result = PermuteHeap(test.nums)
            sortPermutations(result)
            
            if !reflect.DeepEqual(result, test.expected) {
                t.Errorf("Test %d (Heap) failed: nums=%v, got %v, expected %v", 
                    i, test.nums, result, test.expected)
            }
        }
        
        // Test iterative approach
        result = PermuteIterative(test.nums)
        sortPermutations(result)
        
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("Test %d (iterative) failed: nums=%v, got %v, expected %v", 
                i, test.nums, result, test.expected)
        }
    }
}

func BenchmarkPermute(b *testing.B) {
    nums := []int{1, 2, 3, 4, 5}
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        Permute(nums)
    }
}

func BenchmarkPermuteHeap(b *testing.B) {
    nums := []int{1, 2, 3, 4, 5}
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        PermuteHeap(nums)
    }
}

func BenchmarkPermuteIterative(b *testing.B) {
    nums := []int{1, 2, 3, 4, 5}
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        PermuteIterative(nums)
    }
}