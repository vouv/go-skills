// +build OMIT

package main

import "fmt"

func exchange(nums []int) []int {
	// p的左边（不含p）是奇数
	p := 0
	for i := range nums {
		if nums[i]&1 == 1 {
			nums[p], nums[i] = nums[i], nums[p]
			p++
		}
	}
	return nums
}

func main() {
	fmt.Println(exchange([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
}
