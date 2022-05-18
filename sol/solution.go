package sol

import (
	"fmt"
	"strconv"
	"strings"
)

/**
* Definition for a binary tree node.
* type TreeNode struct {
*     Val int
*     Left *TreeNode
*     Right *TreeNode
* }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	result := []string{}
	DFS(root, &result)
	return strings.Join(result, ",")
}
func DFS(root *TreeNode, result *[]string) {
	if root == nil {
		*result = append(*result, "N")
		return
	}
	*result = append(*result, fmt.Sprintf("%d", root.Val))
	DFS(root.Left, result)
	DFS(root.Right, result)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	i := 0
	return DE_DFS(&vals, &i)
}

func DE_DFS(vals *[]string, i *int) *TreeNode {
	if (*vals)[*i] == "N" {
		*i++
		return nil
	}
	num, _ := strconv.Atoi((*vals)[*i])
	*i++
	node := &TreeNode{Val: num}
	node.Left = DE_DFS(vals, i)
	node.Right = DE_DFS(vals, i)
	return node
}
