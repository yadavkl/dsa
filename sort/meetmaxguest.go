package sort

import "sort"

func MeetMaxGuest(arrival []int, departure []int) int {
	i := 1
	j := 0
	res := 1
	cur := 1
	sort.Slice(arrival, func(i, j int) bool { return arrival[i] < arrival[j] })
	sort.Slice(departure, func(i, j int) bool { return departure[i] < departure[j] })
	for i < len(arrival) && j < len(departure) {
		if arrival[i] <= departure[j] {
			cur++
			i++
		} else {
			cur--
			j++
		}
		res = Max(res, cur)
	}
	return res
}
