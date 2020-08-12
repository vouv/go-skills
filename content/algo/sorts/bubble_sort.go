// +build OMIT

package main

import "fmt"

// 冒泡排序
func BubbleSort(arr []int)  {
	n := len(arr)
	for i := 0; i< n; i++ {
		// 循环提前退出标志
		flag := true
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = false
			}
		}
		if flag {return} // 如果一趟扫描没有发生交换，则表明数组已经有序
	}
}

func main() {
	arr := []int{7,8,9,5,6,3,4,2,1}
	fmt.Println("before", arr)
	BubbleSort(arr)
	fmt.Println("after", arr)
}
