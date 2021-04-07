package sort

import "testing"

func TestChocolateDistribution(t *testing.T) {
	arr := []int{7, 3, 2, 4, 9, 12, 56}
	m := 3
	expected := 2
	result := ChocolateDistribution(arr, m)
	if result != expected {
		t.Errorf("Expected: %v Got: %v", expected, result)
	}
}
