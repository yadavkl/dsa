package array

import "testing"

func TestRemoveDuplicate(t *testing.T) {
	type Test struct {
		input          []int
		expectedresult []int
	}
	test := Test{
		input:          []int{20, 20, 20, 30, 40, 40, 40, 50, 50, 60},
		expectedresult: []int{20, 30, 40, 50, 60},
	}

	ret := removeduplicate(test.input)
	if !isEqual(ret, test.expectedresult) {
		t.Errorf("Expected: %v but actual: %v\n", test.expectedresult, ret)
	}
	// 	fmt.Println(ret)
}

func isEqual(sl1 []int, sl2 []int) bool {
	if len(sl1) != len(sl2) {
		return false
	}
	for i, val := range sl1 {
		if val != sl2[i] {
			return false
		}
	}
	return true
}
