package tree

type Node struct {
	Val   int
	left  *Node
	right *Node
}

func NewNode() *Node {
	return &Node{}
}

func BinaryTreeHeight(root *Node) int {
	if root == nil {
		return 0
	}
	return max(BinaryTreeHeight(root.left), BinaryTreeHeight(root.right)) + 1

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
