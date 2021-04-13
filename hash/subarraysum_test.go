package hash

import (
	"fmt"
	"testing"
)

func Test_SubArraySum(t *testing.T) {
	tests := []struct {
		Arr      []int
		Sum      int
		Expected bool
	}{
		{
			[]int{5, 8, 6, 13, 3, -1},
			22,
			true,
		},
	}
	for idx, test := range tests {
		t.Run(fmt.Sprintf("Test-Case-%v", idx), func(t *testing.T) {
			result := SubArraySum(test.Arr, test.Sum)
			if result != test.Expected {
				t.Errorf("Expected: %v Got: %v", test.Expected, result)
			}
		})
	}
}
