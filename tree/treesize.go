package tree

func GetTreeSize(root *Node) int {
	if root == nil {
		return 0
	}
	return 1 + GetTreeSize(root.left) + GetTreeSize(root.right)
}
