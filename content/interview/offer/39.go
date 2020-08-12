// +build OMIT

package main

import "fmt"

func majorityElement(nums []int) int {
	var res int
	var count int
	for _, v := range nums {
		if count == 0 {
			res = v
		}
		if v != res {
			count--
		} else {
			count++
		}
	}
	return res
}

func main() {
	fmt.Println(majorityElement([]int{1, 2, 3, 2, 2, 2, 5, 4, 2}))
}
