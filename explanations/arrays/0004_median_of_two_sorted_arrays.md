# 4. Median of Two Sorted Arrays - Solution Explanation

## Problem Statement
Given two sorted arrays `nums1` and `nums2` of size `m` and `n` respectively, return the median of the two sorted arrays.

The overall run time complexity should be **O(log(min(m, n)))**.

## Difficulty: Hard

## Key Insights
1. **Median Definition**: For combined array of size `(m+n)`:
   - If `(m+n)` is odd: median = middle element
   - If `(m+n)` is even: median = average of two middle elements

2. **Binary Search Approach**: 
   - We can partition both arrays such that:
     - All elements on left side ≤ all elements on right side
     - Number of elements on left = number of elements on right (or left has one more)
   - This partition gives us the median position

3. **Complexity Requirement**: O(log(min(m, n))) suggests binary search on the smaller array

## Solution Approaches

### Approach 1: Binary Search on Partition (Optimal)
**Time Complexity**: O(log(min(m, n)))
**Space Complexity**: O(1)

```go
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    // Ensure nums1 is the smaller array
    if len(nums1) > len(nums2) {
        nums1, nums2 = nums2, nums1
    }
    
    m, n := len(nums1), len(nums2)
    total := m + n
    half := (total + 1) / 2  // Number of elements in left partition
    
    // Binary search on nums1
    left, right := 0, m
    for left <= right {
        // Partition nums1
        i := (left + right) / 2  // Number of elements from nums1 in left partition
        j := half - i            // Number of elements from nums2 in left partition
        
        // Get boundary elements
        nums1Left := math.MinInt64
        if i > 0 {
            nums1Left = nums1[i-1]
        }
        nums1Right := math.MaxInt64
        if i < m {
            nums1Right = nums1[i]
        }
        
        nums2Left := math.MinInt64
        if j > 0 {
            nums2Left = nums2[j-1]
        }
        nums2Right := math.MaxInt64
        if j < n {
            nums2Right = nums2[j]
        }
        
        // Check if partition is correct
        if nums1Left <= nums2Right && nums2Left <= nums1Right {
            // Found correct partition
            if total%2 == 1 {
                // Odd total length
                return float64(max(nums1Left, nums2Left))
            } else {
                // Even total length
                return float64(max(nums1Left, nums2Left)+min(nums1Right, nums2Right)) / 2.0
            }
        } else if nums1Left > nums2Right {
            // Too many elements from nums1 in left partition
            right = i - 1
        } else {
            // Too few elements from nums1 in left partition
            left = i + 1
        }
    }
    
    return 0.0
}
```

### Approach 2: Merge and Find (Simpler but Slower)
**Time Complexity**: O(m+n)
**Space Complexity**: O(1)

```go
func findMedianSortedArraysMerge(nums1 []int, nums2 []int) float64 {
    m, n := len(nums1), len(nums2)
    total := m + n
    medianPos := total / 2
    
    i, j := 0, 0
    prev, curr := 0, 0
    
    // Merge until we reach median position
    for count := 0; count <= medianPos; count++ {
        prev = curr
        if i < m && (j >= n || nums1[i] <= nums2[j]) {
            curr = nums1[i]
            i++
        } else {
            curr = nums2[j]
            j++
        }
    }
    
    if total%2 == 1 {
        return float64(curr)
    } else {
        return float64(prev+curr) / 2.0
    }
}
```

## Step-by-Step Walkthrough (Binary Search Approach)

### Example:
```
nums1 = [1, 3, 8]
nums2 = [2, 7, 9, 10]
```

**Step 1**: Ensure `nums1` is smaller (already true)
- m = 3, n = 4, total = 7, half = 4

**Step 2**: Initial binary search bounds
- left = 0, right = 3

**Step 3**: First iteration (i = 1)
- i = 1 (take 1 element from nums1)
- j = 4 - 1 = 3 (take 3 elements from nums2)
- Partition: nums1 = [1 | 3, 8], nums2 = [2, 7, 9 | 10]
- Check: maxLeft = max(1, 9) = 9, minRight = min(3, 10) = 3
- Condition fails: 9 > 3, so move left

**Step 4**: Second iteration (i = 0)
- i = 0 (take 0 elements from nums1)
- j = 4 - 0 = 4 (take 4 elements from nums2) - but j > n, invalid

**Step 5**: Continue binary search until valid partition found

**Final valid partition**:
- i = 2 (take 2 elements from nums1: [1, 3])
- j = 2 (take 2 elements from nums2: [2, 7])
- Partition: nums1 = [1, 3 | 8], nums2 = [2, 7 | 9, 10]
- maxLeft = max(3, 7) = 7, minRight = min(8, 9) = 8
- Condition satisfied: 7 ≤ 8
- Median = maxLeft = 7 (since total is odd)

## Common Pitfalls
1. **Off-by-one errors** in partition calculations
2. **Forgetting to handle edge cases** (empty arrays, single element arrays)
3. **Incorrect binary search termination condition**
4. **Not ensuring nums1 is the smaller array** for optimal performance

## Related Problems
- Find kth element in two sorted arrays
- Merge k sorted arrays
- Find median in a data stream

## Practice Tips
1. Draw the partition diagram to visualize the approach
2. Test with edge cases: empty arrays, arrays of different sizes
3. Practice deriving the time complexity
4. Try to implement both O(m+n) and O(log(min(m, n))) solutions