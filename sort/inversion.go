package sort

func Inversions(arr []int, l, r int) int {
	res := 0
	if l < r {
		m := l + (r-l)/2
		res += Inversions(arr, l, m)
		res += Inversions(arr, m+1, r)
		res += CountAndMerge(arr, l, m, r)
	}

	return res
}

func CountAndMerge(arr []int, l, m, r int) int {
	left := make([]int, m-l+1)
	right := make([]int, r-m)
	ll := len(left)
	rl := len(right)
	res := 0
	for i := 0; i < ll; i++ {
		left[i] = arr[l+i]
	}
	for i := 0; i < rl; i++ {
		right[i] = arr[m+1+i]
	}
	var i, j int
	k := l

	for i < ll && j < rl {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
			k++
		} else {
			arr[k] = right[j]
			j++
			k++
			res += ll - i
		}
	}
	for i < ll {
		arr[k] = left[i]
		i++
		k++
	}
	for j < rl {
		arr[k] = right[j]
		j++
		k++
	}
	return res
}
