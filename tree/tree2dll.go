package tree

var prev *Node

func Tree2DLL(root *Node) *Node {

	if root == nil {
		return root
	}
	head := Tree2DLL(root.left)
	if prev == nil {
		head = root
	} else {
		root.left = prev
		prev.right = root
	}
	prev = root

	Tree2DLL(root.right)
	return head
}
