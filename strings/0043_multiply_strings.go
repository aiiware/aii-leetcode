package strings

/*
43. Multiply Strings
https://leetcode.com/problems/multiply-strings/

Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.

Note: You must not use any built-in BigInteger library or convert the inputs to integer directly.

Example 1:
Input: num1 = "2", num2 = "3"
Output: "6"

Example 2:
Input: num1 = "123", num2 = "456"
Output: "56088"

Constraints:
- 1 <= num1.length, num2.length <= 200
- num1 and num2 consist of digits only.
- Both num1 and num2 do not contain any leading zero, except the number 0 itself.

Difficulty: Medium
Tags: Math, String, Simulation
Companies: Facebook, Amazon, Microsoft, Google, Bloomberg, Apple, Uber
*/

// Multiply multiplies two numbers represented as strings.
// Time Complexity: O(m * n) where m and n are lengths of num1 and num2
// Space Complexity: O(m + n)
func Multiply(num1 string, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }
    
    m, n := len(num1), len(num2)
    result := make([]byte, m+n)
    
    // Initialize result with '0'
    for i := range result {
        result[i] = '0'
    }
    
    // Multiply digit by digit
    for i := m - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            // Multiply digits
            mul := int((num1[i] - '0') * (num2[j] - '0'))
            
            // Add to current position
            pos1, pos2 := i+j+1, i+j
            sum := mul + int(result[pos1]-'0')
            
            // Update current position and carry
            result[pos1] = byte(sum%10) + '0'
            result[pos2] += byte(sum / 10)
        }
    }
    
    // Convert to string, skipping leading zeros
    start := 0
    for start < len(result) && result[start] == '0' {
        start++
    }
    
    if start == len(result) {
        return "0"
    }
    
    return string(result[start:])
}

// MultiplySimple uses simpler approach with integer conversion (for small numbers)
func MultiplySimple(num1 string, num2 string) string {
    // Convert strings to integers (only works for small numbers)
    // This is for demonstration only - not suitable for large numbers
    n1 := 0
    for i := 0; i < len(num1); i++ {
        n1 = n1*10 + int(num1[i]-'0')
    }
    
    n2 := 0
    for i := 0; i < len(num2); i++ {
        n2 = n2*10 + int(num2[i]-'0')
    }
    
    // Multiply
    product := n1 * n2
    
    // Convert back to string
    if product == 0 {
        return "0"
    }
    
    // Build string in reverse
    var result []byte
    for product > 0 {
        result = append(result, byte(product%10)+'0')
        product /= 10
    }
    
    // Reverse the string
    for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
        result[i], result[j] = result[j], result[i]
    }
    
    return string(result)
}