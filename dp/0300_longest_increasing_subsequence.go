package dp

// 300. Longest Increasing Subsequence
// https://leetcode.com/problems/longest-increasing-subsequence/
//
// Given an integer array nums, return the length of the longest strictly increasing subsequence.
//
// Example 1:
// Input: nums = [10,9,2,5,3,7,101,18]
// Output: 4
// Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
//
// Example 2:
// Input: nums = [0,1,0,3,2,3]
// Output: 4
//
// Example 3:
// Input: nums = [7,7,7,7,7,7,7]
// Output: 1

// lengthOfLISDP uses dynamic programming with O(n²) time complexity
// Time complexity: O(n²)
// Space complexity: O(n)
func lengthOfLISDP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// dp[i] represents the length of LIS ending at index i
	dp := make([]int, len(nums))
	maxLen := 1

	// Initialize dp array
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	// Fill dp array
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}

	return maxLen
}

// lengthOfLISBinarySearch uses patience sorting with binary search
// Time complexity: O(n log n)
// Space complexity: O(n)
func lengthOfLISBinarySearch(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// tails[i] stores the smallest possible tail value for all increasing subsequences of length i+1
	tails := make([]int, 0, len(nums))
	tails = append(tails, nums[0])

	for i := 1; i < len(nums); i++ {
		if nums[i] > tails[len(tails)-1] {
			// nums[i] extends the largest subsequence
			tails = append(tails, nums[i])
		} else {
			// nums[i] will replace the first element in tails that is >= nums[i]
			// This maintains the invariant that tails[i] is the smallest possible tail
			left, right := 0, len(tails)-1
			for left < right {
				mid := left + (right-left)/2
				if tails[mid] < nums[i] {
					left = mid + 1
				} else {
					right = mid
				}
			}
			tails[left] = nums[i]
		}
	}

	return len(tails)
}

// LengthOfLIS is the main function that chooses the optimal algorithm
// It defaults to binary search for better time complexity
func LengthOfLIS(nums []int) int {
	return lengthOfLISBinarySearch(nums)
}

// reconstructLIS reconstructs the actual longest increasing subsequence
// using the binary search method
func reconstructLIS(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	// tails stores the actual values
	// parent stores the index of the previous element in the LIS
	tails := make([]int, 0, len(nums))
	parent := make([]int, len(nums))
	tailsIdx := make([]int, 0, len(nums)) // stores indices in nums

	tails = append(tails, nums[0])
	tailsIdx = append(tailsIdx, 0)
	parent[0] = -1

	for i := 1; i < len(nums); i++ {
		if nums[i] > tails[len(tails)-1] {
			parent[i] = tailsIdx[len(tailsIdx)-1]
			tails = append(tails, nums[i])
			tailsIdx = append(tailsIdx, i)
		} else {
			left, right := 0, len(tails)-1
			for left < right {
				mid := left + (right-left)/2
				if tails[mid] < nums[i] {
					left = mid + 1
				} else {
					right = mid
				}
			}
			tails[left] = nums[i]
			tailsIdx[left] = i
			if left > 0 {
				parent[i] = tailsIdx[left-1]
			} else {
				parent[i] = -1
			}
		}
	}

	// Reconstruct the LIS
	result := make([]int, len(tails))
	idx := tailsIdx[len(tailsIdx)-1]
	for i := len(tails) - 1; i >= 0; i-- {
		result[i] = nums[idx]
		idx = parent[idx]
	}

	return result
}