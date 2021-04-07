package sort

import "testing"

func TestQuickLomuto(t *testing.T) {
	arr := []int{3, 4, 9, 7, 1, 2, 6, 8, 5}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	QuickLomuto(arr, 0, len(arr)-1)

	for idx, val := range expected {
		if arr[idx] != val {
			t.Errorf("Expected: %v Got: %v", val, arr[idx])
		}
	}
}
