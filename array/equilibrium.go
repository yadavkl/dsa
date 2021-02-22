package array

func equilibrium(arr []int) bool {
	prefix_sum := make([]int, len(arr))
	n := len(arr)
	prefix_sum[0] = arr[0]
	for i := 1; i < n; i++ {
		prefix_sum[i] = prefix_sum[i-1] + arr[i]
	}
	if 0 == prefix_sum[n-1]-arr[0] || 0 == prefix_sum[n-2] {
		return true
	}
	for i := 1; i < n-1; i++ {
		if prefix_sum[i] == prefix_sum[n-1]-prefix_sum[i+1] {
			return true
		}
	}
	return false
}
