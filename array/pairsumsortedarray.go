package array

func PairSumSortedArray(arr []int, sum int) bool {
	l := 0
	r := len(arr) - 1

	for l < r {
		if arr[l]+arr[r] == sum {
			return true
		}
		if arr[l]+arr[r] < sum {
			l++
		} else {
			r--
		}
	}
	return false
}
