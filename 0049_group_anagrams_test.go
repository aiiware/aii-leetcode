package leetcode

import (
    "reflect"
    "sort"
    "testing"
)

// Helper to sort groups for comparison
func sortGroups(groups [][]string) {
    // Sort each group
    for _, group := range groups {
        sort.Strings(group)
    }
    
    // Sort groups by first element
    sort.Slice(groups, func(i, j int) bool {
        if len(groups[i]) == 0 && len(groups[j]) == 0 {
            return false
        }
        if len(groups[i]) == 0 {
            return true
        }
        if len(groups[j]) == 0 {
            return false
        }
        return groups[i][0] < groups[j][0]
    })
}

func TestGroupAnagrams(t *testing.T) {
    tests := []struct {
        strs     []string
        expected [][]string
    }{
        {
            []string{"eat", "tea", "tan", "ate", "nat", "bat"},
            [][]string{
                {"bat"},
                {"nat", "tan"},
                {"ate", "eat", "tea"},
            },
        },
        {
            []string{""},
            [][]string{{""}},
        },
        {
            []string{"a"},
            [][]string{{"a"}},
        },
        {
            []string{},
            [][]string{},
        },
        {
            []string{"abc", "bca", "cab", "def", "fed", "xyz", "zyx"},
            [][]string{
                {"def", "fed"},
                {"xyz", "zyx"},
                {"abc", "bca", "cab"},
            },
        },
        {
            []string{"listen", "silent", "enlist", "google", "gogole", "inlets"},
            [][]string{
                {"google", "gogole"},
                {"listen", "silent", "enlist", "inlets"},
            },
        },
    }
    
    for i, test := range tests {
        // Test sorting approach
        result := GroupAnagrams(test.strs)
        sortGroups(result)
        sortGroups(test.expected)
        
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("Test %d (sorting) failed: strs=%v, got %v, expected %v", 
                i, test.strs, result, test.expected)
        }
        
        // Test count approach
        result = GroupAnagramsCount(test.strs)
        sortGroups(result)
        
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("Test %d (count) failed: strs=%v, got %v, expected %v", 
                i, test.strs, result, test.expected)
        }
        
        // Test prime approach
        result = GroupAnagramsPrime(test.strs)
        sortGroups(result)
        
        if !reflect.DeepEqual(result, test.expected) {
            t.Errorf("Test %d (prime) failed: strs=%v, got %v, expected %v", 
                i, test.strs, result, test.expected)
        }
    }
}

func BenchmarkGroupAnagrams(b *testing.B) {
    strs := []string{
        "eat", "tea", "tan", "ate", "nat", "bat",
        "listen", "silent", "enlist", "google", "gogole", "inlets",
        "abc", "bca", "cab", "def", "fed", "xyz", "zyx",
    }
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        GroupAnagrams(strs)
    }
}

func BenchmarkGroupAnagramsCount(b *testing.B) {
    strs := []string{
        "eat", "tea", "tan", "ate", "nat", "bat",
        "listen", "silent", "enlist", "google", "gogole", "inlets",
        "abc", "bca", "cab", "def", "fed", "xyz", "zyx",
    }
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        GroupAnagramsCount(strs)
    }
}

func BenchmarkGroupAnagramsPrime(b *testing.B) {
    strs := []string{
        "eat", "tea", "tan", "ate", "nat", "bat",
        "listen", "silent", "enlist", "google", "gogole", "inlets",
        "abc", "bca", "cab", "def", "fed", "xyz", "zyx",
    }
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        GroupAnagramsPrime(strs)
    }
}