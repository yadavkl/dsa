package hash

import (
	"fmt"
	"testing"
)

func Test_LongestSubarraySum(t *testing.T) {
	tests := []struct {
		Arr      []int
		Sum      int
		Expected int
	}{
		{
			[]int{5, 8, -4, -4, 9, -2, 2},
			0,
			3,
		},
		{
			[]int{3, 1, 0, 1, 8, 2, 3, 6},
			5,
			4,
		},
		{
			[]int{8, 3, 7},
			15,
			0,
		},
	}

	for idx, test := range tests {
		t.Run(fmt.Sprintf("Test-Case-%v", idx), func(t *testing.T) {
			result := LongestSubarraySum(test.Arr, test.Sum)
			if result != test.Expected {
				t.Errorf("Expected: %v Got: %v", test.Expected, result)
			}
		})
	}
}
