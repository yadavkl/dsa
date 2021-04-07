package sort

import "testing"

func TestKthSmallest(t *testing.T) {
	arr := []int{9, 8, 7, 2, 3, 4, 5, 1, 6}
	k := 7
	expected := 7
	result := KthSmallest(arr, k)
	if result != expected {
		t.Errorf("Expected: %v Got: %v", expected, result)
	}
}
