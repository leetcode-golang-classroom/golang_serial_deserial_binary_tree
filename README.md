# golang_serialize_deserialize_binary_tree

Serialization is the process of converting a data structure or object into a sequence of bits so that it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed later in the same or another computer environment.

Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be serialized to a string and this string can be deserialized to the original tree structure.

**Clarification:** The input/output format is the same as [how LeetCode serializes a binary tree](https://leetcode.com/faq/#binary-tree). You do not necessarily need to follow this format, so please be creative and come up with different approaches yourself.

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2020/09/15/serdeser.jpg](https://assets.leetcode.com/uploads/2020/09/15/serdeser.jpg)

```
Input: root = [1,2,3,null,null,4,5]
Output: [1,2,3,null,null,4,5]

```

**Example 2:**

```
Input: root = []
Output: []

```

**Constraints:**

- The number of nodes in the tree is in the range `[0, 104]`.
- `1000 <= Node.val <= 1000`

## 解析

給定一個二元樹結點 root 要求要實作 serial, deserial 方法

可以透過 DFS 實作 Pre-order format

如下圖

![](https://i.imgur.com/rMQXbOx.png)

## 程式碼
```go
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

```
## 困難點

1. Understand DFS

## Solve Point

- [x]  Understand what problem need to solve
- [x]  Analysis Complexity