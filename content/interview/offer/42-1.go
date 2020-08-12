// +build OMIT

package main

import "fmt"

func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	max := nums[0]
	count := 0
	for i := 0; i < n; i++ {
		count += nums[i]
		if count < nums[i] {
			count = nums[i]
		}
		if count > max {
			max = count
		}
	}

	return max
}

func main() {
	fmt.Println("expect 6 and get", maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
