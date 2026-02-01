package trees

import (
	"strconv"
	"strings"
	
	"leetcode/utils"
)

// Codec provides methods to serialize and deserialize a binary tree
type Codec struct{}

// Constructor initializes a new Codec
func Constructor() Codec {
	return Codec{}
}

// Serialize converts a binary tree to a string
// Uses level-order traversal with "null" for nil nodes
func (c *Codec) Serialize(root *utils.TreeNode) string {
	if root == nil {
		return ""
	}

	var result []string
	queue := []*utils.TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, "null")
		} else {
			result = append(result, strconv.Itoa(node.Val))
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}

	// Remove trailing nulls
	for i := len(result) - 1; i >= 0; i-- {
		if result[i] != "null" {
			result = result[:i+1]
			break
		}
	}

	return strings.Join(result, ",")
}

// Deserialize converts a string back to a binary tree
func (c *Codec) Deserialize(data string) *utils.TreeNode {
	if data == "" {
		return nil
	}

	values := strings.Split(data, ",")
	if len(values) == 0 || values[0] == "null" {
		return nil
	}

	// Parse root value
	rootVal, _ := strconv.Atoi(values[0])
	root := &utils.TreeNode{Val: rootVal}

	queue := []*utils.TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(values) {
		node := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(values) && values[i] != "null" {
			leftVal, _ := strconv.Atoi(values[i])
			node.Left = &utils.TreeNode{Val: leftVal}
			queue = append(queue, node.Left)
		}
		i++

		// Right child
		if i < len(values) && values[i] != "null" {
			rightVal, _ := strconv.Atoi(values[i])
			node.Right = &utils.TreeNode{Val: rightVal}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// SerializeTree is a convenience function that serializes a tree
func SerializeTree(root *utils.TreeNode) string {
	codec := Constructor()
	return codec.Serialize(root)
}

// DeserializeTree is a convenience function that deserializes a string to a tree
func DeserializeTree(data string) *utils.TreeNode {
	codec := Constructor()
	return codec.Deserialize(data)
}