package arrays

/*
179. Largest Number

Given a list of non-negative integers nums, arrange them such that they form the largest number and return it.
Since the result may be very large, so you need to return a string instead of an integer.

Example 1:
Input: nums = [10,2]
Output: "210"

Example 2:
Input: nums = [3,30,34,5,9]
Output: "9534330"

Constraints:
- 1 <= nums.length <= 100
- 0 <= nums[i] <= 10^9
*/

/*
Difficulty: Medium
Tags: Array, String, Sorting, Greedy
Companies: Amazon, Facebook, Microsoft, Google, Bloomberg
*/

import (
	"sort"
	"strconv"
)

// largestNumber solves LeetCode problem 179: Largest Number
// Given a list of non-negative integers, arranges them to form the largest number
//
// Time complexity: O(n log n) - sorting takes O(n log n)
// Space complexity: O(n) - for storing string representations
func largestNumber(nums []int) string {
	// Convert all numbers to strings
	strNums := make([]string, len(nums))
	for i, num := range nums {
		strNums[i] = strconv.Itoa(num)
	}

	// Sort using custom comparator
	// For two numbers a and b, compare a+b vs b+a
	// If a+b > b+a, then a should come before b
	sort.Slice(strNums, func(i, j int) bool {
		return strNums[i]+strNums[j] > strNums[j]+strNums[i]
	})

	// Concatenate all strings
	result := ""
	for _, numStr := range strNums {
		result += numStr
	}

	// Handle edge case where all numbers are zero
	if result[0] == '0' {
		return "0"
	}

	return result
}

// largestNumberBubbleSort uses bubble sort (for educational purposes)
// This is less efficient but demonstrates the comparison logic clearly
func largestNumberBubbleSort(nums []int) string {
	// Convert all numbers to strings
	strNums := make([]string, len(nums))
	for i, num := range nums {
		strNums[i] = strconv.Itoa(num)
	}

	// Bubble sort with custom comparator
	n := len(strNums)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			// Compare concatenation order
			if strNums[j]+strNums[j+1] < strNums[j+1]+strNums[j] {
				// Swap if wrong order
				strNums[j], strNums[j+1] = strNums[j+1], strNums[j]
			}
		}
	}

	// Concatenate all strings
	result := ""
	for _, numStr := range strNums {
		result += numStr
	}

	// Handle edge case where all numbers are zero
	if result[0] == '0' {
		return "0"
	}

	return result
}

// largestNumberMergeSort uses merge sort with custom comparator
func largestNumberMergeSort(nums []int) string {
	// Convert all numbers to strings
	strNums := make([]string, len(nums))
	for i, num := range nums {
		strNums[i] = strconv.Itoa(num)
	}

	// Merge sort
	strNums = mergeSort(strNums)

	// Concatenate all strings
	result := ""
	for _, numStr := range strNums {
		result += numStr
	}

	// Handle edge case where all numbers are zero
	if result[0] == '0' {
		return "0"
	}

	return result
}

// mergeSort implements merge sort with custom comparator
func mergeSort(strs []string) []string {
	if len(strs) <= 1 {
		return strs
	}

	mid := len(strs) / 2
	left := mergeSort(strs[:mid])
	right := mergeSort(strs[mid:])

	return merge(left, right)
}

// merge combines two sorted arrays using custom comparator
func merge(left, right []string) []string {
	result := make([]string, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i]+right[j] >= right[j]+left[i] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// Append remaining elements
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

// largestNumberQuickSort uses quicksort with custom comparator
func largestNumberQuickSort(nums []int) string {
	// Convert all numbers to strings
	strNums := make([]string, len(nums))
	for i, num := range nums {
		strNums[i] = strconv.Itoa(num)
	}

	// Quicksort with custom comparator
	quickSort(strNums, 0, len(strNums)-1)

	// Concatenate all strings
	result := ""
	for _, numStr := range strNums {
		result += numStr
	}

	// Handle edge case where all numbers are zero
	if result[0] == '0' {
		return "0"
	}

	return result
}

// quickSort implements quicksort with custom comparator
func quickSort(strs []string, low, high int) {
	if low < high {
		pivotIndex := partition(strs, low, high)
		quickSort(strs, low, pivotIndex-1)
		quickSort(strs, pivotIndex+1, high)
	}
}

// partition partitions array around a pivot using custom comparator
func partition(strs []string, low, high int) int {
	pivot := strs[high]
	i := low - 1

	for j := low; j < high; j++ {
		// If strs[j] should come before pivot
		if strs[j]+pivot >= pivot+strs[j] {
			i++
			strs[i], strs[j] = strs[j], strs[i]
		}
	}

	strs[i+1], strs[high] = strs[high], strs[i+1]
	return i + 1
}
