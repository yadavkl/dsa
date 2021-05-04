package hash

func DistinctElementInWindow(arr []int, k int) (result []int) {

	left := 0
	right := 0
	//<move right to k length
	//create counter map
	countmap := make(map[int]int)
	counter := 0
	for right = 0; right < k; right++ {
		if _, ok := countmap[arr[right]]; !ok {
			counter++
		}
		countmap[arr[right]] += 1
	}
	result = append(result, counter)
	for right = k; right < len(arr); right++ {
		if _, ok := countmap[arr[right]]; !ok {
			counter++
		}
		countmap[arr[right]] += 1
		countmap[arr[left]] -= 1
		if countmap[arr[left]] == 0 {
			counter--
		}
		left++
		result = append(result, counter)
	}
	return
}
