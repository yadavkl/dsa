package sort

//Union of two sorted array
func Union(arr1 []int, arr2 []int) (result []int) {
	n := len(arr1)
	m := len(arr2)
	var i, j int

	for i < n && j < m {
		if i > 0 && arr1[i] == arr1[i-1] {
			i++
			continue
		}
		if j > 0 && arr2[j] == arr2[j-1] {
			j++
			continue
		}
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else if arr1[i] > arr2[j] {
			result = append(result, arr2[j])
			j++
		} else if arr1[i] == arr2[j] {
			result = append(result, arr1[i])
			i++
			j++
		}
	}
	for i > 0 && i < n && arr1[i] != arr1[i-1] {
		result = append(result, arr1[i])
		i++
	}
	for j > 0 && j < m && arr2[j] != arr2[j-1] {
		result = append(result, arr2[j])
		j++
	}
	return
}
