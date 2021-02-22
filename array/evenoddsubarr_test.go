package array

import "testing"

func TestEvenOddSubArr(t *testing.T) {
	arrs := [][]int{[]int{10, 12, 14, 7, 8}, []int{5, 10, 20, 6, 3, 8}}
	expectedresult := []int{3, 3}
	for i, arr := range arrs {
		result := evenoddsubarr(arr)
		if expectedresult[i] != result {
			t.Errorf("Erro: result : %d != expectedresult : %d", result, expectedresult[i])
		}
	}

}
