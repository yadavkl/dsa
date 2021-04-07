package sort

func QuickHoare(arr []int, l, r int) {
	if l < r {
		p := HoarePartition(arr, l, r)
		QuickHoare(arr, l, p)
		QuickHoare(arr, p+1, r)
	}
}

func HoarePartition(arr []int, l, r int) int {
	i := l
	j := r
	pivote := arr[l]
	for {
		for arr[i] < pivote {
			i++
		}
		for arr[j] > pivote {
			j--
		}
		if i >= j {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	return j
}
