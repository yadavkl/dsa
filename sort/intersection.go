package sort

// Intersection of two sorted array
//efficient solution O(n+m)

func Intersection(arr1 []int, arr2 []int) (result []int) {
	n := len(arr1)
	m := len(arr2)

	var i, j int
	for i < n && j < m {
		if i > 0 && arr1[i] == arr1[i-1] {
			i++
			continue
		}
		if arr1[i] < arr2[j] {
			i++
		}
		if arr1[i] > arr2[j] {
			j++
		}
		if arr1[i] == arr2[j] {
			result = append(result, arr1[i])
			i++
			j++
		}
	}
	return
}
