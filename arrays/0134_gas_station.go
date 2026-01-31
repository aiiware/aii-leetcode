package arrays


/*
Difficulty: Hard
Tags: [Add relevant tags]
Companies: [Add company names]
*/

/*
# 0134 - Gas Station
## Problem Description
There are n gas stations along a circular route, where the amount of gas at the ith station is gas[i].

You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from the ith station to its next (i+1)th station. You begin the journey with an empty tank at one of the gas stations.

Given two integer arrays gas and cost, return the starting gas station's index if you can travel around the circuit once in the clockwise direction, otherwise return -1. If there exists a solution, it is guaranteed to be unique.

## Examples
Example 1:
Input: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
Output: 3
Explanation:
Start at station 3 (index 3) and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
Travel to station 4. Your tank = 4 - 1 + 5 = 8
Travel to station 0. Your tank = 8 - 2 + 1 = 7
Travel to station 1. Your tank = 7 - 3 + 2 = 6
Travel to station 2. Your tank = 6 - 4 + 3 = 5
Travel to station 3. The cost is 5. Your tank = 5 - 5 = 0
You arrive back at station 3 with 0 gas left, so return 3.

Example 2:
Input: gas = [2,3,4], cost = [3,4,3]
Output: -1
Explanation:
You can't start at station 0 or 1, as there is not enough gas to travel to the next station.
Let's start at station 2 and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
Travel to station 0. Your tank = 4 - 3 + 2 = 3
Travel to station 1. Your tank = 3 - 4 + 3 = 2
Travel to station 2. The cost is 3. Your tank = 2 - 3 = -1
You cannot travel back to station 2, as it requires more gas than you have, so return -1.

## Constraints
- n == gas.length == cost.length
- 1 <= n <= 10^5
- 0 <= gas[i], cost[i] <= 10^4

## Solution Approach
This problem can be solved using a greedy approach:
1. If total gas is less than total cost, it's impossible to complete the circuit
2. Otherwise, there must be a solution
3. Start from station 0, track current gas in tank
4. If at any point current gas becomes negative, reset starting point to next station
5. The starting point where we can complete the circuit is guaranteed to exist if total gas >= total cost

Time Complexity: O(N) where N is the number of gas stations
Space Complexity: O(1)
*/

// CanCompleteCircuit returns the starting gas station index if possible, otherwise -1
func CanCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	
	// Check if total gas is enough for total cost
	totalGas, totalCost := 0, 0
	for i := 0; i < n; i++ {
		totalGas += gas[i]
		totalCost += cost[i]
	}
	
	// If total gas is less than total cost, impossible to complete circuit
	if totalGas < totalCost {
		return -1
	}
	
	// Find starting station using greedy approach
	currentGas := 0
	startStation := 0
	
	for i := 0; i < n; i++ {
		currentGas += gas[i] - cost[i]
		
		// If we can't reach next station from current start
		if currentGas < 0 {
			// Reset starting point to next station
			startStation = i + 1
			currentGas = 0
		}
	}
	
	return startStation
}