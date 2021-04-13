package hash

func SubArraySumZero(arr []int) bool {
	hash := make(map[int]bool)
	prefixsum := 0
	for _, val := range arr {
		prefixsum += val
		if _, ok := hash[prefixsum]; ok {
			return true
		}
		if prefixsum == 0 {
			return true
		}
		hash[prefixsum] = true
	}
	return false
}
