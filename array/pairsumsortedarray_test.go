package array

import (
	"fmt"
	"testing"
)

func TestPairSumSortedArray(t *testing.T) {
	tests := []struct {
		Arr      []int
		Sum      int
		Expected bool
	}{
		{
			[]int{2, 4, 8, 9, 11, 12, 20, 30},
			23,
			true,
		},
		{
			[]int{3, 8, 13, 18},
			14,
			false,
		},
	}

	for idx, test := range tests {
		t.Run(fmt.Sprintf("Test-Case-%v", idx), func(t *testing.T) {
			result := PairSumSortedArray(test.Arr, test.Sum)
			if result != test.Expected {
				t.Errorf("Expected: %v Got: %v", test.Expected, result)
			}
		})
	}
}
