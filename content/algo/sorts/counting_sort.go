// +build OMIT

package main

import "fmt"
// 计数排序
func CountingSort(arr[]int)  {
	n := len(arr)
	if n < 2 { return }

	// 查找数组中数据的范围
	max := arr[0]
	for i := 1; i < n ; i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	// 根据最大范围申请一个数组来计数
	c := make([]int, max+1)

	// 计算每个元素的个数，放入c中
	for i := 0; i < n; i++ {
		c[arr[i]]++
	}

	// 依次累加
	for i := 1; i <= max; i++ {
		c[i] = c[i-1] + c[i]
	}

	// 临时数组r，存储排序之后的结果
	r := make([]int, n)
	// 计算排序的关键步骤，有点难理解
	for i := n - 1; i >= 0; i-- {
		index := c[arr[i]]-1
		r[index] = arr[i]
		c[arr[i]]--
	}

	// 将结果拷贝给a数组
	for i := 0; i < n; i++ {
		arr[i] = r[i]
	}
}

func main() {
	arr := []int{7,8,9,5,6,3,4,2,1}
	fmt.Println("before", arr)
	CountingSort(arr)
	fmt.Println("after", arr)
}
