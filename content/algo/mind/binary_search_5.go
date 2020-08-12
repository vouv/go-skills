package main

import "fmt"

func BinarySearchMin(arr []int) int {
	n := len(arr)
	lo := 0
	hi := n - 1
	for lo <= hi {
		mid := lo + (hi-lo)>>1
		if arr[0] < arr[mid] {
			lo = mid + 1
		} else {
			if arr[mid-1] > arr[mid] {
				return mid
			}
			hi = mid - 1
		}
	}
	return 0
}

func main() {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println("expect 4 and get", BinarySearchMin(arr))
}
