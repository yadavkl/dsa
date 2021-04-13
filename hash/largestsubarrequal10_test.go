package hash

import (
	"fmt"
	"testing"
)

func Test_LargestSubarrayEqual10(t *testing.T) {
	tests := []struct {
		Arr      []int
		Expected int
	}{
		{
			[]int{1, 0, 1, 1, 1, 0, 0},
			6,
		},
		{
			[]int{1, 1, 1, 1},
			0,
		},
		{
			[]int{0, 0, 1, 1, 1, 1, 1, 0},
			4,
		},
	}
	for idx, test := range tests {
		t.Run(fmt.Sprintf("Test-Case-%v", idx), func(t *testing.T) {
			result := LargestSubarrayWithEqual01(test.Arr)
			if result != test.Expected {
				t.Errorf("Expected: %v Got:%v", test.Expected, result)
			}
		})
	}
}
