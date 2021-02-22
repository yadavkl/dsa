package array

// func main() {
// 	arr := []int{0, 3, 0, 0, 4, 5, 6, 7, 9, 0, 8, 10}
// 	move02end(arr)
// 	fmt.Println(arr)
// }
func move02end(arr []int) {
	nonzero := 0
	for i, val := range arr {
		if val != 0 {
			arr[nonzero], arr[i] = arr[i], arr[nonzero]
			nonzero++
		}
	}
}
