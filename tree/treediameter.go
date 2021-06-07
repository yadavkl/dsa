package tree

func Height(root *Node, dia *int) int {
	if root == nil {
		return 0
	}
	lh := Height(root.left, dia)
	rh := Height(root.right, dia)

	*dia = max(*dia, 1+lh+rh)
	return 1 + max(lh, rh)
}
