package array

import "testing"

func TestEquilibrium(t *testing.T) {
	arr := []int{3, 4, 8, -9, 20, 6}
	expected := true
	result := equilibrium(arr)
	if result != expected {
		t.Errorf("Expected: %v Got: %v", expected, result)
	}
}
