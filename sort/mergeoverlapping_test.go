package sort

import (
	"testing"
)

func TestMergeOverlapping(t *testing.T) {
	arr := []Pair{{5, 10}, {3, 15}, {18, 30}, {2, 7}}
	expected := []Pair{{2, 15}, {18, 30}}
	result := MergeOverlapping(arr)

	for idx, pair := range expected {
		if result[idx] != pair {
			t.Errorf("Expected: %v Got: %v", pair, result[idx])
		}
	}
}
