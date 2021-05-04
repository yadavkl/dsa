package tree

import "math"

func GetMax(root *Node) int {
	if root == nil {
		return math.MinInt32
	}
	return max(root.Val, max(GetMax(root.left), GetMax(root.right)))
}
