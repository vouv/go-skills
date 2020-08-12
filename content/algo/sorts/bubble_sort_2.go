// +build OMIT

package main

import "fmt"

// 冒泡排序
func BubbleSort(arr []int) {
	n := len(arr)
	for 0 < bubble(arr[:n]) {
		n--
	}
}

func bubble(arr []int) int {
	var last = 0
	for j := 1; j < len(arr); j++ {
		if arr[j-1] > arr[j] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			last = j - 1
		}
	}
	return last
}

func main() {
	arr := []int{7, 8, 9, 5, 6, 3, 4, 2, 1}
	fmt.Println("before", arr)
	BubbleSort(arr)
	fmt.Println("after", arr)
}
