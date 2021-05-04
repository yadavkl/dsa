package tree

func KDistNode(root *Node, k int, nodes *[]int) {
	if root == nil {
		return
	}
	if k == 0 {
		*(nodes) = append(*(nodes), root.Val)
	}
	KDistNode(root.left, k-1, nodes)
	KDistNode(root.right, k-1, nodes)
}
