// 0050 - Pow(x, n)
// https://leetcode.com/problems/powx-n/
// Medium - Math, Recursion

package leetcode

// MyPow calculates x raised to the power n.
// Time Complexity: O(log n)
// Space Complexity: O(log n) for recursion, O(1) for iterative
func MyPow(x float64, n int) float64 {
    // Handle negative exponent
    if n < 0 {
        x = 1 / x
        n = -n
    }
    
    return powRecursive(x, n)
}

// powRecursive uses recursive binary exponentiation
func powRecursive(x float64, n int) float64 {
    if n == 0 {
        return 1.0
    }
    
    half := powRecursive(x, n/2)
    
    if n%2 == 0 {
        return half * half
    } else {
        return half * half * x
    }
}

// MyPowIterative uses iterative binary exponentiation
func MyPowIterative(x float64, n int) float64 {
    // Handle negative exponent
    if n < 0 {
        x = 1 / x
        n = -n
    }
    
    result := 1.0
    current := x
    
    for n > 0 {
        if n%2 == 1 {
            result *= current
        }
        current *= current
        n /= 2
    }
    
    return result
}

// MyPowBruteForce uses simple multiplication (for testing)
func MyPowBruteForce(x float64, n int) float64 {
    if n < 0 {
        x = 1 / x
        n = -n
    }
    
    result := 1.0
    for i := 0; i < n; i++ {
        result *= x
    }
    
    return result
}