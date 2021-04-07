package sort

//seggregate negative and posive number such that
//-ve numbers come before +ve numbers
func Sort2Type(arr []int) {
	//Can use Hoare partition algorithm
	i := 0
	j := len(arr) - 1

	for {
		for arr[i] < 0 {
			i++
		}
		for arr[j] > 0 {
			j--
		}
		if i >= j {
			return
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
}
