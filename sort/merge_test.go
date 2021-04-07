package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{10, 9, 2, 3, 4, 8, 7, 6, 5, 1}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr = MergeSort(arr)
	fmt.Println(arr)
	for idx, val := range expected {
		if arr[idx] != val {
			t.Errorf("Expected: %v Got: %v", val, arr[idx])
		}
	}

}
