package search

import (
	"fmt"
	"testing"
)

func TestFirstIndexAccurance(t *testing.T) {
	tests := []struct {
		arr      []int
		key      int
		expected int
	}{
		{
			[]int{5, 10, 10, 15, 20, 20, 20},
			20,
			4,
		},
	}
	for idx, test := range tests {
		t.Run(fmt.Sprintf("Test-Case-%v", idx), func(t *testing.T) {
			result := IndexFirstAccurance(test.arr, 0, len(test.arr)-1, test.key)
			if result != test.expected {
				t.Errorf("expected: %d got: %d\n", test.expected, result)
			}
		})

	}
}

func TestFirstIndexAccuranceItr(t *testing.T) {
	tests := []struct {
		arr      []int
		key      int
		expected int
	}{
		{
			[]int{5, 10, 10, 15, 20, 20, 20},
			20,
			4,
		},
	}
	for idx, test := range tests {
		t.Run(fmt.Sprintf("Test-Case-%v", idx), func(t *testing.T) {
			result := IndexFirstAccuranceItr(test.arr, test.key)
			if result != test.expected {
				t.Errorf("expected: %d got: %d\n", test.expected, result)
			}
		})

	}
}
