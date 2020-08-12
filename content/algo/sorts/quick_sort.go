// +build OMIT

package main

import (
	"fmt"
)

// 快速排序
func QuickSort(arr []int)  {
	n := len(arr)
	if n < 2 { return }
	p := partition(arr) // 获取分区点
	QuickSort(arr[:p])
	QuickSort(arr[p+1:])
}
// 分区并返回分区点，pos左侧为小于pivot的区间，pos右侧为大于pivot的区间
func partition(arr []int) int {
	n := len(arr)
	pivot := arr[n-1]
	var pos = 0 // 定义分区点pos，
	for i := 0; i < n ; i ++ {
		if arr[i] < pivot {
			arr[pos], arr[i] = arr[i], arr[pos]
			pos++
		}
	}
	arr[pos], arr[n-1] = arr[n-1], arr[pos] // 最后把pivot放到分区点pos
	return pos
}

func main() {
	arr := []int{7,8,9,5,6,3,4,2,1}
	fmt.Println("before", arr)
	QuickSort(arr)
	fmt.Println("after", arr)
}
