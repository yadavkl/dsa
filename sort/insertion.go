package sort

func Insertion(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		j := i - 1
		key := arr[i]
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
