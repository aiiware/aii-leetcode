# 0072. Edit Distance

**Difficulty**: Hard  
**Category**: Dynamic Programming  
**Tags**: String, DP, Edit Distance  
**LeetCode Link**: https://leetcode.com/problems/edit-distance/

## Problem Statement

Given two strings `word1` and `word2`, return the minimum number of operations required to convert `word1` to `word2`.

You have the following three operations permitted on a word:
1. **Insert** a character
2. **Delete** a character  
3. **Replace** a character

### Examples

**Example 1:**
```
Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation: 
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')
```

**Example 2:**
```
Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation:
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')
```

### Constraints
- `0 <= word1.length, word2.length <= 500`
- `word1` and `word2` consist of lowercase English letters.

## Solution Approach

### Intuition

The edit distance problem is a classic dynamic programming problem that measures the similarity between two strings. The key insight is that we can build the solution for longer strings from solutions for shorter prefixes.

Think about the last characters of both strings:
- If they match, we don't need any operation for these characters
- If they don't match, we need to perform one of three operations:
  1. **Delete** the last character of `word1`
  2. **Insert** the last character of `word2` into `word1`
  3. **Replace** the last character of `word1` with the last character of `word2`

### Dynamic Programming Formulation

Let `dp[i][j]` represent the minimum edit distance between the first `i` characters of `word1` and the first `j` characters of `word2`.

**Base Cases:**
- `dp[0][j] = j`: Converting empty string to `word2[0:j]` requires `j` insertions
- `dp[i][0] = i`: Converting `word1[0:i]` to empty string requires `i` deletions

**Recurrence Relation:**
```
if word1[i-1] == word2[j-1]:
    dp[i][j] = dp[i-1][j-1]  # Characters match, no operation needed
else:
    dp[i][j] = 1 + min(
        dp[i-1][j],    # Delete operation
        dp[i][j-1],    # Insert operation  
        dp[i-1][j-1]   # Replace operation
    )
```

### Space Optimization

The naive implementation uses O(m×n) space. However, we can optimize to O(min(m,n)) space by noticing that we only need the previous row to compute the current row.

## Implementation

```go
package dp

import "leetcode/utils"

// MinDistance solves LeetCode problem 0072: Edit Distance
// Time complexity: O(m*n), Space complexity: O(min(m,n))
func MinDistance(word1 string, word2 string) int {
    m, n := len(word1), len(word2)
    
    // Use the shorter string for space optimization
    if m < n {
        return MinDistance(word2, word1)
    }
    
    // Create DP array with size n+1
    dp := make([]int, n+1)
    
    // Initialize first row (empty word1 to word2)
    for j := 0; j <= n; j++ {
        dp[j] = j // j insertions
    }
    
    // Fill DP table
    for i := 1; i <= m; i++ {
        prev := dp[0] // dp[i-1][0]
        dp[0] = i     // i deletions
        
        for j := 1; j <= n; j++ {
            temp := dp[j] // dp[i-1][j]
            
            if word1[i-1] == word2[j-1] {
                // Characters match, no operation needed
                dp[j] = prev // dp[i][j] = dp[i-1][j-1]
            } else {
                // Take minimum of three operations:
                // 1. Delete: dp[i-1][j] + 1
                // 2. Insert: dp[i][j-1] + 1  
                // 3. Replace: dp[i-1][j-1] + 1
                dp[j] = utils.Min(temp, utils.Min(dp[j-1], prev)) + 1
            }
            
            prev = temp // Update prev for next iteration
        }
    }
    
    return dp[n]
}
```

### Key Implementation Details

1. **Space Optimization**: The implementation uses O(min(m,n)) space by always processing the shorter string as the inner dimension.
2. **Symmetry Handling**: The function swaps arguments if `word1` is shorter than `word2` to ensure we use the shorter string for space optimization.
3. **In-place Updates**: We update the DP array in place, using a `prev` variable to store `dp[i-1][j-1]`.

## Complexity Analysis

- **Time Complexity**: O(m×n) where m and n are the lengths of the two strings. We need to fill an m×n DP table.
- **Space Complexity**: O(min(m,n)) with space optimization, compared to O(m×n) for the naive approach.

## Test Cases

The solution includes comprehensive test cases covering:
- Basic examples from the problem statement
- Edge cases (empty strings, identical strings)
- Various operations (insert, delete, replace)
- Long strings (up to 500 characters)
- Property tests (symmetry, triangle inequality)

## Applications

Edit distance has numerous real-world applications:

1. **Spell Checking**: Finding the closest dictionary word to a misspelled word
2. **DNA Sequence Alignment**: In bioinformatics, for comparing genetic sequences
3. **Natural Language Processing**: For machine translation and text similarity
4. **Version Control Systems**: For comparing file differences
5. **Plagiarism Detection**: Measuring similarity between documents

## Related Problems

- [0161. One Edit Distance](https://leetcode.com/problems/one-edit-distance/) (Medium)
- [0583. Delete Operation for Two Strings](https://leetcode.com/problems/delete-operation-for-two-strings/) (Medium)
- [0712. Minimum ASCII Delete Sum for Two Strings](https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings/) (Medium)
- [1035. Uncrossed Lines](https://leetcode.com/problems/uncrossed-lines/) (Medium) - Similar to LCS

## Practice Exercises

1. Modify the solution to also return the sequence of operations.
2. Implement the naive O(m×n) space solution first, then optimize.
3. Solve the problem using recursion with memoization.
4. Extend the solution to handle different costs for insert, delete, and replace operations.

## Notes

- The edit distance metric is also known as **Levenshtein distance**
- The algorithm can be extended to handle different operation costs
- For very long strings, more advanced algorithms like the **Myers' diff algorithm** might be more efficient
- The space-optimized version is crucial for handling the maximum constraint of 500 characters (250,000 DP cells in naive approach)