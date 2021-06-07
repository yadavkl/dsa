package tree

import "testing"

func Test_Diameter(t *testing.T) {
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
	root.right.left.Val = 30
	root.right.right = NewNode()
	root.right.right.Val = 40
	expected := 5

	result := 0

	height := Height(root, &result)
	if expected != result {
		t.Errorf("Expected: %v, Got: %v, Height: %v", expected, result, height)
	}
}
