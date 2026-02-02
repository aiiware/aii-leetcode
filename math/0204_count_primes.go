package math

// 0204 - Count Primes (Medium)
// https://leetcode.com/problems/count-primes/

// CountPrimes returns the number of prime numbers less than n
// Uses the Sieve of Eratosthenes algorithm
// Time Complexity: O(n log log n) - Sieve of Eratosthenes
// Space Complexity: O(n) - for the sieve array
func CountPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	
	// Create a sieve array where true means the number is prime
	isPrime := make([]bool, n)
	for i := 2; i < n; i++ {
		isPrime[i] = true
	}
	
	// Sieve of Eratosthenes
	// We only need to check up to sqrt(n)
	for i := 2; i*i < n; i++ {
		if isPrime[i] {
			// Mark all multiples of i as non-prime
			for j := i * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	
	// Count primes
	count := 0
	for i := 2; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}
	
	return count
}

// CountPrimesOptimized uses optimized Sieve of Eratosthenes
// Only considers odd numbers to save space and time
func CountPrimesOptimized(n int) int {
	if n <= 2 {
		return 0
	}
	
	// isPrime[i] represents whether 2*i + 3 is prime
	// We start from 3 and only consider odd numbers
	size := (n - 1) / 2
	isPrime := make([]bool, size)
	for i := range isPrime {
		isPrime[i] = true
	}
	
	// Count 2 as a prime
	count := 1
	
	// Sieve for odd numbers
	for i := 0; i < size; i++ {
		if isPrime[i] {
			// The actual number is 2*i + 3
			p := 2*i + 3
			if p >= n {
				break
			}
			count++
			
			// Mark multiples of p starting from p*p
			// Only odd multiples since p is odd
			start := (p*p - 3) / 2
			if start >= size {
				continue
			}
			for j := start; j < size; j += p {
				isPrime[j] = false
			}
		}
	}
	
	return count
}

// CountPrimesSimple is a simpler implementation for clarity
func CountPrimesSimple(n int) int {
	if n <= 2 {
		return 0
	}
	
	notPrime := make([]bool, n)
	count := 0
	
	for i := 2; i < n; i++ {
		if !notPrime[i] {
			count++
			// Mark multiples of i
			for j := i * i; j < n; j += i {
				notPrime[j] = true
			}
		}
	}
	
	return count
}