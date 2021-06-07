package tree

func Lca(root *Node, n1, n2 int) *Node {
	if root == nil {
		return nil
	}
	if root.Val == n1 || root.Val == n2 {
		return root
	}

	lca1 := Lca(root.left, n1, n2)
	lca2 := Lca(root.right, n1, n2)

	if lca1 != nil && lca2 != nil {
		return root
	}

	if lca1 != nil {
		return lca1
	}
	return lca2
}
