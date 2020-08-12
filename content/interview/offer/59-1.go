// +build OMIT

package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nums
	}

	var res = []int{}
	var buf = []int{} // 保存元素下标的队列

	for i := 0; i < len(nums); i++ {
		if len(buf) > 0 && i-buf[0] >= k {
			buf = buf[1:]
		}
		for len(buf) > 0 && nums[i] > nums[buf[len(buf)-1]] {
			buf = buf[:len(buf)-1]
		}
		buf = append(buf, i)
		if i >= k-1 {
			res = append(res, nums[buf[0]])
		}
	}
	return res
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println("expect [3 3 5 5 6 7] and get", maxSlidingWindow(nums, k))
}
