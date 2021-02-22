package search

import (
	"fmt"
	"testing"
)

func TestBinaryRecursive(t *testing.T) {
	testCases := []struct {
		arr      []int
		key      int
		expected int
	}{
		{
			[]int{10, 20, 30, 40, 50, 60},
			20,
			1,
		},
		{
			[]int{5, 15, 25},
			25,
			2,
		},
		{
			[]int{5, 10, 15, 25, 35},
			20,
			-1,
		},
		{
			[]int{10, 10},
			10,
			1,
		},
	}
	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Test-Case-%v", idx), func(t *testing.T) {
			result := BinaryRecursive(tc.arr, tc.key, 0, len(tc.arr))
			if result != tc.expected {
				t.Errorf("Expected: %v Got: %v\n", tc.expected, result)
			}
		})
	}
}
