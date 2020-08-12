// +build OMIT

package main

import "fmt"

func search(nums []int, target int) int {
	li := bsl(nums, target)
	if li < 0 {
		return 0
	}
	ri := bsr(nums, target)
	return ri - li + 1 // 返回两个下标差值即可
}

// 查找第一个值为给定值的元素下标，不存在返回-1
func bsl(nums []int, target int) int {
	n := len(nums)
	lo := 0
	hi := n - 1
	for lo <= hi {
		mid := lo + (hi-lo)>>1
		if nums[mid] >= target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	if lo > n-1 || nums[lo] != target {
		return -1
	}
	return lo
}

// 查找最后一个值为给定值的下标，不存在返回-1
func bsr(nums []int, target int) int {
	n := len(nums)
	lo := 0
	hi := n - 1
	for lo <= hi {
		mid := lo + (hi-lo)>>1
		if nums[mid] <= target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	if hi < 0 || nums[hi] != target {
		return -1
	}
	return hi
}

func main() {
	fmt.Println("expect 2 and get", search([]int{5, 7, 7, 8, 8, 10}, 8))
}
