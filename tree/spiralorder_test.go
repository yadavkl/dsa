package tree

import "testing"

func Test_SpiralOrder(t *testing.T) {
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

	expected := []int{10, 30, 80, 90, 100, 30, 40}
	result := make([]int, 0)
	SpiralOrder(root, &result)

	if len(expected) != len(result) {
		t.Errorf("Expected: %v, Got: %v", expected, result)
	}
	for idx, val := range expected {
		if val != result[idx] {
			t.Errorf("idx: %v, expected: %v, Got: %v", idx, val, result[idx])
		}
	}
}
