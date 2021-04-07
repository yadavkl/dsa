package sort

func QuickLomuto(arr []int, l, r int) {
	if l < r {
		p := LomutoPartition(arr, l, r)
		QuickLomuto(arr, l, p-1)
		QuickLomuto(arr, p+1, r)
	}
}

func LomutoPartition(arr []int, l, r int) int {
	pivote := arr[r]
	i := l - 1
	for j := l; j < r; j++ {
		if pivote > arr[j] {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[r] = arr[r], arr[i+1]
	return i + 1
}
