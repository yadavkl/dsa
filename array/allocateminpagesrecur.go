package array

import "math"

func AllocateMinPages(arr []int, n int, k int) int {
	if k == 1 {
		return sum(arr, 0, n-1)
	}
	if n == 1 {
		return arr[0]
	}
	res := math.MaxInt32
	for i := 0; i < n; i++ {
		res = min(res, max(AllocateMinPages(arr, i, k-1), sum(arr, i, n-1)))
	}
	return res
}

/*
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
*/
func sum(arr []int, k, e int) int {
	sum := 0
	for i := k; i <= e; i++ {
		sum += arr[i]
	}
	return sum
}
