package sort

import "testing"

func TestUnion(t *testing.T) {
	arr1 := []int{2, 10, 20, 20, 60}
	arr2 := []int{3, 20, 40}
	expected := []int{2, 3, 10, 20, 40, 60}

	result := Union(arr1, arr2)

	for idx, val := range expected {
		if result[idx] != val {
			t.Errorf("Expected: %v Got: %v", val, result[idx])
		}
	}
}
