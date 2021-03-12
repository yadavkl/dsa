package array

import "testing"

func TestPreakElement(t *testing.T) {
	arr := []int{5, 20, 40, 30, 20, 50, 60}
	expected := 40
	result := PeakElement(arr)

	if expected != result {
		t.Errorf("Expected : %d Got: %d", expected, result)
	}
}
