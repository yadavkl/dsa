package array

// If found x return location
// not found return -1

func RotatedSortedArraySearch(arr []int, x int) int {
	l := 0
	r := len(arr) - 1

	for l <= r {
		m := (l + r) / 2
		if arr[m] == x {
			return m
		}
		//Find if left half sorted ?
		// if middle element less than left corner element
		// then left array sorted
		if arr[l] <= arr[m] {
			// element in between left corner and middle then do binary search in
			// sorted array
			// other wise to to right side of middle
			if x >= arr[l] && x < arr[m] {
				r = m - 1
			} else {
				l = m + 1
			}
		} else { //right half sorted
			if x > arr[m] && x <= arr[r] {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	return -1
}
