package hash

import (
	"fmt"
	"testing"
)

func Test_SubArraySumZero(t *testing.T) {
	tests := []struct {
		Arr      []int
		Expected bool
	}{
		{
			[]int{-3, 4, -3, -1, 1},
			true,
		},
		{
			[]int{-3, 2, 1, 4},
			true,
		},
	}
	for idx, test := range tests {
		t.Run(fmt.Sprintf("Test-Case-%v", idx), func(t *testing.T) {
			result := SubArraySumZero(test.Arr)
			if result != test.Expected {
				t.Errorf("Expected: %v Got: %v", test.Expected, result)
			}
		})
	}
}
