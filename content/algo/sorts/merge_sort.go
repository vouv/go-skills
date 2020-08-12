// +build OMIT

package main

import (
	"fmt"
)

// 归并排序
func MergeSort(arr []int)  {
	n := len(arr)
	if n < 2 { return }
	p := n/2 // 中间切分
	MergeSort(arr[:p])
	MergeSort(arr[p:])
	merge(arr[:p], arr[p:], arr)
}
// 合并a和b到result
func merge(a, b, result []int)  {
	na := len(a)
	nb := len(b)
	n := len(result) // assert na + nb = n

	tmp := make([]int, n) // 申请额外数组

	var i,j int
	for k := 0; k < n ; k++ {
		if j == nb || i < na && a[i] <= b[j] {
			tmp[k] = a[i]
			i++
		}else {
			tmp[k] = b[j]
			j++
		}
	}
	// 拷贝数据到原数组
	for i, v := range tmp {
		result[i] = v
	}
}

func main() {
	arr := []int{7,8,9,5,6,3,4,2,1}
	fmt.Println("before", arr)
	MergeSort(arr)
	fmt.Println("after", arr)
}
