package hash

import (
	"fmt"
	"testing"
)

func Test_DistinctElementInWindow(t *testing.T) {
	tests := []struct {
		Arr      []int
		Window   int
		Expected []int
	}{
		{
			[]int{10, 20, 20, 10, 30, 40, 10},
			4,
			[]int{2, 3, 4, 3},
		},
		{
			[]int{10, 10, 10, 10},
			3,
			[]int{1, 1},
		},
		{
			[]int{10, 20, 30, 40},
			3,
			[]int{3, 3},
		},
	}
	for idx, test := range tests {
		t.Run(fmt.Sprintf("Test-Case-%v", idx), func(t *testing.T) {
			result := DistinctElementInWindow(test.Arr, test.Window)
			if len(result) != len(test.Expected) {
				t.Errorf("Expected Length : %v Got: %v", len(test.Expected), len(result))
			}
			for idx, val := range test.Expected {
				if val != result[idx] {
					t.Errorf("Expected : %v Index: %v Got : %v", val, idx, result[idx])
				}
			}
		})
	}
}
