package sort

import "testing"

func TestInsertion(t *testing.T) {
	arr := []int{9, 7, 8, 4, 6, 5, 3, 1, 2}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	Insertion(arr)
	for idx, val := range expected {
		if arr[idx] != val {
			t.Errorf("Expected: %v Got: %v", val, arr[idx])
		}
	}
}
