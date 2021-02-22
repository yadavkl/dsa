package array

import "fmt"

func circularsumsubarr(arr []int) int {
	//kadane ago maximum sum

	maxSubArraySum := func(arr []int) int {
		maxEnding := arr[0]
		maxSum := arr[0]
		n := len(arr)
		for i := 1; i < n; i++ {
			maxEnding = max(maxEnding+arr[i], arr[i])
			maxSum = max(maxSum, maxEnding)
		}
		return maxSum
	}

	mxsas := maxSubArraySum(arr)
	//If max if less than 0 means all negative number
	if mxsas < 0 {
		return mxsas
	}
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		arr[i] = -arr[i]
	}
	fmt.Println(arr)
	mnsas := sum + maxSubArraySum(arr)
	fmt.Println("Max: ", mxsas, "Min: ", mnsas)
	return max(mxsas, mnsas)
}
