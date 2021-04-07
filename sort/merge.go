package sort

// func MergeSort(arr []int, l, r int) {
// 	if l < r {
// 		m := l + (r-l)/2
// 		MergeSort(arr, l, m)
// 		MergeSort(arr, m+1, r)
// 		Merge(arr, l, m, r)
// 	}
// }

// func Merge(arr []int, l, m, r int) {

// 	left := make([]int, (m-l)+1)
// 	right := make([]int, r-m)

// 	for i := 0; i < len(left); i++ {
// 		left[i] = arr[l+i]
// 	}
// 	for i := 0; i < len(right); i++ {
// 		right[i] = arr[m+1+i]
// 	}
// 	ln := len(left)
// 	lm := len(right)

// 	var i, j int
// 	k := l
// 	for i < ln && j < lm {
// 		if left[i] <= right[j] {
// 			arr[k] = left[i]
// 			i++
// 			k++
// 		} else {
// 			arr[k] = right[j]
// 			j++
// 			k++
// 		}
// 	}
// 	for i < ln {
// 		arr[k] = left[i]
// 		i++
// 		k++
// 	}
// 	for j < lm {
// 		arr[k] = right[j]
// 		j++
// 		k++
// 	}
// }

func MergeSort(arr []int) []int {
	n := len(arr)
	if n == 1 {
		return arr
	}
	m := n / 2
	left := make([]int, m)
	right := make([]int, n-m)

	for i := 0; i < m; i++ {
		left[i] = arr[i]
	}
	for i := 0; i < n-m; i++ {
		right[i] = arr[m+i]
	}
	return Merge(MergeSort(left), MergeSort(right))
}

func Merge(left []int, right []int) (result []int) {
	n := len(left)
	m := len(right)
	result = make([]int, n+m)
	var i, j, k int
	for i < n && j < m {
		if left[i] < right[j] {
			result[k] = left[i]
			i++
			k++
		} else {
			result[k] = right[j]
			j++
			k++
		}
	}
	for i < n {
		result[k] = left[i]
		i++
		k++
	}
	for j < m {
		result[k] = right[j]
		j++
		k++
	}
	return
}
