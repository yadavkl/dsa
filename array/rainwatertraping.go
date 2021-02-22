package array

// func main() {
// 	arr := []int{5, 0, 6, 2, 3}
// 	fmt.Println(watertrap(arr))
// }

func watertrap(arr []int) int {
	n := len(arr)
	lmax := make([]int, n)
	rmax := make([]int, n)

	lmax[0] = arr[0]

	for i := 1; i < n-1; i++ {
		lmax[i] = max(lmax[i-1], arr[i])
	}
	rmax[n-1] = arr[n-1]
	for i := n - 2; i >= 0; i-- {
		rmax[i] = max(rmax[i+1], arr[i])
	}
	ret := 0
	for i := 1; i < n-1; i++ {
		ret = ret + min(lmax[i], rmax[i]) - arr[i]
	}
	return ret
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
