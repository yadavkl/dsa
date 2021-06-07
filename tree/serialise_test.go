package tree

import "testing"

func Test_Serialise(t *testing.T) {
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

	expected := []int{10, 80, 90, -1, -1, 100, -1, -1, 30, 35, -1, -1, 40, -1, 50, -1, -1}
	arr := Serialise(root)

	if len(arr) != len(expected) {
		t.Errorf("Expected: %v Got: %v", expected, arr)
	}
	for idx, val := range expected {
		if arr[idx] != val {
			t.Errorf("Expected Val: %v Got: %v", val, arr[idx])
		}
	}
}

func Test_Deserialise(t *testing.T) {
	arr := []int{10, 80, 90, -1, -1, 100, -1, -1, 30, 35, -1, -1, 40, -1, 50, -1, -1}
	root := Deserialise(arr)
	expected := []int{10, 80, 90, -1, -1, 100, -1, -1, 30, 35, -1, -1, 40, -1, 50, -1, -1}
	result := Serialise(root)
	if len(result) != len(expected) {
		t.Errorf("Expected: %v Got: %v", expected, result)
	}
	for idx, val := range expected {
		if result[idx] != val {
			t.Errorf("Expected: %v Got: %v", val, result[idx])
		}
	}
}
