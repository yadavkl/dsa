package sort

import "testing"

func TestSort2Type(t *testing.T) {
	arr := []int{-12, 18, -10, 15}
	expected := []int{-12, -10, 18, 15}
	Sort2Type(arr)
	for idx, val := range expected {
		if arr[idx] != val {
			t.Errorf("Expected: %v Got: %v", val, arr[idx])
		}
	}
}
