package sort

import "testing"

func TestDutchFlag(t *testing.T) {
	arr := []int{0, 1, 0, 2, 1, 2}
	expected := []int{0, 0, 1, 1, 2, 2}
	DutchFlag(arr)
	for idx, val := range expected {
		if arr[idx] != val {
			t.Errorf("Expected: %v Got: %v", val, arr[idx])
		}
	}
}
