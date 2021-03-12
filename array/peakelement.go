package array

func PeakElement(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	l := 0
	h := len(arr) - 1

	for l <= h {
		m := (l + h) / 2
		if (m == 0 || arr[m-1] <= arr[m]) && (m == len(arr)-1 || arr[m+1] <= arr[m]) {
			return arr[m]
		}
		if m > 0 && arr[m-1] >= arr[m] {
			h = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}
