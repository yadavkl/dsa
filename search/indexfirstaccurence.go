package search

func IndexFirstAccurance(arr []int, l, h, x int) int {
	if l > h {
		return -1
	}
	m := (l + h) / 2
	if arr[m] == x {
		if m == 0 || arr[m] != arr[m-1] {
			return m
		}
		return IndexFirstAccurance(arr, l, m-1, x)
	}
	if arr[m] > x {
		return IndexFirstAccurance(arr, l, m-1, x)
	}
	return IndexFirstAccurance(arr, m+1, h, x)
}

func IndexFirstAccuranceItr(arr []int, x int) int {
	l := 0
	h := len(arr) - 1
	for l <= h {
		m := (l + h) / 2
		if x == arr[m] {
			if m == 0 || arr[m-1] != arr[m] {
				return m
			}
			h = m - 1
		}
		if x < arr[m] {
			h = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}
