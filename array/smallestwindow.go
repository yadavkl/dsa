package array

import (
	"fmt"
	"math"
)

/*
 Not considering repeated text in pattern
 If consider repeated test in pattern we need to check
 if map is matched instead leangth
*/
func SmallestWindow(s string, t string) string {
	str := []byte(s)
	tst := []byte(t)
	tmap := make(map[byte]bool)
	for _, val := range tst {
		tmap[val] = true
	}
	type pair struct {
		val byte
		idx int
	}
	var arr []pair
	for idx, val := range str {
		if _, ok := tmap[val]; ok {
			arr = append(arr, pair{val, idx})
		}
	}
	fmt.Println(arr)
	//Now min window
	left := 0
	right := 0
	required := len(tmap)
	minwindow := math.MaxInt32
	var startIdx, endIdx int
	kmap := make(map[byte]int)

	for right < len(arr) {
		kmap[arr[right].val] += 1
		fmt.Println("Kmap Len: ", len(kmap))
		for required == len(kmap) && left < right {
			fmt.Println("Kmap Len: ", len(kmap), "Required: ", required)
			length := arr[right].idx - arr[left].idx

			if minwindow > length {
				minwindow = length
				startIdx = arr[left].idx
				endIdx = arr[right].idx
			}
			kmap[arr[left].val] -= 1
			if kmap[arr[left].val] == 0 {
				delete(kmap, arr[left].val)
			}
			left++
		}
		right++
	}
	return string(str[startIdx : endIdx+1])
}
