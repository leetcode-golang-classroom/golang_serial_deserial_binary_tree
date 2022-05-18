package sol

func Run(root *TreeNode) *TreeNode {
	ser := Constructor()
	deser := Constructor()
	data := ser.serialize(root)
	ans := deser.deserialize(data)
	return ans
}
