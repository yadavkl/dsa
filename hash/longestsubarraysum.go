package hash

func LongestSubarraySum(arr []int, sum int) int {
	presum := 0
	presummap := make(map[int]int)
	maxlen := 0
	for idx, val := range arr {
		presum = presum + val
		if presum == sum {
			return idx + 1
		}
		if _, ok := presummap[presum]; !ok {
			presummap[presum] = idx
		}
		if _, ok := presummap[presum-sum]; ok {
			maxlen = max(maxlen, idx-presummap[presum-sum])
		}
	}
	return maxlen
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
