package array

// func main() {
// 	arr := []int{40, 40, 40}
// 	second := secondlargest(arr)
// 	if second >= 0 {
// 		fmt.Println("Second Largest: ", arr[second])
// 	} else {
// 		fmt.Println("Second Largest: Not found")
// 	}

// }

func secondlargest(arr []int) int {
	second := -1
	first := 0
	for i, val := range arr {
		if arr[first] < val {
			second, first = first, i
			continue
		}
		if arr[first] != val && (second == -1 || arr[second] < val) {
			second = i
		}
	}
	return second
}
