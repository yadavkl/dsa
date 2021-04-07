package sort

import "testing"

func TestQuickHoare(t *testing.T) {
	arr := []int{22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 3, 4, 9, 7, 1, 2, 6, 8, 5, 10}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22}
	QuickHoare(arr, 0, len(arr)-1)

	for idx, val := range expected {
		if arr[idx] != val {
			t.Errorf("Expected: %v Got: %v", val, arr[idx])
		}
	}
}
