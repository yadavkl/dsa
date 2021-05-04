package tree

import "testing"

func Test_IsBalanced(t *testing.T) {
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
	root.right.right.right = NewNode()
	root.right.right.right.Val = 50

	result := IsBalanced(root)

	if result == -1 {
		t.Errorf("Expected: Yes Got: No")
	}
}
