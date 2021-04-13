package hash

func LargestSubarrayWithEqual01(arr []int) int {
	presum := 0
	premap := make(map[int]int)
	maxlen := 0
	for idx, val := range arr {
		if val == 0 {
			presum -= 1
		} else {
			presum += 1
		}
		if presum == 0 {
			maxlen = idx + 1
		}
		if _, ok := premap[presum]; !ok {
			premap[presum] = idx
		} else {
			maxlen = max(maxlen, idx-premap[presum])
		}
	}
	return maxlen
}
