// +build OMIT

package main

import "fmt"

func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)

	dp[0] = nums[0]
	max := dp[0]
	for i := 1; i < n; i++ {
		dp[i] = getMax(nums[i], dp[i-1]+nums[i])
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	fmt.Println("expect 6 and get", maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
