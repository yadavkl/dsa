package array

func findmajority(arr []int) int {
	//find element
	count := 0
	index := 0
	for idx, val := range arr {
		if val == arr[index] {
			count++
		} else {
			count--
		}
		if count == 0 {
			index = idx
			count = 1
		}
	}

	// check for majority element
	count = 0
	for _, val := range arr {
		if val == arr[index] {
			count++
		}
	}
	if count <= len(arr)/2 {
		return -1
	}
	return index
}
