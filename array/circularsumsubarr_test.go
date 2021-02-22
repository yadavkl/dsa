package array

import (
	"fmt"
	"testing"
)

func TestCircularSumSubArr(t *testing.T) {
	testCases := []struct {
		arr      []int
		expected int
	}{
		{
			arr:      []int{8, -4, 3, -5, 4},
			expected: 12,
		},
		{
			arr:      []int{5, -2, 3, 4},
			expected: 12,
		},
		{
			arr:      []int{3, -4, 5, 6, -8, 7},
			expected: 17,
		},
	}

	for idx, tc := range testCases {
		arr := tc.arr
		expected := tc.expected
		t.Run(fmt.Sprintf("Test-Case-%v", idx), func(t *testing.T) {
			res := circularsumsubarr(arr)
			if res != expected {
				t.Errorf("Expected: %v but got: %v", expected, res)
			}
		})
	}
}
