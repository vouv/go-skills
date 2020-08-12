// +build OMIT

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	n := len(nums)

	lo := 0
	hi := n - 1

	for lo < hi {
		sum := nums[lo] + nums[hi]
		if sum == target {
			return []int{nums[lo], nums[hi]}
		} else if sum > target {
			hi--
		} else {
			lo++
		}
	}

	return []int{}
}

func main() {
	arr := []int{2, 7, 11, 15}
	fmt.Println("expect [2, 7] and get", twoSum(arr, 9))
}
