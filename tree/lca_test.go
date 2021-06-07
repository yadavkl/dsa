package tree

import "testing"

func Test_Lca(t *testing.T) {
	root := NewNode()
	root.Val = 10
	root.left = NewNode()
	root.left.Val = 80
	root.left.left = NewNode()
	root.left.left.Val = 90
	root.left.right = NewNode()
	root.left.right.Val = 100
	root.right = NewNode()
	root.right.Val = 30
	root.right.left = NewNode()
	root.right.left.Val = 35
	root.right.right = NewNode()
	root.right.right.Val = 40
	root.right.right.right = NewNode()
	root.right.right.right.Val = 50

	n1 := 35
	n2 := 50
	expected := root.right
	result := Lca(root, n1, n2)
	if result != expected {
		t.Errorf("Expected: %v Got: %v", expected.Val, result.Val)
	}
}
