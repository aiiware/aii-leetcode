package testutils

import (
    "fmt"
    "testing"
    "leetcode/trees"
    "leetcode/utils"
)

func TestCheckSymmetricManually(t *testing.T) {
    // Build the tree manually to ensure it's symmetric
    // Level 0
    root := &utils.TreeNode{Val: 1}
    
    // Level 1
    root.Left = &utils.TreeNode{Val: 2}
    root.Right = &utils.TreeNode{Val: 2}
    
    // Level 2
    root.Left.Left = &utils.TreeNode{Val: 3}
    root.Left.Right = &utils.TreeNode{Val: 4}
    root.Right.Left = &utils.TreeNode{Val: 4}
    root.Right.Right = &utils.TreeNode{Val: 3}
    
    // Level 3
    root.Left.Left.Left = &utils.TreeNode{Val: 5}
    root.Left.Left.Right = &utils.TreeNode{Val: 6}
    root.Left.Right.Left = &utils.TreeNode{Val: 7}
    root.Left.Right.Right = &utils.TreeNode{Val: 7}
    root.Right.Left.Left = &utils.TreeNode{Val: 6}
    root.Right.Left.Right = &utils.TreeNode{Val: 5}
    // Note: root.Right.Right has no children in this symmetric tree
    
    fmt.Println("Manually built tree:")
    fmt.Println("Testing IsSymmetric:", trees.IsSymmetric(root))
    
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
    tree2 := utils.NewTreeFromSlice(slice)
    fmt.Println("\nRebuilt tree from slice:")
    fmt.Println("Testing IsSymmetric:", trees.IsSymmetric(tree2))
}