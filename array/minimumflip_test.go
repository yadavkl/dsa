package array

import (
	"fmt"
	"testing"
)

func TestMinimumFlip(t *testing.T) {
	testCases := []struct {
		arr      []int
		expected string
	}{
		{
			[]int{1, 1, 0, 0, 0, 1, 1, 1, 0, 1, 1, 0, 1},
			"From 2 to 4 From 8 to 8 From 11 to 11 ",
		},
	}
	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Test-Case-%v", idx), func(t *testing.T) {
			result := minimumflip(tc.arr)
			if result != tc.expected {
				t.Errorf("Expected : %q Got: %q", tc.expected, result)
			}
		})
	}
}
