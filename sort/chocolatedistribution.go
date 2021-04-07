package sort

import (
	"fmt"
	"sort"
)

//MinDiff problem
//Each one should get atleast on chocolate packet
//and diffrence between min and max will be min

func ChocolateDistribution(arr []int, m int) int {
	if m > len(arr) {
		return -1
	}

	sort.Slice(arr, func(i, j int) bool { return arr[i] <= arr[j] })
	fmt.Println(arr)
	n := len(arr)
	res := arr[m-1] - arr[0]

	for i := 1; (m + i - 1) < n; i++ {
		res = Min(res, arr[m+i-1]-arr[i])
	}
	return res
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
