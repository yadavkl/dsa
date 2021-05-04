package array

import (
	"fmt"
	"testing"
)

func TestMedianOfSortedArray(t *testing.T) {
	tests := []struct {
		Arr1     []int
		Arr2     []int
		Expected float64
	}{
		{
			[]int{10, 20, 30, 40, 50},
			[]int{5, 15, 25, 30, 35, 55, 65, 75, 85},
			32.5,
		},
		{
			[]int{10, 20, 30},
			[]int{5, 15, 25, 35, 45},
			22.5,
		},
		{
			[]int{30,40,50,40},
			[]int{5,7,8,9,20},
			20,
		},
	}
	for idx, test := range tests {
		t.Run(fmt.Sprintf("Test-Case-%v", idx), func(t *testing.T) {
			result := MedianOfSortedArray(test.Arr1, test.Arr2)
			if result != test.Expected {
				t.Errorf("Expected: %v Got: %v", test.Expected, result)
			}
		})
	}

}
