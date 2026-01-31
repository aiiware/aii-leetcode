package arrays


/*
Difficulty: Hard
Tags: [Add relevant tags]
Companies: [Add company names]
*/

/*
# 0135 - Candy
## Problem Description
There are n children standing in a line. Each child is assigned a rating value given in the integer array ratings.

You are giving candies to these children subjected to the following requirements:
- Each child must have at least one candy.
- Children with a higher rating get more candies than their neighbors.

Return the minimum number of candies you need to have to distribute the candies to the children.

## Examples
Example 1:
Input: ratings = [1,0,2]
Output: 5
Explanation: You can allocate to the first, second and third child with 2, 1, 2 candies respectively.

Example 2:
Input: ratings = [1,2,2]
Output: 4
Explanation: You can allocate to the first, second and third child with 1, 2, 1 candies respectively.
The third child gets 1 candy because it satisfies the above two conditions.

## Constraints
- n == ratings.length
- 1 <= n <= 2 * 10^4
- 0 <= ratings[i] <= 2 * 10^4

## Solution Approach
This problem can be solved using two passes:
1. Left to right pass: Give each child one more candy than left neighbor if rating is higher
2. Right to left pass: Give each child max(current, right neighbor + 1) if rating is higher than right neighbor
3. Sum all candies

Time Complexity: O(N) where N is the number of children
Space Complexity: O(N) for the candies array
*/

// Candy returns the minimum number of candies needed
func Candy(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	
	// Initialize candies array with 1 candy for each child
	candies := make([]int, n)
	for i := range candies {
		candies[i] = 1
	}
	
	// Left to right pass
	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}
	
	// Right to left pass
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			// Take maximum of current candy and right neighbor + 1
			if candies[i] <= candies[i+1] {
				candies[i] = candies[i+1] + 1
			}
		}
	}
	
	// Sum all candies
	total := 0
	for _, candy := range candies {
		total += candy
	}
	
	return total
}