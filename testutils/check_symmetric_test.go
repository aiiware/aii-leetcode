package testutils

import (
    "fmt"
    "testing"
    "leetcode"
)

func TestCheckSymmetricManually(t *testing.T) {
    // Build the tree manually to ensure it's symmetric
    // Level 0
    root := &leetcode.TreeNode{Val: 1}
    
    // Level 1
    root.Left = &leetcode.TreeNode{Val: 2}
    root.Right = &leetcode.TreeNode{Val: 2}
    
    // Level 2
    root.Left.Left = &leetcode.TreeNode{Val: 3}
    root.Left.Right = &leetcode.TreeNode{Val: 4}
    root.Right.Left = &leetcode.TreeNode{Val: 4}
    root.Right.Right = &leetcode.TreeNode{Val: 3}
    
    // Level 3
    root.Left.Left.Left = &leetcode.TreeNode{Val: 5}
    root.Left.Left.Right = &leetcode.TreeNode{Val: 6}
    root.Left.Right.Left = &leetcode.TreeNode{Val: 7}
    root.Left.Right.Right = &leetcode.TreeNode{Val: 7}
    root.Right.Left.Left = &leetcode.TreeNode{Val: 6}
    root.Right.Left.Right = &leetcode.TreeNode{Val: 5}
    // Note: root.Right.Right has no children in this symmetric tree
    
    fmt.Println("Manually built tree:")
    fmt.Println("Testing IsSymmetric:", leetcode.IsSymmetric(root))
    
    // Now let's see what ToSlice() produces
    slice := root.ToSlice()
    fmt.Println("\nToSlice() result:")
    for i, val := range slice {
        if val == nil {
            fmt.Printf("  [%d]: nil\n", i)
        } else {
            fmt.Printf("  [%d]: %d\n", i, *val)
        }
    }
    
    // Now build from slice and test again
    tree2 := leetcode.NewTreeFromSlice(slice)
    fmt.Println("\nRebuilt tree from slice:")
    fmt.Println("Testing IsSymmetric:", leetcode.IsSymmetric(tree2))
}