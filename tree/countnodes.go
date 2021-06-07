package tree

import "math"

func CountNodes(root *Node) int {
	if root == nil {
		return 0
	}
	lh := 0
	rh := 0
	cur := root
	for cur != nil {
		lh++
		cur = cur.left
	}
	cur = root
	for cur != nil {
		rh++
		cur = cur.right
	}

	if lh == rh {
		return int(math.Pow(float64(2), float64(lh))) - 1
	}
	return 1 + CountNodes(root.left) + CountNodes(root.right)

}
