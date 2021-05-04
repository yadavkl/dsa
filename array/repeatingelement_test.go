package array

import (
	"fmt"
	"testing"
)

func TestFindRepeating(t *testing.T) {
	tests := []struct {
		Arr      []int
		Expected int
	}{
		{
			[]int{1, 3, 2, 4, 6, 5, 7, 3},
			3,
		},
	}

	for idx, test := range tests {
		t.Run(fmt.Sprintf("Test-Case-%v", idx), func(t *testing.T) {
			result := FindRepeating(test.Arr)
			if result != test.Expected {
				t.Errorf("Expected: %v Got: %v", test.Expected, result)
			}
		})
	}
}
