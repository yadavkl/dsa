package hash


func SubArraySum(arr []int, sum int)bool{
	hash := make(map[int]bool)
	subsum := 0
	for _, val := range arr{
		subsum += val
		if _, ok := hash[subsum-sum]; ok{
			return true
		}
		if subsum == sum{
			return true
		}
		hash[subsum] = true
	}
	return false
}