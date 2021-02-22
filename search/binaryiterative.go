package search

func BinaryIterative(arr []int, ele int) int {
	l := 0
	r := len(arr)
	for l <= r {
		m := (l + r) / 2
		if arr[m] == ele {
			return m
		} else if arr[m] < ele {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}
