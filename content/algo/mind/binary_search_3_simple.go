// +build OMIT

package main

import "fmt"

func BinarySearchFirstGreater(arr []int, target int) (index int) {
	n := len(arr)
	lo := 0
	hi := n - 1
	for lo <= hi {
		mid := lo + (hi-lo)>>1
		if arr[mid] >= target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	if lo < n && arr[lo] == target {
		return lo
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7}
	fmt.Println("expect 3 and get", BinarySearchFirstGreater(arr, 4))

}
