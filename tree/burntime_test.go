package tree

import "testing"

func Test_BurnTime(t *testing.T) {
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
	expected := 4
	leaf := 100
	time := 0
	dist := -1
	_ = BurnTime(root, leaf, &dist, &time)

	if expected != time {
		t.Errorf("expected: %v Got: %v", expected, time)
	}

}
