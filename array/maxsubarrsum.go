package array

// func main() {
// 	arr := []int{-3, 8, -2, 4, -5, 6}
// 	fmt.Println(maxsubarrsum(arr))
// }

func maxsubarrsum(arr []int) int {
	n := len(arr)

	maxending := arr[0]
	maxval := arr[0]
	for i := 1; i < n; i++ {
		maxending = max(maxending+arr[i], arr[i])
		maxval = max(maxval, maxending)
	}
	return maxval
}
