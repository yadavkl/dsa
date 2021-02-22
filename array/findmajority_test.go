package array

import (
	"fmt"
	"testing"
)

func TestFindMajority(t *testing.T) {
	testCases := []struct {
		arr      []int
		expected int
	}{
		{
			arr:      []int{8, 8, 6, 6, 6, 4, 6},
			expected: 3,
		},
		{
			arr:      []int{6, 8, 4, 8, 8},
			expected: 3,
		},
	}
	for idx, tc := range testCases {
		arr := tc.arr
		expected := tc.expected
		t.Run(fmt.Sprintf("Test-Case-%v", idx), func(t *testing.T) {
			result := findmajority(arr)
			if result != expected {
				t.Errorf("Expected: %v Result: %v", expected, result)
			}
		})
	}

}
