// 0048 - Rotate Image
// https://leetcode.com/problems/rotate-image/
// Medium - Array, Math, Matrix

package leetcode

// Rotate rotates an n x n matrix 90 degrees clockwise in-place.
// Time Complexity: O(n²)
// Space Complexity: O(1)
func Rotate(matrix [][]int) {
    n := len(matrix)
    
    // Transpose the matrix
    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
        }
    }
    
    // Reverse each row
    for i := 0; i < n; i++ {
        for j := 0; j < n/2; j++ {
            matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
        }
    }
}

// RotateLayer rotates matrix layer by layer
func RotateLayer(matrix [][]int) {
    n := len(matrix)
    
    for layer := 0; layer < n/2; layer++ {
        first := layer
        last := n - 1 - layer
        
        for i := first; i < last; i++ {
            offset := i - first
            
            // Save top
            top := matrix[first][i]
            
            // Left -> Top
            matrix[first][i] = matrix[last-offset][first]
            
            // Bottom -> Left
            matrix[last-offset][first] = matrix[last][last-offset]
            
            // Right -> Bottom
            matrix[last][last-offset] = matrix[i][last]
            
            // Top -> Right
            matrix[i][last] = top
        }
    }
}

// RotateCopy creates a new rotated matrix (not in-place)
func RotateCopy(matrix [][]int) [][]int {
    n := len(matrix)
    if n == 0 {
        return [][]int{}
    }
    
    // Create new matrix
    rotated := make([][]int, n)
    for i := range rotated {
        rotated[i] = make([]int, n)
    }
    
    // Fill rotated matrix
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            rotated[j][n-1-i] = matrix[i][j]
        }
    }
    
    return rotated
}