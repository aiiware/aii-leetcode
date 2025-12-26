// 0043 - Multiply Strings
// https://leetcode.com/problems/multiply-strings/
// Medium - Math, String, Simulation

package leetcode

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