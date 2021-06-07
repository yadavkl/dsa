package tree

func BurnTime(root *Node, leaf int, dist *int, time *int) int {
	if root == nil {
		return 0
	}
	if root.Val == leaf {
		*dist = 0
		return 1
	}
	ldist := -1
	rdist := -1
	lh := BurnTime(root.left, leaf, &ldist, time)
	rh := BurnTime(root.right, leaf, &rdist, time)

	if ldist != -1 {
		*dist = ldist + 1
		*time = max(*time, *dist+rh)
	}
	if rdist != -1 {
		*dist = rdist + 1
		*time = max(*time, *dist+lh)
	}
	return max(lh, rh) + 1
}
