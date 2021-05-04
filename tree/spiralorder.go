package tree

import (
	"github.com/dsa/stack"
)

func SpiralOrder(root *Node, result *[]int) {
	s1 := stack.NewStack()
	s2 := stack.NewStack()

	s1.Push(root)

	for !s1.IsEmpty() || !s2.IsEmpty() {
		for !s1.IsEmpty() {
			node := s1.Pop().(*Node)
			*result = append(*result, node.Val)
			if node.left != nil {
				s2.Push(node.left)
			}
			if node.right != nil {
				s2.Push(node.right)
			}
		}
		for !s2.IsEmpty() {
			node := s2.Pop().(*Node)
			*result = append(*result, node.Val)
			if node.right != nil {
				s1.Push(node.right)
			}
			if node.left != nil {
				s1.Push(node.left)
			}
		}
	}
}
