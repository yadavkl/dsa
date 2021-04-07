package sort

import "testing"

func TestSelection(t *testing.T) {
	arr := []int{6, 5, 7, 9, 1, 2, 3, 4}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 9}

	Selection(arr)
	for idx, val := range expected {
		if arr[idx] != val {
			t.Errorf("Expected: %v Gor: %v", val, arr[idx])
		}
	}
}
