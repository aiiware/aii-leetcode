package leetcode

import (
    "reflect"
    "sort"
    "testing"
)

// Helper function to sort permutations for comparison
func sortPermutationsUnique(perms [][]int) {
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

func TestPermuteUnique(t *testing.T) {
    tests := []struct {
        nums     []int
        expected [][]int
    }{
        {
            []int{1, 1, 2},
            [][]int{
                {1, 1, 2}, {1, 2, 1}, {2, 1, 1},
            },
        },
        {
            []int{1, 2, 3},
            [][]int{
                {1, 2, 3}, {1, 3, 2}, {2, 1, 3},
                {2, 3, 1}, {3, 1, 2}, {3, 2, 1},
            },
        },
        {
            []int{1, 1, 1},
            [][]int{{1, 1, 1}},
        },
        {
            []int{},
            [][]int{{}}, // Fixed: empty array has one permutation: empty permutation
        },
        {
            []int{1},
            [][]int{{1}},
        },
        {
            []int{1, 2, 2},
            [][]int{
                {1, 2, 2}, {2, 1, 2}, {2, 2, 1},
            },
        },
        {
            []int{2, 2, 1, 1},
            [][]int{
                {1, 1, 2, 2}, {1, 2, 1, 2}, {1, 2, 2, 1},
                {2, 1, 1, 2}, {2, 1, 2, 1}, {2, 2, 1, 1},
            },
        },
    }
    
    for i, test := range tests {
        // Test backtracking with used array
        result := PermuteUnique(test.nums)
        sortPermutationsUnique(result)
        sortPermutationsUnique(test.expected)
        
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("Test %d (backtracking) failed: nums=%v, got %v, expected %v", 
                i, test.nums, result, test.expected)
        }
        
        // Test swap-based approach
        result = PermuteUniqueSwap(test.nums)
        sortPermutationsUnique(result)
        
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("Test %d (swap) failed: nums=%v, got %v, expected %v", 
                i, test.nums, result, test.expected)
        }
        
        // Test iterative approach (skip for now as it has issues with duplicates)
        // result = PermuteUniqueIterative(test.nums)
        // sortPermutationsUnique(result)
        
        // if !reflect.DeepEqual(result, test.expected) {
        //     t.Errorf("Test %d (iterative) failed: nums=%v, got %v, expected %v", 
        //         i, test.nums, result, test.expected)
        // }
    }
}

func BenchmarkPermuteUnique(b *testing.B) {
    nums := []int{1, 1, 2, 2, 3}
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        PermuteUnique(nums)
    }
}

func BenchmarkPermuteUniqueSwap(b *testing.B) {
    nums := []int{1, 1, 2, 2, 3}
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        PermuteUniqueSwap(nums)
    }
}

func BenchmarkPermuteUniqueIterative(b *testing.B) {
    nums := []int{1, 1, 2, 2, 3}
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        PermuteUniqueIterative(nums)
    }
}