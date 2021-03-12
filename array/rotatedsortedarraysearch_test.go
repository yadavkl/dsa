package array

import (
	"fmt"
	"testing"
)

func TestRotatedSortedArraySearch(t *testing.T) {
	tests := []struct {
		Array    []int
		Key      int
		Expected int
	}{
		{
			[]int{10, 20, 40, 60, 5, 8},
			5,
			4,
		},
		{
			[]int{100, 500, 10, 20, 30, 40, 50},
			40,
			5,
		},
		{
			[]int{100, 200, 500, 1000, 2000, 10, 20},
			40,
			-1,
		},
	}
	for idx, test := range tests {
		t.Run(fmt.Sprintf("Test-Case-%d", idx), func(t *testing.T) {
			result := RotatedSortedArraySearch(test.Array, test.Key)
			if result != test.Expected {
				t.Errorf("Expected: %d Got: %d", test.Expected, result)
			}
		})
	}
}
