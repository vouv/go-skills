package main

import "fmt"

func BinarySearchFirst(arr []int, target int) (index int) {
	n := len(arr)
	lo := 0
	hi := n - 1
	for lo < hi {
		mid := lo + (hi-lo)>>1
		if arr[mid] >= target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	if hi == n || arr[hi] != target {
		return -1
	}
	return hi
}

func main() {
	arr := []int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7}
	fmt.Println("expect 3 and get", BinarySearchFirst(arr, 4))
}
