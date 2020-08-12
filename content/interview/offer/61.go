// +build OMIT

package main

import (
	"fmt"
	"sort"
)

func isStraight(nums []int) bool {
	n := len(nums)
	if n < 5 {
		return false
	}

	sort.Ints(nums)

	var numZero = 0
	var numGap = 0

	for i := 0; i < n && nums[i] == 0; i++ {
		numZero++
	}

	// 统计两数差值是否大于1
	for j := numZero; j < n-1; j++ {
		if nums[j] == nums[j+1] {
			return false
		}
		numGap += nums[j+1] - nums[j] - 1
	}
	return numZero >= numGap
}

func main() {
	fmt.Println("expect true and get", isStraight([]int{1, 2, 3, 4, 5}))
	fmt.Println("expect true and get", isStraight([]int{0, 0, 1, 2, 5}))
	fmt.Println("expect false and get", isStraight([]int{0, 1, 3, 4, 6}))
}
