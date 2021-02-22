package array

// func main() {
// 	arr := []int{20, 20, 20, 30, 40, 40, 40, 50, 50, 60}
// 	ret := removeduplicate(arr)
// 	fmt.Println(ret)
// }
func removeduplicate(arr []int) []int {
	idx := 1
	for _, val := range arr {
		if arr[idx-1] != val {
			arr[idx] = val
			idx++
		}
	}
	return arr[:idx]

}
