package array

import (
	"fmt"
	"testing"
)

func TestAllocateMinPages(t *testing.T) {
	tests := []struct {
		Arr      []int
		K        int
		Expected int
	}{
		{
			[]int{10, 20, 30, 40},
			2,
			60,
		},
		{
			[]int{10, 20, 30, 40},
			1,
			100,
		},
		{
			[]int{40},
			2,
			40,
		},
		{
			[]int{10, 5, 30, 1, 2, 5, 10, 10},
			3,
			30,
		},
	}
	for idx, test := range tests {
		t.Run(fmt.Sprintf("Test-Case-%v", idx), func(t *testing.T) {
			result := AllocateMinPages(test.Arr, len(test.Arr), test.K)

			if result != test.Expected {
				t.Errorf("Expected: %v Got: %v", test.Expected, result)
			}
		})
	}

}
