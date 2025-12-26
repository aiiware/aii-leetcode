package leetcode

import (
    "reflect"
    "testing"
)

func TestRotate(t *testing.T) {
    tests := []struct {
        matrix   [][]int
        expected [][]int
    }{
        {
            [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
            [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
        },
        {
            [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
            [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
        },
        {
            [][]int{{1}},
            [][]int{{1}},
        },
        {
            [][]int{{1, 2}, {3, 4}},
            [][]int{{3, 1}, {4, 2}},
        },
        {
            [][]int{},
            [][]int{},
        },
        {
            [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}},
            [][]int{{13, 9, 5, 1}, {14, 10, 6, 2}, {15, 11, 7, 3}, {16, 12, 8, 4}},
        },
    }
    
    for i, test := range tests {
        // Test transpose + reverse approach
        matrixCopy := make([][]int, len(test.matrix))
        for j := range test.matrix {
            matrixCopy[j] = make([]int, len(test.matrix[j]))
            copy(matrixCopy[j], test.matrix[j])
        }
        
        Rotate(matrixCopy)
        if !reflect.DeepEqual(matrixCopy, test.expected) {
            t.Errorf("Test %d (transpose+reverse) failed: got %v, expected %v", 
                i, matrixCopy, test.expected)
        }
        
        // Test layer by layer approach
        matrixCopy = make([][]int, len(test.matrix))
        for j := range test.matrix {
            matrixCopy[j] = make([]int, len(test.matrix[j]))
            copy(matrixCopy[j], test.matrix[j])
        }
        
        RotateLayer(matrixCopy)
        if !reflect.DeepEqual(matrixCopy, test.expected) {
            t.Errorf("Test %d (layer) failed: got %v, expected %v", 
                i, matrixCopy, test.expected)
        }
        
        // Test copy approach (not in-place)
        rotated := RotateCopy(test.matrix)
        if !reflect.DeepEqual(rotated, test.expected) {
            t.Errorf("Test %d (copy) failed: got %v, expected %v", 
                i, rotated, test.expected)
        }
    }
}

func BenchmarkRotate(b *testing.B) {
    matrix := [][]int{
        {1, 2, 3, 4, 5},
        {6, 7, 8, 9, 10},
        {11, 12, 13, 14, 15},
        {16, 17, 18, 19, 20},
        {21, 22, 23, 24, 25},
    }
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        matrixCopy := make([][]int, len(matrix))
        for j := range matrix {
            matrixCopy[j] = make([]int, len(matrix[j]))
            copy(matrixCopy[j], matrix[j])
        }
        Rotate(matrixCopy)
    }
}

func BenchmarkRotateLayer(b *testing.B) {
    matrix := [][]int{
        {1, 2, 3, 4, 5},
        {6, 7, 8, 9, 10},
        {11, 12, 13, 14, 15},
        {16, 17, 18, 19, 20},
        {21, 22, 23, 24, 25},
    }
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        matrixCopy := make([][]int, len(matrix))
        for j := range matrix {
            matrixCopy[j] = make([]int, len(matrix[j]))
            copy(matrixCopy[j], matrix[j])
        }
        RotateLayer(matrixCopy)
    }
}

func BenchmarkRotateCopy(b *testing.B) {
    matrix := [][]int{
        {1, 2, 3, 4, 5},
        {6, 7, 8, 9, 10},
        {11, 12, 13, 14, 15},
        {16, 17, 18, 19, 20},
        {21, 22, 23, 24, 25},
    }
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        RotateCopy(matrix)
    }
}