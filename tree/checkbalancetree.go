package tree

func IsBalanced(root *Node) int {
	if root == nil {
		return 0
	}
	lh := IsBalanced(root.left)
	if lh == -1 {
		return -1
	}
	rh := IsBalanced(root.right)
	if rh == -1 {
		return -1
	}
	if Abs(lh, rh) > 1 {
		return -1
	}
	return max(lh, rh) + 1
}

func Abs(x, y int) int {
	if x-y > 0 {
		return x - y
	}
	return y - x
}
