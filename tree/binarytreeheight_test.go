package tree

import "testing"

func Test_BinaryTreeHeight(t *testing.T) {
	root := NewNode()
	root.Val = 10
	root.left = NewNode()
	root.left.Val = 80
	root.right = NewNode()
	root.right.Val = 30
	root.right.left = NewNode()
	root.right.left.Val = 30
	root.right.right = NewNode()
	root.right.right.Val = 40

	expected := 3

	result := BinaryTreeHeight(root)
	if result != expected {
		t.Errorf("Expected :%v Got: %v", expected, result)
	}
}
