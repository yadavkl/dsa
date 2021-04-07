package sort

import "testing"

func TestCycleSort(t *testing.T) {
	arr := []int{10, 30, 40, 50, 20}
	expected := []int{10, 20, 30, 40, 50}

	CycleSort(arr)

	for idx, val := range expected {
		if val != arr[idx] {
			t.Errorf("Expected: %v Got: %v", val, arr[idx])
		}
	}
}
