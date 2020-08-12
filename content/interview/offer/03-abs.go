// +build OMIT

package main

import "fmt"

func findRepeatNumber(nums []int) int {
	var zero = false
	for _, v := range nums {
		av := abs(v)
		if av == 0 {
			if zero {
				return 0
			} else {
				zero = true
			}
		}
		if nums[av] < 0 {
			return av
		} else {
			nums[av] = -nums[av]
		}
	}
	return -1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println("expect 2 and get",
		findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3}))
}
