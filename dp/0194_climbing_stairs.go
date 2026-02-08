package dp

// climbStairs calculates the number of distinct ways to climb to the top of stairs
// You can either climb 1 or 2 steps at a time
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	
	prev2 := 1 // ways to reach step 1
	prev1 := 2 // ways to reach step 2
	
	for i := 3; i <= n; i++ {
		current := prev1 + prev2
		prev2 = prev1
		prev1 = current
	}
	
	return prev1
}