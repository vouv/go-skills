// +build OMIT

package main

import "fmt"

func findRepeatNumber(nums []int) int {
	for i, v := range nums {
		for i != v {
			if v == nums[v] {
				return v
			}
			// 把v（nums[i]）放到i的位置
			nums[i], nums[v] = nums[v], nums[i]
			v = nums[i]
		}
	}
	return -1
}

func main() {
	fmt.Println("expect 11 and get",
		findRepeatNumber([]int{0, 1, 2, 3, 4, 11, 6, 7, 8, 9, 10, 12, 11, 13, 14, 15}))
}
