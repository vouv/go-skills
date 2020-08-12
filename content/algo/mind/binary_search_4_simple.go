// +build OMIT

package main

import "fmt"

func BinarySearchLastLess(arr []int, target int) (index int) {
	n := len(arr)
	lo := 0
	hi := n - 1
	for lo <= hi {
		mid := lo + (hi-lo)>>1
		if arr[mid] <= target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	if hi > -1 && arr[hi] == target {
		return hi
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7}
	fmt.Println("expect 6 and get", BinarySearchLastLess(arr, 4))

}
