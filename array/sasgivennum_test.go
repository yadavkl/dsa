package array

import (
	"fmt"
	"testing"
)

func TestSASGivenNum(t *testing.T) {
	testCases := []struct {
		arr      []int
		num      int
		expected bool
	}{
		{
			[]int{1, 4, 20, 3, 10, 5},
			33,
			true,
		},
	}
	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Test-Case-%v", idx), func(t *testing.T) {
			result := subarrsumgivennum(tc.arr, tc.num)
			if result != tc.expected {
				t.Errorf("Expected: %v Got: %v", tc.expected, result)
			}
		})
	}
}
