package sort

func DutchFlag(arr []int) {
	i := 0
	j := len(arr) - 1
	mid := 0
	for mid <= j {
		if arr[mid] == 0 {
			arr[mid], arr[i] = arr[i], arr[mid]
			i++
			mid++
		} else if arr[mid] == 1 {
			mid++
		} else {
			arr[mid], arr[j] = arr[j], arr[mid]
			j--
		}
	}
}
