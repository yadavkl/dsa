package sort

import "testing"

func TestBubble(t *testing.T) {
	arr := []int{5, 4, 3, 40, 50, 32, 21}
	expected := []int{3, 4, 5, 21, 32, 40, 50}
	result := Bubble(arr)
	for idx, val := range expected {
		if result[idx] != val {
			t.Errorf("Expected: %v Got: %v", val, result[idx])
		}
	}
}
