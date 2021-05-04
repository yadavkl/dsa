package array

import (
	"fmt"
	"testing"
)

func Test_SmallestWindow(t *testing.T) {
	tests := []struct {
		String   string
		Pattern  string
		Expected string
	}{
		{
			"ADOBECODEBANC",
			"ABC",
			"BANC",
		},
		{
			"this is a test string",
			"tist",
			"t stri",
		},
		{
			"geeksforgeeks",
			"ork",
			"ksfor",
		},
	}
	for idx, test := range tests {
		t.Run(fmt.Sprintf("Test-Case-%v", idx), func(t *testing.T) {
			result := SmallestWindow(test.String, test.Pattern)
			if result != test.Expected {
				t.Errorf("Expected: %v Got: %v", test.Expected, result)
			}
		})
	}
}
