// +build OMIT

package main

import (
	"fmt"
)

// 插入排序
func InsertionSort(arr []int)  {
	n := len(arr)
	for i := 1; i < n; i++ {
		for j := i ; j > 0 && arr[j] < arr[j-1] ; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

func main() {
	arr := []int{7,8,9,5,6,3,4,2,1}
	fmt.Println("before", arr)
	InsertionSort(arr)
	fmt.Println("after", arr)
}
