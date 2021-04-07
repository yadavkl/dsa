package sort

import "testing"

func TestIntersection(t *testing.T) {
	arr1 := []int{2, 20, 20, 40}
	arr2 := []int{10, 20, 20, 40, 60}
	expected := []int{20, 40}

	result := Intersection(arr1, arr2)

	for idx, val := range expected {
		if result[idx] != val {
			t.Errorf("Expected: %v Got: %v", val, result[idx])
		}
	}
}
