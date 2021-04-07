package sort

func CycleSort(arr []int) {
	n := len(arr)
	for cs := 0; cs < n; cs++ {
		pos := cs
		item := arr[cs]
		for i := cs + 1; i < n; i++ {
			if arr[i] < item {
				pos++
			}
		}
		item, arr[pos] = arr[pos], item
		for pos != cs {
			pos = cs
			for j := cs + 1; j < n; j++ {
				if arr[j] < item {
					pos++
				}
			}
			item, arr[pos] = arr[pos], item
		}
	}
}
