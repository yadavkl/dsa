package sort

func KthSmallest(arr []int, k int) int {
	l := 0
	r := len(arr) - 1

	for l <= r {
		p := partition(arr, l, r)
		if p == k-1 {
			return arr[p]
		}
		if p > k-1 {
			r = p - 1
		} else {
			l = p + 1
		}
	}
	return -1
}

//Lomuto partition
// keeps pivot on correct location
// this is used to get kth element
func partition(arr []int, l, r int) int {
	i := l - 1
	pivot := arr[r]

	for j := l; j < r; j++ {
		if pivot > arr[j] {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[r] = arr[r], arr[i+1]
	return i + 1
}
