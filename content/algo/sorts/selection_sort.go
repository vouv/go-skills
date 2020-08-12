// +build OMIT

package main

import (
	"fmt"
)

// 选择排序
func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n - 1 ; i++ {
		min := i
		for j := i+1 ; j < n ; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}

func main() {
	arr := []int{7,8,9,5,6,3,4,2,1}
	fmt.Println("before", arr)
	SelectionSort(arr)
	fmt.Println("after", arr)
}
