package array

import "fmt"

func minimumflip(arr []int) string {
	var result string
	n := len(arr)
	for i := 1; i < n; i++ {
		if arr[i] != arr[i-1] {
			if arr[i] != arr[0] {
				result += fmt.Sprintf("From %d", i)
			} else {
				result += fmt.Sprintf(" to %d ", i-1)
			}
		}
	}
	if arr[n-1] != arr[0] {
		result += fmt.Sprintf(" to %d", n-1)
	}
	return result
}
