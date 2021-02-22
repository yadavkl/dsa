package array

func evenoddsubarr(arr []int) int {
	count := 1
	maxcount := 1
	for i := 1; i < len(arr); i++ {
		if arr[i-1]%2 != arr[i]%2 {
			count++
			maxcount = max(count, maxcount)
			continue
		}
		count = 1
	}
	return maxcount
}
