// 0049 - Group Anagrams
// https://leetcode.com/problems/group-anagrams/
// Medium - Array, Hash Table, String, Sorting

package leetcode

import "sort"

// GroupAnagrams groups strings by anagram.
// Time Complexity: O(n * k log k) where n is number of strings, k is max string length
// Space Complexity: O(n * k)
func GroupAnagrams(strs []string) [][]string {
    groups := make(map[string][]string)
    
    for _, str := range strs {
        // Sort string to create key
        key := sortString(str)
        groups[key] = append(groups[key], str)
    }
    
    // Convert map to slice
    result := make([][]string, 0, len(groups))
    for _, group := range groups {
        result = append(result, group)
    }
    
    return result
}

// sortString sorts characters in a string
func sortString(s string) string {
    chars := []rune(s)
    sort.Slice(chars, func(i, j int) bool {
        return chars[i] < chars[j]
    })
    return string(chars)
}

// GroupAnagramsCount uses character count as key
func GroupAnagramsCount(strs []string) [][]string {
    groups := make(map[[26]int][]string)
    
    for _, str := range strs {
        // Count characters
        var count [26]int
        for _, ch := range str {
            count[ch-'a']++
        }
        
        groups[count] = append(groups[count], str)
    }
    
    // Convert map to slice
    result := make([][]string, 0, len(groups))
    for _, group := range groups {
        result = append(result, group)
    }
    
    return result
}

// GroupAnagramsPrime uses prime number multiplication as key
func GroupAnagramsPrime(strs []string) [][]string {
    // First 26 prime numbers
    primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
    
    groups := make(map[int][]string)
    
    for _, str := range strs {
        // Calculate product of primes
        product := 1
        for _, ch := range str {
            product *= primes[ch-'a']
        }
        
        groups[product] = append(groups[product], str)
    }
    
    // Convert map to slice
    result := make([][]string, 0, len(groups))
    for _, group := range groups {
        result = append(result, group)
    }
    
    return result
}