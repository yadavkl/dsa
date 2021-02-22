package search

func BinaryRecursive(arr []int, key, l, r int) int {
	if l > r {
		return -1
	}
	m := (l + r) / 2
	if arr[m] == key {
		return m
	}
	if arr[m] > key {
		return BinaryRecursive(arr, key, l, m-1)
	}
	return BinaryRecursive(arr, key, m+1, r)
}
