package array

//window problem positive numbers 
//[]int{1,2,3,40,34,12}
// sum = arr[0]
// 

func subarrsumgivennum(arr []int, num int) bool {
	s := 0
	sum := arr[0]
	n := len(arr)
	for e := 1; e < n; e++ {
		for sum > num && s < e-1 {
			sum -= arr[s]
			s++
		}
		if sum == num {
			return true
		}
		sum += arr[e]
	}
	return false
}
