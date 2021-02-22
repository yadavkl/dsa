package array

// func main() {
// 	arr := []int{1, 5, 3, 8, 12}
// 	//fmt.Println("Max Profit: ", maxprofit(arr, 0, len(arr)-1))
// 	fmt.Println("Max Profit: ", maxprofit(arr))
//}

// func maxprofit(arr []int, start, end int) int {
// 	if end <= start {
// 		return 0
// 	}
// 	profit := 0
// 	for i := start; i < end; i++ {
// 		for j := i + 1; j <= end; j++ {
// 			if arr[i] < arr[j] {
// 				currProfit := arr[j] - arr[i] + maxprofit(arr, start, i-1) + maxprofit(arr, j+1, end)
// 				if currProfit > profit {
// 					profit = currProfit
// 				}
// 			}
// 		}
// 	}
// 	return profit
// }

func maxprofit(arr []int) int {
	profit := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			profit += arr[i] - arr[i-1]
		}
	}
	return profit
}
