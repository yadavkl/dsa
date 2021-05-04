package array

import "fmt"

//repeating element in a array
//O(N)
//o(1) space
//dont change array

func FindRepeating(arr []int) int {
	slow := arr[0]
	fast := arr[0]

	slow = arr[slow]
	fast = arr[arr[fast]]
	//flag := false
	for slow != fast {
		//flag = true
		slow = arr[slow]
		fast = arr[arr[fast]]
		fmt.Println("Phase1-->Slow:", slow, " Fast:", fast)
	}
	slow = arr[0]
	for slow != fast {
		slow = arr[slow]
		fast = arr[fast]
	}
	return slow
}

/*
func main() {
	arr := []int{1, 3, 2, 4, 6, 5, 7, 3}
	result := FindRepeating(arr)
	fmt.Println("Result: ", result)
}
*/
