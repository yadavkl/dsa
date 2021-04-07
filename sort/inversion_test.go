package sort

import "testing"

func TestInversions(t *testing.T) {
	arr := []int{2, 5, 8, 11, 3, 6, 9, 13}
	expected := 6
	result := Inversions(arr, 0, len(arr)-1)
	if result != expected {
		t.Errorf("Expected: %v Got: %v", expected, result)
	}
}
