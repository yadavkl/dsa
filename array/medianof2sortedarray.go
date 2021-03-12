package array

import "math"

func MedianOfSortedArray(arr1 []int, arr2 []int) float64 {
	if len(arr1) < len(arr2) {
		return GetMedian(arr1, arr2)
	}
	return GetMedian(arr2, arr1)
}

func GetMedian(a1 []int, a2 []int) float64 {
	l := 0
	r := len(a1)
	n1 := len(a1)
	n2 := len(a2)
	for l <= r {
		i1 := (l + r) / 2
		i2 := (n1+n2+1)/2 - i1
		min1 := math.MinInt32
		if i1 < n1 {
			min1 = a1[i1]
		}
		max1 := math.MaxInt32
		if i1 > 0 {
			max1 = a1[i1-1]
		}
		min2 := math.MinInt32
		if i2 < n2 {
			min2 = a2[i2]
		}
		max2 := math.MaxInt32
		if i2 > 0 {
			max2 = a2[i2-1]
		}
		if max1 <= min2 && min1 >= max2 {
			if (n1+n2)%2 == 0 {
				return (float64(Min(min1, min2)) + float64(Max(max1, max2))) / 2
			}
			return float64(Max(max1, max2))
		}
		if max1 > min2 {
			r = i1 - 1
		} else {
			l = i1 + 1
		}
	}
	return -1
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
