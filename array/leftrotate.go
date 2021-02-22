package array

// func main() {
// 	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	leftrotate(arr, 3)
// 	fmt.Println(arr)
// }

func leftrotate(arr []int, d int) {
	reverse(arr[:3])
	reverse(arr[3:])
	reverse(arr)
}

func reverse(arr []int) {
	start := 0
	end := len(arr) - 1
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}

}
