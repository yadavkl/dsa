package array

import (
	"math"
)

// func main() {
// 	arr := []int{2, 3, 10, 6, 4, 8, 1}
// 	fmt.Println(maxdiff(arr))
// }

func maxdiff(arr []int) int {
	diff := 0
	minval := arr[0]
	for _, val := range arr {
		diff = int(math.Max(float64(val-minval), float64(diff)))
		minval = int(math.Min(float64(val), float64(minval)))
	}
	return diff
}
