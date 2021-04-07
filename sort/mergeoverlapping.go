package sort

import "sort"

type Pair struct {
	start, end int
}

func MergeOverlapping(arr []Pair) []Pair {
	sort.Slice(arr, func(i, j int) bool { return arr[i].start < arr[j].start })
	idx := 0
	for i := 1; i < len(arr); i++ {
		if arr[idx].end >= arr[i].start {
			arr[idx].start = Min(arr[idx].start, arr[i].start)
			arr[idx].end = Max(arr[idx].end, arr[i].end)
		} else {
			idx++
			arr[idx] = arr[i]
		}
	}
	return arr[:idx+1]
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
