package sort

func Selection(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minidx := i
		for j := i + 1; j < n; j++ {
			if arr[minidx] > arr[j] {
				minidx = j
			}
		}
		arr[minidx], arr[i] = arr[i], arr[minidx]
	}
}
